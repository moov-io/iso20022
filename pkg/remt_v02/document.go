// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentRemt00200102 struct {
	Xmlns       string                      `xml:"xmlns,attr"`
	RmtLctnAdvc RemittanceLocationAdviceV02 `xml:"RmtLctnAdvc"`
}

func (doc DocumentRemt00200102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentRemt00200102) NameSpace() string {
	return utils.DocumentRemt00200102NameSpace
}

func (doc DocumentRemt00200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RmtLctnAdvc RemittanceLocationAdviceV02 `xml:"RmtLctnAdvc"`
	}
	output.RmtLctnAdvc = doc.RmtLctnAdvc
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentRemt00100102 struct {
	Xmlns   string              `xml:"xmlns,attr"`
	RmtAdvc RemittanceAdviceV02 `xml:"RmtAdvc"`
}

func (doc DocumentRemt00100102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentRemt00100102) NameSpace() string {
	return utils.DocumentRemt00100102NameSpace
}

func (doc DocumentRemt00100102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RmtAdvc RemittanceAdviceV02 `xml:"RmtAdvc"`
	}
	output.RmtAdvc = doc.RmtAdvc
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
