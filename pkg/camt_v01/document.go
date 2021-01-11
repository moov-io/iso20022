// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentCamt10100101 struct {
	CretLmt CreateLimitV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.101.001.01 CretLmt"`
}

func (doc DocumentCamt10100101) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt10200101 struct {
	CretStgOrdr CreateStandingOrderV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.102.001.01 CretStgOrdr"`
}

func (doc DocumentCamt10200101) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt10300101 struct {
	CretRsvatn CreateReservationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.103.001.01 CretRsvatn"`
}

func (doc DocumentCamt10300101) Validate() error {
	return utils.Validate(&doc)
}

type DocumentCamt10400101 struct {
	CretMmb CreateMemberV01 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.104.001.01 CretMmb"`
}

func (doc DocumentCamt10400101) Validate() error {
	return utils.Validate(&doc)
}
