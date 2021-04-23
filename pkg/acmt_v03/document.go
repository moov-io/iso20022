// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt00700103 struct {
	XMLName                 *xml.Name                `json:",omitempty"`
	Xmlns                   string                   `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                     `xml:",omitempty" json:",omitempty"`
	AcctOpngReq             AccountOpeningRequestV03 `xml:"AcctOpngReq"`
}

func (doc DocumentAcmt00700103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00700103) NameSpace() string {
	return utils.DocumentAcmt00700103NameSpace
}

func (doc DocumentAcmt00700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctOpngReq AccountOpeningRequestV03 `xml:"AcctOpngReq"`
	}
	output.AcctOpngReq = doc.AcctOpngReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt00800103 struct {
	XMLName                 *xml.Name                         `json:",omitempty"`
	Xmlns                   string                            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                              `xml:",omitempty" json:",omitempty"`
	AcctOpngAmdmntReq       AccountOpeningAmendmentRequestV03 `xml:"AcctOpngAmdmntReq"`
}

func (doc DocumentAcmt00800103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00800103) NameSpace() string {
	return utils.DocumentAcmt00800103NameSpace
}

func (doc DocumentAcmt00800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"AcctOpngAmdmntReq"`
	}
	output.AcctOpngAmdmntReq = doc.AcctOpngAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt00900103 struct {
	XMLName                 *xml.Name                                     `json:",omitempty"`
	Xmlns                   string                                        `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                          `xml:",omitempty" json:",omitempty"`
	AcctOpngAddtlInfReq     AccountOpeningAdditionalInformationRequestV03 `xml:"AcctOpngAddtlInfReq"`
}

func (doc DocumentAcmt00900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctOpngAddtlInfReq AccountOpeningAdditionalInformationRequestV03 `xml:"AcctOpngAddtlInfReq"`
	}
	output.AcctOpngAddtlInfReq = doc.AcctOpngAddtlInfReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

func (doc DocumentAcmt00900103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00900103) NameSpace() string {
	return utils.DocumentAcmt00900103NameSpace
}

type DocumentAcmt01000103 struct {
	XMLName                 *xml.Name                        `json:",omitempty"`
	Xmlns                   string                           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                             `xml:",omitempty" json:",omitempty"`
	AcctReqAck              AccountRequestAcknowledgementV03 `xml:"AcctReqAck"`
}

func (doc DocumentAcmt01000103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctReqAck AccountRequestAcknowledgementV03 `xml:"AcctReqAck"`
	}
	output.AcctReqAck = doc.AcctReqAck
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

func (doc DocumentAcmt01000103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01000103) NameSpace() string {
	return utils.DocumentAcmt01000103NameSpace
}

type DocumentAcmt01100103 struct {
	XMLName                 *xml.Name                  `json:",omitempty"`
	Xmlns                   string                     `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                       `xml:",omitempty" json:",omitempty"`
	AcctReqRjctn            AccountRequestRejectionV03 `xml:"AcctReqRjctn"`
}

func (doc DocumentAcmt01100103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01100103) NameSpace() string {
	return utils.DocumentAcmt01100103NameSpace
}

func (doc DocumentAcmt01100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctReqRjctn AccountRequestRejectionV03 `xml:"AcctReqRjctn"`
	}
	output.AcctReqRjctn = doc.AcctReqRjctn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01200103 struct {
	XMLName                 *xml.Name                              `json:",omitempty"`
	Xmlns                   string                                 `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                   `xml:",omitempty" json:",omitempty"`
	AcctAddtlInfReq         AccountAdditionalInformationRequestV03 `xml:"AcctAddtlInfReq"`
}

func (doc DocumentAcmt01200103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01200103) NameSpace() string {
	return utils.DocumentAcmt01200103NameSpace
}

func (doc DocumentAcmt01200103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctAddtlInfReq AccountAdditionalInformationRequestV03 `xml:"AcctAddtlInfReq"`
	}
	output.AcctAddtlInfReq = doc.AcctAddtlInfReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01300103 struct {
	XMLName                 *xml.Name               `json:",omitempty"`
	Xmlns                   string                  `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                    `xml:",omitempty" json:",omitempty"`
	AcctRptReq              AccountReportRequestV03 `xml:"AcctRptReq"`
}

func (doc DocumentAcmt01300103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01300103) NameSpace() string {
	return utils.DocumentAcmt01300103NameSpace
}

