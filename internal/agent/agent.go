package agent

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	api "github.com/kubev2v/migration-planner/api/v1alpha1"
	"github.com/kubev2v/migration-planner/internal/agent/client"
	"github.com/kubev2v/migration-planner/pkg/log"
	"github.com/lthibault/jitterbug"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

const (
	// name of the file which stores the current inventory
	InventoryFile    = "inventory.json"
	defaultAgentPort = 3333
)

// This varible is set during build time.
// It contains the version of the code.
// For more info take a look into Makefile.
var version string

// New creates a new agent.
func New(id uuid.UUID, log *log.PrefixLogger, config *Config) *Agent {
	return &Agent{
		config:           config,
		log:              log,
		healtCheckStopCh: make(chan chan any),
		id:               id,
	}
}

type Agent struct {
	config           *Config
	log              *log.PrefixLogger
	server           *Server
	healtCheckStopCh chan chan any
	credUrl          string
	id               uuid.UUID
}

func (a *Agent) GetLogPrefix() string {
	return a.log.Prefix()
}

func (a *Agent) Run(ctx context.Context) error {
	var err error
	a.log.Infof("Starting agent: %s", version)
	defer a.log.Infof("Agent stopped")
	a.log.Infof("Configuration: %s", a.config.String())

	defer utilruntime.HandleCrash()

	client, err := newPlannerClient(a.config)
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(ctx)
	a.start(ctx, client)

	select {
	case <-sig:
	case <-ctx.Done():
	}

	a.log.Info("stopping agent...")

	a.Stop()
	cancel()

	return nil
}

func (a *Agent) Stop() {
	serverCh := make(chan any)
	a.server.Stop(serverCh)

	<-serverCh
	a.log.Info("server stopped")

	c := make(chan any)
	a.healtCheckStopCh <- c
	<-c
	a.log.Info("health check stopped")
}

func (a *Agent) start(ctx context.Context, plannerClient client.Planner) {
	inventoryUpdater := NewInventoryUpdater(a.log, a.id, plannerClient)
	statusUpdater := NewStatusUpdater(a.log, a.id, version, a.credUrl, a.config, plannerClient)

	// start server
	a.server = NewServer(defaultAgentPort, a.config.DataDir, a.config.WwwDir)
	go a.server.Start(a.log, statusUpdater)

	// get the credentials url
	a.initializeCredentialUrl()

	// start the health check
	healthChecker, err := NewHealthChecker(
		a.log,
		plannerClient,
		a.config.DataDir,
		time.Duration(a.config.HealthCheckInterval*int64(time.Second)),
	)
	if err != nil {
		a.log.Fatalf("failed to start health check: %w", err)
	}

	// TODO refactor health checker to call it from the main goroutine
	healthChecker.Start(a.healtCheckStopCh)

	collector := NewCollector(a.log, a.config.DataDir)
	collector.collect(ctx)

	updateTicker := jitterbug.New(time.Duration(a.config.UpdateInterval.Duration), &jitterbug.Norm{Stdev: 30 * time.Millisecond, Mean: 0})

	/*
		Main loop
		The status of agent is always computed even if we don't have connectivity with the backend.
		If we're connected to the backend, the agent sends its status and if the status is UpToDate,
		it sends the inventory.
		In case of "source gone", it stops everything and break from the loop.
	*/
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-updateTicker.C:
			}

			// calculate status regardless if we have connectivity withe the backend.
			status, statusInfo, inventory := statusUpdater.CalculateStatus()

			//	check for health. Send requests only if we have connectivity
			if healthChecker.State() == HealthCheckStateConsoleUnreachable {
				continue
			}

			if err := statusUpdater.UpdateStatus(ctx, status, statusInfo); err != nil {
				if errors.Is(err, client.ErrSourceGone) {
					a.log.Info("Source is gone..Stop sending requests")
					// stop the server and the healthchecker
					a.Stop()
					break
				}
				a.log.Errorf("unable to update agent status: %s", err)
				continue // skip inventory update if we cannot update agent's state.
			}

			if status == api.AgentStatusUpToDate {
				inventoryUpdater.UpdateServiceWithInventory(ctx, api.SourceStatusUpToDate, "Inventory collected with success", inventory)
			}
		}
	}()
}

func (a *Agent) initializeCredentialUrl() {
	// Parse the service URL
	parsedURL, err := url.Parse(a.config.PlannerService.Service.Server)
	if err != nil {
		a.log.Errorf("error parsing service URL: %v", err)
		a.credUrl = "N/A"
		return
	}

	// Use either port if specified, or scheme
	port := parsedURL.Port()
	if port == "" {
		port = parsedURL.Scheme
	}

	// Connect to service
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", parsedURL.Hostname(), port))
	if err != nil {
		a.log.Errorf("failed connecting to migration planner: %v", err)
		a.credUrl = "N/A"
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.TCPAddr)
	a.credUrl = fmt.Sprintf("http://%s:%d", localAddr.IP.String(), defaultAgentPort)
	a.log.Infof("Discovered Agent IP address: %s", a.credUrl)
}

func newPlannerClient(cfg *Config) (client.Planner, error) {
	httpClient, err := client.NewFromConfig(&cfg.PlannerService.Config)
	if err != nil {
		return nil, err
	}
	return client.NewPlanner(httpClient), nil
}
