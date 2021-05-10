// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAdmi00400102 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	SysEvtNtfctn SystemEventNotificationV02 `xml:"SysEvtNtfctn"`
}

func (doc DocumentAdmi00400102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00400102) NameSpace() string {
	return utils.DocumentAdmi00400102NameSpace
}

func (doc DocumentAdmi00400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		SysEvtNtfctn SystemEventNotificationV02 `xml:"SysEvtNtfctn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi00900102 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	StatcDataReq StaticDataRequestV02 `xml:"StatcDataReq"`
}

func (doc DocumentAdmi00900102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00900102) NameSpace() string {
	return utils.DocumentAdmi00900102NameSpace
}

func (doc DocumentAdmi00900102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		StatcDataReq StaticDataRequestV02 `xml:"StatcDataReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi01000102 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
	StatcDataRpt StaticDataReportV02 `xml:"StatcDataRpt"`
}

func (doc DocumentAdmi01000102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi01000102) NameSpace() string {
	return utils.DocumentAdmi01000102NameSpace
}

func (doc DocumentAdmi01000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
		StatcDataRpt StaticDataReportV02 `xml:"StatcDataRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}
