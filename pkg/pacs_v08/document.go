// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00300108 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFICstmrDrctDbt FIToFICustomerDirectDebitV08 `xml:"FIToFICstmrDrctDbt"`
}

func (doc DocumentPacs00300108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00300108) NameSpace() string {
	return utils.DocumentPacs00300108NameSpace
}

func (doc DocumentPacs00300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFICstmrDrctDbt FIToFICustomerDirectDebitV08 `xml:"FIToFICstmrDrctDbt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
