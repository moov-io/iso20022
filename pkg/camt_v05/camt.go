// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"encoding/xml"

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
	XMLName         xml.Name             `xml:"GetBizDayInf"`
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
	XMLName     xml.Name             `xml:"Rct"`
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
	XMLName     xml.Name                    `xml:"NtfctnOfCaseAssgnmt"`
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
	XMLName     xml.Name             `xml:"PrtryFrmtInvstgtn"`
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
	XMLName     xml.Name                        `xml:"DbtAuthstnRspn"`
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
	XMLName     xml.Name             `xml:"CaseStsRpt"`
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
	XMLName      xml.Name             `xml:"GetRsvatn"`
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
	XMLName         xml.Name                           `xml:"ModfyRsvatn"`
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
	XMLName     xml.Name                    `xml:"DelRsvatn"`
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
	XMLName     xml.Name                 `xml:"LqdtyCdtTrf"`
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
	TrfdAmt    Amount2Choice                                 `xml:"TrfdAmt"`
	Dbtr       *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct   *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	SttlmDt    *common.ISODate                               `xml:"SttlmDt,omitempty" json:",omitempty"`
}

func (r LiquidityDebitTransfer2) Validate() error {
	return utils.Validate(&r)
}

type LiquidityDebitTransferV05 struct {
	XMLName     xml.Name                `xml:"LqdtyDbtTrf"`
	MsgHdr      MessageHeader1          `xml:"MsgHdr"`
	LqdtyDbtTrf LiquidityDebitTransfer2 `xml:"LqdtyDbtTrf"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r LiquidityDebitTransferV05) Validate() error {
	return utils.Validate(&r)
}

type AccountReportingRequestV05 struct {
	XMLName     xml.Name             `xml:"AcctRptgReq"`
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

type AmendmentInformationDetails10 struct {
	OrgnlMndtId      *common.Max35Text                             `xml:"OrgnlMndtId,omitempty" json:",omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification43                        `xml:"OrgnlCdtrSchmeId,omitempty" json:",omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty" json:",omitempty"`
	OrgnlCdtrAgtAcct *CashAccount24                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlDbtr        *PartyIdentification43                        `xml:"OrgnlDbtr,omitempty" json:",omitempty"`
	OrgnlDbtrAcct    *CashAccount24                                `xml:"OrgnlDbtrAcct,omitempty" json:",omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty" json:",omitempty"`
	OrgnlDbtrAgtAcct *CashAccount24                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlFnlColltnDt *common.ISODate                               `xml:"OrgnlFnlColltnDt,omitempty" json:",omitempty"`
	OrgnlFrqcy       *Frequency21Choice                            `xml:"OrgnlFrqcy,omitempty" json:",omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:",omitempty"`
}

func (r AmendmentInformationDetails10) Validate() error {
	return utils.Validate(&r)
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

func (r AmountType4Choice) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"FinInstnId"`
	BrnchId    *BranchData2                        `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification5) Validate() error {
	return utils.Validate(&r)
}

