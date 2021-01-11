// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v10

type DocumentPain00700110 struct {
	CstmrPmtRvsl CustomerPaymentReversalV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CstmrPmtRvsl"`
}

type DocumentPain00100110 struct {
	CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CstmrCdtTrfInitn"`
}
