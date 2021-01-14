// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v11

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentPain00200111 struct {
	CstmrPmtStsRpt CustomerPaymentStatusReportV11 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CstmrPmtStsRpt"`
}

func (doc DocumentPain00200111) Validate() error {
	return utils.Validate(&doc)
}
