// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

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

type GetMemberV04 struct {
	MsgHdr      MessageHeader9          `xml:"MsgHdr"`
	MmbQryDef   *MemberQueryDefinition4 `xml:"MmbQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r GetMemberV04) Validate() error {
	return utils.Validate(&r)
}

type MemberCriteria4 struct {
	NewQryNm *common.Max35Text       `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []MemberSearchCriteria4 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *MemberReturnCriteria1  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

func (r MemberCriteria4) Validate() error {
	return utils.Validate(&r)
}

type MemberCriteriaDefinition2Choice struct {
	QryNm   common.Max35Text `xml:"QryNm"`
	NewCrit MemberCriteria4  `xml:"NewCrit"`
}

func (r MemberCriteriaDefinition2Choice) Validate() error {
	return utils.Validate(&r)
}

type MemberIdentification3Choice struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Othr        GenericFinancialIdentification1     `xml:"Othr"`
}

func (r MemberIdentification3Choice) Validate() error {
	return utils.Validate(&r)
}

type MemberQueryDefinition4 struct {
	QryTp   *QueryType2Code                  `xml:"QryTp,omitempty" json:",omitempty"`
	MmbCrit *MemberCriteriaDefinition2Choice `xml:"MmbCrit,omitempty" json:",omitempty"`
}

func (r MemberQueryDefinition4) Validate() error {
	return utils.Validate(&r)
}

type MemberReturnCriteria1 struct {
	NmInd        bool `xml:"NmInd,omitempty" json:",omitempty"`
	MmbRtrAdrInd bool `xml:"MmbRtrAdrInd,omitempty" json:",omitempty"`
	AcctInd      bool `xml:"AcctInd,omitempty" json:",omitempty"`
	TpInd        bool `xml:"TpInd,omitempty" json:",omitempty"`
	StsInd       bool `xml:"StsInd,omitempty" json:",omitempty"`
	CtctRefInd   bool `xml:"CtctRefInd,omitempty" json:",omitempty"`
	ComAdrInd    bool `xml:"ComAdrInd,omitempty" json:",omitempty"`
}

func (r MemberReturnCriteria1) Validate() error {
	return utils.Validate(&r)
}

type MemberSearchCriteria4 struct {
	Id  []MemberIdentification3Choice `xml:"Id,omitempty" json:",omitempty"`
	Tp  []SystemMemberType1Choice     `xml:"Tp,omitempty" json:",omitempty"`
	Sts []SystemMemberStatus1Choice   `xml:"Sts,omitempty" json:",omitempty"`
}

