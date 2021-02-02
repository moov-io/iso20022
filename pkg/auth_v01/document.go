// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAuth00100101 struct {
	Xmlns      string                       `xml:"xmlns,attr"`
	InfReqOpng InformationRequestOpeningV01 `xml:"InfReqOpng"`
}

func (doc DocumentAuth00100101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth00100101) NameSpace() string {
	return utils.DocumentAuth00100101NameSpace
}

func (doc DocumentAuth00100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		InfReqOpng InformationRequestOpeningV01 `xml:"InfReqOpng"`
	}
	output.InfReqOpng = doc.InfReqOpng
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentAuth00200101 struct {
	Xmlns      string                        `xml:"xmlns,attr"`
	InfReqRspn InformationRequestResponseV01 `xml:"InfReqRspn"`
}

func (doc DocumentAuth00200101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth00200101) NameSpace() string {
	return utils.DocumentAuth00200101NameSpace
}

func (doc DocumentAuth00200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		InfReqRspn InformationRequestResponseV01 `xml:"InfReqRspn"`
	}
	output.InfReqRspn = doc.InfReqRspn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentAuth00300101 struct {
	Xmlns               string                                        `xml:"xmlns,attr"`
	InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
}

func (doc DocumentAuth00300101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth00300101) NameSpace() string {
	return utils.DocumentAuth00300101NameSpace
}

func (doc DocumentAuth00300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
	}
	output.InfReqStsChngNtfctn = doc.InfReqStsChngNtfctn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
