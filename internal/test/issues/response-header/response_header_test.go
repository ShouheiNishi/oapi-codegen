package header_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	header "github.com/deepmap/oapi-codegen/internal/test/issues/response-header"
	"github.com/stretchr/testify/assert"
)

func getPtr[T any](v T) *T {
	return &v
}

type serverInterfaceTestResponseHeader struct {
	t *testing.T
}

func (s serverInterfaceTestResponseHeader) Test(ctx context.Context, request header.TestRequestObject) (header.TestResponseObject, error) {
	assert.Equal(s.t, header.TestRequestObject{
		Params: header.TestParams{
			Test: &header.Object1{
				Field1: getPtr("value1"),
				Field2: getPtr("value2"),
			},
		},
	}, request)
	return header.Test200Response{
		Headers: header.Test200ResponseHeaders{
			Test: header.Object1{
				Field1: getPtr("value1"),
				Field2: getPtr("value2"),
			},
		},
	}, nil
}

func TestResponseHeader(t *testing.T) {
	h := header.Handler(header.NewStrictHandler(serverInterfaceTestResponseHeader{
		t: t,
	}, nil))
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "field1,value1,field2,value2", r.Header.Get("Test"))
		h.ServeHTTP(w, r)
	}))
	defer srv.Close()
	c, err := header.NewClientWithResponses(srv.URL)
	assert.NoError(t, err)
	res, err := c.TestWithResponse(context.TODO(), &header.TestParams{
		Test: &header.Object1{
			Field1: getPtr("value1"),
			Field2: getPtr("value2"),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "field1,value1,field2,value2", res.HTTPResponse.Header.Get("Test"))
}
