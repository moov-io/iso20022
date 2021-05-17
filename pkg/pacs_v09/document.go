// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00800109 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"FIToFICstmrCdtTrf"`
}

func (doc DocumentPacs00800109) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00800109) NameSpace() string {
	return utils.DocumentPacs00800109NameSpace
}

func (doc DocumentPacs00800109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"FIToFICstmrCdtTrf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentPacs00900109 struct {
	XMLName  xml.Name
	Attrs    []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
	FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf"`
}

func (doc DocumentPacs00900109) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00900109) NameSpace() string {
	return utils.DocumentPacs00900109NameSpace
}

func (doc DocumentPacs00900109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName  xml.Name
		Attrs    []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
		FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"FICdtTrf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
