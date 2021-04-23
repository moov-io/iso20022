// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestDocumentCamt00300107(t *testing.T) {
	sample := DocumentCamt00400108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt00400108{
		Xmlns: sample.NameSpace(),
		RtrAcct: ReturnAccountV08{
			MsgHdr: MessageHeader7{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.004.001.08","RtrAcct":{"MsgHdr":{"MsgId":"Id"},"RptOrErr":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.004.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrAcct><MsgHdr><MsgId>Id</MsgId></MsgHdr><RptOrErr></RptOrErr></RtrAcct></Document>`)
}

func TestDocumentCamt00500108(t *testing.T) {
	sample := DocumentCamt00500108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt00500108{
		Xmlns: sample.NameSpace(),
		GetTx: GetTransactionV08{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.005.001.08","GetTx":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.005.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetTx><MsgHdr><MsgId>Id</MsgId></MsgHdr></GetTx></Document>`)
}

func TestDocumentCamt00800108(t *testing.T) {
	sample := DocumentCamt00800108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentCamt00800108{
		Xmlns: sample.NameSpace(),
		CclTx: CancelTransactionV08{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
			PmtId: PaymentIdentification6Choice{
				TxId: "TxId",
				LngBizId: LongPaymentIdentification2{
					IntrBkSttlmDt: common.ISODate(testTime),
				},
				QId: QueueTransactionIdentification1{
					QId:    "QId",
					PosInQ: "QId",
				},
				PrtryId: "PrtryId",
				ShrtBizId: ShortPaymentIdentification2{
					TxId:          "TxId",
					IntrBkSttlmDt: common.ISODate(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.008.001.08","CclTx":{"MsgHdr":{"MsgId":"Id"},"PmtId":{"TxId":"TxId","QId":{"QId":"QId","PosInQ":"QId"},"LngBizId":{"IntrBkSttlmAmt":0,"IntrBkSttlmDt":"2014-11-12","InstgAgt":{"FinInstnId":{}},"InstdAgt":{"FinInstnId":{}}},"ShrtBizId":{"TxId":"TxId","IntrBkSttlmDt":"2014-11-12","InstgAgt":{"FinInstnId":{}}},"PrtryId":"PrtryId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.008.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CclTx><MsgHdr><MsgId>Id</MsgId></MsgHdr><PmtId><TxId>TxId</TxId><QId><QId>QId</QId><PosInQ>QId</PosInQ></QId><LngBizId><IntrBkSttlmAmt>0</IntrBkSttlmAmt><IntrBkSttlmDt>2014-11-12</IntrBkSttlmDt><InstgAgt><FinInstnId></FinInstnId></InstgAgt><InstdAgt><FinInstnId></FinInstnId></InstdAgt></LngBizId><ShrtBizId><TxId>TxId</TxId><IntrBkSttlmDt>2014-11-12</IntrBkSttlmDt><InstgAgt><FinInstnId></FinInstnId></InstgAgt></ShrtBizId><PrtryId>PrtryId</PrtryId></PmtId></CclTx></Document>`)
}

func TestDocumentCamt00600108(t *testing.T) {
	sample := DocumentCamt00600108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt00600108{
		Xmlns: sample.NameSpace(),
		RtrTx: ReturnTransactionV08{
			MsgHdr: MessageHeader8{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.006.001.08","RtrTx":{"MsgHdr":{"MsgId":"Id"},"RptOrErr":{"BizRpt":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.006.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrTx><MsgHdr><MsgId>Id</MsgId></MsgHdr><RptOrErr><BizRpt></BizRpt></RptOrErr></RtrTx></Document>`)
}

func TestDocumentCamt02700108(t *testing.T) {
	sample := DocumentCamt02700108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02700108{
		Xmlns: sample.NameSpace(),
		ClmNonRct: ClaimNonReceiptV08{
			Assgnmt: CaseAssignment5{
				Id: "Id",
			},
			Undrlyg: UnderlyingTransaction6Choice{
				Initn: UnderlyingPaymentInstruction6{
					OrgnlInstdAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
				IntrBk: UnderlyingPaymentTransaction5{
					OrgnlIntrBkSttlmAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.027.001.08","ClmNonRct":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"0001-01-01T00:00:00"},"Undrlyg":{"Initn":{"OrgnlInstdAmt":{"Value":0,"Ccy":"ABC"}},"IntrBk":{"OrgnlIntrBkSttlmAmt":{"Value":0,"Ccy":"ABC"},"OrgnlIntrBkSttlmDt":"0001-01-01"},"StmtNtry":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.027.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ClmNonRct><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>0001-01-01T00:00:00</CreDtTm></Assgnmt><Undrlyg><Initn><OrgnlInstdAmt Ccy="ABC">0</OrgnlInstdAmt></Initn><IntrBk><OrgnlIntrBkSttlmAmt Ccy="ABC">0</OrgnlIntrBkSttlmAmt><OrgnlIntrBkSttlmDt>0001-01-01</OrgnlIntrBkSttlmDt></IntrBk><StmtNtry></StmtNtry></Undrlyg></ClmNonRct></Document>`)
}

func TestDocumentCamt01000108(t *testing.T) {
	sample := DocumentCamt01000108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01000108{
		Xmlns: sample.NameSpace(),
		RtrLmt: ReturnLimitV08{
			MsgHdr: MessageHeader7{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.010.001.08","RtrLmt":{"MsgHdr":{"MsgId":"Id"},"RptOrErr":{"BizRpt":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.010.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrLmt><MsgHdr><MsgId>Id</MsgId></MsgHdr><RptOrErr><BizRpt></BizRpt></RptOrErr></RtrLmt></Document>`)
}

func TestDocumentCamt02600108(t *testing.T) {
	sample := DocumentCamt02600108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02600108{
		Xmlns: sample.NameSpace(),
		UblToApply: UnableToApplyV08{
			Assgnmt: CaseAssignment5{
				Id: "Id",
			},
			Undrlyg: UnderlyingTransaction6Choice{
				Initn: UnderlyingPaymentInstruction6{
					OrgnlInstdAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
				IntrBk: UnderlyingPaymentTransaction5{
					OrgnlIntrBkSttlmAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.026.001.08","UblToApply":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"0001-01-01T00:00:00"},"Undrlyg":{"Initn":{"OrgnlInstdAmt":{"Value":0,"Ccy":"ABC"}},"IntrBk":{"OrgnlIntrBkSttlmAmt":{"Value":0,"Ccy":"ABC"},"OrgnlIntrBkSttlmDt":"0001-01-01"},"StmtNtry":{}},"Justfn":{"AnyInf":false,"MssngOrIncrrctInf":{},"PssblDplctInstr":false}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.026.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><UblToApply><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>0001-01-01T00:00:00</CreDtTm></Assgnmt><Undrlyg><Initn><OrgnlInstdAmt Ccy="ABC">0</OrgnlInstdAmt></Initn><IntrBk><OrgnlIntrBkSttlmAmt Ccy="ABC">0</OrgnlIntrBkSttlmAmt><OrgnlIntrBkSttlmDt>0001-01-01</OrgnlIntrBkSttlmDt></IntrBk><StmtNtry></StmtNtry></Undrlyg><Justfn><AnyInf>false</AnyInf><MssngOrIncrrctInf></MssngOrIncrrctInf><PssblDplctInstr>false</PssblDplctInstr></Justfn></UblToApply></Document>`)
}

func TestDocumentCamt05400108(t *testing.T) {
	sample := DocumentCamt05400108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentCamt05400108{
		Xmlns: sample.NameSpace(),
		BkToCstmrDbtCdtNtfctn: BankToCustomerDebitCreditNotificationV08{
			GrpHdr: GroupHeader81{
				MsgId:   "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.054.001.08","BkToCstmrDbtCdtNtfctn":{"GrpHdr":{"MsgId":"Id","CreDtTm":"2014-11-12T11:45:26.371"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.054.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><BkToCstmrDbtCdtNtfctn><GrpHdr><MsgId>Id</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></BkToCstmrDbtCdtNtfctn></Document>`)
}

func TestDocumentCamt05200108(t *testing.T) {
	sample := DocumentCamt05200108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentCamt05200108{
		Xmlns: sample.NameSpace(),
		BkToCstmrAcctRpt: BankToCustomerAccountReportV08{
			GrpHdr: GroupHeader81{
				MsgId:   "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08","BkToCstmrAcctRpt":{"GrpHdr":{"MsgId":"Id","CreDtTm":"2014-11-12T11:45:26.371"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><BkToCstmrAcctRpt><GrpHdr><MsgId>Id</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></BkToCstmrAcctRpt></Document>`)
}

func TestDocumentCamt00700108(t *testing.T) {
	sample := DocumentCamt00700108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt00700108{
		Xmlns: sample.NameSpace(),
		ModfyTx: ModifyTransactionV08{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.007.001.08","ModfyTx":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.007.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ModfyTx><MsgHdr><MsgId>Id</MsgId></MsgHdr></ModfyTx></Document>`)
}

func TestDocumentCamt03700108(t *testing.T) {
	sample := DocumentCamt03700108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt03700108{
		Xmlns: sample.NameSpace(),
		DbtAuthstnReq: DebitAuthorisationRequestV08{
			Assgnmt: CaseAssignment5{
				Id: "Id",
			},
			Undrlyg: UnderlyingTransaction6Choice{
				Initn: UnderlyingPaymentInstruction6{
					OrgnlInstdAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
				IntrBk: UnderlyingPaymentTransaction5{
					OrgnlIntrBkSttlmAmt: ActiveOrHistoricCurrencyAndAmount{
						Ccy: "ABC",
					},
				},
			},
			Dtl: DebitAuthorisation2{
				CxlRsn: CancellationReason33Choice{
					Cd:    "111",
					Prtry: "Prtry",
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.037.001.08","DbtAuthstnReq":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"0001-01-01T00:00:00"},"Undrlyg":{"Initn":{"OrgnlInstdAmt":{"Value":0,"Ccy":"ABC"}},"IntrBk":{"OrgnlIntrBkSttlmAmt":{"Value":0,"Ccy":"ABC"},"OrgnlIntrBkSttlmDt":"0001-01-01"},"StmtNtry":{}},"Dtl":{"CxlRsn":{"Cd":"111","Prtry":"Prtry"}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.037.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><DbtAuthstnReq><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>0001-01-01T00:00:00</CreDtTm></Assgnmt><Undrlyg><Initn><OrgnlInstdAmt Ccy="ABC">0</OrgnlInstdAmt></Initn><IntrBk><OrgnlIntrBkSttlmAmt Ccy="ABC">0</OrgnlIntrBkSttlmAmt><OrgnlIntrBkSttlmDt>0001-01-01</OrgnlIntrBkSttlmDt></IntrBk><StmtNtry></StmtNtry></Undrlyg><Dtl><CxlRsn><Cd>111</Cd><Prtry>Prtry</Prtry></CxlRsn></Dtl></DbtAuthstnReq></Document>`)
}

func TestDocumentCamt05300108(t *testing.T) {
	sample := DocumentCamt05300108{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentCamt05300108{
		Xmlns: sample.NameSpace(),
		BkToCstmrStmt: BankToCustomerStatementV08{
			GrpHdr: GroupHeader81{
				MsgId:   "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.053.001.08","BkToCstmrStmt":{"GrpHdr":{"MsgId":"Id","CreDtTm":"2014-11-12T11:45:26.371"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.053.001.08" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><BkToCstmrStmt><GrpHdr><MsgId>Id</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr></BkToCstmrStmt></Document>`)
}

func Test1NestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.Nil(t, AccountOrBusinessError4Choice{}.Validate())
	assert.Nil(t, AccountOrOperationalError4Choice{}.Validate())
	assert.NotNil(t, AccountReport24{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.NotNil(t, BalanceRestrictionType1{}.Validate())
	assert.NotNil(t, BalanceType11Choice{}.Validate())
	assert.NotNil(t, BalanceType9Choice{}.Validate())
	assert.NotNil(t, BilateralLimit3{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.Nil(t, CashAccount37{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CashBalance11{}.Validate())
	assert.NotNil(t, CashBalance13{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, DatePeriodDetails1{}.Validate())
	assert.NotNil(t, ErrorHandling3Choice{}.Validate())
	assert.NotNil(t, ErrorHandling5{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, ExecutionType1Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, Limit5{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProcessingType1Choice{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.NotNil(t, ReturnAccountV08{}.Validate())
	assert.NotNil(t, StandingOrder6{}.Validate())
	assert.Nil(t, StandingOrderTotalAmount1{}.Validate())
	assert.NotNil(t, StandingOrderType1Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, TotalAmountAndCurrency1{}.Validate())
	assert.Nil(t, AccountCashEntryReturnCriteria3{}.Validate())
	assert.NotNil(t, AccountIdentificationSearchCriteria2Choice{}.Validate())
	assert.NotNil(t, ActiveAmountRange3Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmountRange3{}.Validate())
	assert.NotNil(t, ActiveOrHistoricAmountRange2Choice{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmountRange2{}.Validate())
	assert.Nil(t, AmountRangeBoundary1{}.Validate())
	assert.Nil(t, CashAccountEntrySearch6{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification3Choice{}.Validate())
	assert.Nil(t, DateAndDateTimeSearch3Choice{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriodSearch1Choice{}.Validate())
	assert.Nil(t, DateTimePeriod1{}.Validate())
	assert.Nil(t, DateTimePeriod1Choice{}.Validate())
	assert.Nil(t, FromToAmountRange1{}.Validate())
	assert.NotNil(t, GetTransactionV08{}.Validate())
	assert.Nil(t, ImpliedCurrencyAmountRange1Choice{}.Validate())
	assert.Nil(t, ImpliedCurrencyAndAmountRange1{}.Validate())
	assert.Nil(t, InstructionStatusReturnCriteria1{}.Validate())
	assert.Nil(t, InstructionStatusSearch5{}.Validate())
	assert.Nil(t, LongPaymentIdentification2{}.Validate())
	assert.NotNil(t, MessageHeader9{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.NotNil(t, PaymentIdentification6Choice{}.Validate())
	assert.NotNil(t, PaymentOrigin1Choice{}.Validate())
	assert.Nil(t, PaymentReturnCriteria4{}.Validate())
	assert.Nil(t, PaymentSearch8{}.Validate())
	assert.NotNil(t, PaymentStatusCodeSearch2Choice{}.Validate())
	assert.Nil(t, PaymentTransactionParty3{}.Validate())
	assert.NotNil(t, PaymentType4Choice{}.Validate())
	assert.NotNil(t, Priority1Choice{}.Validate())
	assert.NotNil(t, QueueTransactionIdentification1{}.Validate())
	assert.NotNil(t, ShortPaymentIdentification2{}.Validate())
	assert.Nil(t, SystemReturnCriteria2{}.Validate())
	assert.Nil(t, SystemSearch4{}.Validate())
	assert.NotNil(t, TransactionCriteria5Choice{}.Validate())
	assert.Nil(t, TransactionCriteria8{}.Validate())
	assert.Nil(t, TransactionQuery5{}.Validate())
	assert.Nil(t, TransactionReturnCriteria5{}.Validate())
	assert.Nil(t, TransactionSearchCriteria8{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, Amount3Choice{}.Validate())
	assert.NotNil(t, CashAccount39{}.Validate())
	assert.NotNil(t, CashAccountAndEntry3{}.Validate())
	assert.Nil(t, CashEntry2{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.NotNil(t, MessageHeader8{}.Validate())
	assert.Nil(t, NumberAndSumOfTransactions2{}.Validate())
	assert.NotNil(t, Pagination1{}.Validate())
	assert.Nil(t, PaymentCommon4{}.Validate())
	assert.Nil(t, PaymentInstruction32{}.Validate())
	assert.Nil(t, PaymentStatus6{}.Validate())
	assert.NotNil(t, PaymentStatusCode6Choice{}.Validate())
	assert.NotNil(t, PaymentStatusReason1Choice{}.Validate())
	assert.NotNil(t, ProprietaryStatusJustification2{}.Validate())
	assert.NotNil(t, ReturnTransactionV08{}.Validate())
	assert.Nil(t, SecuritiesTransactionReferences1{}.Validate())
	assert.Nil(t, System2{}.Validate())
	assert.Nil(t, Transaction66{}.Validate())
	assert.Nil(t, TransactionOrError4Choice{}.Validate())
	assert.NotNil(t, TransactionReport5{}.Validate())
	assert.Nil(t, TransactionReportOrError4Choice{}.Validate())
	assert.Nil(t, Transactions8{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ModifyTransactionV08{}.Validate())
	assert.Nil(t, PaymentInstruction33{}.Validate())
	assert.NotNil(t, TransactionModification5{}.Validate())
	assert.NotNil(t, CancelTransactionV08{}.Validate())
	assert.NotNil(t, CancellationReason33Choice{}.Validate())
	assert.Nil(t, PaymentCancellationReason5{}.Validate())
	assert.NotNil(t, Limit7{}.Validate())
	assert.NotNil(t, LimitIdentification5{}.Validate())
	assert.NotNil(t, LimitOrError4Choice{}.Validate())
	assert.NotNil(t, LimitReport7{}.Validate())
}

func Test2NestedTypes(t *testing.T) {
	assert.Nil(t, LimitReportOrError4Choice{}.Validate())
	assert.NotNil(t, LimitType1Choice{}.Validate())
	assert.Nil(t, Limits7{}.Validate())
	assert.NotNil(t, ReturnLimitV08{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.Nil(t, AmendmentInformationDetails13{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.Nil(t, CreditTransferMandateData1{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, MandateClassification1Choice{}.Validate())
	assert.Nil(t, MandateRelatedData1Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation14{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, MandateTypeInformation2{}.Validate())
	assert.Nil(t, MissingOrIncorrectInformation3{}.Validate())
	assert.NotNil(t, OriginalGroupInformation29{}.Validate())
	assert.Nil(t, OriginalTransactionReference31{}.Validate())
	assert.Nil(t, PaymentTypeInformation27{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SettlementInstruction7{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.NotNil(t, ModifyTransactionV08{}.Validate())
	assert.Nil(t, PaymentInstruction33{}.Validate())
	assert.NotNil(t, TransactionModification5{}.Validate())
	assert.NotNil(t, CancelTransactionV08{}.Validate())
	assert.NotNil(t, CancellationReason33Choice{}.Validate())
	assert.Nil(t, PaymentCancellationReason5{}.Validate())
	assert.NotNil(t, Limit7{}.Validate())
	assert.NotNil(t, LimitIdentification5{}.Validate())
	assert.NotNil(t, LimitOrError4Choice{}.Validate())
	assert.NotNil(t, LimitReport7{}.Validate())
	assert.Nil(t, LimitReportOrError4Choice{}.Validate())
	assert.NotNil(t, LimitType1Choice{}.Validate())
	assert.Nil(t, Limits7{}.Validate())
	assert.NotNil(t, ReturnLimitV08{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.Nil(t, AmendmentInformationDetails13{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.Nil(t, CreditTransferMandateData1{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
}

func Test3NestedTypes(t *testing.T) {
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, MandateClassification1Choice{}.Validate())
	assert.Nil(t, MandateRelatedData1Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation14{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, MandateTypeInformation2{}.Validate())
	assert.Nil(t, MissingOrIncorrectInformation3{}.Validate())
	assert.NotNil(t, OriginalGroupInformation29{}.Validate())
	assert.Nil(t, OriginalTransactionReference31{}.Validate())
	assert.Nil(t, PaymentTypeInformation27{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SettlementInstruction7{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
	assert.NotNil(t, UnableToApplyIncorrect1{}.Validate())
	assert.Nil(t, UnableToApplyJustification3Choice{}.Validate())
	assert.NotNil(t, UnableToApplyMissing1{}.Validate())
	assert.NotNil(t, UnableToApplyV08{}.Validate())
	assert.NotNil(t, UnderlyingGroupInformation1{}.Validate())
	assert.NotNil(t, UnderlyingPaymentInstruction6{}.Validate())
	assert.NotNil(t, UnderlyingPaymentTransaction5{}.Validate())
	assert.Nil(t, UnderlyingStatementEntry3{}.Validate())
	assert.NotNil(t, UnderlyingTransaction6Choice{}.Validate())
	assert.NotNil(t, ClaimNonReceiptV08{}.Validate())
	assert.Nil(t, InstructionForAssignee1{}.Validate())
	assert.Nil(t, MissingCover4{}.Validate())
	assert.Nil(t, SettlementInstruction6{}.Validate())
	assert.NotNil(t, DebitAuthorisation2{}.Validate())
	assert.NotNil(t, DebitAuthorisationRequestV08{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, AccountInterest4{}.Validate())
	assert.NotNil(t, AccountReport25{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAnd13DecimalAmount{}.Validate())
	assert.Nil(t, AmountAndCurrencyExchange3{}.Validate())
	assert.NotNil(t, AmountAndCurrencyExchangeDetails3{}.Validate())
	assert.NotNil(t, AmountAndCurrencyExchangeDetails4{}.Validate())
	assert.NotNil(t, AmountAndDirection35{}.Validate())
	assert.NotNil(t, BalanceSubType1Choice{}.Validate())
	assert.NotNil(t, BalanceType10Choice{}.Validate())
	assert.NotNil(t, BalanceType13{}.Validate())
	assert.NotNil(t, BankToCustomerAccountReportV08{}.Validate())
	assert.Nil(t, BankTransactionCodeStructure4{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure5{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure6{}.Validate())
	assert.Nil(t, BatchInformation2{}.Validate())
	assert.Nil(t, CardAggregated2{}.Validate())
	assert.Nil(t, CardEntry4{}.Validate())
	assert.Nil(t, CardIndividualTransaction2{}.Validate())
	assert.NotNil(t, CardSecurityInformation1{}.Validate())
	assert.Nil(t, CardSequenceNumberRange1{}.Validate())
	assert.Nil(t, CardTransaction17{}.Validate())
	assert.Nil(t, CardTransaction3Choice{}.Validate())
	assert.NotNil(t, CardholderAuthentication2{}.Validate())
	assert.NotNil(t, CashAvailability1{}.Validate())
	assert.NotNil(t, CashAvailabilityDate1Choice{}.Validate())
	assert.NotNil(t, CashBalance8{}.Validate())
	assert.NotNil(t, CashDeposit1{}.Validate())
	assert.NotNil(t, ChargeType3Choice{}.Validate())
	assert.Nil(t, Charges6{}.Validate())
	assert.NotNil(t, ChargesRecord3{}.Validate())
	assert.NotNil(t, CorporateAction9{}.Validate())
	assert.Nil(t, CreditLine3{}.Validate())
	assert.NotNil(t, CreditLineType1Choice{}.Validate())
	assert.NotNil(t, CurrencyExchange5{}.Validate())
	assert.Nil(t, DateOrDateTimePeriod1Choice{}.Validate())
	assert.NotNil(t, DisplayCapabilities1{}.Validate())
	assert.Nil(t, EntryDetails9{}.Validate())
	assert.NotNil(t, EntryStatus1Choice{}.Validate())
	assert.Nil(t, EntryTransaction10{}.Validate())
	assert.Nil(t, FinancialInstrumentQuantity1Choice{}.Validate())
	assert.NotNil(t, GenericIdentification3{}.Validate())
	assert.NotNil(t, GenericIdentification32{}.Validate())
	assert.NotNil(t, GroupHeader81{}.Validate())
	assert.NotNil(t, IdentificationSource3Choice{}.Validate())
	assert.NotNil(t, InterestRecord2{}.Validate())
	assert.NotNil(t, InterestType1Choice{}.Validate())
	assert.Nil(t, MessageIdentification2{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.Nil(t, NumberAndSumOfTransactions1{}.Validate())
	assert.Nil(t, NumberAndSumOfTransactions4{}.Validate())
	assert.Nil(t, OriginalAndCurrentQuantities1{}.Validate())
	assert.NotNil(t, OtherIdentification1{}.Validate())
	assert.Nil(t, PaymentCard4{}.Validate())
	assert.Nil(t, PaymentContext3{}.Validate())
	assert.Nil(t, PaymentReturnReason5{}.Validate())
	assert.NotNil(t, PlainCardData1{}.Validate())
	assert.NotNil(t, PointOfInteraction1{}.Validate())
	assert.Nil(t, PointOfInteractionCapabilities1{}.Validate())
	assert.NotNil(t, PointOfInteractionComponent1{}.Validate())
	assert.NotNil(t, Price7{}.Validate())
	assert.NotNil(t, PriceRateOrAmount3Choice{}.Validate())
	assert.NotNil(t, Product2{}.Validate())
	assert.NotNil(t, ProprietaryAgent4{}.Validate())
	assert.NotNil(t, ProprietaryBankTransactionCodeStructure1{}.Validate())
	assert.NotNil(t, ProprietaryDate3{}.Validate())
	assert.NotNil(t, ProprietaryParty5{}.Validate())
	assert.NotNil(t, ProprietaryPrice2{}.Validate())
	assert.NotNil(t, ProprietaryQuantity1{}.Validate())
	assert.NotNil(t, ProprietaryReference1{}.Validate())
	assert.NotNil(t, Rate4{}.Validate())
	assert.NotNil(t, RateType4Choice{}.Validate())
	assert.Nil(t, RemittanceLocation7{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.NotNil(t, ReportEntry10{}.Validate())
	assert.NotNil(t, ReportingSource1Choice{}.Validate())
	assert.NotNil(t, ReturnReason5Choice{}.Validate())
	assert.NotNil(t, SecuritiesAccount19{}.Validate())
	assert.Nil(t, SecurityIdentification19{}.Validate())
	assert.NotNil(t, SequenceRange1{}.Validate())
	assert.NotNil(t, SequenceRange1Choice{}.Validate())
	assert.Nil(t, TaxCharges2{}.Validate())
	assert.Nil(t, TaxInformation8{}.Validate())
	assert.NotNil(t, TechnicalInputChannel1Choice{}.Validate())
	assert.Nil(t, TotalTransactions6{}.Validate())
	assert.Nil(t, TotalsPerBankTransactionCode5{}.Validate())
	assert.NotNil(t, TrackData1{}.Validate())
	assert.Nil(t, TransactionAgents5{}.Validate())
	assert.Nil(t, TransactionDates3{}.Validate())
	assert.NotNil(t, TransactionIdentifier1{}.Validate())
	assert.Nil(t, TransactionInterest4{}.Validate())
	assert.Nil(t, TransactionParties6{}.Validate())
	assert.NotNil(t, TransactionPrice4Choice{}.Validate())
	assert.NotNil(t, TransactionQuantities3Choice{}.Validate())
	assert.Nil(t, TransactionReferences6{}.Validate())
	assert.NotNil(t, YieldedOrValueType1Choice{}.Validate())
	assert.NotNil(t, AccountStatement9{}.Validate())
	assert.NotNil(t, BankToCustomerStatementV08{}.Validate())
	assert.NotNil(t, AccountNotification17{}.Validate())
	assert.NotNil(t, BankToCustomerDebitCreditNotificationV08{}.Validate())
}

func Test1Types(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalCashAccountType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalPersonIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalProxyAccountType1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalSystemBalanceType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalSystemErrorHandling1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalSystemEventType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalCashClearingSystem1Code
	assert.NotNil(t, type13.Validate())
	type13 = "tes"
	assert.Nil(t, type13.Validate())

	var type14 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type14.Validate())
	type14 = "tes"
	assert.Nil(t, type14.Validate())

	var type15 ExternalCancellationReason1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalCategoryPurpose1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 ExternalDiscountAmountType1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 ExternalDocumentLineType1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type19 ExternalGarnishmentType1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.Nil(t, type19.Validate())

	var type20 ExternalLocalInstrument1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.Nil(t, type20.Validate())

	var type21 ExternalMandateSetupReason1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.Nil(t, type21.Validate())

	var type22 ExternalPurpose1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.Nil(t, type22.Validate())

	var type23 ExternalServiceLevel1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.Nil(t, type23.Validate())

	var type24 ExternalTaxAmountType1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.Nil(t, type24.Validate())

	var type25 ExternalAgentInstruction1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.Nil(t, type25.Validate())

	var type26 ExternalBalanceSubType1Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.Nil(t, type26.Validate())

	var type27 ExternalBalanceType1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.Nil(t, type27.Validate())

	var type28 ExternalBankTransactionDomain1Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.Nil(t, type28.Validate())

	var type29 ExternalBankTransactionFamily1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.Nil(t, type29.Validate())

	var type30 ExternalBankTransactionSubFamily1Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.Nil(t, type30.Validate())

	var type31 ExternalCardTransactionCategory1Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.Nil(t, type31.Validate())

	var type32 ExternalChargeType1Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.Nil(t, type32.Validate())
}

func Test2Types(t *testing.T) {
	var type1 ExternalCreditLineType1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalEntryStatus1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalFinancialInstrumentIdentificationType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalRePresentmentReason1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalReportingSource1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalReturnReason1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalTechnicalInputChannel1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 BalanceStatus1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "STLD"
	assert.Nil(t, type8.Validate())

	var type9 Frequency2Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.NotNil(t, type9.Validate())
	type9 = "OVNG"
	assert.Nil(t, type9.Validate())

	var type10 PreferredContactMethod1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "CELL"
	assert.Nil(t, type10.Validate())

	var type11 ProcessingType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.NotNil(t, type11.Validate())
	type11 = "CSDB"
	assert.Nil(t, type11.Validate())

	var type12 StandingOrderType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.NotNil(t, type12.Validate())
	type12 = "PSTO"
	assert.Nil(t, type12.Validate())

	var type13 SystemBalanceType2Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.NotNil(t, type13.Validate())
	type13 = "OPNG"
	assert.Nil(t, type13.Validate())

	var type14 CashPaymentStatus2Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.NotNil(t, type14.Validate())
	type14 = "FINL"
	assert.Nil(t, type14.Validate())

	var type15 EntryStatus1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "FUTR"
	assert.Nil(t, type15.Validate())

	var type16 FinalStatusCode
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.NotNil(t, type16.Validate())
	type16 = "FNLD"
	assert.Nil(t, type16.Validate())

	var type17 Instruction1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "TFRO"
	assert.Nil(t, type17.Validate())

	var type18 PaymentInstrument1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "CAN"
	assert.Nil(t, type18.Validate())

	var type19 PaymentType3Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "CBS"
	assert.Nil(t, type19.Validate())

	var type20 PendingStatus4Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "ACPD"
	assert.Nil(t, type20.Validate())

	var type21 Priority5Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "URGT"
	assert.Nil(t, type21.Validate())

	var type22 QueryType2Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "DELD"
	assert.Nil(t, type22.Validate())

	var type23 ReportIndicator1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "PRPR"
	assert.Nil(t, type23.Validate())

	var type24 CancelledStatusReason1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "CSUB"
	assert.Nil(t, type24.Validate())

	var type25 FinalStatus1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "FNLD"
	assert.Nil(t, type25.Validate())

	var type26 PendingFailingSettlement1Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "AWMO"
	assert.Nil(t, type26.Validate())

	var type27 PendingSettlement2Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "AWMO"
	assert.Nil(t, type27.Validate())

	var type28 SuspendedStatusReason1Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "SUBS"
	assert.Nil(t, type28.Validate())

	var type29 UnmatchedStatusReason1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "CMIS"
	assert.Nil(t, type29.Validate())

	var type30 LimitStatus1Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "ENAB"
	assert.Nil(t, type30.Validate())
}

func Test3Types(t *testing.T) {
	var type1 LimitType3Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "MULT"
	assert.Nil(t, type1.Validate())

	var type2 ClearingChannel2Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "RTGS"
	assert.Nil(t, type2.Validate())

	var type3 DocumentType3Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.NotNil(t, type3.Validate())
	type3 = "RADM"
	assert.Nil(t, type3.Validate())

	var type4 DocumentType6Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.NotNil(t, type4.Validate())
	type4 = "MSIN"
	assert.Nil(t, type4.Validate())

	var type5 Frequency6Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.NotNil(t, type5.Validate())
	type5 = "YEAR"
	assert.Nil(t, type5.Validate())

	var type6 PaymentMethod4Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.NotNil(t, type6.Validate())
	type6 = "CHK"
	assert.Nil(t, type6.Validate())

	var type7 Priority2Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.NotNil(t, type7.Validate())
	type7 = "NORM"
	assert.Nil(t, type7.Validate())

	var type8 SequenceType3Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "FRST"
	assert.Nil(t, type8.Validate())

	var type9 SettlementMethod1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.NotNil(t, type9.Validate())
	type9 = "INDA"
	assert.Nil(t, type9.Validate())

	var type10 TaxRecordPeriod1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "MM01"
	assert.Nil(t, type10.Validate())

	var type11 UnableToApplyIncorrectInformation4Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.NotNil(t, type11.Validate())
	type11 = "IN01"
	assert.Nil(t, type11.Validate())

	var type12 UnableToApplyMissingInformation3Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.NotNil(t, type12.Validate())
	type12 = "MS01"
	assert.Nil(t, type12.Validate())

	var type13 AttendanceContext1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.NotNil(t, type13.Validate())
	type13 = "ATTD"
	assert.Nil(t, type13.Validate())

	var type14 AuthenticationEntity1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.NotNil(t, type14.Validate())
	type14 = "ICCD"
	assert.Nil(t, type14.Validate())

	var type15 AuthenticationMethod1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "UKNW"
	assert.Nil(t, type15.Validate())

	var type16 CSCManagement1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.NotNil(t, type16.Validate())
	type16 = "PRST"
	assert.Nil(t, type16.Validate())

	var type17 CardDataReading1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "TAGC"
	assert.Nil(t, type17.Validate())

	var type18 CardPaymentServiceType2Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "AGGR"
	assert.Nil(t, type18.Validate())

	var type19 CardholderVerificationCapability1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "MNSG"
	assert.Nil(t, type19.Validate())

	var type20 ChargeBearerType1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "DEBT"
	assert.Nil(t, type20.Validate())

	var type21 OnLineCapability1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "OFLN"
	assert.Nil(t, type21.Validate())

	var type22 POIComponentType1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "SOFT"
	assert.Nil(t, type22.Validate())

	var type23 PartyType3Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "OPOI"
	assert.Nil(t, type23.Validate())

	var type24 PartyType4Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "MERC"
	assert.Nil(t, type24.Validate())

	var type25 PriceValueType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "DISC"
	assert.Nil(t, type25.Validate())

	var type26 RemittanceLocationMethod2Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "FAXI"
	assert.Nil(t, type26.Validate())

	var type27 TransactionChannel1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "MAIL"
	assert.Nil(t, type27.Validate())

	var type28 TransactionEnvironment1Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "PUBL"
	assert.Nil(t, type28.Validate())

	var type29 UnitOfMeasure1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "PIEC"
	assert.Nil(t, type29.Validate())

	var type30 UserInterface2Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "MDSP"
	assert.Nil(t, type30.Validate())

	var type31 Exact1NumericText
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "000"
	assert.Nil(t, type31.Validate())

	var type32 Exact3NumericText
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "000"
	assert.Nil(t, type32.Validate())

	var type33 ISINOct2015Identifier
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.NotNil(t, type33.Validate())
	type33 = "AA0000000001"
	assert.Nil(t, type33.Validate())

	var type34 ISO2ALanguageCode
	assert.NotNil(t, type34.Validate())
	type34 = "AA"
	assert.NotNil(t, type34.Validate())
	type34 = "aa"
	assert.Nil(t, type34.Validate())

	var type35 EntryTypeIdentifier
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "B00DUM"
	assert.Nil(t, type35.Validate())
}
