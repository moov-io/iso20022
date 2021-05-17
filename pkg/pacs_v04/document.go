// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs01000104 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
	FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"FIDrctDbt"`
}

func (doc DocumentPacs01000104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs01000104) NameSpace() string {
	return utils.DocumentPacs01000104NameSpace
}

func (doc DocumentPacs01000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
		FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"FIDrctDbt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentPacs02800104 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"FIToFIPmtStsReq"`
}

func (doc DocumentPacs02800104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs02800104) NameSpace() string {
	return utils.DocumentPacs02800104NameSpace
}

func (doc DocumentPacs02800104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"FIToFIPmtStsReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
