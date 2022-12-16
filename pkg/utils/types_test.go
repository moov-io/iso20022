package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDocumentFormat(t *testing.T) {
	require.Equal(t, DocumentTypeUnknown, GetDocumentFormat(nil))
	require.Equal(t, DocumentTypeUnknown, GetDocumentFormat([]byte("test")))

	require.Equal(t, DocumentTypeJson, GetDocumentFormat([]byte(`{"a": false}`)))
	require.Equal(t, DocumentTypeJson, GetDocumentFormat([]byte(`{"b": { "c": 12 }}`)))
	require.Equal(t, DocumentTypeJson, GetDocumentFormat([]byte(`{"a": false, "b": { "c": 12, "d": "test" }}`)))

	require.Equal(t, DocumentTypeXml, GetDocumentFormat([]byte(`<b>abc</b>`)))

	// Test deeply nested xml
	doc := strings.Repeat("<span>", 1001)
	doc += "test"
	doc += strings.Repeat("</span>", 1001)
	require.Equal(t, DocumentTypeXml, GetDocumentFormat([]byte(doc)))
}
