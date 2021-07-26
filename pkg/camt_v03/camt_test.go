// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
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

	var type4 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalProxyAccountType1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 QueryType2Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.NotNil(t, type6.Validate())
	type6 = "DELD"
	assert.Nil(t, type6.Validate())

	var type7 StandingOrderQueryType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.NotNil(t, type7.Validate())
	type7 = "SWLS"
	assert.Nil(t, type7.Validate())

	var type8 StandingOrderType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "PSTO"
	assert.Nil(t, type8.Validate())

	var type9 ExternalBankTransactionDomain1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalBankTransactionFamily1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalBankTransactionSubFamily1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalBillingBalanceType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalBillingCompensationType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalBillingRateIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 CompensationMethod1Code
	assert.NotNil(t, type16.Validate())
	type15 = "test"
	assert.NotNil(t, type16.Validate())
	type16 = "DDBT"
	assert.Nil(t, type16.Validate())

	var type17 BillingSubServiceQualifier1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "MACT"
	assert.Nil(t, type17.Validate())

	var type18 BillingTaxCalculationMethod1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "UDFD"
	assert.Nil(t, type18.Validate())

	var type19 BillingStatementStatus1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "TEST"
	assert.Nil(t, type19.Validate())

	var type20 BillingCurrencyType1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "PRCG"
	assert.Nil(t, type20.Validate())

	var type21 BillingCurrencyType2Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "HOST"
	assert.Nil(t, type21.Validate())

	var type22 BillingChargeMethod1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "UPRC"
	assert.Nil(t, type22.Validate())

	var type23 BalanceAdjustmentType1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "CLLD"
	assert.Nil(t, type23.Validate())

	var type24 AccountLevel1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "SMRY"
	assert.Nil(t, type24.Validate())

	var type29 AccountLevel2Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "DETL"
	assert.Nil(t, type29.Validate())

	var type27 PreferredContactMethod1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "CELL"
	assert.Nil(t, type27.Validate())

	var type25 ServiceAdjustmentType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "NCMP"
	assert.Nil(t, type25.Validate())

	var type26 ServicePaymentMethod1Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "FREE"
	assert.Nil(t, type26.Validate())

	var type28 ServiceTaxDesignation1Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "TAXE"
	assert.Nil(t, type28.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriod2Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GetStandingOrderV03{}.Validate())
	assert.NotNil(t, MessageHeader4{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, RequestType3Choice{}.Validate())
	assert.Nil(t, StandingOrderCriteria3{}.Validate())
	assert.NotNil(t, StandingOrderCriteria3Choice{}.Validate())
	assert.Nil(t, StandingOrderQuery3{}.Validate())
	assert.Nil(t, StandingOrderReturnCriteria1{}.Validate())
	assert.Nil(t, StandingOrderSearchCriteria3{}.Validate())
	assert.NotNil(t, StandingOrderType1Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, DeleteStandingOrderV03{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, StandingOrderIdentification4{}.Validate())
	assert.NotNil(t, StandingOrderIdentification5{}.Validate())
	assert.Nil(t, StandingOrderOrAll2Choice{}.Validate())
	assert.NotNil(t, AccountTax1{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AmountAndDirection34{}.Validate())
	assert.NotNil(t, BalanceAdjustment1{}.Validate())
	assert.NotNil(t, BankServicesBillingStatementV03{}.Validate())
	assert.Nil(t, BankTransactionCodeStructure4{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure5{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure6{}.Validate())
	assert.NotNil(t, BillingBalance1{}.Validate())
	assert.NotNil(t, BillingBalanceType1Choice{}.Validate())
	assert.NotNil(t, BillingCompensation1{}.Validate())
	assert.NotNil(t, BillingCompensationType1Choice{}.Validate())
	assert.NotNil(t, BillingMethod1{}.Validate())
	assert.NotNil(t, BillingMethod1Choice{}.Validate())
	assert.NotNil(t, BillingMethod2{}.Validate())
	assert.NotNil(t, BillingMethod3{}.Validate())
	assert.NotNil(t, BillingMethod4{}.Validate())
	assert.Nil(t, BillingPrice1{}.Validate())
	assert.NotNil(t, BillingRate1{}.Validate())
	assert.NotNil(t, BillingRateIdentification1Choice{}.Validate())
	assert.NotNil(t, BillingService2{}.Validate())
	assert.NotNil(t, BillingServiceAdjustment1{}.Validate())
	assert.NotNil(t, BillingServiceCommonIdentification1{}.Validate())
	assert.NotNil(t, BillingServiceIdentification2{}.Validate())
	assert.NotNil(t, BillingServiceIdentification3{}.Validate())
	assert.NotNil(t, BillingServiceParameters2{}.Validate())
	assert.NotNil(t, BillingServiceParameters3{}.Validate())
	assert.NotNil(t, BillingServicesAmount1{}.Validate())
	assert.NotNil(t, BillingServicesAmount2{}.Validate())
	assert.NotNil(t, BillingServicesAmount3{}.Validate())
	assert.NotNil(t, BillingServicesTax1{}.Validate())
	assert.NotNil(t, BillingServicesTax2{}.Validate())
	assert.NotNil(t, BillingServicesTax3{}.Validate())
	assert.NotNil(t, BillingStatement3{}.Validate())
	assert.NotNil(t, BillingSubServiceIdentification1{}.Validate())
	assert.NotNil(t, BillingSubServiceQualifier1Choice{}.Validate())
	assert.Nil(t, BillingTaxIdentification2{}.Validate())
	assert.NotNil(t, BillingTaxRegion2{}.Validate())
	assert.NotNil(t, CashAccountCharacteristics3{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, CurrencyExchange6{}.Validate())
	assert.Nil(t, DatePeriod1{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification19{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.NotNil(t, Pagination1{}.Validate())
	assert.NotNil(t, ParentCashAccount3{}.Validate())
	assert.Nil(t, Party43Choice{}.Validate())
	assert.NotNil(t, PartyIdentification138{}.Validate())
	assert.NotNil(t, ProprietaryBankTransactionCodeStructure1{}.Validate())
	assert.NotNil(t, ReportHeader6{}.Validate())
	assert.NotNil(t, ResidenceLocation1Choice{}.Validate())
	assert.NotNil(t, ServiceTaxDesignation1{}.Validate())
	assert.NotNil(t, StatementGroup3{}.Validate())
	assert.NotNil(t, TaxCalculation1{}.Validate())
	assert.NotNil(t, TaxReason1{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, Case3{}.Validate())
	assert.NotNil(t, CaseAssignment3{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, Party12Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, ProprietaryData3{}.Validate())
	assert.NotNil(t, ProprietaryData4{}.Validate())
	assert.NotNil(t, ProprietaryFormatInvestigationV03{}.Validate())
}
