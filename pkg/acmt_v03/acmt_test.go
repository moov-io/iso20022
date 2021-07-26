// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.Nil(t, AccountContract2{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountOpeningRequestV03{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, Authorisation2{}.Validate())
	assert.Nil(t, BankTransactionCodeStructure4{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure5{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure6{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.Nil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, Channel2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, CodeOrProprietary1Choice{}.Validate())
	assert.NotNil(t, CommunicationFormat1Choice{}.Validate())
	assert.NotNil(t, CommunicationMethod2Choice{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, ContractDocument1{}.Validate())
	assert.NotNil(t, CustomerAccount4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, FixedAmountOrUnlimited1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification13{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, Group4{}.Validate())
	assert.NotNil(t, MaximumAmountByPeriod1{}.Validate())
	assert.NotNil(t, MessageIdentification1{}.Validate())
	assert.NotNil(t, OperationMandate4{}.Validate())
	assert.NotNil(t, Organisation33{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.NotNil(t, PartyAndAuthorisation4{}.Validate())
	assert.Nil(t, PartyAndCertificate4{}.Validate())
	assert.Nil(t, PartyAndSignature3{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PartyIdentification137{}.Validate())
	assert.NotNil(t, PartyOrGroup2Choice{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProprietaryBankTransactionCodeStructure1{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, References4{}.Validate())
	assert.NotNil(t, Restriction1{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.NotNil(t, StatementFrequencyAndForm1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, AccountOpeningAmendmentRequestV03{}.Validate())
	assert.NotNil(t, References3{}.Validate())
	assert.NotNil(t, AccountOpeningAdditionalInformationRequestV03{}.Validate())
	assert.NotNil(t, AccountRequestAcknowledgementV03{}.Validate())
	assert.NotNil(t, AccountRequestRejectionV03{}.Validate())
	assert.NotNil(t, References6{}.Validate())
	assert.NotNil(t, AccountAdditionalInformationRequestV03{}.Validate())
	assert.NotNil(t, AccountReportRequestV03{}.Validate())
	assert.Nil(t, AccountContract3{}.Validate())
	assert.NotNil(t, AccountReport23{}.Validate())
	assert.NotNil(t, AccountReportV03{}.Validate())
	assert.NotNil(t, CustomerAccount5{}.Validate())
	assert.NotNil(t, AccountExcludedMandateMaintenanceRequestV03{}.Validate())
	assert.NotNil(t, References5{}.Validate())
	assert.NotNil(t, AccountStatusModification1{}.Validate())
	assert.Nil(t, AdditionalInformation5{}.Validate())
	assert.Nil(t, AddressModification2{}.Validate())
	assert.Nil(t, AmountModification1{}.Validate())
	assert.Nil(t, DateModification1{}.Validate())
	assert.NotNil(t, FullLegalNameModification1{}.Validate())
	assert.NotNil(t, NameModification1{}.Validate())
	assert.NotNil(t, NumberModification1{}.Validate())
	assert.NotNil(t, OrganisationModification2{}.Validate())
	assert.Nil(t, PartyModification2{}.Validate())
	assert.NotNil(t, PurposeModification1{}.Validate())
	assert.NotNil(t, RestrictionModification1{}.Validate())
	assert.NotNil(t, StatementFrequencyAndFormModification1{}.Validate())
	assert.NotNil(t, TradingNameModification1{}.Validate())
	assert.NotNil(t, TypeModification1{}.Validate())
	assert.NotNil(t, AccountExcludedMandateMaintenanceAmendmentRequestV03{}.Validate())
	assert.NotNil(t, CustomerAccountModification1{}.Validate())
	assert.NotNil(t, AccountMandateMaintenanceRequestV03{}.Validate())
	assert.NotNil(t, Group3{}.Validate())
	assert.NotNil(t, OperationMandate5{}.Validate())
	assert.Nil(t, Organisation34{}.Validate())
	assert.NotNil(t, PartyAndAuthorisation5{}.Validate())
	assert.Nil(t, PartyAndCertificate5{}.Validate())
	assert.NotNil(t, AccountMandateMaintenanceAmendmentRequestV03{}.Validate())
	assert.NotNil(t, AccountClosingRequestV03{}.Validate())
	assert.NotNil(t, AccountForAction2{}.Validate())
	assert.NotNil(t, AccountForAction1{}.Validate())
	assert.NotNil(t, AccountClosingAmendmentRequestV03{}.Validate())
	assert.Nil(t, AccountContract4{}.Validate())
	assert.NotNil(t, AccountClosingAdditionalInformationRequestV03{}.Validate())
	assert.NotNil(t, AccountSwitchInformationRequestV03{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.Nil(t, CitizenshipInformation1{}.Validate())
	assert.Nil(t, CommunicationAddress3{}.Validate())
	assert.NotNil(t, CountryAndResidentialStatusType1{}.Validate())
	assert.NotNil(t, CreditTransferTransaction41{}.Validate())
	assert.NotNil(t, GenericIdentification44{}.Validate())
	assert.NotNil(t, GenericIdentification47{}.Validate())
	assert.NotNil(t, IndividualPerson36{}.Validate())
	assert.NotNil(t, IndividualPersonNameLong2{}.Validate())
	assert.NotNil(t, NewAccount2{}.Validate())
	assert.NotNil(t, Organisation35{}.Validate())
	assert.NotNil(t, OtherIdentification1Choice{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxInformation8{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
	assert.NotNil(t, TransferInstruction1{}.Validate())
	assert.NotNil(t, AccountSwitchBalanceTransferAcknowledgementV03{}.Validate())
	assert.NotNil(t, AccountSwitchDetails1{}.Validate())
	assert.NotNil(t, AmountAndDirection5{}.Validate())
	assert.Nil(t, BalanceTransfer3{}.Validate())
	assert.NotNil(t, BalanceTransferFundingLimit1{}.Validate())
	assert.NotNil(t, BalanceTransferReference1{}.Validate())
	assert.NotNil(t, CashAccount39{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.Nil(t, Cheque11{}.Validate())
	assert.NotNil(t, ChequeDeliveryMethod1Choice{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.Nil(t, EndPoint1Choice{}.Validate())
	assert.Nil(t, Frequency1{}.Validate())
	assert.NotNil(t, Frequency37Choice{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent3{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.NotNil(t, PaymentIdentification6{}.Validate())
	assert.Nil(t, PaymentTypeInformation26{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RegulatoryAuthority2{}.Validate())
	assert.Nil(t, RegulatoryReporting3{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.Nil(t, RemittanceLocation6{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.NotNil(t, SettlementMethod3Choice{}.Validate())
	assert.Nil(t, StructuredRegulatoryReporting3{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.NotNil(t, AccountSwitchInformationResponseV03{}.Validate())
	assert.NotNil(t, DirectDebitInstructionDetails2{}.Validate())
	assert.NotNil(t, PaymentInstruction36{}.Validate())
	assert.NotNil(t, AccountSwitchCancelExistingPaymentV03{}.Validate())
	assert.NotNil(t, AccountSwitchRequestBalanceTransferV03{}.Validate())
	assert.NotNil(t, AccountSwitchRequestPaymentV03{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 CommunicationMethod2Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "POST"
	assert.Nil(t, type1.Validate())

	var type2 CommunicationMethod3Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "ONLI"
	assert.Nil(t, type2.Validate())

	var type3 ExternalAccountIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalBankTransactionDomain1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalBankTransactionFamily1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalBankTransactionSubFamily1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalCashAccountType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalCommunicationFormat1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalPersonIdentification1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalProxyAccountType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 Frequency7Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.NotNil(t, type14.Validate())
	type14 = "INDA"
	assert.Nil(t, type14.Validate())

	var type15 PreferredContactMethod1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "CELL"
	assert.Nil(t, type15.Validate())

	var type16 Unlimited9Text
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 Modification1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "ADDD"
	assert.Nil(t, type17.Validate())

	var type18 BalanceTransferWindow1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "EARL"
	assert.Nil(t, type18.Validate())

	var type19 BusinessDayConvention1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "PREC"
	assert.Nil(t, type19.Validate())

	var type20 ChargeBearerType1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "SLEV"
	assert.Nil(t, type20.Validate())

	var type21 ChequeDelivery1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "RGFA"
	assert.Nil(t, type21.Validate())

	var type22 ChequeDelivery1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "CRFA"
	assert.Nil(t, type22.Validate())

	var type23 ChequeType2Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "ELDR"
	assert.Nil(t, type23.Validate())

	var type24 DocumentType3Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "SCOR"
	assert.Nil(t, type24.Validate())

	var type25 DocumentType6Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "PUOR"
	assert.Nil(t, type25.Validate())

	var type27 ExternalTaxAmountType1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.Nil(t, type27.Validate())

	var type28 Frequency10Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "QURT"
	assert.Nil(t, type28.Validate())

	var type29 Gender1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MALE"
	assert.Nil(t, type29.Validate())

	var type30 OrganisationLegalStatus1Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "PCLG"
	assert.Nil(t, type30.Validate())

	var type31 PersonIdentificationType5Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "GUNL"
	assert.Nil(t, type31.Validate())

	var type32 Priority2Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "NORM"
	assert.Nil(t, type32.Validate())

	var type33 RegulatoryReportingType1Code
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.NotNil(t, type33.Validate())
	type33 = "BOTH"
	assert.Nil(t, type33.Validate())

	var type34 RemittanceLocationMethod2Code
	assert.NotNil(t, type34.Validate())
	type34 = "test"
	assert.NotNil(t, type34.Validate())
	type34 = "SMSM"
	assert.Nil(t, type34.Validate())

	var type35 ResidentialStatus1Code
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "NRES"
	assert.Nil(t, type35.Validate())

	var type36 SwitchType1Code
	assert.NotNil(t, type36.Validate())
	type36 = "test"
	assert.NotNil(t, type36.Validate())
	type36 = "PART"
	assert.Nil(t, type36.Validate())

	var type37 TaxRateMarker1Code
	assert.NotNil(t, type37.Validate())
	type37 = "test"
	assert.NotNil(t, type37.Validate())
	type37 = "GRSS"
	assert.Nil(t, type37.Validate())

	var type38 TaxRecordPeriod1Code
	assert.NotNil(t, type38.Validate())
	type38 = "test"
	assert.NotNil(t, type38.Validate())
	type38 = "MM05"
	assert.Nil(t, type38.Validate())

	var type39 SwitchStatus1Code
	assert.NotNil(t, type39.Validate())
	type39 = "test"
	assert.NotNil(t, type39.Validate())
	type39 = "COMP"
	assert.Nil(t, type39.Validate())

	var type40 ExternalCategoryPurpose1Code
	assert.NotNil(t, type40.Validate())
	type40 = "test"
	assert.Nil(t, type40.Validate())

	var type41 ExternalCreditorAgentInstruction1Code
	assert.NotNil(t, type41.Validate())
	type41 = "test"
	assert.Nil(t, type41.Validate())

	var type42 ExternalDiscountAmountType1Code
	assert.NotNil(t, type42.Validate())
	type42 = "test"
	assert.Nil(t, type42.Validate())

	var type43 ExternalDocumentLineType1Code
	assert.NotNil(t, type43.Validate())
	type43 = "test"
	assert.Nil(t, type43.Validate())

	var type44 ExternalGarnishmentType1Code
	assert.NotNil(t, type44.Validate())
	type44 = "test"
	assert.Nil(t, type44.Validate())

	var type45 ExternalLocalInstrument1Code
	assert.NotNil(t, type45.Validate())
	type45 = "test"
	assert.Nil(t, type45.Validate())

	var type46 ExternalPurpose1Code
	assert.NotNil(t, type46.Validate())
	type46 = "test"
	assert.Nil(t, type46.Validate())

	var type47 ExternalServiceLevel1Code
	assert.NotNil(t, type47.Validate())
	type47 = "test"
	assert.Nil(t, type47.Validate())

	var type48 PaymentMethod3Code
	assert.NotNil(t, type48.Validate())
	type48 = "test"
	assert.NotNil(t, type48.Validate())
	type48 = "TRA"
	assert.Nil(t, type48.Validate())

	var type49 UseCases1Code
	assert.NotNil(t, type49.Validate())
	type49 = "test"
	assert.NotNil(t, type49.Validate())
	type49 = "VIEW"
	assert.Nil(t, type49.Validate())
}
