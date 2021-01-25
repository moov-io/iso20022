// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt06900103 struct {
	GetStgOrdr GetStandingOrderV03 `xml:"GetStgOrdr"`
}

func (doc DocumentCamt06900103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt06900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetStgOrdr GetStandingOrderV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 GetStgOrdr"`
	}
	output.GetStgOrdr = doc.GetStgOrdr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.069.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentCamt07100103 struct {
	DelStgOrdr DeleteStandingOrderV03 `xml:"DelStgOrdr"`
}

func (doc DocumentCamt07100103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt07100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelStgOrdr DeleteStandingOrderV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.071.001.03 DelStgOrdr"`
	}
	output.DelStgOrdr = doc.DelStgOrdr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.071.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentCamt08600103 struct {
	BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
}

func (doc DocumentCamt08600103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt08600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BkSvcsBllgStmt"`
	}
	output.BkSvcsBllgStmt = doc.BkSvcsBllgStmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.086.001.03")
	return e.EncodeElement(&output, start)
}
