// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for GetWithContentTypeParamsContentType.
const (
	Json GetWithContentTypeParamsContentType = "json"
	Text GetWithContentTypeParamsContentType = "text"
)

// EveryTypeOptional defines model for EveryTypeOptional.
type EveryTypeOptional struct {
	ArrayInlineField     *[]int              `json:"array_inline_field,omitempty"`
	ArrayReferencedField *[]SomeObject       `json:"array_referenced_field,omitempty"`
	BoolField            *bool               `json:"bool_field,omitempty"`
	ByteField            *[]byte             `json:"byte_field,omitempty"`
	DateField            *openapi_types.Date `json:"date_field,omitempty"`
	DateTimeField        *time.Time          `json:"date_time_field,omitempty"`
	DoubleField          *float64            `json:"double_field,omitempty"`
	FloatField           *float32            `json:"float_field,omitempty"`
	InlineObjectField    *struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"inline_object_field,omitempty"`
	Int32Field      *int32      `json:"int32_field,omitempty"`
	Int64Field      *int64      `json:"int64_field,omitempty"`
	IntField        *int        `json:"int_field,omitempty"`
	NumberField     *float32    `json:"number_field,omitempty"`
	ReferencedField *SomeObject `json:"referenced_field,omitempty"`
	StringField     *string     `json:"string_field,omitempty"`
}

// EveryTypeRequired defines model for EveryTypeRequired.
type EveryTypeRequired struct {
	ArrayInlineField     []int                `json:"array_inline_field"`
	ArrayReferencedField []SomeObject         `json:"array_referenced_field"`
	BoolField            bool                 `json:"bool_field"`
	ByteField            []byte               `json:"byte_field"`
	DateField            openapi_types.Date   `json:"date_field"`
	DateTimeField        time.Time            `json:"date_time_field"`
	DoubleField          float64              `json:"double_field"`
	EmailField           *openapi_types.Email `json:"email_field,omitempty"`
	FloatField           float32              `json:"float_field"`
	InlineObjectField    struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"inline_object_field"`
	Int32Field      int32      `json:"int32_field"`
	Int64Field      int64      `json:"int64_field"`
	IntField        int        `json:"int_field"`
	NumberField     float32    `json:"number_field"`
	ReferencedField SomeObject `json:"referenced_field"`
	StringField     string     `json:"string_field"`
}

// ReservedKeyword defines model for ReservedKeyword.
type ReservedKeyword struct {
	Channel *string `json:"channel,omitempty"`
}

// Resource defines model for Resource.
type Resource struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

// SomeObject defines model for some_object.
type SomeObject struct {
	Name string `json:"name"`
}

// Argument defines model for argument.
type Argument = string

// ResponseWithReference defines model for ResponseWithReference.
type ResponseWithReference = SomeObject

// SimpleResponse defines model for SimpleResponse.
type SimpleResponse struct {
	Name string `json:"name"`
}

// GetWithArgsParams defines parameters for GetWithArgs.
type GetWithArgsParams struct {
	// OptionalArgument An optional query argument
	OptionalArgument *int64 `form:"optional_argument,omitempty" json:"optional_argument,omitempty"`

	// RequiredArgument An optional query argument
	RequiredArgument int64 `form:"required_argument" json:"required_argument"`

	// HeaderArgument An optional query argument
	HeaderArgument *int32 `json:"header_argument,omitempty"`
}

// GetWithContentTypeParamsContentType defines parameters for GetWithContentType.
type GetWithContentTypeParamsContentType string

// CreateResource2Params defines parameters for CreateResource2.
type CreateResource2Params struct {
	// InlineQueryArgument Some query argument
	InlineQueryArgument *int `form:"inline_query_argument,omitempty" json:"inline_query_argument,omitempty"`
}

// UpdateResource3JSONBody defines parameters for UpdateResource3.
type UpdateResource3JSONBody struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateResourceJSONRequestBody defines body for CreateResource for application/json ContentType.
type CreateResourceJSONRequestBody = EveryTypeRequired

