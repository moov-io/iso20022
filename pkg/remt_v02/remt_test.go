// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package remt_v02

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/stretchr/testify/assert"
)

const (
	testTimeString = "2014-11-12T11:45:26.371Z"
)

func TestDocumentRemt00200102(t *testing.T) {
	sample := DocumentRemt00200102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentRemt00200102{
		RmtLctnAdvc: RemittanceLocationAdviceV02{
			GrpHdr: GroupHeader79{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"RmtLctnAdvc":{"GrpHdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:remt.002.001.02" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><RmtLctnAdvc xmlns="urn:iso:std:iso:20022:tech:xsd:remt.002.001.02"><GrpHdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></GrpHdr></RmtLctnAdvc></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.NotNil(t, Authorisation1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader79{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, RemittanceLocation5{}.Validate())
	assert.NotNil(t, RemittanceLocationAdviceV02{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, TransactionReferences5{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalPersonIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type15 PreferredContactMethod1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "CELL"
	assert.Nil(t, type15.Validate())

	var type16 RemittanceLocationMethod2Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.NotNil(t, type16.Validate())
	type16 = "FAXI"
	assert.Nil(t, type16.Validate())
}
