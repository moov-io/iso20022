// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

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

func TestDocumentAcmt00700103(t *testing.T) {
	sample := DocumentAcmt00700103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt00700103{
		Xmlns: sample.NameSpace(),
		AcctOpngReq: AccountOpeningRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			Acct: CustomerAccount4{Ccy: "ABC"},
			Org: Organisation33{
				FullLglNm: "FullLglNm",
				CtryOfOpr: "AA",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03","AcctOpngReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"Acct":{"Ccy":"ABC"},"AcctSvcrId":{"FinInstnId":{}},"Org":{"FullLglNm":"FullLglNm","CtryOfOpr":"AA","LglAdr":{},"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctOpngReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><Acct><Ccy>ABC</Ccy></Acct><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><Org><FullLglNm>FullLglNm</FullLglNm><CtryOfOpr>AA</CtryOfOpr><LglAdr></LglAdr><OrgId></OrgId></Org></AcctOpngReq></Document>`)
}

func TestDocumentAcmt00800103(t *testing.T) {
	sample := DocumentAcmt00800103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt00800103{
		Xmlns: sample.NameSpace(),
		AcctOpngAmdmntReq: AccountOpeningAmendmentRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			Acct: CustomerAccount4{Ccy: "ABC"},
			Org: Organisation33{
				FullLglNm: "FullLglNm",
				CtryOfOpr: "AA",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03","AcctOpngAmdmntReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"Acct":{"Ccy":"ABC"},"AcctSvcrId":{"FinInstnId":{}},"Org":{"FullLglNm":"FullLglNm","CtryOfOpr":"AA","LglAdr":{},"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.008.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctOpngAmdmntReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><Acct><Ccy>ABC</Ccy></Acct><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><Org><FullLglNm>FullLglNm</FullLglNm><CtryOfOpr>AA</CtryOfOpr><LglAdr></LglAdr><OrgId></OrgId></Org></AcctOpngAmdmntReq></Document>`)
}

func TestDocumentAcmt009001033(t *testing.T) {
	sample := DocumentAcmt00900103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt00900103{
		Xmlns: sample.NameSpace(),
		AcctOpngAddtlInfReq: AccountOpeningAdditionalInformationRequestV03{
			Refs: References3{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				ReqToBeCmpltdId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			Acct: CustomerAccount4{Ccy: "ABC"},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.009.001.03","AcctOpngAddtlInfReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"ReqToBeCmpltdId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"OrgId":{},"Acct":{"Ccy":"ABC"},"AcctSvcrId":{"FinInstnId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.009.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctOpngAddtlInfReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><ReqToBeCmpltdId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></ReqToBeCmpltdId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><OrgId></OrgId><Acct><Ccy>ABC</Ccy></Acct><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId></AcctOpngAddtlInfReq></Document>`)
}

func TestDocumentAcmt01000103(t *testing.T) {
	sample := DocumentAcmt01000103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01000103{
		Xmlns: sample.NameSpace(),
		AcctReqAck: AccountRequestAcknowledgementV03{
			Refs: References5{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				ReqTp: "VIEW",
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.010.001.03","AcctReqAck":{"Refs":{"ReqTp":"VIEW","MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"OrgId":{},"AcctSvcrId":{"FinInstnId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.010.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctReqAck><Refs><ReqTp>VIEW</ReqTp><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><OrgId></OrgId><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId></AcctReqAck></Document>`)
}

func TestDocumentAcmt01100103(t *testing.T) {
	sample := DocumentAcmt01100103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01100103{
		Xmlns: sample.NameSpace(),
		AcctReqRjctn: AccountRequestRejectionV03{
			Refs: References6{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				RjctdReqTp: "VIEW",
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				RjctdReqId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.011.001.03","AcctReqRjctn":{"Refs":{"RjctdReqTp":"VIEW","RjctdReqId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctSvcrId":{"FinInstnId":{}},"OrgId":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.011.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctReqRjctn><Refs><RjctdReqTp>VIEW</RjctdReqTp><RjctdReqId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></RjctdReqId><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><OrgId></OrgId></AcctReqRjctn></Document>`)
}

func TestDocumentAcmt01200103(t *testing.T) {
	sample := DocumentAcmt01200103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01200103{
		Xmlns: sample.NameSpace(),
		AcctAddtlInfReq: AccountAdditionalInformationRequestV03{
			Refs: References3{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				ReqToBeCmpltdId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.012.001.03","AcctAddtlInfReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"ReqToBeCmpltdId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"OrgId":{},"AcctSvcrId":{"FinInstnId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.012.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctAddtlInfReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><ReqToBeCmpltdId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></ReqToBeCmpltdId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><OrgId></OrgId><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId></AcctAddtlInfReq></Document>`)
}

func TestDocumentAcmt01300103(t *testing.T) {
	sample := DocumentAcmt01300103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01300103{
		Xmlns: sample.NameSpace(),
		AcctRptReq: AccountReportRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03","AcctRptReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctSvcrId":{"FinInstnId":{}},"OrgId":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.013.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctRptReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><OrgId></OrgId></AcctRptReq></Document>`)
}

func TestDocumentAcmt01400103(t *testing.T) {
	sample := DocumentAcmt01400103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01400103{
		Xmlns: sample.NameSpace(),
		AcctRpt: AccountReportV03{
			Refs: References5{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				ReqTp: "VIEW",
			},
			Org: Organisation33{
				FullLglNm: "FullLglNm",
				CtryOfOpr: "US",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.03","AcctRpt":{"Refs":{"ReqTp":"VIEW","MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctSvcrId":{"FinInstnId":{}},"Org":{"FullLglNm":"FullLglNm","CtryOfOpr":"US","LglAdr":{},"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.014.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctRpt><Refs><ReqTp>VIEW</ReqTp><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><Org><FullLglNm>FullLglNm</FullLglNm><CtryOfOpr>US</CtryOfOpr><LglAdr></LglAdr><OrgId></OrgId></Org></AcctRpt></Document>`)
}

func TestDocumentAcmt01500103(t *testing.T) {
	sample := DocumentAcmt01500103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01500103{
		Xmlns: sample.NameSpace(),
		AcctExcldMndtMntncReq: AccountExcludedMandateMaintenanceRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			Acct: CustomerAccountModification1{Ccy: "ABC"},
			Org: OrganisationModification2{
				FullLglNm: FullLegalNameModification1{FullLglNm: "FullLglNm"},
				CtryOfOpr: "US",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.03","AcctExcldMndtMntncReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"Acct":{"Ccy":"ABC"},"AcctSvcrId":{"FinInstnId":{}},"Org":{"FullLglNm":{"FullLglNm":"FullLglNm"},"CtryOfOpr":"US","LglAdr":{"Adr":{}},"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.015.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctExcldMndtMntncReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><Acct><Ccy>ABC</Ccy></Acct><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><Org><FullLglNm><FullLglNm>FullLglNm</FullLglNm></FullLglNm><CtryOfOpr>US</CtryOfOpr><LglAdr><Adr></Adr></LglAdr><OrgId></OrgId></Org></AcctExcldMndtMntncReq></Document>`)
}

func TestDocumentAcmt01600103(t *testing.T) {
	sample := DocumentAcmt01600103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01600103{
		Xmlns: sample.NameSpace(),
		AcctExcldMndtMntncAmdmntReq: AccountExcludedMandateMaintenanceAmendmentRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			Acct: CustomerAccountModification1{Ccy: "ABC"},
			Org: OrganisationModification2{
				FullLglNm: FullLegalNameModification1{FullLglNm: "FullLglNm"},
				CtryOfOpr: "US",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.03","AcctExcldMndtMntncAmdmntReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"Acct":{"Ccy":"ABC"},"AcctSvcrId":{"FinInstnId":{}},"Org":{"FullLglNm":{"FullLglNm":"FullLglNm"},"CtryOfOpr":"US","LglAdr":{"Adr":{}},"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.016.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctExcldMndtMntncAmdmntReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><Acct><Ccy>ABC</Ccy></Acct><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><Org><FullLglNm><FullLglNm>FullLglNm</FullLglNm></FullLglNm><CtryOfOpr>US</CtryOfOpr><LglAdr><Adr></Adr></LglAdr><OrgId></OrgId></Org></AcctExcldMndtMntncAmdmntReq></Document>`)
}

func TestDocumentAcmt01700103(t *testing.T) {
	sample := DocumentAcmt01700103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01700103{
		Xmlns: sample.NameSpace(),
		AcctMndtMntncReq: AccountMandateMaintenanceRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.017.001.03","AcctMndtMntncReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctSvcrId":{"FinInstnId":{}},"OrgId":{"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.017.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctMndtMntncReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><OrgId><OrgId></OrgId></OrgId></AcctMndtMntncReq></Document>`)
}

func TestDocumentAcmt01800103(t *testing.T) {
	sample := DocumentAcmt01800103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01800103{
		Xmlns: sample.NameSpace(),
		AcctMndtMntncAmdmntReq: AccountMandateMaintenanceAmendmentRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.018.001.03","AcctMndtMntncAmdmntReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctSvcrId":{"FinInstnId":{}},"OrgId":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.018.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctMndtMntncAmdmntReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><OrgId></OrgId></AcctMndtMntncAmdmntReq></Document>`)
}

func TestDocumentAcmt01900103(t *testing.T) {
	sample := DocumentAcmt01900103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt01900103{
		Xmlns: sample.NameSpace(),
		AcctClsgReq: AccountClosingRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			AcctId: AccountForAction2{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
				Ccy: "ABC",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.019.001.03","AcctClsgReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctId":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}},"Ccy":"ABC"},"AcctSvcrId":{"FinInstnId":{}},"OrgId":{"OrgId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.019.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctClsgReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctId><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id><Ccy>ABC</Ccy></AcctId><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><OrgId><OrgId></OrgId></OrgId></AcctClsgReq></Document>`)
}

func TestDocumentAcmt02000103(t *testing.T) {
	sample := DocumentAcmt02000103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt02000103{
		Xmlns: sample.NameSpace(),
		AcctClsgAmdmntReq: AccountClosingAmendmentRequestV03{
			Refs: References4{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			AcctId: AccountForAction1{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.020.001.03","AcctClsgAmdmntReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"AcctId":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"AcctSvcrId":{"FinInstnId":{}},"OrgId":{}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.020.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctClsgAmdmntReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><AcctId><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></AcctId><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId><OrgId></OrgId></AcctClsgAmdmntReq></Document>`)
}

func TestDocumentAcmt02100103(t *testing.T) {
	sample := DocumentAcmt02100103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt02100103{
		Xmlns: sample.NameSpace(),
		AcctClsgAddtlInfReq: AccountClosingAdditionalInformationRequestV03{
			Refs: References3{
				MsgId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				ReqToBeCmpltdId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
				PrcId: MessageIdentification1{
					Id:      "Id",
					CreDtTm: common.ISODateTime(testTime),
				},
			},
			AcctId: AccountForAction1{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.021.001.03","AcctClsgAddtlInfReq":{"Refs":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"ReqToBeCmpltdId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"PrcId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"}},"OrgId":{},"AcctId":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"AcctSvcrId":{"FinInstnId":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.021.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctClsgAddtlInfReq><Refs><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><ReqToBeCmpltdId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></ReqToBeCmpltdId><PrcId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></PrcId></Refs><OrgId></OrgId><AcctId><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></AcctId><AcctSvcrId><FinInstnId></FinInstnId></AcctSvcrId></AcctClsgAddtlInfReq></Document>`)
}

func TestDocumentAcmt02700103(t *testing.T) {
	sample := DocumentAcmt02700103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt02700103{
		Xmlns: sample.NameSpace(),
		AcctSwtchInfReq: AccountSwitchInformationRequestV03{
			MsgId: MessageIdentification1{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			NewAcct: NewAccount2{
				Acct: CashAccount39{
					Id: AccountIdentification4Choice{
						IBAN: "AA000130",
						Othr: GenericAccountIdentification1{
							Id: "Id",
						},
					},
				},
			},
			OdAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.027.001.03","AcctSwtchInfReq":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"NewAcct":{"Acct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}},"OdAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.027.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchInfReq><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><NewAcct><Acct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></Acct></NewAcct><OdAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></OdAcct></AcctSwtchInfReq></Document>`)
}

func TestDocumentAcmt02800103(t *testing.T) {
	sample := DocumentAcmt02800103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt02800103{
		Xmlns: sample.NameSpace(),
		AcctSwtchInfRspn: AccountSwitchInformationResponseV03{
			MsgId: MessageIdentification1{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			NewAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
			OdAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.028.001.03","AcctSwtchInfRspn":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"NewAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"OdAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.028.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchInfRspn><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><NewAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></NewAcct><OdAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></OdAcct></AcctSwtchInfRspn></Document>`)
}

func TestDocumentAcmt02900103(t *testing.T) {
	sample := DocumentAcmt02900103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt02900103{
		Xmlns: sample.NameSpace(),
		AcctSwtchCclExstgPmt: AccountSwitchCancelExistingPaymentV03{
			MsgId: MessageIdentification1{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			OdAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.029.001.03","AcctSwtchCclExstgPmt":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"OdAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.029.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchCclExstgPmt><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><OdAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></OdAcct></AcctSwtchCclExstgPmt></Document>`)
}

func TestDocumentAcmt03100103(t *testing.T) {
	sample := DocumentAcmt03100103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt03100103{
		Xmlns: sample.NameSpace(),
		AcctSwtchReqBalTrf: AccountSwitchRequestBalanceTransferV03{
			MsgId: MessageIdentification1{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			NewAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.031.001.03","AcctSwtchReqBalTrf":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"NewAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.031.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchReqBalTrf><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><NewAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></NewAcct></AcctSwtchReqBalTrf></Document>`)
}

func TestDocumentAcmt03200103(t *testing.T) {
	sample := DocumentAcmt03200103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt03200103{
		Xmlns: sample.NameSpace(),
		AcctSwtchBalTrfAck: AccountSwitchBalanceTransferAcknowledgementV03{
			MsgId: MessageIdentification1{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			OdAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
			OdAcctBal: AmountAndDirection5{
				Amt: ActiveCurrencyAndAmount{
					Value: 0.1,
					Ccy:   "ABC",
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.032.001.03","AcctSwtchBalTrfAck":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"OdAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"OdAcctBal":{"Amt":{"Value":0.1,"Ccy":"ABC"}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.032.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchBalTrfAck><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><OdAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></OdAcct><OdAcctBal><Amt Ccy="ABC">0.1</Amt></OdAcctBal></AcctSwtchBalTrfAck></Document>`)
}

func TestDocumentAcmt03400103(t *testing.T) {
	sample := DocumentAcmt03400103{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentAcmt03400103{
		Xmlns: sample.NameSpace(),
		AcctSwtchReqPmt: AccountSwitchRequestPaymentV03{
			MsgId: MessageIdentification1{
				Id:      "Id",
				CreDtTm: common.ISODateTime(testTime),
			},
			AcctSwtchDtls: AccountSwitchDetails1{
				UnqRefNb:    "UnqRefNb",
				RtgUnqRefNb: "RtgUnqRefNb",
				SwtchTp:     "PART",
			},
			OdAcct: CashAccount39{
				Id: AccountIdentification4Choice{
					IBAN: "AA000130",
					Othr: GenericAccountIdentification1{
						Id: "Id",
					},
				},
			},
			CdtInstr: CreditTransferTransaction41{
				PmtId: PaymentIdentification6{EndToEndId: "EndToEndId"},
				Amt: ActiveCurrencyAndAmount{
					Value: 0.1,
					Ccy:   "ABC",
				},
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:acmt.034.001.03","AcctSwtchReqPmt":{"MsgId":{"Id":"Id","CreDtTm":"2014-11-12T11:45:26.371"},"AcctSwtchDtls":{"UnqRefNb":"UnqRefNb","RtgUnqRefNb":"RtgUnqRefNb","SwtchTp":"PART"},"OdAcct":{"Id":{"IBAN":"AA000130","Othr":{"Id":"Id"}}},"CdtInstr":{"PmtId":{"EndToEndId":"EndToEndId"},"Amt":{"Value":0.1,"Ccy":"ABC"},"CdtrAgt":{"FinInstnId":{}},"RmtInf":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.034.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><AcctSwtchReqPmt><MsgId><Id>Id</Id><CreDtTm>2014-11-12T11:45:26.371</CreDtTm></MsgId><AcctSwtchDtls><UnqRefNb>UnqRefNb</UnqRefNb><RtgUnqRefNb>RtgUnqRefNb</RtgUnqRefNb><SwtchTp>PART</SwtchTp></AcctSwtchDtls><OdAcct><Id><IBAN>AA000130</IBAN><Othr><Id>Id</Id></Othr></Id></OdAcct><CdtInstr><PmtId><EndToEndId>EndToEndId</EndToEndId></PmtId><Amt Ccy="ABC">0.1</Amt><CdtrAgt><FinInstnId></FinInstnId></CdtrAgt><RmtInf></RmtInf></CdtInstr></AcctSwtchReqPmt></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.Nil(t, AccountContract2{}.Validate())
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountOpeningRequestV03{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, Authorisation2{}.Validate())
	assert.Nil(t, BankTransactionCodeStructure4{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure5{}.Validate())
	assert.NotNil(t, BankTransactionCodeStructure6{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.Nil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, Channel2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.NotNil(t, CodeOrProprietary1Choice{}.Validate())
	assert.NotNil(t, CommunicationFormat1Choice{}.Validate())
	assert.NotNil(t, CommunicationMethod2Choice{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, ContractDocument1{}.Validate())
	assert.NotNil(t, CustomerAccount4{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, FixedAmountOrUnlimited1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification13{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, Group4{}.Validate())
	assert.NotNil(t, MaximumAmountByPeriod1{}.Validate())
	assert.NotNil(t, MessageIdentification1{}.Validate())
	assert.NotNil(t, OperationMandate4{}.Validate())
	assert.NotNil(t, Organisation33{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.NotNil(t, PartyAndAuthorisation4{}.Validate())
	assert.Nil(t, PartyAndCertificate4{}.Validate())
	assert.Nil(t, PartyAndSignature3{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.Nil(t, PartyIdentification137{}.Validate())
	assert.NotNil(t, PartyOrGroup2Choice{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProprietaryBankTransactionCodeStructure1{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, References4{}.Validate())
	assert.NotNil(t, Restriction1{}.Validate())
	assert.Nil(t, SkipPayload{}.Validate())
	assert.NotNil(t, StatementFrequencyAndForm1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.NotNil(t, AccountOpeningAmendmentRequestV03{}.Validate())
	assert.NotNil(t, References3{}.Validate())
	assert.NotNil(t, AccountOpeningAdditionalInformationRequestV03{}.Validate())
	assert.NotNil(t, AccountRequestAcknowledgementV03{}.Validate())
	assert.NotNil(t, AccountRequestRejectionV03{}.Validate())
	assert.NotNil(t, References6{}.Validate())
	assert.NotNil(t, AccountAdditionalInformationRequestV03{}.Validate())
	assert.NotNil(t, AccountReportRequestV03{}.Validate())
	assert.Nil(t, AccountContract3{}.Validate())
	assert.NotNil(t, AccountReport23{}.Validate())
	assert.NotNil(t, AccountReportV03{}.Validate())
	assert.NotNil(t, CustomerAccount5{}.Validate())
	assert.NotNil(t, AccountExcludedMandateMaintenanceRequestV03{}.Validate())
	assert.NotNil(t, References5{}.Validate())
	assert.NotNil(t, AccountStatusModification1{}.Validate())
	assert.Nil(t, AdditionalInformation5{}.Validate())
	assert.Nil(t, AddressModification2{}.Validate())
	assert.Nil(t, AmountModification1{}.Validate())
	assert.Nil(t, DateModification1{}.Validate())
	assert.NotNil(t, FullLegalNameModification1{}.Validate())
	assert.NotNil(t, NameModification1{}.Validate())
	assert.NotNil(t, NumberModification1{}.Validate())
	assert.NotNil(t, OrganisationModification2{}.Validate())
	assert.Nil(t, PartyModification2{}.Validate())
	assert.NotNil(t, PurposeModification1{}.Validate())
	assert.NotNil(t, RestrictionModification1{}.Validate())
	assert.NotNil(t, StatementFrequencyAndFormModification1{}.Validate())
	assert.NotNil(t, TradingNameModification1{}.Validate())
	assert.NotNil(t, TypeModification1{}.Validate())
	assert.NotNil(t, AccountExcludedMandateMaintenanceAmendmentRequestV03{}.Validate())
	assert.NotNil(t, CustomerAccountModification1{}.Validate())
	assert.NotNil(t, AccountMandateMaintenanceRequestV03{}.Validate())
	assert.NotNil(t, Group3{}.Validate())
	assert.NotNil(t, OperationMandate5{}.Validate())
	assert.Nil(t, Organisation34{}.Validate())
	assert.NotNil(t, PartyAndAuthorisation5{}.Validate())
	assert.Nil(t, PartyAndCertificate5{}.Validate())
	assert.NotNil(t, AccountMandateMaintenanceAmendmentRequestV03{}.Validate())
	assert.NotNil(t, AccountClosingRequestV03{}.Validate())
	assert.NotNil(t, AccountForAction2{}.Validate())
	assert.NotNil(t, AccountForAction1{}.Validate())
	assert.NotNil(t, AccountClosingAmendmentRequestV03{}.Validate())
	assert.Nil(t, AccountContract4{}.Validate())
	assert.NotNil(t, AccountClosingAdditionalInformationRequestV03{}.Validate())
	assert.NotNil(t, AccountSwitchInformationRequestV03{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.Nil(t, CitizenshipInformation1{}.Validate())
	assert.Nil(t, CommunicationAddress3{}.Validate())
	assert.NotNil(t, CountryAndResidentialStatusType1{}.Validate())
	assert.NotNil(t, CreditTransferTransaction41{}.Validate())
	assert.NotNil(t, GenericIdentification44{}.Validate())
	assert.NotNil(t, GenericIdentification47{}.Validate())
	assert.NotNil(t, IndividualPerson36{}.Validate())
	assert.NotNil(t, IndividualPersonNameLong2{}.Validate())
	assert.NotNil(t, NewAccount2{}.Validate())
	assert.NotNil(t, Organisation35{}.Validate())
	assert.NotNil(t, OtherIdentification1Choice{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxInformation8{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
	assert.NotNil(t, TransferInstruction1{}.Validate())
	assert.NotNil(t, AccountSwitchBalanceTransferAcknowledgementV03{}.Validate())
	assert.NotNil(t, AccountSwitchDetails1{}.Validate())
	assert.NotNil(t, AmountAndDirection5{}.Validate())
	assert.Nil(t, BalanceTransfer3{}.Validate())
	assert.NotNil(t, BalanceTransferFundingLimit1{}.Validate())
	assert.NotNil(t, BalanceTransferReference1{}.Validate())
	assert.NotNil(t, CashAccount39{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.Nil(t, Cheque11{}.Validate())
	assert.NotNil(t, ChequeDeliveryMethod1Choice{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.Nil(t, EndPoint1Choice{}.Validate())
	assert.Nil(t, Frequency1{}.Validate())
	assert.NotNil(t, Frequency37Choice{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.Nil(t, InstructionForCreditorAgent3{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.NotNil(t, PaymentIdentification6{}.Validate())
	assert.Nil(t, PaymentTypeInformation26{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RegulatoryAuthority2{}.Validate())
	assert.Nil(t, RegulatoryReporting3{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.Nil(t, RemittanceLocation6{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.NotNil(t, SettlementMethod3Choice{}.Validate())
	assert.Nil(t, StructuredRegulatoryReporting3{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.NotNil(t, AccountSwitchInformationResponseV03{}.Validate())
	assert.NotNil(t, DirectDebitInstructionDetails2{}.Validate())
	assert.NotNil(t, PaymentInstruction36{}.Validate())
	assert.NotNil(t, AccountSwitchCancelExistingPaymentV03{}.Validate())
	assert.NotNil(t, AccountSwitchRequestBalanceTransferV03{}.Validate())
	assert.NotNil(t, AccountSwitchRequestPaymentV03{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 CommunicationMethod2Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "POST"
	assert.Nil(t, type1.Validate())

	var type2 CommunicationMethod3Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "ONLI"
	assert.Nil(t, type2.Validate())

	var type3 ExternalAccountIdentification1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalBankTransactionDomain1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type5 ExternalBankTransactionFamily1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalBankTransactionSubFamily1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalCashAccountType1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalCommunicationFormat1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalPersonIdentification1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalProxyAccountType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 Frequency7Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.NotNil(t, type14.Validate())
	type14 = "INDA"
	assert.Nil(t, type14.Validate())

	var type15 PreferredContactMethod1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.NotNil(t, type15.Validate())
	type15 = "CELL"
	assert.Nil(t, type15.Validate())

	var type16 Unlimited9Text
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 Modification1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.NotNil(t, type17.Validate())
	type17 = "ADDD"
	assert.Nil(t, type17.Validate())

	var type18 BalanceTransferWindow1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.NotNil(t, type18.Validate())
	type18 = "EARL"
	assert.Nil(t, type18.Validate())

	var type19 BusinessDayConvention1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.NotNil(t, type19.Validate())
	type19 = "PREC"
	assert.Nil(t, type19.Validate())

	var type20 ChargeBearerType1Code
	assert.NotNil(t, type20.Validate())
	type20 = "test"
	assert.NotNil(t, type20.Validate())
	type20 = "SLEV"
	assert.Nil(t, type20.Validate())

	var type21 ChequeDelivery1Code
	assert.NotNil(t, type21.Validate())
	type21 = "test"
	assert.NotNil(t, type21.Validate())
	type21 = "RGFA"
	assert.Nil(t, type21.Validate())

	var type22 ChequeDelivery1Code
	assert.NotNil(t, type22.Validate())
	type22 = "test"
	assert.NotNil(t, type22.Validate())
	type22 = "CRFA"
	assert.Nil(t, type22.Validate())

	var type23 ChequeType2Code
	assert.NotNil(t, type23.Validate())
	type23 = "test"
	assert.NotNil(t, type23.Validate())
	type23 = "ELDR"
	assert.Nil(t, type23.Validate())

	var type24 DocumentType3Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "SCOR"
	assert.Nil(t, type24.Validate())

	var type25 DocumentType6Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "PUOR"
	assert.Nil(t, type25.Validate())

	var type27 ExternalTaxAmountType1Code
	assert.NotNil(t, type27.Validate())
	type27 = "test"
	assert.Nil(t, type27.Validate())

	var type28 Frequency10Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "QURT"
	assert.Nil(t, type28.Validate())

	var type29 Gender1Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MALE"
	assert.Nil(t, type29.Validate())

	var type30 OrganisationLegalStatus1Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "PCLG"
	assert.Nil(t, type30.Validate())

	var type31 PersonIdentificationType5Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "GUNL"
	assert.Nil(t, type31.Validate())

	var type32 Priority2Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "NORM"
	assert.Nil(t, type32.Validate())

	var type33 RegulatoryReportingType1Code
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.NotNil(t, type33.Validate())
	type33 = "BOTH"
	assert.Nil(t, type33.Validate())

	var type34 RemittanceLocationMethod2Code
	assert.NotNil(t, type34.Validate())
	type34 = "test"
	assert.NotNil(t, type34.Validate())
	type34 = "SMSM"
	assert.Nil(t, type34.Validate())

	var type35 ResidentialStatus1Code
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "NRES"
	assert.Nil(t, type35.Validate())

	var type36 SwitchType1Code
	assert.NotNil(t, type36.Validate())
	type36 = "test"
	assert.NotNil(t, type36.Validate())
	type36 = "PART"
	assert.Nil(t, type36.Validate())

	var type37 TaxRateMarker1Code
	assert.NotNil(t, type37.Validate())
	type37 = "test"
	assert.NotNil(t, type37.Validate())
	type37 = "GRSS"
	assert.Nil(t, type37.Validate())

	var type38 TaxRecordPeriod1Code
	assert.NotNil(t, type38.Validate())
	type38 = "test"
	assert.NotNil(t, type38.Validate())
	type38 = "MM05"
	assert.Nil(t, type38.Validate())

	var type39 SwitchStatus1Code
	assert.NotNil(t, type39.Validate())
	type39 = "test"
	assert.NotNil(t, type39.Validate())
	type39 = "COMP"
	assert.Nil(t, type39.Validate())

	var type40 ExternalCategoryPurpose1Code
	assert.NotNil(t, type40.Validate())
	type40 = "test"
	assert.Nil(t, type40.Validate())

	var type41 ExternalCreditorAgentInstruction1Code
	assert.NotNil(t, type41.Validate())
	type41 = "test"
	assert.Nil(t, type41.Validate())

	var type42 ExternalDiscountAmountType1Code
	assert.NotNil(t, type42.Validate())
	type42 = "test"
	assert.Nil(t, type42.Validate())

	var type43 ExternalDocumentLineType1Code
	assert.NotNil(t, type43.Validate())
	type43 = "test"
	assert.Nil(t, type43.Validate())

	var type44 ExternalGarnishmentType1Code
	assert.NotNil(t, type44.Validate())
	type44 = "test"
	assert.Nil(t, type44.Validate())

	var type45 ExternalLocalInstrument1Code
	assert.NotNil(t, type45.Validate())
	type45 = "test"
	assert.Nil(t, type45.Validate())

	var type46 ExternalPurpose1Code
	assert.NotNil(t, type46.Validate())
	type46 = "test"
	assert.Nil(t, type46.Validate())

	var type47 ExternalServiceLevel1Code
	assert.NotNil(t, type47.Validate())
	type47 = "test"
	assert.Nil(t, type47.Validate())

	var type48 PaymentMethod3Code
	assert.NotNil(t, type48.Validate())
	type48 = "test"
	assert.NotNil(t, type48.Validate())
	type48 = "TRA"
	assert.Nil(t, type48.Validate())

	var type49 UseCases1Code
	assert.NotNil(t, type49.Validate())
	type49 = "test"
	assert.NotNil(t, type49.Validate())
	type49 = "VIEW"
	assert.Nil(t, type49.Validate())
}
