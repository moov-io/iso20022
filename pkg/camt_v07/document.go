// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v07

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt00300107 struct {
	GetAcct GetAccountV07 `xml:"GetAcct"`
}

func (doc DocumentCamt00300107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00300107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetAcct GetAccountV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.003.001.07 GetAcct"`
	}
	output.GetAcct = doc.GetAcct
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.003.001.07")
	return e.EncodeElement(&output, start)
}

type DocumentCamt00900107 struct {
	GetLmt GetLimitV07 `xml:"GetLmt"`
}

func (doc DocumentCamt00900107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00900107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetLmt GetLimitV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.009.001.07 GetLmt"`
	}
	output.GetLmt = doc.GetLmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.009.001.07")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01100107 struct {
	ModfyLmt ModifyLimitV07 `xml:"ModfyLmt"`
}

func (doc DocumentCamt01100107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01100107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyLmt ModifyLimitV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.011.001.07 ModfyLmt"`
	}
	output.ModfyLmt = doc.ModfyLmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.011.001.07")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01900107 struct {
	RtrBizDayInf ReturnBusinessDayInformationV07 `xml:"RtrBizDayInf"`
}

func (doc DocumentCamt01900107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01900107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrBizDayInf ReturnBusinessDayInformationV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.019.001.07 RtrBizDayInf"`
	}
	output.RtrBizDayInf = doc.RtrBizDayInf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.019.001.07")
	return e.EncodeElement(&output, start)
}

type DocumentCamt02300107 struct {
	BckpPmt BackupPaymentV07 `xml:"BckpPmt"`
}

func (doc DocumentCamt02300107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02300107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BckpPmt BackupPaymentV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.023.001.07 BckpPmt"`
	}
	output.BckpPmt = doc.BckpPmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.023.001.07")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01200107 struct {
	DelLmt DeleteLimitV07 `xml:"DelLmt"`
}

func (doc DocumentCamt01200107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01200107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelLmt DeleteLimitV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.012.001.07 DelLmt"`
	}
	output.DelLmt = doc.DelLmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.012.001.07")
	return e.EncodeElement(&output, start)
}

type DocumentCamt08700107 struct {
	ReqToModfyPmt RequestToModifyPaymentV07 `xml:"ReqToModfyPmt"`
}

func (doc DocumentCamt08700107) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt08700107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToModfyPmt RequestToModifyPaymentV07 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.07 ReqToModfyPmt"`
	}
	output.ReqToModfyPmt = doc.ReqToModfyPmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.087.001.07")
	return e.EncodeElement(&output, start)
}
