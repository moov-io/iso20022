// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v07

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type AccountIdentification4Choice struct {
	IBAN *common.IBAN2007Identifier     `IBAN,omitempty" json:",omitempty"`
	Othr *GenericAccountIdentification1 `Othr,omitempty" json:",omitempty"`
}

func (r AccountIdentification4Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountSchemeName1Choice struct {
	Cd    *ExternalAccountIdentification1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                   `Prtry,omitempty" json:",omitempty"`
}

func (r AccountSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type AddressType3Choice struct {
	Cd    *common.AddressType2Code `Cd,omitempty" json:",omitempty"`
	Prtry *GenericIdentification30 `Prtry,omitempty" json:",omitempty"`
}

func (r AddressType3Choice) Validate() error {
	return utils.Validate(&r)
}

type AmountOrRate1Choice struct {
	Amt  *ActiveCurrencyAndAmount `Amt,omitempty" json:",omitempty"`
	Rate float64                  `Rate,omitempty" json:",omitempty"`
}

func (r AmountOrRate1Choice) Validate() error {
	return utils.Validate(&r)
}

type AmountType4Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `InstdAmt,omitempty" json:",omitempty"`
	EqvtAmt  *EquivalentAmount2                 `EqvtAmt,omitempty" json:",omitempty"`
}

func (r AmountType4Choice) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `FinInstnId"`
	BrnchId    *BranchData3                         `BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification6) Validate() error {
	return utils.Validate(&r)
}

type BranchData3 struct {
	Id      *common.Max35Text     `Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData3) Validate() error {
	return utils.Validate(&r)
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice  `Id"`
	Tp   *CashAccountType2Choice       `Tp,omitempty" json:",omitempty"`
	Ccy  *ActiveOrHistoricCurrencyCode `Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text             `Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1  `Prxy,omitempty" json:",omitempty"`
}

func (r CashAccount38) Validate() error {
	return utils.Validate(&r)
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text             `Prtry,omitempty" json:",omitempty"`
}

func (r CashAccountType2Choice) Validate() error {
	return utils.Validate(&r)
}

type CategoryPurpose1Choice struct {
	Cd    *ExternalCategoryPurpose1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text             `Prtry,omitempty" json:",omitempty"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return utils.Validate(&r)
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `Agt"`
}

func (r Charges7) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification2Choice struct {
	Cd    *ExternalClearingSystemIdentification1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                          `Prtry,omitempty" json:",omitempty"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return utils.Validate(&r)
}

type Contact4 struct {
	NmPrfx    *common.NamePrefix2Code      `NmPrfx,omitempty" json:",omitempty"`
	Nm        *common.Max140Text           `Nm,omitempty" json:",omitempty"`
	PhneNb    *common.PhoneNumber          `PhneNb,omitempty" json:",omitempty"`
	MobNb     *common.PhoneNumber          `MobNb,omitempty" json:",omitempty"`
	FaxNb     *common.PhoneNumber          `FaxNb,omitempty" json:",omitempty"`
	EmailAdr  *common.Max2048Text          `EmailAdr,omitempty" json:",omitempty"`
	EmailPurp *common.Max35Text            `EmailPurp,omitempty" json:",omitempty"`
	JobTitl   *common.Max35Text            `JobTitl,omitempty" json:",omitempty"`
	Rspnsblty *common.Max35Text            `Rspnsblty,omitempty" json:",omitempty"`
	Dept      *common.Max70Text            `Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1              `Othr,omitempty" json:",omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `PrefrdMtd,omitempty" json:",omitempty"`
}

func (r Contact4) Validate() error {
	return utils.Validate(&r)
}

type CreditorPaymentActivationRequestStatusReportV07 struct {
	XMLName           xml.Name                       `xml:"CdtrPmtActvtnReqStsRpt"`
	GrpHdr            GroupHeader87                  `GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation30     `OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction31 `OrgnlPmtInfAndSts,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1           `SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorPaymentActivationRequestStatusReportV07) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `Tp,omitempty" json:",omitempty"`
	Ref *common.Max35Text       `Ref,omitempty" json:",omitempty"`
}

