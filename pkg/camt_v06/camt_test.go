// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, ErrorHandling3Choice{}.Validate())
	assert.NotNil(t, ErrorHandling5{}.Validate())
	assert.Nil(t, GeneralBusinessInformation1{}.Validate())
	assert.Nil(t, GeneralBusinessOrError7Choice{}.Validate())
	assert.Nil(t, GeneralBusinessOrError8Choice{}.Validate())
	assert.NotNil(t, GeneralBusinessReport6{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.Nil(t, InformationQualifierType1{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.NotNil(t, ReturnGeneralBusinessInformationV06{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriod2Choice{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, ExecutionType1Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ModifyStandingOrderV06{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.Nil(t, StandingOrder7{}.Validate())
	assert.NotNil(t, StandingOrderIdentification4{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, InvestigationRejectionJustification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, RejectInvestigationV06{}.Validate())
	assert.NotNil(t, RequestForDuplicateV06{}.Validate())
	assert.NotNil(t, DuplicateV06{}.Validate())
	assert.Nil(t, ProprietaryData6{}.Validate())
	assert.NotNil(t, ProprietaryData7{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.Nil(t, CurrentAndDefaultReservation4{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.NotNil(t, Reservation3{}.Validate())
	assert.NotNil(t, ReservationIdentification2{}.Validate())
	assert.Nil(t, ReservationOrError8Choice{}.Validate())
	assert.NotNil(t, ReservationOrError9Choice{}.Validate())
	assert.NotNil(t, ReservationReport6{}.Validate())
	assert.NotNil(t, ReservationStatus1Choice{}.Validate())
	assert.NotNil(t, ReservationType1Choice{}.Validate())
	assert.NotNil(t, ReturnReservationV06{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.NotNil(t, AccountNotification16{}.Validate())
	assert.NotNil(t, NotificationItem7{}.Validate())
	assert.NotNil(t, NotificationToReceiveV06{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.Nil(t, RemittanceLocation7{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
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
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GroupHeader77{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.NotNil(t, NotificationToReceiveCancellationAdviceV06{}.Validate())
	assert.NotNil(t, OriginalItem6{}.Validate())
	assert.Nil(t, OriginalItemReference5{}.Validate())
	assert.NotNil(t, OriginalNotification12{}.Validate())
	assert.Nil(t, OriginalNotificationReference10{}.Validate())
	assert.NotNil(t, GroupHeader84{}.Validate())
	assert.NotNil(t, NotificationToReceiveStatusReportV06{}.Validate())
	assert.NotNil(t, OriginalItemAndStatus6{}.Validate())
	assert.NotNil(t, OriginalNotification11{}.Validate())
	assert.Nil(t, OriginalNotificationReference9{}.Validate())
	assert.Nil(t, AmendmentInformationDetails10{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.Nil(t, CancellationStatusReason2{}.Validate())
	assert.NotNil(t, CancellationStatusReason2Choice{}.Validate())
	assert.NotNil(t, Case3{}.Validate())
	assert.NotNil(t, CaseAssignment3{}.Validate())
	assert.NotNil(t, CashAccount24{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ChargeType3Choice{}.Validate())
	assert.Nil(t, Charges3{}.Validate())
	assert.NotNil(t, ChargesRecord1{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification3Choice{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, CorrectiveGroupInformation1{}.Validate())
	assert.Nil(t, CorrectiveInterbankTransaction1{}.Validate())
	assert.Nil(t, CorrectivePaymentInitiation1{}.Validate())
	assert.Nil(t, CorrectiveTransaction1Choice{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, DatePeriodDetails{}.Validate())
	assert.NotNil(t, EquivalentAmount2{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, Frequency21Choice{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment1{}.Validate())
	assert.NotNil(t, GenericIdentification3{}.Validate())
	assert.NotNil(t, InvestigationStatus3Choice{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation10{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.NotNil(t, NumberOfCancellationsPerStatus1{}.Validate())
	assert.NotNil(t, NumberOfTransactionsPerStatus1{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OriginalGroupHeader5{}.Validate())
	assert.NotNil(t, OriginalGroupInformation3{}.Validate())
	assert.NotNil(t, OriginalPaymentInstruction17{}.Validate())
	assert.Nil(t, OriginalTransactionReference22{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, Party12Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PaymentTransaction66{}.Validate())
	assert.Nil(t, PaymentTransaction67{}.Validate())
	assert.Nil(t, PaymentTypeInformation25{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, RemittanceInformation11{}.Validate())
	assert.Nil(t, ResolutionInformation1{}.Validate())
	assert.NotNil(t, ResolutionOfInvestigationV06{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.NotNil(t, SettlementInstruction4{}.Validate())
	assert.Nil(t, StatementResolutionEntry2{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation13{}.Validate())
	assert.Nil(t, TaxAmount1{}.Validate())
	assert.Nil(t, TaxCharges2{}.Validate())
	assert.Nil(t, TaxInformation4{}.Validate())
	assert.Nil(t, TaxPeriod1{}.Validate())
	assert.Nil(t, TaxRecord1{}.Validate())
	assert.NotNil(t, TaxRecordDetails1{}.Validate())
	assert.Nil(t, UnderlyingTransaction14{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalSystemErrorHandling1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
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

	var type8 ExternalProxyAccountType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalSystemEventType1Code
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

	var type12 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type12.Validate())
	type12 = "tes"
	assert.Nil(t, type12.Validate())

	var type13 ExternalDiscountAmountType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalDocumentLineType1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalGarnishmentType1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalPurpose1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 ExternalTaxAmountType1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 Priority1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "NORM"
	assert.Nil(t, type18.Validate())

	var type19 Frequency2Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "YEAR"
	assert.Nil(t, type19.Validate())

	var type20 InvestigationRejection1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "NFND"
	assert.Nil(t, type20.Validate())

	var type21 PreferredContactMethod1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "LETT"
	assert.Nil(t, type21.Validate())

	var type22 ReservationStatus1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "ENAB"
	assert.Nil(t, type22.Validate())

	var type23 ReservationType2Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "CARE"
	assert.Nil(t, type23.Validate())

	var type24 DocumentType3Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "RADM"
	assert.Nil(t, type24.Validate())

	var type25 DocumentType6Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "MSIN"
	assert.Nil(t, type25.Validate())

	var type26 RemittanceLocationMethod2Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "FAXI"
	assert.Nil(t, type26.Validate())

	var type27 TaxRecordPeriod1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "MM01"
	assert.Nil(t, type27.Validate())

	var type28 NotificationStatus3Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "RCBD"
	assert.Nil(t, type28.Validate())

	var type29 SettlementMethod1Code
	assert.NotNil(t, type29.Validate())
	type29 = "CLRG"
	assert.Nil(t, type29.Validate())

	var type30 ExternalServiceLevel1Code
	assert.NotNil(t, type30.Validate())
	type30 = "CLRG"
	assert.Nil(t, type30.Validate())

	var type31 ClearingChannel2Code
	assert.NotNil(t, type31.Validate())
	type31 = "RTGS"
	assert.Nil(t, type31.Validate())

	var type32 SequenceType3Code
	assert.NotNil(t, type32.Validate())
	type32 = "FRST"
	assert.Nil(t, type32.Validate())

	var type33 Priority2Code
	assert.NotNil(t, type33.Validate())
	type33 = "HIGH"
	assert.Nil(t, type33.Validate())

	var type34 CancellationIndividualStatus1Code
	assert.NotNil(t, type34.Validate())
	type34 = "RJCR"
	assert.Nil(t, type34.Validate())

	var type35 PaymentMethod4Code
	assert.NotNil(t, type35.Validate())
	type35 = "CHK"
	assert.Nil(t, type35.Validate())

	var type36 GroupCancellationStatus1Code
	assert.NotNil(t, type36.Validate())
	type36 = "PACR"
	assert.Nil(t, type36.Validate())

	var type37 TransactionIndividualStatus1Code
	assert.NotNil(t, type37.Validate())
	type37 = "ACTC"
	assert.Nil(t, type37.Validate())

	var type38 ExternalMandateSetupReason1Code
	assert.NotNil(t, type38.Validate())
	type38 = "ACTC"
	assert.Nil(t, type38.Validate())

	var type39 ExternalLocalInstrument1Code
	assert.NotNil(t, type39.Validate())
	type39 = "ACTC"
	assert.Nil(t, type39.Validate())

	var type40 PaymentCancellationRejection2Code
	assert.NotNil(t, type40.Validate())
	type40 = "LEGL"
	assert.Nil(t, type40.Validate())

	var type41 ExternalCategoryPurpose1Code
	assert.NotNil(t, type41.Validate())
	type41 = "LEGL"
	assert.Nil(t, type41.Validate())

	var type42 ExternalChargeType1Code
	assert.NotNil(t, type42.Validate())
	type42 = "LEGL"
	assert.Nil(t, type42.Validate())

	var type43 ChargeBearerType1Code
	assert.NotNil(t, type43.Validate())
	type43 = "DEBT"
	assert.Nil(t, type43.Validate())

	var type44 ChargeBearerType1Code
	assert.NotNil(t, type44.Validate())
	type44 = "DEBT"
	assert.Nil(t, type44.Validate())

	var type45 ExternalCashClearingSystem1Code
	assert.NotNil(t, type45.Validate())
	type45 = "DEBT"
	assert.Nil(t, type45.Validate())

	var type46 Frequency6Code
	assert.NotNil(t, type46.Validate())
	type46 = "YEAR"
	assert.Nil(t, type46.Validate())

	var type47 InvestigationExecutionConfirmation3Code
	assert.NotNil(t, type47.Validate())
	type47 = "PDCR"
	assert.Nil(t, type47.Validate())

	var type48 ModificationRejection2Code
	assert.NotNil(t, type48.Validate())
	type48 = "UM01"
	assert.Nil(t, type48.Validate())
}
