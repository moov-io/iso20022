// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentCamt01800105 struct {
	GetBizDayInf GetBusinessDayInformationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 GetBizDayInf"`
}

func (doc DocumentCamt01800105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt02500105 struct {
	Rct ReceiptV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05 Rct"`
}

func (doc DocumentCamt02500105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03000105 struct {
	NtfctnOfCaseAssgnmt NotificationOfCaseAssignmentV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 NtfctnOfCaseAssgnmt"`
}

func (doc DocumentCamt03000105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03500105 struct {
	PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05 PrtryFrmtInvstgtn"`
}

func (doc DocumentCamt03500105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt04800105 struct {
	ModfyRsvatn ModifyReservationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.048.001.05 ModfyRsvatn"`
}

func (doc DocumentCamt04800105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03600105 struct {
	DbtAuthstnRspn DebitAuthorisationResponseV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.036.001.05 DbtAuthstnRspn"`
}

func (doc DocumentCamt03600105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt04900105 struct {
	DelRsvatn DeleteReservationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.049.001.05 DelRsvatn"`
}

func (doc DocumentCamt04900105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt05000105 struct {
	LqdtyCdtTrf LiquidityCreditTransferV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.050.001.05 LqdtyCdtTrf"`
}

func (doc DocumentCamt05000105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03900105 struct {
	CaseStsRpt CaseStatusReportV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.039.001.05 CaseStsRpt"`
}

func (doc DocumentCamt03900105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt04600105 struct {
	GetRsvatn GetReservationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.046.001.05 GetRsvatn"`
}

func (doc DocumentCamt04600105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt05100105 struct {
	LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LqdtyDbtTrf"`
}

func (doc DocumentCamt05100105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01800108 struct {
	AcctRptgReq AccountReportingRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05 AcctRptgReq"`
}

func (doc DocumentCamt01800108) Validate() error {
	return utils.Validate(&doc)
}
