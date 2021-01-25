// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/stretchr/testify/assert"
)

const (
	testTimeString = "2014-11-12T11:45:26.371Z"
)

func TestDocumentCamt02100106(t *testing.T) {
	sample := DocumentCamt02100106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02100106{
		RtrGnlBizInf: ReturnGeneralBusinessInformationV06{
			MsgHdr: MessageHeader7{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"RtrGnlBizInf":{"MsgHdr":{"MsgId":"MsgId"},"RptOrErr":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.021.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrGnlBizInf xmlns="urn:iso:std:iso:20022:tech:xsd:camt.021.001.06"><MsgHdr><MsgId>MsgId</MsgId></MsgHdr><RptOrErr></RptOrErr></RtrGnlBizInf></Document>`)
}

func TestDocumentCamt02400106(t *testing.T) {
	sample := DocumentCamt02400106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02400106{
		ModfyStgOrdr: ModifyStandingOrderV06{
			MsgHdr: MessageHeader1{
				MsgId: "MsgId",
			},
			StgOrdrId: StandingOrderIdentification4{
				Acct: CashAccount38{
					Id: AccountIdentification4Choice{
						IBAN: "AA000130",
						Othr: GenericAccountIdentification1{
							Id: "Id",
						},
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ModfyStgOrdr":{"MsgHdr":{"MsgId":"MsgId"},"StgOrdrId":{"Acct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}},"NewStgOrdrValSet":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.024.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ModfyStgOrdr xmlns="urn:iso:std:iso:20022:tech:xsd:camt.024.001.06"><MsgHdr><MsgId>MsgId</MsgId></MsgHdr><StgOrdrId><Acct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></Acct></StgOrdrId><NewStgOrdrValSet></NewStgOrdrValSet></ModfyStgOrdr></Document>`)
}

func TestDocumentCamt03100106(t *testing.T) {
	sample := DocumentCamt03100106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03100106{
		RjctInvstgtn: RejectInvestigationV06{
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Justfn: InvestigationRejectionJustification1{
				RjctnRsn: "MROI",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"RjctInvstgtn":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Justfn":{"RjctnRsn":"MROI"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.031.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RjctInvstgtn xmlns="urn:iso:std:iso:20022:tech:xsd:camt.031.001.06"><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><Justfn><RjctnRsn>MROI</RjctnRsn></Justfn></RjctInvstgtn></Document>`)
}

func TestDocumentCamt04700106(t *testing.T) {
	sample := DocumentCamt04700106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt04700106{
		RtrRsvatn: ReturnReservationV06{
			MsgHdr: MessageHeader7{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"RtrRsvatn":{"MsgHdr":{"MsgId":"Id"},"RptOrErr":{"BizRpt":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.047.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RtrRsvatn xmlns="urn:iso:std:iso:20022:tech:xsd:camt.047.001.06"><MsgHdr><MsgId>Id</MsgId></MsgHdr><RptOrErr><BizRpt></BizRpt></RptOrErr></RtrRsvatn></Document>`)
}

func TestDocumentCamt05700106(t *testing.T) {
	sample := DocumentCamt05700106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt05700106{
		NtfctnToRcv: NotificationToReceiveV06{
			GrpHdr: GroupHeader77{
				MsgId:   "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Ntfctn: AccountNotification16{
				Id: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"NtfctnToRcv":{"GrpHdr":{"MsgId":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"Ntfctn":{"Id":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.057.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><NtfctnToRcv xmlns="urn:iso:std:iso:20022:tech:xsd:camt.057.001.06"><GrpHdr><MsgId>Id</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr><Ntfctn><Id>Id</Id></Ntfctn></NtfctnToRcv></Document>`)
}

func TestDocumentCamt03300106(t *testing.T) {
	sample := DocumentCamt03300106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03300106{
		ReqForDplct: RequestForDuplicateV06{
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqForDplct":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.033.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqForDplct xmlns="urn:iso:std:iso:20022:tech:xsd:camt.033.001.06"><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt></ReqForDplct></Document>`)
}

func TestDocumentCamt05800106(t *testing.T) {
	sample := DocumentCamt05800106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt05800106{
		NtfctnToRcvCxlAdvc: NotificationToReceiveCancellationAdviceV06{
			GrpHdr: GroupHeader77{
				MsgId:   "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			OrgnlNtfctn: OriginalNotification12{
				OrgnlMsgId:    "Id",
				OrgnlNtfctnId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"NtfctnToRcvCxlAdvc":{"GrpHdr":{"MsgId":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"OrgnlNtfctn":{"OrgnlMsgId":"Id","OrgnlNtfctnId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.058.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><NtfctnToRcvCxlAdvc xmlns="urn:iso:std:iso:20022:tech:xsd:camt.058.001.06"><GrpHdr><MsgId>Id</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr><OrgnlNtfctn><OrgnlMsgId>Id</OrgnlMsgId><OrgnlNtfctnId>Id</OrgnlNtfctnId></OrgnlNtfctn></NtfctnToRcvCxlAdvc></Document>`)
}

func TestDocumentCamt05900106(t *testing.T) {
	sample := DocumentCamt05900106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt05900106{
		NtfctnToRcvStsRpt: NotificationToReceiveStatusReportV06{
			GrpHdr: GroupHeader84{
				MsgId:   "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			OrgnlNtfctnAndSts: OriginalNotification11{
				OrgnlMsgId:    "Id",
				OrgnlNtfctnId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"NtfctnToRcvStsRpt":{"GrpHdr":{"MsgId":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"OrgnlNtfctnAndSts":{"OrgnlMsgId":"Id","OrgnlNtfctnId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.059.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><NtfctnToRcvStsRpt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.059.001.06"><GrpHdr><MsgId>Id</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></GrpHdr><OrgnlNtfctnAndSts><OrgnlMsgId>Id</OrgnlMsgId><OrgnlNtfctnId>Id</OrgnlNtfctnId></OrgnlNtfctnAndSts></NtfctnToRcvStsRpt></Document>`)
}

func TestDocumentCamt03400106(t *testing.T) {
	sample := DocumentCamt03400106{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03400106{
		Dplct: DuplicateV06{
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Dplct: ProprietaryData7{
				Tp: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Dplct":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Dplct":{"Tp":"Id","Data":{"Any":{"Item":""}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.034.001.06" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Dplct xmlns="urn:iso:std:iso:20022:tech:xsd:camt.034.001.06"><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><Dplct><Tp>Id</Tp><Data><Any><Item></Item></Any></Data></Dplct></Dplct></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, ErrorHandling3Choice{}.Validate())
	assert.NotNil(t, ErrorHandling5{}.Validate())
	assert.Nil(t, GeneralBusinessInformation1{}.Validate())
	assert.Nil(t, GeneralBusinessOrError7Choice{}.Validate())
	assert.Nil(t, GeneralBusinessOrError8Choice{}.Validate())
	assert.NotNil(t, GeneralBusinessReport6{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.Nil(t, InformationQualifierType1{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.NotNil(t, ReturnGeneralBusinessInformationV06{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriod2Choice{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, ExecutionType1Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ModifyStandingOrderV06{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.Nil(t, StandingOrder7{}.Validate())
	assert.NotNil(t, StandingOrderIdentification4{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, InvestigationRejectionJustification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, RejectInvestigationV06{}.Validate())
	assert.NotNil(t, RequestForDuplicateV06{}.Validate())
	assert.NotNil(t, DuplicateV06{}.Validate())
	assert.Nil(t, ProprietaryData6{}.Validate())
	assert.NotNil(t, ProprietaryData7{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.Nil(t, CurrentAndDefaultReservation4{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.NotNil(t, Reservation3{}.Validate())
	assert.NotNil(t, ReservationIdentification2{}.Validate())
	assert.Nil(t, ReservationOrError8Choice{}.Validate())
	assert.NotNil(t, ReservationOrError9Choice{}.Validate())
	assert.NotNil(t, ReservationReport6{}.Validate())
	assert.NotNil(t, ReservationStatus1Choice{}.Validate())
	assert.NotNil(t, ReservationType1Choice{}.Validate())
	assert.NotNil(t, ReturnReservationV06{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.NotNil(t, AccountNotification16{}.Validate())
	assert.NotNil(t, NotificationItem7{}.Validate())
	assert.NotNil(t, NotificationToReceiveV06{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.Nil(t, RemittanceLocation7{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
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
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
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
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GroupHeader77{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.NotNil(t, NotificationToReceiveCancellationAdviceV06{}.Validate())
	assert.NotNil(t, OriginalItem6{}.Validate())
	assert.Nil(t, OriginalItemReference5{}.Validate())
	assert.NotNil(t, OriginalNotification12{}.Validate())
	assert.Nil(t, OriginalNotificationReference10{}.Validate())
	assert.NotNil(t, GroupHeader84{}.Validate())
	assert.NotNil(t, NotificationToReceiveStatusReportV06{}.Validate())
	assert.NotNil(t, OriginalItemAndStatus6{}.Validate())
	assert.NotNil(t, OriginalNotification11{}.Validate())
	assert.Nil(t, OriginalNotificationReference9{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalSystemErrorHandling1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalAccountIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalCashAccountType1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalProxyAccountType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalSystemEventType1Code
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

	var type12 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type12.Validate())
	type12 = "tes"
	assert.Nil(t, type12.Validate())

	var type13 ExternalDiscountAmountType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalDocumentLineType1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalGarnishmentType1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalPurpose1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 ExternalTaxAmountType1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 Priority1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "NORM"
	assert.Nil(t, type18.Validate())

	var type19 Frequency2Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "YEAR"
	assert.Nil(t, type19.Validate())

	var type20 InvestigationRejection1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "NFND"
	assert.Nil(t, type20.Validate())

	var type21 PreferredContactMethod1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "LETT"
	assert.Nil(t, type21.Validate())

	var type22 ReservationStatus1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "ENAB"
	assert.Nil(t, type22.Validate())

	var type23 ReservationType2Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "CARE"
	assert.Nil(t, type23.Validate())

	var type24 DocumentType3Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "RADM"
	assert.Nil(t, type24.Validate())

	var type25 DocumentType6Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "MSIN"
	assert.Nil(t, type25.Validate())

	var type26 RemittanceLocationMethod2Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "FAXI"
	assert.Nil(t, type26.Validate())

	var type27 TaxRecordPeriod1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "MM01"
	assert.Nil(t, type27.Validate())

	var type28 NotificationStatus3Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "RCBD"
	assert.Nil(t, type28.Validate())
}
