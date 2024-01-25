package nullable_test

import (
	"encoding/json"
	"testing"

	newnullable "github.com/deepmap/oapi-codegen/v2/internal/test/issues/nullable-zero-value/new-nullable"
	oldnullable "github.com/deepmap/oapi-codegen/v2/internal/test/issues/nullable-zero-value/old-nullable"
	"github.com/stretchr/testify/assert"
)

// Test treatment additionalProperties in mergeOpenapiSchemas()
func TestIssue1236(t *testing.T) {
	var oldValue oldnullable.Test
	buf, err := json.Marshal(oldValue)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"NullableRequired": null}`, string(buf))

	var newValue newnullable.Test
	buf, err = json.Marshal(newValue)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"NullableRequired": null}`, string(buf))
}
