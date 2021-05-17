// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt10100101 struct {
	XMLName xml.Name
	Attrs   []utils.Attr   `xml:",any,attr,omitempty" json:",omitempty"`
	CretLmt CreateLimitV01 `xml:"CretLmt"`
}

func (doc DocumentCamt10100101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10100101) NameSpace() string {
	return utils.DocumentCamt10100101NameSpace
}

func (doc DocumentCamt10100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr   `xml:",any,attr,omitempty" json:",omitempty"`
		CretLmt CreateLimitV01 `xml:"CretLmt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt10200101 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	CretStgOrdr CreateStandingOrderV01 `xml:"CretStgOrdr"`
}

func (doc DocumentCamt10200101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10200101) NameSpace() string {
	return utils.DocumentCamt10200101NameSpace
}

func (doc DocumentCamt10200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
		CretStgOrdr CreateStandingOrderV01 `xml:"CretStgOrdr"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt10300101 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	CretRsvatn CreateReservationV01 `xml:"CretRsvatn"`
}

func (doc DocumentCamt10300101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10300101) NameSpace() string {
	return utils.DocumentCamt10300101NameSpace
}

func (doc DocumentCamt10300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		CretRsvatn CreateReservationV01 `xml:"CretRsvatn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt10400101 struct {
	XMLName xml.Name
	Attrs   []utils.Attr    `xml:",any,attr,omitempty" json:",omitempty"`
	CretMmb CreateMemberV01 `xml:"CretMmb"`
}

func (doc DocumentCamt10400101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt10400101) NameSpace() string {
	return utils.DocumentCamt10400101NameSpace
}

func (doc DocumentCamt10400101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr    `xml:",any,attr,omitempty" json:",omitempty"`
		CretMmb CreateMemberV01 `xml:"CretMmb"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
