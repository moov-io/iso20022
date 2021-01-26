// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00400110 struct {
	PmtRtr PaymentReturnV10 `xml:"PmtRtr"`
}

func (doc DocumentPacs00400110) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs00400110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PmtRtr PaymentReturnV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PmtRtr"`
	}
	output.PmtRtr = doc.PmtRtr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10")
	return e.EncodeElement(&output, start)
}

type DocumentPacs00700110 struct {
	FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"FIToFIPmtRvsl"`
}

func (doc DocumentPacs00700110) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs00700110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.10 FIToFIPmtRvsl"`
	}
	output.FIToFIPmtRvsl = doc.FIToFIPmtRvsl
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.007.001.10")
	return e.EncodeElement(&output, start)
}
