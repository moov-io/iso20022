// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v09

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      *common.Max35Text                             `xml:"OrgnlMndtId,omitempty" json:",omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification135                       `xml:"OrgnlCdtrSchmeId,omitempty" json:",omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:",omitempty"`
	OrgnlCdtrAgtAcct *CashAccount38                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlDbtr        *PartyIdentification135                       `xml:"OrgnlDbtr,omitempty" json:",omitempty"`
	OrgnlDbtrAcct    *CashAccount38                                `xml:"OrgnlDbtrAcct,omitempty" json:",omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:",omitempty"`
	OrgnlDbtrAgtAcct *CashAccount38                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlFnlColltnDt *common.ISODate                               `xml:"OrgnlFnlColltnDt,omitempty" json:",omitempty"`
	OrgnlFrqcy       *Frequency36Choice                            `xml:"OrgnlFrqcy,omitempty" json:",omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:",omitempty"`
	OrgnlTrckgDays   *Exact2NumericText                            `xml:"OrgnlTrckgDays,omitempty" json:",omitempty"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
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

type CancellationReason33Choice struct {
	Cd    ExternalCancellationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
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

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
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

type ControlData1 struct {
	NbOfTxs common.Max15NumericText `xml:"NbOfTxs"`
	CtrlSum float64                 `xml:"CtrlSum,omitempty" json:",omitempty"`
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

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref *common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

type CustomerPaymentCancellationRequestV09 struct {
	Assgnmt     CaseAssignment5           `xml:"Assgnmt"`
	Case        *Case5                    `xml:"Case,omitempty" json:",omitempty"`
	CtrlData    *ControlData1             `xml:"CtrlData,omitempty" json:",omitempty"`
	Undrlyg     []UnderlyingTransaction27 `xml:"Undrlyg"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     *common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt *common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc *common.Max2048Text           `xml:"Desc,omitempty" json:",omitempty"`
	Amt  *RemittanceAmount3            `xml:"Amt,omitempty" json:",omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
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
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text      `xml:"Issr,omitempty" json:",omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
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

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt *MandateRelatedInformation14 `xml:"DrctDbtMndt,omitempty" json:",omitempty"`
	CdtTrfMndt  *CreditTransferMandateData1  `xml:"CdtTrfMndt,omitempty" json:",omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        *common.Max35Text              `xml:"MndtId,omitempty" json:",omitempty"`
	DtOfSgntr     *common.ISODate                `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	AmdmntInd     bool                           `xml:"AmdmntInd,omitempty" json:",omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:",omitempty"`
	ElctrncSgntr  *common.Max1025Text            `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstColltnDt  *common.ISODate                `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt   *common.ISODate                `xml:"FnlColltnDt,omitempty" json:",omitempty"`
	Frqcy         *Frequency36Choice             `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:",omitempty"`
	TrckgDays     *Exact2NumericText             `xml:"TrckgDays,omitempty" json:",omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
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

type OriginalGroupHeader15 struct {
	GrpCxlId     *common.Max35Text            `xml:"GrpCxlId,omitempty" json:",omitempty"`
	Case         *Case5                       `xml:"Case,omitempty" json:",omitempty"`
	OrgnlMsgId   common.Max35Text             `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text             `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime          `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	NbOfTxs      *common.Max15NumericText     `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum      float64                      `xml:"CtrlSum,omitempty" json:",omitempty"`
	GrpCxl       bool                         `xml:"GrpCxl,omitempty" json:",omitempty"`
	CxlRsnInf    []PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty" json:",omitempty"`
}

