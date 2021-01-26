// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentReda06600101 struct {
	ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
}

func (doc DocumentReda06600101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda06600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 ReqToPayCdtrEnrlmntReq"`
	}
	output.ReqToPayCdtrEnrlmntReq = doc.ReqToPayCdtrEnrlmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.066.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda07200101 struct {
	ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
}

func (doc DocumentReda07200101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda07200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 ReqToPayDbtrActvtnCxlReq"`
	}
	output.ReqToPayDbtrActvtnCxlReq = doc.ReqToPayDbtrActvtnCxlReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.072.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda06700101 struct {
	ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
}

func (doc DocumentReda06700101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda06700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 ReqToPayCdtrEnrlmntAmdmntReq"`
	}
	output.ReqToPayCdtrEnrlmntAmdmntReq = doc.ReqToPayCdtrEnrlmntAmdmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.067.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda07300101 struct {
	ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
}

func (doc DocumentReda07300101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda07300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 ReqToPayDbtrActvtnStsRpt"`
	}
	output.ReqToPayDbtrActvtnStsRpt = doc.ReqToPayDbtrActvtnStsRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.073.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda07100101 struct {
	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
}

func (doc DocumentReda07100101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda07100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ReqToPayDbtrActvtnAmdmntReq"`
	}
	output.ReqToPayDbtrActvtnAmdmntReq = doc.ReqToPayDbtrActvtnAmdmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.071.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda07000101 struct {
	ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
}

func (doc DocumentReda07000101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda07000101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 ReqToPayDbtrActvtnReq"`
	}
	output.ReqToPayDbtrActvtnReq = doc.ReqToPayDbtrActvtnReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.070.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda06800101 struct {
	ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
}

func (doc DocumentReda06800101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda06800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 ReqToPayCdtrEnrlmntCxlReq"`
	}
	output.ReqToPayCdtrEnrlmntCxlReq = doc.ReqToPayCdtrEnrlmntCxlReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.068.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentReda06000101 struct {
	ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
}

func (doc DocumentReda06000101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentReda06000101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 ReqToPayCdtrEnrlmntStsRpt"`
	}
	output.ReqToPayCdtrEnrlmntStsRpt = doc.ReqToPayCdtrEnrlmntStsRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:reda.069.001.01")
	return e.EncodeElement(&output, start)
}
