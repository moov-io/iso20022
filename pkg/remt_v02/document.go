// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v02

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentRemt00200102 struct {
	RmtLctnAdvc RemittanceLocationAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 RmtLctnAdvc"`
}

func (doc DocumentRemt00200102) Validate() error {
	return utils.Validate(&doc)
}
