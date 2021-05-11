// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00700101 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr          `xml:",any,attr,omitempty" json:",omitempty"`
	MndtCpyReq MandateCopyRequestV01 `xml:"MndtCpyReq"`
}

func (doc DocumentPain00700101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00700101) NameSpace() string {
	return utils.DocumentPain00700101NameSpace
}

func (doc DocumentPain00700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr          `xml:",any,attr,omitempty" json:",omitempty"`
		MndtCpyReq MandateCopyRequestV01 `xml:"MndtCpyReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentPain01800101 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr                `xml:",any,attr,omitempty" json:",omitempty"`
	MndtSspnsnReq MandateSuspensionRequestV01 `xml:"MndtSspnsnReq"`
}

func (doc DocumentPain01800101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01800101) NameSpace() string {
	return utils.DocumentPain01800101NameSpace
}

func (doc DocumentPain01800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr                `xml:",any,attr,omitempty" json:",omitempty"`
		MndtSspnsnReq MandateSuspensionRequestV01 `xml:"MndtSspnsnReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
