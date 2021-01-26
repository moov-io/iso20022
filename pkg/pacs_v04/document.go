// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs01000104 struct {
	FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"FIDrctDbt"`
}

func (doc DocumentPacs01000104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs01000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.04 FIDrctDbt"`
	}
	output.FIDrctDbt = doc.FIDrctDbt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.010.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentPacs02800104 struct {
	FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"FIToFIPmtStsReq"`
}

func (doc DocumentPacs02800104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs02800104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04 FIToFIPmtStsReq"`
	}
	output.FIToFIPmtStsReq = doc.FIToFIPmtStsReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04")
	return e.EncodeElement(&output, start)
}
