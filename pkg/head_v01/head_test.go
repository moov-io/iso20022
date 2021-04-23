// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package head_v01

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBusinessApplicationHeaderV01(t *testing.T) {
	sample := BusinessApplicationHeaderV01{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = BusinessApplicationHeaderV01{
		Xmlns:     sample.NameSpace(),
		BizMsgIdr: "BizMsgIdr",
		MsgDefIdr: "MsgDefIdr",
		CreDt:     common.ISONormalisedDateTime(testTime),
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"XMLName":"","Xmlns":"urn:iso:std:iso:20022:tech:xsd:head.001.001.01","Fr":{"OrgId":{},"FIId":{"FinInstnId":{},"BrnchId":{}}},"To":{"OrgId":{},"FIId":{"FinInstnId":{},"BrnchId":{}}},"BizMsgIdr":"BizMsgIdr","MsgDefIdr":"MsgDefIdr","CreDt":"2014-11-12T11:45:26.371"}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<AppHdr xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.01"><Fr><OrgId></OrgId><FIId><FinInstnId></FinInstnId><BrnchId></BrnchId></FIId></Fr><To><OrgId></OrgId><FIId><FinInstnId></FinInstnId><BrnchId></BrnchId></FIId></To><BizMsgIdr>BizMsgIdr</BizMsgIdr><MsgDefIdr>MsgDefIdr</MsgDefIdr><CreDt>2014-11-12T11:45:26.371</CreDt></AppHdr>`)
}

func TestNestedTypes(t *testing.T) {
	assert.Nil(t, BranchAndFinancialInstitutionIdentification5{}.Validate())
	assert.Nil(t, BranchData2{}.Validate())
	assert.NotNil(t, BusinessApplicationHeader1{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, ContactDetails2{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification8{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification7{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, Party10Choice{}.Validate())
	assert.Nil(t, Party9Choice{}.Validate())
	assert.Nil(t, PartyIdentification42{}.Validate())
	assert.Nil(t, PersonIdentification5{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress6{}.Validate())
	assert.Nil(t, SignatureEnvelope{}.Validate())
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
}
