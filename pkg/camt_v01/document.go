// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt10100101 struct {
	Xmlns   string         `xml:"xmlns,attr"`
	CretLmt CreateLimitV01 `xml:"CretLmt"`
}

func (doc DocumentCamt10100101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10100101) NameSpace() string {
	return utils.DocumentCamt10100101NameSpace
}

func (doc DocumentCamt10100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretLmt CreateLimitV01 `xml:"CretLmt"`
	}
	output.CretLmt = doc.CretLmt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt10200101 struct {
	Xmlns       string                 `xml:"xmlns,attr"`
	CretStgOrdr CreateStandingOrderV01 `xml:"CretStgOrdr"`
}

func (doc DocumentCamt10200101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10200101) NameSpace() string {
	return utils.DocumentCamt10200101NameSpace
}

func (doc DocumentCamt10200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretStgOrdr CreateStandingOrderV01 `xml:"CretStgOrdr"`
	}
	output.CretStgOrdr = doc.CretStgOrdr
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt10300101 struct {
	Xmlns      string               `xml:"xmlns,attr"`
	CretRsvatn CreateReservationV01 `xml:"CretRsvatn"`
}

func (doc DocumentCamt10300101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10300101) NameSpace() string {
	return utils.DocumentCamt10300101NameSpace
}

func (doc DocumentCamt10300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretRsvatn CreateReservationV01 `xml:"CretRsvatn"`
	}
	output.CretRsvatn = doc.CretRsvatn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt10400101 struct {
	Xmlns   string          `xml:"xmlns,attr"`
	CretMmb CreateMemberV01 `xml:"CretMmb"`
}

func (doc DocumentCamt10400101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10400101) NameSpace() string {
	return utils.DocumentCamt10400101NameSpace
}

func (doc DocumentCamt10400101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CretMmb CreateMemberV01 `xml:"CretMmb"`
	}
	output.CretMmb = doc.CretMmb
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
