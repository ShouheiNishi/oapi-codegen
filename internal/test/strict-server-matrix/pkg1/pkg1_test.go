// Code generated by generator/generate.go DO NOT EDIT.

package pkg1_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deepmap/oapi-codegen/v2/internal/test/strict-server-matrix/pkg1"
)

type strictServerInterface struct{}

func (s strictServerInterface) TestRefFormdata(ctx context.Context, request pkg1.TestRefFormdataRequestObject) (pkg1.TestRefFormdataResponseObject, error) {
	return pkg1.TestRefFormdatadefaultFormdataResponse{
		Body: pkg1.TestSchema{
			Field1: "bar",
			Field2: 456,
		},
		StatusCode: 200,
	}, nil
}

func TestRefFormdata(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-formdata", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefFormdataWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode())
	assert.Equal(t, "application/x-www-form-urlencoded", res.HTTPResponse.Header.Get("Content-Type"))
	form, err := url.ParseQuery(string(res.Body))
	assert.NoError(t, err)
	assert.Equal(t, url.Values{
		"field1": []string{"bar"},
		"field2": []string{"456"},
	}, form)
}

func (s strictServerInterface) TestRefFixedFormdata(ctx context.Context, request pkg1.TestRefFixedFormdataRequestObject) (pkg1.TestRefFixedFormdataResponseObject, error) {
	return pkg1.TestRefFixedFormdata200FormdataResponse{pkg1.TestRespRefFixedFormdataFormdataResponse(pkg1.TestSchema{
		Field1: "bar",
		Field2: 456,
	})}, nil
}

func TestRefFixedFormdata(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-fixed-formdata", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefFixedFormdataWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode())
	assert.Equal(t, "application/x-www-form-urlencoded", res.HTTPResponse.Header.Get("Content-Type"))
	form, err := url.ParseQuery(string(res.Body))
	assert.NoError(t, err)
	assert.Equal(t, url.Values{
		"field1": []string{"bar"},
		"field2": []string{"456"},
	}, form)
}

func (s strictServerInterface) TestRefHeaderFormdata(ctx context.Context, request pkg1.TestRefHeaderFormdataRequestObject) (pkg1.TestRefHeaderFormdataResponseObject, error) {
	return pkg1.TestRefHeaderFormdatadefaultFormdataResponse{
		Body: pkg1.TestSchema{
			Field1: "bar",
			Field2: 456,
		},
		Headers: pkg1.TestRespRefHeaderFormdataResponseHeaders{
			Header1: "foo",
			Header2: 123,
		},
		StatusCode: 200,
	}, nil
}

func TestRefHeaderFormdata(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-header-formdata", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefHeaderFormdataWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode())
	assert.Equal(t, "foo", res.HTTPResponse.Header.Get("header1"))
	assert.Equal(t, "123", res.HTTPResponse.Header.Get("header2"))
	assert.Equal(t, "application/x-www-form-urlencoded", res.HTTPResponse.Header.Get("Content-Type"))
	form, err := url.ParseQuery(string(res.Body))
	assert.NoError(t, err)
	assert.Equal(t, url.Values{
		"field1": []string{"bar"},
		"field2": []string{"456"},
	}, form)
}

func (s strictServerInterface) TestRefHeaderFixedFormdata(ctx context.Context, request pkg1.TestRefHeaderFixedFormdataRequestObject) (pkg1.TestRefHeaderFixedFormdataResponseObject, error) {
	return pkg1.TestRefHeaderFixedFormdata200FormdataResponse{pkg1.TestRespRefHeaderFixedFormdataFormdataResponse{
		Body: pkg1.TestSchema{
			Field1: "bar",
			Field2: 456,
		},
		Headers: pkg1.TestRespRefHeaderFixedFormdataResponseHeaders{
			Header1: "foo",
			Header2: 123,
		},
	}}, nil
}

func TestRefHeaderFixedFormdata(t *testing.T) {
	hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-header-fixed-formdata", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hh.ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefHeaderFixedFormdataWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode())
	assert.Equal(t, "foo", res.HTTPResponse.Header.Get("header1"))
	assert.Equal(t, "123", res.HTTPResponse.Header.Get("header2"))
	assert.Equal(t, "application/x-www-form-urlencoded", res.HTTPResponse.Header.Get("Content-Type"))
	form, err := url.ParseQuery(string(res.Body))
	assert.NoError(t, err)
	assert.Equal(t, url.Values{
		"field1": []string{"bar"},
		"field2": []string{"456"},
	}, form)
}
