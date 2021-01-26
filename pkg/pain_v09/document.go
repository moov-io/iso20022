// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00800109 struct {
	CstmrDrctDbtInitn CustomerDirectDebitInitiationV09 `xml:"CstmrDrctDbtInitn"`
}

func (doc DocumentPain00800109) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPain00800109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CstmrDrctDbtInitn CustomerDirectDebitInitiationV09 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CstmrDrctDbtInitn"`
	}
	output.CstmrDrctDbtInitn = doc.CstmrDrctDbtInitn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pain.008.001.09")
	return e.EncodeElement(&output, start)
}
