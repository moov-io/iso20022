// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt02800110 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	AddtlPmtInf AdditionalPaymentInformationV10 `xml:"AddtlPmtInf"`
}

func (doc DocumentCamt02800110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02800110) NameSpace() string {
	return utils.DocumentCamt02800110NameSpace
}

func (doc DocumentCamt02800110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		AddtlPmtInf AdditionalPaymentInformationV10 `xml:"AddtlPmtInf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt02800110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt02800110) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentCamt02900110 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	RsltnOfInvstgtn ResolutionOfInvestigationV10 `xml:"RsltnOfInvstgtn"`
}

func (doc DocumentCamt02900110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02900110) NameSpace() string {
	return utils.DocumentCamt02900110NameSpace
}

func (doc DocumentCamt02900110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		RsltnOfInvstgtn ResolutionOfInvestigationV10 `xml:"RsltnOfInvstgtn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt02900110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt02900110) GetAttrs() []utils.Attr {
	return doc.Attrs
}
