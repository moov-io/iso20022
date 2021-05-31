// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAuth01800102 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	CtrctRegnReq ContractRegistrationRequestV02 `xml:"CtrctRegnReq"`
}

func (doc DocumentAuth01800102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth01800102) NameSpace() string {
	return utils.DocumentAuth01800102NameSpace
}

func (doc DocumentAuth01800102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
		CtrctRegnReq ContractRegistrationRequestV02 `xml:"CtrctRegnReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth01800102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth01800102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth01900102 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	CtrctRegnConf ContractRegistrationConfirmationV02 `xml:"CtrctRegnConf"`
}

func (doc DocumentAuth01900102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth01900102) NameSpace() string {
	return utils.DocumentAuth01900102NameSpace
}

func (doc DocumentAuth01900102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		CtrctRegnConf ContractRegistrationConfirmationV02 `xml:"CtrctRegnConf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth01900102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth01900102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02000102 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
	CtrctRegnClsrReq ContractRegistrationClosureRequestV02 `xml:"CtrctRegnClsrReq"`
}

func (doc DocumentAuth02000102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02000102) NameSpace() string {
	return utils.DocumentAuth02000102NameSpace
}

func (doc DocumentAuth02000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
		CtrctRegnClsrReq ContractRegistrationClosureRequestV02 `xml:"CtrctRegnClsrReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02000102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02000102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02100102 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
	CtrctRegnAmdmntReq ContractRegistrationAmendmentRequestV02 `xml:"CtrctRegnAmdmntReq"`
}

func (doc DocumentAuth02100102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02100102) NameSpace() string {
	return utils.DocumentAuth02100102NameSpace
}

func (doc DocumentAuth02100102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
		CtrctRegnAmdmntReq ContractRegistrationAmendmentRequestV02 `xml:"CtrctRegnAmdmntReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02100102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02100102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02200102 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
	CtrctRegnStmt ContractRegistrationStatementV02 `xml:"CtrctRegnStmt"`
}

func (doc DocumentAuth02200102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02200102) NameSpace() string {
	return utils.DocumentAuth02200102NameSpace
}

func (doc DocumentAuth02200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
		CtrctRegnStmt ContractRegistrationStatementV02 `xml:"CtrctRegnStmt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02200102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02200102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02300102 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
	CtrctRegnStmtReq ContractRegistrationStatementRequestV02 `xml:"CtrctRegnStmtReq"`
}

func (doc DocumentAuth02300102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02300102) NameSpace() string {
	return utils.DocumentAuth02300102NameSpace
}

func (doc DocumentAuth02300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
		CtrctRegnStmtReq ContractRegistrationStatementRequestV02 `xml:"CtrctRegnStmtReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02300102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02300102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02400102 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
	PmtRgltryInfNtfctn PaymentRegulatoryInformationNotificationV02 `xml:"PmtRgltryInfNtfctn"`
}

func (doc DocumentAuth02400102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02400102) NameSpace() string {
	return utils.DocumentAuth02400102NameSpace
}

func (doc DocumentAuth02400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
		PmtRgltryInfNtfctn PaymentRegulatoryInformationNotificationV02 `xml:"PmtRgltryInfNtfctn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02400102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02400102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02500102 struct {
	XMLName               xml.Name
	Attrs                 []utils.Attr                                 `xml:",any,attr,omitempty" json:",omitempty"`
	CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"CcyCtrlSpprtgDocDlvry"`
}

func (doc DocumentAuth02500102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02500102) NameSpace() string {
	return utils.DocumentAuth02500102NameSpace
}

func (doc DocumentAuth02500102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName               xml.Name
		Attrs                 []utils.Attr                                 `xml:",any,attr,omitempty" json:",omitempty"`
		CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"CcyCtrlSpprtgDocDlvry"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02500102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02500102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02600102 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
	CcyCtrlReqOrLttr CurrencyControlRequestOrLetterV02 `xml:"CcyCtrlReqOrLttr"`
}

func (doc DocumentAuth02600102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02600102) NameSpace() string {
	return utils.DocumentAuth02600102NameSpace
}

func (doc DocumentAuth02600102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
		CcyCtrlReqOrLttr CurrencyControlRequestOrLetterV02 `xml:"CcyCtrlReqOrLttr"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02600102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02600102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

type DocumentAuth02700102 struct {
	XMLName        xml.Name
	Attrs          []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"CcyCtrlStsAdvc"`
}

func (doc DocumentAuth02700102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAuth02700102) NameSpace() string {
	return utils.DocumentAuth02700102NameSpace
}

func (doc DocumentAuth02700102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName        xml.Name
		Attrs          []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
		CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"CcyCtrlStsAdvc"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAuth02700102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAuth02700102) GetAttrs() []utils.Attr {
	return doc.Attrs
}
