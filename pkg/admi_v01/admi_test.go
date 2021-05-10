// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v01

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentAdmi00200101(t *testing.T) {
	sample := DocumentAdmi00200101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00200101{
		Admi00200101: Admi00200101{
			RltdRef: MessageReference{
				Ref: "Ref",
			},
			Rsn: RejectionReason2{
				RjctgPtyRsn: "RjctgPtyRsn",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"Admi00200101":{"RltdRef":{"Ref":"Ref"},"Rsn":{"RjctgPtyRsn":"RjctgPtyRsn"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00200101><admi.002.001.01><RltdRef><Ref>Ref</Ref></RltdRef><Rsn><RjctgPtyRsn>RjctgPtyRsn</RjctgPtyRsn></Rsn></admi.002.001.01></DocumentAdmi00200101>`,
		string(buf))
}

func TestDocumentAdmi00400101(t *testing.T) {
	sample := DocumentAdmi00400101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00400101{
		Admi00400101: Admi00400101{
			EvtInf: Event1{
				EvtCd: "EvtCd",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"Admi00400101":{"EvtInf":{"EvtCd":"EvtCd"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00400101><admi.004.001.01><EvtInf><EvtCd>EvtCd</EvtCd></EvtInf></admi.004.001.01></DocumentAdmi00400101>`,
		string(buf))
}

func TestDocumentAdmi00500101(t *testing.T) {
	sample := DocumentAdmi00500101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00500101{
		RptQryReq: ReportQueryRequestV01{
			MsgHdr: MessageHeader7{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"RptQryReq":{"MsgHdr":{"MsgId":"MsgId"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00500101><RptQryReq><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></RptQryReq></DocumentAdmi00500101>`,
		string(buf))
}

func TestDocumentAdmi00600101(t *testing.T) {
	sample := DocumentAdmi00600101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00600101{
		RsndReq: ResendRequestV01{
			MsgHdr: MessageHeader7{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"RsndReq":{"MsgHdr":{"MsgId":"MsgId"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00600101><RsndReq><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></RsndReq></DocumentAdmi00600101>`,
		string(buf))
}

func TestDocumentAdmi00700101(t *testing.T) {
	sample := DocumentAdmi00700101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00700101{
		RctAck: ReceiptAcknowledgementV01{
			MsgId: MessageHeader10{
				MsgId: "MsgId",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"RctAck":{"MsgId":{"MsgId":"MsgId"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00700101><RctAck><MsgId><MsgId>MsgId</MsgId></MsgId></RctAck></DocumentAdmi00700101>`,
		string(buf))
}

func TestDocumentAdmi01100101(t *testing.T) {
	sample := DocumentAdmi01100101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi01100101{
		SysEvtAck: SystemEventAcknowledgementV01{
			MsgId: "MsgId",
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"SysEvtAck":{"MsgId":"MsgId"}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi01100101><SysEvtAck><MsgId>MsgId</MsgId></SysEvtAck></DocumentAdmi01100101>`,
		string(buf))
}

func TestDocumentAdmi01700101(t *testing.T) {
	sample := DocumentAdmi01700101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi01700101{
		PrcgReq: ProcessingRequestV01{
			MsgId: "MsgId",
			Req: RequestDetails19{
				Tp: "Tp",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"PrcgReq":{"MsgId":"MsgId","Req":{"Tp":"Tp"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi01700101><PrcgReq><MsgId>MsgId</MsgId><Req><Tp>Tp</Tp></Req></PrcgReq></DocumentAdmi01700101>`,
		string(buf))
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, Admi00200101{}.Validate())
	assert.NotNil(t, MessageReference{}.Validate())
	assert.NotNil(t, RejectionReason2{}.Validate())
	assert.NotNil(t, Admi00400101{}.Validate())
	assert.NotNil(t, Event1{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountIdentificationSearchCriteria2Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, BalanceType11Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashBalance12{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.Nil(t, DateAndDateTimeSearch4Choice{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriodSearch1Choice{}.Validate())
	assert.Nil(t, DateTimePeriod1{}.Validate())
	assert.Nil(t, DateTimePeriod1Choice{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericIdentification36{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.NotNil(t, NameAndAddress5{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, PartyIdentification120Choice{}.Validate())
	assert.NotNil(t, PartyIdentification136{}.Validate())
	assert.NotNil(t, PostalAddress1{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ReportQueryCriteria2{}.Validate())
	assert.NotNil(t, ReportQueryRequestV01{}.Validate())
	assert.NotNil(t, ReportQuerySearchCriteria2{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.NotNil(t, ResendRequestV01{}.Validate())
	assert.NotNil(t, ResendSearchCriteria2{}.Validate())
	assert.NotNil(t, SequenceRange1{}.Validate())
	assert.NotNil(t, SequenceRange1Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, MessageHeader10{}.Validate())
	assert.NotNil(t, MessageReference1{}.Validate())
	assert.NotNil(t, ReceiptAcknowledgementReport2{}.Validate())
	assert.NotNil(t, ReceiptAcknowledgementV01{}.Validate())
	assert.NotNil(t, RequestHandling2{}.Validate())
	assert.NotNil(t, SystemEventAcknowledgementV01{}.Validate())
	assert.NotNil(t, NameAndAddress8{}.Validate())
	assert.NotNil(t, PartyIdentification44{}.Validate())
	assert.Nil(t, PartyIdentification59{}.Validate())
	assert.NotNil(t, PartyIdentification73Choice{}.Validate())
	assert.NotNil(t, ProcessingRequestV01{}.Validate())
	assert.NotNil(t, RequestDetails19{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalSystemBalanceType1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalSystemEventType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 BalanceCounterparty1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "BILA"
	assert.Nil(t, type8.Validate())
}
