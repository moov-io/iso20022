// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt03600101 struct {
	XMLName               xml.Name                          `xml:"Document" json:"-"`
	AcctSwtchTermntnSwtch AccountSwitchTerminationSwitchV01 `xml:"AcctSwtchTermntnSwtch"`
}

func (doc DocumentAcmt03600101) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchTermntnSwtch AccountSwitchTerminationSwitchV01 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01 AcctSwtchTermntnSwtch"`
	}
	output.AcctSwtchTermntnSwtch = doc.AcctSwtchTermntnSwtch
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01")
	return e.EncodeElement(&output, start)
}
