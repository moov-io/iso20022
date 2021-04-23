// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v09

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00800109 struct {
	XMLName                 *xml.Name                        `json:",omitempty"`
	Xmlns                   string                           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                             `xml:",omitempty" json:",omitempty"`
	CstmrDrctDbtInitn       CustomerDirectDebitInitiationV09 `xml:"CstmrDrctDbtInitn"`
}

func (doc DocumentPain00800109) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00800109) NameSpace() string {
	return utils.DocumentPain00800109NameSpace
}

func (doc DocumentPain00800109) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CstmrDrctDbtInitn CustomerDirectDebitInitiationV09 `xml:"CstmrDrctDbtInitn"`
	}
	output.CstmrDrctDbtInitn = doc.CstmrDrctDbtInitn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
