// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v01

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentPain00700101 struct {
	MndtCpyReq MandateCopyRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtCpyReq"`
}

func (doc DocumentPain00700101) Validate() error {
	return utils.Validate(&doc)
}

type DocumentPain01800101 struct {
	MndtSspnsnReq MandateSuspensionRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 MndtSspnsnReq"`
}

func (doc DocumentPain01800101) Validate() error {
	return utils.Validate(&doc)
}
