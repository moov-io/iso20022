// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt02800110 struct {
	XMLName                 *xml.Name                       `json:",omitempty"`
	Xmlns                   string                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                            `xml:",omitempty" json:",omitempty"`
	AddtlPmtInf             AdditionalPaymentInformationV10 `xml:"AddtlPmtInf"`
}

func (doc DocumentCamt02800110) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02800110) NameSpace() string {
	return utils.DocumentCamt02800110NameSpace
}

func (doc DocumentCamt02800110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AddtlPmtInf AdditionalPaymentInformationV10 `xml:"AddtlPmtInf"`
	}
	output.AddtlPmtInf = doc.AddtlPmtInf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt02900110 struct {
	XMLName                 *xml.Name                    `json:",omitempty"`
	Xmlns                   string                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                         `xml:",omitempty" json:",omitempty"`
	RsltnOfInvstgtn         ResolutionOfInvestigationV10 `xml:"RsltnOfInvstgtn"`
}

func (doc DocumentCamt02900110) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02900110) NameSpace() string {
	return utils.DocumentCamt02900110NameSpace
}

func (doc DocumentCamt02900110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RsltnOfInvstgtn ResolutionOfInvestigationV10 `xml:"RsltnOfInvstgtn"`
	}
	output.RsltnOfInvstgtn = doc.RsltnOfInvstgtn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
