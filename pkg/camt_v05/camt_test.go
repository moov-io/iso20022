// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

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

func TestDocumentCamt01800105(t *testing.T) {
	sample := DocumentCamt01800105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt01800105{
		Xmlns: sample.NameSpace(),
		GetBizDayInf: GetBusinessDayInformationV05{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05","GetBizDayInf":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.018.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetBizDayInf><MsgHdr><MsgId>Id</MsgId></MsgHdr></GetBizDayInf></Document>`)
}

func TestDocumentCamt02500105(t *testing.T) {
	sample := DocumentCamt02500105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt02500105{
		Xmlns: sample.NameSpace(),
		Rct: ReceiptV05{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.025.001.05","Rct":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.025.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><Rct><MsgHdr><MsgId>Id</MsgId></MsgHdr></Rct></Document>`)
}

func TestDocumentCamt03000105(t *testing.T) {
	sample := DocumentCamt03000105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03000105{
		Xmlns: sample.NameSpace(),
		NtfctnOfCaseAssgnmt: NotificationOfCaseAssignmentV05{
			Hdr: ReportHeader5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Case: Case5{
				Id: "Id",
			},
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Ntfctn: CaseForwardingNotification3{
				Justfn: "MINE",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05","NtfctnOfCaseAssgnmt":{"Hdr":{"Id":"Id","Fr":{"Pty":{},"Agt":{"FinInstnId":{}}},"To":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Case":{"Id":"Id","Cretr":{"Pty":{},"Agt":{"FinInstnId":{}}}},"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Ntfctn":{"Justfn":"MINE"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.030.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><NtfctnOfCaseAssgnmt><Hdr><Id>Id</Id><Fr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Fr><To><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></To><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Hdr><Case><Id>Id</Id><Cretr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Cretr></Case><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><Ntfctn><Justfn>MINE</Justfn></Ntfctn></NtfctnOfCaseAssgnmt></Document>`)
}

func TestDocumentCamt03500105(t *testing.T) {
	sample := DocumentCamt03500105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03500105{
		Xmlns: sample.NameSpace(),
		PrtryFrmtInvstgtn: ProprietaryFormatInvestigationV05{
			Assgnmt: CaseAssignment5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			PrtryData: ProprietaryData7{
				Tp: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.035.001.05","PrtryFrmtInvstgtn":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"PrtryData":{"Tp":"Id","Data":{"Any":{"Item":""}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.035.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><PrtryFrmtInvstgtn><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><PrtryData><Tp>Id</Tp><Data><Any><Item></Item></Any></Data></PrtryData></PrtryFrmtInvstgtn></Document>`)
}

func TestDocumentCamt04800105(t *testing.T) {
	sample := DocumentCamt04800105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt04800105{
		Xmlns: sample.NameSpace(),
		ModfyRsvatn: ModifyReservationV05{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
			RsvatnId: CurrentOrDefaultReservation2Choice{
				Cur: ReservationIdentification2{
					Tp: ReservationType1Choice{
						Cd:    "BLKD",
						Prtry: "Prtry",
					},
				},
				Dflt: ReservationIdentification2{
					Tp: ReservationType1Choice{
						Cd:    "BLKD",
						Prtry: "Prtry",
					},
				},
			},
			NewRsvatnValSet: Reservation4{
				Amt: Amount2Choice{
					AmtWthCcy: ActiveCurrencyAndAmount{
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
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.048.001.05","ModfyRsvatn":{"MsgHdr":{"MsgId":"Id"},"RsvatnId":{"Cur":{"Tp":{"Cd":"BLKD","Prtry":"Prtry"}},"Dflt":{"Tp":{"Cd":"BLKD","Prtry":"Prtry"}}},"NewRsvatnValSet":{"Amt":{"AmtWthtCcy":0,"AmtWthCcy":{"Value":0,"Ccy":"ABC"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.048.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ModfyRsvatn><MsgHdr><MsgId>Id</MsgId></MsgHdr><RsvatnId><Cur><Tp><Cd>BLKD</Cd><Prtry>Prtry</Prtry></Tp></Cur><Dflt><Tp><Cd>BLKD</Cd><Prtry>Prtry</Prtry></Tp></Dflt></RsvatnId><NewRsvatnValSet><Amt><AmtWthtCcy>0</AmtWthtCcy><AmtWthCcy Ccy="ABC">0</AmtWthCcy></Amt></NewRsvatnValSet></ModfyRsvatn></Document>`)
}

func TestDocumentCamt03600105(t *testing.T) {
	sample := DocumentCamt03600105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03600105{
		Xmlns: sample.NameSpace(),
		DbtAuthstnRspn: DebitAuthorisationResponseV05{
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
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.036.001.05","DbtAuthstnRspn":{"Assgnmt":{"Id":"Id","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Case":{"Id":"Id","Cretr":{"Pty":{},"Agt":{"FinInstnId":{}}}},"Conf":{"DbtAuthstn":false}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.036.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><DbtAuthstnRspn><Assgnmt><Id>Id</Id><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Assgnmt><Case><Id>Id</Id><Cretr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Cretr></Case><Conf><DbtAuthstn>false</DbtAuthstn></Conf></DbtAuthstnRspn></Document>`)
}

func TestDocumentCamt04900105(t *testing.T) {
	sample := DocumentCamt04900105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt04900105{
		Xmlns: sample.NameSpace(),
		DelRsvatn: DeleteReservationV05{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.049.001.05","DelRsvatn":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.049.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><DelRsvatn><MsgHdr><MsgId>Id</MsgId></MsgHdr></DelRsvatn></Document>`)
}

func TestDocumentCamt05000105(t *testing.T) {
	sample := DocumentCamt05000105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt05000105{
		Xmlns: sample.NameSpace(),
		LqdtyCdtTrf: LiquidityCreditTransferV05{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
			LqdtyCdtTrf: LiquidityCreditTransfer2{
				TrfdAmt: Amount2Choice{
					AmtWthCcy: ActiveCurrencyAndAmount{
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
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.050.001.05","LqdtyCdtTrf":{"MsgHdr":{"MsgId":"Id"},"LqdtyCdtTrf":{"TrfdAmt":{"AmtWthtCcy":0,"AmtWthCcy":{"Value":0,"Ccy":"ABC"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.050.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><LqdtyCdtTrf><MsgHdr><MsgId>Id</MsgId></MsgHdr><LqdtyCdtTrf><TrfdAmt><AmtWthtCcy>0</AmtWthtCcy><AmtWthCcy Ccy="ABC">0</AmtWthCcy></TrfdAmt></LqdtyCdtTrf></LqdtyCdtTrf></Document>`)
}

func TestDocumentCamt03900105(t *testing.T) {
	sample := DocumentCamt03900105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentCamt03900105{
		Xmlns: sample.NameSpace(),
		CaseStsRpt: CaseStatusReportV05{
			Hdr: ReportHeader5{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			Case: Case5{
				Id: "Id",
			},
			Sts: CaseStatus2{
				DtTm:    common.ISODateTime(testTime),
				CaseSts: "ODUE",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.039.001.05","CaseStsRpt":{"Hdr":{"Id":"Id","Fr":{"Pty":{},"Agt":{"FinInstnId":{}}},"To":{"Pty":{},"Agt":{"FinInstnId":{}}},"CreDtTm":"2014-11-12T11:45:26.371"},"Case":{"Id":"Id","Cretr":{"Pty":{},"Agt":{"FinInstnId":{}}}},"Sts":{"DtTm":"2014-11-12T11:45:26.371","CaseSts":"ODUE"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.039.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CaseStsRpt><Hdr><Id>Id</Id><Fr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Fr><To><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></To><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></Hdr><Case><Id>Id</Id><Cretr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Cretr></Case><Sts><DtTm>2014-11-12T11:45:26.371</DtTm><CaseSts>ODUE</CaseSts></Sts></CaseStsRpt></Document>`)
}

func TestDocumentCamt04600105(t *testing.T) {
	sample := DocumentCamt04600105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt04600105{
		Xmlns: sample.NameSpace(),
		GetRsvatn: GetReservationV05{
			MsgHdr: MessageHeader9{
				MsgId: "Id",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.046.001.05","GetRsvatn":{"MsgHdr":{"MsgId":"Id"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.046.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><GetRsvatn><MsgHdr><MsgId>Id</MsgId></MsgHdr></GetRsvatn></Document>`)
}

func TestDocumentCamt05100105(t *testing.T) {
	sample := DocumentCamt05100105{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentCamt05100105{
		Xmlns: sample.NameSpace(),
		LqdtyDbtTrf: LiquidityDebitTransferV05{
			MsgHdr: MessageHeader1{
				MsgId: "Id",
			},
			LqdtyDbtTrf: LiquidityDebitTransfer2{
				TrfdAmt: Amount2Choice{
					AmtWthCcy: ActiveCurrencyAndAmount{
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
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05","LqdtyDbtTrf":{"MsgHdr":{"MsgId":"Id"},"LqdtyDbtTrf":{"TrfdAmt":{"AmtWthtCcy":0,"AmtWthCcy":{"Value":0,"Ccy":"ABC"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.051.001.05" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><LqdtyDbtTrf><MsgHdr><MsgId>Id</MsgId></MsgHdr><LqdtyDbtTrf><TrfdAmt><AmtWthtCcy>0</AmtWthtCcy><AmtWthCcy Ccy="ABC">0</AmtWthCcy></TrfdAmt></LqdtyDbtTrf></LqdtyDbtTrf></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.Nil(t, BusinessDayCriteria2{}.Validate())
	assert.NotNil(t, BusinessDayCriteria3Choice{}.Validate())
	assert.Nil(t, BusinessDayQuery2{}.Validate())
	assert.Nil(t, BusinessDayReturnCriteria2{}.Validate())
	assert.Nil(t, BusinessDaySearchCriteria2{}.Validate())
	assert.Nil(t, DateTimePeriod1{}.Validate())
	assert.Nil(t, DateTimePeriod1Choice{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GetBusinessDayInformationV05{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.NotNil(t, MessageHeader9{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, SystemEventType2Choice{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.Nil(t, LongPaymentIdentification2{}.Validate())
	assert.NotNil(t, OriginalMessageAndIssuer1{}.Validate())
	assert.NotNil(t, PaymentIdentification6Choice{}.Validate())
	assert.NotNil(t, PaymentOrigin1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, QueueTransactionIdentification1{}.Validate())
	assert.NotNil(t, Receipt3{}.Validate())
	assert.NotNil(t, ReceiptV05{}.Validate())
	assert.NotNil(t, RequestHandling1{}.Validate())
	assert.NotNil(t, ShortPaymentIdentification2{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, CaseForwardingNotification3{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, NotificationOfCaseAssignmentV05{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, ReportHeader5{}.Validate())
	assert.Nil(t, ProprietaryData6{}.Validate())
	assert.NotNil(t, ProprietaryData7{}.Validate())
	assert.NotNil(t, ProprietaryFormatInvestigationV05{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.Nil(t, DebitAuthorisationConfirmation2{}.Validate())
	assert.NotNil(t, DebitAuthorisationResponseV05{}.Validate())
	assert.NotNil(t, CaseStatus2{}.Validate())
	assert.NotNil(t, CaseStatusReportV05{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GetReservationV05{}.Validate())
	assert.NotNil(t, ReservationCriteria3Choice{}.Validate())
	assert.Nil(t, ReservationCriteria4{}.Validate())
	assert.Nil(t, ReservationQuery3{}.Validate())
	assert.Nil(t, ReservationReturnCriteria1{}.Validate())
	assert.Nil(t, ReservationSearchCriteria3{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.NotNil(t, CurrentOrDefaultReservation2Choice{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, ModifyReservationV05{}.Validate())
	assert.NotNil(t, Reservation4{}.Validate())
	assert.NotNil(t, DeleteReservationV05{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ReservationIdentification2{}.Validate())
	assert.NotNil(t, ReservationType1Choice{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, LiquidityCreditTransfer2{}.Validate())
	assert.NotNil(t, LiquidityCreditTransferV05{}.Validate())
	assert.Nil(t, PaymentIdentification8{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, LiquidityDebitTransfer2{}.Validate())
	assert.NotNil(t, LiquidityDebitTransferV05{}.Validate())
	assert.NotNil(t, AccountReportingRequestV05{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, BalanceSubType1Choice{}.Validate())
	assert.NotNil(t, BalanceType10Choice{}.Validate())
	assert.NotNil(t, BalanceType13{}.Validate())
	assert.Nil(t, DatePeriodDetails1{}.Validate())
	assert.NotNil(t, EntryStatus1Choice{}.Validate())
	assert.NotNil(t, GroupHeader77{}.Validate())
	assert.NotNil(t, Limit2{}.Validate())
	assert.NotNil(t, ReportingPeriod2{}.Validate())
	assert.NotNil(t, ReportingRequest5{}.Validate())
	assert.NotNil(t, SequenceRange1{}.Validate())
	assert.NotNil(t, SequenceRange1Choice{}.Validate())
	assert.Nil(t, TimePeriodDetails1{}.Validate())
	assert.NotNil(t, TransactionType2{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type2.Validate())
	type2 = "tes"
	assert.Nil(t, type2.Validate())

	var type3 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalClearingSystemIdentification1Code
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

	var type7 ExternalPersonIdentification1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalAccountIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalCashAccountType1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalProxyAccountType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalBalanceSubType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalBalanceType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalBalanceType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalEntryStatus1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 EntryTypeIdentifier
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "B00DUM"
	assert.Nil(t, type15.Validate())

	var type17 QueryType2Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "DELD"
	assert.Nil(t, type17.Validate())

	var type18 SystemEventType2Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "LVCO"
	assert.Nil(t, type18.Validate())

	var type19 PaymentInstrument1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "CAN"
	assert.Nil(t, type19.Validate())

	var type20 CaseForwardingNotification3Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "MINE"
	assert.Nil(t, type20.Validate())

	var type21 PreferredContactMethod1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "CELL"
	assert.Nil(t, type21.Validate())

	var type22 CaseStatus2Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "ODUE"
	assert.Nil(t, type22.Validate())

	var type23 ReservationType1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "THRE"
	assert.Nil(t, type23.Validate())

	var type24 ReservationType2Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "BLKD"
	assert.Nil(t, type24.Validate())

	var type25 FloorLimitType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "BOTH"
	assert.Nil(t, type25.Validate())

	var type26 QueryType3Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "MODF"
	assert.Nil(t, type26.Validate())
}
