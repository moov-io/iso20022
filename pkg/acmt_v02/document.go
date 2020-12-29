// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentAcmt02200102 struct {
	IdModAdvc IdentificationModificationAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 IdModAdvc"`
}

func (doc DocumentAcmt02200102) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02300102 struct {
	IdVrfctnReq IdentificationVerificationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02 IdVrfctnReq"`
}

func (doc DocumentAcmt02300102) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt02400102 struct {
	IdVrfctnRpt IdentificationVerificationReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.024.001.02 IdVrfctnRpt"`
}

func (doc DocumentAcmt02400102) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03000102 struct {
	AcctSwtchReqRdrctn AccountSwitchRequestRedirectionV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.030.001.02 AcctSwtchReqRdrctn"`
}

func (doc DocumentAcmt03000102) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03300102 struct {
	AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.033.001.02 AcctSwtchNtfyAcctSwtchCmplt"`
}

func (doc DocumentAcmt03300102) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03500102 struct {
	AcctSwtchPmtRspn AccountSwitchPaymentResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.035.001.02 AcctSwtchPmtRspn"`
}

func (doc DocumentAcmt03500102) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt03700102 struct {
	AcctSwtchTechRjctn AccountSwitchTechnicalRejectionV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.037.001.02 AcctSwtchTechRjctn"`
}

func (doc DocumentAcmt03700102) Validate() error {
	return utils.Validate(&doc)
}
