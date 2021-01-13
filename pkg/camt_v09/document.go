// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v09

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentCamt05500109 struct {
	CstmrPmtCxlReq CustomerPaymentCancellationRequestV09 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.055.001.09 CstmrPmtCxlReq"`
}

func (doc DocumentCamt05500109) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt05600109 struct {
	FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV09 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.09 FIToFIPmtCxlReq"`
}

func (doc DocumentCamt05600109) Validate() error {
	return utils.Validate(&doc)
}
