// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

import "github.com/moov-io/iso20022/pkg/common"

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text  `xml:"Id"`
	SchmeNm *common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type GetMemberV04 struct {
	MsgHdr      MessageHeader9          `xml:"MsgHdr"`
	MmbQryDef   *MemberQueryDefinition4 `xml:"MmbQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MemberCriteria4 struct {
	NewQryNm *common.Max35Text       `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []MemberSearchCriteria4 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *MemberReturnCriteria1  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type MemberCriteriaDefinition2Choice struct {
	QryNm   common.Max35Text `xml:"QryNm"`
	NewCrit MemberCriteria4  `xml:"NewCrit"`
}

type MemberIdentification3Choice struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Othr        GenericFinancialIdentification1     `xml:"Othr"`
}

type MemberQueryDefinition4 struct {
	QryTp   *QueryType2Code                  `xml:"QryTp,omitempty" json:",omitempty"`
	MmbCrit *MemberCriteriaDefinition2Choice `xml:"MmbCrit,omitempty" json:",omitempty"`
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

type MemberSearchCriteria4 struct {
	Id  []MemberIdentification3Choice `xml:"Id,omitempty" json:",omitempty"`
	Tp  []SystemMemberType1Choice     `xml:"Tp,omitempty" json:",omitempty"`
	Sts []SystemMemberStatus1Choice   `xml:"Sts,omitempty" json:",omitempty"`
}

type MessageHeader9 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   *RequestType4Choice `xml:"ReqTp,omitempty" json:",omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemMemberStatus1Choice struct {
	Cd    MemberStatus1Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type SystemMemberType1Choice struct {
	Cd    ExternalSystemMemberType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice         `xml:"Id"`
	Tp   *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type CommunicationAddress10 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"PstlAdr"`
	PhneNb   common.PhoneNumber       `xml:"PhneNb"`
	FaxNb    *common.PhoneNumber      `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max2048Text      `xml:"EmailAdr,omitempty" json:",omitempty"`
}

type ContactIdentificationAndAddress2 struct {
	Nm     *common.Max35Text      `xml:"Nm,omitempty" json:",omitempty"`
	Role   PaymentRole1Choice     `xml:"Role"`
	ComAdr CommunicationAddress10 `xml:"ComAdr"`
}

type ErrorHandling1Choice struct {
	Cd    ErrorHandling1Code          `xml:"Cd"`
	Prtry common.Max4AlphaNumericText `xml:"Prtry"`
}

type ErrorHandling3 struct {
	Err  ErrorHandling1Choice `xml:"Err"`
	Desc *common.Max140Text   `xml:"Desc,omitempty" json:",omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
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

type MemberReport5 struct {
	MmbId    MemberIdentification3Choice `xml:"MmbId"`
	MmbOrErr MemberReportOrError6Choice  `xml:"MmbOrErr"`
}

type MemberReportOrError5Choice struct {
	Rpt     []MemberReport5  `xml:"Rpt"`
	OprlErr []ErrorHandling3 `xml:"OprlErr"`
}

type MemberReportOrError6Choice struct {
	Mmb    Member5        `xml:"Mmb"`
	BizErr ErrorHandling3 `xml:"BizErr"`
}

type MessageHeader7 struct {
	MsgId       common.Max35Text        `xml:"MsgId"`
	CreDtTm     *common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp       *RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       *common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId *common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type PaymentRole1Choice struct {
	Cd    ExternalPaymentRole1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type ReturnMemberV04 struct {
	MsgHdr      MessageHeader7             `xml:"MsgHdr"`
	RptOrErr    MemberReportOrError5Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type CommunicationAddress8 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"PstlAdr"`
	PhneNb   common.PhoneNumber       `xml:"PhneNb"`
	FaxNb    *common.PhoneNumber      `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max256Text       `xml:"EmailAdr,omitempty" json:",omitempty"`
}

type ContactIdentificationAndAddress1 struct {
	Nm     *common.Max35Text     `xml:"Nm,omitempty" json:",omitempty"`
	Role   PaymentRole1Code      `xml:"Role"`
	ComAdr CommunicationAddress8 `xml:"ComAdr"`
}

type LongPostalAddress1Choice struct {
	Ustrd common.Max140Text            `xml:"Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"Strd"`
}

type Member6 struct {
	MmbRtrAdr []MemberIdentification3Choice      `xml:"MmbRtrAdr,omitempty" json:",omitempty"`
	CtctRef   []ContactIdentificationAndAddress1 `xml:"CtctRef,omitempty" json:",omitempty"`
	ComAdr    *CommunicationAddress8             `xml:"ComAdr,omitempty" json:",omitempty"`
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type ModifyMemberV04 struct {
	MsgHdr       MessageHeader1              `xml:"MsgHdr"`
	MmbId        MemberIdentification3Choice `xml:"MmbId"`
	NewMmbValSet Member6                     `xml:"NewMmbValSet"`
	SplmtryData  []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
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

type CurrencyCriteriaDefinition1Choice struct {
	QryNm   common.Max35Text          `xml:"QryNm"`
	NewCrit CurrencyExchangeCriteria2 `xml:"NewCrit"`
}

type CurrencyExchangeCriteria2 struct {
	NewQryNm *common.Max35Text                 `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []CurrencyExchangeSearchCriteria1 `xml:"SchCrit"`
}

type CurrencyExchangeSearchCriteria1 struct {
	SrcCcy  common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`
}

type CurrencyQueryDefinition3 struct {
	QryTp   *QueryType2Code                    `xml:"QryTp,omitempty" json:",omitempty"`
	CcyCrit *CurrencyCriteriaDefinition1Choice `xml:"CcyCrit,omitempty" json:",omitempty"`
}

type GetCurrencyExchangeRateV04 struct {
	MsgHdr      MessageHeader1            `xml:"MsgHdr"`
	CcyQryDef   *CurrencyQueryDefinition3 `xml:"CcyQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type CurrencyExchange7 struct {
	XchgRate float64                             `xml:"XchgRate"`
	QtdCcy   common.ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`
	QtnDt    common.ISODateTime                  `xml:"QtnDt"`
}

type CurrencyExchangeReport3 struct {
	CcyRef       CurrencySourceTarget1            `xml:"CcyRef"`
	CcyXchgOrErr ExchangeRateReportOrError2Choice `xml:"CcyXchgOrErr"`
}

type CurrencySourceTarget1 struct {
	SrcCcy  common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`
}

type ExchangeRateReportOrError1Choice struct {
	CcyXchgRpt []CurrencyExchangeReport3 `xml:"CcyXchgRpt"`
	OprlErr    []ErrorHandling3          `xml:"OprlErr"`
}

type ExchangeRateReportOrError2Choice struct {
	BizErr  []ErrorHandling3  `xml:"BizErr"`
	CcyXchg CurrencyExchange7 `xml:"CcyXchg"`
}

type ReturnCurrencyExchangeRateV04 struct {
	MsgHdr      MessageHeader7                   `xml:"MsgHdr"`
	RptOrErr    ExchangeRateReportOrError1Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type BusinessInformationCriteria1 struct {
	NewQryNm *common.Max35Text                           `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []GeneralBusinessInformationSearchCriteria1 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *GeneralBusinessInformationReturnCriteria1  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type BusinessInformationQueryDefinition3 struct {
	QryTp         *QueryType2Code                                      `xml:"QryTp,omitempty" json:",omitempty"`
	GnlBizInfCrit *GeneralBusinessInformationCriteriaDefinition1Choice `xml:"GnlBizInfCrit,omitempty" json:",omitempty"`
}

type CharacterSearch1Choice struct {
	EQ  common.Max35Text `xml:"EQ"`
	NEQ common.Max35Text `xml:"NEQ"`
	CT  common.Max35Text `xml:"CT"`
	NCT common.Max35Text `xml:"NCT"`
}

type GeneralBusinessInformationCriteriaDefinition1Choice struct {
	QryNm   common.Max35Text             `xml:"QryNm"`
	NewCrit BusinessInformationCriteria1 `xml:"NewCrit"`
}

type GeneralBusinessInformationReturnCriteria1 struct {
	QlfrInd     bool `xml:"QlfrInd,omitempty" json:",omitempty"`
	SbjtInd     bool `xml:"SbjtInd,omitempty" json:",omitempty"`
	SbjtDtlsInd bool `xml:"SbjtDtlsInd,omitempty" json:",omitempty"`
}

type GeneralBusinessInformationSearchCriteria1 struct {
	Ref  []common.Max35Text          `xml:"Ref,omitempty" json:",omitempty"`
	Sbjt []CharacterSearch1Choice    `xml:"Sbjt,omitempty" json:",omitempty"`
	Qlfr []InformationQualifierType1 `xml:"Qlfr,omitempty" json:",omitempty"`
}

type GetGeneralBusinessInformationV04 struct {
	MsgHdr          MessageHeader1                       `xml:"MsgHdr"`
	GnlBizInfQryDef *BusinessInformationQueryDefinition3 `xml:"GnlBizInfQryDef,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type InformationQualifierType1 struct {
	IsFrmtd bool           `xml:"IsFrmtd,omitempty" json:",omitempty"`
	Prty    *Priority1Code `xml:"Prty,omitempty" json:",omitempty"`
}

type CancelCaseAssignmentV04 struct {
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        Case5                `xml:"Case"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData3 struct {
	Id      *common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

type CaseStatusReportRequestV04 struct {
	ReqHdr      ReportHeader5        `xml:"ReqHdr"`
	Case        Case5                `xml:"Case"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
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

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text       `xml:"Issr"`
	SchmeNm *common.Max35Text      `xml:"SchmeNm,omitempty" json:",omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `xml:"ChanlTp"`
	Id      *common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
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

type ReportHeader5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Fr      Party40Choice      `xml:"Fr"`
	To      Party40Choice      `xml:"To"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate  `xml:"FrDt"`
	ToDt *common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"Err"`
	Desc *common.Max140Text   `xml:"Desc,omitempty" json:",omitempty"`
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type ExecutionType1Choice struct {
	Tm  common.ISOTime   `xml:"Tm"`
	Evt EventType1Choice `xml:"Evt"`
}

type MessageHeader6 struct {
	MsgId       common.Max35Text        `xml:"MsgId"`
	CreDtTm     *common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       *common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
	ReqTp       *RequestType3Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
}

type RequestType3Choice struct {
	Cd    StandingOrderQueryType1Code `xml:"Cd"`
	Prtry GenericIdentification1      `xml:"Prtry"`
}

type ReturnStandingOrderV04 struct {
	MsgHdr      MessageHeader6              `xml:"MsgHdr"`
	RptOrErr    StandingOrderOrError5Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
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

type StandingOrderIdentification4 struct {
	Id       *common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	Acct     CashAccount38                                 `xml:"Acct"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

type StandingOrderOrError5Choice struct {
	Rpt     []StandingOrderReport1 `xml:"Rpt"`
	OprlErr []ErrorHandling5       `xml:"OprlErr"`
}

type StandingOrderOrError6Choice struct {
	StgOrdr StandingOrder6   `xml:"StgOrdr"`
	BizErr  []ErrorHandling5 `xml:"BizErr"`
}

type StandingOrderReport1 struct {
	StgOrdrId    StandingOrderIdentification4 `xml:"StgOrdrId"`
	StgOrdrOrErr StandingOrderOrError6Choice  `xml:"StgOrdrOrErr"`
}

type StandingOrderTotalAmount1 struct {
	SetPrdfndOrdr TotalAmountAndCurrency1 `xml:"SetPrdfndOrdr"`
	PdgPrdfndOrdr TotalAmountAndCurrency1 `xml:"PdgPrdfndOrdr"`
	SetStgOrdr    TotalAmountAndCurrency1 `xml:"SetStgOrdr"`
	PdgStgOrdr    TotalAmountAndCurrency1 `xml:"PdgStgOrdr"`
}

type StandingOrderType1Choice struct {
	Cd    StandingOrderType1Code `xml:"Cd"`
	Prtry GenericIdentification1 `xml:"Prtry"`
}

type TotalAmountAndCurrency1 struct {
	TtlAmt    float64                    `xml:"TtlAmt"`
	CdtDbtInd *common.CreditDebitCode    `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Ccy       *common.ActiveCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
}
