// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt03500103 struct {
	XMLName                 *xml.Name                         `json:",omitempty"`
	Xmlns                   string                            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                              `xml:",omitempty" json:",omitempty"`
	PrtryFrmtInvstgtn       ProprietaryFormatInvestigationV03 `xml:"PrtryFrmtInvstgtn"`
}

func (doc DocumentCamt03500103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03500103) NameSpace() string {
	return utils.DocumentCamt03500103NameSpace
}

func (doc DocumentCamt03500103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PrtryFrmtInvstgtn ProprietaryFormatInvestigationV03 `xml:"PrtryFrmtInvstgtn"`
	}
	output.PrtryFrmtInvstgtn = doc.PrtryFrmtInvstgtn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt06900103 struct {
	XMLName                 *xml.Name           `json:",omitempty"`
	Xmlns                   string              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                `xml:",omitempty" json:",omitempty"`
	GetStgOrdr              GetStandingOrderV03 `xml:"GetStgOrdr"`
}

func (doc DocumentCamt06900103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt06900103) NameSpace() string {
	return utils.DocumentCamt06900103NameSpace
}

func (doc DocumentCamt06900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetStgOrdr GetStandingOrderV03 `xml:"GetStgOrdr"`
	}
	output.GetStgOrdr = doc.GetStgOrdr
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt07100103 struct {
	XMLName                 *xml.Name              `json:",omitempty"`
	Xmlns                   string                 `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                   `xml:",omitempty" json:",omitempty"`
	DelStgOrdr              DeleteStandingOrderV03 `xml:"DelStgOrdr"`
}

func (doc DocumentCamt07100103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt07100103) NameSpace() string {
	return utils.DocumentCamt07100103NameSpace
}

func (doc DocumentCamt07100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelStgOrdr DeleteStandingOrderV03 `xml:"DelStgOrdr"`
	}
	output.DelStgOrdr = doc.DelStgOrdr
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt08600103 struct {
	XMLName                 *xml.Name                       `json:",omitempty"`
	Xmlns                   string                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                            `xml:",omitempty" json:",omitempty"`
	BkSvcsBllgStmt          BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
}

func (doc DocumentCamt08600103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt08600103) NameSpace() string {
	return utils.DocumentCamt08600103NameSpace
}

func (doc DocumentCamt08600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkSvcsBllgStmt BankServicesBillingStatementV03 `xml:"BkSvcsBllgStmt"`
	}
	output.BkSvcsBllgStmt = doc.BkSvcsBllgStmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
