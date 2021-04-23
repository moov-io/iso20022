package head_v02

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBusinessApplicationHeaderV02(t *testing.T) {
	sample := BusinessApplicationHeaderV02{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = BusinessApplicationHeaderV02{
		Xmlns:     sample.NameSpace(),
		BizMsgIdr: "BizMsgIdr",
		MsgDefIdr: "MsgDefIdr",
		CreDt:     common.ISODateTime(testTime),
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:head.001.001.02","Fr":{"OrgId":{},"FIId":{"FinInstnId":{}}},"To":{"OrgId":{},"FIId":{"FinInstnId":{}}},"BizMsgIdr":"BizMsgIdr","MsgDefIdr":"MsgDefIdr","CreDt":"2014-11-12T11:45:26.371"}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<AppHdr xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.02"><Fr><OrgId></OrgId><FIId><FinInstnId></FinInstnId></FIId></Fr><To><OrgId></OrgId><FIId><FinInstnId></FinInstnId></FIId></To><BizMsgIdr>BizMsgIdr</BizMsgIdr><MsgDefIdr>MsgDefIdr</MsgDefIdr><CreDt>2014-11-12T11:45:26.371</CreDt></AppHdr>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, BusinessApplicationHeader5{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, ImplementationSpecification1{}.Validate())
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

	var type5 PreferredContactMethod1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.NotNil(t, type5.Validate())
	type5 = "LETT"
	assert.Nil(t, type5.Validate())
}
