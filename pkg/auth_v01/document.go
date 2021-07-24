// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAuth00100101 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	InfReqOpng InformationRequestOpeningV01 `xml:"InfReqOpng"`
}

func (doc DocumentAuth00100101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth00100101) NameSpace() string {
	return utils.DocumentAuth00100101NameSpace
}

func (doc DocumentAuth00100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		InfReqOpng InformationRequestOpeningV01 `xml:"InfReqOpng"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth00100101) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth00100101) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAuth00100101) InspectDocument() interface{} {
	return &doc.InfReqOpng
}

type DocumentAuth00200101 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	InfReqRspn InformationRequestResponseV01 `xml:"InfReqRspn"`
}

func (doc DocumentAuth00200101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth00200101) NameSpace() string {
	return utils.DocumentAuth00200101NameSpace
}

func (doc DocumentAuth00200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		InfReqRspn InformationRequestResponseV01 `xml:"InfReqRspn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth00200101) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth00200101) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAuth00200101) InspectDocument() interface{} {
	return &doc.InfReqRspn
}

type DocumentAuth00300101 struct {
	XMLName             xml.Name
	Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
	InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
}

func (doc DocumentAuth00300101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth00300101) NameSpace() string {
	return utils.DocumentAuth00300101NameSpace
}

func (doc DocumentAuth00300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName             xml.Name
		Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
		InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth00300101) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth00300101) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAuth00300101) InspectDocument() interface{} {
	return &doc.InfReqStsChngNtfctn
}
