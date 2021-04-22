// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt00400108 struct {
	Xmlns   string           `xml:"xmlns,attr"`
	RtrAcct ReturnAccountV08 `xml:"RtrAcct"`
}

func (doc DocumentCamt00400108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00400108) NameSpace() string {
	return utils.DocumentCamt00400108NameSpace
}

func (doc DocumentCamt00400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrAcct ReturnAccountV08 `xml:"RtrAcct"`
	}
	output.RtrAcct = doc.RtrAcct
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt00500108 struct {
	Xmlns string            `xml:"xmlns,attr"`
	GetTx GetTransactionV08 `xml:"GetTx"`
}

func (doc DocumentCamt00500108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00500108) NameSpace() string {
	return utils.DocumentCamt00500108NameSpace
}

func (doc DocumentCamt00500108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetTx GetTransactionV08 `xml:"GetTx"`
	}
	output.GetTx = doc.GetTx
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt00600108 struct {
	Xmlns string               `xml:"xmlns,attr"`
	RtrTx ReturnTransactionV08 `xml:"RtrTx"`
}

func (doc DocumentCamt00600108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00600108) NameSpace() string {
	return utils.DocumentCamt00600108NameSpace
}

func (doc DocumentCamt00600108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrTx ReturnTransactionV08 `xml:"RtrTx"`
	}
	output.RtrTx = doc.RtrTx
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt00700108 struct {
	Xmlns   string               `xml:"xmlns,attr"`
	ModfyTx ModifyTransactionV08 `xml:"ModfyTx"`
}

func (doc DocumentCamt00700108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00700108) NameSpace() string {
	return utils.DocumentCamt00700108NameSpace
}

func (doc DocumentCamt00700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyTx ModifyTransactionV08 `xml:"ModfyTx"`
	}
	output.ModfyTx = doc.ModfyTx
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt00800108 struct {
	Xmlns string               `xml:"xmlns,attr"`
	CclTx CancelTransactionV08 `xml:"CclTx"`
}

func (doc DocumentCamt00800108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00800108) NameSpace() string {
	return utils.DocumentCamt00800108NameSpace
}

func (doc DocumentCamt00800108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CclTx CancelTransactionV08 `xml:"CclTx"`
	}
	output.CclTx = doc.CclTx
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt01000108 struct {
	Xmlns  string         `xml:"xmlns,attr"`
	RtrLmt ReturnLimitV08 `xml:"RtrLmt"`
}

func (doc DocumentCamt01000108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01000108) NameSpace() string {
	return utils.DocumentCamt01000108NameSpace
}

func (doc DocumentCamt01000108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrLmt ReturnLimitV08 `xml:"RtrLmt"`
	}
	output.RtrLmt = doc.RtrLmt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt02600108 struct {
	Xmlns      string           `xml:"xmlns,attr"`
	UblToApply UnableToApplyV08 `xml:"UblToApply"`
}

func (doc DocumentCamt02600108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02600108) NameSpace() string {
	return utils.DocumentCamt02600108NameSpace
}

func (doc DocumentCamt02600108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		UblToApply UnableToApplyV08 `xml:"UblToApply"`
	}
	output.UblToApply = doc.UblToApply
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt02700108 struct {
	Xmlns     string             `xml:"xmlns,attr"`
	ClmNonRct ClaimNonReceiptV08 `xml:"ClmNonRct"`
}

func (doc DocumentCamt02700108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02700108) NameSpace() string {
	return utils.DocumentCamt02700108NameSpace
}

func (doc DocumentCamt02700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ClmNonRct ClaimNonReceiptV08 `xml:"ClmNonRct"`
	}
	output.ClmNonRct = doc.ClmNonRct
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt03700108 struct {
	Xmlns         string                       `xml:"xmlns,attr"`
	DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"DbtAuthstnReq"`
}

func (doc DocumentCamt03700108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03700108) NameSpace() string {
	return utils.DocumentCamt03700108NameSpace
}

func (doc DocumentCamt03700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"DbtAuthstnReq"`
	}
	output.DbtAuthstnReq = doc.DbtAuthstnReq
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05200108 struct {
	Xmlns            string                         `xml:"xmlns,attr"`
	BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"BkToCstmrAcctRpt"`
}

func (doc DocumentCamt05200108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05200108) NameSpace() string {
	return utils.DocumentCamt05200108NameSpace
}

func (doc DocumentCamt05200108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"BkToCstmrAcctRpt"`
	}
	output.BkToCstmrAcctRpt = doc.BkToCstmrAcctRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05300108 struct {
	Xmlns         string                     `xml:"xmlns,attr"`
	BkToCstmrStmt BankToCustomerStatementV08 `xml:"BkToCstmrStmt"`
}

func (doc DocumentCamt05300108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05300108) NameSpace() string {
	return utils.DocumentCamt05300108NameSpace
}

func (doc DocumentCamt05300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkToCstmrStmt BankToCustomerStatementV08 `xml:"BkToCstmrStmt"`
	}
	output.BkToCstmrStmt = doc.BkToCstmrStmt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}

type DocumentCamt05400108 struct {
	Xmlns                 string                                   `xml:"xmlns,attr"`
	BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"BkToCstmrDbtCdtNtfctn"`
}

func (doc DocumentCamt05400108) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05400108) NameSpace() string {
	return utils.DocumentCamt05400108NameSpace
}

func (doc DocumentCamt05400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"BkToCstmrDbtCdtNtfctn"`
	}
	output.BkToCstmrDbtCdtNtfctn = doc.BkToCstmrDbtCdtNtfctn
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
