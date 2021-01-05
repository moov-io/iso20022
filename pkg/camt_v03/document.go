// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

type DocumentCamt06900103 struct {
	GetStgOrdr GetStandingOrderV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.069.001.03 GetStgOrdr"`
}

type DocumentCamt07100103 struct {
	DelStgOrdr DeleteStandingOrderV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.071.001.03 DelStgOrdr"`
}

type DocumentCamt08600103 struct {
	BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.03 BkSvcsBllgStmt"`
}
