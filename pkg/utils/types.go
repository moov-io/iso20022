package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
)

type DocumentType string

const (
	DocumentTypeJson    DocumentType = "json"
	DocumentTypeXml     DocumentType = "xml"
	DocumentTypeUnknown DocumentType = "unknown"

	XmlDefaultNamespace = "xmlns"
)

func GetDocumentFormat(buf []byte) DocumentType {
	if json.Valid(buf) {
		return DocumentTypeJson
	}
	if isValidXML(buf) {
		return DocumentTypeXml
	}
	return DocumentTypeUnknown
}

func isValidXML(buf []byte) bool {
	decoder := xml.NewDecoder(bytes.NewBuffer(buf))
	err := xml.Unmarshal(buf, new(interface{}))
	if err != nil {
		return false
	}
	var attemptsLeft int = 100
	for {
		// Don't let the stream consume resources endlessly
		attemptsLeft -= 1
		if attemptsLeft <= 0 {
			return false
		}
		// Decode the next chunk
		err = decoder.Decode(new(interface{}))
		if err != nil {
			break
		}
	}
	return err == io.EOF
}
