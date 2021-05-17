// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt01800105 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	GetBizDayInf GetBusinessDayInformationV05 `xml:"GetBizDayInf"`
}

func (doc DocumentCamt01800105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01800105) NameSpace() string {
	return utils.DocumentCamt01800105NameSpace
}

func (doc DocumentCamt01800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		GetBizDayInf GetBusinessDayInformationV05 `xml:"GetBizDayInf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt02500105 struct {
	XMLName xml.Name
	Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	Rct     ReceiptV05   `xml:"Rct"`
}

func (doc DocumentCamt02500105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02500105) NameSpace() string {
	return utils.DocumentCamt02500105NameSpace
}

func (doc DocumentCamt02500105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
		Rct     ReceiptV05   `xml:"Rct"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt02600105 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	UblToApply UnableToApplyV05 `xml:"UblToApply"`
}

func (doc DocumentCamt02600105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02600105) NameSpace() string {
	return utils.DocumentCamt02600105NameSpace
}

func (doc DocumentCamt02600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
		UblToApply UnableToApplyV05 `xml:"UblToApply"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt02800105 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	AddtlPmtInf AdditionalPaymentInformationV05 `xml:"AddtlPmtInf"`
}

func (doc DocumentCamt02800105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02800105) NameSpace() string {
	return utils.DocumentCamt02800105NameSpace
}

func (doc DocumentCamt02800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		AddtlPmtInf AdditionalPaymentInformationV05 `xml:"AddtlPmtInf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03000105 struct {
	XMLName             xml.Name
	Attrs               []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	NtfctnOfCaseAssgnmt NotificationOfCaseAssignmentV05 `xml:"NtfctnOfCaseAssgnmt"`
}

func (doc DocumentCamt03000105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03000105) NameSpace() string {
	return utils.DocumentCamt03000105NameSpace
}

func (doc DocumentCamt03000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName             xml.Name
		Attrs               []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		NtfctnOfCaseAssgnmt NotificationOfCaseAssignmentV05 `xml:"NtfctnOfCaseAssgnmt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03500105 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
	PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"PrtryFrmtInvstgtn"`
}

func (doc DocumentCamt03500105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03500105) NameSpace() string {
	return utils.DocumentCamt03500105NameSpace
}

func (doc DocumentCamt03500105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
		PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"PrtryFrmtInvstgtn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03600105 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	DbtAuthstnRspn DebitAuthorisationResponseV05 `xml:"DbtAuthstnRspn"`
}

func (doc DocumentCamt03600105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03600105) NameSpace() string {
	return utils.DocumentCamt03600105NameSpace
}

func (doc DocumentCamt03600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		DbtAuthstnRspn DebitAuthorisationResponseV05 `xml:"DbtAuthstnRspn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03900105 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
	CaseStsRpt CaseStatusReportV05 `xml:"CaseStsRpt"`
}

func (doc DocumentCamt03900105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03900105) NameSpace() string {
	return utils.DocumentCamt03900105NameSpace
}

func (doc DocumentCamt03900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr        `xml:",any,attr,omitempty" json:",omitempty"`
		CaseStsRpt CaseStatusReportV05 `xml:"CaseStsRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt04600105 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
	GetRsvatn GetReservationV05 `xml:"GetRsvatn"`
}

func (doc DocumentCamt04600105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04600105) NameSpace() string {
	return utils.DocumentCamt04600105NameSpace
}

func (doc DocumentCamt04600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
		GetRsvatn GetReservationV05 `xml:"GetRsvatn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt04800105 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	ModfyRsvatn ModifyReservationV05 `xml:"ModfyRsvatn"`
}

func (doc DocumentCamt04800105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04800105) NameSpace() string {
	return utils.DocumentCamt04800105NameSpace
}

func (doc DocumentCamt04800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		ModfyRsvatn ModifyReservationV05 `xml:"ModfyRsvatn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt04900105 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	DelRsvatn DeleteReservationV05 `xml:"DelRsvatn"`
}

func (doc DocumentCamt04900105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04900105) NameSpace() string {
	return utils.DocumentCamt04900105NameSpace
}

func (doc DocumentCamt04900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		DelRsvatn DeleteReservationV05 `xml:"DelRsvatn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt05000105 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	LqdtyCdtTrf LiquidityCreditTransferV05 `xml:"LqdtyCdtTrf"`
}

func (doc DocumentCamt05000105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05000105) NameSpace() string {
	return utils.DocumentCamt05000105NameSpace
}

func (doc DocumentCamt05000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		LqdtyCdtTrf LiquidityCreditTransferV05 `xml:"LqdtyCdtTrf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt05100105 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr              `xml:",any,attr,omitempty" json:",omitempty"`
	LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"LqdtyDbtTrf"`
}

func (doc DocumentCamt05100105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05100105) NameSpace() string {
	return utils.DocumentCamt05100105NameSpace
}

func (doc DocumentCamt05100105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr              `xml:",any,attr,omitempty" json:",omitempty"`
		LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"LqdtyDbtTrf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt05600105 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV05 `xml:"FIToFIPmtCxlReq"`
}

func (doc DocumentCamt05600105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05600105) NameSpace() string {
	return utils.DocumentCamt05600105NameSpace
}

func (doc DocumentCamt05600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV05 `xml:"FIToFIPmtCxlReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt06000105 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	AcctRptgReq AccountReportingRequestV05 `xml:"AcctRptgReq"`
}

func (doc DocumentCamt06000105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt06000105) NameSpace() string {
	return utils.DocumentCamt06000105NameSpace
}

func (doc DocumentCamt06000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		AcctRptgReq AccountReportingRequestV05 `xml:"AcctRptgReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
