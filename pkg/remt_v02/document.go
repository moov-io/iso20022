// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentRemt00100102 struct {
	XMLName xml.Name
	Attrs   []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
	RmtAdvc RemittanceAdviceV02 `xml:"RmtAdvc"`
}

func (doc DocumentRemt00100102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentRemt00100102) NameSpace() string {
	return utils.DocumentRemt00100102NameSpace
}

func (doc DocumentRemt00100102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
		RmtAdvc RemittanceAdviceV02 `xml:"RmtAdvc"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentRemt00100102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentRemt00100102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentRemt00200102 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr                `xml:",any,attr,omitempty" json:",omitempty"`
	RmtLctnAdvc RemittanceLocationAdviceV02 `xml:"RmtLctnAdvc"`
}

func (doc DocumentRemt00200102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentRemt00200102) NameSpace() string {
	return utils.DocumentRemt00200102NameSpace
}

func (doc DocumentRemt00200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr                `xml:",any,attr,omitempty" json:",omitempty"`
		RmtLctnAdvc RemittanceLocationAdviceV02 `xml:"RmtLctnAdvc"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentRemt00200102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentRemt00200102) GetAttrs() []utils.Attr {
	return doc.Attrs
}
