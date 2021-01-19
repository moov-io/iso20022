// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt02200102 struct {
	XMLName   xml.Name                            `xml:"Document" json:"-"`
	IdModAdvc IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
}

func (doc DocumentAcmt02200102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		IdModAdvc IdentificationModificationAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02 IdModAdvc"`
	}
	output.IdModAdvc = doc.IdModAdvc
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02300102 struct {
	XMLName     xml.Name                             `xml:"Document" json:"-"`
	IdVrfctnReq IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
}

func (doc DocumentAcmt02300102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		IdVrfctnReq IdentificationVerificationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02 IdVrfctnReq"`
	}
	output.IdVrfctnReq = doc.IdVrfctnReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02400102 struct {
	XMLName     xml.Name                            `xml:"Document" json:"-"`
	IdVrfctnRpt IdentificationVerificationReportV02 `xml:"IdVrfctnRpt"`
}

func (doc DocumentAcmt02400102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		IdVrfctnRpt IdentificationVerificationReportV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.024.001.02 IdVrfctnRpt"`
	}
	output.IdVrfctnRpt = doc.IdVrfctnRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.024.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03000102 struct {
	XMLName            xml.Name                           `xml:"Document" json:"-"`
	AcctSwtchReqRdrctn AccountSwitchRequestRedirectionV02 `xml:"AcctSwtchReqRdrctn"`
}

func (doc DocumentAcmt03000102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchReqRdrctn AccountSwitchRequestRedirectionV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.030.001.02 AcctSwtchReqRdrctn"`
	}
	output.AcctSwtchReqRdrctn = doc.AcctSwtchReqRdrctn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.030.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03300102 struct {
	XMLName                     xml.Name                                    `xml:"Document" json:"-"`
	AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"AcctSwtchNtfyAcctSwtchCmplt"`
}

func (doc DocumentAcmt03300102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.033.001.02 AcctSwtchNtfyAcctSwtchCmplt"`
	}
	output.AcctSwtchNtfyAcctSwtchCmplt = doc.AcctSwtchNtfyAcctSwtchCmplt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.033.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03500102 struct {
	XMLName          xml.Name                        `xml:"Document" json:"-"`
	AcctSwtchPmtRspn AccountSwitchPaymentResponseV02 `xml:"AcctSwtchPmtRspn"`
}

func (doc DocumentAcmt03500102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03500102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchPmtRspn AccountSwitchPaymentResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.035.001.02 AcctSwtchPmtRspn"`
	}
	output.AcctSwtchPmtRspn = doc.AcctSwtchPmtRspn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.035.001.02")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03700102 struct {
	XMLName            xml.Name                           `xml:"Document" json:"-"`
	AcctSwtchTechRjctn AccountSwitchTechnicalRejectionV02 `xml:"AcctSwtchTechRjctn"`
}

func (doc DocumentAcmt03700102) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03700102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchTechRjctn AccountSwitchTechnicalRejectionV02 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.037.001.02 AcctSwtchTechRjctn"`
	}
	output.AcctSwtchTechRjctn = doc.AcctSwtchTechRjctn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.037.001.02")
	return e.EncodeElement(&output, start)
}
