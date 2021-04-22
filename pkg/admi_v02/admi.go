// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v02

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type Event2 struct {
	EvtCd    common.Max4AlphaNumericText `xml:"EvtCd"`
	EvtParam []common.Max35Text          `xml:"EvtParam,omitempty" json:",omitempty"`
	EvtDesc  *common.Max1000Text         `xml:"EvtDesc,omitempty" json:",omitempty"`
	EvtTm    *common.ISODateTime         `xml:"EvtTm,omitempty" json:",omitempty"`
}

func (r Event2) Validate() error {
	return utils.Validate(&r)
}

type SystemEventNotificationV02 struct {
	EvtInf Event2 `xml:"EvtInf"`
}

func (r SystemEventNotificationV02) Validate() error {
	return utils.Validate(&r)
}

type RequestDetails3 struct {
	Tp  common.Max35Text  `xml:"Tp"`
	Key *common.Max35Text `xml:"Key,omitempty" json:",omitempty"`
}

func (r RequestDetails3) Validate() error {
	return utils.Validate(&r)
}

type StaticDataRequestV02 struct {
	MsgId       common.Max35Text               `xml:"MsgId"`
	SttlmSsnIdr *common.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty" json:",omitempty"`
	DataReqDtls RequestDetails3                `xml:"DataReqDtls"`
	SplmtryData []SupplementaryData1           `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r StaticDataRequestV02) Validate() error {
	return utils.Validate(&r)
}

type ReportParameter1 struct {
	Nm  common.Max70Text  `xml:"Nm"`
	Val common.Max350Text `xml:"Val"`
}

func (r ReportParameter1) Validate() error {
	return utils.Validate(&r)
}

type RequestDetails4 struct {
	Key     common.Max35Text   `xml:"Key"`
	RptData []ReportParameter1 `xml:"RptData,omitempty"`
}

func (r RequestDetails4) Validate() error {
	return utils.Validate(&r)
}

type RequestDetails5 struct {
	Tp     common.Max35Text  `xml:"Tp"`
	ReqRef common.Max35Text  `xml:"ReqRef"`
	RptKey []RequestDetails4 `xml:"RptKey"`
}

func (r RequestDetails5) Validate() error {
	return utils.Validate(&r)
}

type StaticDataReportV02 struct {
	MsgId       common.Max35Text               `xml:"MsgId"`
	SttlmSsnIdr *common.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`
	RptDtls     RequestDetails5                `xml:"RptDtls"`
	SplmtryData []SupplementaryData1           `xml:"SplmtryData,omitempty"`
}

func (r StaticDataReportV02) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
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
