// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v07

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00200107 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFIPmtStsRpt FIToFIPaymentStatusReportV07 `xml:"FIToFIPmtStsRpt"`
}

func (doc DocumentPacs00200107) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00200107) NameSpace() string {
	return utils.DocumentPacs00200107NameSpace
}

func (doc DocumentPacs00200107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtStsRpt FIToFIPaymentStatusReportV07 `xml:"FIToFIPmtStsRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
