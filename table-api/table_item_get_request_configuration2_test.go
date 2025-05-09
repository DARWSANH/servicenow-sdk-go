package tableapi

import (
	"testing"

	"github.com/wiryvalance/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemGetRequestConfiguration2_toConfiguration(t *testing.T) {
	config := &tableItemGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    &TableItemRequestBuilderGetQueryParameters{},
		data:     nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	expected := &core.RequestConfiguration{
		Header:          nil,
		QueryParameters: &TableItemRequestBuilderGetQueryParameters{},
		Data:            nil,
		Response:        &TableItemResponse2[TableEntry]{},
	}

	assert.Equal(t, expected, config.toConfiguration())
}
