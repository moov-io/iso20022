// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v01

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type BusinessApplicationHeaderV01 struct {
	XMLName    string                       `xml:"AppHdr", json:"AppHdr"`
	Attr       *utils.Attr                  `xml:",attr,omitempty" json:",omitempty"`
	Xmlns      string                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
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
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc BusinessApplicationHeaderV01) NameSpace() string {
	return utils.DocumentHead00100101NameSpace
}
