// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00300108 struct {
	FIToFICstmrDrctDbt FIToFICustomerDirectDebitV08 `xml:"FIToFICstmrDrctDbt"`
}

func (doc DocumentPacs00300108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPacs00300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFICstmrDrctDbt FIToFICustomerDirectDebitV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FIToFICstmrDrctDbt"`
	}
	output.FIToFICstmrDrctDbt = doc.FIToFICstmrDrctDbt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08")
	return e.EncodeElement(&output, start)
}
