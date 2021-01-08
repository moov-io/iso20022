// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

type DocumentCamt00400108 struct {
	RtrAcct ReturnAccountV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08 RtrAcct"`
}

type DocumentCamt00500108 struct {
	GetTx GetTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08 GetTx"`
}

type DocumentCamt00800108 struct {
	CclTx CancelTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08 CclTx"`
}

type DocumentCamt00600108 struct {
	RtrTx ReturnTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08 RtrTx"`
}

type DocumentCamt02700108 struct {
	ClmNonRct ClaimNonReceiptV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.027.001.08 ClmNonRct"`
}

type DocumentCamt01000108 struct {
	RtrLmt ReturnLimitV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08 RtrLmt"`
}

type DocumentCamt02600108 struct {
	UblToApply UnableToApplyV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.08 UblToApply"`
}

type DocumentCamt05400108 struct {
	BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.054.001.08 BkToCstmrDbtCdtNtfctn"`
}

type DocumentCamt05200108 struct {
	BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08 BkToCstmrAcctRpt"`
}

type DocumentCamt00700108 struct {
	ModfyTx ModifyTransactionV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08 ModfyTx"`
}

type DocumentCamt03700108 struct {
	DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.037.001.08 DbtAuthstnReq"`
}

type DocumentCamt05300108 struct {
	BkToCstmrStmt BankToCustomerStatementV08 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.053.001.08 BkToCstmrStmt"`
}
