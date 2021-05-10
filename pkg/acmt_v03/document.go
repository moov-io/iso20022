// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt00700103 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
	AcctOpngReq AccountOpeningRequestV03 `xml:"AcctOpngReq"`
}

func (doc DocumentAcmt00700103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00700103) NameSpace() string {
	return utils.DocumentAcmt00700103NameSpace
}

func (doc DocumentAcmt00700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
		AcctOpngReq AccountOpeningRequestV03 `xml:"AcctOpngReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt00800103 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
	AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"AcctOpngAmdmntReq"`
}

func (doc DocumentAcmt00800103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00800103) NameSpace() string {
	return utils.DocumentAcmt00800103NameSpace
}

func (doc DocumentAcmt00800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
		AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"AcctOpngAmdmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt00900103 struct {
	XMLName             xml.Name
	Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
	AcctOpngAddtlInfReq AccountOpeningAdditionalInformationRequestV03 `xml:"AcctOpngAddtlInfReq"`
}

func (doc DocumentAcmt00900103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00900103) NameSpace() string {
	return utils.DocumentAcmt00900103NameSpace
}

func (doc DocumentAcmt00900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName             xml.Name
		Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
		AcctOpngAddtlInfReq AccountOpeningAdditionalInformationRequestV03 `xml:"AcctOpngAddtlInfReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01000103 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
	AcctReqAck AccountRequestAcknowledgementV03 `xml:"AcctReqAck"`
}

func (doc DocumentAcmt01000103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01000103) NameSpace() string {
	return utils.DocumentAcmt01000103NameSpace
}

func (doc DocumentAcmt01000103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
		AcctReqAck AccountRequestAcknowledgementV03 `xml:"AcctReqAck"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01100103 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	AcctReqRjctn AccountRequestRejectionV03 `xml:"AcctReqRjctn"`
}

func (doc DocumentAcmt01100103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01100103) NameSpace() string {
	return utils.DocumentAcmt01100103NameSpace
}

func (doc DocumentAcmt01100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		AcctReqRjctn AccountRequestRejectionV03 `xml:"AcctReqRjctn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01200103 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                           `xml:",any,attr,omitempty" json:",omitempty"`
	AcctAddtlInfReq AccountAdditionalInformationRequestV03 `xml:"AcctAddtlInfReq"`
}

func (doc DocumentAcmt01200103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01200103) NameSpace() string {
	return utils.DocumentAcmt01200103NameSpace
}

func (doc DocumentAcmt01200103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                           `xml:",any,attr,omitempty" json:",omitempty"`
		AcctAddtlInfReq AccountAdditionalInformationRequestV03 `xml:"AcctAddtlInfReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01300103 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr            `xml:",any,attr,omitempty" json:",omitempty"`
	AcctRptReq AccountReportRequestV03 `xml:"AcctRptReq"`
}

func (doc DocumentAcmt01300103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01300103) NameSpace() string {
	return utils.DocumentAcmt01300103NameSpace
}

func (doc DocumentAcmt01300103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr            `xml:",any,attr,omitempty" json:",omitempty"`
		AcctRptReq AccountReportRequestV03 `xml:"AcctRptReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01400103 struct {
	XMLName xml.Name
	Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	AcctRpt AccountReportV03 `xml:"AcctRpt"`
}

func (doc DocumentAcmt01400103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01400103) NameSpace() string {
	return utils.DocumentAcmt01400103NameSpace
}

func (doc DocumentAcmt01400103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
		AcctRpt AccountReportV03 `xml:"AcctRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01500103 struct {
	XMLName               xml.Name
	Attrs                 []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
	AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV03 `xml:"AcctExcldMndtMntncReq"`
}

func (doc DocumentAcmt01500103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01500103) NameSpace() string {
	return utils.DocumentAcmt01500103NameSpace
}

func (doc DocumentAcmt01500103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName               xml.Name
		Attrs                 []utils.Attr                                `xml:",any,attr,omitempty" json:",omitempty"`
		AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV03 `xml:"AcctExcldMndtMntncReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01600103 struct {
	XMLName                     xml.Name
	Attrs                       []utils.Attr                                         `xml:",any,attr,omitempty" json:",omitempty"`
	AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01600103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01600103) NameSpace() string {
	return utils.DocumentAcmt01600103NameSpace
}

func (doc DocumentAcmt01600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                     xml.Name
		Attrs                       []utils.Attr                                         `xml:",any,attr,omitempty" json:",omitempty"`
		AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"AcctExcldMndtMntncAmdmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01700103 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	AcctMndtMntncReq AccountMandateMaintenanceRequestV03 `xml:"AcctMndtMntncReq"`
}

func (doc DocumentAcmt01700103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01700103) NameSpace() string {
	return utils.DocumentAcmt01700103NameSpace
}

func (doc DocumentAcmt01700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		AcctMndtMntncReq AccountMandateMaintenanceRequestV03 `xml:"AcctMndtMntncReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01800103 struct {
	XMLName                xml.Name
	Attrs                  []utils.Attr                                 `xml:",any,attr,omitempty" json:",omitempty"`
	AcctMndtMntncAmdmntReq AccountMandateMaintenanceAmendmentRequestV03 `xml:"AcctMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01800103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01800103) NameSpace() string {
	return utils.DocumentAcmt01800103NameSpace
}

func (doc DocumentAcmt01800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName                xml.Name
		Attrs                  []utils.Attr                                 `xml:",any,attr,omitempty" json:",omitempty"`
		AcctMndtMntncAmdmntReq AccountMandateMaintenanceAmendmentRequestV03 `xml:"AcctMndtMntncAmdmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt01900103 struct {
	XMLName     xml.Name
	Attrs       []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
	AcctClsgReq AccountClosingRequestV03 `xml:"AcctClsgReq"`
}

