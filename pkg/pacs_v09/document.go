// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00800109 struct {
	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"FIToFICstmrCdtTrf"`
}

func (doc DocumentPacs00800109) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs00800109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FIToFICstmrCdtTrf"`
	}
	output.FIToFICstmrCdtTrf = doc.FIToFICstmrCdtTrf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09")
	return e.EncodeElement(&output, start)
}

type DocumentPacs00900109 struct {
	FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf"`
}

func (doc DocumentPacs00900109) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs00900109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.09 FICdtTrf"`
	}
	output.FICdtTrf = doc.FICdtTrf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.009.001.09")
	return e.EncodeElement(&output, start)
}