func (r CreditorReferenceInformation2) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceType1Choice struct {
	Cd    *DocumentType3Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text  `Prtry,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType1Choice) Validate() error {
	return utils.Validate(&r)
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `CdOrPrtry"`
	Issr      *common.Max35Text            `Issr,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType2) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTime2Choice struct {
	Dt   *common.ISODate     `Dt,omitempty" json:",omitempty"`
	DtTm *common.ISODateTime `DtTm,omitempty" json:",omitempty"`
}

func (r DateAndDateTime2Choice) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `BirthDt"`
	PrvcOfBirth *common.Max35Text  `PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `CityOfBirth"`
	CtryOfBirth common.CountryCode `CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2 struct {
	FrDt common.ISODate `FrDt"`
	ToDt common.ISODate `ToDt"`
}

func (r DatePeriod2) Validate() error {
	return utils.Validate(&r)
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice        `Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `Amt"`
}

func (r DiscountAmountAndType1) Validate() error {
	return utils.Validate(&r)
}

type DiscountAmountType1Choice struct {
	Cd    *ExternalDiscountAmountType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                `Prtry,omitempty" json:",omitempty"`
}

func (r DiscountAmountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type Document12 struct {
	Tp        DocumentType1Choice    `Tp"`
	Id        common.Max35Text       `Id"`
	IsseDt    DateAndDateTime2Choice `IsseDt"`
	Nm        *common.Max140Text     `Nm,omitempty" json:",omitempty"`
	LangCd    string                 `LangCd,omitempty" json:",omitempty"`
	Frmt      DocumentFormat1Choice  `Frmt"`
	FileNm    *common.Max140Text     `FileNm,omitempty" json:",omitempty"`
	DgtlSgntr *PartyAndSignature3    `DgtlSgntr,omitempty" json:",omitempty"`
	Nclsr     common.Max10MbBinary   `Nclsr"`
}

func (r Document12) Validate() error {
	return utils.Validate(&r)
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `Amt"`
	CdtDbtInd *common.CreditDebitCode           `CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `AddtlInf,omitempty" json:",omitempty"`
}

func (r DocumentAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type DocumentFormat1Choice struct {
	Cd    *ExternalDocumentFormat1Code `Cd,omitempty" json:",omitempty"`
	Prtry *GenericIdentification1      `Prtry,omitempty" json:",omitempty"`
}

func (r DocumentFormat1Choice) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `Tp,omitempty" json:",omitempty"`
	Nb     *common.Max35Text  `Nb,omitempty" json:",omitempty"`
	RltdDt *common.ISODate    `RltdDt,omitempty" json:",omitempty"`
}

func (r DocumentLineIdentification1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `Id"`
	Desc *common.Max2048Text           `Desc,omitempty" json:",omitempty"`
	Amt  *RemittanceAmount3            `Amt,omitempty" json:",omitempty"`
}

func (r DocumentLineInformation1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `CdOrPrtry"`
	Issr      *common.Max35Text       `Issr,omitempty" json:",omitempty"`
}

func (r DocumentLineType1) Validate() error {
	return utils.Validate(&r)
}

type DocumentLineType1Choice struct {
	Cd    *ExternalDocumentLineType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text              `Prtry,omitempty" json:",omitempty"`
}

func (r DocumentLineType1Choice) Validate() error {
	return utils.Validate(&r)
}

type DocumentType1Choice struct {
	Cd    *ExternalDocumentType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *GenericIdentification1    `Prtry,omitempty" json:",omitempty"`
}

func (r DocumentType1Choice) Validate() error {
	return utils.Validate(&r)
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `CcyOfTrf"`
}

func (r EquivalentAmount2) Validate() error {
	return utils.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    *ExternalFinancialInstitutionIdentification1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                                `Prtry,omitempty" json:",omitempty"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification18) Validate() error {
	return utils.Validate(&r)
}

type Garnishment3 struct {
	Tp                GarnishmentType1                   `Tp"`
	Grnshee           *PartyIdentification135            `Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment3) Validate() error {
	return utils.Validate(&r)
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `CdOrPrtry"`
	Issr      *common.Max35Text      `Issr,omitempty" json:",omitempty"`
}

func (r GarnishmentType1) Validate() error {
	return utils.Validate(&r)
}

type GarnishmentType1Choice struct {
	Cd    *ExternalGarnishmentType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text             `Prtry,omitempty" json:",omitempty"`
}

func (r GarnishmentType1Choice) Validate() error {
	return utils.Validate(&r)
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `Id"`
	SchmeNm *AccountSchemeName1Choice `SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `Issr,omitempty" json:",omitempty"`
}

func (r GenericAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `Issr,omitempty" json:",omitempty"`
}

func (r GenericFinancialIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification1 struct {
	Id      common.Max35Text  `Id"`
	SchmeNm *common.Max35Text `SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text `Issr,omitempty" json:",omitempty"`
}

func (r GenericIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `Id"`
	Issr    common.Max35Text       `Issr"`
	SchmeNm *common.Max35Text      `SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
	return utils.Validate(&r)
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                            `Issr,omitempty" json:",omitempty"`
}

func (r GenericOrganisationIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `Issr,omitempty" json:",omitempty"`
}

func (r GenericPersonIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader87 struct {
	MsgId    common.Max35Text                              `MsgId"`
	CreDtTm  common.ISODateTime                            `CreDtTm"`
	InitgPty PartyIdentification135                        `InitgPty"`
	DbtrAgt  *BranchAndFinancialInstitutionIdentification6 `DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt  *BranchAndFinancialInstitutionIdentification6 `CdtrAgt,omitempty" json:",omitempty"`
}

func (r GroupHeader87) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    *ExternalLocalInstrument1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text             `Prtry,omitempty" json:",omitempty"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
}

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs common.Max15NumericText               `DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `DtldSts"`
	DtldCtrlSum float64                               `DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfTransactionsPerStatus5) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification29) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                        `Prtry,omitempty" json:",omitempty"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupInformation30 struct {
	OrgnlMsgId    common.Max35Text                 `OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                 `OrgnlMsgNmId"`
	OrgnlCreDtTm  *common.ISODateTime              `OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNbOfTxs  *common.Max15NumericText         `OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum  float64                          `OrgnlCtrlSum,omitempty" json:",omitempty"`
	GrpSts        *ExternalPaymentGroupStatus1Code `GrpSts,omitempty" json:",omitempty"`
	StsRsnInf     []StatusReasonInformation12      `StsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `NbOfTxsPerSts,omitempty" json:",omitempty"`
}

func (r OriginalGroupInformation30) Validate() error {
	return utils.Validate(&r)
}

type OriginalPaymentInstruction31 struct {
	OrgnlPmtInfId common.Max35Text                 `OrgnlPmtInfId"`
	OrgnlNbOfTxs  *common.Max15NumericText         `OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum  float64                          `OrgnlCtrlSum,omitempty" json:",omitempty"`
	PmtInfSts     *ExternalPaymentGroupStatus1Code `PmtInfSts,omitempty" json:",omitempty"`
	StsRsnInf     []StatusReasonInformation12      `StsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `NbOfTxsPerSts,omitempty" json:",omitempty"`
	TxInfAndSts   []PaymentTransaction104          `TxInfAndSts,omitempty" json:",omitempty"`
}

func (r OriginalPaymentInstruction31) Validate() error {
	return utils.Validate(&r)
}

type OriginalTransactionReference29 struct {
	Amt         *AmountType4Choice                            `Amt,omitempty" json:",omitempty"`
	ReqdExctnDt *DateAndDateTime2Choice                       `ReqdExctnDt,omitempty" json:",omitempty"`
	XpryDt      *DateAndDateTime2Choice                       `XpryDt,omitempty" json:",omitempty"`
	PmtCond     *PaymentCondition1                            `PmtCond,omitempty" json:",omitempty"`
	PmtTpInf    *PaymentTypeInformation26                     `PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd      *PaymentMethod4Code                           `PmtMtd,omitempty" json:",omitempty"`
	RmtInf      *RemittanceInformation16                      `RmtInf,omitempty" json:",omitempty"`
	NclsdFile   []Document12                                  `NclsdFile,omitempty" json:",omitempty"`
	UltmtDbtr   *PartyIdentification135                       `UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr        *PartyIdentification135                       `Dbtr,omitempty" json:",omitempty"`
	DbtrAcct    *CashAccount38                                `DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt     *BranchAndFinancialInstitutionIdentification6 `DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt     BranchAndFinancialInstitutionIdentification6  `CdtrAgt"`
	Cdtr        PartyIdentification135                        `Cdtr"`
	CdtrAcct    *CashAccount38                                `CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr   *PartyIdentification135                       `UltmtCdtr,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference29) Validate() error {
	return utils.Validate(&r)
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `ChanlTp"`
	Id      *common.Max128Text `Id,omitempty" json:",omitempty"`
}

func (r OtherContact1) Validate() error {
	return utils.Validate(&r)
}

type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `OrgId,omitempty" json:",omitempty"`
	PrvtId *PersonIdentification13       `PrvtId,omitempty" json:",omitempty"`
}

func (r Party38Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `Pty"`
	Sgntr SkipPayload            `Sgntr"`
}

func (r PartyAndSignature3) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification135) Validate() error {
	return utils.Validate(&r)
}

