// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v04

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentPacs01000104 struct {
	FIDrctDbt FinancialInstitutionDirectDebitV04 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.010.001.04 FIDrctDbt"`
}

func (doc DocumentPacs01000104) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPacs02800104 struct {
	FIToFIPmtStsReq FIToFIPaymentStatusRequestV04 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04 FIToFIPmtStsReq"`
}

func (doc DocumentPacs02800104) Validate() error {
	return utils.Validate(&doc)
}
