package issue1301_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	issue1301 "github.com/deepmap/oapi-codegen/v2/internal/test/issues/issue-1301"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testStrictServerInterface struct {
	t *testing.T
}

// (GET /test)
func (s *testStrictServerInterface) Test(ctx context.Context, request issue1301.TestRequestObject) (issue1301.TestResponseObject, error) {
	sBar := "bar"
	return issue1301.Test204Response{
		Headers: issue1301.Test204ResponseHeaders{
			Header2: "foo",
			Header3: &sBar,
		},
	}, nil
}

func TestIssue1301(t *testing.T) {
	g := gin.Default()
	issue1301.RegisterHandlersWithOptions(g,
		issue1301.NewStrictHandler(&testStrictServerInterface{t: t}, nil),
		issue1301.GinServerOptions{})
	ts := httptest.NewServer(g)
	defer ts.Close()

	c, err := issue1301.NewClientWithResponses(ts.URL)
	assert.NoError(t, err)
	res, err := c.TestWithResponse(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode())
	assert.Empty(t, res.HTTPResponse.Header.Values("header1"))
	assert.Equal(t, []string{"foo"}, res.HTTPResponse.Header.Values("header2"))
	assert.Equal(t, []string{"bar"}, res.HTTPResponse.Header.Values("header3"))
}
