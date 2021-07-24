// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPain01400108 struct {
	XMLName                xml.Name
	Attrs                  []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
	CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"CdtrPmtActvtnReqStsRpt"`
}

func (doc DocumentPain01400108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01400108) NameSpace() string {
	return utils.DocumentPain01400108NameSpace
}

func (doc DocumentPain01400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName                xml.Name
		Attrs                  []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
		CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"CdtrPmtActvtnReqStsRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01400108) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01400108) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentPain01400108) InspectDocument() interface{} {
	return &doc.CdtrPmtActvtnReqStsRpt
}

type DocumentPain01300108 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"CdtrPmtActvtnReq"`
}

func (doc DocumentPain01300108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPain01300108) NameSpace() string {
	return utils.DocumentPain01300108NameSpace
}

func (doc DocumentPain01300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"CdtrPmtActvtnReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPain01300108) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPain01300108) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentPain01300108) InspectDocument() interface{} {
	return &doc.CdtrPmtActvtnReq
}
