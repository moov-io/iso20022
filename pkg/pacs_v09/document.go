// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00800109 struct {
	Xmlns             string                          `xml:"xmlns,attr"`
	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"FIToFICstmrCdtTrf"`
}

func (doc DocumentPacs00800109) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00800109) NameSpace() string {
	return utils.DocumentPacs00800109NameSpace
}

func (doc DocumentPacs00800109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"FIToFICstmrCdtTrf"`
	}
	output.FIToFICstmrCdtTrf = doc.FIToFICstmrCdtTrf
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPacs00900109 struct {
	Xmlns    string                                `xml:"xmlns,attr"`
	FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf"`
}

func (doc DocumentPacs00900109) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00900109) NameSpace() string {
	return utils.DocumentPacs00900109NameSpace
}

func (doc DocumentPacs00900109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf"`
	}
	output.FICdtTrf = doc.FICdtTrf
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
