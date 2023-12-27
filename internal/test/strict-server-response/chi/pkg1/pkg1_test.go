// Code generated by generator/generate.go DO NOT EDIT.

package pkg1_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deepmap/oapi-codegen/v2/internal/test/strict-server-response/chi/pkg1"
)

type strictServerInterface struct{}

func (s strictServerInterface) TestHeaderNoContent(ctx context.Context, request pkg1.TestHeaderNoContentRequestObject) (pkg1.TestHeaderNoContentResponseObject, error) {
	return pkg1.TestHeaderNoContentdefaultResponse{
		Headers: pkg1.TestHeaderNoContentdefaultResponseHeaders{
			Header1: "foo",
			Header2: 123,
		},
		StatusCode: 204,
	}, nil
}

func TestHeaderNoContent(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-header-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestHeaderNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "foo", res.HTTPResponse.Header.Get("header1"))
	assert.Equal(t, "123", res.HTTPResponse.Header.Get("header2"))
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}

func (s strictServerInterface) TestHeaderFixedNoContent(ctx context.Context, request pkg1.TestHeaderFixedNoContentRequestObject) (pkg1.TestHeaderFixedNoContentResponseObject, error) {
	return pkg1.TestHeaderFixedNoContent204Response{
		Headers: pkg1.TestHeaderFixedNoContent204ResponseHeaders{
			Header1: "foo",
			Header2: 123,
		},
	}, nil
}

func TestHeaderFixedNoContent(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-header-fixed-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestHeaderFixedNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "foo", res.HTTPResponse.Header.Get("header1"))
	assert.Equal(t, "123", res.HTTPResponse.Header.Get("header2"))
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}

func (s strictServerInterface) TestRefHeaderNoContent(ctx context.Context, request pkg1.TestRefHeaderNoContentRequestObject) (pkg1.TestRefHeaderNoContentResponseObject, error) {
	return pkg1.TestRefHeaderNoContentdefaultResponse{
		Headers: pkg1.TestRespRefHeaderNoContentResponseHeaders{
			Header1: "foo",
			Header2: 123,
		},
		StatusCode: 204,
	}, nil
}

func TestRefHeaderNoContent(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-header-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefHeaderNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "foo", res.HTTPResponse.Header.Get("header1"))
	assert.Equal(t, "123", res.HTTPResponse.Header.Get("header2"))
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}

func (s strictServerInterface) TestRefHeaderFixedNoContent(ctx context.Context, request pkg1.TestRefHeaderFixedNoContentRequestObject) (pkg1.TestRefHeaderFixedNoContentResponseObject, error) {
	return pkg1.TestRespRefHeaderFixedNoContentResponse{
		Headers: pkg1.TestRespRefHeaderFixedNoContentResponseHeaders{
			Header1: "foo",
			Header2: 123,
		},
	}, nil
}

func TestRefHeaderFixedNoContent(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-header-fixed-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefHeaderFixedNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "foo", res.HTTPResponse.Header.Get("header1"))
	assert.Equal(t, "123", res.HTTPResponse.Header.Get("header2"))
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}
