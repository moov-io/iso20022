// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

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

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
}

func (r DatePeriod2Choice) Validate() error {
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

type GetStandingOrderV03 struct {
	MsgHdr        MessageHeader4       `xml:"MsgHdr"`
	StgOrdrQryDef *StandingOrderQuery3 `xml:"StgOrdrQryDef,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r GetStandingOrderV03) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader4 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   *RequestType3Choice `xml:"ReqTp,omitempty" json:",omitempty"`
}

func (r MessageHeader4) Validate() error {
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

type RequestType3Choice struct {
	Cd    StandingOrderQueryType1Code `xml:"Cd"`
	Prtry GenericIdentification1      `xml:"Prtry"`
}

func (r RequestType3Choice) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderCriteria3 struct {
	NewQryNm *common.Max35Text              `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []StandingOrderSearchCriteria3 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  *StandingOrderReturnCriteria1  `xml:"RtrCrit,omitempty" json:",omitempty"`
}

func (r StandingOrderCriteria3) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderCriteria3Choice struct {
	QryNm   common.Max35Text       `xml:"QryNm"`
	NewCrit StandingOrderCriteria3 `xml:"NewCrit"`
}

func (r StandingOrderCriteria3Choice) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderQuery3 struct {
	QryTp       *QueryType2Code               `xml:"QryTp,omitempty" json:",omitempty"`
	StgOrdrCrit *StandingOrderCriteria3Choice `xml:"StgOrdrCrit,omitempty" json:",omitempty"`
}

func (r StandingOrderQuery3) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderReturnCriteria1 struct {
	StgOrdrIdInd    bool `xml:"StgOrdrIdInd,omitempty" json:",omitempty"`
	TpInd           bool `xml:"TpInd,omitempty" json:",omitempty"`
	SysMmbInd       bool `xml:"SysMmbInd,omitempty" json:",omitempty"`
	RspnsblPtyInd   bool `xml:"RspnsblPtyInd,omitempty" json:",omitempty"`
	CcyInd          bool `xml:"CcyInd,omitempty" json:",omitempty"`
	DbtrAcctInd     bool `xml:"DbtrAcctInd,omitempty" json:",omitempty"`
	CdtrAcctInd     bool `xml:"CdtrAcctInd,omitempty" json:",omitempty"`
	AssoctdPoolAcct bool `xml:"AssoctdPoolAcct,omitempty" json:",omitempty"`
	FrqcyInd        bool `xml:"FrqcyInd,omitempty" json:",omitempty"`
	ExctnTpInd      bool `xml:"ExctnTpInd,omitempty" json:",omitempty"`
	VldtyFrInd      bool `xml:"VldtyFrInd,omitempty" json:",omitempty"`
	VldToInd        bool `xml:"VldToInd,omitempty" json:",omitempty"`
	LkSetIdInd      bool `xml:"LkSetIdInd,omitempty" json:",omitempty"`
	LkSetOrdrIdInd  bool `xml:"LkSetOrdrIdInd,omitempty" json:",omitempty"`
	LkSetOrdrSeqInd bool `xml:"LkSetOrdrSeqInd,omitempty" json:",omitempty"`
	TtlAmtInd       bool `xml:"TtlAmtInd,omitempty" json:",omitempty"`
	ZeroSweepInd    bool `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
}

func (r StandingOrderReturnCriteria1) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderSearchCriteria3 struct {
	KeyAttrbtsInd   bool                                          `xml:"KeyAttrbtsInd,omitempty" json:",omitempty"`
	StgOrdrId       *common.Max35Text                             `xml:"StgOrdrId,omitempty" json:",omitempty"`
	Tp              *StandingOrderType1Choice                     `xml:"Tp,omitempty" json:",omitempty"`
	Acct            *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	Ccy             *common.ActiveCurrencyCode                    `xml:"Ccy,omitempty" json:",omitempty"`
	VldtyPrd        *DatePeriod2Choice                            `xml:"VldtyPrd,omitempty" json:",omitempty"`
	SysMmb          *BranchAndFinancialInstitutionIdentification6 `xml:"SysMmb,omitempty" json:",omitempty"`
	RspnsblPty      *BranchAndFinancialInstitutionIdentification6 `xml:"RspnsblPty,omitempty" json:",omitempty"`
	AssoctdPoolAcct *AccountIdentification4Choice                 `xml:"AssoctdPoolAcct,omitempty" json:",omitempty"`
	LkSetId         *common.Max35Text                             `xml:"LkSetId,omitempty" json:",omitempty"`
	LkSetOrdrId     *common.Max35Text                             `xml:"LkSetOrdrId,omitempty" json:",omitempty"`
	LkSetOrdrSeq    float64                                       `xml:"LkSetOrdrSeq,omitempty" json:",omitempty"`
	ZeroSweepInd    bool                                          `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
}

func (r StandingOrderSearchCriteria3) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderType1Choice struct {
	Cd    StandingOrderType1Code `xml:"Cd"`
	Prtry GenericIdentification1 `xml:"Prtry"`
}

func (r StandingOrderType1Choice) Validate() error {
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

type DeleteStandingOrderV03 struct {
	MsgHdr      MessageHeader1            `xml:"MsgHdr"`
	StgOrdrDtls StandingOrderOrAll2Choice `xml:"StgOrdrDtls"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DeleteStandingOrderV03) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r MessageHeader1) Validate() error {
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

type StandingOrderIdentification5 struct {
	Acct     CashAccount38                                 `xml:"Acct"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

func (r StandingOrderIdentification5) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderOrAll2Choice struct {
	StgOrdr     []StandingOrderIdentification4 `xml:"StgOrdr" json:",omitempty"`
	AllStgOrdrs []StandingOrderIdentification5 `xml:"AllStgOrdrs" json:",omitempty"`
}

func (r StandingOrderOrAll2Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountTax1 struct {
	ClctnMtd   BillingTaxCalculationMethod1Code `xml:"ClctnMtd"`
	Rgn        *common.Max40Text                `xml:"Rgn,omitempty" json:",omitempty"`
	NonResCtry *ResidenceLocation1Choice        `xml:"NonResCtry,omitempty" json:",omitempty"`
}

func (r AccountTax1) Validate() error {
	return utils.Validate(&r)
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type AmountAndDirection34 struct {
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	Sgn bool                              `xml:"Sgn"`
}

func (r AmountAndDirection34) Validate() error {
	return utils.Validate(&r)
}

type BalanceAdjustment1 struct {
	Tp                BalanceAdjustmentType1Code `xml:"Tp"`
	Desc              common.Max105Text          `xml:"Desc"`
	BalAmt            AmountAndDirection34       `xml:"BalAmt"`
	AvrgAmt           *AmountAndDirection34      `xml:"AvrgAmt,omitempty" json:",omitempty"`
	ErrDt             *common.ISODate            `xml:"ErrDt,omitempty" json:",omitempty"`
	PstngDt           *common.ISODate            `xml:"PstngDt"`
	Days              float64                    `xml:"Days,omitempty" json:",omitempty"`
	EarngsAdjstmntAmt *AmountAndDirection34      `xml:"EarngsAdjstmntAmt,omitempty" json:",omitempty"`
}

func (r BalanceAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type BankServicesBillingStatementV03 struct {
	RptHdr      ReportHeader6     `xml:"RptHdr"`
	BllgStmtGrp []StatementGroup3 `xml:"BllgStmtGrp" json:",omitempty"`
}

func (r BankServicesBillingStatementV03) Validate() error {
	return utils.Validate(&r)
}

type BankTransactionCodeStructure4 struct {
	Domn  *BankTransactionCodeStructure5            `xml:"Domn,omitempty" json:",omitempty"`
	Prtry *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r BankTransactionCodeStructure4) Validate() error {
	return utils.Validate(&r)
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"Fmly"`
}

func (r BankTransactionCodeStructure5) Validate() error {
	return utils.Validate(&r)
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"SubFmlyCd"`
}

func (r BankTransactionCodeStructure6) Validate() error {
	return utils.Validate(&r)
}

type BillingBalance1 struct {
	Tp    BillingBalanceType1Choice `xml:"Tp"`
	Val   AmountAndDirection34      `xml:"Val"`
	CcyTp *BillingCurrencyType1Code `xml:"CcyTp,omitempty" json:",omitempty"`
}

func (r BillingBalance1) Validate() error {
	return utils.Validate(&r)
}

type BillingBalanceType1Choice struct {
	Cd    ExternalBillingBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r BillingBalanceType1Choice) Validate() error {
	return utils.Validate(&r)
}

type BillingCompensation1 struct {
	Tp    BillingCompensationType1Choice `xml:"Tp"`
	Val   AmountAndDirection34           `xml:"Val"`
	CcyTp *BillingCurrencyType2Code      `xml:"CcyTp,omitempty" json:",omitempty"`
}

func (r BillingCompensation1) Validate() error {
	return utils.Validate(&r)
}

type BillingCompensationType1Choice struct {
	Cd    ExternalBillingCompensationType1Code `xml:"Cd"`
	Prtry common.Max35Text                     `xml:"Prtry"`
}

func (r BillingCompensationType1Choice) Validate() error {
	return utils.Validate(&r)
}

type BillingMethod1 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"SvcTax"`
	TtlChrg       BillingServicesAmount2 `xml:"TtlChrg"`
	TaxId         []BillingServicesTax1  `xml:"TaxId" json:",omitempty"`
}

func (r BillingMethod1) Validate() error {
	return utils.Validate(&r)
}

type BillingMethod1Choice struct {
	MtdA BillingMethod1 `xml:"MtdA"`
	MtdB BillingMethod2 `xml:"MtdB"`
	MtdD BillingMethod3 `xml:"MtdD"`
}

func (r BillingMethod1Choice) Validate() error {
	return utils.Validate(&r)
}

type BillingMethod2 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"SvcTax"`
	TaxId         []BillingServicesTax1  `xml:"TaxId" json:",omitempty"`
}

func (r BillingMethod2) Validate() error {
	return utils.Validate(&r)
}

type BillingMethod3 struct {
	SvcTaxPricAmt AmountAndDirection34  `xml:"SvcTaxPricAmt"`
	TaxId         []BillingServicesTax2 `xml:"TaxId" json:",omitempty"`
}

func (r BillingMethod3) Validate() error {
	return utils.Validate(&r)
}

type BillingMethod4 struct {
	SvcDtl   []BillingServiceParameters2 `xml:"SvcDtl" json:",omitempty"`
	TaxClctn TaxCalculation1             `xml:"TaxClctn"`
}

func (r BillingMethod4) Validate() error {
	return utils.Validate(&r)
}

type BillingPrice1 struct {
	Ccy      *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	UnitPric *AmountAndDirection34                `xml:"UnitPric,omitempty" json:",omitempty"`
	Mtd      *BillingChargeMethod1Code            `xml:"Mtd,omitempty" json:",omitempty"`
	Rule     *common.Max20Text                    `xml:"Rule,omitempty" json:",omitempty"`
}

func (r BillingPrice1) Validate() error {
	return utils.Validate(&r)
}

type BillingRate1 struct {
	Id        BillingRateIdentification1Choice `xml:"Id"`
	Val       float64                          `xml:"Val"`
	DaysInPrd float64                          `xml:"DaysInPrd,omitempty" json:",omitempty"`
	DaysInYr  float64                          `xml:"DaysInYr,omitempty" json:",omitempty"`
}

func (r BillingRate1) Validate() error {
	return utils.Validate(&r)
}

type BillingRateIdentification1Choice struct {
	Cd    ExternalBillingRateIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                       `xml:"Prtry"`
}

func (r BillingRateIdentification1Choice) Validate() error {
	return utils.Validate(&r)
}

type BillingService2 struct {
	SvcDtl            BillingServiceParameters3 `xml:"SvcDtl"`
	Pric              *BillingPrice1            `xml:"Pric,omitempty" json:",omitempty"`
	PmtMtd            ServicePaymentMethod1Code `xml:"PmtMtd"`
	OrgnlChrgPric     AmountAndDirection34      `xml:"OrgnlChrgPric"`
	OrgnlChrgSttlmAmt *AmountAndDirection34     `xml:"OrgnlChrgSttlmAmt,omitempty" json:",omitempty"`
	BalReqrdAcctAmt   *AmountAndDirection34     `xml:"BalReqrdAcctAmt,omitempty" json:",omitempty"`
	TaxDsgnt          ServiceTaxDesignation1    `xml:"TaxDsgnt"`
	TaxClctn          *BillingMethod1Choice     `xml:"TaxClctn,omitempty" json:",omitempty"`
}

func (r BillingService2) Validate() error {
	return utils.Validate(&r)
}

type BillingServiceAdjustment1 struct {
	Tp           ServiceAdjustmentType1Code        `xml:"Tp"`
	Desc         common.Max140Text                 `xml:"Desc"`
	Amt          AmountAndDirection34              `xml:"Amt"`
	BalReqrdAmt  *AmountAndDirection34             `xml:"BalReqrdAmt,omitempty" json:",omitempty"`
	ErrDt        *common.ISODate                   `xml:"ErrDt,omitempty" json:",omitempty"`
	AdjstmntId   *common.Max35Text                 `xml:"AdjstmntId,omitempty" json:",omitempty"`
	SubSvc       *BillingSubServiceIdentification1 `xml:"SubSvc,omitempty" json:",omitempty"`
	PricChng     *AmountAndDirection34             `xml:"PricChng,omitempty" json:",omitempty"`
	OrgnlPric    *AmountAndDirection34             `xml:"OrgnlPric,omitempty" json:",omitempty"`
	NewPric      *AmountAndDirection34             `xml:"NewPric,omitempty" json:",omitempty"`
	VolChng      float64                           `xml:"VolChng,omitempty" json:",omitempty"`
	OrgnlVol     float64                           `xml:"OrgnlVol,omitempty" json:",omitempty"`
	NewVol       float64                           `xml:"NewVol,omitempty" json:",omitempty"`
	OrgnlChrgAmt *AmountAndDirection34             `xml:"OrgnlChrgAmt,omitempty" json:",omitempty"`
	NewChrgAmt   *AmountAndDirection34             `xml:"NewChrgAmt,omitempty" json:",omitempty"`
}

func (r BillingServiceAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type BillingServiceCommonIdentification1 struct {
	Issr common.Max6Text `xml:"Issr"`
	Id   common.Max8Text `xml:"Id"`
}

func (r BillingServiceCommonIdentification1) Validate() error {
	return utils.Validate(&r)
}

type BillingServiceIdentification2 struct {
	Id     common.Max35Text                  `xml:"Id"`
	SubSvc *BillingSubServiceIdentification1 `xml:"SubSvc,omitempty" json:",omitempty"`
	Desc   common.Max70Text                  `xml:"Desc"`
}

func (r BillingServiceIdentification2) Validate() error {
	return utils.Validate(&r)
}

type BillingServiceIdentification3 struct {
	Id     common.Max35Text                     `xml:"Id"`
	SubSvc *BillingSubServiceIdentification1    `xml:"SubSvc,omitempty" json:",omitempty"`
	Desc   common.Max70Text                     `xml:"Desc"`
	CmonCd *BillingServiceCommonIdentification1 `xml:"CmonCd,omitempty" json:",omitempty"`
	BkTxCd *BankTransactionCodeStructure4       `xml:"BkTxCd,omitempty" json:",omitempty"`
	SvcTp  *common.Max12Text                    `xml:"SvcTp,omitempty" json:",omitempty"`
}

func (r BillingServiceIdentification3) Validate() error {
	return utils.Validate(&r)
}

type BillingServiceParameters2 struct {
	BkSvc      BillingServiceIdentification2 `xml:"BkSvc"`
	Vol        float64                       `xml:"Vol,omitempty" json:",omitempty"`
	UnitPric   *AmountAndDirection34         `xml:"UnitPric,omitempty" json:",omitempty"`
	SvcChrgAmt *AmountAndDirection34         `xml:"SvcChrgAmt"`
}

func (r BillingServiceParameters2) Validate() error {
	return utils.Validate(&r)
}

type BillingServiceParameters3 struct {
	BkSvc BillingServiceIdentification3 `xml:"BkSvc"`
	Vol   float64                       `xml:"Vol,omitempty" json:",omitempty"`
}

func (r BillingServiceParameters3) Validate() error {
	return utils.Validate(&r)
}

type BillingServicesAmount1 struct {
	HstAmt   AmountAndDirection34  `xml:"HstAmt"`
	PricgAmt *AmountAndDirection34 `xml:"PricgAmt,omitempty" json:",omitempty"`
}

func (r BillingServicesAmount1) Validate() error {
	return utils.Validate(&r)
}

type BillingServicesAmount2 struct {
	HstAmt   AmountAndDirection34  `xml:"HstAmt"`
	SttlmAmt *AmountAndDirection34 `xml:"SttlmAmt,omitempty" json:",omitempty"`
	PricgAmt *AmountAndDirection34 `xml:"PricgAmt,omitempty" json:",omitempty"`
}

func (r BillingServicesAmount2) Validate() error {
	return utils.Validate(&r)
}

type BillingServicesAmount3 struct {
	SrcAmt AmountAndDirection34 `xml:"SrcAmt"`
	HstAmt AmountAndDirection34 `xml:"HstAmt"`
}

func (r BillingServicesAmount3) Validate() error {
	return utils.Validate(&r)
}

type BillingServicesTax1 struct {
	Nb       common.Max35Text      `xml:"Nb"`
	Desc     *common.Max40Text     `xml:"Desc,omitempty" json:",omitempty"`
	Rate     float64               `xml:"Rate"`
	HstAmt   AmountAndDirection34  `xml:"HstAmt"`
	PricgAmt *AmountAndDirection34 `xml:"PricgAmt,omitempty" json:",omitempty"`
}

func (r BillingServicesTax1) Validate() error {
	return utils.Validate(&r)
}

type BillingServicesTax2 struct {
	Nb       common.Max35Text     `xml:"Nb"`
	Desc     *common.Max40Text    `xml:"Desc,omitempty" json:",omitempty"`
	Rate     float64              `xml:"Rate"`
	PricgAmt AmountAndDirection34 `xml:"PricgAmt"`
}

func (r BillingServicesTax2) Validate() error {
	return utils.Validate(&r)
}

type BillingServicesTax3 struct {
	Nb        common.Max35Text     `xml:"Nb"`
	Desc      *common.Max40Text    `xml:"Desc,omitempty" json:",omitempty"`
	Rate      float64              `xml:"Rate"`
	TtlTaxAmt AmountAndDirection34 `xml:"TtlTaxAmt"`
}

func (r BillingServicesTax3) Validate() error {
	return utils.Validate(&r)
}

type BillingStatement3 struct {
	StmtId      common.Max35Text            `xml:"StmtId"`
	FrToDt      DatePeriod1                 `xml:"FrToDt"`
	CreDtTm     common.ISODateTime          `xml:"CreDtTm"`
	Sts         BillingStatementStatus1Code `xml:"Sts"`
	AcctChrtcs  CashAccountCharacteristics3 `xml:"AcctChrtcs"`
	RateData    []BillingRate1              `xml:"RateData,omitempty" json:",omitempty"`
	CcyXchg     []CurrencyExchange6         `xml:"CcyXchg,omitempty" json:",omitempty"`
	Bal         []BillingBalance1           `xml:"Bal,omitempty" json:",omitempty"`
	Compstn     []BillingCompensation1      `xml:"Compstn,omitempty" json:",omitempty"`
	Svc         []BillingService2           `xml:"Svc,omitempty" json:",omitempty"`
	TaxRgn      []BillingTaxRegion2         `xml:"TaxRgn,omitempty" json:",omitempty"`
	BalAdjstmnt []BalanceAdjustment1        `xml:"BalAdjstmnt,omitempty" json:",omitempty"`
	SvcAdjstmnt []BillingServiceAdjustment1 `xml:"SvcAdjstmnt,omitempty" json:",omitempty"`
}

func (r BillingStatement3) Validate() error {
	return utils.Validate(&r)
}

type BillingSubServiceIdentification1 struct {
	Issr BillingSubServiceQualifier1Choice `xml:"Issr"`
	Id   common.Max35Text                  `xml:"Id"`
}

func (r BillingSubServiceIdentification1) Validate() error {
	return utils.Validate(&r)
}

type BillingSubServiceQualifier1Choice struct {
	Cd    BillingSubServiceQualifier1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r BillingSubServiceQualifier1Choice) Validate() error {
	return utils.Validate(&r)
}

type BillingTaxIdentification2 struct {
	VATRegnNb *common.Max35Text `xml:"VATRegnNb,omitempty" json:",omitempty"`
	TaxRegnNb *common.Max35Text `xml:"TaxRegnNb,omitempty" json:",omitempty"`
	TaxCtct   *Contact4         `xml:"TaxCtct,omitempty" json:",omitempty"`
}

func (r BillingTaxIdentification2) Validate() error {
	return utils.Validate(&r)
}

type BillingTaxRegion2 struct {
	RgnNb       common.Max40Text           `xml:"RgnNb"`
	RgnNm       common.Max40Text           `xml:"RgnNm"`
	CstmrTaxId  common.Max40Text           `xml:"CstmrTaxId"`
	PtDt        *common.ISODate            `xml:"PtDt,omitempty" json:",omitempty"`
	SndgFI      *BillingTaxIdentification2 `xml:"SndgFI,omitempty" json:",omitempty"`
	InvcNb      *common.Max40Text          `xml:"InvcNb,omitempty" json:",omitempty"`
	MtdC        *BillingMethod4            `xml:"MtdC,omitempty" json:",omitempty"`
	SttlmAmt    AmountAndDirection34       `xml:"SttlmAmt"`
	TaxDueToRgn AmountAndDirection34       `xml:"TaxDueToRgn"`
}

func (r BillingTaxRegion2) Validate() error {
	return utils.Validate(&r)
}

type CashAccountCharacteristics3 struct {
	AcctLvl      AccountLevel2Code                             `xml:"AcctLvl"`
	CshAcct      CashAccount38                                 `xml:"CshAcct"`
	AcctSvcr     *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	PrntAcct     *ParentCashAccount3                           `xml:"PrntAcct,omitempty" json:",omitempty"`
	CompstnMtd   CompensationMethod1Code                       `xml:"CompstnMtd"`
	DbtAcct      *AccountIdentification4Choice                 `xml:"DbtAcct,omitempty" json:",omitempty"`
	DelydDbtDt   *common.ISODate                               `xml:"DelydDbtDt,omitempty" json:",omitempty"`
	SttlmAdvc    *common.Max105Text                            `xml:"SttlmAdvc,omitempty" json:",omitempty"`
	AcctBalCcyCd common.ActiveOrHistoricCurrencyCode           `xml:"AcctBalCcyCd"`
	SttlmCcyCd   *common.ActiveOrHistoricCurrencyCode          `xml:"SttlmCcyCd,omitempty" json:",omitempty"`
	HstCcyCd     *common.ActiveOrHistoricCurrencyCode          `xml:"HstCcyCd,omitempty" json:",omitempty"`
	Tax          *AccountTax1                                  `xml:"Tax,omitempty" json:",omitempty"`
	AcctSvcrCtct Contact4                                      `xml:"AcctSvcrCtct"`
}

func (r CashAccountCharacteristics3) Validate() error {
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

type CurrencyExchange6 struct {
	SrcCcy   common.ActiveOrHistoricCurrencyCode  `xml:"SrcCcy"`
	TrgtCcy  common.ActiveOrHistoricCurrencyCode  `xml:"TrgtCcy"`
	XchgRate float64                              `xml:"XchgRate"`
	Desc     *common.Max40Text                    `xml:"Desc,omitempty" json:",omitempty"`
	UnitCcy  *common.ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty" json:",omitempty"`
	Cmnts    *common.Max70Text                    `xml:"Cmnts,omitempty" json:",omitempty"`
	QtnDt    *common.ISODateTime                  `xml:"QtnDt,omitempty" json:",omitempty"`
}

func (r CurrencyExchange6) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod1 struct {
	FrDt *common.ISODate `xml:"FrDt,omitempty" json:",omitempty"`
	ToDt common.ISODate  `xml:"ToDt"`
}

func (r DatePeriod1) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification19 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification19) Validate() error {
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

type Pagination1 struct {
	PgNb      common.Max5NumericText `xml:"PgNb"`
	LastPgInd bool                   `xml:"LastPgInd"`
}

func (r Pagination1) Validate() error {
	return utils.Validate(&r)
}

type ParentCashAccount3 struct {
	Lvl  *AccountLevel1Code                            `xml:"Lvl,omitempty" json:",omitempty"`
	Id   CashAccount38                                 `xml:"Id"`
	Svcr *BranchAndFinancialInstitutionIdentification6 `xml:"Svcr,omitempty" json:",omitempty"`
}

func (r ParentCashAccount3) Validate() error {
	return utils.Validate(&r)
}

type Party43Choice struct {
	OrgId OrganisationIdentification29         `xml:"OrgId"`
	FIId  FinancialInstitutionIdentification19 `xml:"FIId"`
}

func (r Party43Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification138 struct {
	Nm        common.Max140Text   `xml:"Nm"`
	LglNm     *common.Max140Text  `xml:"LglNm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        Party43Choice       `xml:"Id"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification138) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   common.Max35Text  `xml:"Cd"`
	Issr *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r ProprietaryBankTransactionCodeStructure1) Validate() error {
	return utils.Validate(&r)
}

type ReportHeader6 struct {
	RptId    common.Max35Text `xml:"RptId"`
	MsgPgntn *Pagination1     `xml:"MsgPgntn,omitempty" json:",omitempty"`
}

func (r ReportHeader6) Validate() error {
	return utils.Validate(&r)
}

type ResidenceLocation1Choice struct {
	Ctry common.CountryCode `xml:"Ctry"`
	Area common.Max35Text   `xml:"Area"`
}

func (r ResidenceLocation1Choice) Validate() error {
	return utils.Validate(&r)
}

type ServiceTaxDesignation1 struct {
	Cd     ServiceTaxDesignation1Code `xml:"Cd"`
	Rgn    *common.Max35Text          `xml:"Rgn,omitempty" json:",omitempty"`
	TaxRsn []TaxReason1               `xml:"TaxRsn,omitempty" json:",omitempty"`
}

func (r ServiceTaxDesignation1) Validate() error {
	return utils.Validate(&r)
}

type StatementGroup3 struct {
	GrpId        common.Max35Text       `xml:"GrpId"`
	Sndr         PartyIdentification138 `xml:"Sndr"`
	SndrIndvCtct []Contact4             `xml:"SndrIndvCtct,omitempty" json:",omitempty"`
	Rcvr         PartyIdentification138 `xml:"Rcvr"`
	RcvrIndvCtct []Contact4             `xml:"RcvrIndvCtct,omitempty" json:",omitempty"`
	BllgStmt     []BillingStatement3    `xml:"BllgStmt"`
}

func (r StatementGroup3) Validate() error {
	return utils.Validate(&r)
}

type TaxCalculation1 struct {
	HstCcy                common.ActiveOrHistoricCurrencyCode `xml:"HstCcy"`
	TaxblSvcChrgConvs     []BillingServicesAmount3            `xml:"TaxblSvcChrgConvs"`
	TtlTaxblSvcChrgHstAmt AmountAndDirection34                `xml:"TtlTaxblSvcChrgHstAmt"`
	TaxId                 []BillingServicesTax3               `xml:"TaxId"`
	TtlTax                AmountAndDirection34                `xml:"TtlTax"`
}

func (r TaxCalculation1) Validate() error {
	return utils.Validate(&r)
}

type TaxReason1 struct {
	Cd     common.Max10Text  `xml:"Cd"`
	Expltn common.Max105Text `xml:"Expltn"`
}

func (r TaxReason1) Validate() error {
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

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
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

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericPersonIdentification1) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification8 struct {
	AnyBIC *common.AnyBICIdentifier             `xml:"AnyBIC,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification8) Validate() error {
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

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification5) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress6 struct {
	AdrTp       *common.AddressType2Code `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text        `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text        `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text        `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text        `xml:"BldgNb,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text        `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text        `xml:"TwnNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text        `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode      `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text       `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress6) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryData3 struct {
	Item string `xml:",any"`
}

func (r ProprietaryData3) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryData4 struct {
	Tp   common.Max35Text `xml:"Tp"`
	Data ProprietaryData3 `xml:"Data"`
}

func (r ProprietaryData4) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryFormatInvestigationV03 struct {
	Assgnmt     CaseAssignment3      `xml:"Assgnmt"`
	Case        Case3                `xml:"Case"`
	PrtryData   ProprietaryData4     `xml:"PrtryData"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ProprietaryFormatInvestigationV03) Validate() error {
	return utils.Validate(&r)
}
