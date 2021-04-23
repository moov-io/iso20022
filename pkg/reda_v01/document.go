// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentReda06600101 struct {
	XMLName                 *xml.Name                               `json:",omitempty"`
	Xmlns                   string                                  `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                    `xml:",omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntReq  RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
}

func (doc DocumentReda06600101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06600101) NameSpace() string {
	return utils.DocumentReda06600101NameSpace
}

func (doc DocumentReda06600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntReq RequestToPayCreditorEnrolmentRequestV01 `xml:"ReqToPayCdtrEnrlmntReq"`
	}
	output.ReqToPayCdtrEnrlmntReq = doc.ReqToPayCdtrEnrlmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda06700101 struct {
	XMLName                      *xml.Name                                        `json:",omitempty"`
	Xmlns                        string                                           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace      bool                                             `xml:",omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
}

func (doc DocumentReda06700101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06700101) NameSpace() string {
	return utils.DocumentReda06700101NameSpace
}

func (doc DocumentReda06700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntAmdmntReq RequestToPayCreditorEnrolmentAmendmentRequestV01 `xml:"ReqToPayCdtrEnrlmntAmdmntReq"`
	}
	output.ReqToPayCdtrEnrlmntAmdmntReq = doc.ReqToPayCdtrEnrlmntAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda06800101 struct {
	XMLName                   *xml.Name                                           `json:",omitempty"`
	Xmlns                     string                                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace   bool                                                `xml:",omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
}

func (doc DocumentReda06800101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06800101) NameSpace() string {
	return utils.DocumentReda06800101NameSpace
}

func (doc DocumentReda06800101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntCxlReq RequestToPayCreditorEnrolmentCancellationRequestV01 `xml:"ReqToPayCdtrEnrlmntCxlReq"`
	}
	output.ReqToPayCdtrEnrlmntCxlReq = doc.ReqToPayCdtrEnrlmntCxlReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda06900101 struct {
	XMLName                   *xml.Name                                    `json:",omitempty"`
	Xmlns                     string                                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace   bool                                         `xml:",omitempty" json:",omitempty"`
	ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
}

func (doc DocumentReda06900101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda06900101) NameSpace() string {
	return utils.DocumentReda06900101NameSpace
}

func (doc DocumentReda06900101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayCdtrEnrlmntStsRpt RequestToPayCreditorEnrolmentStatusReportV01 `xml:"ReqToPayCdtrEnrlmntStsRpt"`
	}
	output.ReqToPayCdtrEnrlmntStsRpt = doc.ReqToPayCdtrEnrlmntStsRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda07000101 struct {
	XMLName                 *xml.Name                              `json:",omitempty"`
	Xmlns                   string                                 `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                   `xml:",omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnReq   RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
}

func (doc DocumentReda07000101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07000101) NameSpace() string {
	return utils.DocumentReda07000101NameSpace
}

func (doc DocumentReda07000101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnReq RequestToPayDebtorActivationRequestV01 `xml:"ReqToPayDbtrActvtnReq"`
	}
	output.ReqToPayDbtrActvtnReq = doc.ReqToPayDbtrActvtnReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda07100101 struct {
	XMLName                     *xml.Name                                       `json:",omitempty"`
	Xmlns                       string                                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace     bool                                            `xml:",omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
}

func (doc DocumentReda07100101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07100101) NameSpace() string {
	return utils.DocumentReda07100101NameSpace
}

func (doc DocumentReda07100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:"ReqToPayDbtrActvtnAmdmntReq"`
	}
	output.ReqToPayDbtrActvtnAmdmntReq = doc.ReqToPayDbtrActvtnAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda07200101 struct {
	XMLName                  *xml.Name                                          `json:",omitempty"`
	Xmlns                    string                                             `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace  bool                                               `xml:",omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
}

func (doc DocumentReda07200101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07200101) NameSpace() string {
	return utils.DocumentReda07200101NameSpace
}

func (doc DocumentReda07200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnCxlReq RequestToPayDebtorActivationCancellationRequestV01 `xml:"ReqToPayDbtrActvtnCxlReq"`
	}
	output.ReqToPayDbtrActvtnCxlReq = doc.ReqToPayDbtrActvtnCxlReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentReda07300101 struct {
	XMLName                  *xml.Name                                   `json:",omitempty"`
	Xmlns                    string                                      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace  bool                                        `xml:",omitempty" json:",omitempty"`
	ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
}

func (doc DocumentReda07300101) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentReda07300101) NameSpace() string {
	return utils.DocumentReda07300101NameSpace
}

func (doc DocumentReda07300101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToPayDbtrActvtnStsRpt RequestToPayDebtorActivationStatusReportV01 `xml:"ReqToPayDbtrActvtnStsRpt"`
	}
	output.ReqToPayDbtrActvtnStsRpt = doc.ReqToPayDbtrActvtnStsRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
