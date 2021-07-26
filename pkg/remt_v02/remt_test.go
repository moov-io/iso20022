// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, Authorisation1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader79{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, RemittanceLocation5{}.Validate())
	assert.NotNil(t, RemittanceLocationAdviceV02{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, TransactionReferences5{}.Validate())
	assert.Nil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, TaxRecordDetails1{}.Validate())
	assert.Nil(t, TaxPeriod1{}.Validate())
	assert.Nil(t, TaxRecord1{}.Validate())
	assert.NotNil(t, TaxAmount1{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AmountType3Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.Nil(t, DatePeriodDetails{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.Nil(t, CashAccount24{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.NotNil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, Garnishment1{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.Nil(t, GroupHeader62{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.Nil(t, OriginalPaymentInformation6{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PaymentTypeInformation19{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAdviceV02{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation12{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation13{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation4{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TransactionReferences4{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.Nil(t, ExchangeRate1{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalPersonIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 PreferredContactMethod1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.NotNil(t, type5.Validate())
	type5 = "CELL"
	assert.Nil(t, type5.Validate())

	var type6 RemittanceLocationMethod2Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.NotNil(t, type6.Validate())
	type6 = "FAXI"
	assert.Nil(t, type6.Validate())

	var type7 ExternalAccountIdentification1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 TaxRecordPeriod1Code
	assert.NotNil(t, type8.Validate())
	type8 = "MM01"
	assert.Nil(t, type8.Validate())

	var type9 AddressType2Code
	assert.NotNil(t, type9.Validate())
	type9 = "ADDR"
	assert.Nil(t, type9.Validate())

	var type10 Authorisation1Code
	assert.NotNil(t, type10.Validate())
	type10 = "AUTH"
	assert.Nil(t, type10.Validate())

	var type11 ExternalCashAccountType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "AUTH"
	assert.Nil(t, type11.Validate())

	var type12 DocumentType6Code
	assert.NotNil(t, type12.Validate())
	type12 = "MSIN"
	assert.Nil(t, type12.Validate())

	var type13 DocumentType3Code
	assert.NotNil(t, type13.Validate())
	type13 = "RADM"
	assert.Nil(t, type13.Validate())

	var type14 ExchangeRateType1Code
	assert.NotNil(t, type14.Validate())
	type14 = "SPOT"
	assert.Nil(t, type14.Validate())

	var type15 ExternalCategoryPurpose1Code
	assert.NotNil(t, type15.Validate())
	type15 = "SPOT"
	assert.Nil(t, type15.Validate())

	var type16 ExternalDiscountAmountType1Code
	assert.NotNil(t, type16.Validate())
	type16 = "SPOT"
	assert.Nil(t, type16.Validate())

	var type17 ExternalDocumentLineType1Code
	assert.NotNil(t, type17.Validate())
	type17 = "SPOT"
	assert.Nil(t, type17.Validate())

	var type18 ExternalGarnishmentType1Code
	assert.NotNil(t, type18.Validate())
	type18 = "SPOT"
	assert.Nil(t, type18.Validate())

	var type19 ExternalLocalInstrument1Code
	assert.NotNil(t, type19.Validate())
	type19 = "SPOT"
	assert.Nil(t, type19.Validate())

	var type20 ExternalServiceLevel1Code
	assert.NotNil(t, type20.Validate())
	type20 = "SPOT"
	assert.Nil(t, type20.Validate())

	var type21 ExternalTaxAmountType1Code
	assert.NotNil(t, type21.Validate())
	type21 = "SPOT"
	assert.Nil(t, type21.Validate())

	var type22 Priority2Code
	assert.NotNil(t, type22.Validate())
	type22 = "HIGH"
	assert.Nil(t, type22.Validate())
}
