// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type Admi00200101 struct {
	XMLName xml.Name         `xml:"admi.002.001.01"`
	RltdRef MessageReference `xml:"RltdRef"`
	Rsn     RejectionReason2 `xml:"Rsn"`
}

func (r Admi00200101) Validate() error {
	return utils.Validate(&r)
}

type MessageReference struct {
	Ref common.Max35Text `xml:"Ref"`
}

func (r MessageReference) Validate() error {
	return utils.Validate(&r)
}

type RejectionReason2 struct {
	RjctgPtyRsn common.Max35Text     `xml:"RjctgPtyRsn"`
	RjctnDtTm   *common.ISODateTime  `xml:"RjctnDtTm,omitempty" json:",omitempty"`
	ErrLctn     *common.Max350Text   `xml:"ErrLctn,omitempty" json:",omitempty"`
	RsnDesc     *common.Max350Text   `xml:"RsnDesc,omitempty" json:",omitempty"`
	AddtlData   *common.Max20000Text `xml:"AddtlData,omitempty" json:",omitempty"`
}

type AddtlRawData struct {
	Content common.Max20000Text `xml:",cdata"`
}

func (r RejectionReason2) Validate() error {
	return utils.Validate(&r)
}

type Admi00400101 struct {
	XMLName xml.Name `xml:"admi.004.001.01"`
	EvtInf  Event1   `xml:"EvtInf"`
}

func (r Admi00400101) Validate() error {
	return utils.Validate(&r)
}

type Event1 struct {
	EvtCd    common.Max4AlphaNumericText `xml:"EvtCd"`
	EvtParam []common.Max35Text          `xml:"EvtParam,omitempty" json:",omitempty"`
	EvtDesc  *common.Max350Text          `xml:"EvtDesc,omitempty" json:",omitempty"`
	EvtTm    *common.ISODateTime         `xml:"EvtTm,omitempty" json:",omitempty"`
}

func (r Event1) Validate() error {
	return utils.Validate(&r)
}

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

func (r AccountIdentification4Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountIdentificationSearchCriteria2Choice struct {
	EQ     AccountIdentification4Choice `xml:"EQ"`
	CTTxt  common.Max35Text             `xml:"CTTxt"`
	NCTTxt common.Max35Text             `xml:"NCTTxt"`
}

func (r AccountIdentificationSearchCriteria2Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

func (r AccountSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
	return utils.Validate(&r)
}

type BalanceType11Choice struct {
	Cd    ExternalSystemBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text               `xml:"Prtry"`
}

func (r BalanceType11Choice) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification6) Validate() error {
	return utils.Validate(&r)
}

type BranchData3 struct {
	Id      *common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData3) Validate() error {
	return utils.Validate(&r)
}

type CashBalance12 struct {
	Tp       []BalanceType11Choice                          `xml:"Tp,omitempty" json:",omitempty"`
	CtrPtyTp BalanceCounterparty1Code                       `xml:"CtrPtyTp"`
	CtrPtyId []BranchAndFinancialInstitutionIdentification6 `xml:"CtrPtyId,omitempty" json:",omitempty"`
	ValDt    []DateAndDateTimeSearch4Choice                 `xml:"ValDt,omitempty" json:",omitempty"`
	PrcgDt   *DateAndDateTimeSearch4Choice                  `xml:"PrcgDt,omitempty" json:",omitempty"`
}

func (r CashBalance12) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTimeSearch4Choice struct {
	DtTm DateTimeSearch2Choice   `xml:"DtTm"`
	Dt   DatePeriodSearch1Choice `xml:"Dt"`
}

func (r DateAndDateTimeSearch4Choice) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodSearch1Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
	EQDt   common.ISODate `xml:"EQDt"`
	NEQDt  common.ISODate `xml:"NEQDt"`
}

func (r DatePeriodSearch1Choice) Validate() error {
	return utils.Validate(&r)
}

type DateTimePeriod1 struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
}

func (r DateTimePeriod1) Validate() error {
	return utils.Validate(&r)
}

type DateTimePeriod1Choice struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
	DtTmRg DateTimePeriod1    `xml:"DtTmRg"`
}

func (r DateTimePeriod1Choice) Validate() error {
	return utils.Validate(&r)
}

type DateTimeSearch2Choice struct {
	FrDtTm   common.ISODateTime `xml:"FrDtTm"`
	ToDtTm   common.ISODateTime `xml:"ToDtTm"`
	FrToDtTm DateTimePeriod1    `xml:"FrToDtTm"`
	EQDtTm   common.ISODateTime `xml:"EQDtTm"`
	NEQDtTm  common.ISODateTime `xml:"NEQDtTm"`
}

