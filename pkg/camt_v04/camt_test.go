// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
