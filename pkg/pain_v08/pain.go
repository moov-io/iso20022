// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v08

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type AdviceType1 struct {
	CdtAdvc AdviceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtAdvc,omitempty"`
	DbtAdvc AdviceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DbtAdvc,omitempty"`
}

type AdviceType1Choice struct {
	Cd    AdviceType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type AmountOrRate1Choice struct {
	Amt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
	Rate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rate"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 EqvtAmt"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LEI,omitempty"`
	Nm      common.Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	Tp   CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type Cheque11 struct {
	ChqTp       ChequeType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChqTp,omitempty"`
	ChqNb       common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChqNb,omitempty"`
	ChqFr       NameAndAddress16            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChqFr,omitempty"`
	DlvryMtd    ChequeDeliveryMethod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DlvryMtd,omitempty"`
	DlvrTo      NameAndAddress16            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DlvrTo,omitempty"`
	InstrPrty   Priority2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InstrPrty,omitempty"`
	ChqMtrtyDt  common.ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChqMtrtyDt,omitempty"`
	FrmsCd      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FrmsCd,omitempty"`
	MemoFld     []common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MemoFld,omitempty"`
	RgnlClrZone common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RgnlClrZone,omitempty"`
	PrtLctn     common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PrtLctn,omitempty"`
	Sgntr       []common.Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Sgntr,omitempty"`
}

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	PhneNb    common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PrefrdMtd,omitempty"`
}

type CreditTransferMandateData1 struct {
	MndtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MndtId,omitempty"`
	Tp           MandateTypeInformation2   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	DtOfSgntr    common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DtOfSgntr,omitempty"`
	DtOfVrfctn   common.ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DtOfVrfctn,omitempty"`
	ElctrncSgntr common.Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ElctrncSgntr,omitempty"`
	FrstPmtDt    common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FrstPmtDt,omitempty"`
	FnlPmtDt     common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FnlPmtDt,omitempty"`
	Frqcy        Frequency36Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Frqcy,omitempty"`
	Rsn          MandateSetupReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rsn,omitempty"`
}

type CreditTransferTransaction42 struct {
	PmtId           PaymentIdentification6                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtId"`
	PmtTpInf        PaymentTypeInformation26                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtTpInf,omitempty"`
	PmtCond         PaymentCondition1                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtCond,omitempty"`
	Amt             AmountType4Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChrgBr"`
	MndtRltdInf     CreditTransferMandateData1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MndtRltdInf,omitempty"`
	ChqInstr        Cheque11                                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChqInstr,omitempty"`
	UltmtDbtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 UltmtDbtr,omitempty"`
	IntrmyAgt1      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 IntrmyAgt1,omitempty"`
	IntrmyAgt2      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 IntrmyAgt2,omitempty"`
	IntrmyAgt3      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 IntrmyAgt3,omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtrAgt"`
	Cdtr            PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cdtr"`
	CdtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtrAcct,omitempty"`
	UltmtCdtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent3               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InstrForCdtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RgltryRptg,omitempty"`
	Tax             TaxInformation8                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RltdRmtInf,omitempty"`
	RmtInf          RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RmtInf,omitempty"`
	NclsdFile       []Document12                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 NclsdFile,omitempty"`
	SplmtryData     []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SplmtryData,omitempty"`
}

type CreditorPaymentActivationRequestV08 struct {
	GrpHdr      GroupHeader78          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 GrpHdr"`
	PmtInf      []PaymentInstruction35 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtInf"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SplmtryData,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Ref common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdOrPrtry"`
	Issr      common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dt"`
	DtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type Document12 struct {
	Tp        DocumentType1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp"`
	Id        common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	IsseDt    DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 IsseDt"`
	Nm        common.Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	LangCd    string                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LangCd,omitempty"`
	Frmt      DocumentFormat1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Frmt"`
	FileNm    common.Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FileNm,omitempty"`
	DgtlSgntr PartyAndSignature3     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DgtlSgntr,omitempty"`
	Nclsr     common.Max10MbBinary   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nclsr"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AddtlInf,omitempty"`
}

type DocumentFormat1Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry GenericIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Nb     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nb,omitempty"`
	RltdDt common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	Desc common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdOrPrtry"`
	Issr      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry GenericIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CcyOfTrf"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LEI,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 GrnshmtAdmstr,omitempty"`
	RefNb             common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RefNb,omitempty"`
	Dt                common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdOrPrtry"`
	Issr      common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	SchmeNm common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	Issr    common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr"`
	SchmeNm common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type GroupHeader78 struct {
	MsgId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 MsgId"`
	CreDtTm  common.ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CreDtTm"`
	NbOfTxs  common.Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 NbOfTxs"`
	CtrlSum  float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtrlSum,omitempty"`
	InitgPty PartyIdentification135  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InitgPty"`
}

type InstructionForCreditorAgent3 struct {
	Cd       ExternalCreditorAgentInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd,omitempty"`
	InstrInf common.Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InstrInf,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Clssfctn,omitempty"`
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm"`
	Adr PostalAddress24   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Adr"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PrvtId"`
}

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Pty"`
	Sgntr SkipPayload            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Sgntr"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	PstlAdr   PostalAddress24    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PstlAdr,omitempty"`
	Id        Party38Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtryOfRes,omitempty"`
	CtctDtls  Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtctDtls,omitempty"`
}

type PaymentCondition1 struct {
	AmtModAllwd   bool                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AmtModAllwd"`
	EarlyPmtAllwd bool                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 EarlyPmtAllwd"`
	DelyPnlty     common.Max140Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DelyPnlty,omitempty"`
	ImdtPmtRbt    AmountOrRate1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ImdtPmtRbt,omitempty"`
	GrntedPmtReqd bool                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 GrntedPmtReqd"`
}

type PaymentIdentification6 struct {
	InstrId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InstrId,omitempty"`
	EndToEndId common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 EndToEndId"`
	UETR       common.UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 UETR,omitempty"`
}

type PaymentInstruction35 struct {
	PmtInfId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtInfId,omitempty"`
	PmtMtd      PaymentMethod7Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtMtd"`
	ReqdAdvcTp  AdviceType1                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ReqdAdvcTp,omitempty"`
	PmtTpInf    PaymentTypeInformation26                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtTpInf,omitempty"`
	ReqdExctnDt DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ReqdExctnDt"`
	XpryDt      DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 XpryDt,omitempty"`
	PmtCond     PaymentCondition1                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PmtCond,omitempty"`
	Dbtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dbtr"`
	DbtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DbtrAcct,omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DbtrAgt"`
	UltmtDbtr   PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 UltmtDbtr,omitempty"`
	ChrgBr      ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ChrgBr,omitempty"`
	CdtTrfTx    []CreditTransferTransaction42                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtTrfTx"`
}

type PaymentTypeInformation26 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 InstrPrty,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Id common.Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Nb       common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
	Ctry common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dtls,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PstlAdr,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type SkipPayload struct {
	Item string `xml:",any"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Dt   common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dt,omitempty"`
	Ctry common.CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ctry,omitempty"`
	Cd   common.Max10Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt,omitempty"`
	Inf  []common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Inf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 GrnshmtRmt,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Titl,omitempty"`
	Nm   common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 UltmtDbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rcrd,omitempty"`
}

type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxId,omitempty"`
	RegnId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxId,omitempty"`
	RegnId  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 RegnId,omitempty"`
	TaxTp   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 FrmsCd,omitempty"`
	Prd      TaxPeriod2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prd,omitempty"`
	TaxAmt   TaxAmount2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08 Amt"`
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Agt"`
}

type CreditorPaymentActivationRequestStatusReportV08 struct {
	GrpHdr            GroupHeader87                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation30     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction39 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlPmtInfAndSts,omitempty"`
	SplmtryData       []SupplementaryData1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 SplmtryData,omitempty"`
}

type GroupHeader87 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CreDtTm"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 InitgPty"`
	DbtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DbtrAgt,omitempty"`
	CdtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CdtrAgt,omitempty"`
}

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs common.Max15NumericText               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DtldSts"`
	DtldCtrlSum float64                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DtldCtrlSum,omitempty"`
}

