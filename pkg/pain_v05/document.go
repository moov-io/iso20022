// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain00900105 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr                `xml:",any,attr,omitempty" json:",omitempty"`
	MndtInitnReq MandateInitiationRequestV05 `xml:"MndtInitnReq"`
}

func (doc DocumentPain00900105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain00900105) NameSpace() string {
	return utils.DocumentPain00900105NameSpace
}

func (doc DocumentPain00900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr                `xml:",any,attr,omitempty" json:",omitempty"`
		MndtInitnReq MandateInitiationRequestV05 `xml:"MndtInitnReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain00900105) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain00900105) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentPain01000105 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	MndtAmdmntReq MandateAmendmentRequestV05 `xml:"MndtAmdmntReq"`
}

func (doc DocumentPain01000105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01000105) NameSpace() string {
	return utils.DocumentPain01000105NameSpace
}

func (doc DocumentPain01000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		MndtAmdmntReq MandateAmendmentRequestV05 `xml:"MndtAmdmntReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01000105) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01000105) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentPain01200105 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	MndtAccptncRpt MandateAcceptanceReportV05 `xml:"MndtAccptncRpt"`
}

func (doc DocumentPain01200105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01200105) NameSpace() string {
	return utils.DocumentPain01200105NameSpace
}

func (doc DocumentPain01200105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		MndtAccptncRpt MandateAcceptanceReportV05 `xml:"MndtAccptncRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01200105) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01200105) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentPain01100105 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	MndtCxlReq MandateCancellationRequestV05 `xml:"MndtCxlReq"`
}

func (doc DocumentPain01100105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01100105) NameSpace() string {
	return utils.DocumentPain01100105NameSpace
}

func (doc DocumentPain01100105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		MndtCxlReq MandateCancellationRequestV05 `xml:"MndtCxlReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01100105) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01100105) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentPain01300105 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV05 `xml:"CdtrPmtActvtnReq"`
}

func (doc DocumentPain01300105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01300105) NameSpace() string {
	return utils.DocumentPain01300105NameSpace
}

func (doc DocumentPain01300105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		CdtrPmtActvtnReq CreditorPaymentActivationRequestV05 `xml:"CdtrPmtActvtnReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01300105) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01300105) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentPain01400105 struct {
	XMLName                xml.Name
	Attrs                  []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV05 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (doc DocumentPain01400105) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01400105) NameSpace() string {
	return utils.DocumentPain01400105NameSpace
}

func (doc DocumentPain01400105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName                xml.Name
		Attrs                  []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
		CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV05 `xml:"CdtrPmtActvtnReqStsRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01400105) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01400105) GetAttrs() []utils.Attr {
	return doc.Attrs
}
