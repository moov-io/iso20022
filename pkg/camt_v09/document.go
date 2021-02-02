// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt05500109 struct {
	Xmlns          string                                `xml:"xmlns,attr"`
	CstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:"CstmrPmtCxlReq"`
}

func (doc DocumentCamt05500109) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05500109) NameSpace() string {
	return utils.DocumentCamt05500109NameSpace
}

func (doc DocumentCamt05500109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.055.001.09 CstmrPmtCxlReq"`
	}
	output.CstmrPmtCxlReq = doc.CstmrPmtCxlReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05600109 struct {
	Xmlns           string                              `xml:"xmlns,attr"`
	FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV09 `xml:"FIToFIPmtCxlReq"`
}

func (doc DocumentCamt05600109) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05600109) NameSpace() string {
	return utils.DocumentCamt05600109NameSpace
}

func (doc DocumentCamt05600109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV09 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.09 FIToFIPmtCxlReq"`
	}
	output.FIToFIPmtCxlReq = doc.FIToFIPmtCxlReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
