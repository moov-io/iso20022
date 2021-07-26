// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalPersonIdentification1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalVerificationReason1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 BalanceTransferWindow1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.NotNil(t, type7.Validate())
	type7 = "EARL"
	assert.Nil(t, type7.Validate())

	var type8 PreferredContactMethod1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "CELL"
	assert.Nil(t, type8.Validate())

	var type9 SwitchStatus1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.NotNil(t, type9.Validate())
	type9 = "TMTN"
	assert.Nil(t, type9.Validate())

	var type10 SwitchType1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "PART"
	assert.Nil(t, type10.Validate())

	var type11 ExternalCashAccountType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalProxyAccountType1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, IdentificationAssignment2{}.Validate())
	assert.Nil(t, IdentificationInformation2{}.Validate())
	assert.NotNil(t, IdentificationModification2{}.Validate())
	assert.NotNil(t, IdentificationModificationAdviceV02{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, OriginalTransactionReference18{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, Party12Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.NotNil(t, PaymentIdentification4{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, IdentificationVerification2{}.Validate())
	assert.NotNil(t, IdentificationVerificationRequestV02{}.Validate())
	assert.NotNil(t, IdentificationVerificationReportV02{}.Validate())
	assert.Nil(t, MessageIdentification5{}.Validate())
	assert.NotNil(t, VerificationReason1Choice{}.Validate())
	assert.NotNil(t, VerificationReport2{}.Validate())
	assert.NotNil(t, AccountSwitchDetails1{}.Validate())
	assert.NotNil(t, AccountSwitchRequestRedirectionV02{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashAccount39{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, MessageIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, AccountSwitchNotifyAccountSwitchCompleteV02{}.Validate())
	assert.NotNil(t, AccountSwitchPaymentResponseV02{}.Validate())
	assert.NotNil(t, AccountSwitchTechnicalRejectionV02{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
}
