// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

type DocumentPain01400108 struct {
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CdtrPmtActvtnReqStsRpt"`
}

type DocumentPain01300108 struct {
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtrPmtActvtnReq"`
}