type OriginalGroupInformation29 struct {
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

type OriginalPaymentInstruction36 struct {
	PmtCxlId      *common.Max35Text            `xml:"PmtCxlId,omitempty" json:",omitempty"`
	Case          *Case5                       `xml:"Case,omitempty" json:",omitempty"`
	OrgnlPmtInfId common.Max35Text             `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf   *OriginalGroupInformation29  `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	NbOfTxs       *common.Max15NumericText     `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum       float64                      `xml:"CtrlSum,omitempty" json:",omitempty"`
	PmtInfCxl     bool                         `xml:"PmtInfCxl,omitempty" json:",omitempty"`
	CxlRsnInf     []PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	TxInf         []PaymentTransaction124      `xml:"TxInf,omitempty" json:",omitempty"`
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    *DateAndDateTime2Choice                       `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    *PartyIdentification135                       `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       *SettlementInstruction7                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       *PaymentTypeInformation27                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    *MandateRelatedData1Choice                    `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      *Party40Choice                                `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    *CashAccount38                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           *Party40Choice                                `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      *Party40Choice                                `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Purp           *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
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

type PaymentCancellationReason5 struct {
	Orgtr    *PartyIdentification135     `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *CancellationReason33Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text         `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type PaymentTransaction124 struct {
	CxlId             *common.Max35Text                  `xml:"CxlId,omitempty" json:",omitempty"`
	Case              *Case5                             `xml:"Case,omitempty" json:",omitempty"`
	OrgnlInstrId      *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId   *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR         *common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlInstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty" json:",omitempty"`
	OrgnlReqdExctnDt  *DateAndDateTime2Choice            `xml:"OrgnlReqdExctnDt,omitempty" json:",omitempty"`
	OrgnlReqdColltnDt *common.ISODate                    `xml:"OrgnlReqdColltnDt,omitempty" json:",omitempty"`
	CxlRsnInf         []PaymentCancellationReason5       `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	OrgnlTxRef        *OriginalTransactionReference31    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1               `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type PaymentTypeInformation27 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     *SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
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

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `xml:"Tp,omitempty" json:",omitempty"`
	Nb       *common.Max35Text          `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   *common.ISODate            `xml:"RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:",omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      *common.Max35Text           `xml:"Issr,omitempty" json:",omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
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
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
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

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty" json:",omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

type TaxAuthorisation1 struct {
	Titl *common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
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

type TaxParty1 struct {
	TaxId  *common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId *common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  *common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

type TaxParty2 struct {
	TaxId   *common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  *common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   *common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

type TaxPeriod2 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
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

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type UnderlyingTransaction27 struct {
	OrgnlGrpInfAndCxl *OriginalGroupHeader15         `xml:"OrgnlGrpInfAndCxl,omitempty" json:",omitempty"`
	OrgnlPmtInfAndCxl []OriginalPaymentInstruction36 `xml:"OrgnlPmtInfAndCxl,omitempty" json:",omitempty"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type FIToFIPaymentCancellationRequestV09 struct {
	Assgnmt     CaseAssignment5           `xml:"Assgnmt"`
	Case        *Case5                    `xml:"Case,omitempty" json:",omitempty"`
	CtrlData    *ControlData1             `xml:"CtrlData,omitempty" json:",omitempty"`
	Undrlyg     []UnderlyingTransaction26 `xml:"Undrlyg"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type PaymentTransaction120 struct {
	CxlId               *common.Max35Text                             `xml:"CxlId,omitempty" json:",omitempty"`
	Case                *Case5                                        `xml:"Case,omitempty" json:",omitempty"`
	OrgnlGrpInf         *OriginalGroupInformation29                   `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                             `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                             `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                             `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlUETR           *common.UUIDv4Identifier                      `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	OrgnlClrSysRef      *common.Max35Text                             `xml:"OrgnlClrSysRef,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmDt  *common.ISODate                               `xml:"OrgnlIntrBkSttlmDt,omitempty" json:",omitempty"`
	Assgnr              *BranchAndFinancialInstitutionIdentification6 `xml:"Assgnr,omitempty" json:",omitempty"`
	Assgne              *BranchAndFinancialInstitutionIdentification6 `xml:"Assgne,omitempty" json:",omitempty"`
	CxlRsnInf           []PaymentCancellationReason5                  `xml:"CxlRsnInf,omitempty" json:",omitempty"`
	OrgnlTxRef          *OriginalTransactionReference31               `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	SplmtryData         []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type UnderlyingTransaction26 struct {
	OrgnlGrpInfAndCxl *OriginalGroupHeader15  `xml:"OrgnlGrpInfAndCxl,omitempty" json:",omitempty"`
	TxInf             []PaymentTransaction120 `xml:"TxInf,omitempty" json:",omitempty"`
}
