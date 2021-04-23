// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v07

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00200107 struct {
	XMLName                 *xml.Name                    `json:",omitempty"`
	Xmlns                   string                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                         `xml:",omitempty" json:",omitempty"`
	FIToFIPmtStsRpt         FIToFIPaymentStatusReportV07 `xml:"FIToFIPmtStsRpt"`
}

func (doc DocumentPacs00200107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00200107) NameSpace() string {
	return utils.DocumentPacs00200107NameSpace
}

func (doc DocumentPacs00200107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtStsRpt FIToFIPaymentStatusReportV07 `xml:"FIToFIPmtStsRpt"`
	}
	output.FIToFIPmtStsRpt = doc.FIToFIPmtStsRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
