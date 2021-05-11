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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:"CstmrPmtCxlReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV09 `xml:"FIToFIPmtCxlReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
