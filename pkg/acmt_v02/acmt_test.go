// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestDocumentAcmt02200102(t *testing.T) {
	sample := DocumentAcmt02200102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt02200102{
		Xmlns: sample.NameSpace(),
		IdModAdvc: IdentificationModificationAdviceV02{
			Assgnmt: IdentificationAssignment2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02","IdModAdvc":{"Assgnmt":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}}}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><IdModAdvc><Assgnmt><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne></Assgnmt></IdModAdvc></Document>`, string(buf))
}

func TestDocumentAcmt02300102(t *testing.T) {
	sample := DocumentAcmt02300102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt02300102{
		Xmlns: sample.NameSpace(),
		IdVrfctnReq: IdentificationVerificationRequestV02{
			Assgnmt: IdentificationAssignment2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02","IdVrfctnReq":{"Assgnmt":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}}}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.023.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><IdVrfctnReq><Assgnmt><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne></Assgnmt></IdVrfctnReq></Document>`, string(buf))
}

func TestDocumentAcmt02400102(t *testing.T) {
	sample := DocumentAcmt02400102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt02400102{
		Xmlns: sample.NameSpace(),
		IdVrfctnRpt: IdentificationVerificationReportV02{
			Assgnmt: IdentificationAssignment2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.024.001.02","IdVrfctnRpt":{"Assgnmt":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","Assgnr":{"Pty":{},"Agt":{"FinInstnId":{}}},"Assgne":{"Pty":{},"Agt":{"FinInstnId":{}}}}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.024.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><IdVrfctnRpt><Assgnmt><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><Assgnr><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgnr><Assgne><Pty></Pty><Agt><FinInstnId></FinInstnId></Agt></Assgne></Assgnmt></IdVrfctnRpt></Document>`, string(buf))
}

func TestDocumentAcmt03000102(t *testing.T) {
	sample := DocumentAcmt03000102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt03000102{
		Xmlns: sample.NameSpace(),
		AcctSwtchReqRdrctn: AccountSwitchRequestRedirectionV02{
			MsgId: MessageIdentification1{
				Id:      "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			NewAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{Id: "Id"},
				},
			},
			OdAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{Id: "Id"},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.030.001.02","AcctSwtchReqRdrctn":{"MsgId":{"Id":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"NewAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"OdAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.030.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchReqRdrctn><MsgId><Id>MsgId</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><NewAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></NewAcct><OdAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></OdAcct></AcctSwtchReqRdrctn></Document>`, string(buf))
}

func TestDocumentAcmt03300102(t *testing.T) {
	sample := DocumentAcmt03300102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt03300102{
		Xmlns: sample.NameSpace(),
		AcctSwtchNtfyAcctSwtchCmplt: AccountSwitchNotifyAccountSwitchCompleteV02{
			MsgId: MessageIdentification1{
				Id:      "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.033.001.02","AcctSwtchNtfyAcctSwtchCmplt":{"MsgId":{"Id":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.033.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchNtfyAcctSwtchCmplt><MsgId><Id>MsgId</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls></AcctSwtchNtfyAcctSwtchCmplt></Document>`, string(buf))
}

func TestDocumentAcmt03500102(t *testing.T) {
	sample := DocumentAcmt03500102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt03500102{
		Xmlns: sample.NameSpace(),
		AcctSwtchPmtRspn: AccountSwitchPaymentResponseV02{
			MsgId: MessageIdentification1{
				Id:      "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.035.001.02","AcctSwtchPmtRspn":{"MsgId":{"Id":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.035.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchPmtRspn><MsgId><Id>MsgId</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls></AcctSwtchPmtRspn></Document>`, string(buf))
}

func TestDocumentAcmt03700102(t *testing.T) {
	sample := DocumentAcmt03700102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentAcmt03700102{
		Xmlns: sample.NameSpace(),
		AcctSwtchTechRjctn: AccountSwitchTechnicalRejectionV02{
			MsgId: MessageIdentification1{
				Id:      "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.037.001.02","AcctSwtchTechRjctn":{"MsgId":{"Id":"MsgId","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.037.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchTechRjctn><MsgId><Id>MsgId</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls></AcctSwtchTechRjctn></Document>`, string(buf))
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

	var type3 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalPersonIdentification1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalVerificationReason1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 BalanceTransferWindow1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.NotNil(t, type7.Validate())
	type7 = "EARL"
	assert.Nil(t, type7.Validate())

	var type8 PreferredContactMethod1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "CELL"
	assert.Nil(t, type8.Validate())

	var type9 SwitchStatus1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.NotNil(t, type9.Validate())
	type9 = "TMTN"
	assert.Nil(t, type9.Validate())

	var type10 SwitchType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "PART"
	assert.Nil(t, type10.Validate())

	var type11 ExternalCashAccountType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalProxyAccountType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, IdentificationAssignment2{}.Validate())
	assert.Nil(t, IdentificationInformation2{}.Validate())
	assert.NotNil(t, IdentificationModification2{}.Validate())
	assert.NotNil(t, IdentificationModificationAdviceV02{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, OriginalTransactionReference18{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, Party12Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.NotNil(t, PaymentIdentification4{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, IdentificationVerification2{}.Validate())
	assert.NotNil(t, IdentificationVerificationRequestV02{}.Validate())
	assert.NotNil(t, IdentificationVerificationReportV02{}.Validate())
	assert.Nil(t, MessageIdentification5{}.Validate())
	assert.NotNil(t, VerificationReason1Choice{}.Validate())
	assert.NotNil(t, VerificationReport2{}.Validate())
	assert.NotNil(t, AccountSwitchDetails1{}.Validate())
	assert.NotNil(t, AccountSwitchRequestRedirectionV02{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashAccount39{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, MessageIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, AccountSwitchNotifyAccountSwitchCompleteV02{}.Validate())
	assert.NotNil(t, AccountSwitchPaymentResponseV02{}.Validate())
	assert.NotNil(t, AccountSwitchTechnicalRejectionV02{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
}
