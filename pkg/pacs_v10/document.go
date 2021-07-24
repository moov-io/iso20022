// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00200110 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFIPmtStsRpt FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt"`
}

func (doc DocumentPacs00200110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00200110) NameSpace() string {
	return utils.DocumentPacs00200110NameSpace
}

func (doc DocumentPacs00200110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtStsRpt FIToFIPaymentStatusReportV10 `xml:"FIToFIPmtStsRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPacs00200110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPacs00200110) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentPacs00200110) InspectDocument() interface{} {
	return &doc.FIToFIPmtStsRpt
}

type DocumentPacs00400110 struct {
	XMLName xml.Name
	Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	PmtRtr  PaymentReturnV10 `xml:"PmtRtr"`
}

func (doc DocumentPacs00400110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00400110) NameSpace() string {
	return utils.DocumentPacs00400110NameSpace
}

func (doc DocumentPacs00400110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
		PmtRtr  PaymentReturnV10 `xml:"PmtRtr"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPacs00400110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPacs00400110) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentPacs00400110) InspectDocument() interface{} {
	return &doc.PmtRtr
}

type DocumentPacs00700110 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"FIToFIPmtRvsl"`
}

func (doc DocumentPacs00700110) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00700110) NameSpace() string {
	return utils.DocumentPacs00700110NameSpace
}

func (doc DocumentPacs00700110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"FIToFIPmtRvsl"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPacs00700110) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPacs00700110) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentPacs00700110) InspectDocument() interface{} {
	return &doc.FIToFIPmtRvsl
}
