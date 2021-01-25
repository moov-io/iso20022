// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt01800105 struct {
	GetBizDayInf GetBusinessDayInformationV05 `xml:"GetBizDayInf"`
}

func (doc DocumentCamt01800105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetBizDayInf GetBusinessDayInformationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 GetBizDayInf"`
	}
	output.GetBizDayInf = doc.GetBizDayInf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.018.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt02500105 struct {
	Rct ReceiptV05 `xml:"Rct"`
}

func (doc DocumentCamt02500105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02500105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		Rct ReceiptV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Rct"`
	}
	output.Rct = doc.Rct
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.025.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03000105 struct {
	NtfctnOfCaseAssgnmt NotificationOfCaseAssignmentV05 `xml:"NtfctnOfCaseAssgnmt"`
}

func (doc DocumentCamt03000105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnOfCaseAssgnmt NotificationOfCaseAssignmentV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 NtfctnOfCaseAssgnmt"`
	}
	output.NtfctnOfCaseAssgnmt = doc.NtfctnOfCaseAssgnmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.030.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03500105 struct {
	PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"PrtryFrmtInvstgtn"`
}

func (doc DocumentCamt03500105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03500105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryFrmtInvstgtn"`
	}
	output.PrtryFrmtInvstgtn = doc.PrtryFrmtInvstgtn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.035.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt04800105 struct {
	ModfyRsvatn ModifyReservationV05 `xml:"ModfyRsvatn"`
}

func (doc DocumentCamt04800105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt04800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyRsvatn ModifyReservationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.048.001.05 ModfyRsvatn"`
	}
	output.ModfyRsvatn = doc.ModfyRsvatn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.048.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03600105 struct {
	DbtAuthstnRspn DebitAuthorisationResponseV05 `xml:"DbtAuthstnRspn"`
}

func (doc DocumentCamt03600105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DbtAuthstnRspn DebitAuthorisationResponseV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.036.001.05 DbtAuthstnRspn"`
	}
	output.DbtAuthstnRspn = doc.DbtAuthstnRspn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.036.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt04900105 struct {
	DelRsvatn DeleteReservationV05 `xml:"DelRsvatn"`
}

func (doc DocumentCamt04900105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt04900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelRsvatn DeleteReservationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.049.001.05 DelRsvatn"`
	}
	output.DelRsvatn = doc.DelRsvatn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.049.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05000105 struct {
	LqdtyCdtTrf LiquidityCreditTransferV05 `xml:"LqdtyCdtTrf"`
}

func (doc DocumentCamt05000105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		LqdtyCdtTrf LiquidityCreditTransferV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.050.001.05 LqdtyCdtTrf"`
	}
	output.LqdtyCdtTrf = doc.LqdtyCdtTrf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.050.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03900105 struct {
	CaseStsRpt CaseStatusReportV05 `xml:"CaseStsRpt"`
}

func (doc DocumentCamt03900105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CaseStsRpt CaseStatusReportV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.039.001.05 CaseStsRpt"`
	}
	output.CaseStsRpt = doc.CaseStsRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.039.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt04600105 struct {
	GetRsvatn GetReservationV05 `xml:"GetRsvatn"`
}

func (doc DocumentCamt04600105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt04600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetRsvatn GetReservationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.05 GetRsvatn"`
	}
	output.GetRsvatn = doc.GetRsvatn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.046.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05100105 struct {
	LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"LqdtyDbtTrf"`
}

func (doc DocumentCamt05100105) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05100105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LqdtyDbtTrf"`
	}
	output.LqdtyDbtTrf = doc.LqdtyDbtTrf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.051.001.05")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01800108 struct {
	AcctRptgReq AccountReportingRequestV05 `xml:"AcctRptgReq"`
}

func (doc DocumentCamt01800108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01800108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctRptgReq AccountReportingRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05 AcctRptgReq"`
	}
	output.AcctRptgReq = doc.AcctRptgReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05")
	return e.EncodeElement(&output, start)
}
