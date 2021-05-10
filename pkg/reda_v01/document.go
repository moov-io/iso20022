// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentReda06600101 struct {
	XMLName                xml.Name
	Attrs                  []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
}

func (doc DocumentReda06600101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06600101) NameSpace() string {
	return utils.DocumentReda06600101NameSpace
}

func (doc DocumentReda06600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                xml.Name
		Attrs                  []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda06700101 struct {
	XMLName                      xml.Name
	Attrs                        []utils.Attr                                     `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
}

func (doc DocumentReda06700101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06700101) NameSpace() string {
	return utils.DocumentReda06700101NameSpace
}

func (doc DocumentReda06700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                      xml.Name
		Attrs                        []utils.Attr                                     `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda06800101 struct {
	XMLName                   xml.Name
	Attrs                     []utils.Attr                                        `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
}

func (doc DocumentReda06800101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06800101) NameSpace() string {
	return utils.DocumentReda06800101NameSpace
}

func (doc DocumentReda06800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                   xml.Name
		Attrs                     []utils.Attr                                        `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda06900101 struct {
	XMLName                   xml.Name
	Attrs                     []utils.Attr                                 `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
}

func (doc DocumentReda06900101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06900101) NameSpace() string {
	return utils.DocumentReda06900101NameSpace
}

func (doc DocumentReda06900101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                   xml.Name
		Attrs                     []utils.Attr                                 `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda07000101 struct {
	XMLName               xml.Name
	Attrs                 []utils.Attr                           `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
}

func (doc DocumentReda07000101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07000101) NameSpace() string {
	return utils.DocumentReda07000101NameSpace
}

func (doc DocumentReda07000101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName               xml.Name
		Attrs                 []utils.Attr                           `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda07100101 struct {
	XMLName                     xml.Name
	Attrs                       []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
}

func (doc DocumentReda07100101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07100101) NameSpace() string {
	return utils.DocumentReda07100101NameSpace
}

func (doc DocumentReda07100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                     xml.Name
		Attrs                       []utils.Attr                                    `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda07200101 struct {
	XMLName                  xml.Name
	Attrs                    []utils.Attr                                       `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
}

func (doc DocumentReda07200101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07200101) NameSpace() string {
	return utils.DocumentReda07200101NameSpace
}

func (doc DocumentReda07200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                  xml.Name
		Attrs                    []utils.Attr                                       `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentReda07300101 struct {
	XMLName                  xml.Name
	Attrs                    []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
}

func (doc DocumentReda07300101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07300101) NameSpace() string {
	return utils.DocumentReda07300101NameSpace
}

func (doc DocumentReda07300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                  xml.Name
		Attrs                    []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
		ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}
