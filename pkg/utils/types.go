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

type Attr xml.Attr

func (a *Attr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	name.Local = a.Name.Local
	return xml.Attr{Name: name, Value: a.Value}, nil
}

func (a *Attr) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Name.Space != "" {
		attr.Name.Local = attr.Name.Space + ":" + attr.Name.Local
		attr.Name.Space = ""
	}
	*a = Attr(attr)
	return nil
}

const (
	TestTimeString      = "2014-11-12T11:45:26.371Z"
	XmlDefaultNamespace = "xmlns"
)
