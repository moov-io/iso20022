// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

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

func TestDocumentAcmt03600101(t *testing.T) {
	sample := DocumentAcmt03600101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt03600101{
		Xmlns: sample.NameSpace(),
		AcctSwtchTermntnSwtch: AccountSwitchTerminationSwitchV01{
			MsgId: MessageIdentification1{
				Id:      "ID",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchDt:     common.ISODate(testTime),
				SwtchTp:     "PART",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01","AcctSwtchTermntnSwtch":{"MsgId":{"Id":"ID","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchDt":"2014-11-12","SwtchTp":"PART"}}}`, string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.036.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchTermntnSwtch><MsgId><Id>ID</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchDt>2014-11-12</SwtchDt><SwtchTp>PART</SwtchTp></AcctSwtchDtls></AcctSwtchTermntnSwtch></Document>`, string(buf))
}

func TestTypes(t *testing.T) {
	var type1 BalanceTransferWindow1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "DAYH"
	assert.Nil(t, type1.Validate())

	var type2 SwitchStatus1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "TMTN"
	assert.Nil(t, type2.Validate())

	var type3 SwitchType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.NotNil(t, type3.Validate())
	type3 = "PART"
	assert.Nil(t, type3.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountSwitchDetails1{}.Validate())
	assert.NotNil(t, AccountSwitchTerminationSwitchV01{}.Validate())
	assert.NotNil(t, MessageIdentification1{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
}
