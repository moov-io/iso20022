// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt01300104 struct {
	XMLName xml.Name
	Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	GetMmb  GetMemberV04 `xml:"GetMmb"`
}

func (doc DocumentCamt01300104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01300104) NameSpace() string {
	return utils.DocumentCamt01300104NameSpace
}

func (doc DocumentCamt01300104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
		GetMmb  GetMemberV04 `xml:"GetMmb"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt01400104 struct {
	XMLName xml.Name
	Attrs   []utils.Attr    `xml:",any,attr,omitempty" json:",omitempty"`
	RtrMmb  ReturnMemberV04 `xml:"RtrMmb"`
}

func (doc DocumentCamt01400104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01400104) NameSpace() string {
	return utils.DocumentCamt01400104NameSpace
}

func (doc DocumentCamt01400104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr    `xml:",any,attr,omitempty" json:",omitempty"`
		RtrMmb  ReturnMemberV04 `xml:"RtrMmb"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt01500104 struct {
	XMLName  xml.Name
	Attrs    []utils.Attr    `xml:",any,attr,omitempty" json:",omitempty"`
	ModfyMmb ModifyMemberV04 `xml:"ModfyMmb"`
}

func (doc DocumentCamt01500104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01500104) NameSpace() string {
	return utils.DocumentCamt01500104NameSpace
}

func (doc DocumentCamt01500104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName  xml.Name
		Attrs    []utils.Attr    `xml:",any,attr,omitempty" json:",omitempty"`
		ModfyMmb ModifyMemberV04 `xml:"ModfyMmb"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt01600104 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"GetCcyXchgRate"`
}

func (doc DocumentCamt01600104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01600104) NameSpace() string {
	return utils.DocumentCamt01600104NameSpace
}

func (doc DocumentCamt01600104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		GetCcyXchgRate GetCurrencyExchangeRateV04 `xml:"GetCcyXchgRate"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt01700104 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"RtrCcyXchgRate"`
}

func (doc DocumentCamt01700104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01700104) NameSpace() string {
	return utils.DocumentCamt01700104NameSpace
}

func (doc DocumentCamt01700104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		RtrCcyXchgRate ReturnCurrencyExchangeRateV04 `xml:"RtrCcyXchgRate"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt02000104 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
	GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"GetGnlBizInf"`
}

func (doc DocumentCamt02000104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02000104) NameSpace() string {
	return utils.DocumentCamt02000104NameSpace
}

func (doc DocumentCamt02000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
		GetGnlBizInf GetGeneralBusinessInformationV04 `xml:"GetGnlBizInf"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt03200104 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr            `xml:",any,attr,omitempty" json:",omitempty"`
	CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"CclCaseAssgnmt"`
}

func (doc DocumentCamt03200104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03200104) NameSpace() string {
	return utils.DocumentCamt03200104NameSpace
}

func (doc DocumentCamt03200104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr            `xml:",any,attr,omitempty" json:",omitempty"`
		CclCaseAssgnmt CancelCaseAssignmentV04 `xml:"CclCaseAssgnmt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt03800104 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	CaseStsRptReq CaseStatusReportRequestV04 `xml:"CaseStsRptReq"`
}

func (doc DocumentCamt03800104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03800104) NameSpace() string {
	return utils.DocumentCamt03800104NameSpace
}

func (doc DocumentCamt03800104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		CaseStsRptReq CaseStatusReportRequestV04 `xml:"CaseStsRptReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt07000104 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	RtrStgOrdr ReturnStandingOrderV04 `xml:"RtrStgOrdr"`
}

func (doc DocumentCamt07000104) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt07000104) NameSpace() string {
	return utils.DocumentCamt07000104NameSpace
}

func (doc DocumentCamt07000104) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
		RtrStgOrdr ReturnStandingOrderV04 `xml:"RtrStgOrdr"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
