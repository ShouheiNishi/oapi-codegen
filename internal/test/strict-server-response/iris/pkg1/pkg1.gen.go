// Package pkg1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package pkg1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/kataras/iris/v12"
	strictiris "github.com/oapi-codegen/runtime/strictmiddleware/iris"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// TestHeaderFixedNoContent request
	TestHeaderFixedNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TestHeaderNoContent request
	TestHeaderNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TestRefHeaderFixedNoContent request
	TestRefHeaderFixedNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// TestRefHeaderNoContent request
	TestRefHeaderNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) TestHeaderFixedNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTestHeaderFixedNoContentRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TestHeaderNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTestHeaderNoContentRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TestRefHeaderFixedNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTestRefHeaderFixedNoContentRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) TestRefHeaderNoContent(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewTestRefHeaderNoContentRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewTestHeaderFixedNoContentRequest generates requests for TestHeaderFixedNoContent
func NewTestHeaderFixedNoContentRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test-header-fixed-nocontent")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTestHeaderNoContentRequest generates requests for TestHeaderNoContent
func NewTestHeaderNoContentRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test-header-nocontent")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTestRefHeaderFixedNoContentRequest generates requests for TestRefHeaderFixedNoContent
func NewTestRefHeaderFixedNoContentRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test-ref-header-fixed-nocontent")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewTestRefHeaderNoContentRequest generates requests for TestRefHeaderNoContent
func NewTestRefHeaderNoContentRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/test-ref-header-nocontent")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// TestHeaderFixedNoContentWithResponse request
	TestHeaderFixedNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestHeaderFixedNoContentResponse, error)

	// TestHeaderNoContentWithResponse request
	TestHeaderNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestHeaderNoContentResponse, error)

	// TestRefHeaderFixedNoContentWithResponse request
	TestRefHeaderFixedNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestRefHeaderFixedNoContentResponse, error)

	// TestRefHeaderNoContentWithResponse request
	TestRefHeaderNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestRefHeaderNoContentResponse, error)
}

type TestHeaderFixedNoContentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r TestHeaderFixedNoContentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TestHeaderFixedNoContentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TestHeaderNoContentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r TestHeaderNoContentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TestHeaderNoContentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TestRefHeaderFixedNoContentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r TestRefHeaderFixedNoContentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TestRefHeaderFixedNoContentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type TestRefHeaderNoContentResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r TestRefHeaderNoContentResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r TestRefHeaderNoContentResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// TestHeaderFixedNoContentWithResponse request returning *TestHeaderFixedNoContentResponse
func (c *ClientWithResponses) TestHeaderFixedNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestHeaderFixedNoContentResponse, error) {
	rsp, err := c.TestHeaderFixedNoContent(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTestHeaderFixedNoContentResponse(rsp)
}

// TestHeaderNoContentWithResponse request returning *TestHeaderNoContentResponse
func (c *ClientWithResponses) TestHeaderNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestHeaderNoContentResponse, error) {
	rsp, err := c.TestHeaderNoContent(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTestHeaderNoContentResponse(rsp)
}

// TestRefHeaderFixedNoContentWithResponse request returning *TestRefHeaderFixedNoContentResponse
func (c *ClientWithResponses) TestRefHeaderFixedNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestRefHeaderFixedNoContentResponse, error) {
	rsp, err := c.TestRefHeaderFixedNoContent(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTestRefHeaderFixedNoContentResponse(rsp)
}

// TestRefHeaderNoContentWithResponse request returning *TestRefHeaderNoContentResponse
func (c *ClientWithResponses) TestRefHeaderNoContentWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*TestRefHeaderNoContentResponse, error) {
	rsp, err := c.TestRefHeaderNoContent(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseTestRefHeaderNoContentResponse(rsp)
}

// ParseTestHeaderFixedNoContentResponse parses an HTTP response from a TestHeaderFixedNoContentWithResponse call
func ParseTestHeaderFixedNoContentResponse(rsp *http.Response) (*TestHeaderFixedNoContentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TestHeaderFixedNoContentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseTestHeaderNoContentResponse parses an HTTP response from a TestHeaderNoContentWithResponse call
func ParseTestHeaderNoContentResponse(rsp *http.Response) (*TestHeaderNoContentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TestHeaderNoContentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseTestRefHeaderFixedNoContentResponse parses an HTTP response from a TestRefHeaderFixedNoContentWithResponse call
func ParseTestRefHeaderFixedNoContentResponse(rsp *http.Response) (*TestRefHeaderFixedNoContentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TestRefHeaderFixedNoContentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseTestRefHeaderNoContentResponse parses an HTTP response from a TestRefHeaderNoContentWithResponse call
func ParseTestRefHeaderNoContentResponse(rsp *http.Response) (*TestRefHeaderNoContentResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &TestRefHeaderNoContentResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /test-header-fixed-nocontent)
	TestHeaderFixedNoContent(ctx iris.Context)

	// (GET /test-header-nocontent)
	TestHeaderNoContent(ctx iris.Context)

	// (GET /test-ref-header-fixed-nocontent)
	TestRefHeaderFixedNoContent(ctx iris.Context)

	// (GET /test-ref-header-nocontent)
	TestRefHeaderNoContent(ctx iris.Context)
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc iris.Handler

// TestHeaderFixedNoContent converts iris context to params.
func (w *ServerInterfaceWrapper) TestHeaderFixedNoContent(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.TestHeaderFixedNoContent(ctx)
}

// TestHeaderNoContent converts iris context to params.
func (w *ServerInterfaceWrapper) TestHeaderNoContent(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.TestHeaderNoContent(ctx)
}

// TestRefHeaderFixedNoContent converts iris context to params.
func (w *ServerInterfaceWrapper) TestRefHeaderFixedNoContent(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.TestRefHeaderFixedNoContent(ctx)
}

// TestRefHeaderNoContent converts iris context to params.
func (w *ServerInterfaceWrapper) TestRefHeaderNoContent(ctx iris.Context) {

	// Invoke the callback with all the unmarshaled arguments
	w.Handler.TestRefHeaderNoContent(ctx)
}

// IrisServerOption is the option for iris server
type IrisServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *iris.Application, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, IrisServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *iris.Application, si ServerInterface, options IrisServerOptions) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.Get(options.BaseURL+"/test-header-fixed-nocontent", wrapper.TestHeaderFixedNoContent)
	router.Get(options.BaseURL+"/test-header-nocontent", wrapper.TestHeaderNoContent)
	router.Get(options.BaseURL+"/test-ref-header-fixed-nocontent", wrapper.TestRefHeaderFixedNoContent)
	router.Get(options.BaseURL+"/test-ref-header-nocontent", wrapper.TestRefHeaderNoContent)

	router.Build()
}

type TestRespRefHeaderFixedNoContentResponseHeaders struct {
	Header1 string
	Header2 int
}
type TestRespRefHeaderFixedNoContentResponse struct {
	Headers TestRespRefHeaderFixedNoContentResponseHeaders
}

type TestRespRefHeaderNoContentResponseHeaders struct {
	Header1 string
	Header2 int
}
type TestRespRefHeaderNoContentResponse struct {
	Headers TestRespRefHeaderNoContentResponseHeaders
}

type TestHeaderFixedNoContentRequestObject struct {
}

type TestHeaderFixedNoContentResponseObject interface {
	VisitTestHeaderFixedNoContentResponse(ctx iris.Context) error
}

type TestHeaderFixedNoContent204ResponseHeaders struct {
	Header1 string
	Header2 int
}

type TestHeaderFixedNoContent204Response struct {
	Headers TestHeaderFixedNoContent204ResponseHeaders
}

func (response TestHeaderFixedNoContent204Response) VisitTestHeaderFixedNoContentResponse(ctx iris.Context) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.StatusCode(204)
	return nil
}

type TestHeaderNoContentRequestObject struct {
}

type TestHeaderNoContentResponseObject interface {
	VisitTestHeaderNoContentResponse(ctx iris.Context) error
}

type TestHeaderNoContentdefaultResponseHeaders struct {
	Header1 string
	Header2 int
}

type TestHeaderNoContentdefaultResponse struct {
	Headers TestHeaderNoContentdefaultResponseHeaders

	StatusCode int
}

func (response TestHeaderNoContentdefaultResponse) VisitTestHeaderNoContentResponse(ctx iris.Context) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.StatusCode(response.StatusCode)
	return nil
}

type TestRefHeaderFixedNoContentRequestObject struct {
}

