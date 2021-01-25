// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type BusinessDayCriteria2 struct {
	NewQryNm *common.Max35Text            `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []BusinessDaySearchCriteria2 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *BusinessDayReturnCriteria2  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

func (r BusinessDayCriteria2) Validate() error {
	return utils.Validate(&r)
}

type BusinessDayCriteria3Choice struct {
	QryNm   common.Max35Text     `xml:"QryNm"`
	NewCrit BusinessDayCriteria2 `xml:"NewCrit"`
}

func (r BusinessDayCriteria3Choice) Validate() error {
	return utils.Validate(&r)
}

type BusinessDayQuery2 struct {
	QryTp *QueryType2Code             `xml:"QryTp,omitempty" json:",omitempty"`
	Crit  *BusinessDayCriteria3Choice `xml:"Crit,omitempty" json:",omitempty"`
}

func (r BusinessDayQuery2) Validate() error {
	return utils.Validate(&r)
}

type BusinessDayReturnCriteria2 struct {
	SysDtInd   bool `xml:"SysDtInd,omitempty" json:",omitempty"`
	SysStsInd  bool `xml:"SysStsInd,omitempty" json:",omitempty"`
	SysCcyInd  bool `xml:"SysCcyInd,omitempty" json:",omitempty"`
	ClsrPrdInd bool `xml:"ClsrPrdInd,omitempty" json:",omitempty"`
	EvtInd     bool `xml:"EvtInd,omitempty" json:",omitempty"`
	SsnPrdInd  bool `xml:"SsnPrdInd,omitempty" json:",omitempty"`
	EvtTpInd   bool `xml:"EvtTpInd,omitempty" json:",omitempty"`
}

func (r BusinessDayReturnCriteria2) Validate() error {
	return utils.Validate(&r)
}

type BusinessDaySearchCriteria2 struct {
	SysDt   *common.ISODate               `xml:"SysDt,omitempty" json:",omitempty"`
	SysId   []SystemIdentification2Choice `xml:"SysId,omitempty" json:",omitempty"`
	SysCcy  []common.ActiveCurrencyCode   `xml:"SysCcy,omitempty" json:",omitempty"`
	EvtTp   *SystemEventType2Choice       `xml:"EvtTp,omitempty" json:",omitempty"`
	ClsrPrd *DateTimePeriod1Choice        `xml:"ClsrPrd,omitempty" json:",omitempty"`
}

func (r BusinessDaySearchCriteria2) Validate() error {
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

type GenericIdentification1 struct {
	Id      common.Max35Text  `xml:"Id"`
	SchmeNm *common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GetBusinessDayInformationV05 struct {
	MsgHdr          MessageHeader9       `xml:"MsgHdr"`
	BizDayInfQryDef *BusinessDayQuery2   `xml:"BizDayInfQryDef,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r GetBusinessDayInformationV05) Validate() error {
	return utils.Validate(&r)
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r MarketInfrastructureIdentification1Choice) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader9 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   *RequestType4Choice `xml:"ReqTp,omitempty" json:",omitempty"`
}

