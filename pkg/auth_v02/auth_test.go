// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, BenchmarkCurveName4Choice{}.Validate())
	assert.Nil(t, BinaryFile1{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashCollateral5{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, ContractBalance1{}.Validate())
	assert.NotNil(t, ContractBalanceType1Choice{}.Validate())
	assert.NotNil(t, ContractCollateral1{}.Validate())
	assert.NotNil(t, ContractRegistration3{}.Validate())
	assert.NotNil(t, ContractRegistration4{}.Validate())
	assert.NotNil(t, ContractRegistrationRequestV02{}.Validate())
	assert.NotNil(t, CurrencyControlHeader4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, DocumentGeneralInformation3{}.Validate())
	assert.NotNil(t, DocumentIdentification22{}.Validate())
	assert.Nil(t, ExchangeRate1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, FloatingInterestRate4{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, InterestPaymentDateRange1{}.Validate())
	assert.NotNil(t, InterestPaymentDateRange2{}.Validate())
	assert.Nil(t, InterestPaymentSchedule1Choice{}.Validate())
	assert.NotNil(t, InterestRate2Choice{}.Validate())
	assert.NotNil(t, InterestRateContractTerm1{}.Validate())
	assert.Nil(t, LegalOrganisation2{}.Validate())
	assert.NotNil(t, LoanContract2{}.Validate())
	assert.NotNil(t, LoanContractTranche1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PaymentDateRange1{}.Validate())
	assert.NotNil(t, PaymentDateRange2{}.Validate())
	assert.Nil(t, PaymentSchedule1Choice{}.Validate())
	assert.NotNil(t, PaymentScheduleType1Choice{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.Nil(t, ShipmentDateRange1{}.Validate())
	assert.Nil(t, ShipmentDateRange2{}.Validate())
	assert.Nil(t, ShipmentSchedule2Choice{}.Validate())
	assert.Nil(t, SignatureEnvelopeReference{}.Validate())
	assert.NotNil(t, SpecialCondition1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, SyndicatedLoan2{}.Validate())
	assert.NotNil(t, TaxExemptionReasonFormat1Choice{}.Validate())
	assert.Nil(t, TaxParty4{}.Validate())
	assert.NotNil(t, TradeContract2{}.Validate())
	assert.Nil(t, TradeParty5{}.Validate())
	assert.NotNil(t, UnderlyingContract2Choice{}.Validate())
	assert.NotNil(t, ContractClosureReason1Choice{}.Validate())
	assert.NotNil(t, ContractRegistrationConfirmationV02{}.Validate())
	assert.NotNil(t, CurrencyControlHeader6{}.Validate())
	assert.Nil(t, DocumentIdentification28{}.Validate())
	assert.NotNil(t, DocumentIdentification29{}.Validate())
	assert.NotNil(t, RegisteredContract7{}.Validate())
	assert.Nil(t, RegisteredContractAmendment1{}.Validate())
	assert.NotNil(t, RegisteredContractCommunication1{}.Validate())
	assert.NotNil(t, RegisteredContractJournal2{}.Validate())
	assert.NotNil(t, ContractRegistrationClosureRequestV02{}.Validate())
	assert.NotNil(t, RegisteredContract6{}.Validate())
	assert.NotNil(t, ContractRegistrationAmendmentRequestV02{}.Validate())
	assert.NotNil(t, RegisteredContract10{}.Validate())
	assert.NotNil(t, RegisteredContract9{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.Nil(t, CertificateIdentification1{}.Validate())
	assert.Nil(t, CertificateReference1{}.Validate())
	assert.NotNil(t, ContractRegistrationReference1Choice{}.Validate())
	assert.NotNil(t, ContractRegistrationStatement2{}.Validate())
	assert.NotNil(t, ContractRegistrationStatementV02{}.Validate())
	assert.Nil(t, DatePeriod3{}.Validate())
	assert.Nil(t, DocumentAmendment1{}.Validate())
	assert.NotNil(t, GenericValidationRuleIdentification1{}.Validate())
	assert.NotNil(t, ProprietaryReference1{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, RegisteredContract8{}.Validate())
	assert.NotNil(t, ReportingPeriod4{}.Validate())
	assert.NotNil(t, ShipmentAttribute1{}.Validate())
	assert.NotNil(t, SupportingDocument2{}.Validate())
	assert.NotNil(t, SupportingDocumentEntry1{}.Validate())
	assert.Nil(t, TimePeriod2{}.Validate())
	assert.NotNil(t, TransactionCertificate2{}.Validate())
	assert.NotNil(t, TransactionCertificate3{}.Validate())
	assert.Nil(t, TransactionCertificateContract1{}.Validate())
	assert.NotNil(t, TransactionCertificateRecord1{}.Validate())
	assert.NotNil(t, ValidationRuleSchemeName1Choice{}.Validate())
	assert.Nil(t, ContractRegistrationStatementCriteria1{}.Validate())
	assert.NotNil(t, ContractRegistrationStatementRequest2{}.Validate())
	assert.NotNil(t, ContractRegistrationStatementRequestV02{}.Validate())
	assert.NotNil(t, CurrencyControlHeader5{}.Validate())
	assert.Nil(t, Party40Choice{}.Validate())
	assert.NotNil(t, PaymentRegulatoryInformationNotificationV02{}.Validate())
	assert.NotNil(t, RegulatoryReportingNotification2{}.Validate())
	assert.NotNil(t, CurrencyControlSupportingDocumentDeliveryV02{}.Validate())
	assert.NotNil(t, CurrencyControlRequestOrLetterV02{}.Validate())
	assert.NotNil(t, OriginalMessage4{}.Validate())
	assert.NotNil(t, SupportingDocumentRequestOrLetter2{}.Validate())
	assert.NotNil(t, CurrencyControlGroupStatus2{}.Validate())
	assert.NotNil(t, CurrencyControlPackageStatus2{}.Validate())
	assert.NotNil(t, CurrencyControlRecordStatus2{}.Validate())
	assert.NotNil(t, CurrencyControlStatusAdviceV02{}.Validate())
	assert.NotNil(t, OriginalMessage5{}.Validate())
	assert.Nil(t, Period2{}.Validate())
	assert.Nil(t, Period4Choice{}.Validate())
	assert.NotNil(t, StatusReason6Choice{}.Validate())
	assert.Nil(t, ValidationStatusReason2{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 BenchmarkCurveName2Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "BBSW"
	assert.Nil(t, type1.Validate())

	var type2 CreditDebit3Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "DBIT"
	assert.Nil(t, type2.Validate())

	var type3 DepositType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.NotNil(t, type3.Validate())
	type3 = "CALL"
	assert.Nil(t, type3.Validate())

	var type4 Exact1NumericText
	assert.NotNil(t, type4.Validate())
	type4 = "111"
	assert.Nil(t, type4.Validate())

	var type5 Exact5NumericText
	assert.NotNil(t, type5.Validate())
	type5 = "11111"
	assert.Nil(t, type5.Validate())

	var type6 ExchangeRateType1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.NotNil(t, type6.Validate())
	type6 = "AGRD"
	assert.Nil(t, type6.Validate())

	var type7 ExternalAccountIdentification1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalContractBalanceType1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalDocumentType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type14 ExternalPersonIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ISINOct2015Identifier
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "AA00000000011"
	assert.Nil(t, type15.Validate())

	var type16 PaymentScheduleType1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.NotNil(t, type16.Validate())
	type16 = "ESTM"
	assert.Nil(t, type16.Validate())

	var type17 PreferredContactMethod1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "CELL"
	assert.Nil(t, type17.Validate())

	var type18 Priority2Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "NORM"
	assert.Nil(t, type18.Validate())

	var type19 RateBasis1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "YEAR"
	assert.Nil(t, type19.Validate())

	var type20 CommunicationMethod4Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "SWMX"
	assert.Nil(t, type20.Validate())

	var type21 TaxExemptReason1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "PRYP"
	assert.Nil(t, type21.Validate())

	var type22 ExternalContractClosureReason1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.Nil(t, type22.Validate())

	var type23 ExternalCashAccountType1Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.Nil(t, type23.Validate())

	var type24 ExternalProxyAccountType1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.Nil(t, type24.Validate())

	var type25 ExternalShipmentCondition1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.Nil(t, type25.Validate())

	var type26 ExternalValidationRuleIdentification1Code
	assert.NotNil(t, type26.Validate())
	type26 = "test"
	assert.Nil(t, type26.Validate())

	var type27 QueryType3Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.NotNil(t, type27.Validate())
	type27 = "MODF"
	assert.Nil(t, type27.Validate())

	var type28 QueryType3Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "MODF"
	assert.Nil(t, type28.Validate())

	var type29 SupportDocumentType1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "SUPP"
	assert.Nil(t, type29.Validate())

	var type30 ExternalStatusReason1Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.Nil(t, type30.Validate())

	var type31 StatisticalReportingStatus1Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "CRPT"
	assert.Nil(t, type31.Validate())
}