func (r MemberSearchCriteria4) Validate() error {
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

type SystemMemberStatus1Choice struct {
	Cd    MemberStatus1Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

func (r SystemMemberStatus1Choice) Validate() error {
	return utils.Validate(&r)
}

type SystemMemberType1Choice struct {
	Cd    ExternalSystemMemberType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

func (r SystemMemberType1Choice) Validate() error {
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

type CommunicationAddress10 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"PstlAdr"`
	PhneNb   common.PhoneNumber       `xml:"PhneNb"`
	FaxNb    *common.PhoneNumber      `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max2048Text      `xml:"EmailAdr,omitempty" json:",omitempty"`
}

func (r CommunicationAddress10) Validate() error {
	return utils.Validate(&r)
}

type ContactIdentificationAndAddress2 struct {
	Nm     *common.Max35Text      `xml:"Nm,omitempty" json:",omitempty"`
	Role   PaymentRole1Choice     `xml:"Role"`
	ComAdr CommunicationAddress10 `xml:"ComAdr"`
}

func (r ContactIdentificationAndAddress2) Validate() error {
	return utils.Validate(&r)
}

type ErrorHandling1Choice struct {
	Cd    ErrorHandling1Code          `xml:"Cd"`
	Prtry common.Max4AlphaNumericText `xml:"Prtry"`
}

func (r ErrorHandling1Choice) Validate() error {
	return utils.Validate(&r)
}

type ErrorHandling3 struct {
	Err  ErrorHandling1Choice `xml:"Err"`
	Desc *common.Max140Text   `xml:"Desc,omitempty" json:",omitempty"`
}

func (r ErrorHandling3) Validate() error {
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

type Member5 struct {
	Nm      *common.Max35Text                  `xml:"Nm,omitempty" json:",omitempty"`
	RtrAdr  []MemberIdentification3Choice      `xml:"RtrAdr,omitempty" json:",omitempty"`
	Acct    []CashAccount38                    `xml:"Acct,omitempty" json:",omitempty"`
	Tp      *SystemMemberType1Choice           `xml:"Tp,omitempty" json:",omitempty"`
	Sts     *SystemMemberStatus1Choice         `xml:"Sts,omitempty" json:",omitempty"`
	CtctRef []ContactIdentificationAndAddress2 `xml:"CtctRef,omitempty" json:",omitempty"`
	ComAdr  CommunicationAddress10             `xml:"ComAdr,omitempty" json:",omitempty"`
}

func (r Member5) Validate() error {
	return utils.Validate(&r)
}

type MemberReport5 struct {
	MmbId    MemberIdentification3Choice `xml:"MmbId"`
	MmbOrErr MemberReportOrError6Choice  `xml:"MmbOrErr"`
}

func (r MemberReport5) Validate() error {
	return utils.Validate(&r)
}

type MemberReportOrError5Choice struct {
	Rpt     []MemberReport5  `xml:"Rpt" json:",omitempty"`
	OprlErr []ErrorHandling3 `xml:"OprlErr" json:",omitempty"`
}

func (r MemberReportOrError5Choice) Validate() error {
	return utils.Validate(&r)
}

type MemberReportOrError6Choice struct {
	Mmb    Member5        `xml:"Mmb"`
	BizErr ErrorHandling3 `xml:"BizErr"`
}

func (r MemberReportOrError6Choice) Validate() error {
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

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId *common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalBusinessQuery1) Validate() error {
	return utils.Validate(&r)
}

type PaymentRole1Choice struct {
	Cd    ExternalPaymentRole1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r PaymentRole1Choice) Validate() error {
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

type ReturnMemberV04 struct {
	MsgHdr      MessageHeader7             `xml:"MsgHdr"`
	RptOrErr    MemberReportOrError5Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReturnMemberV04) Validate() error {
	return utils.Validate(&r)
}

type CommunicationAddress8 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"PstlAdr"`
	PhneNb   common.PhoneNumber       `xml:"PhneNb"`
	FaxNb    *common.PhoneNumber      `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max256Text       `xml:"EmailAdr,omitempty" json:",omitempty"`
}

func (r CommunicationAddress8) Validate() error {
	return utils.Validate(&r)
}

type ContactIdentificationAndAddress1 struct {
	Nm     *common.Max35Text     `xml:"Nm,omitempty" json:",omitempty"`
	Role   PaymentRole1Code      `xml:"Role"`
	ComAdr CommunicationAddress8 `xml:"ComAdr"`
}

func (r ContactIdentificationAndAddress1) Validate() error {
	return utils.Validate(&r)
}

type LongPostalAddress1Choice struct {
	Ustrd common.Max140Text            `xml:"Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"Strd"`
}

func (r LongPostalAddress1Choice) Validate() error {
	return utils.Validate(&r)
}

type Member6 struct {
	MmbRtrAdr []MemberIdentification3Choice      `xml:"MmbRtrAdr,omitempty" json:",omitempty"`
	CtctRef   []ContactIdentificationAndAddress1 `xml:"CtctRef,omitempty" json:",omitempty"`
	ComAdr    *CommunicationAddress8             `xml:"ComAdr,omitempty" json:",omitempty"`
}

func (r Member6) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r MessageHeader1) Validate() error {
	return utils.Validate(&r)
}

type ModifyMemberV04 struct {
	MsgHdr       MessageHeader1              `xml:"MsgHdr"`
	MmbId        MemberIdentification3Choice `xml:"MmbId"`
	NewMmbValSet Member6                     `xml:"NewMmbValSet"`
	SplmtryData  []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ModifyMemberV04) Validate() error {
	return utils.Validate(&r)
}

type StructuredLongPostalAddress1 struct {
	BldgNm     *common.Max35Text  `xml:"BldgNm,omitempty" json:",omitempty"`
	StrtNm     *common.Max35Text  `xml:"StrtNm,omitempty" json:",omitempty"`
	StrtBldgId *common.Max35Text  `xml:"StrtBldgId,omitempty" json:",omitempty"`
	Flr        *common.Max16Text  `xml:"Flr,omitempty" json:",omitempty"`
	TwnNm      common.Max35Text   `xml:"TwnNm"`
	DstrctNm   *common.Max35Text  `xml:"DstrctNm,omitempty" json:",omitempty"`
	RgnId      *common.Max35Text  `xml:"RgnId,omitempty" json:",omitempty"`
	Stat       *common.Max35Text  `xml:"Stat,omitempty" json:",omitempty"`
	CtyId      common.Max35Text   `xml:"CtyId,omitempty" json:",omitempty"`
	Ctry       common.CountryCode `xml:"Ctry"`
	PstCdId    common.Max16Text   `xml:"PstCdId"`
	POB        *common.Max16Text  `xml:"POB,omitempty" json:",omitempty"`
}

func (r StructuredLongPostalAddress1) Validate() error {
	return utils.Validate(&r)
}

type CurrencyCriteriaDefinition1Choice struct {
	QryNm   common.Max35Text          `xml:"QryNm"`
	NewCrit CurrencyExchangeCriteria2 `xml:"NewCrit"`
}

func (r CurrencyCriteriaDefinition1Choice) Validate() error {
	return utils.Validate(&r)
}

type CurrencyExchangeCriteria2 struct {
	NewQryNm *common.Max35Text                 `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []CurrencyExchangeSearchCriteria1 `xml:"SchCrit" json:",omitempty"`
}

func (r CurrencyExchangeCriteria2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyExchangeSearchCriteria1 struct {
	SrcCcy  common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`
}

func (r CurrencyExchangeSearchCriteria1) Validate() error {
	return utils.Validate(&r)
}

type CurrencyQueryDefinition3 struct {
	QryTp   *QueryType2Code                    `xml:"QryTp,omitempty" json:",omitempty"`
	CcyCrit *CurrencyCriteriaDefinition1Choice `xml:"CcyCrit,omitempty" json:",omitempty"`
}

func (r CurrencyQueryDefinition3) Validate() error {
	return utils.Validate(&r)
}

type GetCurrencyExchangeRateV04 struct {
	MsgHdr      MessageHeader1            `xml:"MsgHdr"`
	CcyQryDef   *CurrencyQueryDefinition3 `xml:"CcyQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r GetCurrencyExchangeRateV04) Validate() error {
	return utils.Validate(&r)
}

type CurrencyExchange7 struct {
	XchgRate float64                             `xml:"XchgRate"`
	QtdCcy   common.ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`
	QtnDt    common.ISODateTime                  `xml:"QtnDt"`
}

func (r CurrencyExchange7) Validate() error {
	return utils.Validate(&r)
}

type CurrencyExchangeReport3 struct {
	CcyRef       CurrencySourceTarget1            `xml:"CcyRef"`
	CcyXchgOrErr ExchangeRateReportOrError2Choice `xml:"CcyXchgOrErr"`
}

func (r CurrencyExchangeReport3) Validate() error {
	return utils.Validate(&r)
}

type CurrencySourceTarget1 struct {
	SrcCcy  common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`
}

func (r CurrencySourceTarget1) Validate() error {
	return utils.Validate(&r)
}

type ExchangeRateReportOrError1Choice struct {
	CcyXchgRpt []CurrencyExchangeReport3 `xml:"CcyXchgRpt" json:",omitempty"`
	OprlErr    []ErrorHandling3          `xml:"OprlErr" json:",omitempty"`
}

func (r ExchangeRateReportOrError1Choice) Validate() error {
	return utils.Validate(&r)
}

type ExchangeRateReportOrError2Choice struct {
	BizErr  []ErrorHandling3  `xml:"BizErr" json:",omitempty"`
	CcyXchg CurrencyExchange7 `xml:"CcyXchg"`
}

func (r ExchangeRateReportOrError2Choice) Validate() error {
	return utils.Validate(&r)
}

type ReturnCurrencyExchangeRateV04 struct {
	MsgHdr      MessageHeader7                   `xml:"MsgHdr"`
	RptOrErr    ExchangeRateReportOrError1Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReturnCurrencyExchangeRateV04) Validate() error {
	return utils.Validate(&r)
}

type BusinessInformationCriteria1 struct {
	NewQryNm *common.Max35Text                           `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []GeneralBusinessInformationSearchCriteria1 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *GeneralBusinessInformationReturnCriteria1  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

func (r BusinessInformationCriteria1) Validate() error {
	return utils.Validate(&r)
}

type BusinessInformationQueryDefinition3 struct {
	QryTp         *QueryType2Code                                      `xml:"QryTp,omitempty" json:",omitempty"`
	GnlBizInfCrit *GeneralBusinessInformationCriteriaDefinition1Choice `xml:"GnlBizInfCrit,omitempty" json:",omitempty"`
}

func (r BusinessInformationQueryDefinition3) Validate() error {
	return utils.Validate(&r)
}

type CharacterSearch1Choice struct {
	EQ  common.Max35Text `xml:"EQ"`
	NEQ common.Max35Text `xml:"NEQ"`
	CT  common.Max35Text `xml:"CT"`
	NCT common.Max35Text `xml:"NCT"`
}

func (r CharacterSearch1Choice) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessInformationCriteriaDefinition1Choice struct {
	QryNm   common.Max35Text             `xml:"QryNm"`
	NewCrit BusinessInformationCriteria1 `xml:"NewCrit"`
}

func (r GeneralBusinessInformationCriteriaDefinition1Choice) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessInformationReturnCriteria1 struct {
	QlfrInd     bool `xml:"QlfrInd,omitempty" json:",omitempty"`
	SbjtInd     bool `xml:"SbjtInd,omitempty" json:",omitempty"`
	SbjtDtlsInd bool `xml:"SbjtDtlsInd,omitempty" json:",omitempty"`
}

func (r GeneralBusinessInformationReturnCriteria1) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessInformationSearchCriteria1 struct {
	Ref  []common.Max35Text          `xml:"Ref,omitempty" json:",omitempty"`
	Sbjt []CharacterSearch1Choice    `xml:"Sbjt,omitempty" json:",omitempty"`
	Qlfr []InformationQualifierType1 `xml:"Qlfr,omitempty" json:",omitempty"`
}

func (r GeneralBusinessInformationSearchCriteria1) Validate() error {
	return utils.Validate(&r)
}

type GetGeneralBusinessInformationV04 struct {
	MsgHdr          MessageHeader1                       `xml:"MsgHdr"`
	GnlBizInfQryDef *BusinessInformationQueryDefinition3 `xml:"GnlBizInfQryDef,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r GetGeneralBusinessInformationV04) Validate() error {
	return utils.Validate(&r)
}

type InformationQualifierType1 struct {
	IsFrmtd bool           `xml:"IsFrmtd,omitempty" json:",omitempty"`
	Prty    *Priority1Code `xml:"Prty,omitempty" json:",omitempty"`
}

func (r InformationQualifierType1) Validate() error {
	return utils.Validate(&r)
}

type CancelCaseAssignmentV04 struct {
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        Case5                `xml:"Case"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CancelCaseAssignmentV04) Validate() error {
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

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

func (r Case5) Validate() error {
	return utils.Validate(&r)
}

type CaseStatusReportRequestV04 struct {
	ReqHdr      ReportHeader5        `xml:"ReqHdr"`
	Case        Case5                `xml:"Case"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CaseStatusReportRequestV04) Validate() error {
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

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
	return utils.Validate(&r)
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
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

type ReportHeader5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Fr      Party40Choice      `xml:"Fr"`
	To      Party40Choice      `xml:"To"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r ReportHeader5) Validate() error {
	return utils.Validate(&r)
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

func (r Amount2Choice) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate  `xml:"FrDt"`
	ToDt *common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

func (r DatePeriodDetails1) Validate() error {
	return utils.Validate(&r)
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

func (r ErrorHandling3Choice) Validate() error {
	return utils.Validate(&r)
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"Err"`
	Desc *common.Max140Text   `xml:"Desc,omitempty" json:",omitempty"`
}

func (r ErrorHandling5) Validate() error {
	return utils.Validate(&r)
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r EventType1Choice) Validate() error {
	return utils.Validate(&r)
}

type ExecutionType1Choice struct {
	Tm  common.ISOTime   `xml:"Tm"`
	Evt EventType1Choice `xml:"Evt"`
}

func (r ExecutionType1Choice) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader6 struct {
	MsgId       common.Max35Text        `xml:"MsgId"`
	CreDtTm     *common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       *common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
	ReqTp       *RequestType3Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
}

func (r MessageHeader6) Validate() error {
	return utils.Validate(&r)
}

type RequestType3Choice struct {
	Cd    StandingOrderQueryType1Code `xml:"Cd"`
	Prtry GenericIdentification1      `xml:"Prtry"`
}

func (r RequestType3Choice) Validate() error {
	return utils.Validate(&r)
}

type ReturnStandingOrderV04 struct {
	MsgHdr      MessageHeader6              `xml:"MsgHdr"`
	RptOrErr    StandingOrderOrError5Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReturnStandingOrderV04) Validate() error {
	return utils.Validate(&r)
}

type StandingOrder6 struct {
	Amt             Amount2Choice                                 `xml:"Amt"`
	CdtDbtInd       common.CreditDebitCode                        `xml:"CdtDbtInd"`
	Ccy             *common.ActiveCurrencyCode                    `xml:"Ccy,omitempty" json:",omitempty"`
	Tp              *StandingOrderType1Choice                     `xml:"Tp,omitempty" json:",omitempty"`
	AssoctdPoolAcct *AccountIdentification4Choice                 `xml:"AssoctdPoolAcct,omitempty" json:",omitempty"`
	Ref             *common.Max35Text                             `xml:"Ref,omitempty" json:",omitempty"`
	Frqcy           *Frequency2Code                               `xml:"Frqcy,omitempty" json:",omitempty"`
	VldtyPrd        *DatePeriodDetails1                           `xml:"VldtyPrd,omitempty" json:",omitempty"`
	SysMmb          *BranchAndFinancialInstitutionIdentification6 `xml:"SysMmb,omitempty" json:",omitempty"`
	RspnsblPty      *BranchAndFinancialInstitutionIdentification6 `xml:"RspnsblPty,omitempty" json:",omitempty"`
	LkSetId         *common.Max35Text                             `xml:"LkSetId,omitempty" json:",omitempty"`
	LkSetOrdrId     *common.Max35Text                             `xml:"LkSetOrdrId,omitempty" json:",omitempty"`
	LkSetOrdrSeq    float64                                       `xml:"LkSetOrdrSeq,omitempty" json:",omitempty"`
	ExctnTp         *ExecutionType1Choice                         `xml:"ExctnTp,omitempty" json:",omitempty"`
	Cdtr            *BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct        *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	Dbtr            *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct        *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	TtlsPerStgOrdr  *StandingOrderTotalAmount1                    `xml:"TtlsPerStgOrdr,omitempty" json:",omitempty"`
	ZeroSweepInd    bool                                          `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
}

func (r StandingOrder6) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderIdentification4 struct {
	Id       *common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	Acct     CashAccount38                                 `xml:"Acct"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

func (r StandingOrderIdentification4) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderOrError5Choice struct {
	Rpt     []StandingOrderReport1 `xml:"Rpt" json:",omitempty"`
	OprlErr []ErrorHandling5       `xml:"OprlErr" json:",omitempty"`
}

func (r StandingOrderOrError5Choice) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderOrError6Choice struct {
	StgOrdr StandingOrder6   `xml:"StgOrdr"`
	BizErr  []ErrorHandling5 `xml:"BizErr" json:",omitempty"`
}

func (r StandingOrderOrError6Choice) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderReport1 struct {
	StgOrdrId    StandingOrderIdentification4 `xml:"StgOrdrId"`
	StgOrdrOrErr StandingOrderOrError6Choice  `xml:"StgOrdrOrErr"`
}

func (r StandingOrderReport1) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderTotalAmount1 struct {
	SetPrdfndOrdr TotalAmountAndCurrency1 `xml:"SetPrdfndOrdr"`
	PdgPrdfndOrdr TotalAmountAndCurrency1 `xml:"PdgPrdfndOrdr"`
	SetStgOrdr    TotalAmountAndCurrency1 `xml:"SetStgOrdr"`
	PdgStgOrdr    TotalAmountAndCurrency1 `xml:"PdgStgOrdr"`
}

func (r StandingOrderTotalAmount1) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderType1Choice struct {
	Cd    StandingOrderType1Code `xml:"Cd"`
	Prtry GenericIdentification1 `xml:"Prtry"`
}

func (r StandingOrderType1Choice) Validate() error {
	return utils.Validate(&r)
}

type TotalAmountAndCurrency1 struct {
	TtlAmt    float64                    `xml:"TtlAmt"`
	CdtDbtInd *common.CreditDebitCode    `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Ccy       *common.ActiveCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
}

func (r TotalAmountAndCurrency1) Validate() error {
	return utils.Validate(&r)
}
