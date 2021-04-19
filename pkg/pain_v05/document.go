// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00900105 struct {
	Xmlns        string                      `xml:"xmlns,attr"`
	MndtInitnReq MandateInitiationRequestV05 `xml:"MndtInitnReq"`
}

func (doc DocumentPain00900105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00900105) NameSpace() string {
	return utils.DocumentPain00900105NameSpace
}

func (doc DocumentPain00900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtInitnReq MandateInitiationRequestV05 `xml:"MndtInitnReq"`
	}
	output.MndtInitnReq = doc.MndtInitnReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain01000105 struct {
	Xmlns         string                     `xml:"xmlns,attr"`
	MndtAmdmntReq MandateAmendmentRequestV05 `xml:"MndtAmdmntReq"`
}

func (doc DocumentPain01000105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01000105) NameSpace() string {
	return utils.DocumentPain01000105NameSpace
}

func (doc DocumentPain01000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtAmdmntReq MandateAmendmentRequestV05 `xml:"MndtAmdmntReq"`
	}
	output.MndtAmdmntReq = doc.MndtAmdmntReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain01200105 struct {
	Xmlns          string                     `xml:"xmlns,attr"`
	MndtAccptncRpt MandateAcceptanceReportV05 `xml:"MndtAccptncRpt"`
}

func (doc DocumentPain01200105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01200105) NameSpace() string {
	return utils.DocumentPain01200105NameSpace
}

func (doc DocumentPain01200105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtAccptncRpt MandateAcceptanceReportV05 `xml:"MndtAccptncRpt"`
	}
	output.MndtAccptncRpt = doc.MndtAccptncRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain01100105 struct {
	Xmlns      string                        `xml:"xmlns,attr"`
	MndtCxlReq MandateCancellationRequestV05 `xml:"MndtCxlReq"`
}

func (doc DocumentPain01100105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01100105) NameSpace() string {
	return utils.DocumentPain01100105NameSpace
}

func (doc DocumentPain01100105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtCxlReq MandateCancellationRequestV05 `xml:"MndtCxlReq"`
	}
	output.MndtCxlReq = doc.MndtCxlReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain01300105 struct {
	Xmlns            string                              `xml:"xmlns,attr"`
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV05 `xml:"CdtrPmtActvtnReq"`
}

func (doc DocumentPain01300105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01300105) NameSpace() string {
	return utils.DocumentPain01300105NameSpace
}

func (doc DocumentPain01300105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CdtrPmtActvtnReq CreditorPaymentActivationRequestV05 `xml:"CdtrPmtActvtnReq"`
	}
	output.CdtrPmtActvtnReq = doc.CdtrPmtActvtnReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain01400105 struct {
	Xmlns                  string                                          `xml:"xmlns,attr"`
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV05 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (doc DocumentPain01400105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01400105) NameSpace() string {
	return utils.DocumentPain01400105NameSpace
}

func (doc DocumentPain01400105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV05 `xml:"CdtrPmtActvtnReqStsRpt"`
	}
	output.CdtrPmtActvtnReqStsRpt = doc.CdtrPmtActvtnReqStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
