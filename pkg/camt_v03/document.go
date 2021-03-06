// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt03500103 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
	PrtryFrmtInvstgtn ProprietaryFormatInvestigationV03 `xml:"PrtryFrmtInvstgtn"`
}

func (doc DocumentCamt03500103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03500103) NameSpace() string {
	return utils.DocumentCamt03500103NameSpace
}

func (doc DocumentCamt03500103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
		PrtryFrmtInvstgtn ProprietaryFormatInvestigationV03 `xml:"PrtryFrmtInvstgtn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt03500103) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt03500103) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentCamt06900103 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
	GetStgOrdr GetStandingOrderV03 `xml:"GetStgOrdr"`
}

func (doc DocumentCamt06900103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt06900103) NameSpace() string {
	return utils.DocumentCamt06900103NameSpace
}

func (doc DocumentCamt06900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
		GetStgOrdr GetStandingOrderV03 `xml:"GetStgOrdr"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt06900103) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt06900103) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentCamt07100103 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	DelStgOrdr DeleteStandingOrderV03 `xml:"DelStgOrdr"`
}

func (doc DocumentCamt07100103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt07100103) NameSpace() string {
	return utils.DocumentCamt07100103NameSpace
}

func (doc DocumentCamt07100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
		DelStgOrdr DeleteStandingOrderV03 `xml:"DelStgOrdr"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt07100103) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt07100103) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentCamt08600103 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
}

func (doc DocumentCamt08600103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt08600103) NameSpace() string {
	return utils.DocumentCamt08600103NameSpace
}

func (doc DocumentCamt08600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentCamt08600103) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentCamt08600103) GetAttrs() []utils.Attr {
	return doc.Attrs
}
