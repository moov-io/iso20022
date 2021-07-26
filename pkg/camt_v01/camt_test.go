// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

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

	var type4 ExternalMarketInfrastructure1Code
	assert.NotNil(t, type4.Validate())
	type4 = "tes"
	assert.Nil(t, type4.Validate())

	var type5 LimitType3Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.NotNil(t, type5.Validate())
	type5 = "MULT"
	assert.Nil(t, type5.Validate())

	var type6 ExternalCashAccountType1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalProxyAccountType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalSystemEventType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 Frequency2Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.NotNil(t, type9.Validate())
	type9 = "OVNG"
	assert.Nil(t, type9.Validate())

	var type10 ReservationType2Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.NotNil(t, type10.Validate())
	type10 = "BLKD"
	assert.Nil(t, type10.Validate())

	var type11 PaymentRole1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.NotNil(t, type11.Validate())
	type11 = "STMG"
	assert.Nil(t, type11.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, Amount2Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, CreateLimitV01{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, LimitIdentification5{}.Validate())
	assert.NotNil(t, LimitStructure4{}.Validate())
	assert.NotNil(t, LimitType1Choice{}.Validate())
	assert.NotNil(t, MarketInfrastructureIdentification1Choice{}.Validate())
	assert.NotNil(t, MessageHeader1{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, SystemIdentification2Choice{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CreateStandingOrderV01{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriod2Choice{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, ExecutionType1Choice{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.Nil(t, StandingOrder7{}.Validate())
	assert.NotNil(t, StandingOrderIdentification4{}.Validate())
	assert.NotNil(t, CreateReservationV01{}.Validate())
	assert.NotNil(t, Reservation4{}.Validate())
	assert.NotNil(t, ReservationIdentification2{}.Validate())
	assert.NotNil(t, ReservationType1Choice{}.Validate())
	assert.NotNil(t, CommunicationAddress8{}.Validate())
	assert.NotNil(t, ContactIdentificationAndAddress1{}.Validate())
	assert.NotNil(t, CreateMemberV01{}.Validate())
	assert.NotNil(t, LongPostalAddress1Choice{}.Validate())
	assert.Nil(t, Member6{}.Validate())
	assert.NotNil(t, MemberIdentification3Choice{}.Validate())
	assert.NotNil(t, StructuredLongPostalAddress1{}.Validate())
}
