// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type AccountSwitchDetails1 struct {
	UnqRefNb      common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 UnqRefNb"`
	RtgUnqRefNb   common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 RtgUnqRefNb"`
	SwtchRcvdDtTm common.ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchRcvdDtTm,omitempty"`
	SwtchDt       common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchDt,omitempty"`
	SwtchTp       SwitchType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchTp"`
	SwtchSts      SwitchStatus1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SwtchSts,omitempty"`
	BalTrfWndw    BalanceTransferWindow1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 BalTrfWndw,omitempty"`
	Rspn          []ResponseDetails1         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 Rspn,omitempty"`
}

func (r AccountSwitchDetails1) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchTerminationSwitchV01 struct {
	MsgId         MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 SplmtryData,omitempty"`
}

func (r AccountSwitchTerminationSwitchV01) Validate() error {
	return utils.Validate(&r)
}

type MessageIdentification1 struct {
	Id      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 Id"`
	CreDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 CreDtTm"`
}

func (r MessageIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ResponseDetails1 struct {
	RspnCd    common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 RspnCd"`
	AddtlDtls common.Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 AddtlDtls,omitempty"`
}

func (r ResponseDetails1) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 Envlp"`
}

func (r SupplementaryData1) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

func (r SupplementaryDataEnvelope1) Validate() error {
	return utils.Validate(&r)
}
