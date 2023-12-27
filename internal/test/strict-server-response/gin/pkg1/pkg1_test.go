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

func (s strictServerInterface) TestSpecialJSON(ctx context.Context, request pkg1.TestSpecialJSONRequestObject) (pkg1.TestSpecialJSONResponseObject, error) {
	return pkg1.TestSpecialJSONdefaultApplicationTestPlusJSONResponse{
		Body: pkg1.TestSchema{
			Field1: "bar",
			Field2: 456,
		},
		StatusCode: 200,
	}, nil
}

func TestSpecialJSON(t *testing.T) {
	g := gin.New()
	pkg1.RegisterHandlers(g, pkg1.NewStrictHandler(strictServerInterface{}, nil))

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !assert.Equal(t, "/test-special-json", r.URL.Path) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		g.Handler().ServeHTTP(w, r)
	}))
	defer ts.Close()

	c, err := pkg1.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestSpecialJSONWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode())
	assert.Equal(t, "application/test+json", res.HTTPResponse.Header.Get("Content-Type"))
	assert.Equal(t, &pkg1.TestSchema{
		Field1: "bar",
		Field2: 456,
	}, res.ApplicationtestJSONDefault)
}
