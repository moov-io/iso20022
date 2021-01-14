// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v04

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentRemt00100104 struct {
	RmtAdvc RemittanceAdviceV04 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtAdvc"`
}

func (doc DocumentRemt00100104) Validate() error {
	return utils.Validate(&doc)
}
