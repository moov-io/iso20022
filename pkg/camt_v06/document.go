// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt02100106 struct {
	RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"RtrGnlBizInf"`
}

func (doc DocumentCamt02100106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02100106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.021.001.06 RtrGnlBizInf"`
	}
	output.RtrGnlBizInf = doc.RtrGnlBizInf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.021.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt02400106 struct {
	ModfyStgOrdr ModifyStandingOrderV06 `xml:"ModfyStgOrdr"`
}

func (doc DocumentCamt02400106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02400106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyStgOrdr ModifyStandingOrderV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.024.001.06 ModfyStgOrdr"`
	}
	output.ModfyStgOrdr = doc.ModfyStgOrdr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.024.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03100106 struct {
	RjctInvstgtn RejectInvestigationV06 `xml:"RjctInvstgtn"`
}

func (doc DocumentCamt03100106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03100106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RjctInvstgtn RejectInvestigationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.031.001.06 RjctInvstgtn"`
	}
	output.RjctInvstgtn = doc.RjctInvstgtn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.031.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt04700106 struct {
	RtrRsvatn ReturnReservationV06 `xml:"RtrRsvatn"`
}

func (doc DocumentCamt04700106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt04700106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrRsvatn ReturnReservationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.047.001.06 RtrRsvatn"`
	}
	output.RtrRsvatn = doc.RtrRsvatn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.047.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05700106 struct {
	NtfctnToRcv NotificationToReceiveV06 `xml:"NtfctnToRcv"`
}

func (doc DocumentCamt05700106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05700106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnToRcv NotificationToReceiveV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.057.001.06 NtfctnToRcv"`
	}
	output.NtfctnToRcv = doc.NtfctnToRcv
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.057.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03300106 struct {
	ReqForDplct RequestForDuplicateV06 `xml:"ReqForDplct"`
}

func (doc DocumentCamt03300106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03300106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqForDplct RequestForDuplicateV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.033.001.06 ReqForDplct"`
	}
	output.ReqForDplct = doc.ReqForDplct
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.033.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05800106 struct {
	NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"NtfctnToRcvCxlAdvc"`
}

func (doc DocumentCamt05800106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05800106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.058.001.06 NtfctnToRcvCxlAdvc"`
	}
	output.NtfctnToRcvCxlAdvc = doc.NtfctnToRcvCxlAdvc
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.058.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05900106 struct {
	NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"NtfctnToRcvStsRpt"`
}

func (doc DocumentCamt05900106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05900106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.06 NtfctnToRcvStsRpt"`
	}
	output.NtfctnToRcvStsRpt = doc.NtfctnToRcvStsRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.059.001.06")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03400106 struct {
	Dplct DuplicateV06 `xml:"Dplct"`
}

func (doc DocumentCamt03400106) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03400106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		Dplct DuplicateV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.034.001.06 Dplct"`
	}
	output.Dplct = doc.Dplct
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.034.001.06")
	return e.EncodeElement(&output, start)
}
