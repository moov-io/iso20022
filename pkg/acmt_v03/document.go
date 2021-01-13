// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentAcmt00700103 struct {
	AcctOpngReq AccountOpeningRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03 AcctOpngReq"`
}

func (doc DocumentAcmt00700103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt00800103 struct {
	AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AcctOpngAmdmntReq"`
}

func (doc DocumentAcmt00800103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt00900103 struct {
	AcctOpngAddtlInfReq AccountOpeningAdditionalInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.03 AcctOpngAddtlInfReq"`
}

func (doc DocumentAcmt00900103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01000103 struct {
	AcctReqAck AccountRequestAcknowledgementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.010.001.03 AcctReqAck"`
}

func (doc DocumentAcmt01000103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01100103 struct {
	AcctReqRjctn AccountRequestRejectionV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.011.001.03 AcctReqRjctn"`
}

func (doc DocumentAcmt01100103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01200103 struct {
	AcctAddtlInfReq AccountAdditionalInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.012.001.03 AcctAddtlInfReq"`
}

func (doc DocumentAcmt01200103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01300103 struct {
	AcctRptReq AccountReportRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AcctRptReq"`
}

func (doc DocumentAcmt01300103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01400103 struct {
	AcctRpt AccountReportV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.03 AcctRpt"`
}

func (doc DocumentAcmt01400103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01500103 struct {
	AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.03 AcctExcldMndtMntncReq"`
}

func (doc DocumentAcmt01500103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01600103 struct {
	AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.03 AcctExcldMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01600103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01700103 struct {
	AcctMndtMntncReq AccountMandateMaintenanceRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.017.001.03 AcctMndtMntncReq"`
}

func (doc DocumentAcmt01700103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01800103 struct {
	AcctMndtMntncAmdmntReq AccountMandateMaintenanceAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.018.001.03 AcctMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01800103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01900103 struct {
	AcctClsgReq AccountClosingRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.019.001.03 AcctClsgReq"`
}

func (doc DocumentAcmt01900103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02000103 struct {
	AcctClsgAmdmntReq AccountClosingAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.020.001.03 AcctClsgAmdmntReq"`
}

func (doc DocumentAcmt02000103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02100103 struct {
	AcctClsgAddtlInfReq AccountClosingAdditionalInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.021.001.03 AcctClsgAddtlInfReq"`
}

func (doc DocumentAcmt02100103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02700103 struct {
	AcctSwtchInfReq AccountSwitchInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.027.001.03 AcctSwtchInfReq"`
}

func (doc DocumentAcmt02700103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02800103 struct {
	AcctSwtchInfRspn AccountSwitchInformationResponseV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.028.001.03 AcctSwtchInfRspn"`
}

func (doc DocumentAcmt02800103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02900103 struct {
	AcctSwtchCclExstgPmt AccountSwitchCancelExistingPaymentV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.029.001.03 AcctSwtchCclExstgPmt"`
}

func (doc DocumentAcmt02900103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03100103 struct {
	AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.03 AcctSwtchReqBalTrf"`
}

func (doc DocumentAcmt03100103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03200103 struct {
	AcctSwtchBalTrfAck AccountSwitchBalanceTransferAcknowledgementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.032.001.03 AcctSwtchBalTrfAck"`
}

func (doc DocumentAcmt03200103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03400103 struct {
	AcctSwtchReqPmt AccountSwitchRequestPaymentV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.034.001.03 AcctSwtchReqPmt"`
}

func (doc DocumentAcmt03400103) Validate() error {
	return utils.Validate(&doc)
}
