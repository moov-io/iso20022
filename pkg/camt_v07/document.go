// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v07

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentCamt00300107 struct {
	GetAcct GetAccountV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.003.001.07 GetAcct"`
}

func (doc DocumentCamt00300107) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt00900107 struct {
	GetLmt GetLimitV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.009.001.07 GetLmt"`
}

func (doc DocumentCamt00900107) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01100107 struct {
	ModfyLmt ModifyLimitV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.011.001.07 ModfyLmt"`
}

func (doc DocumentCamt01100107) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01900107 struct {
	RtrBizDayInf ReturnBusinessDayInformationV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.019.001.07 RtrBizDayInf"`
}

func (doc DocumentCamt01900107) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt02300107 struct {
	BckpPmt BackupPaymentV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.023.001.07 BckpPmt"`
}

func (doc DocumentCamt02300107) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt01200107 struct {
	DelLmt DeleteLimitV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.012.001.07 DelLmt"`
}

func (doc DocumentCamt01200107) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt08700107 struct {
	ReqToModfyPmt RequestToModifyPaymentV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.07 ReqToModfyPmt"`
}

func (doc DocumentCamt08700107) Validate() error {
	return utils.Validate(&doc)
}
