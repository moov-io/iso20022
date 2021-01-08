// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountOrBusinessError4Choice struct {
	Acct   CashAccount37    `xml:"Acct"`
	BizErr []ErrorHandling5 `xml:"BizErr"`
}

type AccountOrOperationalError4Choice struct {
	AcctRpt []AccountReport24 `xml:"AcctRpt"`
	OprlErr []ErrorHandling5  `xml:"OprlErr"`
}

type AccountReport24 struct {
	AcctId    AccountIdentification4Choice  `xml:"AcctId"`
	AcctOrErr AccountOrBusinessError4Choice `xml:"AcctOrErr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

type BalanceRestrictionType1 struct {
	Tp     GenericIdentification1 `xml:"Tp"`
	Desc   common.Max140Text      `xml:"Desc,omitempty" json:",omitempty"`
	PrcgTp ProcessingType1Choice  `xml:"PrcgTp,omitempty" json:",omitempty"`
}

type BalanceType11Choice struct {
	Cd    ExternalSystemBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text               `xml:"Prtry"`
}

type BalanceType9Choice struct {
	Cd    SystemBalanceType2Code `xml:"Cd"`
	Prtry common.Max35Text       `xml:"Prtry"`
}

type BilateralLimit3 struct {
	CtrPtyId  BranchAndFinancialInstitutionIdentification6 `xml:"CtrPtyId"`
	LmtAmt    Amount2Choice                                `xml:"LmtAmt"`
	CdtDbtInd common.CreditDebitCode                       `xml:"CdtDbtInd"`
	BilBal    []CashBalance11                              `xml:"BilBal,omitempty" json:",omitempty"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    BranchData3                          `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type CashAccount37 struct {
	Nm        common.Max70Text                             `xml:"Nm,omitempty" json:",omitempty"`
	Tp        CashAccountType2Choice                       `xml:"Tp,omitempty" json:",omitempty"`
	Ccy       common.ActiveOrHistoricCurrencyCode          `xml:"Ccy,omitempty" json:",omitempty"`
	Prxy      ProxyAccountIdentification1                  `xml:"Prxy,omitempty" json:",omitempty"`
	CurMulLmt Limit5                                       `xml:"CurMulLmt,omitempty" json:",omitempty"`
	Ownr      PartyIdentification135                       `xml:"Ownr,omitempty" json:",omitempty"`
	Svcr      BranchAndFinancialInstitutionIdentification6 `xml:"Svcr,omitempty" json:",omitempty"`
	MulBal    []CashBalance13                              `xml:"MulBal,omitempty" json:",omitempty"`
	CurBilLmt []BilateralLimit3                            `xml:"CurBilLmt,omitempty" json:",omitempty"`
	StgOrdr   []StandingOrder6                             `xml:"StgOrdr,omitempty" json:",omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"Id"`
	Tp   CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type CashBalance11 struct {
	Amt       float64                `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode `xml:"CdtDbtInd"`
	Tp        BalanceType9Choice     `xml:"Tp,omitempty" json:",omitempty"`
	Sts       BalanceStatus1Code     `xml:"Sts,omitempty" json:",omitempty"`
	ValDt     DateAndDateTime2Choice `xml:"ValDt,omitempty" json:",omitempty"`
	NbOfPmts  float64                `xml:"NbOfPmts,omitempty" json:",omitempty"`
}

type CashBalance13 struct {
	Amt       float64                 `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode  `xml:"CdtDbtInd"`
	Tp        BalanceType11Choice     `xml:"Tp,omitempty" json:",omitempty"`
	Sts       BalanceStatus1Code      `xml:"Sts,omitempty" json:",omitempty"`
	ValDt     DateAndDateTime2Choice  `xml:"ValDt,omitempty" json:",omitempty"`
	PrcgDt    DateAndDateTime2Choice  `xml:"PrcgDt,omitempty" json:",omitempty"`
	NbOfPmts  float64                 `xml:"NbOfPmts,omitempty" json:",omitempty"`
	RstrctnTp BalanceRestrictionType1 `xml:"RstrctnTp,omitempty" json:",omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                    `xml:"MmbId"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm        common.Max140Text           `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb    common.PhoneNumber          `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb     common.PhoneNumber          `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb     common.PhoneNumber          `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr  common.Max2048Text          `xml:"EmailAdr,omitempty" json:",omitempty"`
	EmailPurp common.Max35Text            `xml:"EmailPurp,omitempty" json:",omitempty"`
	JobTitl   common.Max35Text            `xml:"JobTitl,omitempty" json:",omitempty"`
	Rspnsblty common.Max35Text            `xml:"Rspnsblty,omitempty" json:",omitempty"`
	Dept      common.Max70Text            `xml:"Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1             `xml:"Othr,omitempty" json:",omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:",omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"Err"`
	Desc common.Max140Text    `xml:"Desc,omitempty" json:",omitempty"`
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type ExecutionType1Choice struct {
	Tm  common.ISOTime   `xml:"Tm"`
	Evt EventType1Choice `xml:"Evt"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"Id"`
	SchmeNm AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text `xml:"Id"`
	SchmeNm common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text       `xml:"Issr"`
	SchmeNm common.Max35Text       `xml:"SchmeNm,omitempty" json:",omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

type Limit5 struct {
	Amt       Amount2Choice          `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode `xml:"CdtDbtInd"`
}

type MessageHeader7 struct {
	MsgId       common.Max35Text       `xml:"MsgId"`
	CreDtTm     common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp       RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    common.LEIIdentifier                 `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	MsgNmId common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"ChanlTp"`
	Id      common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	BldgNm      common.Max35Text   `xml:"BldgNm,omitempty" json:",omitempty"`
	Flr         common.Max70Text   `xml:"Flr,omitempty" json:",omitempty"`
	PstBx       common.Max16Text   `xml:"PstBx,omitempty" json:",omitempty"`
	Room        common.Max70Text   `xml:"Room,omitempty" json:",omitempty"`
	PstCd       common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    common.Max35Text   `xml:"DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text `xml:"AdrLine,omitempty" json:",omitempty"`
}

type ProcessingType1Choice struct {
	Cd    ProcessingType1Code `xml:"Cd"`
	Prtry common.Max35Text    `xml:"Prtry"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text      `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

type ReturnAccountV08 struct {
	MsgHdr      MessageHeader7                   `xml:"MsgHdr"`
	RptOrErr    AccountOrOperationalError4Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type StandingOrder6 struct {
	Amt             Amount2Choice                                `xml:"Amt"`
	CdtDbtInd       common.CreditDebitCode                       `xml:"CdtDbtInd"`
	Ccy             common.ActiveCurrencyCode                    `xml:"Ccy,omitempty" json:",omitempty"`
	Tp              StandingOrderType1Choice                     `xml:"Tp,omitempty" json:",omitempty"`
	AssoctdPoolAcct AccountIdentification4Choice                 `xml:"AssoctdPoolAcct,omitempty" json:",omitempty"`
	Ref             common.Max35Text                             `xml:"Ref,omitempty" json:",omitempty"`
	Frqcy           Frequency2Code                               `xml:"Frqcy,omitempty" json:",omitempty"`
	VldtyPrd        DatePeriodDetails1                           `xml:"VldtyPrd,omitempty" json:",omitempty"`
	SysMmb          BranchAndFinancialInstitutionIdentification6 `xml:"SysMmb,omitempty" json:",omitempty"`
	RspnsblPty      BranchAndFinancialInstitutionIdentification6 `xml:"RspnsblPty,omitempty" json:",omitempty"`
	LkSetId         common.Max35Text                             `xml:"LkSetId,omitempty" json:",omitempty"`
	LkSetOrdrId     common.Max35Text                             `xml:"LkSetOrdrId,omitempty" json:",omitempty"`
	LkSetOrdrSeq    float64                                      `xml:"LkSetOrdrSeq,omitempty" json:",omitempty"`
	ExctnTp         ExecutionType1Choice                         `xml:"ExctnTp,omitempty" json:",omitempty"`
	Cdtr            BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct        CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	Dbtr            BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct        CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	TtlsPerStgOrdr  StandingOrderTotalAmount1                    `xml:"TtlsPerStgOrdr,omitempty" json:",omitempty"`
	ZeroSweepInd    bool                                         `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
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

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TotalAmountAndCurrency1 struct {
	TtlAmt    float64                   `xml:"TtlAmt"`
	CdtDbtInd common.CreditDebitCode    `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Ccy       common.ActiveCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
}

type AccountCashEntryReturnCriteria3 struct {
	NtryRefInd  bool `xml:"NtryRefInd,omitempty" json:",omitempty"`
	AcctTpInd   bool `xml:"AcctTpInd,omitempty" json:",omitempty"`
	NtryAmtInd  bool `xml:"NtryAmtInd,omitempty" json:",omitempty"`
	AcctCcyInd  bool `xml:"AcctCcyInd,omitempty" json:",omitempty"`
	NtryStsInd  bool `xml:"NtryStsInd,omitempty" json:",omitempty"`
	NtryDtInd   bool `xml:"NtryDtInd,omitempty" json:",omitempty"`
	AcctSvcrInd bool `xml:"AcctSvcrInd,omitempty" json:",omitempty"`
	AcctOwnrInd bool `xml:"AcctOwnrInd,omitempty" json:",omitempty"`
}

type AccountIdentificationSearchCriteria2Choice struct {
	EQ     AccountIdentification4Choice `xml:"EQ"`
	CTTxt  common.Max35Text             `xml:"CTTxt"`
	NCTTxt common.Max35Text             `xml:"NCTTxt"`
}

type ActiveAmountRange3Choice struct {
	ImpldCcyAndAmtRg ImpliedCurrencyAndAmountRange1 `xml:"ImpldCcyAndAmtRg"`
	CcyAndAmtRg      ActiveCurrencyAndAmountRange3  `xml:"CcyAndAmtRg"`
}

type ActiveCurrencyAndAmountRange3 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Ccy       common.ActiveCurrencyCode         `xml:"Ccy"`
}

type ActiveOrHistoricAmountRange2Choice struct {
	ImpldCcyAndAmtRg ImpliedCurrencyAndAmountRange1          `xml:"ImpldCcyAndAmtRg"`
	CcyAndAmtRg      ActiveOrHistoricCurrencyAndAmountRange2 `xml:"CcyAndAmtRg"`
}

type ActiveOrHistoricCurrencyAndAmountRange2 struct {
	Amt       ImpliedCurrencyAmountRange1Choice   `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode              `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Ccy       common.ActiveOrHistoricCurrencyCode `xml:"Ccy"`
}

type AmountRangeBoundary1 struct {
	BdryAmt float64 `xml:"BdryAmt"`
	Incl    bool    `xml:"Incl"`
}

type CashAccountEntrySearch6 struct {
	AcctId     []AccountIdentificationSearchCriteria2Choice `xml:"AcctId,omitempty" json:",omitempty"`
	NtryAmt    []ActiveOrHistoricAmountRange2Choice         `xml:"NtryAmt,omitempty" json:",omitempty"`
	NtryAmtCcy []common.ActiveOrHistoricCurrencyCode        `xml:"NtryAmtCcy,omitempty" json:",omitempty"`
	CdtDbtInd  common.CreditDebitCode                       `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	NtrySts    []EntryStatus1Code                           `xml:"NtrySts,omitempty" json:",omitempty"`
	NtryDt     []DateAndDateTimeSearch3Choice               `xml:"NtryDt,omitempty" json:",omitempty"`
	AcctOwnr   PartyIdentification135                       `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type DateAndDateTimeSearch3Choice struct {
	DtTmSch DateTimePeriod1Choice   `xml:"DtTmSch"`
	DtSch   DatePeriodSearch1Choice `xml:"DtSch"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type DatePeriodSearch1Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
	EQDt   common.ISODate `xml:"EQDt"`
	NEQDt  common.ISODate `xml:"NEQDt"`
}

type DateTimePeriod1 struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
}

type DateTimePeriod1Choice struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
	DtTmRg DateTimePeriod1    `xml:"DtTmRg"`
}

type FromToAmountRange1 struct {
	FrAmt AmountRangeBoundary1 `xml:"FrAmt"`
	ToAmt AmountRangeBoundary1 `xml:"ToAmt"`
}

type GetTransactionV08 struct {
	MsgHdr      MessageHeader9       `xml:"MsgHdr"`
	TxQryDef    TransactionQuery5    `xml:"TxQryDef,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type ImpliedCurrencyAmountRange1Choice struct {
	FrAmt   AmountRangeBoundary1 `xml:"FrAmt"`
	ToAmt   AmountRangeBoundary1 `xml:"ToAmt"`
	FrToAmt FromToAmountRange1   `xml:"FrToAmt"`
	EQAmt   float64              `xml:"EQAmt"`
	NEQAmt  float64              `xml:"NEQAmt"`
}

type ImpliedCurrencyAndAmountRange1 struct {
	Amt       ImpliedCurrencyAmountRange1Choice `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

type InstructionStatusReturnCriteria1 struct {
	PmtInstrStsInd     bool `xml:"PmtInstrStsInd"`
	PmtInstrStsDtTmInd bool `xml:"PmtInstrStsDtTmInd,omitempty" json:",omitempty"`
	PmtInstrStsRsnInd  bool `xml:"PmtInstrStsRsnInd,omitempty" json:",omitempty"`
}

type InstructionStatusSearch5 struct {
	PmtInstrSts     PaymentStatusCodeSearch2Choice `xml:"PmtInstrSts,omitempty" json:",omitempty"`
	PmtInstrStsDtTm DateTimePeriod1Choice          `xml:"PmtInstrStsDtTm,omitempty" json:",omitempty"`
	PrtryStsRsn     common.Max4AlphaNumericText    `xml:"PrtryStsRsn,omitempty" json:",omitempty"`
}

type LongPaymentIdentification2 struct {
	TxId           common.Max35Text                             `xml:"TxId,omitempty" json:",omitempty"`
	UETR           common.UUIDv4Identifier                      `xml:"UETR,omitempty" json:",omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                               `xml:"IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"PmtMtd,omitempty" json:",omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"NtryTp,omitempty" json:",omitempty"`
	EndToEndId     common.Max35Text                             `xml:"EndToEndId,omitempty" json:",omitempty"`
}

type MessageHeader9 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   RequestType4Choice `xml:"ReqTp,omitempty" json:",omitempty"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

type PaymentIdentification6Choice struct {
	TxId      common.Max35Text                `xml:"TxId"`
	QId       QueueTransactionIdentification1 `xml:"QId"`
	LngBizId  LongPaymentIdentification2      `xml:"LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"ShrtBizId"`
	PrtryId   common.Max70Text                `xml:"PrtryId"`
}

type PaymentOrigin1Choice struct {
	FINMT    common.Max3NumericText `xml:"FINMT"`
	XMLMsgNm common.Max35Text       `xml:"XMLMsgNm"`
	Prtry    common.Max35Text       `xml:"Prtry"`
	Instrm   PaymentInstrument1Code `xml:"Instrm"`
}

type PaymentReturnCriteria4 struct {
	MsgIdInd            bool                             `xml:"MsgIdInd,omitempty" json:",omitempty"`
	ReqdExctnDtInd      bool                             `xml:"ReqdExctnDtInd,omitempty" json:",omitempty"`
	InstrInd            bool                             `xml:"InstrInd,omitempty" json:",omitempty"`
	InstrStsRtrCrit     InstructionStatusReturnCriteria1 `xml:"InstrStsRtrCrit,omitempty" json:",omitempty"`
	InstdAmtInd         bool                             `xml:"InstdAmtInd,omitempty" json:",omitempty"`
	CdtDbtInd           bool                             `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	IntrBkSttlmAmtInd   bool                             `xml:"IntrBkSttlmAmtInd,omitempty" json:",omitempty"`
	PrtyInd             bool                             `xml:"PrtyInd,omitempty" json:",omitempty"`
	PrcgVldtyTmInd      bool                             `xml:"PrcgVldtyTmInd,omitempty" json:",omitempty"`
	PurpInd             bool                             `xml:"PurpInd,omitempty" json:",omitempty"`
	InstrCpyInd         bool                             `xml:"InstrCpyInd,omitempty" json:",omitempty"`
	PmtMTInd            bool                             `xml:"PmtMTInd,omitempty" json:",omitempty"`
	PmtTpInd            bool                             `xml:"PmtTpInd,omitempty" json:",omitempty"`
	TxIdInd             bool                             `xml:"TxIdInd,omitempty" json:",omitempty"`
	IntrBkSttlmDtInd    bool                             `xml:"IntrBkSttlmDtInd,omitempty" json:",omitempty"`
	EndToEndIdInd       bool                             `xml:"EndToEndIdInd,omitempty" json:",omitempty"`
	PmtMtdInd           bool                             `xml:"PmtMtdInd,omitempty" json:",omitempty"`
	DbtrInd             bool                             `xml:"DbtrInd,omitempty" json:",omitempty"`
	DbtrAgtInd          bool                             `xml:"DbtrAgtInd,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtInd bool                             `xml:"InstgRmbrsmntAgtInd,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtInd bool                             `xml:"InstdRmbrsmntAgtInd,omitempty" json:",omitempty"`
	IntrmyInd           bool                             `xml:"IntrmyInd,omitempty" json:",omitempty"`
	CdtrAgtInd          bool                             `xml:"CdtrAgtInd,omitempty" json:",omitempty"`
	CdtrInd             bool                             `xml:"CdtrInd,omitempty" json:",omitempty"`
}

type PaymentSearch8 struct {
	MsgId             []common.Max35Text                    `xml:"MsgId,omitempty" json:",omitempty"`
	ReqdExctnDt       []DateAndDateTimeSearch3Choice        `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	PmtId             []PaymentIdentification6Choice        `xml:"PmtId,omitempty" json:",omitempty"`
	Sts               []InstructionStatusSearch5            `xml:"Sts,omitempty" json:",omitempty"`
	InstdAmt          []ActiveOrHistoricAmountRange2Choice  `xml:"InstdAmt,omitempty" json:",omitempty"`
	InstdAmtCcy       []common.ActiveOrHistoricCurrencyCode `xml:"InstdAmtCcy,omitempty" json:",omitempty"`
	CdtDbtInd         common.CreditDebitCode                `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	IntrBkSttlmAmt    []ActiveAmountRange3Choice            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	IntrBkSttlmAmtCcy []common.ActiveCurrencyCode           `xml:"IntrBkSttlmAmtCcy,omitempty" json:",omitempty"`
	PmtMtd            []PaymentOrigin1Choice                `xml:"PmtMtd,omitempty" json:",omitempty"`
	PmtTp             []PaymentType4Choice                  `xml:"PmtTp,omitempty" json:",omitempty"`
	Prty              []Priority1Choice                     `xml:"Prty,omitempty" json:",omitempty"`
	PrcgVldtyTm       []DateTimePeriod1Choice               `xml:"PrcgVldtyTm,omitempty" json:",omitempty"`
	Instr             []Instruction1Code                    `xml:"Instr,omitempty" json:",omitempty"`
	TxId              []common.Max35Text                    `xml:"TxId,omitempty" json:",omitempty"`
	IntrBkSttlmDt     []common.ISODate                      `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	EndToEndId        []common.Max35Text                    `xml:"EndToEndId,omitempty" json:",omitempty"`
	Pties             PaymentTransactionParty3              `xml:"Pties,omitempty" json:",omitempty"`
}

type PaymentStatusCodeSearch2Choice struct {
	PdgSts       PendingStatus4Code     `xml:"PdgSts"`
	FnlSts       FinalStatusCode        `xml:"FnlSts"`
	PdgAndFnlSts CashPaymentStatus2Code `xml:"PdgAndFnlSts"`
}

type PaymentTransactionParty3 struct {
	InstgAgt         BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:",omitempty"`
	InstdAgt         BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:",omitempty"`
	UltmtDbtr        Party40Choice                                `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr             Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt          BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	IntrmyAgt1       BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt2       BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt3       BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	CdtrAgt          BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	Cdtr             Party40Choice                                `xml:"Cdtr,omitempty" json:",omitempty"`
	UltmtCdtr        Party40Choice                                `xml:"UltmtCdtr,omitempty" json:",omitempty"`
}

type PaymentType4Choice struct {
	Cd    PaymentType3Code `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

type Priority1Choice struct {
	Cd    Priority5Code    `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

type QueueTransactionIdentification1 struct {
	QId    common.Max16Text `xml:"QId"`
	PosInQ common.Max16Text `xml:"PosInQ"`
}

type ShortPaymentIdentification2 struct {
	TxId          common.Max35Text                             `xml:"TxId"`
	IntrBkSttlmDt common.ISODate                               `xml:"IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
}

type SystemReturnCriteria2 struct {
	SysIdInd  bool `xml:"SysIdInd,omitempty" json:",omitempty"`
	MmbIdInd  bool `xml:"MmbIdInd,omitempty" json:",omitempty"`
	CtryIdInd bool `xml:"CtryIdInd,omitempty" json:",omitempty"`
	AcctIdInd bool `xml:"AcctIdInd,omitempty" json:",omitempty"`
}

type SystemSearch4 struct {
	SysId  []ClearingSystemIdentification3Choice          `xml:"SysId,omitempty" json:",omitempty"`
	MmbId  []BranchAndFinancialInstitutionIdentification6 `xml:"MmbId,omitempty" json:",omitempty"`
	Ctry   common.CountryCode                             `xml:"Ctry,omitempty" json:",omitempty"`
	AcctId AccountIdentification4Choice                   `xml:"AcctId,omitempty" json:",omitempty"`
}

type TransactionCriteria5Choice struct {
	QryNm   common.Max35Text     `xml:"QryNm"`
	NewCrit TransactionCriteria8 `xml:"NewCrit"`
}

type TransactionCriteria8 struct {
	NewQryNm common.Max35Text             `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []TransactionSearchCriteria8 `xml:"SchCrit,omitempty" json:",omitempty"`
	StmtRpt  ReportIndicator1Code         `xml:"StmtRpt,omitempty" json:",omitempty"`
	RtrCrit  TransactionReturnCriteria5   `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type TransactionQuery5 struct {
	QryTp  QueryType2Code             `xml:"QryTp,omitempty" json:",omitempty"`
	TxCrit TransactionCriteria5Choice `xml:"TxCrit,omitempty" json:",omitempty"`
}

type TransactionReturnCriteria5 struct {
	PmtToRtrCrit       SystemReturnCriteria2           `xml:"PmtToRtrCrit,omitempty" json:",omitempty"`
	PmtFrRtrCrit       SystemReturnCriteria2           `xml:"PmtFrRtrCrit,omitempty" json:",omitempty"`
	AcctCshNtryRtrCrit AccountCashEntryReturnCriteria3 `xml:"AcctCshNtryRtrCrit,omitempty" json:",omitempty"`
	PmtRtrCrit         PaymentReturnCriteria4          `xml:"PmtRtrCrit,omitempty" json:",omitempty"`
}

type TransactionSearchCriteria8 struct {
	PmtTo       []SystemSearch4         `xml:"PmtTo,omitempty" json:",omitempty"`
	PmtFr       []SystemSearch4         `xml:"PmtFr,omitempty" json:",omitempty"`
	PmtSch      PaymentSearch8          `xml:"PmtSch,omitempty" json:",omitempty"`
	AcctNtrySch CashAccountEntrySearch6 `xml:"AcctNtrySch,omitempty" json:",omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type Amount3Choice struct {
	AmtWthCcy  ActiveOrHistoricCurrencyAndAmount `xml:"AmtWthCcy"`
	AmtWthtCcy float64                           `xml:"AmtWthtCcy"`
}

type CashAccount39 struct {
	Id   AccountIdentification4Choice                 `xml:"Id"`
	Tp   CashAccountType2Choice                       `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode          `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   common.Max70Text                             `xml:"Nm,omitempty" json:",omitempty"`
	Prxy ProxyAccountIdentification1                  `xml:"Prxy,omitempty" json:",omitempty"`
	Ownr PartyIdentification135                       `xml:"Ownr,omitempty" json:",omitempty"`
	Svcr BranchAndFinancialInstitutionIdentification6 `xml:"Svcr,omitempty" json:",omitempty"`
}

type CashAccountAndEntry3 struct {
	Acct CashAccount39 `xml:"Acct"`
	Ntry CashEntry2    `xml:"Ntry,omitempty" json:",omitempty"`
}

type CashEntry2 struct {
	Amt          ActiveCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Dt           DateAndDateTime2Choice  `xml:"Dt,omitempty" json:",omitempty"`
	Sts          EntryStatus1Code        `xml:"Sts,omitempty" json:",omitempty"`
	Id           common.Max35Text        `xml:"Id,omitempty" json:",omitempty"`
	StmtId       common.Max35Text        `xml:"StmtId,omitempty" json:",omitempty"`
	AcctSvcrRef  float64                 `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	AddtlNtryInf []common.Max140Text     `xml:"AddtlNtryInf,omitempty" json:",omitempty"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type MessageHeader8 struct {
	MsgId       common.Max35Text       `xml:"MsgId"`
	CreDtTm     common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	MsgPgntn    Pagination1            `xml:"MsgPgntn,omitempty" json:",omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	ReqTp       RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	QryNm       common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

type NumberAndSumOfTransactions2 struct {
	NbOfNtries    common.Max15NumericText `xml:"NbOfNtries,omitempty" json:",omitempty"`
	Sum           float64                 `xml:"Sum,omitempty" json:",omitempty"`
	TtlNetNtryAmt float64                 `xml:"TtlNetNtryAmt,omitempty" json:",omitempty"`
	CdtDbtInd     common.CreditDebitCode  `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

type Pagination1 struct {
	PgNb      Max5NumericText `xml:"PgNb"`
	LastPgInd bool            `xml:"LastPgInd"`
}

type PaymentCommon4 struct {
	PmtFr       System2                `xml:"PmtFr,omitempty" json:",omitempty"`
	PmtTo       System2                `xml:"PmtTo,omitempty" json:",omitempty"`
	CmonSts     []PaymentStatus6       `xml:"CmonSts,omitempty" json:",omitempty"`
	ReqdExctnDt DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	NtryDt      DateAndDateTime2Choice `xml:"NtryDt,omitempty" json:",omitempty"`
	CdtDbtInd   common.CreditDebitCode `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	PmtMtd      PaymentOrigin1Choice   `xml:"PmtMtd,omitempty" json:",omitempty"`
}

type PaymentInstruction32 struct {
	MsgId          common.Max35Text         `xml:"MsgId,omitempty" json:",omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice   `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	Sts            []PaymentStatus6         `xml:"Sts,omitempty" json:",omitempty"`
	InstdAmt       Amount3Choice            `xml:"InstdAmt,omitempty" json:",omitempty"`
	IntrBkSttlmAmt Amount2Choice            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Purp           common.Max10Text         `xml:"Purp,omitempty" json:",omitempty"`
	PmtMtd         PaymentOrigin1Choice     `xml:"PmtMtd,omitempty" json:",omitempty"`
	Prty           Priority1Choice          `xml:"Prty,omitempty" json:",omitempty"`
	PrcgVldtyTm    DateTimePeriod1Choice    `xml:"PrcgVldtyTm,omitempty" json:",omitempty"`
	InstrCpy       common.Max20000Text      `xml:"InstrCpy,omitempty" json:",omitempty"`
	Tp             PaymentType4Choice       `xml:"Tp,omitempty" json:",omitempty"`
	GnrtdOrdr      bool                     `xml:"GnrtdOrdr,omitempty" json:",omitempty"`
	TxId           common.Max35Text         `xml:"TxId,omitempty" json:",omitempty"`
	IntrBkSttlmDt  common.ISODate           `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	EndToEndId     common.Max35Text         `xml:"EndToEndId,omitempty" json:",omitempty"`
	Pties          PaymentTransactionParty3 `xml:"Pties,omitempty" json:",omitempty"`
}

type PaymentStatus6 struct {
	Cd   PaymentStatusCode6Choice     `xml:"Cd,omitempty" json:",omitempty"`
	DtTm DateAndDateTime2Choice       `xml:"DtTm,omitempty" json:",omitempty"`
	Rsn  []PaymentStatusReason1Choice `xml:"Rsn,omitempty" json:",omitempty"`
}

type PaymentStatusCode6Choice struct {
	Pdg   PendingStatus4Code          `xml:"Pdg"`
	Fnl   FinalStatus1Code            `xml:"Fnl"`
	RTGS  common.Max4AlphaNumericText `xml:"RTGS"`
	Sttlm common.Max4AlphaNumericText `xml:"Sttlm"`
	Prtry common.Max35Text            `xml:"Prtry"`
}

type PaymentStatusReason1Choice struct {
	Umtchd       UnmatchedStatusReason1Code      `xml:"Umtchd"`
	Canc         CancelledStatusReason1Code      `xml:"Canc"`
	Sspd         SuspendedStatusReason1Code      `xml:"Sspd"`
	PdgFlngSttlm PendingFailingSettlement1Code   `xml:"PdgFlngSttlm"`
	PdgSttlm     PendingSettlement2Code          `xml:"PdgSttlm"`
	PrtryRjctn   ProprietaryStatusJustification2 `xml:"PrtryRjctn"`
	Prtry        common.Max35Text                `xml:"Prtry"`
}

type ProprietaryStatusJustification2 struct {
	PrtryStsRsn common.Max4AlphaNumericText `xml:"PrtryStsRsn"`
	Rsn         common.Max256Text           `xml:"Rsn"`
}

type ReturnTransactionV08 struct {
	MsgHdr      MessageHeader8                  `xml:"MsgHdr"`
	RptOrErr    TransactionReportOrError4Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type SecuritiesTransactionReferences1 struct {
	AcctOwnrTxId      common.Max35Text `xml:"AcctOwnrTxId,omitempty" json:",omitempty"`
	AcctSvcrTxId      common.Max35Text `xml:"AcctSvcrTxId,omitempty" json:",omitempty"`
	MktInfrstrctrTxId common.Max35Text `xml:"MktInfrstrctrTxId,omitempty" json:",omitempty"`
	PrcgId            common.Max35Text `xml:"PrcgId,omitempty" json:",omitempty"`
}

type System2 struct {
	SysId  MarketInfrastructureIdentification1Choice    `xml:"SysId,omitempty" json:",omitempty"`
	MmbId  BranchAndFinancialInstitutionIdentification6 `xml:"MmbId,omitempty" json:",omitempty"`
	Ctry   common.CountryCode                           `xml:"Ctry,omitempty" json:",omitempty"`
	AcctId AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type Transaction66 struct {
	PmtTo        System2                          `xml:"PmtTo,omitempty" json:",omitempty"`
	PmtFr        System2                          `xml:"PmtFr,omitempty" json:",omitempty"`
	CdtDbtInd    common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Pmt          PaymentInstruction32             `xml:"Pmt,omitempty" json:",omitempty"`
	AcctNtry     CashAccountAndEntry3             `xml:"AcctNtry,omitempty" json:",omitempty"`
	SctiesTxRefs SecuritiesTransactionReferences1 `xml:"SctiesTxRefs,omitempty" json:",omitempty"`
}

type TransactionOrError4Choice struct {
	Tx     Transaction66    `xml:"Tx"`
	BizErr []ErrorHandling5 `xml:"BizErr"`
}

type TransactionReport5 struct {
	PmtId   PaymentIdentification6Choice `xml:"PmtId"`
	TxOrErr TransactionOrError4Choice    `xml:"TxOrErr"`
}

type TransactionReportOrError4Choice struct {
	BizRpt  Transactions8    `xml:"BizRpt"`
	OprlErr []ErrorHandling5 `xml:"OprlErr"`
}

type Transactions8 struct {
	PmtCmonInf PaymentCommon4              `xml:"PmtCmonInf,omitempty" json:",omitempty"`
	TxsSummry  NumberAndSumOfTransactions2 `xml:"TxsSummry,omitempty" json:",omitempty"`
	TxRpt      []TransactionReport5        `xml:"TxRpt"`
}

type MessageHeader1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type ModifyTransactionV08 struct {
	MsgHdr      MessageHeader1             `xml:"MsgHdr"`
	Mod         []TransactionModification5 `xml:"Mod"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type PaymentInstruction33 struct {
	Instr       Instruction1Code      `xml:"Instr,omitempty" json:",omitempty"`
	Tp          PaymentType4Choice    `xml:"Tp,omitempty" json:",omitempty"`
	Prty        Priority1Choice       `xml:"Prty,omitempty" json:",omitempty"`
	PrcgVldtyTm DateTimePeriod1Choice `xml:"PrcgVldtyTm,omitempty" json:",omitempty"`
}

type TransactionModification5 struct {
	PmtId        PaymentIdentification6Choice `xml:"PmtId"`
	NewPmtValSet PaymentInstruction33         `xml:"NewPmtValSet"`
}

type CancelTransactionV08 struct {
	MsgHdr      MessageHeader9               `xml:"MsgHdr"`
	PmtId       PaymentIdentification6Choice `xml:"PmtId"`
	CshAcct     CashAccount38                `xml:"CshAcct,omitempty" json:",omitempty"`
	CxlRsn      PaymentCancellationReason5   `xml:"CxlRsn,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type CancellationReason33Choice struct {
	Cd    ExternalCancellationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type PaymentCancellationReason5 struct {
	Orgtr    PartyIdentification135     `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      CancellationReason33Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text        `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type Limit7 struct {
	Amt             Amount2Choice          `xml:"Amt"`
	CdtDbtInd       common.CreditDebitCode `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Sts             LimitStatus1Code       `xml:"Sts,omitempty" json:",omitempty"`
	StartDtTm       DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	UsdAmt          Amount2Choice          `xml:"UsdAmt,omitempty" json:",omitempty"`
	UsdAmtCdtDbtInd common.CreditDebitCode `xml:"UsdAmtCdtDbtInd,omitempty" json:",omitempty"`
	UsdPctg         float64                `xml:"UsdPctg,omitempty" json:",omitempty"`
	RmngAmt         Amount2Choice          `xml:"RmngAmt,omitempty" json:",omitempty"`
}

type LimitIdentification5 struct {
	SysId          SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	BilLmtCtrPtyId BranchAndFinancialInstitutionIdentification6 `xml:"BilLmtCtrPtyId,omitempty" json:",omitempty"`
	Tp             LimitType1Choice                             `xml:"Tp"`
	AcctOwnr       BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId         AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type LimitOrError4Choice struct {
	Lmt    Limit7           `xml:"Lmt"`
	BizErr []ErrorHandling5 `xml:"BizErr"`
}

type LimitReport7 struct {
	LmtId    LimitIdentification5 `xml:"LmtId"`
	LmtOrErr LimitOrError4Choice  `xml:"LmtOrErr"`
}

type LimitReportOrError4Choice struct {
	BizRpt  Limits7          `xml:"BizRpt"`
	OprlErr []ErrorHandling5 `xml:"OprlErr"`
}

type LimitType1Choice struct {
	Cd    LimitType3Code   `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

type Limits7 struct {
	CurLmt  []LimitReport7 `xml:"CurLmt,omitempty" json:",omitempty"`
	DfltLmt []LimitReport7 `xml:"DfltLmt,omitempty" json:",omitempty"`
}

type ReturnLimitV08 struct {
	MsgHdr      MessageHeader7            `xml:"MsgHdr"`
	RptOrErr    LimitReportOrError4Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      common.Max35Text                             `xml:"OrgnlMndtId,omitempty" json:",omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:",omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:",omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:",omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:",omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:",omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlFnlColltnDt common.ISODate                               `xml:"OrgnlFnlColltnDt,omitempty" json:",omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:",omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:",omitempty"`
	OrgnlTrckgDays   Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty" json:",omitempty"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type CreditTransferMandateData1 struct {
	MndtId       common.Max35Text          `xml:"MndtId,omitempty" json:",omitempty"`
	Tp           MandateTypeInformation2   `xml:"Tp,omitempty" json:",omitempty"`
	DtOfSgntr    common.ISODate            `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	DtOfVrfctn   common.ISODateTime        `xml:"DtOfVrfctn,omitempty" json:",omitempty"`
	ElctrncSgntr common.Max10KBinary       `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstPmtDt    common.ISODate            `xml:"FrstPmtDt,omitempty" json:",omitempty"`
	FnlPmtDt     common.ISODate            `xml:"FnlPmtDt,omitempty" json:",omitempty"`
	Frqcy        Frequency36Choice         `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn          MandateSetupReason1Choice `xml:"Rsn,omitempty" json:",omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text             `xml:"Issr,omitempty" json:",omitempty"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       common.Max4Text                   `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  common.Max140Text                 `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc common.Max2048Text            `xml:"Desc,omitempty" json:",omitempty"`
	Amt  RemittanceAmount3             `xml:"Amt,omitempty" json:",omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text        `xml:"Issr,omitempty" json:",omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp"`
	Prd    FrequencyPeriod1    `xml:"Prd"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"Tp"`
	PtInTm Exact2NumericText `xml:"PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"Tp"`
	Grnshee           PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                              `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:",omitempty"`
	CdtTrfMndt  CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:",omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        common.Max35Text              `xml:"MndtId,omitempty" json:",omitempty"`
	DtOfSgntr     common.ISODate                `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	AmdmntInd     bool                          `xml:"AmdmntInd,omitempty" json:",omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:",omitempty"`
	ElctrncSgntr  common.Max1025Text            `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstColltnDt  common.ISODate                `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt   common.ISODate                `xml:"FnlColltnDt,omitempty" json:",omitempty"`
	Frqcy         Frequency36Choice             `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:",omitempty"`
	TrckgDays     Exact2NumericText             `xml:"TrckgDays,omitempty" json:",omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
}

type MissingOrIncorrectInformation3 struct {
	AMLReq     bool                      `xml:"AMLReq,omitempty" json:",omitempty"`
	MssngInf   []UnableToApplyMissing1   `xml:"MssngInf,omitempty" json:",omitempty"`
	IncrrctInf []UnableToApplyIncorrect1 `xml:"IncrrctInf,omitempty" json:",omitempty"`
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   common.Max35Text   `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text   `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           Party40Choice                                `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
}

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"Tp,omitempty" json:",omitempty"`
	Nb       common.Max35Text           `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   common.ISODate             `xml:"RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:",omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

type SettlementInstruction7 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       PartyIdentification135         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      PartyIdentification135         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      TaxInformation7                `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"Dtls,omitempty" json:",omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:",omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      TaxPeriod2        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   TaxAmount2        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type UnableToApplyIncorrect1 struct {
	Cd              UnableToApplyIncorrectInformation4Code `xml:"Cd"`
	AddtlIncrrctInf common.Max140Text                      `xml:"AddtlIncrrctInf,omitempty" json:",omitempty"`
}

type UnableToApplyJustification3Choice struct {
	AnyInf            bool                           `xml:"AnyInf"`
	MssngOrIncrrctInf MissingOrIncorrectInformation3 `xml:"MssngOrIncrrctInf"`
	PssblDplctInstr   bool                           `xml:"PssblDplctInstr"`
}

type UnableToApplyMissing1 struct {
	Cd            UnableToApplyMissingInformation3Code `xml:"Cd"`
	AddtlMssngInf common.Max140Text                    `xml:"AddtlMssngInf,omitempty" json:",omitempty"`
}

type UnableToApplyV08 struct {
	Assgnmt     CaseAssignment5                   `xml:"Assgnmt"`
	Case        Case5                             `xml:"Case,omitempty" json:",omitempty"`
	Undrlyg     UnderlyingTransaction6Choice      `xml:"Undrlyg"`
	Justfn      UnableToApplyJustification3Choice `xml:"Justfn"`
	SplmtryData []SupplementaryData1              `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type UnderlyingGroupInformation1 struct {
	OrgnlMsgId         common.Max35Text   `xml:"OrgnlMsgId"`
	OrgnlMsgNmId       common.Max35Text   `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm       common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlMsgDlvryChanl common.Max35Text   `xml:"OrgnlMsgDlvryChanl,omitempty" json:",omitempty"`
}

type UnderlyingPaymentInstruction6 struct {
	OrgnlGrpInf     UnderlyingGroupInformation1       `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlPmtInfId   common.Max35Text                  `xml:"OrgnlPmtInfId,omitempty" json:",omitempty"`
	OrgnlInstrId    common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`
	ReqdExctnDt     DateAndDateTime2Choice            `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt    common.ISODate                    `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	OrgnlTxRef      OriginalTransactionReference31    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

type UnderlyingPaymentTransaction5 struct {
	OrgnlGrpInf         UnderlyingGroupInformation1       `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           common.Max35Text                  `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlUETR           common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`
	OrgnlIntrBkSttlmDt  common.ISODate                    `xml:"OrgnlIntrBkSttlmDt"`
	OrgnlTxRef          OriginalTransactionReference31    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

type UnderlyingStatementEntry3 struct {
	OrgnlGrpInf OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlStmtId common.Max35Text           `xml:"OrgnlStmtId,omitempty" json:",omitempty"`
	OrgnlNtryId common.Max35Text           `xml:"OrgnlNtryId,omitempty" json:",omitempty"`
	OrgnlUETR   common.UUIDv4Identifier    `xml:"OrgnlUETR,omitempty" json:",omitempty"`
}

type UnderlyingTransaction6Choice struct {
	Initn    UnderlyingPaymentInstruction6 `xml:"Initn"`
	IntrBk   UnderlyingPaymentTransaction5 `xml:"IntrBk"`
	StmtNtry UnderlyingStatementEntry3     `xml:"StmtNtry"`
}

type ClaimNonReceiptV08 struct {
	Assgnmt        CaseAssignment5              `xml:"Assgnmt"`
	Case           Case5                        `xml:"Case,omitempty" json:",omitempty"`
	Undrlyg        UnderlyingTransaction6Choice `xml:"Undrlyg"`
	CoverDtls      MissingCover4                `xml:"CoverDtls,omitempty" json:",omitempty"`
	InstrForAssgne InstructionForAssignee1      `xml:"InstrForAssgne,omitempty" json:",omitempty"`
	SplmtryData    []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type InstructionForAssignee1 struct {
	Cd       ExternalAgentInstruction1Code `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf common.Max140Text             `xml:"InstrInf,omitempty" json:",omitempty"`
}

type MissingCover4 struct {
	MssngCoverInd bool                   `xml:"MssngCoverInd"`
	CoverCrrctn   SettlementInstruction6 `xml:"CoverCrrctn,omitempty" json:",omitempty"`
}

type SettlementInstruction6 struct {
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

type DebitAuthorisation2 struct {
	CxlRsn         CancellationReason33Choice        `xml:"CxlRsn"`
	AmtToDbt       ActiveOrHistoricCurrencyAndAmount `xml:"AmtToDbt,omitempty" json:",omitempty"`
	ValDtToDbt     common.ISODate                    `xml:"ValDtToDbt,omitempty" json:",omitempty"`
	AddtlCxlRsnInf []common.Max105Text               `xml:"AddtlCxlRsnInf,omitempty" json:",omitempty"`
}

type DebitAuthorisationRequestV08 struct {
	Assgnmt     CaseAssignment5              `xml:"Assgnmt"`
	Case        Case5                        `xml:"Case,omitempty" json:",omitempty"`
	Undrlyg     UnderlyingTransaction6Choice `xml:"Undrlyg"`
	Dtl         DebitAuthorisation2          `xml:"Dtl"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       PartyIdentification135         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      PartyIdentification135         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      TaxInformation7                `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

type AccountInterest4 struct {
	Tp     InterestType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Rate   []Rate4             `xml:"Rate,omitempty" json:",omitempty"`
	FrToDt DateTimePeriod1     `xml:"FrToDt,omitempty" json:",omitempty"`
	Rsn    common.Max35Text    `xml:"Rsn,omitempty" json:",omitempty"`
	Tax    TaxCharges2         `xml:"Tax,omitempty" json:",omitempty"`
}

type AccountReport25 struct {
	Id           common.Max35Text          `xml:"Id"`
	RptPgntn     Pagination1               `xml:"RptPgntn,omitempty" json:",omitempty"`
	ElctrncSeqNb float64                   `xml:"ElctrncSeqNb,omitempty" json:",omitempty"`
	RptgSeq      SequenceRange1Choice      `xml:"RptgSeq,omitempty" json:",omitempty"`
	LglSeqNb     float64                   `xml:"LglSeqNb,omitempty" json:",omitempty"`
	CreDtTm      common.ISODateTime        `xml:"CreDtTm,omitempty" json:",omitempty"`
	FrToDt       DateTimePeriod1           `xml:"FrToDt,omitempty" json:",omitempty"`
	CpyDplctInd  common.CopyDuplicate1Code `xml:"CpyDplctInd,omitempty" json:",omitempty"`
	RptgSrc      ReportingSource1Choice    `xml:"RptgSrc,omitempty" json:",omitempty"`
	Acct         CashAccount39             `xml:"Acct"`
	RltdAcct     CashAccount38             `xml:"RltdAcct,omitempty" json:",omitempty"`
	Intrst       []AccountInterest4        `xml:"Intrst,omitempty" json:",omitempty"`
	Bal          []CashBalance8            `xml:"Bal,omitempty" json:",omitempty"`
	TxsSummry    TotalTransactions6        `xml:"TxsSummry,omitempty" json:",omitempty"`
	Ntry         []ReportEntry10           `xml:"Ntry,omitempty" json:",omitempty"`
	AddtlRptInf  common.Max500Text         `xml:"AddtlRptInf,omitempty" json:",omitempty"`
}

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AmountAndCurrencyExchange3 struct {
	InstdAmt      AmountAndCurrencyExchangeDetails3   `xml:"InstdAmt,omitempty" json:",omitempty"`
	TxAmt         AmountAndCurrencyExchangeDetails3   `xml:"TxAmt,omitempty" json:",omitempty"`
	CntrValAmt    AmountAndCurrencyExchangeDetails3   `xml:"CntrValAmt,omitempty" json:",omitempty"`
	AnncdPstngAmt AmountAndCurrencyExchangeDetails3   `xml:"AnncdPstngAmt,omitempty" json:",omitempty"`
	PrtryAmt      []AmountAndCurrencyExchangeDetails4 `xml:"PrtryAmt,omitempty" json:",omitempty"`
}

type AmountAndCurrencyExchangeDetails3 struct {
	Amt     ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyXchg CurrencyExchange5                 `xml:"CcyXchg,omitempty" json:",omitempty"`
}

type AmountAndCurrencyExchangeDetails4 struct {
	Tp      common.Max35Text                  `xml:"Tp"`
	Amt     ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CcyXchg CurrencyExchange5                 `xml:"CcyXchg,omitempty" json:",omitempty"`
}

type AmountAndDirection35 struct {
	Amt       float64                `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode `xml:"CdtDbtInd"`
}

type BalanceSubType1Choice struct {
	Cd    ExternalBalanceSubType1Code `xml:"Cd"`
	Prtry common.Max35Text            `xml:"Prtry"`
}

type BalanceType10Choice struct {
	Cd    ExternalBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type BalanceType13 struct {
	CdOrPrtry BalanceType10Choice   `xml:"CdOrPrtry"`
	SubTp     BalanceSubType1Choice `xml:"SubTp,omitempty" json:",omitempty"`
}

type BankToCustomerAccountReportV08 struct {
	GrpHdr      GroupHeader81        `xml:"GrpHdr"`
	Rpt         []AccountReport25    `xml:"Rpt"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type BankTransactionCodeStructure4 struct {
	Domn  BankTransactionCodeStructure5            `xml:"Domn,omitempty" json:",omitempty"`
	Prtry ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty" json:",omitempty"`
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"Fmly"`
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"SubFmlyCd"`
}

type BatchInformation2 struct {
	MsgId     common.Max35Text                  `xml:"MsgId,omitempty" json:",omitempty"`
	PmtInfId  common.Max35Text                  `xml:"PmtInfId,omitempty" json:",omitempty"`
	NbOfTxs   common.Max15NumericText           `xml:"NbOfTxs,omitempty" json:",omitempty"`
	TtlAmt    ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

type CardAggregated2 struct {
	AddtlSvc      CardPaymentServiceType2Code          `xml:"AddtlSvc,omitempty" json:",omitempty"`
	TxCtgy        ExternalCardTransactionCategory1Code `xml:"TxCtgy,omitempty" json:",omitempty"`
	SaleRcncltnId common.Max35Text                     `xml:"SaleRcncltnId,omitempty" json:",omitempty"`
	SeqNbRg       CardSequenceNumberRange1             `xml:"SeqNbRg,omitempty" json:",omitempty"`
	TxDtRg        DateOrDateTimePeriod1Choice          `xml:"TxDtRg,omitempty" json:",omitempty"`
}

type CardEntry4 struct {
	Card      PaymentCard4        `xml:"Card,omitempty" json:",omitempty"`
	POI       PointOfInteraction1 `xml:"POI,omitempty" json:",omitempty"`
	AggtdNtry CardAggregated2     `xml:"AggtdNtry,omitempty" json:",omitempty"`
	PrePdAcct CashAccount38       `xml:"PrePdAcct,omitempty" json:",omitempty"`
}

type CardIndividualTransaction2 struct {
	ICCRltdData    common.Max1025Text                   `xml:"ICCRltdData,omitempty" json:",omitempty"`
	PmtCntxt       PaymentContext3                      `xml:"PmtCntxt,omitempty" json:",omitempty"`
	AddtlSvc       CardPaymentServiceType2Code          `xml:"AddtlSvc,omitempty" json:",omitempty"`
	TxCtgy         ExternalCardTransactionCategory1Code `xml:"TxCtgy,omitempty" json:",omitempty"`
	SaleRcncltnId  common.Max35Text                     `xml:"SaleRcncltnId,omitempty" json:",omitempty"`
	SaleRefNb      common.Max35Text                     `xml:"SaleRefNb,omitempty" json:",omitempty"`
	RePresntmntRsn ExternalRePresentmentReason1Code     `xml:"RePresntmntRsn,omitempty" json:",omitempty"`
	SeqNb          common.Max35Text                     `xml:"SeqNb,omitempty" json:",omitempty"`
	TxId           TransactionIdentifier1               `xml:"TxId,omitempty" json:",omitempty"`
	Pdct           Product2                             `xml:"Pdct,omitempty" json:",omitempty"`
	VldtnDt        common.ISODate                       `xml:"VldtnDt,omitempty" json:",omitempty"`
	VldtnSeqNb     common.Max35Text                     `xml:"VldtnSeqNb,omitempty" json:",omitempty"`
}

type CardSecurityInformation1 struct {
	CSCMgmt CSCManagement1Code         `xml:"CSCMgmt"`
	CSCVal  common.Min3Max4NumericText `xml:"CSCVal,omitempty" json:",omitempty"`
}

type CardSequenceNumberRange1 struct {
	FrstTx common.Max35Text `xml:"FrstTx,omitempty" json:",omitempty"`
	LastTx common.Max35Text `xml:"LastTx,omitempty" json:",omitempty"`
}

type CardTransaction17 struct {
	Card      PaymentCard4           `xml:"Card,omitempty" json:",omitempty"`
	POI       PointOfInteraction1    `xml:"POI,omitempty" json:",omitempty"`
	Tx        CardTransaction3Choice `xml:"Tx,omitempty" json:",omitempty"`
	PrePdAcct CashAccount38          `xml:"PrePdAcct,omitempty" json:",omitempty"`
}

type CardTransaction3Choice struct {
	Aggtd CardAggregated2            `xml:"Aggtd"`
	Indv  CardIndividualTransaction2 `xml:"Indv"`
}

type CardholderAuthentication2 struct {
	AuthntcnMtd  AuthenticationMethod1Code `xml:"AuthntcnMtd"`
	AuthntcnNtty AuthenticationEntity1Code `xml:"AuthntcnNtty"`
}

type CashAvailability1 struct {
	Dt        CashAvailabilityDate1Choice       `xml:"Dt"`
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd"`
}

type CashAvailabilityDate1Choice struct {
	NbOfDays common.Max15PlusSignedNumericText `xml:"NbOfDays"`
	ActlDt   common.ISODate                    `xml:"ActlDt"`
}

type CashBalance8 struct {
	Tp        BalanceType13                     `xml:"Tp"`
	CdtLine   []CreditLine3                     `xml:"CdtLine,omitempty" json:",omitempty"`
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd"`
	Dt        DateAndDateTime2Choice            `xml:"Dt"`
	Avlbty    []CashAvailability1               `xml:"Avlbty,omitempty" json:",omitempty"`
}

type CashDeposit1 struct {
	NoteDnmtn ActiveCurrencyAndAmount `xml:"NoteDnmtn"`
	NbOfNotes common.Max15NumericText `xml:"NbOfNotes"`
	Amt       ActiveCurrencyAndAmount `xml:"Amt"`
}

type ChargeType3Choice struct {
	Cd    ExternalChargeType1Code `xml:"Cd"`
	Prtry GenericIdentification3  `xml:"Prtry"`
}

type Charges6 struct {
	TtlChrgsAndTaxAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty" json:",omitempty"`
	Rcrd              []ChargesRecord3                  `xml:"Rcrd,omitempty" json:",omitempty"`
}

type ChargesRecord3 struct {
	Amt         ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	CdtDbtInd   common.CreditDebitCode                       `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	ChrgInclInd bool                                         `xml:"ChrgInclInd,omitempty" json:",omitempty"`
	Tp          ChargeType3Choice                            `xml:"Tp,omitempty" json:",omitempty"`
	Rate        float64                                      `xml:"Rate,omitempty" json:",omitempty"`
	Br          ChargeBearerType1Code                        `xml:"Br,omitempty" json:",omitempty"`
	Agt         BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:",omitempty"`
	Tax         TaxCharges2                                  `xml:"Tax,omitempty" json:",omitempty"`
}

type CorporateAction9 struct {
	EvtTp common.Max35Text `xml:"EvtTp"`
	EvtId common.Max35Text `xml:"EvtId"`
}

type CreditLine3 struct {
	Incl bool                              `xml:"Incl"`
	Tp   CreditLineType1Choice             `xml:"Tp,omitempty" json:",omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Dt   DateAndDateTime2Choice            `xml:"Dt,omitempty" json:",omitempty"`
}

type CreditLineType1Choice struct {
	Cd    ExternalCreditLineType1Code `xml:"Cd"`
	Prtry common.Max35Text            `xml:"Prtry"`
}

type CurrencyExchange5 struct {
	SrcCcy   common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy  common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy,omitempty" json:",omitempty"`
	UnitCcy  common.ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty" json:",omitempty"`
	XchgRate float64                             `xml:"XchgRate"`
	CtrctId  common.Max35Text                    `xml:"CtrctId,omitempty" json:",omitempty"`
	QtnDt    common.ISODateTime                  `xml:"QtnDt,omitempty" json:",omitempty"`
}

type DateOrDateTimePeriod1Choice struct {
	Dt   DatePeriod2     `xml:"Dt"`
	DtTm DateTimePeriod1 `xml:"DtTm"`
}

type DisplayCapabilities1 struct {
	DispTp    UserInterface2Code     `xml:"DispTp"`
	NbOfLines common.Max3NumericText `xml:"NbOfLines"`
	LineWidth common.Max3NumericText `xml:"LineWidth"`
}

type EntryDetails9 struct {
	Btch   BatchInformation2    `xml:"Btch,omitempty" json:",omitempty"`
	TxDtls []EntryTransaction10 `xml:"TxDtls,omitempty" json:",omitempty"`
}

type EntryStatus1Choice struct {
	Cd    ExternalEntryStatus1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type EntryTransaction10 struct {
	Refs        TransactionReferences6            `xml:"Refs,omitempty" json:",omitempty"`
	Amt         ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	CdtDbtInd   common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	AmtDtls     AmountAndCurrencyExchange3        `xml:"AmtDtls,omitempty" json:",omitempty"`
	Avlbty      []CashAvailability1               `xml:"Avlbty,omitempty" json:",omitempty"`
	BkTxCd      BankTransactionCodeStructure4     `xml:"BkTxCd,omitempty" json:",omitempty"`
	Chrgs       Charges6                          `xml:"Chrgs,omitempty" json:",omitempty"`
	Intrst      TransactionInterest4              `xml:"Intrst,omitempty" json:",omitempty"`
	RltdPties   TransactionParties6               `xml:"RltdPties,omitempty" json:",omitempty"`
	RltdAgts    TransactionAgents5                `xml:"RltdAgts,omitempty" json:",omitempty"`
	LclInstrm   LocalInstrument2Choice            `xml:"LclInstrm,omitempty" json:",omitempty"`
	Purp        Purpose2Choice                    `xml:"Purp,omitempty" json:",omitempty"`
	RltdRmtInf  []RemittanceLocation7             `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf      RemittanceInformation16           `xml:"RmtInf,omitempty" json:",omitempty"`
	RltdDts     TransactionDates3                 `xml:"RltdDts,omitempty" json:",omitempty"`
	RltdPric    TransactionPrice4Choice           `xml:"RltdPric,omitempty" json:",omitempty"`
	RltdQties   []TransactionQuantities3Choice    `xml:"RltdQties,omitempty" json:",omitempty"`
	FinInstrmId SecurityIdentification19          `xml:"FinInstrmId,omitempty" json:",omitempty"`
	Tax         TaxInformation8                   `xml:"Tax,omitempty" json:",omitempty"`
	RtrInf      PaymentReturnReason5              `xml:"RtrInf,omitempty" json:",omitempty"`
	CorpActn    CorporateAction9                  `xml:"CorpActn,omitempty" json:",omitempty"`
	SfkpgAcct   SecuritiesAccount19               `xml:"SfkpgAcct,omitempty" json:",omitempty"`
	CshDpst     []CashDeposit1                    `xml:"CshDpst,omitempty" json:",omitempty"`
	CardTx      CardTransaction17                 `xml:"CardTx,omitempty" json:",omitempty"`
	AddtlTxInf  common.Max500Text                 `xml:"AddtlTxInf,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1              `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"Unit"`
	FaceAmt  float64 `xml:"FaceAmt"`
	AmtsdVal float64 `xml:"AmtsdVal"`
}

type GenericIdentification3 struct {
	Id   common.Max35Text `xml:"Id"`
	Issr common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification32 struct {
	Id     common.Max35Text `xml:"Id"`
	Tp     PartyType3Code   `xml:"Tp,omitempty" json:",omitempty"`
	Issr   PartyType4Code   `xml:"Issr,omitempty" json:",omitempty"`
	ShrtNm common.Max35Text `xml:"ShrtNm,omitempty" json:",omitempty"`
}

type GroupHeader81 struct {
	MsgId       common.Max35Text       `xml:"MsgId"`
	CreDtTm     common.ISODateTime     `xml:"CreDtTm"`
	MsgRcpt     PartyIdentification135 `xml:"MsgRcpt,omitempty" json:",omitempty"`
	MsgPgntn    Pagination1            `xml:"MsgPgntn,omitempty" json:",omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	AddtlInf    common.Max500Text      `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"Cd"`
	Prtry common.Max35Text                                   `xml:"Prtry"`
}

type InterestRecord2 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd"`
	Tp        InterestType1Choice               `xml:"Tp,omitempty" json:",omitempty"`
	Rate      Rate4                             `xml:"Rate,omitempty" json:",omitempty"`
	FrToDt    DateTimePeriod1                   `xml:"FrToDt,omitempty" json:",omitempty"`
	Rsn       common.Max35Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	Tax       TaxCharges2                       `xml:"Tax,omitempty" json:",omitempty"`
}

type InterestType1Choice struct {
	Cd    common.InterestType1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type MessageIdentification2 struct {
	MsgNmId common.Max35Text `xml:"MsgNmId,omitempty" json:",omitempty"`
	MsgId   common.Max35Text `xml:"MsgId,omitempty" json:",omitempty"`
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress24   `xml:"Adr"`
}

type NumberAndSumOfTransactions1 struct {
	NbOfNtries common.Max15NumericText `xml:"NbOfNtries,omitempty" json:",omitempty"`
	Sum        float64                 `xml:"Sum,omitempty" json:",omitempty"`
}

type NumberAndSumOfTransactions4 struct {
	NbOfNtries common.Max15NumericText `xml:"NbOfNtries,omitempty" json:",omitempty"`
	Sum        float64                 `xml:"Sum,omitempty" json:",omitempty"`
	TtlNetNtry AmountAndDirection35    `xml:"TtlNetNtry,omitempty" json:",omitempty"`
}

type OriginalAndCurrentQuantities1 struct {
	FaceAmt  float64 `xml:"FaceAmt"`
	AmtsdVal float64 `xml:"AmtsdVal"`
}

type OtherIdentification1 struct {
	Id  common.Max35Text            `xml:"Id"`
	Sfx common.Max16Text            `xml:"Sfx,omitempty" json:",omitempty"`
	Tp  IdentificationSource3Choice `xml:"Tp"`
}

type PaymentCard4 struct {
	PlainCardData PlainCardData1         `xml:"PlainCardData,omitempty" json:",omitempty"`
	CardCtryCd    Exact3NumericText      `xml:"CardCtryCd,omitempty" json:",omitempty"`
	CardBrnd      GenericIdentification1 `xml:"CardBrnd,omitempty" json:",omitempty"`
	AddtlCardData common.Max70Text       `xml:"AddtlCardData,omitempty" json:",omitempty"`
}

type PaymentContext3 struct {
	CardPres       bool                        `xml:"CardPres,omitempty" json:",omitempty"`
	CrdhldrPres    bool                        `xml:"CrdhldrPres,omitempty" json:",omitempty"`
	OnLineCntxt    bool                        `xml:"OnLineCntxt,omitempty" json:",omitempty"`
	AttndncCntxt   AttendanceContext1Code      `xml:"AttndncCntxt,omitempty" json:",omitempty"`
	TxEnvt         TransactionEnvironment1Code `xml:"TxEnvt,omitempty" json:",omitempty"`
	TxChanl        TransactionChannel1Code     `xml:"TxChanl,omitempty" json:",omitempty"`
	AttndntMsgCpbl bool                        `xml:"AttndntMsgCpbl,omitempty" json:",omitempty"`
	AttndntLang    ISO2ALanguageCode           `xml:"AttndntLang,omitempty" json:",omitempty"`
	CardDataNtryMd CardDataReading1Code        `xml:"CardDataNtryMd"`
	FllbckInd      bool                        `xml:"FllbckInd,omitempty" json:",omitempty"`
	AuthntcnMtd    CardholderAuthentication2   `xml:"AuthntcnMtd,omitempty" json:",omitempty"`
}

type PaymentReturnReason5 struct {
	OrgnlBkTxCd BankTransactionCodeStructure4 `xml:"OrgnlBkTxCd,omitempty" json:",omitempty"`
	Orgtr       PartyIdentification135        `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn         ReturnReason5Choice           `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf    []common.Max105Text           `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type PlainCardData1 struct {
	PAN        common.Min8Max28NumericText `xml:"PAN"`
	CardSeqNb  common.Min2Max3NumericText  `xml:"CardSeqNb,omitempty" json:",omitempty"`
	FctvDt     common.ISOYearMonth         `xml:"FctvDt,omitempty" json:",omitempty"`
	XpryDt     common.ISOYearMonth         `xml:"XpryDt"`
	SvcCd      Exact3NumericText           `xml:"SvcCd,omitempty" json:",omitempty"`
	TrckData   []TrackData1                `xml:"TrckData,omitempty" json:",omitempty"`
	CardSctyCd CardSecurityInformation1    `xml:"CardSctyCd,omitempty" json:",omitempty"`
}

type PointOfInteraction1 struct {
	Id       GenericIdentification32         `xml:"Id"`
	SysNm    common.Max70Text                `xml:"SysNm,omitempty" json:",omitempty"`
	GrpId    common.Max35Text                `xml:"GrpId,omitempty" json:",omitempty"`
	Cpblties PointOfInteractionCapabilities1 `xml:"Cpblties,omitempty" json:",omitempty"`
	Cmpnt    []PointOfInteractionComponent1  `xml:"Cmpnt,omitempty" json:",omitempty"`
}

type PointOfInteractionCapabilities1 struct {
	CardRdngCpblties      []CardDataReading1Code                  `xml:"CardRdngCpblties,omitempty" json:",omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty" json:",omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"OnLineCpblties,omitempty" json:",omitempty"`
	DispCpblties          []DisplayCapabilities1                  `xml:"DispCpblties,omitempty" json:",omitempty"`
	PrtLineWidth          common.Max3NumericText                  `xml:"PrtLineWidth,omitempty" json:",omitempty"`
}

type PointOfInteractionComponent1 struct {
	POICmpntTp POIComponentType1Code `xml:"POICmpntTp"`
	ManfctrId  common.Max35Text      `xml:"ManfctrId,omitempty" json:",omitempty"`
	Mdl        common.Max35Text      `xml:"Mdl,omitempty" json:",omitempty"`
	VrsnNb     common.Max16Text      `xml:"VrsnNb,omitempty" json:",omitempty"`
	SrlNb      common.Max35Text      `xml:"SrlNb,omitempty" json:",omitempty"`
	ApprvlNb   []common.Max70Text    `xml:"ApprvlNb,omitempty" json:",omitempty"`
}

type Price7 struct {
	Tp  YieldedOrValueType1Choice `xml:"Tp"`
	Val PriceRateOrAmount3Choice  `xml:"Val"`
}

type PriceRateOrAmount3Choice struct {
	Rate float64                                    `xml:"Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`
}

type Product2 struct {
	PdctCd       common.Max70Text   `xml:"PdctCd"`
	UnitOfMeasr  UnitOfMeasure1Code `xml:"UnitOfMeasr,omitempty" json:",omitempty"`
	PdctQty      float64            `xml:"PdctQty,omitempty" json:",omitempty"`
	UnitPric     float64            `xml:"UnitPric,omitempty" json:",omitempty"`
	PdctAmt      float64            `xml:"PdctAmt,omitempty" json:",omitempty"`
	TaxTp        common.Max35Text   `xml:"TaxTp,omitempty" json:",omitempty"`
	AddtlPdctInf common.Max35Text   `xml:"AddtlPdctInf,omitempty" json:",omitempty"`
}

type ProprietaryAgent4 struct {
	Tp  common.Max35Text                             `xml:"Tp"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   common.Max35Text `xml:"Cd"`
	Issr common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type ProprietaryDate3 struct {
	Tp common.Max35Text       `xml:"Tp"`
	Dt DateAndDateTime2Choice `xml:"Dt"`
}

type ProprietaryParty5 struct {
	Tp  common.Max35Text `xml:"Tp"`
	Pty Party40Choice    `xml:"Pty"`
}

type ProprietaryPrice2 struct {
	Tp   common.Max35Text                  `xml:"Tp"`
	Pric ActiveOrHistoricCurrencyAndAmount `xml:"Pric"`
}

type ProprietaryQuantity1 struct {
	Tp  common.Max35Text `xml:"Tp"`
	Qty common.Max35Text `xml:"Qty"`
}

type ProprietaryReference1 struct {
	Tp  common.Max35Text `xml:"Tp"`
	Ref common.Max35Text `xml:"Ref"`
}

type Rate4 struct {
	Tp      RateType4Choice                         `xml:"Tp"`
	VldtyRg ActiveOrHistoricCurrencyAndAmountRange2 `xml:"VldtyRg,omitempty" json:",omitempty"`
}

type RateType4Choice struct {
	Pctg float64          `xml:"Pctg"`
	Othr common.Max35Text `xml:"Othr"`
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:",omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    NameAndAddress16              `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type ReportEntry10 struct {
	NtryRef       common.Max35Text                  `xml:"NtryRef,omitempty" json:",omitempty"`
	Amt           ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd     common.CreditDebitCode            `xml:"CdtDbtInd"`
	RvslInd       bool                              `xml:"RvslInd,omitempty" json:",omitempty"`
	Sts           EntryStatus1Choice                `xml:"Sts"`
	BookgDt       DateAndDateTime2Choice            `xml:"BookgDt,omitempty" json:",omitempty"`
	ValDt         DateAndDateTime2Choice            `xml:"ValDt,omitempty" json:",omitempty"`
	AcctSvcrRef   common.Max35Text                  `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	Avlbty        []CashAvailability1               `xml:"Avlbty,omitempty" json:",omitempty"`
	BkTxCd        BankTransactionCodeStructure4     `xml:"BkTxCd"`
	ComssnWvrInd  bool                              `xml:"ComssnWvrInd,omitempty" json:",omitempty"`
	AddtlInfInd   MessageIdentification2            `xml:"AddtlInfInd,omitempty" json:",omitempty"`
	AmtDtls       AmountAndCurrencyExchange3        `xml:"AmtDtls,omitempty" json:",omitempty"`
	Chrgs         Charges6                          `xml:"Chrgs,omitempty" json:",omitempty"`
	TechInptChanl TechnicalInputChannel1Choice      `xml:"TechInptChanl,omitempty" json:",omitempty"`
	Intrst        TransactionInterest4              `xml:"Intrst,omitempty" json:",omitempty"`
	CardTx        CardEntry4                        `xml:"CardTx,omitempty" json:",omitempty"`
	NtryDtls      []EntryDetails9                   `xml:"NtryDtls,omitempty" json:",omitempty"`
	AddtlNtryInf  common.Max500Text                 `xml:"AddtlNtryInf,omitempty" json:",omitempty"`
}

type ReportingSource1Choice struct {
	Cd    ExternalReportingSource1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type ReturnReason5Choice struct {
	Cd    ExternalReturnReason1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

type SecuritiesAccount19 struct {
	Id common.Max35Text        `xml:"Id"`
	Tp GenericIdentification30 `xml:"Tp,omitempty" json:",omitempty"`
	Nm common.Max70Text        `xml:"Nm,omitempty" json:",omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"ISIN,omitempty" json:",omitempty"`
	OthrId []OtherIdentification1 `xml:"OthrId,omitempty" json:",omitempty"`
	Desc   common.Max140Text      `xml:"Desc,omitempty" json:",omitempty"`
}

type SequenceRange1 struct {
	FrSeq common.Max35Text `xml:"FrSeq"`
	ToSeq common.Max35Text `xml:"ToSeq"`
}

type SequenceRange1Choice struct {
	FrSeq   common.Max35Text   `xml:"FrSeq"`
	ToSeq   common.Max35Text   `xml:"ToSeq"`
	FrToSeq []SequenceRange1   `xml:"FrToSeq"`
	EQSeq   []common.Max35Text `xml:"EQSeq"`
	NEQSeq  []common.Max35Text `xml:"NEQSeq"`
}

type TaxCharges2 struct {
	Id   common.Max35Text                  `xml:"Id,omitempty" json:",omitempty"`
	Rate float64                           `xml:"Rate,omitempty" json:",omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
}

type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:",omitempty"`
}

type TechnicalInputChannel1Choice struct {
	Cd    ExternalTechnicalInputChannel1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type TotalTransactions6 struct {
	TtlNtries          NumberAndSumOfTransactions4     `xml:"TtlNtries,omitempty" json:",omitempty"`
	TtlCdtNtries       NumberAndSumOfTransactions1     `xml:"TtlCdtNtries,omitempty" json:",omitempty"`
	TtlDbtNtries       NumberAndSumOfTransactions1     `xml:"TtlDbtNtries,omitempty" json:",omitempty"`
	TtlNtriesPerBkTxCd []TotalsPerBankTransactionCode5 `xml:"TtlNtriesPerBkTxCd,omitempty" json:",omitempty"`
}

type TotalsPerBankTransactionCode5 struct {
	NbOfNtries common.Max15NumericText       `xml:"NbOfNtries,omitempty" json:",omitempty"`
	Sum        float64                       `xml:"Sum,omitempty" json:",omitempty"`
	TtlNetNtry AmountAndDirection35          `xml:"TtlNetNtry,omitempty" json:",omitempty"`
	CdtNtries  NumberAndSumOfTransactions1   `xml:"CdtNtries,omitempty" json:",omitempty"`
	DbtNtries  NumberAndSumOfTransactions1   `xml:"DbtNtries,omitempty" json:",omitempty"`
	FcstInd    bool                          `xml:"FcstInd,omitempty" json:",omitempty"`
	BkTxCd     BankTransactionCodeStructure4 `xml:"BkTxCd"`
	Avlbty     []CashAvailability1           `xml:"Avlbty,omitempty" json:",omitempty"`
	Dt         DateAndDateTime2Choice        `xml:"Dt,omitempty" json:",omitempty"`
}

type TrackData1 struct {
	TrckNb  Exact1NumericText `xml:"TrckNb,omitempty" json:",omitempty"`
	TrckVal common.Max140Text `xml:"TrckVal"`
}

type TransactionAgents5 struct {
	InstgAgt   BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:",omitempty"`
	InstdAgt   BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:",omitempty"`
	DbtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt1 BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt2 BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt3 BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	RcvgAgt    BranchAndFinancialInstitutionIdentification6 `xml:"RcvgAgt,omitempty" json:",omitempty"`
	DlvrgAgt   BranchAndFinancialInstitutionIdentification6 `xml:"DlvrgAgt,omitempty" json:",omitempty"`
	IssgAgt    BranchAndFinancialInstitutionIdentification6 `xml:"IssgAgt,omitempty" json:",omitempty"`
	SttlmPlc   BranchAndFinancialInstitutionIdentification6 `xml:"SttlmPlc,omitempty" json:",omitempty"`
	Prtry      []ProprietaryAgent4                          `xml:"Prtry,omitempty" json:",omitempty"`
}

type TransactionDates3 struct {
	AccptncDtTm             common.ISODateTime `xml:"AccptncDtTm,omitempty" json:",omitempty"`
	TradActvtyCtrctlSttlmDt common.ISODate     `xml:"TradActvtyCtrctlSttlmDt,omitempty" json:",omitempty"`
	TradDt                  common.ISODate     `xml:"TradDt,omitempty" json:",omitempty"`
	IntrBkSttlmDt           common.ISODate     `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	StartDt                 common.ISODate     `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt                   common.ISODate     `xml:"EndDt,omitempty" json:",omitempty"`
	TxDtTm                  common.ISODateTime `xml:"TxDtTm,omitempty" json:",omitempty"`
	Prtry                   []ProprietaryDate3 `xml:"Prtry,omitempty" json:",omitempty"`
}

type TransactionIdentifier1 struct {
	TxDtTm common.ISODateTime `xml:"TxDtTm"`
	TxRef  common.Max35Text   `xml:"TxRef"`
}

type TransactionInterest4 struct {
	TtlIntrstAndTaxAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlIntrstAndTaxAmt,omitempty" json:",omitempty"`
	Rcrd               []InterestRecord2                 `xml:"Rcrd,omitempty" json:",omitempty"`
}

type TransactionParties6 struct {
	InitgPty  Party40Choice       `xml:"InitgPty,omitempty" json:",omitempty"`
	Dbtr      Party40Choice       `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct  CashAccount38       `xml:"DbtrAcct,omitempty" json:",omitempty"`
	UltmtDbtr Party40Choice       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Cdtr      Party40Choice       `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct  CashAccount38       `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr Party40Choice       `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	TradgPty  Party40Choice       `xml:"TradgPty,omitempty" json:",omitempty"`
	Prtry     []ProprietaryParty5 `xml:"Prtry,omitempty" json:",omitempty"`
}

type TransactionPrice4Choice struct {
	DealPric Price7              `xml:"DealPric"`
	Prtry    []ProprietaryPrice2 `xml:"Prtry"`
}

type TransactionQuantities3Choice struct {
	Qty                FinancialInstrumentQuantity1Choice `xml:"Qty"`
	OrgnlAndCurFaceAmt OriginalAndCurrentQuantities1      `xml:"OrgnlAndCurFaceAmt"`
	Prtry              ProprietaryQuantity1               `xml:"Prtry"`
}

type TransactionReferences6 struct {
	MsgId             common.Max35Text        `xml:"MsgId,omitempty" json:",omitempty"`
	AcctSvcrRef       common.Max35Text        `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	PmtInfId          common.Max35Text        `xml:"PmtInfId,omitempty" json:",omitempty"`
	InstrId           common.Max35Text        `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId        common.Max35Text        `xml:"EndToEndId,omitempty" json:",omitempty"`
	UETR              common.UUIDv4Identifier `xml:"UETR,omitempty" json:",omitempty"`
	TxId              common.Max35Text        `xml:"TxId,omitempty" json:",omitempty"`
	MndtId            common.Max35Text        `xml:"MndtId,omitempty" json:",omitempty"`
	ChqNb             common.Max35Text        `xml:"ChqNb,omitempty" json:",omitempty"`
	ClrSysRef         common.Max35Text        `xml:"ClrSysRef,omitempty" json:",omitempty"`
	AcctOwnrTxId      common.Max35Text        `xml:"AcctOwnrTxId,omitempty" json:",omitempty"`
	AcctSvcrTxId      common.Max35Text        `xml:"AcctSvcrTxId,omitempty" json:",omitempty"`
	MktInfrstrctrTxId common.Max35Text        `xml:"MktInfrstrctrTxId,omitempty" json:",omitempty"`
	PrcgId            common.Max35Text        `xml:"PrcgId,omitempty" json:",omitempty"`
	Prtry             []ProprietaryReference1 `xml:"Prtry,omitempty" json:",omitempty"`
}

type YieldedOrValueType1Choice struct {
	Yldd  bool                `xml:"Yldd"`
	ValTp PriceValueType1Code `xml:"ValTp"`
}

type AccountStatement9 struct {
	Id           common.Max35Text          `xml:"Id"`
	StmtPgntn    Pagination1               `xml:"StmtPgntn,omitempty" json:",omitempty"`
	ElctrncSeqNb float64                   `xml:"ElctrncSeqNb,omitempty" json:",omitempty"`
	RptgSeq      SequenceRange1Choice      `xml:"RptgSeq,omitempty" json:",omitempty"`
	LglSeqNb     float64                   `xml:"LglSeqNb,omitempty" json:",omitempty"`
	CreDtTm      common.ISODateTime        `xml:"CreDtTm,omitempty" json:",omitempty"`
	FrToDt       DateTimePeriod1           `xml:"FrToDt,omitempty" json:",omitempty"`
	CpyDplctInd  common.CopyDuplicate1Code `xml:"CpyDplctInd,omitempty" json:",omitempty"`
	RptgSrc      ReportingSource1Choice    `xml:"RptgSrc,omitempty" json:",omitempty"`
	Acct         CashAccount39             `xml:"Acct"`
	RltdAcct     CashAccount38             `xml:"RltdAcct,omitempty" json:",omitempty"`
	Intrst       []AccountInterest4        `xml:"Intrst,omitempty" json:",omitempty"`
	Bal          []CashBalance8            `xml:"Bal"`
	TxsSummry    TotalTransactions6        `xml:"TxsSummry,omitempty" json:",omitempty"`
	Ntry         []ReportEntry10           `xml:"Ntry,omitempty" json:",omitempty"`
	AddtlStmtInf common.Max500Text         `xml:"AddtlStmtInf,omitempty" json:",omitempty"`
}

type BankToCustomerStatementV08 struct {
	GrpHdr      GroupHeader81        `xml:"GrpHdr"`
	Stmt        []AccountStatement9  `xml:"Stmt"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type AccountNotification17 struct {
	Id             common.Max35Text          `xml:"Id"`
	NtfctnPgntn    Pagination1               `xml:"NtfctnPgntn,omitempty" json:",omitempty"`
	ElctrncSeqNb   float64                   `xml:"ElctrncSeqNb,omitempty" json:",omitempty"`
	RptgSeq        SequenceRange1Choice      `xml:"RptgSeq,omitempty" json:",omitempty"`
	LglSeqNb       float64                   `xml:"LglSeqNb,omitempty" json:",omitempty"`
	CreDtTm        common.ISODateTime        `xml:"CreDtTm,omitempty" json:",omitempty"`
	FrToDt         DateTimePeriod1           `xml:"FrToDt,omitempty" json:",omitempty"`
	CpyDplctInd    common.CopyDuplicate1Code `xml:"CpyDplctInd,omitempty" json:",omitempty"`
	RptgSrc        ReportingSource1Choice    `xml:"RptgSrc,omitempty" json:",omitempty"`
	Acct           CashAccount39             `xml:"Acct"`
	RltdAcct       CashAccount38             `xml:"RltdAcct,omitempty" json:",omitempty"`
	Intrst         []AccountInterest4        `xml:"Intrst,omitempty" json:",omitempty"`
	TxsSummry      TotalTransactions6        `xml:"TxsSummry,omitempty" json:",omitempty"`
	Ntry           []ReportEntry10           `xml:"Ntry,omitempty" json:",omitempty"`
	AddtlNtfctnInf common.Max500Text         `xml:"AddtlNtfctnInf,omitempty" json:",omitempty"`
}

type BankToCustomerDebitCreditNotificationV08 struct {
	GrpHdr      GroupHeader81           `xml:"GrpHdr"`
	Ntfctn      []AccountNotification17 `xml:"Ntfctn"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}
