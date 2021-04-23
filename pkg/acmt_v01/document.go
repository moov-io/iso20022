// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt03600101 struct {
	XMLName                 *xml.Name                         `json:",omitempty"`
	Xmlns                   string                            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                              `xml:",omitempty" json:",omitempty"`
	AcctSwtchTermntnSwtch   AccountSwitchTerminationSwitchV01 `xml:"AcctSwtchTermntnSwtch"`
}

func (doc DocumentAcmt03600101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03600101) NameSpace() string {
	return utils.DocumentAcmt03600101NameSpace
}

func (doc DocumentAcmt03600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchTermntnSwtch AccountSwitchTerminationSwitchV01 `xml:"AcctSwtchTermntnSwtch"`
	}

	output.AcctSwtchTermntnSwtch = doc.AcctSwtchTermntnSwtch
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