func (doc DocumentAcmt01300103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctRptReq AccountReportRequestV03 `xml:"AcctRptReq"`
	}
	output.AcctRptReq = doc.AcctRptReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01400103 struct {
	XMLName                 *xml.Name        `json:",omitempty"`
	Xmlns                   string           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool             `xml:",omitempty" json:",omitempty"`
	AcctRpt                 AccountReportV03 `xml:"AcctRpt"`
}

func (doc DocumentAcmt01400103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01400103) NameSpace() string {
	return utils.DocumentAcmt01400103NameSpace
}

func (doc DocumentAcmt01400103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctRpt AccountReportV03 `xml:"AcctRpt"`
	}
	output.AcctRpt = doc.AcctRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01500103 struct {
	XMLName                 *xml.Name                                   `json:",omitempty"`
	Xmlns                   string                                      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                        `xml:",omitempty" json:",omitempty"`
	AcctExcldMndtMntncReq   AccountExcludedMandateMaintenanceRequestV03 `xml:"AcctExcldMndtMntncReq"`
}

func (doc DocumentAcmt01500103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01500103) NameSpace() string {
	return utils.DocumentAcmt01500103NameSpace
}

func (doc DocumentAcmt01500103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV03 `xml:"AcctExcldMndtMntncReq"`
	}
	output.AcctExcldMndtMntncReq = doc.AcctExcldMndtMntncReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01600103 struct {
	XMLName                     *xml.Name                                            `json:",omitempty"`
	Xmlns                       string                                               `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace     bool                                                 `xml:",omitempty" json:",omitempty"`
	AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01600103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01600103) NameSpace() string {
	return utils.DocumentAcmt01600103NameSpace
}

func (doc DocumentAcmt01600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"AcctExcldMndtMntncAmdmntReq"`
	}
	output.AcctExcldMndtMntncAmdmntReq = doc.AcctExcldMndtMntncAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01700103 struct {
	XMLName                 *xml.Name                           `json:",omitempty"`
	Xmlns                   string                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                `xml:",omitempty" json:",omitempty"`
	AcctMndtMntncReq        AccountMandateMaintenanceRequestV03 `xml:"AcctMndtMntncReq"`
}

func (doc DocumentAcmt01700103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01700103) NameSpace() string {
	return utils.DocumentAcmt01700103NameSpace
}

func (doc DocumentAcmt01700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctMndtMntncReq AccountMandateMaintenanceRequestV03 `xml:"AcctMndtMntncReq"`
	}
	output.AcctMndtMntncReq = doc.AcctMndtMntncReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01800103 struct {
	XMLName                 *xml.Name                                    `json:",omitempty"`
	Xmlns                   string                                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                         `xml:",omitempty" json:",omitempty"`
	AcctMndtMntncAmdmntReq  AccountMandateMaintenanceAmendmentRequestV03 `xml:"AcctMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01800103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01800103) NameSpace() string {
	return utils.DocumentAcmt01800103NameSpace
}

func (doc DocumentAcmt01800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctMndtMntncAmdmntReq AccountMandateMaintenanceAmendmentRequestV03 `xml:"AcctMndtMntncAmdmntReq"`
	}
	output.AcctMndtMntncAmdmntReq = doc.AcctMndtMntncAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01900103 struct {
	XMLName                 *xml.Name                `json:",omitempty"`
	Xmlns                   string                   `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                     `xml:",omitempty" json:",omitempty"`
	AcctClsgReq             AccountClosingRequestV03 `xml:"AcctClsgReq"`
}

func (doc DocumentAcmt01900103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01900103) NameSpace() string {
	return utils.DocumentAcmt01900103NameSpace
}

func (doc DocumentAcmt01900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctClsgReq AccountClosingRequestV03 `xml:"AcctClsgReq"`
	}
	output.AcctClsgReq = doc.AcctClsgReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02000103 struct {
	XMLName                 *xml.Name                         `json:",omitempty"`
	Xmlns                   string                            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                              `xml:",omitempty" json:",omitempty"`
	AcctClsgAmdmntReq       AccountClosingAmendmentRequestV03 `xml:"AcctClsgAmdmntReq"`
}

func (doc DocumentAcmt02000103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02000103) NameSpace() string {
	return utils.DocumentAcmt02000103NameSpace
}

