// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v01

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type BusinessApplicationHeaderV01 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	CharSet    string                       `xml:"CharSet,omitempty" json:",omitempty"`
	Fr         Party9Choice                 `xml:"Fr"`
	To         Party9Choice                 `xml:"To"`
	BizMsgIdr  common.Max35Text             `xml:"BizMsgIdr"`
	MsgDefIdr  common.Max35Text             `xml:"MsgDefIdr"`
	BizSvc     *common.Max35Text            `xml:"BizSvc,omitempty" json:",omitempty"`
	CreDt      common.ISONormalisedDateTime `xml:"CreDt"`
	CpyDplct   *common.CopyDuplicate1Code   `xml:"CpyDplct,omitempty" json:",omitempty"`
	PssblDplct bool                         `xml:"PssblDplct,omitempty" json:",omitempty"`
	Prty       string                       `xml:"Prty,omitempty" json:",omitempty"`
	Sgntr      *SignatureEnvelope           `xml:"Sgntr,omitempty" json:",omitempty"`
	Rltd       *BusinessApplicationHeader1  `xml:"Rltd,omitempty" json:",omitempty"`
}

func (doc BusinessApplicationHeaderV01) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc BusinessApplicationHeaderV01) NameSpace() string {
	return utils.DocumentHead00100101NameSpace
}

func (doc BusinessApplicationHeaderV01) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		CharSet    string                       `xml:"CharSet,omitempty" json:",omitempty"`
		Fr         Party9Choice                 `xml:"Fr"`
		To         Party9Choice                 `xml:"To"`
		BizMsgIdr  common.Max35Text             `xml:"BizMsgIdr"`
		MsgDefIdr  common.Max35Text             `xml:"MsgDefIdr"`
		BizSvc     *common.Max35Text            `xml:"BizSvc,omitempty" json:",omitempty"`
		CreDt      common.ISONormalisedDateTime `xml:"CreDt"`
		CpyDplct   *common.CopyDuplicate1Code   `xml:"CpyDplct,omitempty" json:",omitempty"`
		PssblDplct bool                         `xml:"PssblDplct,omitempty" json:",omitempty"`
		Prty       string                       `xml:"Prty,omitempty" json:",omitempty"`
		Sgntr      *SignatureEnvelope           `xml:"Sgntr,omitempty" json:",omitempty"`
		Rltd       *BusinessApplicationHeader1  `xml:"Rltd,omitempty" json:",omitempty"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *BusinessApplicationHeaderV01) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *BusinessApplicationHeaderV01) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *BusinessApplicationHeaderV01) InspectDocument() interface{} {
	return nil
}
