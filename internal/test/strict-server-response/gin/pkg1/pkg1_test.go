// Code generated by generator/generate.go DO NOT EDIT.

package pkg1_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/deepmap/oapi-codegen/v2/internal/test/strict-server-response/gin/pkg1"
)

type strictServerInterface struct{}

func (s strictServerInterface) TestNoContent(ctx context.Context, request pkg1.TestNoContentRequestObject) (pkg1.TestNoContentResponseObject, error) {
	return pkg1.TestNoContentdefaultResponse{
		StatusCode: 204,
	}, nil
}

func TestNoContent(t *testing.T) {
	g := gin.New()
	pkg1.RegisterHandlers(g, pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		g.Handler().ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}

func (s strictServerInterface) TestFixedNoContent(ctx context.Context, request pkg1.TestFixedNoContentRequestObject) (pkg1.TestFixedNoContentResponseObject, error) {
	return pkg1.TestFixedNoContent204Response{}, nil
}

func TestFixedNoContent(t *testing.T) {
	g := gin.New()
	pkg1.RegisterHandlers(g, pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-fixed-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		g.Handler().ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestFixedNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}

func (s strictServerInterface) TestRefNoContent(ctx context.Context, request pkg1.TestRefNoContentRequestObject) (pkg1.TestRefNoContentResponseObject, error) {
	return pkg1.TestRefNoContentdefaultResponse{
		StatusCode: 204,
	}, nil
}

func TestRefNoContent(t *testing.T) {
	g := gin.New()
	pkg1.RegisterHandlers(g, pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		g.Handler().ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}

func (s strictServerInterface) TestRefFixedNoContent(ctx context.Context, request pkg1.TestRefFixedNoContentRequestObject) (pkg1.TestRefFixedNoContentResponseObject, error) {
	return pkg1.TestRespRefFixedNoContentResponse{}, nil
}

func TestRefFixedNoContent(t *testing.T) {
	g := gin.New()
	pkg1.RegisterHandlers(g, pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-ref-fixed-nocontent", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		g.Handler().ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestRefFixedNoContentWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 204, res.StatusCode())
	assert.Equal(t, "", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, []byte{}, res.Body)
}
