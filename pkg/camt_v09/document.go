// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt05500109 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
	CstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:"CstmrPmtCxlReq"`
}

func (doc DocumentCamt05500109) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05500109) NameSpace() string {
	return utils.DocumentCamt05500109NameSpace
}

func (doc DocumentCamt05500109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:"CstmrPmtCxlReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt05500109) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt05500109) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentCamt05500109) InspectDocument() interface{} {
	return &doc.CstmrPmtCxlReq
}

type DocumentCamt05600109 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV09 `xml:"FIToFIPmtCxlReq"`
}

func (doc DocumentCamt05600109) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05600109) NameSpace() string {
	return utils.DocumentCamt05600109NameSpace
}

func (doc DocumentCamt05600109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV09 `xml:"FIToFIPmtCxlReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt05600109) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt05600109) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentCamt05600109) InspectDocument() interface{} {
	return &doc.FIToFIPmtCxlReq
}
