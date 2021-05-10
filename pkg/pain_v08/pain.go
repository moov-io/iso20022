// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

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

type AdviceType1 struct {
	CdtAdvc *AdviceType1Choice `xml:"CdtAdvc,omitempty" json:",omitempty"`
	DbtAdvc *AdviceType1Choice `xml:"DbtAdvc,omitempty" json:",omitempty"`
}

func (r AdviceType1) Validate() error {
	return utils.Validate(&r)
}

type AdviceType1Choice struct {
	Cd    AdviceType1Code  `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

func (r AdviceType1Choice) Validate() error {
	return utils.Validate(&r)
}

type AmountOrRate1Choice struct {
	Amt  ActiveCurrencyAndAmount `xml:"Amt"`
	Rate float64                 `xml:"Rate"`
}

func (r AmountOrRate1Choice) Validate() error {
	return utils.Validate(&r)
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

func (r AmountType4Choice) Validate() error {
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

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return utils.Validate(&r)
}

type Cheque11 struct {
	ChqTp       *ChequeType2Code             `xml:"ChqTp,omitempty" json:",omitempty"`
	ChqNb       *common.Max35Text            `xml:"ChqNb,omitempty" json:",omitempty"`
	ChqFr       *NameAndAddress16            `xml:"ChqFr,omitempty" json:",omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty" json:",omitempty"`
	DlvrTo      *NameAndAddress16            `xml:"DlvrTo,omitempty" json:",omitempty"`
	InstrPrty   *Priority2Code               `xml:"InstrPrty,omitempty" json:",omitempty"`
	ChqMtrtyDt  *common.ISODate              `xml:"ChqMtrtyDt,omitempty" json:",omitempty"`
	FrmsCd      *common.Max35Text            `xml:"FrmsCd,omitempty" json:",omitempty"`
	MemoFld     []common.Max35Text           `xml:"MemoFld,omitempty" json:",omitempty"`
	RgnlClrZone *common.Max35Text            `xml:"RgnlClrZone,omitempty" json:",omitempty"`
	PrtLctn     *common.Max35Text            `xml:"PrtLctn,omitempty" json:",omitempty"`
	Sgntr       []common.Max70Text           `xml:"Sgntr,omitempty" json:",omitempty"`
}

func (r Cheque11) Validate() error {
	return utils.Validate(&r)
}

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"Cd"`
	Prtry common.Max35Text    `xml:"Prtry"`
}

