// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestDocumentCamt01300104(t *testing.T) {
	sample := DocumentCamt01300104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01300104{
		Xmlns: sample.NameSpace(),
		GetMmb: GetMemberV04{
			MsgHdr: MessageHeader9{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.013.001.04","GetMmb":{"MsgHdr":{"MsgId":"MsgId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.013.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetMmb><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></GetMmb></Document>`)
}

func TestDocumentCamt01400104(t *testing.T) {
	sample := DocumentCamt01400104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01400104{
		Xmlns: sample.NameSpace(),
		RtrMmb: ReturnMemberV04{
			MsgHdr: MessageHeader7{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.014.001.04","RtrMmb":{"MsgHdr":{"MsgId":"MsgId"},"RptOrErr":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.014.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrMmb><MsgHdr><MsgId>MsgId</MsgId></MsgHdr><RptOrErr></RptOrErr></RtrMmb></Document>`)
}

func TestDocumentCamt01500104(t *testing.T) {
	sample := DocumentCamt01500104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01500104{
		Xmlns: sample.NameSpace(),
		ModfyMmb: ModifyMemberV04{
			MsgHdr: MessageHeader1{
				MsgId: "MsgId",
			},
			MmbId: MemberIdentification3Choice{
				BICFI:       "0000AA00111",
				ClrSysMmbId: ClearingSystemMemberIdentification2{MmbId: "MmbId"},
				Othr:        GenericFinancialIdentification1{Id: "MmbId"},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.015.001.04","ModfyMmb":{"MsgHdr":{"MsgId":"MsgId"},"MmbId":{"BICFI":"0000AA00111","ClrSysMmbId":{"MmbId":"MmbId"},"Othr":{"Id":"MmbId"}},"NewMmbValSet":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.015.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ModfyMmb><MsgHdr><MsgId>MsgId</MsgId></MsgHdr><MmbId><BICFI>0000AA00111</BICFI><ClrSysMmbId><MmbId>MmbId</MmbId></ClrSysMmbId><Othr><Id>MmbId</Id></Othr></MmbId><NewMmbValSet></NewMmbValSet></ModfyMmb></Document>`)
}

func TestDocumentCamt01600104(t *testing.T) {
	sample := DocumentCamt01600104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01600104{
		Xmlns: sample.NameSpace(),
		GetCcyXchgRate: GetCurrencyExchangeRateV04{
			MsgHdr: MessageHeader1{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.016.001.04","GetCcyXchgRate":{"MsgHdr":{"MsgId":"MsgId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.016.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetCcyXchgRate><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></GetCcyXchgRate></Document>`)
}

func TestDocumentCamt02000104(t *testing.T) {
	sample := DocumentCamt02000104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02000104{
		Xmlns: sample.NameSpace(),
		GetGnlBizInf: GetGeneralBusinessInformationV04{
			MsgHdr: MessageHeader1{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.020.001.04","GetGnlBizInf":{"MsgHdr":{"MsgId":"MsgId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.020.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetGnlBizInf><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></GetGnlBizInf></Document>`)
}

func TestDocumentCamt01700104(t *testing.T) {
	sample := DocumentCamt01700104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01700104{
		Xmlns: sample.NameSpace(),
		RtrCcyXchgRate: ReturnCurrencyExchangeRateV04{
			MsgHdr: MessageHeader7{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.017.001.04","RtrCcyXchgRate":{"MsgHdr":{"MsgId":"MsgId"},"RptOrErr":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.017.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrCcyXchgRate><MsgHdr><MsgId>MsgId</MsgId></MsgHdr><RptOrErr></RptOrErr></RtrCcyXchgRate></Document>`)
}

func TestDocumentCamt03200104(t *testing.T) {
	sample := DocumentCamt03200104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentCamt03200104{
		Xmlns: sample.NameSpace(),
		CclCaseAssgnmt: CancelCaseAssignmentV04{
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Case: Case5{
				Id: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.032.001.04","CclCaseAssgnmt":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Case":{"Id":"Id","Cretr":{"Pty":{},"Agt":{"FinInstnId":{}}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.032.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CclCaseAssgnmt><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><Case><Id>Id</Id><Cretr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Cretr></Case></CclCaseAssgnmt></Document>`)
}

func TestDocumentCamt03800104(t *testing.T) {
	sample := DocumentCamt03800104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentCamt03800104{
		Xmlns: sample.NameSpace(),
		CaseStsRptReq: CaseStatusReportRequestV04{
			ReqHdr: ReportHeader5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Case: Case5{
				Id: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.038.001.04","CaseStsRptReq":{"ReqHdr":{"Id":"Id","Fr":{"Pty":{},"Agt":{"FinInstnId":{}}},"To":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Case":{"Id":"Id","Cretr":{"Pty":{},"Agt":{"FinInstnId":{}}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.038.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CaseStsRptReq><ReqHdr><Id>Id</Id><Fr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Fr><To><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></To><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></ReqHdr><Case><Id>Id</Id><Cretr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Cretr></Case></CaseStsRptReq></Document>`)
}

func TestDocumentCamt07000104(t *testing.T) {
	sample := DocumentCamt07000104{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt07000104{
		Xmlns: sample.NameSpace(),
		RtrStgOrdr: ReturnStandingOrderV04{
			MsgHdr: MessageHeader6{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.070.001.04","RtrStgOrdr":{"MsgHdr":{"MsgId":"Id"},"RptOrErr":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.070.001.04" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrStgOrdr><MsgHdr><MsgId>Id</MsgId></MsgHdr><RptOrErr></RptOrErr></RtrStgOrdr></Document>`)
}

func TestTypes(t *testing.T) {
	var type1 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalSystemMemberType1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalAccountIdentification1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalCashAccountType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalPaymentRole1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalProxyAccountType1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalPersonIdentification1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalSystemErrorHandling1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalSystemEventType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 MemberStatus1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.NotNil(t, type14.Validate())
	type14 = "JOIN"
	assert.Nil(t, type14.Validate())

	var type15 QueryType2Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "DELD"
	assert.Nil(t, type15.Validate())

	var type16 ErrorHandling1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.NotNil(t, type16.Validate())
	type16 = "X050"
	assert.Nil(t, type16.Validate())

	var type17 PaymentRole1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "STMG"
	assert.Nil(t, type17.Validate())

	var type18 Priority1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "LOWW"
	assert.Nil(t, type18.Validate())

	var type19 PreferredContactMethod1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "CELL"
	assert.Nil(t, type19.Validate())

	var type20 Frequency2Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "OVNG"
	assert.Nil(t, type20.Validate())

	var type21 StandingOrderQueryType1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "SWLS"
	assert.Nil(t, type21.Validate())

	var type22 StandingOrderType1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "PSTO"
	assert.Nil(t, type22.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GetMemberV04{}.Validate())
	assert.Nil(t, MemberCriteria4{}.Validate())
	assert.NotNil(t, MemberCriteriaDefinition2Choice{}.Validate())
	assert.NotNil(t, MemberIdentification3Choice{}.Validate())
	assert.Nil(t, MemberQueryDefinition4{}.Validate())
	assert.Nil(t, MemberReturnCriteria1{}.Validate())
	assert.Nil(t, MemberSearchCriteria4{}.Validate())
	assert.NotNil(t, MessageHeader9{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, SystemMemberStatus1Choice{}.Validate())
	assert.NotNil(t, SystemMemberType1Choice{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CommunicationAddress10{}.Validate())
	assert.NotNil(t, ContactIdentificationAndAddress2{}.Validate())
	assert.NotNil(t, ErrorHandling1Choice{}.Validate())
	assert.NotNil(t, ErrorHandling3{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, Member5{}.Validate())
	assert.NotNil(t, MemberReport5{}.Validate())
	assert.Nil(t, MemberReportOrError5Choice{}.Validate())
	assert.NotNil(t, MemberReportOrError6Choice{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, PaymentRole1Choice{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, ReturnMemberV04{}.Validate())
	assert.NotNil(t, CommunicationAddress8{}.Validate())
	assert.NotNil(t, ContactIdentificationAndAddress1{}.Validate())
	assert.NotNil(t, LongPostalAddress1Choice{}.Validate())
	assert.Nil(t, Member6{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ModifyMemberV04{}.Validate())
	assert.NotNil(t, StructuredLongPostalAddress1{}.Validate())
	assert.NotNil(t, CurrencyCriteriaDefinition1Choice{}.Validate())
	assert.Nil(t, CurrencyExchangeCriteria2{}.Validate())
	assert.NotNil(t, CurrencyExchangeSearchCriteria1{}.Validate())
	assert.Nil(t, CurrencyQueryDefinition3{}.Validate())
	assert.NotNil(t, GetCurrencyExchangeRateV04{}.Validate())
	assert.NotNil(t, CurrencyExchange7{}.Validate())
	assert.NotNil(t, CurrencyExchangeReport3{}.Validate())
	assert.NotNil(t, CurrencySourceTarget1{}.Validate())
	assert.Nil(t, ExchangeRateReportOrError1Choice{}.Validate())
	assert.NotNil(t, ExchangeRateReportOrError2Choice{}.Validate())
	assert.NotNil(t, ReturnCurrencyExchangeRateV04{}.Validate())
	assert.Nil(t, BusinessInformationCriteria1{}.Validate())
	assert.Nil(t, BusinessInformationQueryDefinition3{}.Validate())
	assert.NotNil(t, CharacterSearch1Choice{}.Validate())
	assert.NotNil(t, GeneralBusinessInformationCriteriaDefinition1Choice{}.Validate())
	assert.Nil(t, GeneralBusinessInformationReturnCriteria1{}.Validate())
	assert.Nil(t, GeneralBusinessInformationSearchCriteria1{}.Validate())
	assert.NotNil(t, GetGeneralBusinessInformationV04{}.Validate())
	assert.Nil(t, InformationQualifierType1{}.Validate())
	assert.NotNil(t, CancelCaseAssignmentV04{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseStatusReportRequestV04{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ReportHeader5{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.Nil(t, DatePeriodDetails1{}.Validate())
	assert.NotNil(t, ErrorHandling3Choice{}.Validate())
	assert.NotNil(t, ErrorHandling5{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, ExecutionType1Choice{}.Validate())
	assert.NotNil(t, MessageHeader6{}.Validate())
	assert.NotNil(t, RequestType3Choice{}.Validate())
	assert.NotNil(t, ReturnStandingOrderV04{}.Validate())
	assert.NotNil(t, StandingOrder6{}.Validate())
	assert.NotNil(t, StandingOrderIdentification4{}.Validate())
	assert.Nil(t, StandingOrderOrError5Choice{}.Validate())
	assert.NotNil(t, StandingOrderOrError6Choice{}.Validate())
	assert.NotNil(t, StandingOrderReport1{}.Validate())
	assert.Nil(t, StandingOrderTotalAmount1{}.Validate())
	assert.NotNil(t, StandingOrderType1Choice{}.Validate())
	assert.Nil(t, TotalAmountAndCurrency1{}.Validate())
}
