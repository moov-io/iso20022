// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt01300104 struct {
	Xmlns  string       `xml:"xmlns,attr"`
	GetMmb GetMemberV04 `xml:"GetMmb"`
}

func (doc DocumentCamt01300104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01300104) NameSpace() string {
	return utils.DocumentCamt01300104NameSpace
}

func (doc DocumentCamt01300104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetMmb GetMemberV04 `xml:"GetMmb"`
	}
	output.GetMmb = doc.GetMmb
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt01400104 struct {
	Xmlns  string          `xml:"xmlns,attr"`
	RtrMmb ReturnMemberV04 `xml:"RtrMmb"`
}

func (doc DocumentCamt01400104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01400104) NameSpace() string {
	return utils.DocumentCamt01400104NameSpace
}

func (doc DocumentCamt01400104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrMmb ReturnMemberV04 `xml:"RtrMmb"`
	}
	output.RtrMmb = doc.RtrMmb
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt01500104 struct {
	Xmlns    string          `xml:"xmlns,attr"`
	ModfyMmb ModifyMemberV04 `xml:"ModfyMmb"`
}

func (doc DocumentCamt01500104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01500104) NameSpace() string {
	return utils.DocumentCamt01500104NameSpace
}

func (doc DocumentCamt01500104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyMmb ModifyMemberV04 `xml:"ModfyMmb"`
	}
	output.ModfyMmb = doc.ModfyMmb
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt01600104 struct {
	Xmlns          string                     `xml:"xmlns,attr"`
	GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"GetCcyXchgRate"`
}

func (doc DocumentCamt01600104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01600104) NameSpace() string {
	return utils.DocumentCamt01600104NameSpace
}

func (doc DocumentCamt01600104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"GetCcyXchgRate"`
	}
	output.GetCcyXchgRate = doc.GetCcyXchgRate
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt02000104 struct {
	Xmlns        string                           `xml:"xmlns,attr"`
	GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"GetGnlBizInf"`
}

func (doc DocumentCamt02000104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02000104) NameSpace() string {
	return utils.DocumentCamt02000104NameSpace
}

func (doc DocumentCamt02000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"GetGnlBizInf"`
	}
	output.GetGnlBizInf = doc.GetGnlBizInf
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt01700104 struct {
	Xmlns          string                        `xml:"xmlns,attr"`
	RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"RtrCcyXchgRate"`
}

func (doc DocumentCamt01700104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01700104) NameSpace() string {
	return utils.DocumentCamt01700104NameSpace
}

func (doc DocumentCamt01700104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"RtrCcyXchgRate"`
	}
	output.RtrCcyXchgRate = doc.RtrCcyXchgRate
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt03200104 struct {
	Xmlns          string                  `xml:"xmlns,attr"`
	CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"CclCaseAssgnmt"`
}

func (doc DocumentCamt03200104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03200104) NameSpace() string {
	return utils.DocumentCamt03200104NameSpace
}

func (doc DocumentCamt03200104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"CclCaseAssgnmt"`
	}
	output.CclCaseAssgnmt = doc.CclCaseAssgnmt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt03800104 struct {
	Xmlns         string                     `xml:"xmlns,attr"`
	CaseStsRptReq CaseStatusReportRequestV04 `xml:"CaseStsRptReq"`
}

func (doc DocumentCamt03800104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03800104) NameSpace() string {
	return utils.DocumentCamt03800104NameSpace
}

func (doc DocumentCamt03800104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CaseStsRptReq CaseStatusReportRequestV04 `xml:"CaseStsRptReq"`
	}
	output.CaseStsRptReq = doc.CaseStsRptReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt07000104 struct {
	Xmlns      string                 `xml:"xmlns,attr"`
	RtrStgOrdr ReturnStandingOrderV04 `xml:"RtrStgOrdr"`
}

func (doc DocumentCamt07000104) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt07000104) NameSpace() string {
	return utils.DocumentCamt07000104NameSpace
}

func (doc DocumentCamt07000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrStgOrdr ReturnStandingOrderV04 `xml:"RtrStgOrdr"`
	}
	output.RtrStgOrdr = doc.RtrStgOrdr
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
