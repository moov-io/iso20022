// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.Nil(t, AccountIdentification4Choice{}.Validate())
	assert.Nil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.Nil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, AmountOrRate1Choice{}.Validate())
	assert.Nil(t, AmountType4Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, CashAccountType2Choice{}.Validate())
	assert.Nil(t, CategoryPurpose1Choice{}.Validate())
	assert.Nil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.Nil(t, CashAccount38{}.Validate())
	assert.Nil(t, CashAccountType2Choice{}.Validate())
	assert.Nil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, Charges7{}.Validate())
	assert.Nil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, LocalInstrument2Choice{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.Nil(t, CreditorReferenceType1Choice{}.Validate())
	assert.Nil(t, CreditorReferenceType2{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.Nil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, Document12{}.Validate())
	assert.Nil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentFormat1Choice{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.Nil(t, DocumentLineType1{}.Validate())
	assert.Nil(t, DocumentLineType1Choice{}.Validate())
	assert.Nil(t, DocumentType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.Nil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.Nil(t, Garnishment3{}.Validate())
	assert.Nil(t, GarnishmentType1{}.Validate())
	assert.Nil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader87{}.Validate())
	assert.Nil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, NumberOfTransactionsPerStatus5{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.Nil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OriginalGroupInformation30{}.Validate())
	assert.NotNil(t, OriginalPaymentInstruction31{}.Validate())
	assert.Nil(t, OriginalTransactionReference29{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyAndSignature3{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PaymentCondition1{}.Validate())
	assert.Nil(t, PaymentConditionStatus1{}.Validate())
	assert.Nil(t, PaymentTransaction104{}.Validate())
	assert.Nil(t, PaymentTypeInformation26{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.Nil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.Nil(t, ProxyAccountType1Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.Nil(t, ReferredDocumentType3Choice{}.Validate())
	assert.Nil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.Nil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.Nil(t, StatusReason6Choice{}.Validate())
	assert.Nil(t, StatusReasonInformation12{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.Nil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
	assert.Nil(t, Cheque11{}.Validate())
	assert.Nil(t, ChequeDeliveryMethod1Choice{}.Validate())
	assert.NotNil(t, CreditTransferTransaction35{}.Validate())
	assert.NotNil(t, CreditorPaymentActivationRequestV07{}.Validate())
	assert.NotNil(t, GroupHeader78{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent1{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.NotNil(t, PaymentIdentification6{}.Validate())
	assert.NotNil(t, PaymentInstruction31{}.Validate())
	assert.Nil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, RegulatoryAuthority2{}.Validate())
	assert.Nil(t, RegulatoryReporting3{}.Validate())
	assert.Nil(t, RemittanceLocation7{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.Nil(t, StructuredRegulatoryReporting3{}.Validate())
	assert.Nil(t, TaxInformation8{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalCashAccountType1Code
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

	var type10 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type12 ExternalLocalInstrument1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type14 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalPersonIdentification1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type18 ExternalServiceLevel1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type29 DocumentType6Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MSIN"
	assert.Nil(t, type29.Validate())

	var type32 TaxRecordPeriod1Code
	assert.NotNil(t, type32.Validate())
	type32 = "MM01"
	assert.Nil(t, type32.Validate())

	var type33 RemittanceLocationMethod2Code
	assert.NotNil(t, type33.Validate())
	type33 = "FAXI"
	assert.Nil(t, type33.Validate())

	var type35 ChargeBearerType1Code
	assert.NotNil(t, type35.Validate())
	type35 = "DEBT"
	assert.Nil(t, type35.Validate())

	var type36 ChequeDelivery1Code
	assert.NotNil(t, type36.Validate())
	type36 = "MLDB"
	assert.Nil(t, type36.Validate())

	var type37 ChequeType2Code
	assert.NotNil(t, type37.Validate())
	type37 = "CCHQ"
	assert.Nil(t, type37.Validate())

	var type38 DocumentType3Code
	assert.NotNil(t, type38.Validate())
	type38 = "RADM"
	assert.Nil(t, type38.Validate())

	var type39 Instruction3Code
	assert.NotNil(t, type39.Validate())
	type39 = "CHQB"
	assert.Nil(t, type39.Validate())

	var type40 Priority2Code
	assert.NotNil(t, type40.Validate())
	type40 = "HIGH"
	assert.Nil(t, type40.Validate())

	var type41 PaymentMethod7Code
	assert.NotNil(t, type41.Validate())
	type41 = "CHK"
	assert.Nil(t, type41.Validate())

	var type42 RegulatoryReportingType1Code
	assert.NotNil(t, type42.Validate())
	type42 = "CRED"
	assert.Nil(t, type42.Validate())

	var type43 ExternalDiscountAmountType1Code
	assert.NotNil(t, type43.Validate())
	type43 = "CRED"
	assert.Nil(t, type43.Validate())

	var type44 ExternalDocumentLineType1Code
	assert.NotNil(t, type44.Validate())
	type44 = "CRED"
	assert.Nil(t, type44.Validate())

	var type45 ExternalGarnishmentType1Code
	assert.NotNil(t, type45.Validate())
	type45 = "CRED"
	assert.Nil(t, type45.Validate())

	var type46 ExternalPurpose1Code
	assert.NotNil(t, type46.Validate())
	type46 = "CRED"
	assert.Nil(t, type46.Validate())

	var type47 ExternalTaxAmountType1Code
	assert.NotNil(t, type47.Validate())
	type47 = "CRED"
	assert.Nil(t, type47.Validate())

	var type48 ExternalStatusReason1Code
	assert.NotNil(t, type48.Validate())
	type48 = "CRED"
	assert.Nil(t, type48.Validate())

	var type49 PaymentMethod4Code
	assert.NotNil(t, type49.Validate())
	type49 = "CHK"
	assert.Nil(t, type49.Validate())
}
