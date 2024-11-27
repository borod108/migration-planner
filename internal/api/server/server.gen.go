// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	. "github.com/kubev2v/migration-planner/api/v1alpha1"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/agents)
	ListAgents(w http.ResponseWriter, r *http.Request)

	// (DELETE /api/v1/agents/{id})
	DeleteAgent(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (DELETE /api/v1/sources)
	DeleteSources(w http.ResponseWriter, r *http.Request)

	// (GET /api/v1/sources)
	ListSources(w http.ResponseWriter, r *http.Request)

	// (DELETE /api/v1/sources/{id})
	DeleteSource(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (GET /api/v1/sources/{id})
	ReadSource(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (GET /api/v1/sources/{id}/image)
	GetSourceImage(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (HEAD /api/v1/sources/{id}/image)
	HeadSourceImage(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)

	// (GET /health)
	Health(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /api/v1/agents)
func (_ Unimplemented) ListAgents(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /api/v1/agents/{id})
func (_ Unimplemented) DeleteAgent(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /api/v1/sources)
func (_ Unimplemented) DeleteSources(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /api/v1/sources)
func (_ Unimplemented) ListSources(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /api/v1/sources/{id})
func (_ Unimplemented) DeleteSource(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /api/v1/sources/{id})
func (_ Unimplemented) ReadSource(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /api/v1/sources/{id}/image)
func (_ Unimplemented) GetSourceImage(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (HEAD /api/v1/sources/{id}/image)
func (_ Unimplemented) HeadSourceImage(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /health)
func (_ Unimplemented) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListAgents operation middleware
func (siw *ServerInterfaceWrapper) ListAgents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListAgents(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteAgent operation middleware
func (siw *ServerInterfaceWrapper) DeleteAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteAgent(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteSources operation middleware
func (siw *ServerInterfaceWrapper) DeleteSources(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteSources(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListSources operation middleware
func (siw *ServerInterfaceWrapper) ListSources(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListSources(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteSource operation middleware
func (siw *ServerInterfaceWrapper) DeleteSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteSource(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ReadSource operation middleware
func (siw *ServerInterfaceWrapper) ReadSource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ReadSource(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSourceImage operation middleware
func (siw *ServerInterfaceWrapper) GetSourceImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSourceImage(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// HeadSourceImage operation middleware
func (siw *ServerInterfaceWrapper) HeadSourceImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HeadSourceImage(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Health operation middleware
func (siw *ServerInterfaceWrapper) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Health(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/agents", wrapper.ListAgents)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/agents/{id}", wrapper.DeleteAgent)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/sources", wrapper.DeleteSources)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/sources", wrapper.ListSources)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/v1/sources/{id}", wrapper.DeleteSource)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/sources/{id}", wrapper.ReadSource)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/v1/sources/{id}/image", wrapper.GetSourceImage)
	})
	r.Group(func(r chi.Router) {
		r.Head(options.BaseURL+"/api/v1/sources/{id}/image", wrapper.HeadSourceImage)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/health", wrapper.Health)
	})

	return r
}

type ListAgentsRequestObject struct {
}

type ListAgentsResponseObject interface {
	VisitListAgentsResponse(w http.ResponseWriter) error
}

type ListAgents200JSONResponse AgentList

func (response ListAgents200JSONResponse) VisitListAgentsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListAgents401JSONResponse Error

func (response ListAgents401JSONResponse) VisitListAgentsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type ListAgents500JSONResponse Error

func (response ListAgents500JSONResponse) VisitListAgentsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteAgentRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type DeleteAgentResponseObject interface {
	VisitDeleteAgentResponse(w http.ResponseWriter) error
}

type DeleteAgent200JSONResponse Agent

func (response DeleteAgent200JSONResponse) VisitDeleteAgentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteAgent400JSONResponse Error

func (response DeleteAgent400JSONResponse) VisitDeleteAgentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteAgent401JSONResponse Error

func (response DeleteAgent401JSONResponse) VisitDeleteAgentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type DeleteAgent404JSONResponse Error

func (response DeleteAgent404JSONResponse) VisitDeleteAgentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteAgent500JSONResponse Error

func (response DeleteAgent500JSONResponse) VisitDeleteAgentResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSourcesRequestObject struct {
}

type DeleteSourcesResponseObject interface {
	VisitDeleteSourcesResponse(w http.ResponseWriter) error
}

type DeleteSources200JSONResponse Status

func (response DeleteSources200JSONResponse) VisitDeleteSourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSources401JSONResponse Error

func (response DeleteSources401JSONResponse) VisitDeleteSourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type ListSourcesRequestObject struct {
}

type ListSourcesResponseObject interface {
	VisitListSourcesResponse(w http.ResponseWriter) error
}

type ListSources200JSONResponse SourceList

func (response ListSources200JSONResponse) VisitListSourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListSources401JSONResponse Error

func (response ListSources401JSONResponse) VisitListSourcesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSourceRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type DeleteSourceResponseObject interface {
	VisitDeleteSourceResponse(w http.ResponseWriter) error
}

type DeleteSource200JSONResponse Source

func (response DeleteSource200JSONResponse) VisitDeleteSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSource400JSONResponse Error

func (response DeleteSource400JSONResponse) VisitDeleteSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSource401JSONResponse Error

func (response DeleteSource401JSONResponse) VisitDeleteSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type DeleteSource404JSONResponse Error

func (response DeleteSource404JSONResponse) VisitDeleteSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type ReadSourceRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type ReadSourceResponseObject interface {
	VisitReadSourceResponse(w http.ResponseWriter) error
}

type ReadSource200JSONResponse Source

func (response ReadSource200JSONResponse) VisitReadSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ReadSource400JSONResponse Error

func (response ReadSource400JSONResponse) VisitReadSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ReadSource401JSONResponse Error

func (response ReadSource401JSONResponse) VisitReadSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type ReadSource404JSONResponse Error

func (response ReadSource404JSONResponse) VisitReadSourceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetSourceImageRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type GetSourceImageResponseObject interface {
	VisitGetSourceImageResponse(w http.ResponseWriter) error
}

type GetSourceImage200ApplicationoctetStreamResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response GetSourceImage200ApplicationoctetStreamResponse) VisitGetSourceImageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type GetSourceImage400JSONResponse Error

func (response GetSourceImage400JSONResponse) VisitGetSourceImageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetSourceImage401JSONResponse Error

func (response GetSourceImage401JSONResponse) VisitGetSourceImageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetSourceImage404JSONResponse Error

func (response GetSourceImage404JSONResponse) VisitGetSourceImageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetSourceImage500JSONResponse Error

func (response GetSourceImage500JSONResponse) VisitGetSourceImageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type HeadSourceImageRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type HeadSourceImageResponseObject interface {
	VisitHeadSourceImageResponse(w http.ResponseWriter) error
}

type HeadSourceImage200Response struct {
}

func (response HeadSourceImage200Response) VisitHeadSourceImageResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type HeadSourceImage400Response struct {
}

func (response HeadSourceImage400Response) VisitHeadSourceImageResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type HeadSourceImage401Response struct {
}

func (response HeadSourceImage401Response) VisitHeadSourceImageResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type HeadSourceImage404Response struct {
}

func (response HeadSourceImage404Response) VisitHeadSourceImageResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type HeadSourceImage500Response struct {
}

func (response HeadSourceImage500Response) VisitHeadSourceImageResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

type HealthRequestObject struct {
}

type HealthResponseObject interface {
	VisitHealthResponse(w http.ResponseWriter) error
}

type Health200Response struct {
}

func (response Health200Response) VisitHealthResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /api/v1/agents)
	ListAgents(ctx context.Context, request ListAgentsRequestObject) (ListAgentsResponseObject, error)

	// (DELETE /api/v1/agents/{id})
	DeleteAgent(ctx context.Context, request DeleteAgentRequestObject) (DeleteAgentResponseObject, error)

	// (DELETE /api/v1/sources)
	DeleteSources(ctx context.Context, request DeleteSourcesRequestObject) (DeleteSourcesResponseObject, error)

	// (GET /api/v1/sources)
	ListSources(ctx context.Context, request ListSourcesRequestObject) (ListSourcesResponseObject, error)

	// (DELETE /api/v1/sources/{id})
	DeleteSource(ctx context.Context, request DeleteSourceRequestObject) (DeleteSourceResponseObject, error)

	// (GET /api/v1/sources/{id})
	ReadSource(ctx context.Context, request ReadSourceRequestObject) (ReadSourceResponseObject, error)

	// (GET /api/v1/sources/{id}/image)
	GetSourceImage(ctx context.Context, request GetSourceImageRequestObject) (GetSourceImageResponseObject, error)

	// (HEAD /api/v1/sources/{id}/image)
	HeadSourceImage(ctx context.Context, request HeadSourceImageRequestObject) (HeadSourceImageResponseObject, error)

	// (GET /health)
	Health(ctx context.Context, request HealthRequestObject) (HealthResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// ListAgents operation middleware
func (sh *strictHandler) ListAgents(w http.ResponseWriter, r *http.Request) {
	var request ListAgentsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ListAgents(ctx, request.(ListAgentsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListAgents")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ListAgentsResponseObject); ok {
		if err := validResponse.VisitListAgentsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteAgent operation middleware
func (sh *strictHandler) DeleteAgent(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request DeleteAgentRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteAgent(ctx, request.(DeleteAgentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteAgent")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteAgentResponseObject); ok {
		if err := validResponse.VisitDeleteAgentResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteSources operation middleware
func (sh *strictHandler) DeleteSources(w http.ResponseWriter, r *http.Request) {
	var request DeleteSourcesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteSources(ctx, request.(DeleteSourcesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteSources")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteSourcesResponseObject); ok {
		if err := validResponse.VisitDeleteSourcesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ListSources operation middleware
func (sh *strictHandler) ListSources(w http.ResponseWriter, r *http.Request) {
	var request ListSourcesRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ListSources(ctx, request.(ListSourcesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListSources")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ListSourcesResponseObject); ok {
		if err := validResponse.VisitListSourcesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteSource operation middleware
func (sh *strictHandler) DeleteSource(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request DeleteSourceRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteSource(ctx, request.(DeleteSourceRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteSource")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteSourceResponseObject); ok {
		if err := validResponse.VisitDeleteSourceResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ReadSource operation middleware
func (sh *strictHandler) ReadSource(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request ReadSourceRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ReadSource(ctx, request.(ReadSourceRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ReadSource")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ReadSourceResponseObject); ok {
		if err := validResponse.VisitReadSourceResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetSourceImage operation middleware
func (sh *strictHandler) GetSourceImage(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request GetSourceImageRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetSourceImage(ctx, request.(GetSourceImageRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetSourceImage")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetSourceImageResponseObject); ok {
		if err := validResponse.VisitGetSourceImageResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// HeadSourceImage operation middleware
func (sh *strictHandler) HeadSourceImage(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var request HeadSourceImageRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.HeadSourceImage(ctx, request.(HeadSourceImageRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "HeadSourceImage")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HeadSourceImageResponseObject); ok {
		if err := validResponse.VisitHeadSourceImageResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Health operation middleware
func (sh *strictHandler) Health(w http.ResponseWriter, r *http.Request) {
	var request HealthRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Health(ctx, request.(HealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Health")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HealthResponseObject); ok {
		if err := validResponse.VisitHealthResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