type OriginalGroupInformation30 struct {
	OrgnlMsgId    common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlMsgNmId"`
	OrgnlCreDtTm  common.ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlCtrlSum,omitempty"`
	GrpSts        ExternalPaymentGroupStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 NbOfTxsPerSts,omitempty"`
}

type OriginalPaymentInstruction39 struct {
	OrgnlPmtInfId common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlPmtInfId"`
	OrgnlNbOfTxs  common.Max15NumericText          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlCtrlSum,omitempty"`
	PmtInfSts     ExternalPaymentGroupStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 PmtInfSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 NbOfTxsPerSts,omitempty"`
	TxInfAndSts   []PaymentTransaction128          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 TxInfAndSts,omitempty"`
}

type OriginalTransactionReference33 struct {
	Amt         AmountType4Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Amt,omitempty"`
	ReqdExctnDt DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 ReqdExctnDt,omitempty"`
	XpryDt      DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 XpryDt,omitempty"`
	PmtCond     PaymentCondition1                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 PmtCond,omitempty"`
	PmtTpInf    PaymentTypeInformation26                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 PmtTpInf,omitempty"`
	PmtMtd      PaymentMethod4Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 PmtMtd,omitempty"`
	MndtRltdInf CreditTransferMandateData1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 MndtRltdInf,omitempty"`
	RmtInf      RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 RmtInf,omitempty"`
	NclsdFile   []Document12                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 NclsdFile,omitempty"`
	UltmtDbtr   PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 UltmtDbtr,omitempty"`
	Dbtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Dbtr,omitempty"`
	DbtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DbtrAcct,omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DbtrAgt,omitempty"`
	CdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CdtrAgt"`
	Cdtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Cdtr"`
	CdtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 CdtrAcct,omitempty"`
	UltmtCdtr   PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 UltmtCdtr,omitempty"`
}

