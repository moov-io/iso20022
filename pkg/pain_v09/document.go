// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00800109 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
	CstmrDrctDbtInitn CustomerDirectDebitInitiationV09 `xml:"CstmrDrctDbtInitn"`
}

func (doc DocumentPain00800109) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00800109) NameSpace() string {
	return utils.DocumentPain00800109NameSpace
}

func (doc DocumentPain00800109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrDrctDbtInitn CustomerDirectDebitInitiationV09 `xml:"CstmrDrctDbtInitn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain00800109) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain00800109) GetAttrs() []utils.Attr {
	return doc.Attrs
}
