// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt01300104 struct {
	GetMmb GetMemberV04 `xml:"GetMmb"`
}

func (doc DocumentCamt01300104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01300104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetMmb GetMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.013.001.04 GetMmb"`
	}
	output.GetMmb = doc.GetMmb
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.013.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01400104 struct {
	RtrMmb ReturnMemberV04 `xml:"RtrMmb"`
}

func (doc DocumentCamt01400104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01400104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrMmb ReturnMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RtrMmb"`
	}
	output.RtrMmb = doc.RtrMmb
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.014.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01500104 struct {
	ModfyMmb ModifyMemberV04 `xml:"ModfyMmb"`
}

func (doc DocumentCamt01500104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01500104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyMmb ModifyMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.015.001.04 ModfyMmb"`
	}
	output.ModfyMmb = doc.ModfyMmb
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.015.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01600104 struct {
	GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"GetCcyXchgRate"`
}

func (doc DocumentCamt01600104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01600104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.016.001.04 GetCcyXchgRate"`
	}
	output.GetCcyXchgRate = doc.GetCcyXchgRate
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.016.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt02000104 struct {
	GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"GetGnlBizInf"`
}

func (doc DocumentCamt02000104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 GetGnlBizInf"`
	}
	output.GetGnlBizInf = doc.GetGnlBizInf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.020.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01700104 struct {
	RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"RtrCcyXchgRate"`
}

func (doc DocumentCamt01700104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01700104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.04 RtrCcyXchgRate"`
	}
	output.RtrCcyXchgRate = doc.RtrCcyXchgRate
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.017.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03200104 struct {
	CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"CclCaseAssgnmt"`
}

func (doc DocumentCamt03200104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03200104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.032.001.04 CclCaseAssgnmt"`
	}
	output.CclCaseAssgnmt = doc.CclCaseAssgnmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.032.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03800104 struct {
	CaseStsRptReq CaseStatusReportRequestV04 `xml:"CaseStsRptReq"`
}

func (doc DocumentCamt03800104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03800104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CaseStsRptReq CaseStatusReportRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.038.001.04 CaseStsRptReq"`
	}
	output.CaseStsRptReq = doc.CaseStsRptReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.038.001.04")
	return e.EncodeElement(&output, start)
}

type DocumentCamt07000104 struct {
	RtrStgOrdr ReturnStandingOrderV04 `xml:"RtrStgOrdr"`
}

func (doc DocumentCamt07000104) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt07000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrStgOrdr ReturnStandingOrderV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 RtrStgOrdr"`
	}
	output.RtrStgOrdr = doc.RtrStgOrdr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.070.001.04")
	return e.EncodeElement(&output, start)
}
