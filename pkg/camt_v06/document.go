// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt02100106 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"RtrGnlBizInf"`
}

func (doc DocumentCamt02100106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02100106) NameSpace() string {
	return utils.DocumentCamt02100106NameSpace
}

func (doc DocumentCamt02100106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		RtrGnlBizInf ReturnGeneralBusinessInformationV06 `xml:"RtrGnlBizInf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt02400106 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	ModfyStgOrdr ModifyStandingOrderV06 `xml:"ModfyStgOrdr"`
}

func (doc DocumentCamt02400106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02400106) NameSpace() string {
	return utils.DocumentCamt02400106NameSpace
}

func (doc DocumentCamt02400106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
		ModfyStgOrdr ModifyStandingOrderV06 `xml:"ModfyStgOrdr"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt02900106 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	RsltnOfInvstgtn ResolutionOfInvestigationV06 `xml:"RsltnOfInvstgtn"`
}

func (doc DocumentCamt02900106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02900106) NameSpace() string {
	return utils.DocumentCamt02900106NameSpace
}

func (doc DocumentCamt02900106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		RsltnOfInvstgtn ResolutionOfInvestigationV06 `xml:"RsltnOfInvstgtn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03100106 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	RjctInvstgtn RejectInvestigationV06 `xml:"RjctInvstgtn"`
}

func (doc DocumentCamt03100106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03100106) NameSpace() string {
	return utils.DocumentCamt03100106NameSpace
}

func (doc DocumentCamt03100106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
		RjctInvstgtn RejectInvestigationV06 `xml:"RjctInvstgtn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03300106 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	ReqForDplct RequestForDuplicateV06 `xml:"ReqForDplct"`
}

func (doc DocumentCamt03300106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03300106) NameSpace() string {
	return utils.DocumentCamt03300106NameSpace
}

func (doc DocumentCamt03300106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
		ReqForDplct RequestForDuplicateV06 `xml:"ReqForDplct"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt03400106 struct {
	XMLName xml.Name
	Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	Dplct   DuplicateV06 `xml:"Dplct"`
}

func (doc DocumentCamt03400106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03400106) NameSpace() string {
	return utils.DocumentCamt03400106NameSpace
}

func (doc DocumentCamt03400106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
		Dplct   DuplicateV06 `xml:"Dplct"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt04700106 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	RtrRsvatn ReturnReservationV06 `xml:"RtrRsvatn"`
}

func (doc DocumentCamt04700106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04700106) NameSpace() string {
	return utils.DocumentCamt04700106NameSpace
}

func (doc DocumentCamt04700106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		RtrRsvatn ReturnReservationV06 `xml:"RtrRsvatn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt05700106 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
	NtfctnToRcv NotificationToReceiveV06 `xml:"NtfctnToRcv"`
}

func (doc DocumentCamt05700106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05700106) NameSpace() string {
	return utils.DocumentCamt05700106NameSpace
}

func (doc DocumentCamt05700106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
		NtfctnToRcv NotificationToReceiveV06 `xml:"NtfctnToRcv"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt05800106 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                               `xml:",any,attr,omitempty" json:",omitempty"`
	NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"NtfctnToRcvCxlAdvc"`
}

func (doc DocumentCamt05800106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05800106) NameSpace() string {
	return utils.DocumentCamt05800106NameSpace
}

func (doc DocumentCamt05800106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                               `xml:",any,attr,omitempty" json:",omitempty"`
		NtfctnToRcvCxlAdvc NotificationToReceiveCancellationAdviceV06 `xml:"NtfctnToRcvCxlAdvc"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

type DocumentCamt05900106 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                         `xml:",any,attr,omitempty" json:",omitempty"`
	NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"NtfctnToRcvStsRpt"`
}

func (doc DocumentCamt05900106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05900106) NameSpace() string {
	return utils.DocumentCamt05900106NameSpace
}

func (doc DocumentCamt05900106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                         `xml:",any,attr,omitempty" json:",omitempty"`
		NtfctnToRcvStsRpt NotificationToReceiveStatusReportV06 `xml:"NtfctnToRcvStsRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}
