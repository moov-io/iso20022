// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAuth01800102 struct {
	XMLName                 *xml.Name                      `json:",omitempty"`
	Xmlns                   string                         `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                           `xml:",omitempty" json:",omitempty"`
	CtrctRegnReq            ContractRegistrationRequestV02 `xml:"CtrctRegnReq"`
}

func (doc DocumentAuth01800102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth01800102) NameSpace() string {
	return utils.DocumentAuth01800102NameSpace
}

func (doc DocumentAuth01800102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnReq ContractRegistrationRequestV02 `xml:"CtrctRegnReq"`
	}
	output.CtrctRegnReq = doc.CtrctRegnReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth01900102 struct {
	XMLName                 *xml.Name                           `json:",omitempty"`
	Xmlns                   string                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                `xml:",omitempty" json:",omitempty"`
	CtrctRegnConf           ContractRegistrationConfirmationV02 `xml:"CtrctRegnConf"`
}

func (doc DocumentAuth01900102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth01900102) NameSpace() string {
	return utils.DocumentAuth01900102NameSpace
}

func (doc DocumentAuth01900102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnConf ContractRegistrationConfirmationV02 `xml:"CtrctRegnConf"`
	}
	output.CtrctRegnConf = doc.CtrctRegnConf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02000102 struct {
	XMLName                 *xml.Name                             `json:",omitempty"`
	Xmlns                   string                                `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                  `xml:",omitempty" json:",omitempty"`
	CtrctRegnClsrReq        ContractRegistrationClosureRequestV02 `xml:"CtrctRegnClsrReq"`
}

func (doc DocumentAuth02000102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02000102) NameSpace() string {
	return utils.DocumentAuth02000102NameSpace
}

func (doc DocumentAuth02000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnClsrReq ContractRegistrationClosureRequestV02 `xml:"CtrctRegnClsrReq"`
	}
	output.CtrctRegnClsrReq = doc.CtrctRegnClsrReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02100102 struct {
	XMLName                 *xml.Name                               `json:",omitempty"`
	Xmlns                   string                                  `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                    `xml:",omitempty" json:",omitempty"`
	CtrctRegnAmdmntReq      ContractRegistrationAmendmentRequestV02 `xml:"CtrctRegnAmdmntReq"`
}

func (doc DocumentAuth02100102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02100102) NameSpace() string {
	return utils.DocumentAuth02100102NameSpace
}

func (doc DocumentAuth02100102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnAmdmntReq ContractRegistrationAmendmentRequestV02 `xml:"CtrctRegnAmdmntReq"`
	}
	output.CtrctRegnAmdmntReq = doc.CtrctRegnAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02200102 struct {
	XMLName                 *xml.Name                        `json:",omitempty"`
	Xmlns                   string                           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                             `xml:",omitempty" json:",omitempty"`
	CtrctRegnStmt           ContractRegistrationStatementV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctRegnStmt"`
}

func (doc DocumentAuth02200102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02200102) NameSpace() string {
	return utils.DocumentAuth02200102NameSpace
}

func (doc DocumentAuth02200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnStmt ContractRegistrationStatementV02 `xml:"CtrctRegnStmt"`
	}
	output.CtrctRegnStmt = doc.CtrctRegnStmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02300102 struct {
	XMLName                 *xml.Name                               `json:",omitempty"`
	Xmlns                   string                                  `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                    `xml:",omitempty" json:",omitempty"`
	CtrctRegnStmtReq        ContractRegistrationStatementRequestV02 `xml:"CtrctRegnStmtReq"`
}

func (doc DocumentAuth02300102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02300102) NameSpace() string {
	return utils.DocumentAuth02300102NameSpace
}

func (doc DocumentAuth02300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CtrctRegnStmtReq ContractRegistrationStatementRequestV02 `xml:"CtrctRegnStmtReq"`
	}
	output.CtrctRegnStmtReq = doc.CtrctRegnStmtReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02400102 struct {
	XMLName                 *xml.Name                                   `json:",omitempty"`
	Xmlns                   string                                      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                        `xml:",omitempty" json:",omitempty"`
	PmtRgltryInfNtfctn      PaymentRegulatoryInformationNotificationV02 `xml:"PmtRgltryInfNtfctn"`
}

func (doc DocumentAuth02400102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02400102) NameSpace() string {
	return utils.DocumentAuth02400102NameSpace
}

func (doc DocumentAuth02400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PmtRgltryInfNtfctn PaymentRegulatoryInformationNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 PmtRgltryInfNtfctn"`
	}
	output.PmtRgltryInfNtfctn = doc.PmtRgltryInfNtfctn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02500102 struct {
	XMLName                 *xml.Name                                    `json:",omitempty"`
	Xmlns                   string                                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                         `xml:",omitempty" json:",omitempty"`
	CcyCtrlSpprtgDocDlvry   CurrencyControlSupportingDocumentDeliveryV02 `xml:"CcyCtrlSpprtgDocDlvry"`
}

func (doc DocumentAuth02500102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02500102) NameSpace() string {
	return utils.DocumentAuth02500102NameSpace
}

func (doc DocumentAuth02500102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"CcyCtrlSpprtgDocDlvry"`
	}
	output.CcyCtrlSpprtgDocDlvry = doc.CcyCtrlSpprtgDocDlvry
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02600102 struct {
	XMLName                 *xml.Name                         `json:",omitempty"`
	Xmlns                   string                            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                              `xml:",omitempty" json:",omitempty"`
	CcyCtrlReqOrLttr        CurrencyControlRequestOrLetterV02 `xml:"CcyCtrlReqOrLttr"`
}

func (doc DocumentAuth02600102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02600102) NameSpace() string {
	return utils.DocumentAuth02600102NameSpace
}

func (doc DocumentAuth02600102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CcyCtrlReqOrLttr CurrencyControlRequestOrLetterV02 `xml:"CcyCtrlReqOrLttr"`
	}
	output.CcyCtrlReqOrLttr = doc.CcyCtrlReqOrLttr
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAuth02700102 struct {
	XMLName                 *xml.Name                      `json:",omitempty"`
	Xmlns                   string                         `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                           `xml:",omitempty" json:",omitempty"`
	CcyCtrlStsAdvc          CurrencyControlStatusAdviceV02 `xml:"CcyCtrlStsAdvc"`
}

func (doc DocumentAuth02700102) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02700102) NameSpace() string {
	return utils.DocumentAuth02700102NameSpace
}

func (doc DocumentAuth02700102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"CcyCtrlStsAdvc"`
	}
	output.CcyCtrlStsAdvc = doc.CcyCtrlStsAdvc
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
