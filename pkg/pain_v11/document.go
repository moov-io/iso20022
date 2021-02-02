// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v11

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00200111 struct {
	Xmlns          string                         `xml:"xmlns,attr"`
	CstmrPmtStsRpt CustomerPaymentStatusReportV11 `xml:"CstmrPmtStsRpt"`
}

func (doc DocumentPain00200111) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00200111) NameSpace() string {
	return utils.DocumentPain00200111NameSpace
}

func (doc DocumentPain00200111) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CstmrPmtStsRpt CustomerPaymentStatusReportV11 `xml:"CstmrPmtStsRpt"`
	}
	output.CstmrPmtStsRpt = doc.CstmrPmtStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
