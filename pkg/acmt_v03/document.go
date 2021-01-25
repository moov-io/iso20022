// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAcmt00700103 struct {
	XMLName     xml.Name                 `xml:"Document" json:"-"`
	AcctOpngReq AccountOpeningRequestV03 `xml:"AcctOpngReq"`
}

func (doc DocumentAcmt00700103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctOpngReq AccountOpeningRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03 AcctOpngReq"`
	}
	output.AcctOpngReq = doc.AcctOpngReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt00800103 struct {
	XMLName           xml.Name                          `xml:"Document" json:"-"`
	AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"AcctOpngAmdmntReq"`
}

func (doc DocumentAcmt00800103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt00800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctOpngAmdmntReq AccountOpeningAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03 AcctOpngAmdmntReq"`
	}
	output.AcctOpngAmdmntReq = doc.AcctOpngAmdmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt00900103 struct {
	XMLName             xml.Name                                      `xml:"Document" json:"-"`
	AcctOpngAddtlInfReq AccountOpeningAdditionalInformationRequestV03 `xml:"AcctOpngAddtlInfReq"`
}

func (doc DocumentAcmt00900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctOpngAddtlInfReq AccountOpeningAdditionalInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.03 AcctOpngAddtlInfReq"`
	}
	output.AcctOpngAddtlInfReq = doc.AcctOpngAddtlInfReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.009.001.03")
	return e.EncodeElement(&output, start)
}

func (doc DocumentAcmt00900103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01000103 struct {
	XMLName    xml.Name                         `xml:"Document" json:"-"`
	AcctReqAck AccountRequestAcknowledgementV03 `xml:"AcctReqAck"`
}

func (doc DocumentAcmt01000103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctReqAck AccountRequestAcknowledgementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.010.001.03 AcctReqAck"`
	}
	output.AcctReqAck = doc.AcctReqAck
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.010.001.03")
	return e.EncodeElement(&output, start)
}

func (doc DocumentAcmt01000103) Validate() error {
	return utils.Validate(&doc)
}

type DocumentAcmt01100103 struct {
	XMLName      xml.Name                   `xml:"Document" json:"-"`
	AcctReqRjctn AccountRequestRejectionV03 `xml:"AcctReqRjctn"`
}

func (doc DocumentAcmt01100103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctReqRjctn AccountRequestRejectionV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.011.001.03 AcctReqRjctn"`
	}
	output.AcctReqRjctn = doc.AcctReqRjctn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.011.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01200103 struct {
	XMLName         xml.Name                               `xml:"Document" json:"-"`
	AcctAddtlInfReq AccountAdditionalInformationRequestV03 `xml:"AcctAddtlInfReq"`
}

func (doc DocumentAcmt01200103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01200103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctAddtlInfReq AccountAdditionalInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.012.001.03 AcctAddtlInfReq"`
	}
	output.AcctAddtlInfReq = doc.AcctAddtlInfReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.012.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01300103 struct {
	XMLName    xml.Name                `xml:"Document" json:"-"`
	AcctRptReq AccountReportRequestV03 `xml:"AcctRptReq"`
}

func (doc DocumentAcmt01300103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01300103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctRptReq AccountReportRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03 AcctRptReq"`
	}
	output.AcctRptReq = doc.AcctRptReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01400103 struct {
	XMLName xml.Name         `xml:"Document" json:"-"`
	AcctRpt AccountReportV03 `xml:"AcctRpt"`
}

func (doc DocumentAcmt01400103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01400103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctRpt AccountReportV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.03 AcctRpt"`
	}
	output.AcctRpt = doc.AcctRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.014.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01500103 struct {
	XMLName               xml.Name                                    `xml:"Document" json:"-"`
	AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV03 `xml:"AcctExcldMndtMntncReq"`
}

func (doc DocumentAcmt01500103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01500103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.03 AcctExcldMndtMntncReq"`
	}
	output.AcctExcldMndtMntncReq = doc.AcctExcldMndtMntncReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.015.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01600103 struct {
	XMLName                     xml.Name                                             `xml:"Document" json:"-"`
	AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01600103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01600103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctExcldMndtMntncAmdmntReq AccountExcludedMandateMaintenanceAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.03 AcctExcldMndtMntncAmdmntReq"`
	}
	output.AcctExcldMndtMntncAmdmntReq = doc.AcctExcldMndtMntncAmdmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.016.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01700103 struct {
	XMLName          xml.Name                            `xml:"Document" json:"-"`
	AcctMndtMntncReq AccountMandateMaintenanceRequestV03 `xml:"AcctMndtMntncReq"`
}

func (doc DocumentAcmt01700103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctMndtMntncReq AccountMandateMaintenanceRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.017.001.03 AcctMndtMntncReq"`
	}
	output.AcctMndtMntncReq = doc.AcctMndtMntncReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.017.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01800103 struct {
	XMLName                xml.Name                                     `xml:"Document" json:"-"`
	AcctMndtMntncAmdmntReq AccountMandateMaintenanceAmendmentRequestV03 `xml:"AcctMndtMntncAmdmntReq"`
}

