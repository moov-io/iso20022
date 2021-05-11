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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrDrctDbtInitn CustomerDirectDebitInitiationV09 `xml:"CstmrDrctDbtInitn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
