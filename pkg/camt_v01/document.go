// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt10100101 struct {
	CretLmt CreateLimitV01 `xml:"CretLmt"`
}

func (doc DocumentCamt10100101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt10100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretLmt CreateLimitV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.101.001.01 CretLmt"`
	}
	output.CretLmt = doc.CretLmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.101.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentCamt10200101 struct {
	CretStgOrdr CreateStandingOrderV01 `xml:"CretStgOrdr"`
}

func (doc DocumentCamt10200101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt10200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretStgOrdr CreateStandingOrderV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 CretStgOrdr"`
	}
	output.CretStgOrdr = doc.CretStgOrdr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.102.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentCamt10300101 struct {
	CretRsvatn CreateReservationV01 `xml:"CretRsvatn"`
}

func (doc DocumentCamt10300101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt10300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretRsvatn CreateReservationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.103.001.01 CretRsvatn"`
	}
	output.CretRsvatn = doc.CretRsvatn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.103.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentCamt10400101 struct {
	CretMmb CreateMemberV01 `xml:"CretMmb"`
}

func (doc DocumentCamt10400101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt10400101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretMmb CreateMemberV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.104.001.01 CretMmb"`
	}
	output.CretMmb = doc.CretMmb
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.104.001.01")
	return e.EncodeElement(&output, start)
}
