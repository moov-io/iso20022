// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentCamt02100106 struct {
	RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.021.001.06 RtrGnlBizInf"`
}

func (doc DocumentCamt02100106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt02400106 struct {
	ModfyStgOrdr ModifyStandingOrderV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.024.001.06 ModfyStgOrdr"`
}

func (doc DocumentCamt02400106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03100106 struct {
	RjctInvstgtn RejectInvestigationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.031.001.06 RjctInvstgtn"`
}

func (doc DocumentCamt03100106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt04700106 struct {
	RtrRsvatn ReturnReservationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.047.001.06 RtrRsvatn"`
}

func (doc DocumentCamt04700106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt05700106 struct {
	NtfctnToRcv NotificationToReceiveV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.057.001.06 NtfctnToRcv"`
}

func (doc DocumentCamt05700106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03300106 struct {
	ReqForDplct RequestForDuplicateV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.033.001.06 ReqForDplct"`
}

func (doc DocumentCamt03300106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt05800106 struct {
	NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.058.001.06 NtfctnToRcvCxlAdvc"`
}

func (doc DocumentCamt05800106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt05900106 struct {
	NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.06 NtfctnToRcvStsRpt"`
}

func (doc DocumentCamt05900106) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt03400106 struct {
	Dplct DuplicateV06 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.034.001.06 Dplct"`
}

func (doc DocumentCamt03400106) Validate() error {
	return utils.Validate(&doc)
}
