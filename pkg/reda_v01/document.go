// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

type DocumentReda06600101 struct {
	ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 ReqToPayCdtrEnrlmntReq"`
}

type DocumentReda07200101 struct {
	ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 ReqToPayDbtrActvtnCxlReq"`
}

type DocumentReda06700101 struct {
	ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 ReqToPayCdtrEnrlmntAmdmntReq"`
}

type DocumentReda07300101 struct {
	ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 ReqToPayDbtrActvtnStsRpt"`
}

type DocumentReda07100101 struct {
	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ReqToPayDbtrActvtnAmdmntReq"`
}

type DocumentReda07000101 struct {
	ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 ReqToPayDbtrActvtnReq"`
}

type DocumentReda06800101 struct {
	ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 ReqToPayCdtrEnrlmntCxlReq"`
}

type DocumentReda06000101 struct {
	ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 ReqToPayCdtrEnrlmntStsRpt"`
}
