// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentRemt00100104 struct {
	XMLName xml.Name
	Attrs   []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
	RmtAdvc RemittanceAdviceV04 `xml:"RmtAdvc"`
}

func (doc DocumentRemt00100104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentRemt00100104) NameSpace() string {
	return utils.DocumentRemt00100104NameSpace
}

func (doc DocumentRemt00100104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
		RmtAdvc RemittanceAdviceV04 `xml:"RmtAdvc"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