type PaymentCondition1 struct {
	AmtModAllwd   bool                 `AmtModAllwd"`
	EarlyPmtAllwd bool                 `EarlyPmtAllwd"`
	DelyPnlty     *common.Max140Text   `DelyPnlty,omitempty" json:",omitempty"`
	ImdtPmtRbt    *AmountOrRate1Choice `ImdtPmtRbt,omitempty" json:",omitempty"`
	GrntedPmtReqd bool                 `GrntedPmtReqd"`
}

func (r PaymentCondition1) Validate() error {
	return utils.Validate(&r)
}

type PaymentConditionStatus1 struct {
	AccptdAmt *ActiveCurrencyAndAmount `AccptdAmt,omitempty" json:",omitempty"`
	GrntedPmt bool                     `GrntedPmt"`
	EarlyPmt  bool                     `EarlyPmt"`
}

func (r PaymentConditionStatus1) Validate() error {
	return utils.Validate(&r)
}

type PaymentTransaction104 struct {
	StsId           *common.Max35Text                      `StsId,omitempty" json:",omitempty"`
	OrgnlInstrId    *common.Max35Text                      `OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId *common.Max35Text                      `OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       *common.UUIDv4Identifier               `OrgnlUETR,omitempty" json:",omitempty"`
	TxSts           *ExternalPaymentTransactionStatus1Code `TxSts,omitempty" json:",omitempty"`
	StsRsnInf       []StatusReasonInformation12            `StsRsnInf,omitempty" json:",omitempty"`
	PmtCondSts      *PaymentConditionStatus1               `PmtCondSts,omitempty" json:",omitempty"`
	ChrgsInf        []Charges7                             `ChrgsInf,omitempty" json:",omitempty"`
	DbtrDcsnDtTm    *common.ISODateTime                    `DbtrDcsnDtTm,omitempty" json:",omitempty"`
	AccptncDtTm     *common.ISODateTime                    `AccptncDtTm,omitempty" json:",omitempty"`
	AcctSvcrRef     *common.Max35Text                      `AcctSvcrRef,omitempty" json:",omitempty"`
	ClrSysRef       *common.Max35Text                      `ClrSysRef,omitempty" json:",omitempty"`
	OrgnlTxRef      *OriginalTransactionReference29        `OrgnlTxRef,omitempty" json:",omitempty"`
	NclsdFile       []Document12                           `NclsdFile,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                   `SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentTransaction104) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation26 struct {
	InstrPrty *Priority2Code          `InstrPrty,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice   `SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation26) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification13) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text                  `Prtry,omitempty" json:",omitempty"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text   `Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text   `SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text   `StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text   `BldgNb,omitempty" json:",omitempty"`
	BldgNm      *common.Max35Text   `BldgNm,omitempty" json:",omitempty"`
	Flr         *common.Max70Text   `Flr,omitempty" json:",omitempty"`
	PstBx       *common.Max16Text   `PstBx,omitempty" json:",omitempty"`
	Room        *common.Max70Text   `Room,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text   `PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text   `TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   *common.Max35Text   `TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    *common.Max35Text   `DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text   `CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode `Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text  `AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress24) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `Id"`
}