// CreateResource2JSONRequestBody defines body for CreateResource2 for application/json ContentType.
type CreateResource2JSONRequestBody = Resource

// UpdateResource3JSONRequestBody defines body for UpdateResource3 for application/json ContentType.
type UpdateResource3JSONRequestBody UpdateResource3JSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get every type optional
	// (GET /every-type-optional)
	GetEveryTypeOptional(w http.ResponseWriter, r *http.Request)
	// Get resource via simple path
	// (GET /get-simple)
	GetSimple(w http.ResponseWriter, r *http.Request)
	// Getter with referenced parameter and referenced response
	// (GET /get-with-args)
	GetWithArgs(w http.ResponseWriter, r *http.Request, params GetWithArgsParams)
	// Getter with referenced parameter and referenced response
	// (GET /get-with-references/{global_argument}/{argument})
	GetWithReferences(w http.ResponseWriter, r *http.Request, globalArgument int64, argument Argument)
	// Get an object by ID
	// (GET /get-with-type/{content_type})
	GetWithContentType(w http.ResponseWriter, r *http.Request, contentType GetWithContentTypeParamsContentType)
	// get with reserved keyword
	// (GET /reserved-keyword)
	GetReservedKeyword(w http.ResponseWriter, r *http.Request)
	// Create a resource
	// (POST /resource/{argument})
	CreateResource(w http.ResponseWriter, r *http.Request, argument Argument)
	// Create a resource with inline parameter
	// (POST /resource2/{inline_argument})
	CreateResource2(w http.ResponseWriter, r *http.Request, inlineArgument int, params CreateResource2Params)
	// Update a resource with inline body. The parameter name is a reserved
	// keyword, so make sure that gets prefixed to avoid syntax errors
	// (PUT /resource3/{fallthrough})
	UpdateResource3(w http.ResponseWriter, r *http.Request, pFallthrough int)
	// get response with reference
	// (GET /response-with-reference)
	GetResponseWithReference(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// get every type optional
// (GET /every-type-optional)
func (_ Unimplemented) GetEveryTypeOptional(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get resource via simple path
// (GET /get-simple)
func (_ Unimplemented) GetSimple(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Getter with referenced parameter and referenced response
// (GET /get-with-args)
func (_ Unimplemented) GetWithArgs(w http.ResponseWriter, r *http.Request, params GetWithArgsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Getter with referenced parameter and referenced response
// (GET /get-with-references/{global_argument}/{argument})
func (_ Unimplemented) GetWithReferences(w http.ResponseWriter, r *http.Request, globalArgument int64, argument Argument) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get an object by ID
// (GET /get-with-type/{content_type})
func (_ Unimplemented) GetWithContentType(w http.ResponseWriter, r *http.Request, contentType GetWithContentTypeParamsContentType) {
	w.WriteHeader(http.StatusNotImplemented)
}

// get with reserved keyword
// (GET /reserved-keyword)
func (_ Unimplemented) GetReservedKeyword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a resource
// (POST /resource/{argument})
func (_ Unimplemented) CreateResource(w http.ResponseWriter, r *http.Request, argument Argument) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a resource with inline parameter
// (POST /resource2/{inline_argument})
func (_ Unimplemented) CreateResource2(w http.ResponseWriter, r *http.Request, inlineArgument int, params CreateResource2Params) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a resource with inline body. The parameter name is a reserved
// keyword, so make sure that gets prefixed to avoid syntax errors
// (PUT /resource3/{fallthrough})
func (_ Unimplemented) UpdateResource3(w http.ResponseWriter, r *http.Request, pFallthrough int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// get response with reference
// (GET /response-with-reference)
func (_ Unimplemented) GetResponseWithReference(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetEveryTypeOptional operation middleware
func (siw *ServerInterfaceWrapper) GetEveryTypeOptional(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetEveryTypeOptional(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSimple operation middleware
func (siw *ServerInterfaceWrapper) GetSimple(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSimple(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetWithArgs operation middleware
func (siw *ServerInterfaceWrapper) GetWithArgs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetWithArgsParams

	// ------------- Optional query parameter "optional_argument" -------------

	err = runtime.BindQueryParameter("form", true, false, "optional_argument", r.URL.Query(), &params.OptionalArgument)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "optional_argument", Err: err})
		return
	}

	// ------------- Required query parameter "required_argument" -------------

	if paramValue := r.URL.Query().Get("required_argument"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "required_argument"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "required_argument", r.URL.Query(), &params.RequiredArgument)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "required_argument", Err: err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "header_argument" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header_argument")]; found {
		var HeaderArgument int32
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "header_argument", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "header_argument", valueList[0], &HeaderArgument, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "header_argument", Err: err})
			return
		}

		params.HeaderArgument = &HeaderArgument

	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithArgs(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetWithReferences operation middleware
func (siw *ServerInterfaceWrapper) GetWithReferences(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "global_argument" -------------
	var globalArgument int64

	err = runtime.BindStyledParameterWithOptions("simple", "global_argument", chi.URLParam(r, "global_argument"), &globalArgument, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "global_argument", Err: err})
		return
	}

	// ------------- Path parameter "argument" -------------
	var argument Argument

	err = runtime.BindStyledParameterWithOptions("simple", "argument", chi.URLParam(r, "argument"), &argument, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "argument", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithReferences(w, r, globalArgument, argument)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetWithContentType operation middleware
func (siw *ServerInterfaceWrapper) GetWithContentType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "content_type" -------------
	var contentType GetWithContentTypeParamsContentType

	err = runtime.BindStyledParameterWithOptions("simple", "content_type", chi.URLParam(r, "content_type"), &contentType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "content_type", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithContentType(w, r, contentType)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetReservedKeyword operation middleware
func (siw *ServerInterfaceWrapper) GetReservedKeyword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetReservedKeyword(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateResource operation middleware
func (siw *ServerInterfaceWrapper) CreateResource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "argument" -------------
	var argument Argument

	err = runtime.BindStyledParameterWithOptions("simple", "argument", chi.URLParam(r, "argument"), &argument, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "argument", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateResource(w, r, argument)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateResource2 operation middleware
func (siw *ServerInterfaceWrapper) CreateResource2(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "inline_argument" -------------
	var inlineArgument int

	err = runtime.BindStyledParameterWithOptions("simple", "inline_argument", chi.URLParam(r, "inline_argument"), &inlineArgument, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inline_argument", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateResource2Params

	// ------------- Optional query parameter "inline_query_argument" -------------

	err = runtime.BindQueryParameter("form", true, false, "inline_query_argument", r.URL.Query(), &params.InlineQueryArgument)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "inline_query_argument", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateResource2(w, r, inlineArgument, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateResource3 operation middleware
func (siw *ServerInterfaceWrapper) UpdateResource3(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "fallthrough" -------------
	var pFallthrough int

	err = runtime.BindStyledParameterWithOptions("simple", "fallthrough", chi.URLParam(r, "fallthrough"), &pFallthrough, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "fallthrough", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateResource3(w, r, pFallthrough)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetResponseWithReference operation middleware
func (siw *ServerInterfaceWrapper) GetResponseWithReference(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetResponseWithReference(w, r)
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
		r.Get(options.BaseURL+"/every-type-optional", wrapper.GetEveryTypeOptional)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/get-simple", wrapper.GetSimple)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/get-with-args", wrapper.GetWithArgs)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/get-with-references/{global_argument}/{argument}", wrapper.GetWithReferences)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/get-with-type/{content_type}", wrapper.GetWithContentType)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/reserved-keyword", wrapper.GetReservedKeyword)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/resource/{argument}", wrapper.CreateResource)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/resource2/{inline_argument}", wrapper.CreateResource2)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/resource3/{fallthrough}", wrapper.UpdateResource3)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/response-with-reference", wrapper.GetResponseWithReference)
	})

	return r
}
