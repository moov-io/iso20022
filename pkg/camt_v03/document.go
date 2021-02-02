// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt06900103 struct {
	Xmlns      string              `xml:"xmlns,attr"`
	GetStgOrdr GetStandingOrderV03 `xml:"GetStgOrdr"`
}

func (doc DocumentCamt06900103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt06900103) NameSpace() string {
	return utils.DocumentCamt06900103NameSpace
}

func (doc DocumentCamt06900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetStgOrdr GetStandingOrderV03 `xml:"GetStgOrdr"`
	}
	output.GetStgOrdr = doc.GetStgOrdr
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt07100103 struct {
	Xmlns      string                 `xml:"xmlns,attr"`
	DelStgOrdr DeleteStandingOrderV03 `xml:"DelStgOrdr"`
}

func (doc DocumentCamt07100103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt07100103) NameSpace() string {
	return utils.DocumentCamt07100103NameSpace
}

func (doc DocumentCamt07100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelStgOrdr DeleteStandingOrderV03 `xml:"DelStgOrdr"`
	}
	output.DelStgOrdr = doc.DelStgOrdr
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt08600103 struct {
	Xmlns          string                          `xml:"xmlns,attr"`
	BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
}

func (doc DocumentCamt08600103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt08600103) NameSpace() string {
	return utils.DocumentCamt08600103NameSpace
}

func (doc DocumentCamt08600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
	}
	output.BkSvcsBllgStmt = doc.BkSvcsBllgStmt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
