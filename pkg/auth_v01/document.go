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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		InfReqOpng InformationRequestOpeningV01 `xml:"InfReqOpng"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		InfReqRspn InformationRequestResponseV01 `xml:"InfReqRspn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName             xml.Name
		Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
		InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
