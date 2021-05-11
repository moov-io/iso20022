// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v11

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00200111 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	CstmrPmtStsRpt CustomerPaymentStatusReportV11 `xml:"CstmrPmtStsRpt"`
}

func (doc DocumentPain00200111) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00200111) NameSpace() string {
	return utils.DocumentPain00200111NameSpace
}

func (doc DocumentPain00200111) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrPmtStsRpt CustomerPaymentStatusReportV11 `xml:"CstmrPmtStsRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