type BranchData2 struct {
	Id      *common.Max35Text  `xml:"Id,omitempty" json:",omitempty"`
	Nm      *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress6    `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData2) Validate() error {
	return utils.Validate(&r)
}

type CancellationReason14Choice struct {
	Cd    CancellationReason5Code `xml:"Cd"`
	Prtry common.Max35Text        `xml:"Prtry"`
}

func (r CancellationReason14Choice) Validate() error {
	return utils.Validate(&r)
}

type Case3 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party12Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

func (r Case3) Validate() error {
	return utils.Validate(&r)
}

type CaseAssignment3 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party12Choice      `xml:"Assgnr"`
	Assgne  Party12Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r CaseAssignment3) Validate() error {
	return utils.Validate(&r)
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice         `xml:"Id"`
	Tp  *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm  *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
}

func (r CashAccount24) Validate() error {
	return utils.Validate(&r)
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r ClearingSystemIdentification3Choice) Validate() error {
	return utils.Validate(&r)
}

type ContactDetails2 struct {
	NmPrfx   *common.NamePrefix1Code `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm       *common.Max140Text      `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb   *common.PhoneNumber     `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb    *common.PhoneNumber     `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb    *common.PhoneNumber     `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max2048Text     `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr     *common.Max35Text       `xml:"Othr,omitempty" json:",omitempty"`
}

func (r ContactDetails2) Validate() error {
	return utils.Validate(&r)
}

type ControlData1 struct {
	NbOfTxs common.Max15NumericText `xml:"NbOfTxs"`
	CtrlSum float64                 `xml:"CtrlSum,omitempty" json:",omitempty"`
}

func (r ControlData1) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref *common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
}

func (r CreditorReferenceInformation2) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

func (r CreditorReferenceType1Choice) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType2) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriodDetails) Validate() error {
	return utils.Validate(&r)
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r DiscountAmountAndType1) Validate() error {
	return utils.Validate(&r)
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r DiscountAmountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd *CreditDebitCode                  `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DocumentAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     *common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt *common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

func (r DocumentLineIdentification1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc *common.Max2048Text           `xml:"Desc,omitempty" json:",omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty" json:",omitempty"`
}

func (r DocumentLineInformation1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
}

func (r DocumentLineType1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

func (r DocumentLineType1Choice) Validate() error {
	return utils.Validate(&r)
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

func (r EquivalentAmount2) Validate() error {
	return utils.Validate(&r)
}

type FIToFIPaymentCancellationRequestV05 struct {
	XMLName     xml.Name                  `xml:"FIToFIPmtCxlReq"`
	Assgnmt     CaseAssignment3           `xml:"Assgnmt"`
	Case        *Case3                    `xml:"Case,omitempty" json:",omitempty"`
	CtrlData    *ControlData1             `xml:"CtrlData,omitempty" json:",omitempty"`
	Undrlyg     []UnderlyingTransaction13 `xml:"Undrlyg"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r FIToFIPaymentCancellationRequestV05) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *common.BICFIIdentifier              `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification8) Validate() error {
	return utils.Validate(&r)
}

type Frequency21Choice struct {
	Tp  Frequency6Code   `xml:"Tp"`
	Prd FrequencyPeriod1 `xml:"Prd"`
}

func (r Frequency21Choice) Validate() error {
	return utils.Validate(&r)
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

func (r FrequencyPeriod1) Validate() error {
	return utils.Validate(&r)
}

type Garnishment1 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification43             `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification43             `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment1) Validate() error {
	return utils.Validate(&r)
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text      `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GarnishmentType1) Validate() error {
	return utils.Validate(&r)
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r GarnishmentType1Choice) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateRelatedInformation10 struct {
	MndtId        *common.Max35Text              `xml:"MndtId,omitempty" json:",omitempty"`
	DtOfSgntr     *common.ISODate                `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	AmdmntInd     bool                           `xml:"AmdmntInd,omitempty" json:",omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails10 `xml:"AmdmntInfDtls,omitempty" json:",omitempty"`
	ElctrncSgntr  *common.Max1025Text            `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstColltnDt  *common.ISODate                `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt   *common.ISODate                `xml:"FnlColltnDt,omitempty" json:",omitempty"`
	Frqcy         *Frequency21Choice             `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:",omitempty"`
}

func (r MandateRelatedInformation10) Validate() error {
	return utils.Validate(&r)
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

func (r MandateSetupReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification8 struct {
	AnyBIC *common.AnyBICIdentifier             `xml:"AnyBIC,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification8) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupHeader4 struct {
	GrpCxlId     *common.Max35Text            `xml:"GrpCxlId,omitempty" json:",omitempty"`
	Case         *Case3                       `xml:"Case,omitempty" json:",omitempty"`
	OrgnlMsgId   common.Max35Text             `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text             `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime          `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	NbOfTxs      *common.Max15NumericText     `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum      float64                      `xml:"CtrlSum,omitempty" json:",omitempty"`
	GrpCxl       bool                         `xml:"GrpCxl,omitempty" json:",omitempty"`
	CxlRsnInf    []PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty" json:",omitempty"`
}

func (r OriginalGroupHeader4) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalGroupInformation3) Validate() error {
	return utils.Validate(&r)
}

type OriginalTransactionReference22 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    *common.ISODate                               `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    *PartyIdentification43                        `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       *SettlementInstruction4                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       *PaymentTypeInformation25                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    *MandateRelatedInformation10                  `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         *RemittanceInformation11                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           *PartyIdentification43                        `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    *CashAccount24                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    *CashAccount24                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           *PartyIdentification43                        `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference22) Validate() error {
	return utils.Validate(&r)
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"OrgId"`
	PrvtId PersonIdentification5       `xml:"PrvtId"`
}

func (r Party11Choice) Validate() error {
	return utils.Validate(&r)
}

type Party12Choice struct {
	Pty PartyIdentification43                        `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (r Party12Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification43 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress6     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party11Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *ContactDetails2    `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification43) Validate() error {
	return utils.Validate(&r)
}

type PaymentCancellationReason2 struct {
	Orgtr    *PartyIdentification43      `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *CancellationReason14Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text         `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r PaymentCancellationReason2) Validate() error {
	return utils.Validate(&r)
}

type PaymentTransaction62 struct {
	CxlId               *common.Max35Text                             `xml:"CxlId,omitempty" json:",omitempty"`
	Case                *Case3                                        `xml:"Case,omitempty" json:",omitempty"`
	OrgnlGrpInf         *OriginalGroupInformation3                    `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                             `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                             `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                             `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlClrSysRef      *common.Max35Text                             `xml:"OrgnlClrSysRef,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmDt  *common.ISODate                               `xml:"OrgnlIntrBkSttlmDt,omitempty" json:",omitempty"`
	Assgnr              *BranchAndFinancialInstitutionIdentification5 `xml:"Assgnr,omitempty" json:",omitempty"`
	Assgne              *BranchAndFinancialInstitutionIdentification5 `xml:"Assgne,omitempty" json:",omitempty"`
	CxlRsnInf           []PaymentCancellationReason2                  `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	OrgnlTxRef          *OriginalTransactionReference22               `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	SplmtryData         []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentTransaction62) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation25 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     *SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation25) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification5) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress6 struct {
	AdrTp       *AddressType2Code   `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text  `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress6) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty" json:",omitempty"`
	Nb       *common.Max35Text          `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   *common.ISODate            `xml:"RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:",omitempty"`
}

func (r ReferredDocumentInformation7) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

func (r ReferredDocumentType3Choice) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text           `xml:"Issr,omitempty" json:",omitempty"`
}

