// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v02

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type BusinessApplicationHeaderV02 struct {
	XMLName    string                        `xml:"AppHdr", json:"AppHdr"`
	Attr       *utils.Attr                   `xml:",attr,omitempty" json:",omitempty"`
	Xmlns      string                        `xml:"xmlns,attr,omitempty" json:",omitempty"`
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
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc BusinessApplicationHeaderV02) NameSpace() string {
	return utils.DocumentHead00100102NameSpace
}
