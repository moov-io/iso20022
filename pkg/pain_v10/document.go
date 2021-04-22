// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00100110 struct {
	Xmlns            string                              `xml:"xmlns,attr"`
	CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 `xml:"CstmrCdtTrfInitn"`
}

func (doc DocumentPain00100110) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00100110) NameSpace() string {
	return utils.DocumentPain00100110NameSpace
}

func (doc DocumentPain00100110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 `xml:"CstmrCdtTrfInitn"`
	}
	output.CstmrCdtTrfInitn = doc.CstmrCdtTrfInitn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain00700110 struct {
	Xmlns        string                     `xml:"xmlns,attr"`
	CstmrPmtRvsl CustomerPaymentReversalV10 `xml:"CstmrPmtRvsl"`
}

func (doc DocumentPain00700110) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00700110) NameSpace() string {
	return utils.DocumentPain00700110NameSpace
}

func (doc DocumentPain00700110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CstmrPmtRvsl CustomerPaymentReversalV10 `xml:"CstmrPmtRvsl"`
	}
	output.CstmrPmtRvsl = doc.CstmrPmtRvsl
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
