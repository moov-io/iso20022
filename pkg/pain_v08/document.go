// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain01400108 struct {
	Xmlns                  string                                          `xml:"xmlns,attr"`
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (doc DocumentPain01400108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01400108) NameSpace() string {
	return utils.DocumentPain01400108NameSpace
}

func (doc DocumentPain01400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"CdtrPmtActvtnReqStsRpt"`
	}
	output.CdtrPmtActvtnReqStsRpt = doc.CdtrPmtActvtnReqStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPain01300108 struct {
	Xmlns            string                              `xml:"xmlns,attr"`
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"CdtrPmtActvtnReq"`
}

func (doc DocumentPain01300108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01300108) NameSpace() string {
	return utils.DocumentPain01300108NameSpace
}

func (doc DocumentPain01300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"CdtrPmtActvtnReq"`
	}
	output.CdtrPmtActvtnReq = doc.CdtrPmtActvtnReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
