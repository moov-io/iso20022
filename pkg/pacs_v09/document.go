// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v09

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentPacs00800109 struct {
	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV09 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FIToFICstmrCdtTrf"`
}

func (doc DocumentPacs00800109) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPacs00900109 struct {
	FICdtTrf FinancialInstitutionCreditTransferV09 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.09 FICdtTrf"`
}

func (doc DocumentPacs00900109) Validate() error {
	return utils.Validate(&doc)
}