func (doc DocumentAcmt02000103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctClsgAmdmntReq AccountClosingAmendmentRequestV03 `xml:"AcctClsgAmdmntReq"`
	}
	output.AcctClsgAmdmntReq = doc.AcctClsgAmdmntReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02100103 struct {
	XMLName                 *xml.Name                                     `json:",omitempty"`
	Xmlns                   string                                        `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                          `xml:",omitempty" json:",omitempty"`
	AcctClsgAddtlInfReq     AccountClosingAdditionalInformationRequestV03 `xml:"AcctClsgAddtlInfReq"`
}

func (doc DocumentAcmt02100103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02100103) NameSpace() string {
	return utils.DocumentAcmt02100103NameSpace
}

func (doc DocumentAcmt02100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctClsgAddtlInfReq AccountClosingAdditionalInformationRequestV03 `xml:"AcctClsgAddtlInfReq"`
	}
	output.AcctClsgAddtlInfReq = doc.AcctClsgAddtlInfReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02700103 struct {
	XMLName                 *xml.Name                          `json:",omitempty"`
	Xmlns                   string                             `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                               `xml:",omitempty" json:",omitempty"`
	AcctSwtchInfReq         AccountSwitchInformationRequestV03 `xml:"AcctSwtchInfReq"`
}

func (doc DocumentAcmt02700103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02700103) NameSpace() string {
	return utils.DocumentAcmt02700103NameSpace
}

func (doc DocumentAcmt02700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchInfReq AccountSwitchInformationRequestV03 `xml:"AcctSwtchInfReq"`
	}
	output.AcctSwtchInfReq = doc.AcctSwtchInfReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02800103 struct {
	XMLName                 *xml.Name                           `json:",omitempty"`
	Xmlns                   string                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                `xml:",omitempty" json:",omitempty"`
	AcctSwtchInfRspn        AccountSwitchInformationResponseV03 `xml:"AcctSwtchInfRspn"`
}

func (doc DocumentAcmt02800103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02800103) NameSpace() string {
	return utils.DocumentAcmt02800103NameSpace
}

func (doc DocumentAcmt02800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchInfRspn AccountSwitchInformationResponseV03 `xml:"AcctSwtchInfRspn"`
	}
	output.AcctSwtchInfRspn = doc.AcctSwtchInfRspn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02900103 struct {
	XMLName                 *xml.Name                             `json:",omitempty"`
	Xmlns                   string                                `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                  `xml:",omitempty" json:",omitempty"`
	AcctSwtchCclExstgPmt    AccountSwitchCancelExistingPaymentV03 `xml:"AcctSwtchCclExstgPmt"`
}

func (doc DocumentAcmt02900103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02900103) NameSpace() string {
	return utils.DocumentAcmt02900103NameSpace
}

func (doc DocumentAcmt02900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchCclExstgPmt AccountSwitchCancelExistingPaymentV03 `xml:"AcctSwtchCclExstgPmt"`
	}
	output.AcctSwtchCclExstgPmt = doc.AcctSwtchCclExstgPmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03100103 struct {
	XMLName                 *xml.Name                              `json:",omitempty"`
	Xmlns                   string                                 `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                   `xml:",omitempty" json:",omitempty"`
	AcctSwtchReqBalTrf      AccountSwitchRequestBalanceTransferV03 `xml:"AcctSwtchReqBalTrf"`
}

func (doc DocumentAcmt03100103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03100103) NameSpace() string {
	return utils.DocumentAcmt03100103NameSpace
}

func (doc DocumentAcmt03100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV03 `xml:"AcctSwtchReqBalTrf"`
	}
	output.AcctSwtchReqBalTrf = doc.AcctSwtchReqBalTrf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03200103 struct {
	XMLName                 *xml.Name                                      `json:",omitempty"`
	Xmlns                   string                                         `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                           `xml:",omitempty" json:",omitempty"`
	AcctSwtchBalTrfAck      AccountSwitchBalanceTransferAcknowledgementV03 `xml:"AcctSwtchBalTrfAck"`
}

func (doc DocumentAcmt03200103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03200103) NameSpace() string {
	return utils.DocumentAcmt03200103NameSpace
}

func (doc DocumentAcmt03200103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchBalTrfAck AccountSwitchBalanceTransferAcknowledgementV03 `xml:"AcctSwtchBalTrfAck"`
	}
	output.AcctSwtchBalTrfAck = doc.AcctSwtchBalTrfAck
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03400103 struct {
	XMLName                 *xml.Name                      `json:",omitempty"`
	Xmlns                   string                         `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                           `xml:",omitempty" json:",omitempty"`
	AcctSwtchReqPmt         AccountSwitchRequestPaymentV03 `xml:"AcctSwtchReqPmt"`
}

func (doc DocumentAcmt03400103) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03400103) NameSpace() string {
	return utils.DocumentAcmt03400103NameSpace
}

func (doc DocumentAcmt03400103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchReqPmt AccountSwitchRequestPaymentV03 `xml:"AcctSwtchReqPmt"`
	}
	output.AcctSwtchReqPmt = doc.AcctSwtchReqPmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