func (r ReferredDocumentType4) Validate() error {
	return utils.Validate(&r)
}

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount2) Validate() error {
	return utils.Validate(&r)
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount3) Validate() error {
	return utils.Validate(&r)
}

type RemittanceInformation11 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation13 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation11) Validate() error {
	return utils.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceLevel8Choice) Validate() error {
	return utils.Validate(&r)
}

type SettlementInstruction4 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount24                                `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount24                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount24                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount24                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

func (r SettlementInstruction4) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation13 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification43         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification43         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation4               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment1                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation13) Validate() error {
	return utils.Validate(&r)
}

type TaxAmount1 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails1                `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount1) Validate() error {
	return utils.Validate(&r)
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxAmountAndType1) Validate() error {
	return utils.Validate(&r)
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

func (r TaxAmountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type TaxAuthorisation1 struct {
	Titl *common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
}

func (r TaxAuthorisation1) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation4 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord1                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation4) Validate() error {
	return utils.Validate(&r)
}

type TaxParty1 struct {
	TaxId  *common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId *common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  *common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

func (r TaxParty1) Validate() error {
	return utils.Validate(&r)
}

type TaxParty2 struct {
	TaxId   *common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  *common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   *common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

func (r TaxParty2) Validate() error {
	return utils.Validate(&r)
}

type TaxPeriod1 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriodDetails    `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod1) Validate() error {
	return utils.Validate(&r)
}

type TaxRecord1 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod1        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount1        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord1) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails1) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingTransaction13 struct {
	OrgnlGrpInfAndCxl *OriginalGroupHeader4  `xml:"OrgnlGrpInfAndCxl,omitempty" json:",omitempty"`
	TxInf             []PaymentTransaction62 `xml:"TxInf,omitempty" json:",omitempty"`
}

func (r UnderlyingTransaction13) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTimeChoice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

func (r DateAndDateTimeChoice) Validate() error {
	return utils.Validate(&r)
}

type MissingOrIncorrectInformation3 struct {
	AMLReq     bool                      `xml:"AMLReq,omitempty" json:",omitempty"`
	MssngInf   []UnableToApplyMissing1   `xml:"MssngInf,omitempty" json:",omitempty"`
	IncrrctInf []UnableToApplyIncorrect1 `xml:"IncrrctInf,omitempty" json:",omitempty"`
}

func (r MissingOrIncorrectInformation3) Validate() error {
	return utils.Validate(&r)
}

type UnableToApplyIncorrect1 struct {
	Cd              UnableToApplyIncorrectInformation4Code `xml:"Cd"`
	AddtlIncrrctInf *common.Max140Text                     `xml:"AddtlIncrrctInf,omitempty" json:",omitempty"`
}

func (r UnableToApplyIncorrect1) Validate() error {
	return utils.Validate(&r)
}

type UnableToApplyJustification3Choice struct {
	AnyInf            bool                           `xml:"AnyInf"`
	MssngOrIncrrctInf MissingOrIncorrectInformation3 `xml:"MssngOrIncrrctInf"`
	PssblDplctInstr   bool                           `xml:"PssblDplctInstr"`
}

func (r UnableToApplyJustification3Choice) Validate() error {
	return utils.Validate(&r)
}

type UnableToApplyMissing1 struct {
	Cd            UnableToApplyMissingInformation3Code `xml:"Cd"`
	AddtlMssngInf *common.Max140Text                   `xml:"AddtlMssngInf,omitempty" json:",omitempty"`
}

func (r UnableToApplyMissing1) Validate() error {
	return utils.Validate(&r)
}

type UnableToApplyV05 struct {
	XMLName     xml.Name                          `xml:"UblToApply"`
	Assgnmt     CaseAssignment3                   `xml:"Assgnmt"`
	Case        Case3                             `xml:"Case"`
	Undrlyg     UnderlyingTransaction3Choice      `xml:"Undrlyg"`
	Justfn      UnableToApplyJustification3Choice `xml:"Justfn"`
	SplmtryData []SupplementaryData1              `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r UnableToApplyV05) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId         common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId       common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm       *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlMsgDlvryChanl *common.Max35Text   `xml:"OrgnlMsgDlvryChanl,omitempty" json:",omitempty"`
}

func (r UnderlyingGroupInformation1) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingPaymentInstruction3 struct {
	OrgnlGrpInf     *UnderlyingGroupInformation1      `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlPmtInfId   *common.Max35Text                 `xml:"OrgnlPmtInfId,omitempty" json:",omitempty"`
	OrgnlInstrId    *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`
	ReqdExctnDt     *DateAndDateTimeChoice            `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt    *common.ISODate                   `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
}

func (r UnderlyingPaymentInstruction3) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingPaymentTransaction2 struct {
	OrgnlGrpInf         *UnderlyingGroupInformation1      `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                 `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  common.ISODate                    `xml:"OrgnlIntrBkSttlmDt"`
}

func (r UnderlyingPaymentTransaction2) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingStatementEntry1 struct {
	OrgnlGrpInf *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlStmtId *common.Max35Text          `xml:"OrgnlStmtId,omitempty" json:",omitempty"`
	OrgnlNtryId *common.Max35Text          `xml:"OrgnlNtryId,omitempty" json:",omitempty"`
}

func (r UnderlyingStatementEntry1) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingTransaction3Choice struct {
	Initn    UnderlyingPaymentInstruction3 `xml:"Initn"`
	IntrBk   UnderlyingPaymentTransaction2 `xml:"IntrBk"`
	StmtNtry UnderlyingStatementEntry1     `xml:"StmtNtry"`
}

func (r UnderlyingTransaction3Choice) Validate() error {
	return utils.Validate(&r)
}

type AdditionalPaymentInformationV05 struct {
	XMLName     xml.Name                         `xml:"AddtlPmtInf"`
	Assgnmt     CaseAssignment3                  `xml:"Assgnmt"`
	Case        Case3                            `xml:"Case"`
	Undrlyg     UnderlyingTransaction2Choice     `xml:"Undrlyg"`
	Inf         PaymentComplementaryInformation4 `xml:"Inf"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AdditionalPaymentInformationV05) Validate() error {
	return utils.Validate(&r)
}

type InstructionForCreditorAgent1 struct {
	Cd       *Instruction3Code  `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForCreditorAgent1) Validate() error {
	return utils.Validate(&r)
}

type InstructionForNextAgent1 struct {
	Cd       *Instruction4Code  `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForNextAgent1) Validate() error {
	return utils.Validate(&r)
}

type PaymentComplementaryInformation4 struct {
	InstrId          *common.Max35Text                             `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId       *common.Max35Text                             `xml:"EndToEndId,omitempty" json:",omitempty"`
	TxId             *common.Max35Text                             `xml:"TxId,omitempty" json:",omitempty"`
	PmtTpInf         *PaymentTypeInformation25                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt      *common.ISODate                               `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt     *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	IntrBkSttlmDt    *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	Amt              *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmAmt   *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	ChrgBr           *ChargeBearerType1Code                        `xml:"ChrgBr,omitempty" json:",omitempty"`
	UltmtDbtr        *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr             *PartyIdentification43                        `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct         *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt          *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct      *CashAccount24                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	SttlmInf         *SettlementInstruction1                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	IntrmyAgt1       *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt1Acct   *CashAccount24                                `xml:"IntrmyAgt1Acct,omitempty" json:",omitempty"`
	IntrmyAgt2       *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt2Acct   *CashAccount24                                `xml:"IntrmyAgt2Acct,omitempty" json:",omitempty"`
	IntrmyAgt3       *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	IntrmyAgt3Acct   *CashAccount24                                `xml:"IntrmyAgt3Acct,omitempty" json:",omitempty"`
	CdtrAgt          *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct      *CashAccount24                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr             *PartyIdentification43                        `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct         *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr        *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp             *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	InstrForDbtrAgt  *common.Max140Text                            `xml:"InstrForDbtrAgt,omitempty" json:",omitempty"`
	PrvsInstgAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt,omitempty" json:",omitempty"`
	PrvsInstgAgtAcct *CashAccount24                                `xml:"PrvsInstgAgtAcct,omitempty" json:",omitempty"`
	InstrForNxtAgt   []InstructionForNextAgent1                    `xml:"InstrForNxtAgt,omitempty" json:",omitempty"`
	InstrForCdtrAgt  []InstructionForCreditorAgent1                `xml:"InstrForCdtrAgt,omitempty" json:",omitempty"`
	RmtInf           *RemittanceInformation10                      `xml:"RmtInf,omitempty" json:",omitempty"`
}

func (r PaymentComplementaryInformation4) Validate() error {
	return utils.Validate(&r)
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r Purpose2Choice) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentInformation6 struct {
	Tp     *ReferredDocumentType4 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     *common.Max35Text      `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt *common.ISODate        `xml:"RltdDt,omitempty" json:",omitempty"`
}

func (r ReferredDocumentInformation6) Validate() error {
	return utils.Validate(&r)
}

type RemittanceInformation10 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation12 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation10) Validate() error {
	return utils.Validate(&r)
}

type SettlementInstruction1 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount24                                `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount24                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount24                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount24                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

func (r SettlementInstruction1) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation12 struct {
	RfrdDocInf  []ReferredDocumentInformation6 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification43         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification43         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation4               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment1                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation12) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingPaymentInstruction2 struct {
	OrgnlGrpInf     *UnderlyingGroupInformation1      `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlPmtInfId   *common.Max35Text                 `xml:"OrgnlPmtInfId,omitempty" json:",omitempty"`
	OrgnlInstrId    *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`
	ReqdExctnDt     *common.ISODate                   `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt    *common.ISODate                   `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
}

func (r UnderlyingPaymentInstruction2) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingTransaction2Choice struct {
	Initn    UnderlyingPaymentInstruction2 `xml:"Initn"`
	IntrBk   UnderlyingPaymentTransaction2 `xml:"IntrBk"`
	StmtNtry UnderlyingStatementEntry1     `xml:"StmtNtry"`
}

func (r UnderlyingTransaction2Choice) Validate() error {
	return utils.Validate(&r)
}
