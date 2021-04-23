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
	assert.Equal(t, string(buf), `{"Admi00200101":{"RltdRef":{"Ref":"Ref"},"Rsn":{"RjctgPtyRsn":"RjctgPtyRsn"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.002.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><admi.002.001.01><RltdRef><Ref>Ref</Ref></RltdRef><Rsn><RjctgPtyRsn>RjctgPtyRsn</RjctgPtyRsn></Rsn></admi.002.001.01></Document>`)
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
	assert.Equal(t, string(buf), `{"Admi00400101":{"EvtInf":{"EvtCd":"EvtCd"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.004.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><admi.004.001.01><EvtInf><EvtCd>EvtCd</EvtCd></EvtInf></admi.004.001.01></Document>`)
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
	assert.Equal(t, string(buf), `{"RptQryReq":{"MsgHdr":{"MsgId":"MsgId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.005.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RptQryReq><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></RptQryReq></Document>`)
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
	assert.Equal(t, string(buf), `{"RsndReq":{"MsgHdr":{"MsgId":"MsgId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.006.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RsndReq><MsgHdr><MsgId>MsgId</MsgId></MsgHdr></RsndReq></Document>`)
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
	assert.Equal(t, string(buf), `{"RctAck":{"MsgId":{"MsgId":"MsgId"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.007.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RctAck><MsgId><MsgId>MsgId</MsgId></MsgId></RctAck></Document>`)
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
	assert.Equal(t, string(buf), `{"SysEvtAck":{"MsgId":"MsgId"}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.011.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><SysEvtAck><MsgId>MsgId</MsgId></SysEvtAck></Document>`)
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
	assert.Equal(t, string(buf), `{"PrcgReq":{"MsgId":"MsgId","Req":{"Tp":"Tp"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.017.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><PrcgReq><MsgId>MsgId</MsgId><Req><Tp>Tp</Tp></Req></PrcgReq></Document>`)
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