func (r ProxyAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountType1Choice struct {
	Cd    *ExternalProxyAccountType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text              `Prtry,omitempty" json:",omitempty"`
}

func (r ProxyAccountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4     `Tp,omitempty" json:",omitempty"`
	Nb       *common.Max35Text          `Nb,omitempty" json:",omitempty"`
	RltdDt   *common.ISODate            `RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `LineDtls,omitempty" json:",omitempty"`
}

func (r ReferredDocumentInformation7) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentType3Choice struct {
	Cd    *DocumentType6Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text  `Prtry,omitempty" json:",omitempty"`
}

func (r ReferredDocumentType3Choice) Validate() error {
	return utils.Validate(&r)
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `CdOrPrtry"`
	Issr      *common.Max35Text           `Issr,omitempty" json:",omitempty"`
}

func (r ReferredDocumentType4) Validate() error {
	return utils.Validate(&r)
}

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount2) Validate() error {
	return utils.Validate(&r)
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1           `DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1                `TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1              `AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `RmtdAmt,omitempty" json:",omitempty"`
}

func (r RemittanceAmount3) Validate() error {
	return utils.Validate(&r)
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    *ExternalServiceLevel1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text          `Prtry,omitempty" json:",omitempty"`
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

type StatusReason6Choice struct {
	Cd    *ExternalStatusReason1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text          `Prtry,omitempty" json:",omitempty"`
}

func (r StatusReason6Choice) Validate() error {
	return utils.Validate(&r)
}

type StatusReasonInformation12 struct {
	Orgtr    *PartyIdentification135 `Orgtr,omitempty" json:",omitempty"`
	Rsn      *StatusReason6Choice    `Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text     `AddtlInf,omitempty" json:",omitempty"`
}

func (r StatusReasonInformation12) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification135        `Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification135        `Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation7               `TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment3                  `GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text          `PlcAndNm,omitempty" json:",omitempty"`
	Envlp    *SupplementaryDataEnvelope1 `Envlp"`
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
	Rate         float64                            `Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2                `Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount2) Validate() error {
	return utils.Validate(&r)
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice             `Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `Amt"`
}

