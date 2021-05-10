// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v02

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type BusinessApplicationHeaderV02 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	CharSet    string                        `xml:"CharSet,omitempty" json:",omitempty"`
	Fr         Party44Choice                 `xml:"Fr"`
	To         Party44Choice                 `xml:"To"`
	BizMsgIdr  common.Max35Text              `xml:"BizMsgIdr"`
	MsgDefIdr  common.Max35Text              `xml:"MsgDefIdr"`
	BizSvc     *common.Max35Text             `xml:"BizSvc,omitempty" json:",omitempty"`
	MktPrctc   *ImplementationSpecification1 `xml:"MktPrctc,omitempty" json:",omitempty"`
	CreDt      common.ISODateTime            `xml:"CreDt"`
	BizPrcgDt  *common.ISODateTime           `xml:"BizPrcgDt,omitempty" json:",omitempty"`
	CpyDplct   *common.CopyDuplicate1Code    `xml:"CpyDplct,omitempty" json:",omitempty"`
	PssblDplct bool                          `xml:"PssblDplct,omitempty" json:",omitempty"`
	Prty       string                        `xml:"Prty,omitempty" json:",omitempty"`
	Sgntr      *SignatureEnvelope            `xml:"Sgntr,omitempty" json:",omitempty"`
	Rltd       []BusinessApplicationHeader5  `xml:"Rltd,omitempty" json:",omitempty"`
}

func (doc BusinessApplicationHeaderV02) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc BusinessApplicationHeaderV02) NameSpace() string {
	return utils.DocumentHead00100102NameSpace
}

func (doc BusinessApplicationHeaderV02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		CharSet    string                        `xml:"CharSet,omitempty" json:",omitempty"`
		Fr         Party44Choice                 `xml:"Fr"`
		To         Party44Choice                 `xml:"To"`
		BizMsgIdr  common.Max35Text              `xml:"BizMsgIdr"`
		MsgDefIdr  common.Max35Text              `xml:"MsgDefIdr"`
		BizSvc     *common.Max35Text             `xml:"BizSvc,omitempty" json:",omitempty"`
		MktPrctc   *ImplementationSpecification1 `xml:"MktPrctc,omitempty" json:",omitempty"`
		CreDt      common.ISODateTime            `xml:"CreDt"`
		BizPrcgDt  *common.ISODateTime           `xml:"BizPrcgDt,omitempty" json:",omitempty"`
		CpyDplct   *common.CopyDuplicate1Code    `xml:"CpyDplct,omitempty" json:",omitempty"`
		PssblDplct bool                          `xml:"PssblDplct,omitempty" json:",omitempty"`
		Prty       string                        `xml:"Prty,omitempty" json:",omitempty"`
		Sgntr      *SignatureEnvelope            `xml:"Sgntr,omitempty" json:",omitempty"`
		Rltd       []BusinessApplicationHeader5  `xml:"Rltd,omitempty" json:",omitempty"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}
