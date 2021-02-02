// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v10

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00400110 struct {
	Xmlns  string           `xml:"xmlns,attr"`
	PmtRtr PaymentReturnV10 `xml:"PmtRtr"`
}

func (doc DocumentPacs00400110) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00400110) NameSpace() string {
	return utils.DocumentPacs00400110NameSpace
}

func (doc DocumentPacs00400110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PmtRtr PaymentReturnV10 `xml:"PmtRtr"`
	}
	output.PmtRtr = doc.PmtRtr
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentPacs00700110 struct {
	Xmlns         string                   `xml:"xmlns,attr"`
	FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"FIToFIPmtRvsl"`
}

func (doc DocumentPacs00700110) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00700110) NameSpace() string {
	return utils.DocumentPacs00700110NameSpace
}

func (doc DocumentPacs00700110) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtRvsl FIToFIPaymentReversalV10 `xml:"FIToFIPmtRvsl"`
	}
	output.FIToFIPmtRvsl = doc.FIToFIPmtRvsl
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
