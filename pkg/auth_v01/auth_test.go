// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, InformationRequestOpeningV01{}.Validate())
	assert.NotNil(t, InformationRequestResponseV01{}.Validate())
	assert.NotNil(t, InformationRequestStatusChangeNotificationV01{}.Validate())
	assert.NotNil(t, AccountAndParties1{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, AuthorityInvestigation2{}.Validate())
	assert.NotNil(t, AuthorityRequestType1{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, CashAccount25{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.Nil(t, CustomerIdentification1{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth{}.Validate())
	assert.Nil(t, DateOrDateTimePeriodChoice{}.Validate())
	assert.Nil(t, DatePeriodDetails{}.Validate())
	assert.Nil(t, DateTimePeriodDetails{}.Validate())
	assert.Nil(t, DueDate1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, InvestigatedParties1Choice{}.Validate())
	assert.NotNil(t, LegalMandate1{}.Validate())
	assert.Nil(t, OrganisationIdentification8{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, Party11Choice{}.Validate())
	assert.Nil(t, PartyIdentification43{}.Validate())
	assert.NotNil(t, PaymentInstrumentType1{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.NotNil(t, RequestType1{}.Validate())
	assert.NotNil(t, SearchCriteria1Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, InvestigationResult1Choice{}.Validate())
	assert.NotNil(t, ReturnIndicator1{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 Min8Max28NumericText
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "1111111111"
	assert.Nil(t, type1.Validate())

	var type2 InvestigationStatus1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "NOAP"
	assert.Nil(t, type2.Validate())

	var type3 TransactionRequestType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.NotNil(t, type3.Validate())
	type3 = "OREC"
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

	var type8 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalPersonIdentification1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 StatusResponse1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "PART"
	assert.Nil(t, type10.Validate())

	var type11 InvestigatedParties1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.NotNil(t, type11.Validate())
	type11 = "OWNE"
	assert.Nil(t, type11.Validate())
}
