// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AuthenticationChannel1Choice{}.Validate())
	assert.NotNil(t, Authorisation1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, CashAccount24{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, DatePeriodDetails1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, Frequency37Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader47{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, Mandate10{}.Validate())
	assert.Nil(t, MandateAdjustment1{}.Validate())
	assert.Nil(t, MandateAuthentication1{}.Validate())
	assert.NotNil(t, MandateClassification1Choice{}.Validate())
	assert.NotNil(t, MandateInitiationRequestV05{}.Validate())
	assert.NotNil(t, MandateOccurrences4{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, MandateTypeInformation2{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, ReferredMandateDocument1{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, Mandate8{}.Validate())
	assert.NotNil(t, Mandate9{}.Validate())
	assert.NotNil(t, MandateAmendment5{}.Validate())
	assert.NotNil(t, MandateAmendmentReason1{}.Validate())
	assert.NotNil(t, MandateAmendmentRequestV05{}.Validate())
	assert.NotNil(t, MandateReason1Choice{}.Validate())
	assert.NotNil(t, OriginalMandate4Choice{}.Validate())
	assert.NotNil(t, OriginalMessageInformation1{}.Validate())
	assert.NotNil(t, MandateCancellation5{}.Validate())
	assert.NotNil(t, MandateCancellationRequestV05{}.Validate())
	assert.NotNil(t, PaymentCancellationReason1{}.Validate())
	assert.Nil(t, AcceptanceResult6{}.Validate())
	assert.Nil(t, Mandate11{}.Validate())
	assert.Nil(t, MandateAcceptance5{}.Validate())
	assert.NotNil(t, MandateAcceptanceReportV05{}.Validate())
	assert.NotNil(t, OriginalMandate5Choice{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.Nil(t, Cheque7{}.Validate())
	assert.NotNil(t, CreditTransferTransaction22{}.Validate())
	assert.NotNil(t, CreditorPaymentActivationRequestV05{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.Nil(t, DatePeriodDetails{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.NotNil(t, Garnishment1{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GroupHeader45{}.Validate())
	assert.NotNil(t, NameAndAddress10{}.Validate())
	assert.NotNil(t, PaymentIdentification1{}.Validate())
	assert.NotNil(t, PaymentInstruction19{}.Validate())
	assert.Nil(t, PaymentTypeInformation19{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.Nil(t, RegulatoryAuthority2{}.Validate())
	assert.Nil(t, RegulatoryReporting3{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation11{}.Validate())
	assert.Nil(t, RemittanceLocation4{}.Validate())
	assert.NotNil(t, RemittanceLocationDetails1{}.Validate())
	assert.NotNil(t, StructuredRegulatoryReporting3{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation13{}.Validate())
	assert.Nil(t, TaxAmount1{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation3{}.Validate())
	assert.Nil(t, TaxInformation4{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod1{}.Validate())
	assert.Nil(t, TaxRecord1{}.Validate())
	assert.NotNil(t, TaxRecordDetails1{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent1{}.Validate())
	assert.NotNil(t, ChequeDeliveryMethod1Choice{}.Validate())
	assert.NotNil(t, Charges2{}.Validate())
	assert.NotNil(t, CreditorPaymentActivationRequestStatusReportV05{}.Validate())
	assert.NotNil(t, GroupHeader46{}.Validate())
	assert.NotNil(t, NumberOfTransactionsPerStatus3{}.Validate())
	assert.NotNil(t, OriginalGroupInformation25{}.Validate())
	assert.Nil(t, OriginalPaymentInstruction19{}.Validate())
	assert.Nil(t, OriginalTransactionReference23{}.Validate())
	assert.Nil(t, PaymentTransaction69{}.Validate())
	assert.NotNil(t, StatusReason6Choice{}.Validate())
	assert.Nil(t, StatusReasonInformation9{}.Validate())
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

	var type5 ExternalAuthenticationChannel1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type10 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type12 ExternalLocalInstrument1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalMandateSetupReason1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalPersonIdentification1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type17 ExternalMandateReason1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 ExternalServiceLevel1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type24 Frequency10Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "NEVR"
	assert.Nil(t, type24.Validate())

	var type29 DocumentType6Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MSIN"
	assert.Nil(t, type29.Validate())

	var type30 Frequency6Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "YEAR"
	assert.Nil(t, type30.Validate())

	var type31 SequenceType2Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "RCUR"
	assert.Nil(t, type31.Validate())

	var type32 TaxRecordPeriod1Code
	assert.NotNil(t, type32.Validate())
	type32 = "MM01"
	assert.Nil(t, type32.Validate())

	var type33 RemittanceLocationMethod2Code
	assert.NotNil(t, type33.Validate())
	type33 = "FAXI"
	assert.Nil(t, type33.Validate())

	var type34 AddressType2Code
	assert.NotNil(t, type34.Validate())
	type34 = "ADDR"
	assert.Nil(t, type34.Validate())

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

	var type50 TransactionGroupStatus3Code
	assert.NotNil(t, type50.Validate())
	type50 = "ACTC"
	assert.Nil(t, type50.Validate())

	var type51 TransactionIndividualStatus3Code
	assert.NotNil(t, type51.Validate())
	type51 = "ACTC"
	assert.Nil(t, type51.Validate())
}
