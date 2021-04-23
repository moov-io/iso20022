// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00700101 struct {
	XMLName                 *xml.Name             `json:",omitempty"`
	Xmlns                   string                `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                  `xml:",omitempty" json:",omitempty"`
	MndtCpyReq              MandateCopyRequestV01 `xml:"MndtCpyReq"`
}

func (doc DocumentPain00700101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00700101) NameSpace() string {
	return utils.DocumentPain00700101NameSpace
}

func (doc DocumentPain00700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtCpyReq MandateCopyRequestV01 `xml:"MndtCpyReq"`
	}
	output.MndtCpyReq = doc.MndtCpyReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentPain01800101 struct {
	XMLName                 *xml.Name                   `json:",omitempty"`
	Xmlns                   string                      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                        `xml:",omitempty" json:",omitempty"`
	MndtSspnsnReq           MandateSuspensionRequestV01 `xml:"MndtSspnsnReq"`
}

func (doc DocumentPain01800101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01800101) NameSpace() string {
	return utils.DocumentPain01800101NameSpace
}

func (doc DocumentPain01800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		MndtSspnsnReq MandateSuspensionRequestV01 `xml:"MndtSspnsnReq"`
	}
	output.MndtSspnsnReq = doc.MndtSspnsnReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