func (r DateTimeSearch2Choice) Validate() error {
	return utils.Validate(&r)
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r EventType1Choice) Validate() error {
	return utils.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification18) Validate() error {
	return utils.Validate(&r)
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericFinancialIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification1 struct {
	Id      common.Max35Text  `xml:"Id"`
	SchmeNm *common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification36 struct {
	Id      common.Max35Text  `xml:"Id"`
	Issr    common.Max35Text  `xml:"Issr"`
	SchmeNm *common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification36) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader7 struct {
	MsgId       common.Max35Text        `xml:"MsgId"`
	CreDtTm     *common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp       *RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       *common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

func (r MessageHeader7) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress5 struct {
	Nm  common.Max350Text `xml:"Nm"`
	Adr *PostalAddress1   `xml:"Adr,omitempty" json:",omitempty"`
}

func (r NameAndAddress5) Validate() error {
	return utils.Validate(&r)
}

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId *common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalBusinessQuery1) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification120Choice struct {
	AnyBIC   common.AnyBICDec2014Identifier `xml:"AnyBIC"`
	PrtryId  GenericIdentification36        `xml:"PrtryId"`
	NmAndAdr NameAndAddress5                `xml:"NmAndAdr"`
}

func (r PartyIdentification120Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification136 struct {
	Id  PartyIdentification120Choice `xml:"Id"`
	LEI *common.LEIIdentifier        `xml:"LEI,omitempty" json:",omitempty"`
}

func (r PartyIdentification136) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress1 struct {
	AdrTp       *common.AddressType2Code `xml:"AdrTp,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text       `xml:"AdrLine,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text        `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text        `xml:"BldgNb,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text        `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text        `xml:"TwnNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text        `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        common.CountryCode       `xml:"Ctry"`
}

func (r PostalAddress1) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	BldgNm      *common.Max35Text   `xml:"BldgNm,omitempty" json:",omitempty"`
	Flr         *common.Max70Text   `xml:"Flr,omitempty" json:",omitempty"`
	PstBx       *common.Max16Text   `xml:"PstBx,omitempty" json:",omitempty"`
	Room        *common.Max70Text   `xml:"Room,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   *common.Max35Text   `xml:"TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    *common.Max35Text   `xml:"DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text  `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress24) Validate() error {
	return utils.Validate(&r)
}

type ReportQueryCriteria2 struct {
	NewQryNm common.Max35Text           `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  ReportQuerySearchCriteria2 `xml:"SchCrit"`
}

func (r ReportQueryCriteria2) Validate() error {
	return utils.Validate(&r)
}

type ReportQueryRequestV01 struct {
	XMLName     xml.Name               `xml:"RptQryReq"`
	MsgHdr      MessageHeader7         `xml:"MsgHdr"`
	RptQryCrit  []ReportQueryCriteria2 `xml:"RptQryCrit,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReportQueryRequestV01) Validate() error {
	return utils.Validate(&r)
}

type ReportQuerySearchCriteria2 struct {
	AcctId       []AccountIdentificationSearchCriteria2Choice `xml:"AcctId,omitempty" json:",omitempty"`
	Bal          []CashBalance12                              `xml:"Bal,omitempty" json:",omitempty"`
	RptNm        *common.Max4AlphaNumericText                 `xml:"RptNm,omitempty" json:",omitempty"`
	MsgNmId      *common.Max35Text                            `xml:"MsgNmId,omitempty" json:",omitempty"`
	PtyId        PartyIdentification136                       `xml:"PtyId"`
	RspnsblPtyId *PartyIdentification136                      `xml:"RspnsblPtyId,omitempty" json:",omitempty"`
	DtSch        *DatePeriodSearch1Choice                     `xml:"DtSch,omitempty" json:",omitempty"`
	SchdldTm     *DateTimePeriod1Choice                       `xml:"SchdldTm,omitempty" json:",omitempty"`
	Evt          *EventType1Choice                            `xml:"Evt,omitempty" json:",omitempty"`
}

func (r ReportQuerySearchCriteria2) Validate() error {
	return utils.Validate(&r)
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

func (r RequestType4Choice) Validate() error {
	return utils.Validate(&r)
}

type ResendRequestV01 struct {
	XMLName     xml.Name                `xml:"RsndReq"`
	MsgHdr      MessageHeader7          `xml:"MsgHdr"`
	RsndSchCrit []ResendSearchCriteria2 `xml:"RsndSchCrit,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ResendRequestV01) Validate() error {
	return utils.Validate(&r)
}

type ResendSearchCriteria2 struct {
	BizDt        *common.ISODate        `xml:"BizDt,omitempty" json:",omitempty"`
	SeqNb        *common.Max35Text      `xml:"SeqNb,omitempty" json:",omitempty"`
	SeqRg        *SequenceRange1Choice  `xml:"SeqRg,omitempty" json:",omitempty"`
	OrgnlMsgNmId *common.Max35Text      `xml:"OrgnlMsgNmId,omitempty" json:",omitempty"`
	FileRef      *common.Max35Text      `xml:"FileRef,omitempty" json:",omitempty"`
	Rcpt         PartyIdentification136 `xml:"Rcpt"`
}

func (r ResendSearchCriteria2) Validate() error {
	return utils.Validate(&r)
}

type SequenceRange1 struct {
	FrSeq common.Max35Text `xml:"FrSeq"`
	ToSeq common.Max35Text `xml:"ToSeq"`
}

func (r SequenceRange1) Validate() error {
	return utils.Validate(&r)
}

type SequenceRange1Choice struct {
	FrSeq   common.Max35Text   `xml:"FrSeq"`
	ToSeq   common.Max35Text   `xml:"ToSeq"`
	FrToSeq []SequenceRange1   `xml:"FrToSeq"`
	EQSeq   []common.Max35Text `xml:"EQSeq"`
	NEQSeq  []common.Max35Text `xml:"NEQSeq"`
}

func (r SequenceRange1Choice) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

func (r SupplementaryData1) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

func (r SupplementaryDataEnvelope1) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader10 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	QryNm   *common.Max35Text   `xml:"QryNm,omitempty" json:",omitempty"`
}

func (r MessageHeader10) Validate() error {
	return utils.Validate(&r)
}

type MessageReference1 struct {
	Ref     common.Max35Text       `xml:"Ref"`
	MsgNm   *common.Max35Text      `xml:"MsgNm,omitempty" json:",omitempty"`
	RefIssr PartyIdentification136 `xml:"RefIssr,omitempty" json:",omitempty"`
}

func (r MessageReference1) Validate() error {
	return utils.Validate(&r)
}

type ReceiptAcknowledgementReport2 struct {
	RltdRef MessageReference1 `xml:"RltdRef"`
	ReqHdlg RequestHandling2  `xml:"ReqHdlg"`
}

func (r ReceiptAcknowledgementReport2) Validate() error {
	return utils.Validate(&r)
}

type ReceiptAcknowledgementV01 struct {
	XMLName     xml.Name                        `xml:"RctAck"`
	MsgId       MessageHeader10                 `xml:"MsgId"`
	Rpt         []ReceiptAcknowledgementReport2 `xml:"Rpt,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReceiptAcknowledgementV01) Validate() error {
	return utils.Validate(&r)
}

type RequestHandling2 struct {
	StsCd   common.Max4AlphaNumericText `xml:"StsCd"`
	StsDtTm *common.ISODateTime         `xml:"StsDtTm,omitempty" json:",omitempty"`
	Desc    *common.Max140Text          `xml:"Desc,omitempty" json:",omitempty"`
}

func (r RequestHandling2) Validate() error {
	return utils.Validate(&r)
}

type SystemEventAcknowledgementV01 struct {
	XMLName     xml.Name                       `xml:"SysEvtAck"`
	MsgId       common.Max35Text               `xml:"MsgId"`
	OrgtrRef    *common.Max35Text              `xml:"OrgtrRef,omitempty" json:",omitempty"`
	SttlmSsnIdr *common.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty" json:",omitempty"`
	AckDtls     *Event1                        `xml:"AckDtls,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1           `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r SystemEventAcknowledgementV01) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress8 struct {
	Nm         common.Max350Text  `xml:"Nm"`
	Adr        *PostalAddress1    `xml:"Adr,omitempty" json:",omitempty"`
	AltrntvIdr []common.Max35Text `xml:"AltrntvIdr,omitempty" json:",omitempty"`
}

func (r NameAndAddress8) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification44 struct {
	AnyBIC     common.AnyBICIdentifier `xml:"AnyBIC"`
	AltrntvIdr []common.Max35Text      `xml:"AltrntvIdr,omitempty" json:",omitempty"`
}

func (r PartyIdentification44) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification59 struct {
	PtyNm      *common.Max34Text                    `xml:"PtyNm,omitempty" json:",omitempty"`
	AnyBIC     *PartyIdentification44               `xml:"AnyBIC,omitempty" json:",omitempty"`
	AcctNb     *common.Max34Text                    `xml:"AcctNb,omitempty" json:",omitempty"`
	Adr        *common.Max105Text                   `xml:"Adr,omitempty" json:",omitempty"`
	ClrSysId   *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	LglNttyIdr *common.LEIIdentifier                `xml:"LglNttyIdr,omitempty" json:",omitempty"`
}

func (r PartyIdentification59) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification73Choice struct {
	NmAndAdr NameAndAddress8       `xml:"NmAndAdr"`
	AnyBIC   PartyIdentification44 `xml:"AnyBIC"`
	PtyId    PartyIdentification59 `xml:"PtyId"`
}

func (r PartyIdentification73Choice) Validate() error {
	return utils.Validate(&r)
}

type ProcessingRequestV01 struct {
	XMLName     xml.Name                       `xml:"PrcgReq"`
	MsgId       common.Max35Text               `xml:"MsgId"`
	SttlmSsnIdr *common.Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty" json:",omitempty"`
	Req         RequestDetails19               `xml:"Req"`
}

func (r ProcessingRequestV01) Validate() error {
	return utils.Validate(&r)
}

type RequestDetails19 struct {
	Tp          common.Max35Text             `xml:"Tp"`
	RqstrId     *PartyIdentification73Choice `xml:"RqstrId,omitempty" json:",omitempty"`
	AddtlReqInf []common.Max35Text           `xml:"AddtlReqInf,omitempty" json:",omitempty"`
}

func (r RequestDetails19) Validate() error {
	return utils.Validate(&r)
}
