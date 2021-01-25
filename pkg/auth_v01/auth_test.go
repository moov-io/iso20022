// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

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

func TestDocumentAuth001001V01(t *testing.T) {
	sample := DocumentAuth001001V01{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAuth001001V01{
		InfReqOpng: InformationRequestOpeningV01{
			InvstgtnId: "InvstgtnId",
			LglMndtBsis: LegalMandate1{
				Prgrph: "Prgrph",
			},
			InvstgtnPrd: DateOrDateTimePeriodChoice{
				Dt: DatePeriodDetails{
					FrDt: common.ISODate(testTime),
					ToDt: common.ISODate(testTime),
				},
				DtTm: DateTimePeriodDetails{
					FrDtTm: common.ISODateTime(testTime),
					ToDtTm: common.ISODateTime(testTime),
				},
			},
			SchCrit: SearchCriteria1Choice{
				Acct: AccountAndParties1{
					Id: CashAccount25{
						Id: AccountIdentification4Choice{
							IBAN: "AA000130",
							Othr: GenericAccountIdentification1{
								Id: "Id",
							},
						},
					},
					InvstgtdPties: InvestigatedParties1Choice{
						Cd:    "OWNE",
						Prtry: "Prtry",
					},
				},
				PmtInstrm: PaymentInstrumentType1{
					CardNb: "11111111111",
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"InfReqOpng":{"InvstgtnId":"InvstgtnId","LglMndtBsis":{"Prgrph":"Prgrph"},"CnfdtltySts":false,"InvstgtnPrd":{"Dt":{"FrDt":"2014-11-12","ToDt":"2014-11-12"},"DtTm":{"FrDtTm":"2014-11-12T11:45:26.371","ToDtTm":"2014-11-12T11:45:26.371"}},"SchCrit":{"Acct":{"Id":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"InvstgtdPties":{"Cd":"OWNE","Prtry":"Prtry"}},"CstmrId":{"Pty":{}},"PmtInstrm":{"CardNb":"11111111111"}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:auth.001.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><InfReqOpng xmlns="urn:iso:std:iso:20022:tech:xsd:auth.001.001.01"><InvstgtnId>InvstgtnId</InvstgtnId><LglMndtBsis><Prgrph>Prgrph</Prgrph></LglMndtBsis><CnfdtltySts>false</CnfdtltySts><InvstgtnPrd><Dt><FrDt>2014-11-12</FrDt><ToDt>2014-11-12</ToDt></Dt><DtTm><FrDtTm>2014-11-12T11:45:26.371</FrDtTm><ToDtTm>2014-11-12T11:45:26.371</ToDtTm></DtTm></InvstgtnPrd><SchCrit><Acct><Id><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></Id><InvstgtdPties><Cd>OWNE</Cd><Prtry>Prtry</Prtry></InvstgtdPties></Acct><CstmrId><Pty></Pty></CstmrId><PmtInstrm><CardNb>11111111111</CardNb></PmtInstrm></SchCrit></InfReqOpng></Document>`)
}

func TestDocumentAuth002001V01(t *testing.T) {
	sample := DocumentAuth002001V01{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAuth002001V01{
		InfReqRspn: InformationRequestResponseV01{
			InvstgtnId: "InvstgtnId",
			RspnId:     "RspnId",
			RspnSts:    "PART",
			SchCrit: SearchCriteria1Choice{
				Acct: AccountAndParties1{
					Id: CashAccount25{
						Id: AccountIdentification4Choice{
							IBAN: "AA000130",
							Othr: GenericAccountIdentification1{
								Id: "Id",
							},
						},
					},
					InvstgtdPties: InvestigatedParties1Choice{
						Cd:    "OWNE",
						Prtry: "Prtry",
					},
				},
				PmtInstrm: PaymentInstrumentType1{
					CardNb: "11111111111",
				},
			},
		},
	}

	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"InfReqRspn":{"RspnId":"RspnId","InvstgtnId":"InvstgtnId","RspnSts":"PART","SchCrit":{"Acct":{"Id":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"InvstgtdPties":{"Cd":"OWNE","Prtry":"Prtry"}},"CstmrId":{"Pty":{}},"PmtInstrm":{"CardNb":"11111111111"}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:auth.002.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><InfReqRspn xmlns="urn:iso:std:iso:20022:tech:xsd:auth.002.001.01"><RspnId>RspnId</RspnId><InvstgtnId>InvstgtnId</InvstgtnId><RspnSts>PART</RspnSts><SchCrit><Acct><Id><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></Id><InvstgtdPties><Cd>OWNE</Cd><Prtry>Prtry</Prtry></InvstgtdPties></Acct><CstmrId><Pty></Pty></CstmrId><PmtInstrm><CardNb>11111111111</CardNb></PmtInstrm></SchCrit></InfReqRspn></Document>`)
}

func TestDocumentAuth003001V01(t *testing.T) {
	sample := DocumentAuth003001V01{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAuth003001V01{
		InfReqStsChngNtfctn: InformationRequestStatusChangeNotificationV01{
			OrgnlBizQry: "OrgnlBizQry",
		},
	}

	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"InfReqStsChngNtfctn":{"OrgnlBizQry":"OrgnlBizQry","CnfdtltySts":false}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:auth.003.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><InfReqStsChngNtfctn xmlns="urn:iso:std:iso:20022:tech:xsd:auth.003.001.01"><OrgnlBizQry>OrgnlBizQry</OrgnlBizQry><CnfdtltySts>false</CnfdtltySts></InfReqStsChngNtfctn></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, InformationRequestOpeningV01{}.Validate())
	assert.NotNil(t, InformationRequestResponseV01{}.Validate())
	assert.NotNil(t, InformationRequestStatusChangeNotificationV01{}.Validate())
	assert.NotNil(t, AccountAndParties1{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, AuthorityInvestigation2{}.Validate())
	assert.NotNil(t, AuthorityRequestType1{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, CashAccount25{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.Nil(t, CustomerIdentification1{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, DateOrDateTimePeriodChoice{}.Validate())
	assert.Nil(t, DatePeriodDetails{}.Validate())
	assert.Nil(t, DateTimePeriodDetails{}.Validate())
	assert.Nil(t, DueDate1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, InvestigatedParties1Choice{}.Validate())
	assert.NotNil(t, LegalMandate1{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.NotNil(t, PaymentInstrumentType1{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.NotNil(t, RequestType1{}.Validate())
	assert.NotNil(t, SearchCriteria1Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, InvestigationResult1Choice{}.Validate())
	assert.NotNil(t, ReturnIndicator1{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 Min8Max28NumericText
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "1111111111"
	assert.Nil(t, type1.Validate())

	var type2 InvestigationStatus1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "NOAP"
	assert.Nil(t, type2.Validate())

	var type3 TransactionRequestType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.NotNil(t, type3.Validate())
	type3 = "OREC"
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

	var type8 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalPersonIdentification1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 StatusResponse1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "PART"
	assert.Nil(t, type10.Validate())

	var type11 InvestigatedParties1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.NotNil(t, type11.Validate())
	type11 = "OWNE"
	assert.Nil(t, type11.Validate())
}
