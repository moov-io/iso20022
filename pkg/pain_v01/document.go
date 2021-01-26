// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00700101 struct {
	MndtCpyReq MandateCopyRequestV01 `xml:"MndtCpyReq"`
}

func (doc DocumentPain00700101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPain00700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtCpyReq MandateCopyRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtCpyReq"`
	}
	output.MndtCpyReq = doc.MndtCpyReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pain.017.001.01")
	return e.EncodeElement(&output, start)
}

type DocumentPain01800101 struct {
	MndtSspnsnReq MandateSuspensionRequestV01 `xml:"MndtSspnsnReq"`
}

func (doc DocumentPain01800101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentPain01800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtSspnsnReq MandateSuspensionRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 MndtSspnsnReq"`
	}
	output.MndtSspnsnReq = doc.MndtSspnsnReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:pain.018.001.01")
	return e.EncodeElement(&output, start)
}
