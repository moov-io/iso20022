// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentPain00900105 struct {
	MndtInitnReq MandateInitiationRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtInitnReq"`
}

func (doc DocumentPain00900105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPain01000105 struct {
	MndtAmdmntReq MandateAmendmentRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtAmdmntReq"`
}

func (doc DocumentPain01000105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPain01200105 struct {
	MndtAccptncRpt MandateAcceptanceReportV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 MndtAccptncRpt"`
}

func (doc DocumentPain01200105) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPain01100105 struct {
	MndtCxlReq MandateCancellationRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 MndtCxlReq"`
}

func (doc DocumentPain01100105) Validate() error {
	return utils.Validate(&doc)
}
