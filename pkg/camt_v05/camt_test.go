// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.Nil(t, BusinessDayCriteria2{}.Validate())
	assert.NotNil(t, BusinessDayCriteria3Choice{}.Validate())
	assert.Nil(t, BusinessDayQuery2{}.Validate())
	assert.Nil(t, BusinessDayReturnCriteria2{}.Validate())
	assert.Nil(t, BusinessDaySearchCriteria2{}.Validate())
	assert.Nil(t, DateTimePeriod1{}.Validate())
	assert.Nil(t, DateTimePeriod1Choice{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GetBusinessDayInformationV05{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.NotNil(t, MessageHeader9{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, SystemEventType2Choice{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.Nil(t, LongPaymentIdentification2{}.Validate())
	assert.NotNil(t, OriginalMessageAndIssuer1{}.Validate())
	assert.NotNil(t, PaymentIdentification6Choice{}.Validate())
	assert.NotNil(t, PaymentOrigin1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, QueueTransactionIdentification1{}.Validate())
	assert.NotNil(t, Receipt3{}.Validate())
	assert.NotNil(t, ReceiptV05{}.Validate())
	assert.NotNil(t, RequestHandling1{}.Validate())
	assert.NotNil(t, ShortPaymentIdentification2{}.Validate())
	assert.NotNil(t, Case5{}.Validate())
	assert.NotNil(t, CaseAssignment5{}.Validate())
	assert.NotNil(t, CaseForwardingNotification3{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, NotificationOfCaseAssignmentV05{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, ReportHeader5{}.Validate())
	assert.Nil(t, ProprietaryData6{}.Validate())
	assert.NotNil(t, ProprietaryData7{}.Validate())
	assert.NotNil(t, ProprietaryFormatInvestigationV05{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.Nil(t, DebitAuthorisationConfirmation2{}.Validate())
	assert.NotNil(t, DebitAuthorisationResponseV05{}.Validate())
	assert.NotNil(t, CaseStatus2{}.Validate())
	assert.NotNil(t, CaseStatusReportV05{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GetReservationV05{}.Validate())
	assert.NotNil(t, ReservationCriteria3Choice{}.Validate())
	assert.Nil(t, ReservationCriteria4{}.Validate())
	assert.Nil(t, ReservationQuery3{}.Validate())
	assert.Nil(t, ReservationReturnCriteria1{}.Validate())
	assert.Nil(t, ReservationSearchCriteria3{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.NotNil(t, CurrentOrDefaultReservation2Choice{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, ModifyReservationV05{}.Validate())
	assert.NotNil(t, Reservation4{}.Validate())
	assert.NotNil(t, DeleteReservationV05{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.NotNil(t, ReservationIdentification2{}.Validate())
	assert.NotNil(t, ReservationType1Choice{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, LiquidityCreditTransfer2{}.Validate())
	assert.NotNil(t, LiquidityCreditTransferV05{}.Validate())
	assert.Nil(t, PaymentIdentification8{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, LiquidityDebitTransfer2{}.Validate())
	assert.NotNil(t, LiquidityDebitTransferV05{}.Validate())
	assert.NotNil(t, AccountReportingRequestV05{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, BalanceSubType1Choice{}.Validate())
	assert.NotNil(t, BalanceType10Choice{}.Validate())
	assert.NotNil(t, BalanceType13{}.Validate())
	assert.Nil(t, DatePeriodDetails1{}.Validate())
	assert.NotNil(t, EntryStatus1Choice{}.Validate())
	assert.NotNil(t, GroupHeader77{}.Validate())
	assert.NotNil(t, Limit2{}.Validate())
	assert.NotNil(t, ReportingPeriod2{}.Validate())
	assert.NotNil(t, ReportingRequest5{}.Validate())
	assert.NotNil(t, SequenceRange1{}.Validate())
	assert.NotNil(t, SequenceRange1Choice{}.Validate())
	assert.Nil(t, TimePeriodDetails1{}.Validate())
	assert.NotNil(t, TransactionType2{}.Validate())
	assert.Nil(t, AmendmentInformationDetails10{}.Validate())
	assert.NotNil(t, AmountType4Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, CancellationReason14Choice{}.Validate())
	assert.NotNil(t, Case3{}.Validate())
	assert.NotNil(t, CaseAssignment3{}.Validate())
	assert.NotNil(t, CashAccount24{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification3Choice{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.Nil(t, ControlData1{}.Validate())
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
	assert.NotNil(t, FIToFIPaymentCancellationRequestV05{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, Frequency21Choice{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment1{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation10{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OriginalGroupHeader4{}.Validate())
	assert.NotNil(t, OriginalGroupInformation3{}.Validate())
	assert.Nil(t, OriginalTransactionReference22{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, Party12Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.Nil(t, PaymentCancellationReason2{}.Validate())
	assert.Nil(t, PaymentTransaction62{}.Validate())
	assert.Nil(t, PaymentTypeInformation25{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation11{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.NotNil(t, SettlementInstruction4{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation13{}.Validate())
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
	assert.Nil(t, UnderlyingTransaction13{}.Validate())
	assert.Nil(t, DateAndDateTimeChoice{}.Validate())
	assert.Nil(t, MissingOrIncorrectInformation3{}.Validate())
	assert.NotNil(t, UnableToApplyIncorrect1{}.Validate())
	assert.Nil(t, UnableToApplyJustification3Choice{}.Validate())
	assert.NotNil(t, UnableToApplyMissing1{}.Validate())
	assert.NotNil(t, UnableToApplyV05{}.Validate())
	assert.NotNil(t, UnderlyingGroupInformation1{}.Validate())
	assert.NotNil(t, UnderlyingPaymentInstruction3{}.Validate())
	assert.NotNil(t, UnderlyingPaymentTransaction2{}.Validate())
	assert.Nil(t, UnderlyingStatementEntry1{}.Validate())
	assert.NotNil(t, UnderlyingTransaction3Choice{}.Validate())
	assert.NotNil(t, AdditionalPaymentInformationV05{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent1{}.Validate())
	assert.Nil(t, InstructionForNextAgent1{}.Validate())
	assert.Nil(t, PaymentComplementaryInformation4{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation6{}.Validate())
	assert.Nil(t, RemittanceInformation10{}.Validate())
	assert.NotNil(t, SettlementInstruction1{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation12{}.Validate())
	assert.NotNil(t, UnderlyingPaymentInstruction2{}.Validate())
	assert.NotNil(t, UnderlyingTransaction2Choice{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type2.Validate())
	type2 = "tes"
	assert.Nil(t, type2.Validate())

	var type3 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalPersonIdentification1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalAccountIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalCashAccountType1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalProxyAccountType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalBalanceSubType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalBalanceType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalBalanceType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalEntryStatus1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 EntryTypeIdentifier
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "B00DUM"
	assert.Nil(t, type15.Validate())

	var type17 QueryType2Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "DELD"
	assert.Nil(t, type17.Validate())

	var type18 SystemEventType2Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "LVCO"
	assert.Nil(t, type18.Validate())

	var type19 PaymentInstrument1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "CAN"
	assert.Nil(t, type19.Validate())

	var type20 CaseForwardingNotification3Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "MINE"
	assert.Nil(t, type20.Validate())

	var type21 PreferredContactMethod1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "CELL"
	assert.Nil(t, type21.Validate())

	var type22 CaseStatus2Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "ODUE"
	assert.Nil(t, type22.Validate())

	var type23 ReservationType1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "THRE"
	assert.Nil(t, type23.Validate())

	var type24 ReservationType2Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "BLKD"
	assert.Nil(t, type24.Validate())

	var type25 FloorLimitType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "BOTH"
	assert.Nil(t, type25.Validate())

	var type26 QueryType3Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.NotNil(t, type26.Validate())
	type26 = "MODF"
	assert.Nil(t, type26.Validate())

	var type27 CancellationReason5Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "DUPL"
	assert.Nil(t, type27.Validate())

	var type28 DocumentType3Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "RADM"
	assert.Nil(t, type28.Validate())

	var type29 AddressType2Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "ADDR"
	assert.Nil(t, type29.Validate())

	var type30 DocumentType6Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "MSIN"
	assert.Nil(t, type30.Validate())

	var type31 ExternalServiceLevel1Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.Nil(t, type31.Validate())

	var type32 SettlementMethod1Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "COVE"
	assert.Nil(t, type32.Validate())

	var type33 ExternalTaxAmountType1Code
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.Nil(t, type33.Validate())

	var type34 TaxRecordPeriod1Code
	assert.NotNil(t, type34.Validate())
	type34 = "test"
	assert.NotNil(t, type34.Validate())
	type34 = "MM01"
	assert.Nil(t, type34.Validate())

	var type35 UnableToApplyIncorrectInformation4Code
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "IN01"
	assert.Nil(t, type35.Validate())

	var type36 UnableToApplyMissingInformation3Code
	assert.NotNil(t, type36.Validate())
	type36 = "test"
	assert.NotNil(t, type36.Validate())
	type36 = "MS01"
	assert.Nil(t, type36.Validate())

	var type37 ChargeBearerType1Code
	assert.NotNil(t, type37.Validate())
	type37 = "test"
	assert.NotNil(t, type37.Validate())
	type37 = "DEBT"
	assert.Nil(t, type37.Validate())

	var type38 ExternalPurpose1Code
	assert.NotNil(t, type38.Validate())
	type38 = "test"
	assert.Nil(t, type38.Validate())

	var type39 Instruction3Code
	assert.NotNil(t, type39.Validate())
	type39 = "test"
	assert.NotNil(t, type39.Validate())
	type39 = "CHQB"
	assert.Nil(t, type39.Validate())

	var type40 Instruction4Code
	assert.NotNil(t, type40.Validate())
	type40 = "test"
	assert.NotNil(t, type40.Validate())
	type40 = "PHOA"
	assert.Nil(t, type40.Validate())

	var type41 PaymentMethod4Code
	assert.NotNil(t, type41.Validate())
	type41 = "test"
	assert.NotNil(t, type41.Validate())
	type41 = "CHK"
	assert.Nil(t, type41.Validate())

	var type42 Priority2Code
	assert.NotNil(t, type42.Validate())
	type42 = "test"
	assert.NotNil(t, type42.Validate())
	type42 = "HIGH"
	assert.Nil(t, type42.Validate())

	var type43 ClearingChannel2Code
	assert.NotNil(t, type43.Validate())
	type43 = "test"
	assert.NotNil(t, type43.Validate())
	type43 = "RTGS"
	assert.Nil(t, type43.Validate())

	var type44 SequenceType3Code
	assert.NotNil(t, type44.Validate())
	type44 = "test"
	assert.NotNil(t, type44.Validate())
	type44 = "FRST"
	assert.Nil(t, type44.Validate())

	var type45 CreditDebitCode
	assert.NotNil(t, type45.Validate())
	type45 = "test"
	assert.NotNil(t, type45.Validate())
	type45 = "CRDT"
	assert.Nil(t, type45.Validate())

	var type46 ExternalCategoryPurpose1Code
	assert.NotNil(t, type46.Validate())
	type46 = "test"
	assert.Nil(t, type46.Validate())

	var type47 ExternalCashClearingSystem1Code
	assert.NotNil(t, type47.Validate())
	type47 = "test"
	assert.Nil(t, type47.Validate())

	var type48 ExternalDiscountAmountType1Code
	assert.NotNil(t, type48.Validate())
	type48 = "test"
	assert.Nil(t, type48.Validate())

	var type49 ExternalDocumentLineType1Code
	assert.NotNil(t, type49.Validate())
	type49 = "test"
	assert.Nil(t, type49.Validate())

	var type50 CreditDebitCode
	assert.NotNil(t, type50.Validate())
	type50 = "test"
	assert.NotNil(t, type50.Validate())
	type50 = "CRDT"
	assert.Nil(t, type50.Validate())

	var type51 Frequency6Code
	assert.NotNil(t, type51.Validate())
	type51 = "test"
	assert.NotNil(t, type51.Validate())
	type51 = "YEAR"
	assert.Nil(t, type51.Validate())

	var type52 ExternalGarnishmentType1Code
	assert.NotNil(t, type52.Validate())
	type52 = "test"
	assert.Nil(t, type52.Validate())

	var type53 ExternalLocalInstrument1Code
	assert.NotNil(t, type53.Validate())
	type53 = "test"
	assert.Nil(t, type53.Validate())

	var type54 ExternalMandateSetupReason1Code
	assert.NotNil(t, type54.Validate())
	type54 = "test"
	assert.Nil(t, type54.Validate())
}
