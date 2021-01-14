// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v08

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

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
	return utils.Validate(&r)
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
	OrgnlTrckgDays   common.Exact2NumericText                     `xml:"OrgnlTrckgDays,omitempty" json:",omitempty"`
}

func (r AmendmentInformationDetails13) Validate() error {
	return utils.Validate(&r)
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"Cd"`
	Prtry common.Max128Text         `xml:"Prtry"`
}

func (r Authorisation1Choice) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    BranchData3                          `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification6) Validate() error {
	return utils.Validate(&r)
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData3) Validate() error {
	return utils.Validate(&r)
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"Id"`
	Tp   CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
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

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return utils.Validate(&r)
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (r Charges7) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r ClearingSystemIdentification3Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                    `xml:"MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return utils.Validate(&r)
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

func (r Contact4) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
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
	Issr      common.Max35Text             `xml:"Issr,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType2) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
	return utils.Validate(&r)
}

type DirectDebitTransaction10 struct {
	MndtRltdInf MandateRelatedInformation14 `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	CdtrSchmeId PartyIdentification135      `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	PreNtfctnId common.Max35Text            `xml:"PreNtfctnId,omitempty" json:",omitempty"`
	PreNtfctnDt common.ISODate              `xml:"PreNtfctnDt,omitempty" json:",omitempty"`
}

func (r DirectDebitTransaction10) Validate() error {
	return utils.Validate(&r)
}

type DirectDebitTransactionInformation24 struct {
	PmtId          PaymentIdentification7                       `xml:"PmtId"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	IntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	SttlmPrty      Priority3Code                                `xml:"SttlmPrty,omitempty" json:",omitempty"`
	SttlmTmIndctn  SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty" json:",omitempty"`
	InstdAmt       ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty" json:",omitempty"`
	XchgRate       float64                                      `xml:"XchgRate,omitempty" json:",omitempty"`
	ChrgBr         ChargeBearerType1Code                        `xml:"ChrgBr"`
	ChrgsInf       []Charges7                                   `xml:"ChrgsInf,omitempty" json:",omitempty"`
	ReqdColltnDt   common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	DrctDbtTx      DirectDebitTransaction10                     `xml:"DrctDbtTx,omitempty" json:",omitempty"`
	Cdtr           PartyIdentification135                       `xml:"Cdtr"`
	CdtrAcct       CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	CdtrAgtAcct    CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	UltmtCdtr      PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	InitgPty       PartyIdentification135                       `xml:"InitgPty,omitempty" json:",omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:",omitempty"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:",omitempty"`
	IntrmyAgt1     BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt1Acct CashAccount38                                `xml:"IntrmyAgt1Acct,omitempty" json:",omitempty"`
	IntrmyAgt2     BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt2Acct CashAccount38                                `xml:"IntrmyAgt2Acct,omitempty" json:",omitempty"`
	IntrmyAgt3     BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	IntrmyAgt3Acct CashAccount38                                `xml:"IntrmyAgt3Acct,omitempty" json:",omitempty"`
	Dbtr           PartyIdentification135                       `xml:"Dbtr"`
	DbtrAcct       CashAccount38                                `xml:"DbtrAcct"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	DbtrAgtAcct    CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	UltmtDbtr      PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Purp           Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RgltryRptg     []RegulatoryReporting3                       `xml:"RgltryRptg,omitempty" json:",omitempty"`
	RltdRmtInf     []RemittanceLocation7                        `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	SplmtryData    []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DirectDebitTransactionInformation24) Validate() error {
	return utils.Validate(&r)
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"Tp,omitempty" json:",omitempty"`
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
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       common.Max4Text                   `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  common.Max140Text                 `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DocumentAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

func (r DocumentLineIdentification1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc common.Max2048Text            `xml:"Desc,omitempty" json:",omitempty"`
	Amt  RemittanceAmount3             `xml:"Amt,omitempty" json:",omitempty"`
}

func (r DocumentLineInformation1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text        `xml:"Issr,omitempty" json:",omitempty"`
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

type FIToFICustomerDirectDebitV08 struct {
	GrpHdr       GroupHeader94                         `xml:"GrpHdr"`
	DrctDbtTxInf []DirectDebitTransactionInformation24 `xml:"DrctDbtTxInf"`
	SplmtryData  []SupplementaryData1                  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r FIToFICustomerDirectDebitV08) Validate() error {
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
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification18) Validate() error {
	return utils.Validate(&r)
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"Tp"`
	Prd    FrequencyPeriod1    `xml:"Prd"`
	PtInTm FrequencyAndMoment1 `xml:"PtInTm"`
}

func (r Frequency36Choice) Validate() error {
	return utils.Validate(&r)
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code           `xml:"Tp"`
	PtInTm common.Exact2NumericText `xml:"PtInTm"`
}

