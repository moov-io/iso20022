package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
)

const (
	DocumentTypeJson    = "json"
	DocumentTypeXml     = "xml"
	DocumentTypeUnknown = "unknown"
)

func isValidXML(buf []byte) bool {
	decoder := xml.NewDecoder(bytes.NewBuffer(buf))
	err := xml.Unmarshal(buf, new(interface{}))
	if err != nil {
		return false
	}
	for {
		err = decoder.Decode(new(interface{}))
		if err != nil {
			break
		}
	}
	return err == io.EOF
}

func isValidJSON(buf []byte) bool {
	var dummy map[string]interface{}
	if err := json.Unmarshal(buf, &dummy); err != nil {
		return false
	}
	return true
}

// Get buffer format
func GetBufferFormat(buf []byte) string {
	if isValidJSON(buf) {
		return DocumentTypeJson
	} else if isValidXML(buf) {
		return DocumentTypeXml
	}
	return DocumentTypeUnknown
}

type Attr struct {
	Name  string
	Value string
}

func (attr Attr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if len(attr.Name) == 0 {
		return xml.Attr{}, nil
	}
	return xml.Attr{Name: xml.Name{Local: attr.Name}, Value: attr.Value}, nil
}

const (
	TestTimeString = "2014-11-12T11:45:26.371Z"
)