func (r TaxAmountAndType1) Validate() error {
	return utils.Validate(&r)
}

type TaxAmountType1Choice struct {
	Cd    *ExternalTaxAmountType1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text           `Prtry,omitempty" json:",omitempty"`
}

func (r TaxAmountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type TaxAuthorisation1 struct {
	Titl *common.Max35Text  `Titl,omitempty" json:",omitempty"`
	Nm   *common.Max140Text `Nm,omitempty" json:",omitempty"`
}

func (r TaxAuthorisation1) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `Dbtr,omitempty" json:",omitempty"`
	UltmtDbtr       *TaxParty2                         `UltmtDbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                       `Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation7) Validate() error {
	return utils.Validate(&r)
}

type TaxParty1 struct {
	TaxId  *common.Max35Text `TaxId,omitempty" json:",omitempty"`
	RegnId *common.Max35Text `RegnId,omitempty" json:",omitempty"`
	TaxTp  *common.Max35Text `TaxTp,omitempty" json:",omitempty"`
}

func (r TaxParty1) Validate() error {
	return utils.Validate(&r)
}

type TaxParty2 struct {
	TaxId   *common.Max35Text  `TaxId,omitempty" json:",omitempty"`
	RegnId  *common.Max35Text  `RegnId,omitempty" json:",omitempty"`
	TaxTp   *common.Max35Text  `TaxTp,omitempty" json:",omitempty"`
	Authstn *TaxAuthorisation1 `Authstn,omitempty" json:",omitempty"`
}

