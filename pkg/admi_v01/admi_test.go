// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, Admi00200101{}.Validate())
	assert.NotNil(t, MessageReference{}.Validate())
	assert.NotNil(t, RejectionReason2{}.Validate())
	assert.NotNil(t, Admi00400101{}.Validate())
	assert.NotNil(t, Event1{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountIdentificationSearchCriteria2Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, BalanceType11Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashBalance12{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.Nil(t, DateAndDateTimeSearch4Choice{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DatePeriodSearch1Choice{}.Validate())
	assert.Nil(t, DateTimePeriod1{}.Validate())
	assert.Nil(t, DateTimePeriod1Choice{}.Validate())
	assert.NotNil(t, EventType1Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericIdentification36{}.Validate())
	assert.NotNil(t, MessageHeader7{}.Validate())
	assert.NotNil(t, NameAndAddress5{}.Validate())
	assert.NotNil(t, OriginalBusinessQuery1{}.Validate())
	assert.NotNil(t, PartyIdentification120Choice{}.Validate())
	assert.NotNil(t, PartyIdentification136{}.Validate())
	assert.NotNil(t, PostalAddress1{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ReportQueryCriteria2{}.Validate())
	assert.NotNil(t, ReportQueryRequestV01{}.Validate())
	assert.NotNil(t, ReportQuerySearchCriteria2{}.Validate())
	assert.NotNil(t, RequestType4Choice{}.Validate())
	assert.NotNil(t, ResendRequestV01{}.Validate())
	assert.NotNil(t, ResendSearchCriteria2{}.Validate())
	assert.NotNil(t, SequenceRange1{}.Validate())
	assert.NotNil(t, SequenceRange1Choice{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, MessageHeader10{}.Validate())
	assert.NotNil(t, MessageReference1{}.Validate())
	assert.NotNil(t, ReceiptAcknowledgementReport2{}.Validate())
	assert.NotNil(t, ReceiptAcknowledgementV01{}.Validate())
	assert.NotNil(t, RequestHandling2{}.Validate())
	assert.NotNil(t, SystemEventAcknowledgementV01{}.Validate())
	assert.NotNil(t, NameAndAddress8{}.Validate())
	assert.NotNil(t, PartyIdentification44{}.Validate())
	assert.Nil(t, PartyIdentification59{}.Validate())
	assert.NotNil(t, PartyIdentification73Choice{}.Validate())
	assert.NotNil(t, ProcessingRequestV01{}.Validate())
	assert.NotNil(t, RequestDetails19{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalEnquiryRequestType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalPaymentControlRequestType1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalSystemBalanceType1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalSystemEventType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 BalanceCounterparty1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.NotNil(t, type8.Validate())
	type8 = "BILA"
	assert.Nil(t, type8.Validate())
}
