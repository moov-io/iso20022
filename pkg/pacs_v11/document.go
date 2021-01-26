// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v11

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00200111 struct {
	FIToFIPmtStsRpt FIToFIPaymentStatusReportV11 `xml:"FIToFIPmtStsRpt"`
}

func (doc DocumentPacs00200111) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs00200111) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtStsRpt FIToFIPaymentStatusReportV11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11 FIToFIPmtStsRpt"`
	}
	output.FIToFIPmtStsRpt = doc.FIToFIPmtStsRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11")
	return e.EncodeElement(&output, start)
}
