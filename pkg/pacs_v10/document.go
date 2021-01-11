// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v10

type DocumentPacs00400110 struct {
	PmtRtr PaymentReturnV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PmtRtr"`
}

type DocumentPacs00700110 struct {
	FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.007.001.10 FIToFIPmtRvsl"`
}
