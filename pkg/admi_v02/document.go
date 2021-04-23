// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAdmi00400102 struct {
	XMLName                 *xml.Name                  `json:",omitempty"`
	DisableDefaultNamespace bool                       `xml:",omitempty" json:",omitempty"`
	SysEvtNtfctn            SystemEventNotificationV02 `xml:"SysEvtNtfctn"`
}

func (doc DocumentAdmi00400102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00400102) NameSpace() string {
	return utils.DocumentAdmi00400102NameSpace
}

func (doc DocumentAdmi00400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		SysEvtNtfctn SystemEventNotificationV02 `xml:"SysEvtNtfctn"`
	}

	output.SysEvtNtfctn = doc.SysEvtNtfctn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAdmi00900102 struct {
	XMLName                 *xml.Name            `json:",omitempty"`
	DisableDefaultNamespace bool                 `xml:",omitempty" json:",omitempty"`
	StatcDataReq            StaticDataRequestV02 `xml:"StatcDataReq"`
}

func (doc DocumentAdmi00900102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00900102) NameSpace() string {
	return utils.DocumentAdmi00900102NameSpace
}

func (doc DocumentAdmi00900102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		StatcDataReq StaticDataRequestV02 `xml:"StatcDataReq"`
	}

	output.StatcDataReq = doc.StatcDataReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAdmi01000102 struct {
	XMLName                 *xml.Name           `json:",omitempty"`
	DisableDefaultNamespace bool                `xml:",omitempty" json:",omitempty"`
	StatcDataRpt            StaticDataReportV02 `xml:"StatcDataRpt"`
}

func (doc DocumentAdmi01000102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAdmi01000102) NameSpace() string {
	return utils.DocumentAdmi01000102NameSpace
}

func (doc DocumentAdmi01000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		StatcDataRpt StaticDataReportV02 `xml:"StatcDataRpt"`
	}

	output.StatcDataRpt = doc.StatcDataRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
