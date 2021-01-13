// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentCamt01300104 struct {
	GetMmb GetMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.013.001.04 GetMmb"`
}

func (doc DocumentCamt01300104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01400104 struct {
	RtrMmb ReturnMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RtrMmb"`
}

func (doc DocumentCamt01400104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01500104 struct {
	ModfyMmb ModifyMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.015.001.04 ModfyMmb"`
}

func (doc DocumentCamt01500104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01600104 struct {
	GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.016.001.04 GetCcyXchgRate"`
}

func (doc DocumentCamt01600104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt02000104 struct {
	GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 GetGnlBizInf"`
}

func (doc DocumentCamt02000104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01700104 struct {
	RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.04 RtrCcyXchgRate"`
}

func (doc DocumentCamt01700104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03200104 struct {
	CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.032.001.04 CclCaseAssgnmt"`
}

func (doc DocumentCamt03200104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03800104 struct {
	CaseStsRptReq CaseStatusReportRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.038.001.04 CaseStsRptReq"`
}

func (doc DocumentCamt03800104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt07000104 struct {
	RtrStgOrdr ReturnStandingOrderV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 RtrStgOrdr"`
}

func (doc DocumentCamt07000104) Validate() error {
	return utils.Validate(&doc)
}
