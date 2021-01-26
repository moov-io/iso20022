// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

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

func TestDocumentReda06600101(t *testing.T) {
	sample := DocumentReda06600101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda06600101{
		ReqToPayCdtrEnrlmntReq: RequestToPayCreditorEnrolmentRequestV01{
			Hdr: EnrolmentHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayCdtrEnrlmntReq":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}},"CdtrEnrlmnt":null,"ActvtnData":{"LtdPresntmntInd":false,"ActvtnReqDlvryPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.066.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayCdtrEnrlmntReq xmlns="urn:iso:std:iso:20022:tech:xsd:reda.066.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr><ActvtnData><LtdPresntmntInd>false</LtdPresntmntInd><ActvtnReqDlvryPty></ActvtnReqDlvryPty></ActvtnData></ReqToPayCdtrEnrlmntReq></Document>`)
}

func TestDocumentReda07200101(t *testing.T) {
	sample := DocumentReda07200101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda07200101{
		ReqToPayDbtrActvtnCxlReq: RequestToPayDebtorActivationCancellationRequestV01{
			Hdr: ActivationHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayDbtrActvtnCxlReq":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.072.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayDbtrActvtnCxlReq xmlns="urn:iso:std:iso:20022:tech:xsd:reda.072.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr></ReqToPayDbtrActvtnCxlReq></Document>`)
}

func TestDocumentReda06700101(t *testing.T) {
	sample := DocumentReda06700101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda06700101{
		ReqToPayCdtrEnrlmntAmdmntReq: RequestToPayCreditorEnrolmentAmendmentRequestV01{
			Hdr: EnrolmentHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayCdtrEnrlmntAmdmntReq":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.067.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayCdtrEnrlmntAmdmntReq xmlns="urn:iso:std:iso:20022:tech:xsd:reda.067.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr></ReqToPayCdtrEnrlmntAmdmntReq></Document>`)
}

func TestDocumentReda07300101(t *testing.T) {
	sample := DocumentReda07300101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda07300101{
		ReqToPayDbtrActvtnStsRpt: RequestToPayDebtorActivationStatusReportV01{
			Hdr: ActivationHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayDbtrActvtnStsRpt":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.073.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayDbtrActvtnStsRpt xmlns="urn:iso:std:iso:20022:tech:xsd:reda.073.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr></ReqToPayDbtrActvtnStsRpt></Document>`)
}

func TestDocumentReda07100101(t *testing.T) {
	sample := DocumentReda07100101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda07100101{
		ReqToPayDbtrActvtnAmdmntReq: RequestToPayDebtorActivationAmendmentRequestV01{
			Hdr: ActivationHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayDbtrActvtnAmdmntReq":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.071.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayDbtrActvtnAmdmntReq xmlns="urn:iso:std:iso:20022:tech:xsd:reda.071.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr></ReqToPayDbtrActvtnAmdmntReq></Document>`)
}

func TestDocumentReda07000101(t *testing.T) {
	sample := DocumentReda07000101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda07000101{
		ReqToPayDbtrActvtnReq: RequestToPayDebtorActivationRequestV01{
			Hdr: ActivationHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
			ElctrncInvcData: ElectronicInvoice1{
				PresntmntTp: "FULL",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayDbtrActvtnReq":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}},"ElctrncInvcData":{"PresntmntTp":"FULL"}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.070.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayDbtrActvtnReq xmlns="urn:iso:std:iso:20022:tech:xsd:reda.070.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr><ElctrncInvcData><PresntmntTp>FULL</PresntmntTp></ElctrncInvcData></ReqToPayDbtrActvtnReq></Document>`)
}

func TestDocumentReda06800101(t *testing.T) {
	sample := DocumentReda06800101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda06800101{
		ReqToPayCdtrEnrlmntCxlReq: RequestToPayCreditorEnrolmentCancellationRequestV01{
			Hdr: EnrolmentHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayCdtrEnrlmntCxlReq":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.068.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayCdtrEnrlmntCxlReq xmlns="urn:iso:std:iso:20022:tech:xsd:reda.068.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr></ReqToPayCdtrEnrlmntCxlReq></Document>`)
}

func TestDocumentReda06000101(t *testing.T) {
	sample := DocumentReda06000101{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, testTimeString)
	sample = DocumentReda06000101{
		ReqToPayCdtrEnrlmntStsRpt: RequestToPayCreditorEnrolmentStatusReportV01{
			Hdr: EnrolmentHeader2{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"ReqToPayCdtrEnrlmntStsRpt":{"Hdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:reda.069.001.01" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ReqToPayCdtrEnrlmntStsRpt xmlns="urn:iso:std:iso:20022:tech:xsd:reda.069.001.01"><Hdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><InitgPty></InitgPty></Hdr></ReqToPayCdtrEnrlmntStsRpt></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.NotNil(t, CreditorEnrolment3{}.Validate())
	assert.Nil(t, CreditorInvoice3{}.Validate())
	assert.Nil(t, CreditorServiceEnrolment1{}.Validate())
	assert.Nil(t, CustomerTypeRequest2{}.Validate())
	assert.Nil(t, DateAndDateTime2Choice{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.NotNil(t, DocumentFormat2Choice{}.Validate())
	assert.NotNil(t, DocumentType1Choice{}.Validate())
	assert.NotNil(t, EnrolmentHeader2{}.Validate())
	assert.NotNil(t, GenericIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericOrganisationType1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonType1{}.Validate())
	assert.Nil(t, OrganisationIdentification37{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, OrganisationType2{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party49Choice{}.Validate())
	assert.Nil(t, PersonIdentification17{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PersonType2{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.Nil(t, RTPPartyIdentification1{}.Validate())
	assert.NotNil(t, RequestToPayCreditorEnrolmentRequestV01{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, Visibilty1{}.Validate())
	assert.Nil(t, CreditorEnrolment4{}.Validate())
	assert.NotNil(t, CreditorEnrolmentAmendment3{}.Validate())
	assert.Nil(t, CreditorEnrolmentAmendment4{}.Validate())
	assert.NotNil(t, CreditorEnrolmentAmendmentReason1Choice{}.Validate())
	assert.NotNil(t, CreditorEnrolmentAmendmentReason2{}.Validate())
	assert.Nil(t, CreditorInvoice4{}.Validate())
	assert.NotNil(t, OriginalBusinessInstruction1{}.Validate())
	assert.NotNil(t, OriginalEnrolment2Choice{}.Validate())
	assert.NotNil(t, RequestToPayCreditorEnrolmentAmendmentRequestV01{}.Validate())
	assert.NotNil(t, CreditorEnrolmentCancellation2{}.Validate())
	assert.NotNil(t, CreditorEnrolmentCancellationReason1Choice{}.Validate())
	assert.NotNil(t, CreditorEnrolmentCancellationReason2{}.Validate())
	assert.NotNil(t, RequestToPayCreditorEnrolmentCancellationRequestV01{}.Validate())
	assert.NotNil(t, CreditorEnrolmentStatusReason2{}.Validate())
	assert.NotNil(t, CreditorEnrolmentStatusReason2Choice{}.Validate())
	assert.NotNil(t, EnrolmentStatus2{}.Validate())
	assert.NotNil(t, RequestToPayCreditorEnrolmentStatusReportV01{}.Validate())
	assert.NotNil(t, ServiceStatus1Choice{}.Validate())
	assert.NotNil(t, ActivationHeader2{}.Validate())
	assert.NotNil(t, ContractReference1{}.Validate())
	assert.Nil(t, DebtorActivation3{}.Validate())
	assert.NotNil(t, ElectronicInvoice1{}.Validate())
	assert.NotNil(t, RequestToPayDebtorActivationRequestV01{}.Validate())
	assert.Nil(t, DebtorActivation4{}.Validate())
	assert.Nil(t, DebtorActivationAmendment3{}.Validate())
	assert.Nil(t, DebtorActivationAmendment4{}.Validate())
	assert.NotNil(t, DebtorActivationAmendmentReason1Choice{}.Validate())
	assert.NotNil(t, DebtorActivationAmendmentReason2{}.Validate())
	assert.NotNil(t, RequestToPayDebtorActivationAmendmentRequestV01{}.Validate())
	assert.Nil(t, DebtorActivationCancellation2{}.Validate())
	assert.NotNil(t, DebtorActivationCancellationReason1Choice{}.Validate())
	assert.NotNil(t, DebtorActivationCancellationReason2{}.Validate())
	assert.Nil(t, OriginalActivation2Choice{}.Validate())
	assert.NotNil(t, RequestToPayDebtorActivationCancellationRequestV01{}.Validate())
	assert.NotNil(t, ActivationStatus2{}.Validate())
	assert.NotNil(t, DebtorActivationStatusReason1Choice{}.Validate())
	assert.NotNil(t, DebtorActivationStatusReason2{}.Validate())
	assert.NotNil(t, RequestToPayDebtorActivationStatusReportV01{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalDocumentFormat1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalDocumentType1Code
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

	var type5 ExternalCreditorEnrolmentAmendmentReason1Code
	assert.NotNil(t, type5.Validate())
	type5 = "test"
	assert.Nil(t, type5.Validate())

	var type6 ExternalCreditorEnrolmentCancellationReason1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type7 ExternalCreditorEnrolmentStatusReason1Code
	assert.NotNil(t, type7.Validate())
	type7 = "test"
	assert.Nil(t, type7.Validate())

	var type8 ExternalDebtorActivationAmendmentReason1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type9 ExternalDebtorActivationCancellationReason1Code
	assert.NotNil(t, type9.Validate())
	type9 = "test"
	assert.Nil(t, type9.Validate())

	var type10 ExternalDebtorActivationStatusReason1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 PreferredContactMethod1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.NotNil(t, type11.Validate())
	type11 = "LETT"
	assert.Nil(t, type11.Validate())

	var type12 ServiceRequestStatus1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.NotNil(t, type12.Validate())
	type12 = "ACPT"
	assert.Nil(t, type12.Validate())

	var type13 PresentmentType1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.NotNil(t, type13.Validate())
	type13 = "FULL"
	assert.Nil(t, type13.Validate())
}
