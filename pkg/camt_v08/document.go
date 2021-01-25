// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt00400108 struct {
	RtrAcct ReturnAccountV08 `xml:"RtrAcct"`
}

func (doc DocumentCamt00400108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrAcct ReturnAccountV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 RtrAcct"`
	}
	output.RtrAcct = doc.RtrAcct
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.004.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt00500108 struct {
	GetTx GetTransactionV08 `xml:"GetTx"`
}

func (doc DocumentCamt00500108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00500108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetTx GetTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 GetTx"`
	}
	output.GetTx = doc.GetTx
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.005.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt00800108 struct {
	CclTx CancelTransactionV08 `xml:"CclTx"`
}

func (doc DocumentCamt00800108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00800108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CclTx CancelTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CclTx"`
	}
	output.CclTx = doc.CclTx
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.008.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt00600108 struct {
	RtrTx ReturnTransactionV08 `xml:"RtrTx"`
}

func (doc DocumentCamt00600108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00600108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrTx ReturnTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 RtrTx"`
	}
	output.RtrTx = doc.RtrTx
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.006.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt02700108 struct {
	ClmNonRct ClaimNonReceiptV08 `xml:"ClmNonRct"`
}

func (doc DocumentCamt02700108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ClmNonRct ClaimNonReceiptV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.08 ClmNonRct"`
	}
	output.ClmNonRct = doc.ClmNonRct
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.027.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt01000108 struct {
	RtrLmt ReturnLimitV08 `xml:"RtrLmt"`
}

func (doc DocumentCamt01000108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt01000108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrLmt ReturnLimitV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 RtrLmt"`
	}
	output.RtrLmt = doc.RtrLmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.010.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt02600108 struct {
	UblToApply UnableToApplyV08 `xml:"UblToApply"`
}

func (doc DocumentCamt02600108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt02600108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		UblToApply UnableToApplyV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.08 UblToApply"`
	}
	output.UblToApply = doc.UblToApply
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.026.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05400108 struct {
	BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"BkToCstmrDbtCdtNtfctn"`
}

func (doc DocumentCamt05400108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.054.001.08 BkToCstmrDbtCdtNtfctn"`
	}
	output.BkToCstmrDbtCdtNtfctn = doc.BkToCstmrDbtCdtNtfctn
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.054.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05200108 struct {
	BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"BkToCstmrAcctRpt"`
}

func (doc DocumentCamt05200108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05200108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 BkToCstmrAcctRpt"`
	}
	output.BkToCstmrAcctRpt = doc.BkToCstmrAcctRpt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt00700108 struct {
	ModfyTx ModifyTransactionV08 `xml:"ModfyTx"`
}

func (doc DocumentCamt00700108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt00700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyTx ModifyTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ModfyTx"`
	}
	output.ModfyTx = doc.ModfyTx
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.007.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt03700108 struct {
	DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"DbtAuthstnReq"`
}

func (doc DocumentCamt03700108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt03700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.037.001.08 DbtAuthstnReq"`
	}
	output.DbtAuthstnReq = doc.DbtAuthstnReq
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.037.001.08")
	return e.EncodeElement(&output, start)
}

type DocumentCamt05300108 struct {
	BkToCstmrStmt BankToCustomerStatementV08 `xml:"BkToCstmrStmt"`
}

func (doc DocumentCamt05300108) Validate() error {
	return utils.Validate(&doc)
}

func (doc DocumentCamt05300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkToCstmrStmt BankToCustomerStatementV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.053.001.08 BkToCstmrStmt"`
	}
	output.BkToCstmrStmt = doc.BkToCstmrStmt
	utils.XmlElement(&start, "urn:iso:std:iso:20022:tech:xsd:camt.053.001.08")
	return e.EncodeElement(&output, start)
}
