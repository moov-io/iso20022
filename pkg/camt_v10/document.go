// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v10

type DocumentCamt02800110 struct {
	AddtlPmtInf AdditionalPaymentInformationV10 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.028.001.10 AddtlPmtInf"`
}

type DocumentCamt02900110 struct {
	RsltnOfInvstgtn ResolutionOfInvestigationV10 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.10 RsltnOfInvstgtn"`
}
