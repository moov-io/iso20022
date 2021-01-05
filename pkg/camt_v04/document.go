// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

type DocumentCamt01300104 struct {
	GetMmb GetMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.013.001.04 GetMmb"`
}

type DocumentCamt01400104 struct {
	RtrMmb ReturnMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04 RtrMmb"`
}

type DocumentCamt01500104 struct {
	ModfyMmb ModifyMemberV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.015.001.04 ModfyMmb"`
}

type DocumentCamt01600104 struct {
	GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.016.001.04 GetCcyXchgRate"`
}

type DocumentCamt02000104 struct {
	GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04 GetGnlBizInf"`
}

type DocumentCamt01700104 struct {
	RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.017.001.04 RtrCcyXchgRate"`
}

type DocumentCamt03200104 struct {
	CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.032.001.04 CclCaseAssgnmt"`
}

type DocumentCamt03800104 struct {
	CaseStsRptReq CaseStatusReportRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.038.001.04 CaseStsRptReq"`
}

type DocumentCamt07000104 struct {
	RtrStgOrdr ReturnStandingOrderV04 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04 RtrStgOrdr"`
}
