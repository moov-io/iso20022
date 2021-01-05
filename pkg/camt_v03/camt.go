// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
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

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                    `xml:"MmbId"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type DatePeriod2Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
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

type GetStandingOrderV03 struct {
	MsgHdr        MessageHeader4       `xml:"MsgHdr"`
	StgOrdrQryDef StandingOrderQuery3  `xml:"StgOrdrQryDef,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MessageHeader4 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   RequestType3Choice `xml:"ReqTp,omitempty" json:",omitempty"`
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

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text      `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type RequestType3Choice struct {
	Cd    StandingOrderQueryType1Code `xml:"Cd"`
	Prtry GenericIdentification1      `xml:"Prtry"`
}

type StandingOrderCriteria3 struct {
	NewQryNm common.Max35Text               `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []StandingOrderSearchCriteria3 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  StandingOrderReturnCriteria1   `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type StandingOrderCriteria3Choice struct {
	QryNm   common.Max35Text       `xml:"QryNm"`
	NewCrit StandingOrderCriteria3 `xml:"NewCrit"`
}

type StandingOrderQuery3 struct {
	QryTp       QueryType2Code               `xml:"QryTp,omitempty" json:",omitempty"`
	StgOrdrCrit StandingOrderCriteria3Choice `xml:"StgOrdrCrit,omitempty" json:",omitempty"`
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

type StandingOrderSearchCriteria3 struct {
	KeyAttrbtsInd   bool                                         `xml:"KeyAttrbtsInd,omitempty" json:",omitempty"`
	StgOrdrId       common.Max35Text                             `xml:"StgOrdrId,omitempty" json:",omitempty"`
	Tp              StandingOrderType1Choice                     `xml:"Tp,omitempty" json:",omitempty"`
	Acct            CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	Ccy             common.ActiveCurrencyCode                    `xml:"Ccy,omitempty" json:",omitempty"`
	VldtyPrd        DatePeriod2Choice                            `xml:"VldtyPrd,omitempty" json:",omitempty"`
	SysMmb          BranchAndFinancialInstitutionIdentification6 `xml:"SysMmb,omitempty" json:",omitempty"`
	RspnsblPty      BranchAndFinancialInstitutionIdentification6 `xml:"RspnsblPty,omitempty" json:",omitempty"`
	AssoctdPoolAcct AccountIdentification4Choice                 `xml:"AssoctdPoolAcct,omitempty" json:",omitempty"`
	LkSetId         common.Max35Text                             `xml:"LkSetId,omitempty" json:",omitempty"`
	LkSetOrdrId     common.Max35Text                             `xml:"LkSetOrdrId,omitempty" json:",omitempty"`
	LkSetOrdrSeq    float64                                      `xml:"LkSetOrdrSeq,omitempty" json:",omitempty"`
	ZeroSweepInd    bool                                         `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
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

