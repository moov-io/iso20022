// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v02

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentAdmi00400102(t *testing.T) {
	sample := DocumentAdmi00400102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00400102{
		SysEvtNtfctn: SystemEventNotificationV02{
			EvtInf: Event2{
				EvtCd: "ABCD",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"SysEvtNtfctn":{"EvtInf":{"EvtCd":"ABCD"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00400102><SysEvtNtfctn><EvtInf><EvtCd>ABCD</EvtCd></EvtInf></SysEvtNtfctn></DocumentAdmi00400102>`,
		string(buf))
}

func TestDocumentAdmi00900102(t *testing.T) {
	sample := DocumentAdmi00900102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi00900102{
		StatcDataReq: StaticDataRequestV02{
			MsgId: "MsgId",
			DataReqDtls: RequestDetails3{
				Tp: "Tp",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"StatcDataReq":{"MsgId":"MsgId","DataReqDtls":{"Tp":"Tp"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi00900102><StatcDataReq><MsgId>MsgId</MsgId><DataReqDtls><Tp>Tp</Tp></DataReqDtls></StatcDataReq></DocumentAdmi00900102>`,
		string(buf))
}

func TestDocumentAdmi01000102(t *testing.T) {
	sample := DocumentAdmi01000102{}
	err := sample.Validate()
	assert.NotNil(t, err)

	sample = DocumentAdmi01000102{
		StatcDataRpt: StaticDataReportV02{
			MsgId: "MsgId",
			RptDtls: RequestDetails5{
				Tp:     "Tp",
				ReqRef: "ReqRef",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`{"XMLName":{"Space":"","Local":""},"StatcDataRpt":{"MsgId":"MsgId","RptDtls":{"Tp":"Tp","ReqRef":"ReqRef"}}}`,
		string(buf))

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t,
		`<DocumentAdmi01000102><StatcDataRpt><MsgId>MsgId</MsgId><RptDtls><Tp>Tp</Tp><ReqRef>ReqRef</ReqRef></RptDtls></StatcDataRpt></DocumentAdmi01000102>`,
		string(buf))
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, Event2{}.Validate())
	assert.NotNil(t, SystemEventNotificationV02{}.Validate())
	assert.NotNil(t, RequestDetails3{}.Validate())
	assert.NotNil(t, StaticDataRequestV02{}.Validate())
	assert.NotNil(t, ReportParameter1{}.Validate())
	assert.NotNil(t, RequestDetails4{}.Validate())
	assert.NotNil(t, RequestDetails5{}.Validate())
	assert.NotNil(t, StaticDataReportV02{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
}
