// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AdditionalPaymentInformationV10{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, AmendmentInformationDetails13{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification3Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.Nil(t, CreditTransferMandateData1{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent3{}.Validate())
	assert.Nil(t, InstructionForNextAgent1{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, MandateClassification1Choice{}.Validate())
	assert.Nil(t, MandateRelatedData1Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation14{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, MandateTypeInformation2{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OriginalGroupInformation29{}.Validate())
	assert.Nil(t, OriginalTransactionReference31{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PaymentComplementaryInformation9{}.Validate())
	assert.Nil(t, PaymentTypeInformation27{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, SettlementInstruction7{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
	assert.NotNil(t, UnderlyingGroupInformation1{}.Validate())
	assert.NotNil(t, UnderlyingPaymentInstruction6{}.Validate())
	assert.NotNil(t, UnderlyingPaymentTransaction5{}.Validate())
	assert.Nil(t, UnderlyingStatementEntry3{}.Validate())
	assert.NotNil(t, UnderlyingTransaction6Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, CancellationStatusReason3Choice{}.Validate())
	assert.Nil(t, CancellationStatusReason4{}.Validate())
	assert.NotNil(t, ChargeType3Choice{}.Validate())
	assert.Nil(t, Charges6{}.Validate())
	assert.NotNil(t, Charges8{}.Validate())
	assert.NotNil(t, ChargesRecord3{}.Validate())
	assert.Nil(t, ClaimNonReceipt2{}.Validate())
	assert.NotNil(t, ClaimNonReceipt2Choice{}.Validate())
	assert.NotNil(t, ClaimNonReceiptRejectReason1Choice{}.Validate())
	assert.NotNil(t, Compensation3{}.Validate())
	assert.NotNil(t, CompensationReason1Choice{}.Validate())
	assert.NotNil(t, CorrectiveGroupInformation1{}.Validate())
	assert.NotNil(t, CorrectiveInterbankTransaction2{}.Validate())
	assert.NotNil(t, CorrectivePaymentInitiation4{}.Validate())
	assert.NotNil(t, CorrectiveTransaction4Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, GenericIdentification3{}.Validate())
	assert.NotNil(t, InvestigationStatus5Choice{}.Validate())
	assert.NotNil(t, ModificationStatusReason1Choice{}.Validate())
	assert.Nil(t, ModificationStatusReason2{}.Validate())
	assert.NotNil(t, NumberOfCancellationsPerStatus1{}.Validate())
	assert.NotNil(t, NumberOfTransactionsPerStatus1{}.Validate())
	assert.Nil(t, OriginalGroupHeader14{}.Validate())
	assert.Nil(t, OriginalPaymentInstruction35{}.Validate())
	assert.NotNil(t, PaymentTransaction116{}.Validate())
	assert.Nil(t, PaymentTransaction122{}.Validate())
	assert.Nil(t, PaymentTransaction127{}.Validate())
	assert.Nil(t, ResolutionData2{}.Validate())
	assert.NotNil(t, ResolutionOfInvestigationV10{}.Validate())
	assert.Nil(t, StatementResolutionEntry4{}.Validate())
	assert.NotNil(t, TaxCharges2{}.Validate())
	assert.Nil(t, UnderlyingTransaction25{}.Validate())
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

	var type3 ExternalCashClearingSystem1Code
	assert.NotNil(t, type3.Validate())
	type3 = "tes"
	assert.Nil(t, type3.Validate())

	var type4 ExternalCategoryPurpose1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalCreditorAgentInstruction1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalDiscountAmountType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalDocumentLineType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalGarnishmentType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalLocalInstrument1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalMandateSetupReason1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalPersonIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalProxyAccountType1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalPurpose1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 ExternalServiceLevel1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 ExternalTaxAmountType1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type19 ExternalChargeType1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.Nil(t, type19.Validate())

	var type20 ExternalClaimNonReceiptRejection1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.Nil(t, type20.Validate())

	var type21 ExternalInvestigationExecutionConfirmation1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.Nil(t, type21.Validate())

	var type22 ExternalPaymentCancellationRejection1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.Nil(t, type22.Validate())

	var type23 ExternalPaymentCompensationReason1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.Nil(t, type23.Validate())

	var type24 ExternalPaymentModificationRejection1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.Nil(t, type24.Validate())

	var type25 ChargeBearerType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "SLEV"
	assert.Nil(t, type25.Validate())

	var type26 ClearingChannel2Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "BOOK"
	assert.Nil(t, type26.Validate())

	var type27 DocumentType3Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "SCOR"
	assert.Nil(t, type27.Validate())

	var type28 DocumentType6Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "MSIN"
	assert.Nil(t, type28.Validate())

	var type29 Frequency6Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "YEAR"
	assert.Nil(t, type29.Validate())

	var type30 Instruction4Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "TELA"
	assert.Nil(t, type30.Validate())

	var type31 PaymentMethod4Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "TRA"
	assert.Nil(t, type31.Validate())

	var type32 PreferredContactMethod1Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "CELL"
	assert.Nil(t, type32.Validate())

	var type33 Priority2Code
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.NotNil(t, type33.Validate())
	type33 = "NORM"
	assert.Nil(t, type33.Validate())

	var type34 SequenceType3Code
	assert.NotNil(t, type34.Validate())
	type34 = "test"
	assert.NotNil(t, type34.Validate())
	type34 = "RPRE"
	assert.Nil(t, type34.Validate())

	var type35 SettlementMethod1Code
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "CLRG"
	assert.Nil(t, type35.Validate())

	var type36 TaxRecordPeriod1Code
	assert.NotNil(t, type36.Validate())
	type36 = "test"
	assert.NotNil(t, type36.Validate())
	type36 = "MM01"
	assert.Nil(t, type36.Validate())

	var type37 TransactionIndividualStatus1Code
	assert.NotNil(t, type37.Validate())
	type37 = "test"
	assert.NotNil(t, type37.Validate())
	type37 = "ACTC"
	assert.Nil(t, type37.Validate())

	var type38 CancellationIndividualStatus1Code
	assert.NotNil(t, type38.Validate())
	type38 = "test"
	assert.NotNil(t, type38.Validate())
	type38 = "PDCR"
	assert.Nil(t, type38.Validate())

	var type39 GroupCancellationStatus1Code
	assert.NotNil(t, type39.Validate())
	type39 = "test"
	assert.NotNil(t, type39.Validate())
	type39 = "PDCR"
	assert.Nil(t, type39.Validate())
}