type DeleteStandingOrderV03 struct {
	MsgHdr      MessageHeader1            `xml:"MsgHdr"`
	StgOrdrDtls StandingOrderOrAll2Choice `xml:"StgOrdrDtls"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MessageHeader1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type StandingOrderIdentification4 struct {
	Id       common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	Acct     CashAccount38                                `xml:"Acct"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

type StandingOrderIdentification5 struct {
	Acct     CashAccount38                                `xml:"Acct"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

type StandingOrderOrAll2Choice struct {
	StgOrdr     []StandingOrderIdentification4 `xml:"StgOrdr"`
	AllStgOrdrs []StandingOrderIdentification5 `xml:"AllStgOrdrs"`
}

type AccountTax1 struct {
	ClctnMtd   BillingTaxCalculationMethod1Code `xml:"ClctnMtd"`
	Rgn        common.Max40Text                 `xml:"Rgn,omitempty" json:",omitempty"`
	NonResCtry ResidenceLocation1Choice         `xml:"NonResCtry,omitempty" json:",omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AmountAndDirection34 struct {
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	Sgn bool                              `xml:"Sgn"`
}

type BalanceAdjustment1 struct {
	Tp                BalanceAdjustmentType1Code `xml:"Tp"`
	Desc              common.Max105Text          `xml:"Desc"`
	BalAmt            AmountAndDirection34       `xml:"BalAmt"`
	AvrgAmt           AmountAndDirection34       `xml:"AvrgAmt,omitempty" json:",omitempty"`
	ErrDt             common.ISODate             `xml:"ErrDt,omitempty" json:",omitempty"`
	PstngDt           common.ISODate             `xml:"PstngDt"`
	Days              float64                    `xml:"Days,omitempty" json:",omitempty"`
	EarngsAdjstmntAmt AmountAndDirection34       `xml:"EarngsAdjstmntAmt,omitempty" json:",omitempty"`
}

type BankServicesBillingStatementV03 struct {
	RptHdr      ReportHeader6     `xml:"RptHdr"`
	BllgStmtGrp []StatementGroup3 `xml:"BllgStmtGrp"`
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

type BillingBalance1 struct {
	Tp    BillingBalanceType1Choice `xml:"Tp"`
	Val   AmountAndDirection34      `xml:"Val"`
	CcyTp BillingCurrencyType1Code  `xml:"CcyTp,omitempty" json:",omitempty"`
}

type BillingBalanceType1Choice struct {
	Cd    ExternalBillingBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type BillingCompensation1 struct {
	Tp    BillingCompensationType1Choice `xml:"Tp"`
	Val   AmountAndDirection34           `xml:"Val"`
	CcyTp BillingCurrencyType2Code       `xml:"CcyTp,omitempty" json:",omitempty"`
}

type BillingCompensationType1Choice struct {
	Cd    ExternalBillingCompensationType1Code `xml:"Cd"`
	Prtry common.Max35Text                     `xml:"Prtry"`
}

type BillingMethod1 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"SvcTax"`
	TtlChrg       BillingServicesAmount2 `xml:"TtlChrg"`
	TaxId         []BillingServicesTax1  `xml:"TaxId"`
}

type BillingMethod1Choice struct {
	MtdA BillingMethod1 `xml:"MtdA"`
	MtdB BillingMethod2 `xml:"MtdB"`
	MtdD BillingMethod3 `xml:"MtdD"`
}

type BillingMethod2 struct {
	SvcChrgHstAmt AmountAndDirection34   `xml:"SvcChrgHstAmt"`
	SvcTax        BillingServicesAmount1 `xml:"SvcTax"`
	TaxId         []BillingServicesTax1  `xml:"TaxId"`
}

type BillingMethod3 struct {
	SvcTaxPricAmt AmountAndDirection34  `xml:"SvcTaxPricAmt"`
	TaxId         []BillingServicesTax2 `xml:"TaxId"`
}

type BillingMethod4 struct {
	SvcDtl   []BillingServiceParameters2 `xml:"SvcDtl"`
	TaxClctn TaxCalculation1             `xml:"TaxClctn"`
}

type BillingPrice1 struct {
	Ccy      common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	UnitPric AmountAndDirection34                `xml:"UnitPric,omitempty" json:",omitempty"`
	Mtd      BillingChargeMethod1Code            `xml:"Mtd,omitempty" json:",omitempty"`
	Rule     common.Max20Text                    `xml:"Rule,omitempty" json:",omitempty"`
}

type BillingRate1 struct {
	Id        BillingRateIdentification1Choice `xml:"Id"`
	Val       float64                          `xml:"Val"`
	DaysInPrd float64                          `xml:"DaysInPrd,omitempty" json:",omitempty"`
	DaysInYr  float64                          `xml:"DaysInYr,omitempty" json:",omitempty"`
}

type BillingRateIdentification1Choice struct {
	Cd    ExternalBillingRateIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                       `xml:"Prtry"`
}

type BillingService2 struct {
	SvcDtl            BillingServiceParameters3 `xml:"SvcDtl"`
	Pric              BillingPrice1             `xml:"Pric,omitempty" json:",omitempty"`
	PmtMtd            ServicePaymentMethod1Code `xml:"PmtMtd"`
	OrgnlChrgPric     AmountAndDirection34      `xml:"OrgnlChrgPric"`
	OrgnlChrgSttlmAmt AmountAndDirection34      `xml:"OrgnlChrgSttlmAmt,omitempty" json:",omitempty"`
	BalReqrdAcctAmt   AmountAndDirection34      `xml:"BalReqrdAcctAmt,omitempty" json:",omitempty"`
	TaxDsgnt          ServiceTaxDesignation1    `xml:"TaxDsgnt"`
	TaxClctn          BillingMethod1Choice      `xml:"TaxClctn,omitempty" json:",omitempty"`
}

type BillingServiceAdjustment1 struct {
	Tp           ServiceAdjustmentType1Code       `xml:"Tp"`
	Desc         common.Max140Text                `xml:"Desc"`
	Amt          AmountAndDirection34             `xml:"Amt"`
	BalReqrdAmt  AmountAndDirection34             `xml:"BalReqrdAmt,omitempty" json:",omitempty"`
	ErrDt        common.ISODate                   `xml:"ErrDt,omitempty" json:",omitempty"`
	AdjstmntId   common.Max35Text                 `xml:"AdjstmntId,omitempty" json:",omitempty"`
	SubSvc       BillingSubServiceIdentification1 `xml:"SubSvc,omitempty" json:",omitempty"`
	PricChng     AmountAndDirection34             `xml:"PricChng,omitempty" json:",omitempty"`
	OrgnlPric    AmountAndDirection34             `xml:"OrgnlPric,omitempty" json:",omitempty"`
	NewPric      AmountAndDirection34             `xml:"NewPric,omitempty" json:",omitempty"`
	VolChng      float64                          `xml:"VolChng,omitempty" json:",omitempty"`
	OrgnlVol     float64                          `xml:"OrgnlVol,omitempty" json:",omitempty"`
	NewVol       float64                          `xml:"NewVol,omitempty" json:",omitempty"`
	OrgnlChrgAmt AmountAndDirection34             `xml:"OrgnlChrgAmt,omitempty" json:",omitempty"`
	NewChrgAmt   AmountAndDirection34             `xml:"NewChrgAmt,omitempty" json:",omitempty"`
}

type BillingServiceCommonIdentification1 struct {
	Issr common.Max6Text `xml:"Issr"`
	Id   common.Max8Text `xml:"Id"`
}

type BillingServiceIdentification2 struct {
	Id     common.Max35Text                 `xml:"Id"`
	SubSvc BillingSubServiceIdentification1 `xml:"SubSvc,omitempty" json:",omitempty"`
	Desc   common.Max70Text                 `xml:"Desc"`
}

type BillingServiceIdentification3 struct {
	Id     common.Max35Text                    `xml:"Id"`
	SubSvc BillingSubServiceIdentification1    `xml:"SubSvc,omitempty" json:",omitempty"`
	Desc   common.Max70Text                    `xml:"Desc"`
	CmonCd BillingServiceCommonIdentification1 `xml:"CmonCd,omitempty" json:",omitempty"`
	BkTxCd BankTransactionCodeStructure4       `xml:"BkTxCd,omitempty" json:",omitempty"`
	SvcTp  common.Max12Text                    `xml:"SvcTp,omitempty" json:",omitempty"`
}

type BillingServiceParameters2 struct {
	BkSvc      BillingServiceIdentification2 `xml:"BkSvc"`
	Vol        float64                       `xml:"Vol,omitempty" json:",omitempty"`
	UnitPric   AmountAndDirection34          `xml:"UnitPric,omitempty" json:",omitempty"`
	SvcChrgAmt AmountAndDirection34          `xml:"SvcChrgAmt"`
}

type BillingServiceParameters3 struct {
	BkSvc BillingServiceIdentification3 `xml:"BkSvc"`
	Vol   float64                       `xml:"Vol,omitempty" json:",omitempty"`
}

type BillingServicesAmount1 struct {
	HstAmt   AmountAndDirection34 `xml:"HstAmt"`
	PricgAmt AmountAndDirection34 `xml:"PricgAmt,omitempty" json:",omitempty"`
}

type BillingServicesAmount2 struct {
	HstAmt   AmountAndDirection34 `xml:"HstAmt"`
	SttlmAmt AmountAndDirection34 `xml:"SttlmAmt,omitempty" json:",omitempty"`
	PricgAmt AmountAndDirection34 `xml:"PricgAmt,omitempty" json:",omitempty"`
}

type BillingServicesAmount3 struct {
	SrcAmt AmountAndDirection34 `xml:"SrcAmt"`
	HstAmt AmountAndDirection34 `xml:"HstAmt"`
}

type BillingServicesTax1 struct {
	Nb       common.Max35Text     `xml:"Nb"`
	Desc     common.Max40Text     `xml:"Desc,omitempty" json:",omitempty"`
	Rate     float64              `xml:"Rate"`
	HstAmt   AmountAndDirection34 `xml:"HstAmt"`
	PricgAmt AmountAndDirection34 `xml:"PricgAmt,omitempty" json:",omitempty"`
}

type BillingServicesTax2 struct {
	Nb       common.Max35Text     `xml:"Nb"`
	Desc     common.Max40Text     `xml:"Desc,omitempty" json:",omitempty"`
	Rate     float64              `xml:"Rate"`
	PricgAmt AmountAndDirection34 `xml:"PricgAmt"`
}

type BillingServicesTax3 struct {
	Nb        common.Max35Text     `xml:"Nb"`
	Desc      common.Max40Text     `xml:"Desc,omitempty" json:",omitempty"`
	Rate      float64              `xml:"Rate"`
	TtlTaxAmt AmountAndDirection34 `xml:"TtlTaxAmt"`
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

type BillingSubServiceIdentification1 struct {
	Issr BillingSubServiceQualifier1Choice `xml:"Issr"`
	Id   common.Max35Text                  `xml:"Id"`
}

type BillingSubServiceQualifier1Choice struct {
	Cd    BillingSubServiceQualifier1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type BillingTaxIdentification2 struct {
	VATRegnNb common.Max35Text `xml:"VATRegnNb,omitempty" json:",omitempty"`
	TaxRegnNb common.Max35Text `xml:"TaxRegnNb,omitempty" json:",omitempty"`
	TaxCtct   Contact4         `xml:"TaxCtct,omitempty" json:",omitempty"`
}

type BillingTaxRegion2 struct {
	RgnNb       common.Max40Text          `xml:"RgnNb"`
	RgnNm       common.Max40Text          `xml:"RgnNm"`
	CstmrTaxId  common.Max40Text          `xml:"CstmrTaxId"`
	PtDt        common.ISODate            `xml:"PtDt,omitempty" json:",omitempty"`
	SndgFI      BillingTaxIdentification2 `xml:"SndgFI,omitempty" json:",omitempty"`
	InvcNb      common.Max40Text          `xml:"InvcNb,omitempty" json:",omitempty"`
	MtdC        BillingMethod4            `xml:"MtdC,omitempty" json:",omitempty"`
	SttlmAmt    AmountAndDirection34      `xml:"SttlmAmt"`
	TaxDueToRgn AmountAndDirection34      `xml:"TaxDueToRgn"`
}

type CashAccountCharacteristics3 struct {
	AcctLvl      AccountLevel2Code                            `xml:"AcctLvl"`
	CshAcct      CashAccount38                                `xml:"CshAcct"`
	AcctSvcr     BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	PrntAcct     ParentCashAccount3                           `xml:"PrntAcct,omitempty" json:",omitempty"`
	CompstnMtd   CompensationMethod1Code                      `xml:"CompstnMtd"`
	DbtAcct      AccountIdentification4Choice                 `xml:"DbtAcct,omitempty" json:",omitempty"`
	DelydDbtDt   common.ISODate                               `xml:"DelydDbtDt,omitempty" json:",omitempty"`
	SttlmAdvc    common.Max105Text                            `xml:"SttlmAdvc,omitempty" json:",omitempty"`
	AcctBalCcyCd common.ActiveOrHistoricCurrencyCode          `xml:"AcctBalCcyCd"`
	SttlmCcyCd   common.ActiveOrHistoricCurrencyCode          `xml:"SttlmCcyCd,omitempty" json:",omitempty"`
	HstCcyCd     common.ActiveOrHistoricCurrencyCode          `xml:"HstCcyCd,omitempty" json:",omitempty"`
	Tax          AccountTax1                                  `xml:"Tax,omitempty" json:",omitempty"`
	AcctSvcrCtct Contact4                                     `xml:"AcctSvcrCtct"`
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

type CurrencyExchange6 struct {
	SrcCcy   common.ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`
	TrgtCcy  common.ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`
	XchgRate float64                             `xml:"XchgRate"`
	Desc     common.Max40Text                    `xml:"Desc,omitempty" json:",omitempty"`
	UnitCcy  common.ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty" json:",omitempty"`
	Cmnts    common.Max70Text                    `xml:"Cmnts,omitempty" json:",omitempty"`
	QtnDt    common.ISODateTime                  `xml:"QtnDt,omitempty" json:",omitempty"`
}

type DatePeriod1 struct {
	FrDt common.ISODate `xml:"FrDt,omitempty" json:",omitempty"`
	ToDt common.ISODate `xml:"ToDt"`
}

type FinancialInstitutionIdentification19 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
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

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"ChanlTp"`
	Id      common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

type Pagination1 struct {
	PgNb      common.Max5NumericText `xml:"PgNb"`
	LastPgInd bool                   `xml:"LastPgInd"`
}

type ParentCashAccount3 struct {
	Lvl  AccountLevel1Code                            `xml:"Lvl,omitempty" json:",omitempty"`
	Id   CashAccount38                                `xml:"Id"`
	Svcr BranchAndFinancialInstitutionIdentification6 `xml:"Svcr,omitempty" json:",omitempty"`
}

type Party43Choice struct {
	OrgId OrganisationIdentification29         `xml:"OrgId"`
	FIId  FinancialInstitutionIdentification19 `xml:"FIId"`
}

type PartyIdentification138 struct {
	Nm        common.Max140Text  `xml:"Nm"`
	LglNm     common.Max140Text  `xml:"LglNm,omitempty" json:",omitempty"`
	PstlAdr   PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        Party43Choice      `xml:"Id"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   common.Max35Text `xml:"Cd"`
	Issr common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type ReportHeader6 struct {
	RptId    common.Max35Text `xml:"RptId"`
	MsgPgntn Pagination1      `xml:"MsgPgntn,omitempty" json:",omitempty"`
}

type ResidenceLocation1Choice struct {
	Ctry common.CountryCode `xml:"Ctry"`
	Area common.Max35Text   `xml:"Area"`
}

type ServiceTaxDesignation1 struct {
	Cd     ServiceTaxDesignation1Code `xml:"Cd"`
	Rgn    common.Max35Text           `xml:"Rgn,omitempty" json:",omitempty"`
	TaxRsn []TaxReason1               `xml:"TaxRsn,omitempty" json:",omitempty"`
}

type StatementGroup3 struct {
	GrpId        common.Max35Text       `xml:"GrpId"`
	Sndr         PartyIdentification138 `xml:"Sndr"`
	SndrIndvCtct []Contact4             `xml:"SndrIndvCtct,omitempty" json:",omitempty"`
	Rcvr         PartyIdentification138 `xml:"Rcvr"`
	RcvrIndvCtct []Contact4             `xml:"RcvrIndvCtct,omitempty" json:",omitempty"`
	BllgStmt     []BillingStatement3    `xml:"BllgStmt"`
}

type TaxCalculation1 struct {
	HstCcy                common.ActiveOrHistoricCurrencyCode `xml:"HstCcy"`
	TaxblSvcChrgConvs     []BillingServicesAmount3            `xml:"TaxblSvcChrgConvs"`
	TtlTaxblSvcChrgHstAmt AmountAndDirection34                `xml:"TtlTaxblSvcChrgHstAmt"`
	TaxId                 []BillingServicesTax3               `xml:"TaxId"`
	TtlTax                AmountAndDirection34                `xml:"TtlTax"`
}

type TaxReason1 struct {
	Cd     common.Max10Text  `xml:"Cd"`
	Expltn common.Max105Text `xml:"Expltn"`
}