func (r FrequencyAndMoment1) Validate() error {
	return utils.Validate(&r)
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

func (r FrequencyPeriod1) Validate() error {
	return utils.Validate(&r)
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

func (r Garnishment3) Validate() error {
	return utils.Validate(&r)
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
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

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"Id"`
	SchmeNm AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericFinancialIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm common.Max35Text              `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
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
	Id      common.Max35Text                      `xml:"Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericPersonIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader94 struct {
	MsgId             common.Max35Text                             `xml:"MsgId"`
	CreDtTm           common.ISODateTime                           `xml:"CreDtTm"`
	Authstn           []Authorisation1Choice                       `xml:"Authstn,omitempty" json:",omitempty"`
	BtchBookg         bool                                         `xml:"BtchBookg,omitempty" json:",omitempty"`
	NbOfTxs           common.Max15NumericText                      `xml:"NbOfTxs"`
	CtrlSum           float64                                      `xml:"CtrlSum,omitempty" json:",omitempty"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	IntrBkSttlmDt     common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	SttlmInf          SettlementInstruction8                       `xml:"SttlmInf"`
	PmtTpInf          PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:",omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:",omitempty"`
}

func (r GroupHeader94) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
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
	TrckgDays     common.Exact2NumericText      `xml:"TrckgDays,omitempty" json:",omitempty"`
}

func (r MandateRelatedInformation14) Validate() error {
	return utils.Validate(&r)
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

func (r MandateSetupReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress24   `xml:"Adr"`
}

func (r NameAndAddress16) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    common.LEIIdentifier                 `xml:"LEI,omitempty" json:",omitempty"`
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
	ChanlTp common.Max4Text   `xml:"ChanlTp"`
	Id      common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
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

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification135) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification7 struct {
	InstrId    common.Max35Text        `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text        `xml:"EndToEndId"`
	TxId       common.Max35Text        `xml:"TxId,omitempty" json:",omitempty"`
	UETR       common.UUIDv4Identifier `xml:"UETR,omitempty" json:",omitempty"`
	ClrSysRef  common.Max35Text        `xml:"ClrSysRef,omitempty" json:",omitempty"`
}

func (r PaymentIdentification7) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation27) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
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

func (r PostalAddress24) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text      `xml:"Id"`
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

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r Purpose2Choice) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"Tp,omitempty" json:",omitempty"`
	Nb       common.Max35Text           `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   common.ISODate             `xml:"RltdDt,omitempty" json:",omitempty"`
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
	Issr      common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

func (r ReferredDocumentType4) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	Ctry common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
}

func (r RegulatoryAuthority2) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"DbtCdtRptgInd,omitempty" json:",omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"Authrty,omitempty" json:",omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r RegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount2) Validate() error {
	return utils.Validate(&r)
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount3) Validate() error {
	return utils.Validate(&r)
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:",omitempty"`
}

func (r RemittanceLocation7) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    NameAndAddress16              `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r RemittanceLocationData1) Validate() error {
	return utils.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceLevel8Choice) Validate() error {
	return utils.Validate(&r)
}

type SettlementDateTimeIndication1 struct {
	DbtDtTm common.ISODateTime `xml:"DbtDtTm,omitempty" json:",omitempty"`
	CdtDtTm common.ISODateTime `xml:"CdtDtTm,omitempty" json:",omitempty"`
}

func (r SettlementDateTimeIndication1) Validate() error {
	return utils.Validate(&r)
}

type SettlementInstruction8 struct {
	SttlmMtd  SettlementMethod2Code               `xml:"SttlmMtd"`
	SttlmAcct CashAccount38                       `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys    ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty" json:",omitempty"`
}

func (r SettlementInstruction8) Validate() error {
	return utils.Validate(&r)
}

type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                  `xml:"Tp,omitempty" json:",omitempty"`
	Dt   common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	Ctry common.CountryCode                `xml:"Ctry,omitempty" json:",omitempty"`
	Cd   common.Max10Text                  `xml:"Cd,omitempty" json:",omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Inf  []common.Max35Text                `xml:"Inf,omitempty" json:",omitempty"`
}

func (r StructuredRegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
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

func (r StructuredRemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"PlcAndNm,omitempty" json:",omitempty"`
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

type TaxAmount2 struct {
	Rate         float64                           `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount2) Validate() error {
	return utils.Validate(&r)
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"Tp,omitempty" json:",omitempty"`
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
	Titl common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
}

func (r TaxAuthorisation1) Validate() error {
	return utils.Validate(&r)
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

func (r TaxInformation7) Validate() error {
	return utils.Validate(&r)
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

func (r TaxParty1) Validate() error {
	return utils.Validate(&r)
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

func (r TaxParty2) Validate() error {
	return utils.Validate(&r)
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod2) Validate() error {
	return utils.Validate(&r)
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

func (r TaxRecord2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails2) Validate() error {
	return utils.Validate(&r)
}