func (doc DocumentAcmt01800103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctMndtMntncAmdmntReq AccountMandateMaintenanceAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.018.001.03 AcctMndtMntncAmdmntReq"`
	}
	output.AcctMndtMntncAmdmntReq = doc.AcctMndtMntncAmdmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.018.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt01900103 struct {
	XMLName     xml.Name                 `xml:"Document" json:"-"`
	AcctClsgReq AccountClosingRequestV03 `xml:"AcctClsgReq"`
}

func (doc DocumentAcmt01900103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt01900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctClsgReq AccountClosingRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.019.001.03 AcctClsgReq"`
	}
	output.AcctClsgReq = doc.AcctClsgReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.019.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02000103 struct {
	XMLName           xml.Name                          `xml:"Document" json:"-"`
	AcctClsgAmdmntReq AccountClosingAmendmentRequestV03 `xml:"AcctClsgAmdmntReq"`
}

func (doc DocumentAcmt02000103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02000103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctClsgAmdmntReq AccountClosingAmendmentRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.020.001.03 AcctClsgAmdmntReq"`
	}
	output.AcctClsgAmdmntReq = doc.AcctClsgAmdmntReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.020.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02100103 struct {
	XMLName             xml.Name                                      `xml:"Document" json:"-"`
	AcctClsgAddtlInfReq AccountClosingAdditionalInformationRequestV03 `xml:"AcctClsgAddtlInfReq"`
}

func (doc DocumentAcmt02100103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctClsgAddtlInfReq AccountClosingAdditionalInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.021.001.03 AcctClsgAddtlInfReq"`
	}
	output.AcctClsgAddtlInfReq = doc.AcctClsgAddtlInfReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.021.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02700103 struct {
	XMLName         xml.Name                           `xml:"Document" json:"-"`
	AcctSwtchInfReq AccountSwitchInformationRequestV03 `xml:"AcctSwtchInfReq"`
}

func (doc DocumentAcmt02700103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02700103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchInfReq AccountSwitchInformationRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.027.001.03 AcctSwtchInfReq"`
	}
	output.AcctSwtchInfReq = doc.AcctSwtchInfReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.027.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02800103 struct {
	XMLName          xml.Name                            `xml:"Document" json:"-"`
	AcctSwtchInfRspn AccountSwitchInformationResponseV03 `xml:"AcctSwtchInfRspn"`
}

func (doc DocumentAcmt02800103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02800103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchInfRspn AccountSwitchInformationResponseV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.028.001.03 AcctSwtchInfRspn"`
	}
	output.AcctSwtchInfRspn = doc.AcctSwtchInfRspn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.028.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt02900103 struct {
	XMLName              xml.Name                              `xml:"Document" json:"-"`
	AcctSwtchCclExstgPmt AccountSwitchCancelExistingPaymentV03 `xml:"AcctSwtchCclExstgPmt"`
}

func (doc DocumentAcmt02900103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt02900103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchCclExstgPmt AccountSwitchCancelExistingPaymentV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.029.001.03 AcctSwtchCclExstgPmt"`
	}
	output.AcctSwtchCclExstgPmt = doc.AcctSwtchCclExstgPmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.029.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03100103 struct {
	XMLName            xml.Name                               `xml:"Document" json:"-"`
	AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV03 `xml:"AcctSwtchReqBalTrf"`
}

func (doc DocumentAcmt03100103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03100103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchReqBalTrf AccountSwitchRequestBalanceTransferV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.03 AcctSwtchReqBalTrf"`
	}
	output.AcctSwtchReqBalTrf = doc.AcctSwtchReqBalTrf
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.031.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03200103 struct {
	XMLName            xml.Name                                       `xml:"Document" json:"-"`
	AcctSwtchBalTrfAck AccountSwitchBalanceTransferAcknowledgementV03 `xml:"AcctSwtchBalTrfAck"`
}

func (doc DocumentAcmt03200103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03200103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchBalTrfAck AccountSwitchBalanceTransferAcknowledgementV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.032.001.03 AcctSwtchBalTrfAck"`
	}
	output.AcctSwtchBalTrfAck = doc.AcctSwtchBalTrfAck
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.032.001.03")
	return e.EncodeElement(&output, start)
}

type DocumentAcmt03400103 struct {
	XMLName         xml.Name                       `xml:"Document" json:"-"`
	AcctSwtchReqPmt AccountSwitchRequestPaymentV03 `xml:"AcctSwtchReqPmt"`
}

func (doc DocumentAcmt03400103) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentAcmt03400103) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctSwtchReqPmt AccountSwitchRequestPaymentV03 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.034.001.03 AcctSwtchReqPmt"`
	}
	output.AcctSwtchReqPmt = doc.AcctSwtchReqPmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:acmt.034.001.03")
	return e.EncodeElement(&output, start)
}
