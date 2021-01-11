// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import "github.com/moov-io/iso20022/pkg/utils"

type DocumentAuth001001V01 struct {
	InfReqOpng InformationRequestOpeningV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.001.001.01 InfReqOpng"`
}

func (doc DocumentAuth001001V01) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth002001V01 struct {
	InfReqRspn InformationRequestResponseV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.002.001.01 InfReqRspn"`
}

func (doc DocumentAuth002001V01) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAuth003001V01 struct {
	InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 InfReqStsChngNtfctn"`
}

func (doc DocumentAuth003001V01) Validate() error {
	return utils.Validate(&doc)
}