type PaymentConditionStatus1 struct {
	AccptdAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 AccptdAmt,omitempty"`
	GrntedPmt bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 GrntedPmt"`
	EarlyPmt  bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 EarlyPmt"`
}

type PaymentTransaction128 struct {
	StsId           common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 StsId,omitempty"`
	OrgnlInstrId    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlEndToEndId,omitempty"`
	OrgnlUETR       common.UUIDv4Identifier               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlUETR,omitempty"`
	TxSts           ExternalPaymentTransactionStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 TxSts,omitempty"`
	StsRsnInf       []StatusReasonInformation12           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 StsRsnInf,omitempty"`
	PmtCondSts      PaymentConditionStatus1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 PmtCondSts,omitempty"`
	ChrgsInf        []Charges7                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 ChrgsInf,omitempty"`
	DbtrDcsnDtTm    common.ISODateTime                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 DbtrDcsnDtTm,omitempty"`
	AccptncDtTm     common.ISODateTime                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 AccptncDtTm,omitempty"`
	AcctSvcrRef     common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 AcctSvcrRef,omitempty"`
	ClrSysRef       common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 ClrSysRef,omitempty"`
	OrgnlTxRef      OriginalTransactionReference33        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 OrgnlTxRef,omitempty"`
	NclsdFile       []Document12                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 NclsdFile,omitempty"`
	SplmtryData     []SupplementaryData1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 SplmtryData,omitempty"`
}

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Prtry"`
}

type StatusReasonInformation12 struct {
	Orgtr    PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Orgtr,omitempty"`
	Rsn      StatusReason6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 Rsn,omitempty"`
	AddtlInf []common.Max105Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08 AddtlInf,omitempty"`
}
