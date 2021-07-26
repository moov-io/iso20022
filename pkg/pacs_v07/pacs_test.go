package pacs_v07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.Nil(t, AmendmentInformationDetails10{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, CashAccount24{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, Charges2{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification3Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, DatePeriodDetails{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.NotNil(t, FIToFIPaymentStatusReportV07{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, Frequency21Choice{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment1{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader53{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation10{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.NotNil(t, NumberOfTransactionsPerStatus3{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OriginalGroupHeader1{}.Validate())
	assert.NotNil(t, OriginalGroupInformation3{}.Validate())
	assert.Nil(t, OriginalTransactionReference22{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PaymentTransaction63{}.Validate())
	assert.Nil(t, PaymentTypeInformation25{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation11{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.NotNil(t, SettlementInstruction4{}.Validate())
	assert.NotNil(t, StatusReason6Choice{}.Validate())
	assert.Nil(t, StatusReasonInformation9{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation13{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, TaxAmount1{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation4{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod1{}.Validate())
	assert.Nil(t, TaxRecord1{}.Validate())
	assert.NotNil(t, TaxRecordDetails1{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalCashAccountType1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalCashClearingSystem1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalCategoryPurpose1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalDiscountAmountType1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalDocumentLineType1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalGarnishmentType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalLocalInstrument1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalMandateSetupReason1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type11.Validate())
	type11 = "tes"
	assert.Nil(t, type11.Validate())

	var type12 ExternalPersonIdentification1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalServiceLevel1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalStatusReason1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalMandateSetupReason1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalTaxAmountType1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 AddressType2Code
	assert.NotNil(t, type17.Validate())
	type17 = "DLVY"
	assert.Nil(t, type17.Validate())

	var type18 ClearingChannel2Code
	assert.NotNil(t, type18.Validate())
	type18 = "BOOK"
	assert.Nil(t, type18.Validate())

	var type19 DocumentType3Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "SCOR"
	assert.Nil(t, type19.Validate())

	var type20 DocumentType6Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "MSIN"
	assert.Nil(t, type20.Validate())

	var type21 Frequency6Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "YEAR"
	assert.Nil(t, type21.Validate())

	var type22 PaymentMethod4Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "TRA"
	assert.Nil(t, type22.Validate())

	var type23 Priority2Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "NORM"
	assert.Nil(t, type23.Validate())

	var type24 SequenceType3Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "FRST"
	assert.Nil(t, type24.Validate())

	var type25 SettlementMethod1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "INDA"
	assert.Nil(t, type25.Validate())

	var type26 TaxRecordPeriod1Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "MM01"
	assert.Nil(t, type26.Validate())

	var type27 SequenceType3Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "FRST"
	assert.Nil(t, type27.Validate())

	var type29 TaxRecordPeriod1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MM01"
	assert.Nil(t, type29.Validate())

	var type30 TransactionGroupStatus3Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "ACTC"
	assert.Nil(t, type30.Validate())

	var type31 TransactionIndividualStatus3Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "ACTC"
	assert.Nil(t, type31.Validate())
}
