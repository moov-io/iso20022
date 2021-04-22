// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentReda06600101 struct {
	Xmlns                  string                                  `xml:"xmlns,attr"`
	ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
}

func (doc DocumentReda06600101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06600101) NameSpace() string {
	return utils.DocumentReda06600101NameSpace
}

func (doc DocumentReda06600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
	}
	output.ReqToPayCdtrEnrlmntReq = doc.ReqToPayCdtrEnrlmntReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda06700101 struct {
	Xmlns                        string                                           `xml:"xmlns,attr"`
	ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
}

func (doc DocumentReda06700101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06700101) NameSpace() string {
	return utils.DocumentReda06700101NameSpace
}

func (doc DocumentReda06700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
	}
	output.ReqToPayCdtrEnrlmntAmdmntReq = doc.ReqToPayCdtrEnrlmntAmdmntReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda06800101 struct {
	Xmlns                     string                                              `xml:"xmlns,attr"`
	ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
}

func (doc DocumentReda06800101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06800101) NameSpace() string {
	return utils.DocumentReda06800101NameSpace
}

func (doc DocumentReda06800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
	}
	output.ReqToPayCdtrEnrlmntCxlReq = doc.ReqToPayCdtrEnrlmntCxlReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda06900101 struct {
	Xmlns                     string                                       `xml:"xmlns,attr"`
	ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
}

func (doc DocumentReda06900101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06900101) NameSpace() string {
	return utils.DocumentReda06900101NameSpace
}

func (doc DocumentReda06900101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
	}
	output.ReqToPayCdtrEnrlmntStsRpt = doc.ReqToPayCdtrEnrlmntStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda07000101 struct {
	Xmlns                 string                                 `xml:"xmlns,attr"`
	ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
}

func (doc DocumentReda07000101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07000101) NameSpace() string {
	return utils.DocumentReda07000101NameSpace
}

func (doc DocumentReda07000101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
	}
	output.ReqToPayDbtrActvtnReq = doc.ReqToPayDbtrActvtnReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda07100101 struct {
	Xmlns                       string                                          `xml:"xmlns,attr"`
	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
}

func (doc DocumentReda07100101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07100101) NameSpace() string {
	return utils.DocumentReda07100101NameSpace
}

func (doc DocumentReda07100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
	}
	output.ReqToPayDbtrActvtnAmdmntReq = doc.ReqToPayDbtrActvtnAmdmntReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda07200101 struct {
	Xmlns                    string                                             `xml:"xmlns,attr"`
	ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
}

func (doc DocumentReda07200101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07200101) NameSpace() string {
	return utils.DocumentReda07200101NameSpace
}

func (doc DocumentReda07200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
	}
	output.ReqToPayDbtrActvtnCxlReq = doc.ReqToPayDbtrActvtnCxlReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentReda07300101 struct {
	Xmlns                    string                                      `xml:"xmlns,attr"`
	ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
}

func (doc DocumentReda07300101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07300101) NameSpace() string {
	return utils.DocumentReda07300101NameSpace
}

func (doc DocumentReda07300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
	}
	output.ReqToPayDbtrActvtnStsRpt = doc.ReqToPayDbtrActvtnStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
