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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                xml.Name
		Attrs                  []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
		CdtrPmtActvtnReqStsRpt CreditorPaymentActivationRequestStatusReportV08 `xml:"CdtrPmtActvtnReqStsRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
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
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		CdtrPmtActvtnReq CreditorPaymentActivationRequestV08 `xml:"CdtrPmtActvtnReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}
