// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAuth018001V02 struct {
	XMLName      xml.Name                       `xml:"Document" json:"-"`
	CtrctRegnReq ContractRegistrationRequestV02 `xml:"CtrctRegnReq"`
}

func (doc DocumentAuth018001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth018001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnReq ContractRegistrationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctRegnReq"`
	}
	output.CtrctRegnReq = doc.CtrctRegnReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.018.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth019001V02 struct {
	XMLName       xml.Name                            `xml:"Document" json:"-"`
	CtrctRegnConf ContractRegistrationConfirmationV02 `xml:"CtrctRegnConf"`
}

func (doc DocumentAuth019001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth019001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnConf ContractRegistrationConfirmationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 CtrctRegnConf"`
	}
	output.CtrctRegnConf = doc.CtrctRegnConf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.019.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth020001V02 struct {
	XMLName          xml.Name                              `xml:"Document" json:"-"`
	CtrctRegnClsrReq ContractRegistrationClosureRequestV02 `xml:"CtrctRegnClsrReq"`
}

func (doc DocumentAuth020001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth020001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnClsrReq ContractRegistrationClosureRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 CtrctRegnClsrReq"`
	}
	output.CtrctRegnClsrReq = doc.CtrctRegnClsrReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.020.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth021001V02 struct {
	XMLName            xml.Name                                `xml:"Document" json:"-"`
	CtrctRegnAmdmntReq ContractRegistrationAmendmentRequestV02 `xml:"CtrctRegnAmdmntReq"`
}

func (doc DocumentAuth021001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth021001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnAmdmntReq ContractRegistrationAmendmentRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 CtrctRegnAmdmntReq"`
	}
	output.CtrctRegnAmdmntReq = doc.CtrctRegnAmdmntReq
	utils.XmlElement(&start, "urn:urn:iso:std:iso:20022:tech:xsd:auth.021.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth022001V02 struct {
	XMLName       xml.Name                         `xml:"Document" json:"-"`
	CtrctRegnStmt ContractRegistrationStatementV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctRegnStmt"`
}

func (doc DocumentAuth022001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth022001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnStmt ContractRegistrationStatementV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctRegnStmt"`
	}
	output.CtrctRegnStmt = doc.CtrctRegnStmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.022.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth023001V02 struct {
	XMLName          xml.Name                                `xml:"Document" json:"-"`
	CtrctRegnStmtReq ContractRegistrationStatementRequestV02 `xml:"CtrctRegnStmtReq"`
}

func (doc DocumentAuth023001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth023001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnStmtReq ContractRegistrationStatementRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 CtrctRegnStmtReq"`
	}
	output.CtrctRegnStmtReq = doc.CtrctRegnStmtReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.023.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth024001V02 struct {
	XMLName            xml.Name                                    `xml:"Document" json:"-"`
	PmtRgltryInfNtfctn PaymentRegulatoryInformationNotificationV02 `xml:"PmtRgltryInfNtfctn"`
}

func (doc DocumentAuth024001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth024001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PmtRgltryInfNtfctn PaymentRegulatoryInformationNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 PmtRgltryInfNtfctn"`
	}
	output.PmtRgltryInfNtfctn = doc.PmtRgltryInfNtfctn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.024.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth025001V02 struct {
	XMLName               xml.Name                                     `xml:"Document" json:"-"`
	CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"CcyCtrlSpprtgDocDlvry"`
}

func (doc DocumentAuth025001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth025001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CcyCtrlSpprtgDocDlvry"`
	}
	output.CcyCtrlSpprtgDocDlvry = doc.CcyCtrlSpprtgDocDlvry
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.025.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth026001V02 struct {
	XMLName          xml.Name                          `xml:"Document" json:"-"`
	CcyCtrlReqOrLttr CurrencyControlRequestOrLetterV02 `xml:"CcyCtrlReqOrLttr"`
}

func (doc DocumentAuth026001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth026001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CcyCtrlReqOrLttr CurrencyControlRequestOrLetterV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 CcyCtrlReqOrLttr"`
	}
	output.CcyCtrlReqOrLttr = doc.CcyCtrlReqOrLttr
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.026.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAuth027001V02 struct {
	XMLName        xml.Name                       `xml:"Document" json:"-"`
	CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"CcyCtrlStsAdvc"`
}

func (doc DocumentAuth027001V02) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAuth027001V02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CcyCtrlStsAdvc"`
	}
	output.CcyCtrlStsAdvc = doc.CcyCtrlStsAdvc
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:auth.027.001.02")
	return e.EncodeElement(&output, start)
}