func (r MessageHeader9) Validate() error {
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

type SystemEventType2Choice struct {
	Cd    SystemEventType2Code   `xml:"Cd"`
	Prtry GenericIdentification1 `xml:"Prtry"`
}

func (r SystemEventType2Choice) Validate() error {
	return utils.Validate(&r)
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

func (r SystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
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

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericFinancialIdentification1) Validate() error {
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

type LongPaymentIdentification2 struct {
	TxId           *common.Max35Text                            `xml:"TxId,omitempty" json:",omitempty"`
	UETR           *common.UUIDv4Identifier                     `xml:"UETR,omitempty" json:",omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                               `xml:"IntrBkSttlmDt"`
	PmtMtd         *PaymentOrigin1Choice                        `xml:"PmtMtd,omitempty" json:",omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
	NtryTp         *EntryTypeIdentifier                         `xml:"NtryTp,omitempty" json:",omitempty"`
	EndToEndId     *common.Max35Text                            `xml:"EndToEndId,omitempty" json:",omitempty"`
}

func (r LongPaymentIdentification2) Validate() error {
	return utils.Validate(&r)
}

type OriginalMessageAndIssuer1 struct {
	MsgId   common.Max35Text  `xml:"MsgId"`
	MsgNmId *common.Max35Text `xml:"MsgNmId,omitempty" json:",omitempty"`
	OrgtrNm *common.Max70Text `xml:"OrgtrNm,omitempty" json:",omitempty"`
}

func (r OriginalMessageAndIssuer1) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification6Choice struct {
	TxId      common.Max35Text                `xml:"TxId"`
	QId       QueueTransactionIdentification1 `xml:"QId"`
	LngBizId  LongPaymentIdentification2      `xml:"LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"ShrtBizId"`
	PrtryId   common.Max70Text                `xml:"PrtryId"`
}

func (r PaymentIdentification6Choice) Validate() error {
	return utils.Validate(&r)
}

type PaymentOrigin1Choice struct {
	FINMT    common.Max3NumericText `xml:"FINMT"`
	XMLMsgNm common.Max35Text       `xml:"XMLMsgNm"`
	Prtry    common.Max35Text       `xml:"Prtry"`
	Instrm   PaymentInstrument1Code `xml:"Instrm"`
}

func (r PaymentOrigin1Choice) Validate() error {
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

type QueueTransactionIdentification1 struct {
	QId    common.Max16Text `xml:"QId"`
	PosInQ common.Max16Text `xml:"PosInQ"`
}

func (r QueueTransactionIdentification1) Validate() error {
	return utils.Validate(&r)
}

type Receipt3 struct {
	OrgnlMsgId OriginalMessageAndIssuer1     `xml:"OrgnlMsgId"`
	OrgnlPmtId *PaymentIdentification6Choice `xml:"OrgnlPmtId,omitempty" json:",omitempty"`
	ReqHdlg    []RequestHandling1            `xml:"ReqHdlg,omitempty" json:",omitempty"`
}

func (r Receipt3) Validate() error {
	return utils.Validate(&r)
}

type ReceiptV05 struct {
	MsgHdr      MessageHeader9       `xml:"MsgHdr"`
	RctDtls     []Receipt3           `xml:"RctDtls" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReceiptV05) Validate() error {
	return utils.Validate(&r)
}

type RequestHandling1 struct {
	StsCd common.Max4AlphaNumericText `xml:"StsCd"`
	Desc  *common.Max140Text          `xml:"Desc,omitempty" json:",omitempty"`
}

func (r RequestHandling1) Validate() error {
	return utils.Validate(&r)
}

type ShortPaymentIdentification2 struct {
	TxId          common.Max35Text                             `xml:"TxId"`
	IntrBkSttlmDt common.ISODate                               `xml:"IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
}

func (r ShortPaymentIdentification2) Validate() error {
	return utils.Validate(&r)
}

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

func (r Case5) Validate() error {
	return utils.Validate(&r)
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r CaseAssignment5) Validate() error {
	return utils.Validate(&r)
}

type CaseForwardingNotification3 struct {
	Justfn CaseForwardingNotification3Code `xml:"Justfn"`
}

func (r CaseForwardingNotification3) Validate() error {
	return utils.Validate(&r)
}

type Contact4 struct {
	NmPrfx    *common.NamePrefix2Code      `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm        *common.Max140Text           `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb    *common.PhoneNumber          `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb     *common.PhoneNumber          `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb     *common.PhoneNumber          `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr  *common.Max2048Text          `xml:"EmailAdr,omitempty" json:",omitempty"`
	EmailPurp *common.Max35Text            `xml:"EmailPurp,omitempty" json:",omitempty"`
	JobTitl   *common.Max35Text            `xml:"JobTitl,omitempty" json:",omitempty"`
	Rspnsblty *common.Max35Text            `xml:"Rspnsblty,omitempty" json:",omitempty"`
	Dept      *common.Max70Text            `xml:"Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1              `xml:"Othr,omitempty" json:",omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:",omitempty"`
}

func (r Contact4) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
	return utils.Validate(&r)
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericOrganisationIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericPersonIdentification1) Validate() error {
	return utils.Validate(&r)
}

type NotificationOfCaseAssignmentV05 struct {
	Hdr         ReportHeader5               `xml:"Hdr"`
	Case        Case5                       `xml:"Case"`
	Assgnmt     CaseAssignment5             `xml:"Assgnmt"`
	Ntfctn      CaseForwardingNotification3 `xml:"Ntfctn"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r NotificationOfCaseAssignmentV05) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification29) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `xml:"ChanlTp"`
	Id      *common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

func (r OtherContact1) Validate() error {
	return utils.Validate(&r)
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

func (r Party38Choice) Validate() error {
	return utils.Validate(&r)
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (r Party40Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification135) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification13) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type ReportHeader5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Fr      Party40Choice      `xml:"Fr"`
	To      Party40Choice      `xml:"To"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r ReportHeader5) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryData6 struct {
	Any SkipPayload `xml:"Any"`
}

func (r ProprietaryData6) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryData7 struct {
	Tp   common.Max35Text `xml:"Tp"`
	Data ProprietaryData6 `xml:"Data"`
}

func (r ProprietaryData7) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryFormatInvestigationV05 struct {
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        *Case5               `xml:"Case,omitempty" json:",omitempty"`
	PrtryData   ProprietaryData7     `xml:"PrtryData"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ProprietaryFormatInvestigationV05) Validate() error {
	return utils.Validate(&r)
}

type SkipPayload struct {
	Item string `xml:",any"`
}

func (r SkipPayload) Validate() error {
	return utils.Validate(&r)
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type DebitAuthorisationConfirmation2 struct {
	DbtAuthstn bool                     `xml:"DbtAuthstn"`
	AmtToDbt   *ActiveCurrencyAndAmount `xml:"AmtToDbt,omitempty" json:",omitempty"`
	ValDtToDbt *common.ISODate          `xml:"ValDtToDbt,omitempty" json:",omitempty"`
	Rsn        *common.Max140Text       `xml:"Rsn,omitempty" json:",omitempty"`
}

func (r DebitAuthorisationConfirmation2) Validate() error {
	return utils.Validate(&r)
}

type DebitAuthorisationResponseV05 struct {
	Assgnmt     CaseAssignment5                 `xml:"Assgnmt"`
	Case        Case5                           `xml:"Case,omitempty" json:",omitempty"`
	Conf        DebitAuthorisationConfirmation2 `xml:"Conf"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DebitAuthorisationResponseV05) Validate() error {
	return utils.Validate(&r)
}

type CaseStatus2 struct {
	DtTm    common.ISODateTime `xml:"DtTm"`
	CaseSts CaseStatus2Code    `xml:"CaseSts"`
	Rsn     *common.Max140Text `xml:"Rsn,omitempty" json:",omitempty"`
}

func (r CaseStatus2) Validate() error {
	return utils.Validate(&r)
}

type CaseStatusReportV05 struct {
	Hdr         ReportHeader5        `xml:"Hdr"`
	Case        Case5                `xml:"Case"`
	Sts         CaseStatus2          `xml:"Sts"`
	NewAssgnmt  *CaseAssignment5     `xml:"NewAssgnmt,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CaseStatusReportV05) Validate() error {
	return utils.Validate(&r)
}

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

func (r AccountIdentification4Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

func (r AccountSchemeName1Choice) Validate() error {
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

type GetReservationV05 struct {
	MsgHdr       MessageHeader9       `xml:"MsgHdr"`
	RsvatnQryDef *ReservationQuery3   `xml:"RsvatnQryDef,omitempty" json:",omitempty"`
	SplmtryData  []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r GetReservationV05) Validate() error {
	return utils.Validate(&r)
}

type ReservationCriteria3Choice struct {
	QryNm   common.Max35Text     `xml:"QryNm"`
	NewCrit ReservationCriteria4 `xml:"NewCrit"`
}

func (r ReservationCriteria3Choice) Validate() error {
	return utils.Validate(&r)
}

type ReservationCriteria4 struct {
	NewQryNm *common.Max35Text            `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []ReservationSearchCriteria3 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *ReservationReturnCriteria1  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

func (r ReservationCriteria4) Validate() error {
	return utils.Validate(&r)
}

type ReservationQuery3 struct {
	QryTp      *QueryType2Code             `xml:"QryTp,omitempty" json:",omitempty"`
	RsvatnCrit *ReservationCriteria3Choice `xml:"RsvatnCrit,omitempty" json:",omitempty"`
}

func (r ReservationQuery3) Validate() error {
	return utils.Validate(&r)
}

type ReservationReturnCriteria1 struct {
	StartDtTmInd bool `xml:"StartDtTmInd,omitempty" json:",omitempty"`
	StsInd       bool `xml:"StsInd,omitempty" json:",omitempty"`
}

func (r ReservationReturnCriteria1) Validate() error {
	return utils.Validate(&r)
}

type ReservationSearchCriteria3 struct {
	SysId        *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	DfltRsvatnTp []ReservationType1Code                        `xml:"DfltRsvatnTp,omitempty" json:",omitempty"`
	CurRsvatnTp  []ReservationType1Code                        `xml:"CurRsvatnTp,omitempty" json:",omitempty"`
	AcctOwnr     *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId       *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

func (r ReservationSearchCriteria3) Validate() error {
	return utils.Validate(&r)
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

func (r Amount2Choice) Validate() error {
	return utils.Validate(&r)
}

type CurrentOrDefaultReservation2Choice struct {
	Cur  ReservationIdentification2 `xml:"Cur"`
	Dflt ReservationIdentification2 `xml:"Dflt"`
}

func (r CurrentOrDefaultReservation2Choice) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

func (r DateAndDateTime2Choice) Validate() error {
	return utils.Validate(&r)
}

type ModifyReservationV05 struct {
	MsgHdr          MessageHeader1                     `xml:"MsgHdr"`
	RsvatnId        CurrentOrDefaultReservation2Choice `xml:"RsvatnId"`
	NewRsvatnValSet Reservation4                       `xml:"NewRsvatnValSet"`
	SplmtryData     []SupplementaryData1               `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ModifyReservationV05) Validate() error {
	return utils.Validate(&r)
}

type Reservation4 struct {
	StartDtTm *DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice           `xml:"Amt"`
}

func (r Reservation4) Validate() error {
	return utils.Validate(&r)
}

type DeleteReservationV05 struct {
	MsgHdr      MessageHeader1              `xml:"MsgHdr"`
	CurRsvatn   *ReservationIdentification2 `xml:"CurRsvatn,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DeleteReservationV05) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r MessageHeader1) Validate() error {
	return utils.Validate(&r)
}

type ReservationIdentification2 struct {
	RsvatnId *common.Max35Text                             `xml:"RsvatnId,omitempty" json:",omitempty"`
	SysId    *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	Tp       ReservationType1Choice                        `xml:"Tp"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId   *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

func (r ReservationIdentification2) Validate() error {
	return utils.Validate(&r)
}

type ReservationType1Choice struct {
	Cd    ReservationType2Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r ReservationType1Choice) Validate() error {
	return utils.Validate(&r)
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice         `xml:"Id"`
	Tp   *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

func (r CashAccount38) Validate() error {
	return utils.Validate(&r)
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CashAccountType2Choice) Validate() error {
	return utils.Validate(&r)
}

type LiquidityCreditTransfer2 struct {
	LqdtyTrfId *PaymentIdentification8                       `xml:"LqdtyTrfId,omitempty" json:",omitempty"`
	Cdtr       *BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct   *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	TrfdAmt    Amount2Choice                                 `xml:"TrfdAmt"`
	Dbtr       *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct   *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	SttlmDt    *common.ISODate                               `xml:"SttlmDt,omitempty" json:",omitempty"`
}

func (r LiquidityCreditTransfer2) Validate() error {
	return utils.Validate(&r)
}

type LiquidityCreditTransferV05 struct {
	MsgHdr      MessageHeader1           `xml:"MsgHdr"`
	LqdtyCdtTrf LiquidityCreditTransfer2 `xml:"LqdtyCdtTrf"`
	SplmtryData []SupplementaryData1     `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r LiquidityCreditTransferV05) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification8 struct {
	InstrId    *common.Max35Text        `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId *common.Max35Text        `xml:"EndToEndId"`
	TxId       *common.Max35Text        `xml:"TxId,omitempty" json:",omitempty"`
	UETR       *common.UUIDv4Identifier `xml:"UETR,omitempty" json:",omitempty"`
}

func (r PaymentIdentification8) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

func (r ProxyAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

func (r ProxyAccountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type LiquidityDebitTransfer2 struct {
	LqdtyTrfId *PaymentIdentification8                       `xml:"LqdtyTrfId,omitempty" json:",omitempty"`
	Cdtr       *BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct   *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	TrfdAmt    *Amount2Choice                                `xml:"TrfdAmt"`
	Dbtr       *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct   *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	SttlmDt    *common.ISODate                               `xml:"SttlmDt,omitempty" json:",omitempty"`
}

func (r LiquidityDebitTransfer2) Validate() error {
	return utils.Validate(&r)
}

type LiquidityDebitTransferV05 struct {
	MsgHdr      MessageHeader1          `xml:"MsgHdr"`
	LqdtyDbtTrf LiquidityDebitTransfer2 `xml:"LqdtyDbtTrf"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r LiquidityDebitTransferV05) Validate() error {
	return utils.Validate(&r)
}

type AccountReportingRequestV05 struct {
	GrpHdr      GroupHeader77        `xml:"GrpHdr"`
	RptgReq     []ReportingRequest5  `xml:"RptgReq" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountReportingRequestV05) Validate() error {
	return utils.Validate(&r)
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type BalanceSubType1Choice struct {
	Cd    ExternalBalanceSubType1Code `xml:"Cd"`
	Prtry common.Max35Text            `xml:"Prtry"`
}

func (r BalanceSubType1Choice) Validate() error {
	return utils.Validate(&r)
}

type BalanceType10Choice struct {
	Cd    ExternalBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r BalanceType10Choice) Validate() error {
	return utils.Validate(&r)
}

type BalanceType13 struct {
	CdOrPrtry BalanceType10Choice    `xml:"CdOrPrtry"`
	SubTp     *BalanceSubType1Choice `xml:"SubTp,omitempty" json:",omitempty"`
}

func (r BalanceType13) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate  `xml:"FrDt"`
	ToDt *common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

func (r DatePeriodDetails1) Validate() error {
	return utils.Validate(&r)
}

type EntryStatus1Choice struct {
	Cd    ExternalEntryStatus1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r EntryStatus1Choice) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader77 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
	MsgSndr *Party40Choice     `xml:"MsgSndr,omitempty" json:",omitempty"`
}

func (r GroupHeader77) Validate() error {
	return utils.Validate(&r)
}

type Limit2 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd FloorLimitType1Code               `xml:"CdtDbtInd"`
}

func (r Limit2) Validate() error {
	return utils.Validate(&r)
}

type ReportingPeriod2 struct {
	FrToDt DatePeriodDetails1  `xml:"FrToDt"`
	FrToTm *TimePeriodDetails1 `xml:"FrToTm,omitempty" json:",omitempty"`
	Tp     QueryType3Code      `xml:"Tp"`
}

func (r ReportingPeriod2) Validate() error {
	return utils.Validate(&r)
}

type ReportingRequest5 struct {
	Id          *common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	ReqdMsgNmId common.Max35Text                              `xml:"ReqdMsgNmId"`
	Acct        *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr    Party40Choice                                 `xml:"AcctOwnr"`
	AcctSvcr    *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RptgPrd     *ReportingPeriod2                             `xml:"RptgPrd,omitempty" json:",omitempty"`
	RptgSeq     *SequenceRange1Choice                         `xml:"RptgSeq,omitempty" json:",omitempty"`
	ReqdTxTp    *TransactionType2                             `xml:"ReqdTxTp,omitempty" json:",omitempty"`
	ReqdBalTp   []BalanceType13                               `xml:"ReqdBalTp,omitempty" json:",omitempty"`
}

func (r ReportingRequest5) Validate() error {
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
	FrToSeq []SequenceRange1   `xml:"FrToSeq" json:",omitempty"`
	EQSeq   []common.Max35Text `xml:"EQSeq" json:",omitempty"`
	NEQSeq  []common.Max35Text `xml:"NEQSeq" json:",omitempty"`
}

func (r SequenceRange1Choice) Validate() error {
	return utils.Validate(&r)
}

type TimePeriodDetails1 struct {
	FrTm common.ISOTime  `xml:"FrTm"`
	ToTm *common.ISOTime `xml:"ToTm,omitempty" json:",omitempty"`
}

func (r TimePeriodDetails1) Validate() error {
	return utils.Validate(&r)
}

type TransactionType2 struct {
	Sts       EntryStatus1Choice     `xml:"Sts"`
	CdtDbtInd common.CreditDebitCode `xml:"CdtDbtInd"`
	FlrLmt    []Limit2               `xml:"FlrLmt,omitempty" json:",omitempty"`
}

func (r TransactionType2) Validate() error {
	return utils.Validate(&r)
}