func (r ChequeDeliveryMethod1Choice) Validate() error {
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

type CreditTransferMandateData1 struct {
	MndtId       *common.Max35Text          `xml:"MndtId,omitempty" json:",omitempty"`
	Tp           *MandateTypeInformation2   `xml:"Tp,omitempty" json:",omitempty"`
	DtOfSgntr    *common.ISODate            `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	DtOfVrfctn   *common.ISODateTime        `xml:"DtOfVrfctn,omitempty" json:",omitempty"`
	ElctrncSgntr *common.Max10KBinary       `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstPmtDt    *common.ISODate            `xml:"FrstPmtDt,omitempty" json:",omitempty"`
	FnlPmtDt     *common.ISODate            `xml:"FnlPmtDt,omitempty" json:",omitempty"`
	Frqcy        *Frequency36Choice         `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn          *MandateSetupReason1Choice `xml:"Rsn,omitempty" json:",omitempty"`
}

func (r CreditTransferMandateData1) Validate() error {
	return utils.Validate(&r)
}

type CreditTransferTransaction42 struct {
	PmtId           PaymentIdentification6                        `xml:"PmtId"`
	PmtTpInf        *PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtCond         *PaymentCondition1                            `xml:"PmtCond,omitempty" json:",omitempty"`
	Amt             AmountType4Choice                             `xml:"Amt"`
	ChrgBr          ChargeBearerType1Code                         `xml:"ChrgBr"`
	MndtRltdInf     *CreditTransferMandateData1                   `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	ChqInstr        *Cheque11                                     `xml:"ChqInstr,omitempty" json:",omitempty"`
	UltmtDbtr       *PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	Cdtr            PartyIdentification135                        `xml:"Cdtr"`
	CdtrAcct        *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr       *PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent3                `xml:"InstrForCdtrAgt,omitempty" json:",omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty" json:",omitempty"`
	Tax             *TaxInformation8                              `xml:"Tax,omitempty" json:",omitempty"`
	RltdRmtInf      []RemittanceLocation7                         `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf          *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	NclsdFile       []Document12                                  `xml:"NclsdFile,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditTransferTransaction42) Validate() error {
	return utils.Validate(&r)
}

type CreditorPaymentActivationRequestV08 struct {
	Attr        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      GroupHeader78          `xml:"GrpHdr"`
	PmtInf      []PaymentInstruction35 `xml:"PmtInf" json:",omitempty"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorPaymentActivationRequestV08) Validate() error {
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

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

func (r DateAndDateTime2Choice) Validate() error {
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

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
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

type Document12 struct {
	Tp        DocumentType1Choice    `xml:"Tp"`
	Id        common.Max35Text       `xml:"Id"`
	IsseDt    DateAndDateTime2Choice `xml:"IsseDt"`
	Nm        *common.Max140Text     `xml:"Nm,omitempty" json:",omitempty"`
	LangCd    string                 `xml:"LangCd,omitempty" json:",omitempty"`
	Frmt      DocumentFormat1Choice  `xml:"Frmt"`
	FileNm    *common.Max140Text     `xml:"FileNm,omitempty" json:",omitempty"`
	DgtlSgntr *PartyAndSignature3    `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	Nclsr     common.Max10MbBinary   `xml:"Nclsr"`
}

func (r Document12) Validate() error {
	return utils.Validate(&r)
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DocumentAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type DocumentFormat1Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"Cd"`
	Prtry GenericIdentification1      `xml:"Prtry"`
}

func (r DocumentFormat1Choice) Validate() error {
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
	Id   []DocumentLineIdentification1 `xml:"Id" json:",omitempty"`
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

type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"Cd"`
	Prtry GenericIdentification1    `xml:"Prtry"`
}

func (r DocumentType1Choice) Validate() error {
	return utils.Validate(&r)
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

func (r EquivalentAmount2) Validate() error {
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
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment3) Validate() error {
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

type GroupHeader78 struct {
	MsgId    common.Max35Text        `xml:"MsgId"`
	CreDtTm  common.ISODateTime      `xml:"CreDtTm"`
	NbOfTxs  common.Max15NumericText `xml:"NbOfTxs"`
	CtrlSum  float64                 `xml:"CtrlSum,omitempty" json:",omitempty"`
	InitgPty PartyIdentification135  `xml:"InitgPty"`
}

func (r GroupHeader78) Validate() error {
	return utils.Validate(&r)
}

type InstructionForCreditorAgent3 struct {
	Cd       *ExternalCreditorAgentInstruction1Code `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text                     `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForCreditorAgent3) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r MandateClassification1Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

func (r MandateSetupReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
}

func (r MandateTypeInformation2) Validate() error {
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

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"Pty"`
	Sgntr SkipPayload            `xml:"Sgntr"`
}

func (r PartyAndSignature3) Validate() error {
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

type PaymentCondition1 struct {
	AmtModAllwd   bool                 `xml:"AmtModAllwd"`
	EarlyPmtAllwd bool                 `xml:"EarlyPmtAllwd"`
	DelyPnlty     *common.Max140Text   `xml:"DelyPnlty,omitempty" json:",omitempty"`
	ImdtPmtRbt    *AmountOrRate1Choice `xml:"ImdtPmtRbt,omitempty" json:",omitempty"`
	GrntedPmtReqd bool                 `xml:"GrntedPmtReqd"`
}

func (r PaymentCondition1) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification6 struct {
	InstrId    *common.Max35Text        `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text         `xml:"EndToEndId"`
	UETR       *common.UUIDv4Identifier `xml:"UETR,omitempty" json:",omitempty"`
}

func (r PaymentIdentification6) Validate() error {
	return utils.Validate(&r)
}

type PaymentInstruction35 struct {
	PmtInfId    *common.Max35Text                            `xml:"PmtInfId,omitempty" json:",omitempty"`
	PmtMtd      PaymentMethod7Code                           `xml:"PmtMtd"`
	ReqdAdvcTp  *AdviceType1                                 `xml:"ReqdAdvcTp,omitempty" json:",omitempty"`
	PmtTpInf    *PaymentTypeInformation26                    `xml:"PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt DateAndDateTime2Choice                       `xml:"ReqdExctnDt"`
	XpryDt      *DateAndDateTime2Choice                      `xml:"XpryDt,omitempty" json:",omitempty"`
	PmtCond     *PaymentCondition1                           `xml:"PmtCond,omitempty" json:",omitempty"`
	Dbtr        PartyIdentification135                       `xml:"Dbtr"`
	DbtrAcct    *CashAccount38                               `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	UltmtDbtr   *PartyIdentification135                      `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	ChrgBr      *ChargeBearerType1Code                       `xml:"ChrgBr,omitempty" json:",omitempty"`
	CdtTrfTx    []CreditTransferTransaction42                `xml:"CdtTrfTx" json:",omitempty"`
}

func (r PaymentInstruction35) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation26 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation26) Validate() error {
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

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r Purpose2Choice) Validate() error {
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

type RegulatoryAuthority2 struct {
	Nm   *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	Ctry *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
}

func (r RegulatoryAuthority2) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd *RegulatoryReportingType1Code    `xml:"DbtCdtRptgInd,omitempty" json:",omitempty"`
	Authrty       *RegulatoryAuthority2            `xml:"Authrty,omitempty" json:",omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r RegulatoryReporting3) Validate() error {
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

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocation7 struct {
	RmtId       *common.Max35Text         `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:",omitempty"`
}

func (r RemittanceLocation7) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr *common.Max2048Text           `xml:"ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    *NameAndAddress16             `xml:"PstlAdr,omitempty" json:",omitempty"`
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

type SkipPayload struct {
	Item string `xml:",any"`
}

func (r SkipPayload) Validate() error {
	return utils.Validate(&r)
}

type StructuredRegulatoryReporting3 struct {
	Tp   *common.Max35Text                  `xml:"Tp,omitempty" json:",omitempty"`
	Dt   *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	Ctry *common.CountryCode                `xml:"Ctry,omitempty" json:",omitempty"`
	Cd   *common.Max10Text                  `xml:"Cd,omitempty" json:",omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Inf  []common.Max35Text                 `xml:"Inf,omitempty" json:",omitempty"`
}

func (r StructuredRegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification135        `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification135        `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation7               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment3                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation16) Validate() error {
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

type TaxAmount2 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount2) Validate() error {
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

type TaxInformation7 struct {
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
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation7) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation8) Validate() error {
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

type TaxPeriod2 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecord2 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod2        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount2        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails2) Validate() error {
	return utils.Validate(&r)
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (r Charges7) Validate() error {
	return utils.Validate(&r)
}

type CreditorPaymentActivationRequestStatusReportV08 struct {
	Attr              []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr            GroupHeader87                  `xml:"GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation30     `xml:"OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction39 `xml:"OrgnlPmtInfAndSts,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1           `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorPaymentActivationRequestStatusReportV08) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader87 struct {
	MsgId    common.Max35Text                             `xml:"MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"CreDtTm"`
	InitgPty PartyIdentification135                       `xml:"InitgPty"`
	DbtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
}

func (r GroupHeader87) Validate() error {
	return utils.Validate(&r)
}

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs common.Max15NumericText               `xml:"DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `xml:"DtldSts"`
	DtldCtrlSum float64                               `xml:"DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfTransactionsPerStatus5) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupInformation30 struct {
	OrgnlMsgId    common.Max35Text                 `xml:"OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                 `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm  common.ISODateTime               `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText          `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum  float64                          `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	GrpSts        ExternalPaymentGroupStatus1Code  `xml:"GrpSts,omitempty" json:",omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"StsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty" json:",omitempty"`
}

func (r OriginalGroupInformation30) Validate() error {
	return utils.Validate(&r)
}

type OriginalPaymentInstruction39 struct {
	OrgnlPmtInfId common.Max35Text                 `xml:"OrgnlPmtInfId"`
	OrgnlNbOfTxs  common.Max15NumericText          `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum  float64                          `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	PmtInfSts     ExternalPaymentGroupStatus1Code  `xml:"PmtInfSts,omitempty" json:",omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"StsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty" json:",omitempty"`
	TxInfAndSts   []PaymentTransaction128          `xml:"TxInfAndSts,omitempty" json:",omitempty"`
}

func (r OriginalPaymentInstruction39) Validate() error {
	return utils.Validate(&r)
}

type OriginalTransactionReference33 struct {
	Amt         AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	ReqdExctnDt DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	XpryDt      DateAndDateTime2Choice                       `xml:"XpryDt,omitempty" json:",omitempty"`
	PmtCond     PaymentCondition1                            `xml:"PmtCond,omitempty" json:",omitempty"`
	PmtTpInf    PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd      PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf CreditTransferMandateData1                   `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf      RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	NclsdFile   []Document12                                 `xml:"NclsdFile,omitempty" json:",omitempty"`
	UltmtDbtr   PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr        PartyIdentification135                       `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct    CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	Cdtr        PartyIdentification135                       `xml:"Cdtr"`
	CdtrAcct    CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr   PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference33) Validate() error {
	return utils.Validate(&r)
}

type PaymentConditionStatus1 struct {
	AccptdAmt ActiveCurrencyAndAmount `xml:"AccptdAmt,omitempty" json:",omitempty"`
	GrntedPmt bool                    `xml:"GrntedPmt"`
	EarlyPmt  bool                    `xml:"EarlyPmt"`
}

func (r PaymentConditionStatus1) Validate() error {
	return utils.Validate(&r)
}

type PaymentTransaction128 struct {
	StsId           common.Max35Text                      `xml:"StsId,omitempty" json:",omitempty"`
	OrgnlInstrId    common.Max35Text                      `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId common.Max35Text                      `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       common.UUIDv4Identifier               `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	TxSts           ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty" json:",omitempty"`
	StsRsnInf       []StatusReasonInformation12           `xml:"StsRsnInf,omitempty" json:",omitempty"`
	PmtCondSts      PaymentConditionStatus1               `xml:"PmtCondSts,omitempty" json:",omitempty"`
	ChrgsInf        []Charges7                            `xml:"ChrgsInf,omitempty" json:",omitempty"`
	DbtrDcsnDtTm    common.ISODateTime                    `xml:"DbtrDcsnDtTm,omitempty" json:",omitempty"`
	AccptncDtTm     common.ISODateTime                    `xml:"AccptncDtTm,omitempty" json:",omitempty"`
	AcctSvcrRef     common.Max35Text                      `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	ClrSysRef       common.Max35Text                      `xml:"ClrSysRef,omitempty" json:",omitempty"`
	OrgnlTxRef      OriginalTransactionReference33        `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	NclsdFile       []Document12                          `xml:"NclsdFile,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentTransaction128) Validate() error {
	return utils.Validate(&r)
}

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r StatusReason6Choice) Validate() error {
	return utils.Validate(&r)
}

type StatusReasonInformation12 struct {
	Orgtr    PartyIdentification135 `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      StatusReason6Choice    `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text    `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r StatusReasonInformation12) Validate() error {
	return utils.Validate(&r)
}
