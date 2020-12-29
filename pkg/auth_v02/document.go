// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentAuth018001V02 struct {
	CtrctRegnReq ContractRegistrationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctRegnReq"`
}

func (doc DocumentAuth018001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth019001V02 struct {
	CtrctRegnConf ContractRegistrationConfirmationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 CtrctRegnConf"`
}

func (doc DocumentAuth019001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth020001V02 struct {
	CtrctRegnClsrReq ContractRegistrationClosureRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 CtrctRegnClsrReq"`
}

func (doc DocumentAuth020001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth021001V02 struct {
	CtrctRegnAmdmntReq ContractRegistrationAmendmentRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 CtrctRegnAmdmntReq"`
}

func (doc DocumentAuth021001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth022001V02 struct {
	CtrctRegnStmt ContractRegistrationStatementV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctRegnStmt"`
}

func (doc DocumentAuth022001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth023001V02 struct {
	CtrctRegnStmtReq ContractRegistrationStatementRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 CtrctRegnStmtReq"`
}

func (doc DocumentAuth023001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth024001V02 struct {
	PmtRgltryInfNtfctn PaymentRegulatoryInformationNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 PmtRgltryInfNtfctn"`
}

func (doc DocumentAuth024001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth025001V02 struct {
	CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CcyCtrlSpprtgDocDlvry"`
}

func (doc DocumentAuth025001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth026001V02 struct {
	CcyCtrlReqOrLttr CurrencyControlRequestOrLetterV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 CcyCtrlReqOrLttr"`
}

func (doc DocumentAuth026001V02) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth027001V02 struct {
	CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CcyCtrlStsAdvc"`
}

func (doc DocumentAuth027001V02) Validate() error {
	return utils.Validate(&doc)
}
