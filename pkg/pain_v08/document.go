// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain01400108 struct {
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (doc DocumentPain01400108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPain01400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CdtrPmtActvtnReqStsRpt"`
	}
	output.CdtrPmtActvtnReqStsRpt = doc.CdtrPmtActvtnReqStsRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pain.014.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentPain01300108 struct {
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"CdtrPmtActvtnReq"`
}

func (doc DocumentPain01300108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPain01300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtrPmtActvtnReq"`
	}
	output.CdtrPmtActvtnReq = doc.CdtrPmtActvtnReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pain.013.001.08")
	return e.EncodeElement(&output, start)
}
