// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAuth001001V01 struct {
	XMLName    xml.Name                     `xml:"Document" json:"-"`
	InfReqOpng InformationRequestOpeningV01 `xml:"InfReqOpng"`
}

func (doc DocumentAuth001001V01) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth001001V01) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		InfReqOpng InformationRequestOpeningV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.001.001.01 InfReqOpng"`
	}
	output.InfReqOpng = doc.InfReqOpng
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.001.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentAuth002001V01 struct {
	XMLName    xml.Name                      `xml:"Document" json:"-"`
	InfReqRspn InformationRequestResponseV01 `xml:"InfReqRspn"`
}

func (doc DocumentAuth002001V01) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth002001V01) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		InfReqRspn InformationRequestResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.002.001.01 InfReqRspn"`
	}
	output.InfReqRspn = doc.InfReqRspn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.002.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentAuth003001V01 struct {
	XMLName             xml.Name                                      `xml:"Document" json:"-"`
	InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"InfReqStsChngNtfctn"`
}

func (doc DocumentAuth003001V01) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth003001V01) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 InfReqStsChngNtfctn"`
	}
	output.InfReqStsChngNtfctn = doc.InfReqStsChngNtfctn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.003.001.01")
	return e.EncodeElement(&output, start)
}