func (doc DocumentAcmt01900103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01900103) NameSpace() string {
	return utils.DocumentAcmt01900103NameSpace
}

func (doc DocumentAcmt01900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName     xml.Name
		Attrs       []utils.Attr             `xml:",any,attr,omitempty" json:",omitempty"`
		AcctClsgReq AccountClosingRequestV03 `xml:"AcctClsgReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt02000103 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
	AcctClsgAmdmntReq AccountClosingAmendmentRequestV03 `xml:"AcctClsgAmdmntReq"`
}

func (doc DocumentAcmt02000103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02000103) NameSpace() string {
	return utils.DocumentAcmt02000103NameSpace
}

func (doc DocumentAcmt02000103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                      `xml:",any,attr,omitempty" json:",omitempty"`
		AcctClsgAmdmntReq AccountClosingAmendmentRequestV03 `xml:"AcctClsgAmdmntReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt02100103 struct {
	XMLName             xml.Name
	Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
	AcctClsgAddtlInfReq AccountClosingAdditionalInformationRequestV03 `xml:"AcctClsgAddtlInfReq"`
}

func (doc DocumentAcmt02100103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02100103) NameSpace() string {
	return utils.DocumentAcmt02100103NameSpace
}

func (doc DocumentAcmt02100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName             xml.Name
		Attrs               []utils.Attr                                  `xml:",any,attr,omitempty" json:",omitempty"`
		AcctClsgAddtlInfReq AccountClosingAdditionalInformationRequestV03 `xml:"AcctClsgAddtlInfReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt02700103 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchInfReq AccountSwitchInformationRequestV03 `xml:"AcctSwtchInfReq"`
}

func (doc DocumentAcmt02700103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02700103) NameSpace() string {
	return utils.DocumentAcmt02700103NameSpace
}

func (doc DocumentAcmt02700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchInfReq AccountSwitchInformationRequestV03 `xml:"AcctSwtchInfReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt02800103 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchInfRspn AccountSwitchInformationResponseV03 `xml:"AcctSwtchInfRspn"`
}

func (doc DocumentAcmt02800103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02800103) NameSpace() string {
	return utils.DocumentAcmt02800103NameSpace
}

func (doc DocumentAcmt02800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                        `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchInfRspn AccountSwitchInformationResponseV03 `xml:"AcctSwtchInfRspn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt02900103 struct {
	XMLName              xml.Name
	Attrs                []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchCclExstgPmt AccountSwitchCancelExistingPaymentV03 `xml:"AcctSwtchCclExstgPmt"`
}

func (doc DocumentAcmt02900103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02900103) NameSpace() string {
	return utils.DocumentAcmt02900103NameSpace
}

func (doc DocumentAcmt02900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName              xml.Name
		Attrs                []utils.Attr                          `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchCclExstgPmt AccountSwitchCancelExistingPaymentV03 `xml:"AcctSwtchCclExstgPmt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt03100103 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                           `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV03 `xml:"AcctSwtchReqBalTrf"`
}

func (doc DocumentAcmt03100103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03100103) NameSpace() string {
	return utils.DocumentAcmt03100103NameSpace
}

func (doc DocumentAcmt03100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                           `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV03 `xml:"AcctSwtchReqBalTrf"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt03200103 struct {
	XMLName            xml.Name
	Attrs              []utils.Attr                                   `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchBalTrfAck AccountSwitchBalanceTransferAcknowledgementV03 `xml:"AcctSwtchBalTrfAck"`
}

func (doc DocumentAcmt03200103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03200103) NameSpace() string {
	return utils.DocumentAcmt03200103NameSpace
}

func (doc DocumentAcmt03200103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName            xml.Name
		Attrs              []utils.Attr                                   `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchBalTrfAck AccountSwitchBalanceTransferAcknowledgementV03 `xml:"AcctSwtchBalTrfAck"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAcmt03400103 struct {
	XMLName         xml.Name
	Attrs           []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	AcctSwtchReqPmt AccountSwitchRequestPaymentV03 `xml:"AcctSwtchReqPmt"`
}

func (doc DocumentAcmt03400103) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03400103) NameSpace() string {
	return utils.DocumentAcmt03400103NameSpace
}

func (doc DocumentAcmt03400103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName         xml.Name
		Attrs           []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
		AcctSwtchReqPmt AccountSwitchRequestPaymentV03 `xml:"AcctSwtchReqPmt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}
