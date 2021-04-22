// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt02100106 struct {
	Xmlns        string                              `xml:"xmlns,attr"`
	RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"RtrGnlBizInf"`
}

func (doc DocumentCamt02100106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02100106) NameSpace() string {
	return utils.DocumentCamt02100106NameSpace
}

func (doc DocumentCamt02100106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"RtrGnlBizInf"`
	}
	output.RtrGnlBizInf = doc.RtrGnlBizInf
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt02400106 struct {
	Xmlns        string                 `xml:"xmlns,attr"`
	ModfyStgOrdr ModifyStandingOrderV06 `xml:"ModfyStgOrdr"`
}

func (doc DocumentCamt02400106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02400106) NameSpace() string {
	return utils.DocumentCamt02400106NameSpace
}

func (doc DocumentCamt02400106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyStgOrdr ModifyStandingOrderV06 `xml:"ModfyStgOrdr"`
	}
	output.ModfyStgOrdr = doc.ModfyStgOrdr
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt02900106 struct {
	Xmlns           string                       `xml:"xmlns,attr"`
	RsltnOfInvstgtn ResolutionOfInvestigationV06 `xml:"RsltnOfInvstgtn"`
	Dplct           DuplicateV06                 `xml:"Dplct"`
}

func (doc DocumentCamt02900106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02900106) NameSpace() string {
	return utils.DocumentCamt02900106NameSpace
}

func (doc DocumentCamt02900106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RsltnOfInvstgtn ResolutionOfInvestigationV06 `xml:"RsltnOfInvstgtn"`
	}
	output.RsltnOfInvstgtn = doc.RsltnOfInvstgtn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt03100106 struct {
	Xmlns        string                 `xml:"xmlns,attr"`
	RjctInvstgtn RejectInvestigationV06 `xml:"RjctInvstgtn"`
}

func (doc DocumentCamt03100106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03100106) NameSpace() string {
	return utils.DocumentCamt03100106NameSpace
}

func (doc DocumentCamt03100106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RjctInvstgtn RejectInvestigationV06 `xml:"RjctInvstgtn"`
	}
	output.RjctInvstgtn = doc.RjctInvstgtn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt03300106 struct {
	Xmlns       string                 `xml:"xmlns,attr"`
	ReqForDplct RequestForDuplicateV06 `xml:"ReqForDplct"`
}

func (doc DocumentCamt03300106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03300106) NameSpace() string {
	return utils.DocumentCamt03300106NameSpace
}

func (doc DocumentCamt03300106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqForDplct RequestForDuplicateV06 `xml:"ReqForDplct"`
	}
	output.ReqForDplct = doc.ReqForDplct
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt03400106 struct {
	Xmlns string       `xml:"xmlns,attr"`
	Dplct DuplicateV06 `xml:"Dplct"`
}

func (doc DocumentCamt03400106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03400106) NameSpace() string {
	return utils.DocumentCamt03400106NameSpace
}

func (doc DocumentCamt03400106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		Dplct DuplicateV06 `xml:"Dplct"`
	}
	output.Dplct = doc.Dplct
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt04700106 struct {
	Xmlns     string               `xml:"xmlns,attr"`
	RtrRsvatn ReturnReservationV06 `xml:"RtrRsvatn"`
}

func (doc DocumentCamt04700106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04700106) NameSpace() string {
	return utils.DocumentCamt04700106NameSpace
}

func (doc DocumentCamt04700106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrRsvatn ReturnReservationV06 `xml:"RtrRsvatn"`
	}
	output.RtrRsvatn = doc.RtrRsvatn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05700106 struct {
	Xmlns       string                   `xml:"xmlns,attr"`
	NtfctnToRcv NotificationToReceiveV06 `xml:"NtfctnToRcv"`
}

func (doc DocumentCamt05700106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05700106) NameSpace() string {
	return utils.DocumentCamt05700106NameSpace
}

func (doc DocumentCamt05700106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnToRcv NotificationToReceiveV06 `xml:"NtfctnToRcv"`
	}
	output.NtfctnToRcv = doc.NtfctnToRcv
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05800106 struct {
	Xmlns              string                                     `xml:"xmlns,attr"`
	NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"NtfctnToRcvCxlAdvc"`
}

func (doc DocumentCamt05800106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05800106) NameSpace() string {
	return utils.DocumentCamt05800106NameSpace
}

func (doc DocumentCamt05800106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"NtfctnToRcvCxlAdvc"`
	}
	output.NtfctnToRcvCxlAdvc = doc.NtfctnToRcvCxlAdvc
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05900106 struct {
	Xmlns             string                               `xml:"xmlns,attr"`
	NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"NtfctnToRcvStsRpt"`
}

func (doc DocumentCamt05900106) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05900106) NameSpace() string {
	return utils.DocumentCamt05900106NameSpace
}

func (doc DocumentCamt05900106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"NtfctnToRcvStsRpt"`
	}
	output.NtfctnToRcvStsRpt = doc.NtfctnToRcvStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
