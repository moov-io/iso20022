// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt03600101 struct {
	XMLName               xml.Name
	Attrs                 []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchTermntnSwtch AccountSwitchTerminationSwitchV01 `xml:"AcctSwtchTermntnSwtch"`
}

func (doc DocumentAcmt03600101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03600101) NameSpace() string {
	return utils.DocumentAcmt03600101NameSpace
}

func (doc DocumentAcmt03600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName               xml.Name
		Attrs                 []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchTermntnSwtch AccountSwitchTerminationSwitchV01 `xml:"AcctSwtchTermntnSwtch"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