func (r TaxParty2) Validate() error {
	return utils.Validate(&r)
}

type TaxPeriod2 struct {
	Yr     *common.ISODate       `Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriod2          `FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecord2 struct {
	Tp       *common.Max35Text  `Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod2        `Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount2        `TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `Amt"`
}

func (r TaxRecordDetails2) Validate() error {
	return utils.Validate(&r)
}

type Cheque11 struct {
	ChqTp       *ChequeType2Code             `ChqTp,omitempty" json:",omitempty"`
	ChqNb       *common.Max35Text            `ChqNb,omitempty" json:",omitempty"`
	ChqFr       *NameAndAddress16            `ChqFr,omitempty" json:",omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `DlvryMtd,omitempty" json:",omitempty"`
	DlvrTo      *NameAndAddress16            `DlvrTo,omitempty" json:",omitempty"`
	InstrPrty   *Priority2Code               `InstrPrty,omitempty" json:",omitempty"`
	ChqMtrtyDt  *common.ISODate              `ChqMtrtyDt,omitempty" json:",omitempty"`
	FrmsCd      *common.Max35Text            `FrmsCd,omitempty" json:",omitempty"`
	MemoFld     []common.Max35Text           `MemoFld,omitempty" json:",omitempty"`
	RgnlClrZone *common.Max35Text            `RgnlClrZone,omitempty" json:",omitempty"`
	PrtLctn     *common.Max35Text            `PrtLctn,omitempty" json:",omitempty"`
	Sgntr       []common.Max70Text           `Sgntr,omitempty" json:",omitempty"`
}

func (r Cheque11) Validate() error {
	return utils.Validate(&r)
}

type ChequeDeliveryMethod1Choice struct {
	Cd    *ChequeDelivery1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text    `Prtry,omitempty" json:",omitempty"`
}

func (r ChequeDeliveryMethod1Choice) Validate() error {
	return utils.Validate(&r)
}

type CreditTransferTransaction35 struct {
	PmtId           PaymentIdentification6                        `PmtId"`
	PmtTpInf        *PaymentTypeInformation26                     `PmtTpInf,omitempty" json:",omitempty"`
	PmtCond         *PaymentCondition1                            `PmtCond,omitempty" json:",omitempty"`
	Amt             AmountType4Choice                             `Amt"`
	ChrgBr          ChargeBearerType1Code                         `ChrgBr"`
	ChqInstr        *Cheque11                                     `ChqInstr,omitempty" json:",omitempty"`
	UltmtDbtr       *PartyIdentification135                       `UltmtDbtr,omitempty" json:",omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification6 `IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification6 `IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification6 `IntrmyAgt3,omitempty" json:",omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification6  `CdtrAgt"`
	Cdtr            PartyIdentification135                        `Cdtr"`
	CdtrAcct        *CashAccount38                                `CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr       *PartyIdentification135                       `UltmtCdtr,omitempty" json:",omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                `InstrForCdtrAgt,omitempty" json:",omitempty"`
	Purp            *Purpose2Choice                               `Purp,omitempty" json:",omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `RgltryRptg,omitempty" json:",omitempty"`
	Tax             *TaxInformation8                              `Tax,omitempty" json:",omitempty"`
	RltdRmtInf      []RemittanceLocation7                         `RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf          *RemittanceInformation16                      `RmtInf,omitempty" json:",omitempty"`
	NclsdFile       []Document12                                  `NclsdFile,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                          `SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditTransferTransaction35) Validate() error {
	return utils.Validate(&r)
}

type CreditorPaymentActivationRequestV07 struct {
	XMLName     xml.Name               `xml:"CdtrPmtActvtnReq"`
	GrpHdr      GroupHeader78          `GrpHdr"`
	PmtInf      []PaymentInstruction31 `PmtInf"`
	SplmtryData []SupplementaryData1   `SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorPaymentActivationRequestV07) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader78 struct {
	MsgId    common.Max35Text        `MsgId"`
	CreDtTm  common.ISODateTime      `CreDtTm"`
	NbOfTxs  common.Max15NumericText `NbOfTxs"`
	CtrlSum  float64                 `CtrlSum,omitempty" json:",omitempty"`
	InitgPty PartyIdentification135  `InitgPty"`
}

func (r GroupHeader78) Validate() error {
	return utils.Validate(&r)
}

type InstructionForCreditorAgent1 struct {
	Cd       *Instruction3Code  `Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForCreditorAgent1) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `Nm"`
	Adr PostalAddress24   `Adr"`
}

func (r NameAndAddress16) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification6 struct {
	InstrId    *common.Max35Text        `InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text         `EndToEndId"`
	UETR       *common.UUIDv4Identifier `UETR,omitempty" json:",omitempty"`
}

func (r PaymentIdentification6) Validate() error {
	return utils.Validate(&r)
}

type PaymentInstruction31 struct {
	PmtInfId    *common.Max35Text                            `PmtInfId,omitempty" json:",omitempty"`
	PmtMtd      PaymentMethod7Code                           `PmtMtd"`
	PmtTpInf    *PaymentTypeInformation26                    `PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt DateAndDateTime2Choice                       `ReqdExctnDt"`
	XpryDt      *DateAndDateTime2Choice                      `XpryDt,omitempty" json:",omitempty"`
	PmtCond     *PaymentCondition1                           `PmtCond,omitempty" json:",omitempty"`
	Dbtr        PartyIdentification135                       `Dbtr"`
	DbtrAcct    *CashAccount38                               `DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification6 `DbtrAgt"`
	UltmtDbtr   *PartyIdentification135                      `UltmtDbtr,omitempty" json:",omitempty"`
	ChrgBr      *ChargeBearerType1Code                       `ChrgBr,omitempty" json:",omitempty"`
	CdtTrfTx    []CreditTransferTransaction35                `CdtTrfTx,omitempty" json:",omitempty"`
}

func (r PaymentInstruction31) Validate() error {
	return utils.Validate(&r)
}

type Purpose2Choice struct {
	Cd    *ExternalPurpose1Code `Cd,omitempty" json:",omitempty"`
	Prtry *common.Max35Text     `Prtry,omitempty" json:",omitempty"`
}

func (r Purpose2Choice) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryAuthority2 struct {
	Nm   *common.Max140Text  `Nm,omitempty" json:",omitempty"`
	Ctry *common.CountryCode `Ctry,omitempty" json:",omitempty"`
}

func (r RegulatoryAuthority2) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd *RegulatoryReportingType1Code    `DbtCdtRptgInd,omitempty" json:",omitempty"`
	Authrty       *RegulatoryAuthority2            `Authrty,omitempty" json:",omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `Dtls,omitempty" json:",omitempty"`
}

func (r RegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocation7 struct {
	RmtId       *common.Max35Text         `RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `RmtLctnDtls,omitempty" json:",omitempty"`
}

func (r RemittanceLocation7) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `Mtd"`
	ElctrncAdr *common.Max2048Text           `ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    *NameAndAddress16             `PstlAdr,omitempty" json:",omitempty"`
}

func (r RemittanceLocationData1) Validate() error {
	return utils.Validate(&r)
}

type StructuredRegulatoryReporting3 struct {
	Tp   *common.Max35Text                  `Tp,omitempty" json:",omitempty"`
	Dt   *common.ISODate                    `Dt,omitempty" json:",omitempty"`
	Ctry *common.CountryCode                `Ctry,omitempty" json:",omitempty"`
	Cd   *common.Max10Text                  `Cd,omitempty" json:",omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `Amt,omitempty" json:",omitempty"`
	Inf  []common.Max35Text                 `Inf,omitempty" json:",omitempty"`
}

func (r StructuredRegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `Dbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                       `Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation8) Validate() error {
	return utils.Validate(&r)
}