type TestRefHeaderFixedNoContentResponseObject interface {
	VisitTestRefHeaderFixedNoContentResponse(ctx iris.Context) error
}

type TestRefHeaderFixedNoContent204Response = TestRespRefHeaderFixedNoContentResponse

func (response TestRefHeaderFixedNoContent204Response) VisitTestRefHeaderFixedNoContentResponse(ctx iris.Context) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.StatusCode(204)
	return nil
}

type TestRefHeaderNoContentRequestObject struct {
}

type TestRefHeaderNoContentResponseObject interface {
	VisitTestRefHeaderNoContentResponse(ctx iris.Context) error
}

type TestRefHeaderNoContentdefaultResponse struct {
	Headers TestRespRefHeaderNoContentResponseHeaders

	StatusCode int
}

func (response TestRefHeaderNoContentdefaultResponse) VisitTestRefHeaderNoContentResponse(ctx iris.Context) error {
	ctx.Response().Header.Set("header1", fmt.Sprint(response.Headers.Header1))
	ctx.Response().Header.Set("header2", fmt.Sprint(response.Headers.Header2))
	ctx.StatusCode(response.StatusCode)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /test-header-fixed-nocontent)
	TestHeaderFixedNoContent(ctx context.Context, request TestHeaderFixedNoContentRequestObject) (TestHeaderFixedNoContentResponseObject, error)

	// (GET /test-header-nocontent)
	TestHeaderNoContent(ctx context.Context, request TestHeaderNoContentRequestObject) (TestHeaderNoContentResponseObject, error)

	// (GET /test-ref-header-fixed-nocontent)
	TestRefHeaderFixedNoContent(ctx context.Context, request TestRefHeaderFixedNoContentRequestObject) (TestRefHeaderFixedNoContentResponseObject, error)

	// (GET /test-ref-header-nocontent)
	TestRefHeaderNoContent(ctx context.Context, request TestRefHeaderNoContentRequestObject) (TestRefHeaderNoContentResponseObject, error)
}

type StrictHandlerFunc = strictiris.StrictIrisHandlerFunc
type StrictMiddlewareFunc = strictiris.StrictIrisMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// TestHeaderFixedNoContent operation middleware
func (sh *strictHandler) TestHeaderFixedNoContent(ctx iris.Context) {
	var request TestHeaderFixedNoContentRequestObject

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TestHeaderFixedNoContent(ctx, request.(TestHeaderFixedNoContentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TestHeaderFixedNoContent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(TestHeaderFixedNoContentResponseObject); ok {
		if err := validResponse.VisitTestHeaderFixedNoContentResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// TestHeaderNoContent operation middleware
func (sh *strictHandler) TestHeaderNoContent(ctx iris.Context) {
	var request TestHeaderNoContentRequestObject

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TestHeaderNoContent(ctx, request.(TestHeaderNoContentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TestHeaderNoContent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(TestHeaderNoContentResponseObject); ok {
		if err := validResponse.VisitTestHeaderNoContentResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// TestRefHeaderFixedNoContent operation middleware
func (sh *strictHandler) TestRefHeaderFixedNoContent(ctx iris.Context) {
	var request TestRefHeaderFixedNoContentRequestObject

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TestRefHeaderFixedNoContent(ctx, request.(TestRefHeaderFixedNoContentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TestRefHeaderFixedNoContent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(TestRefHeaderFixedNoContentResponseObject); ok {
		if err := validResponse.VisitTestRefHeaderFixedNoContentResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}

// TestRefHeaderNoContent operation middleware
func (sh *strictHandler) TestRefHeaderNoContent(ctx iris.Context) {
	var request TestRefHeaderNoContentRequestObject

	handler := func(ctx iris.Context, request interface{}) (interface{}, error) {
		return sh.ssi.TestRefHeaderNoContent(ctx, request.(TestRefHeaderNoContentRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "TestRefHeaderNoContent")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, err)
		return
	} else if validResponse, ok := response.(TestRefHeaderNoContentResponseObject); ok {
		if err := validResponse.VisitTestRefHeaderNoContentResponse(ctx); err != nil {
			ctx.StopWithError(http.StatusBadRequest, err)
			return
		}
	} else if response != nil {
		ctx.Writef("Unexpected response type: %T", response)
		return
	}
}
