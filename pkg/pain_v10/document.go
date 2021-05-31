// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00100110 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 `xml:"CstmrCdtTrfInitn"`
}

func (doc DocumentPain00100110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00100110) NameSpace() string {
	return utils.DocumentPain00100110NameSpace
}

func (doc DocumentPain00100110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 `xml:"CstmrCdtTrfInitn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain00100110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain00100110) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentPain00700110 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	CstmrPmtRvsl CustomerPaymentReversalV10 `xml:"CstmrPmtRvsl"`
}

func (doc DocumentPain00700110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00700110) NameSpace() string {
	return utils.DocumentPain00700110NameSpace
}

func (doc DocumentPain00700110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		CstmrPmtRvsl CustomerPaymentReversalV10 `xml:"CstmrPmtRvsl"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain00700110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain00700110) GetAttrs() []utils.Attr {
	return doc.Attrs
}
