// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt02200102 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	IdModAdvc IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
}

func (doc DocumentAcmt02200102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02200102) NameSpace() string {
	return utils.DocumentAcmt02200102NameSpace
}

func (doc DocumentAcmt02200102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		IdModAdvc IdentificationModificationAdviceV02 `xml:"IdModAdvc"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt02200102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt02200102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt02200102) InspectDocument() interface{} {
	return &doc.IdModAdvc
}

type DocumentAcmt02300102 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr                         `xml:",any,attr,omitempty" json:",omitempty"`
	IdVrfctnReq IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
}

func (doc DocumentAcmt02300102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02300102) NameSpace() string {
	return utils.DocumentAcmt02300102NameSpace
}

func (doc DocumentAcmt02300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr                         `xml:",any,attr,omitempty" json:",omitempty"`
		IdVrfctnReq IdentificationVerificationRequestV02 `xml:"IdVrfctnReq"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt02300102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt02300102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt02300102) InspectDocument() interface{} {
	return &doc.IdVrfctnReq
}

type DocumentAcmt02400102 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	IdVrfctnRpt IdentificationVerificationReportV02 `xml:"IdVrfctnRpt"`
}

func (doc DocumentAcmt02400102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02400102) NameSpace() string {
	return utils.DocumentAcmt02400102NameSpace
}

func (doc DocumentAcmt02400102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		IdVrfctnRpt IdentificationVerificationReportV02 `xml:"IdVrfctnRpt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt02400102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt02400102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt02400102) InspectDocument() interface{} {
	return &doc.IdVrfctnRpt
}

type DocumentAcmt03000102 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchReqRdrctn AccountSwitchRequestRedirectionV02 `xml:"AcctSwtchReqRdrctn"`
}

func (doc DocumentAcmt03000102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03000102) NameSpace() string {
	return utils.DocumentAcmt03000102NameSpace
}

func (doc DocumentAcmt03000102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchReqRdrctn AccountSwitchRequestRedirectionV02 `xml:"AcctSwtchReqRdrctn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt03000102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt03000102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt03000102) InspectDocument() interface{} {
	return &doc.AcctSwtchReqRdrctn
}

type DocumentAcmt03300102 struct {
	XMLName                     xml.Name
	Attrs                       []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"AcctSwtchNtfyAcctSwtchCmplt"`
}

func (doc DocumentAcmt03300102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03300102) NameSpace() string {
	return utils.DocumentAcmt03300102NameSpace
}

func (doc DocumentAcmt03300102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName                     xml.Name
		Attrs                       []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchNtfyAcctSwtchCmplt AccountSwitchNotifyAccountSwitchCompleteV02 `xml:"AcctSwtchNtfyAcctSwtchCmplt"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt03300102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt03300102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt03300102) InspectDocument() interface{} {
	return &doc.AcctSwtchNtfyAcctSwtchCmplt
}

type DocumentAcmt03500102 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchPmtRspn AccountSwitchPaymentResponseV02 `xml:"AcctSwtchPmtRspn"`
}

func (doc DocumentAcmt03500102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03500102) NameSpace() string {
	return utils.DocumentAcmt03500102NameSpace
}

func (doc DocumentAcmt03500102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchPmtRspn AccountSwitchPaymentResponseV02 `xml:"AcctSwtchPmtRspn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt03500102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt03500102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt03500102) InspectDocument() interface{} {
	return &doc.AcctSwtchPmtRspn
}

type DocumentAcmt03700102 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchTechRjctn AccountSwitchTechnicalRejectionV02 `xml:"AcctSwtchTechRjctn"`
}

func (doc DocumentAcmt03700102) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03700102) NameSpace() string {
	return utils.DocumentAcmt03700102NameSpace
}

func (doc DocumentAcmt03700102) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchTechRjctn AccountSwitchTechnicalRejectionV02 `xml:"AcctSwtchTechRjctn"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentAcmt03700102) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentAcmt03700102) GetAttrs() []utils.Attr {
	return doc.Attrs
}

func (doc *DocumentAcmt03700102) InspectDocument() interface{} {
	return &doc.AcctSwtchTechRjctn
}
