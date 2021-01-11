// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v10

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type AdviceType1 struct {
	CdtAdvc AdviceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtAdvc,omitempty"`
	DbtAdvc AdviceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DbtAdvc,omitempty"`
}

type AdviceType1Choice struct {
	Cd    AdviceType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 EqvtAmt"`
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 LEI,omitempty"`
	Nm      common.Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
	PstlAdr PostalAddress24      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	Tp   CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type Cheque11 struct {
	ChqTp       ChequeType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChqTp,omitempty"`
	ChqNb       common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChqNb,omitempty"`
	ChqFr       NameAndAddress16            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChqFr,omitempty"`
	DlvryMtd    ChequeDeliveryMethod1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DlvryMtd,omitempty"`
	DlvrTo      NameAndAddress16            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DlvrTo,omitempty"`
	InstrPrty   Priority2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrPrty,omitempty"`
	ChqMtrtyDt  common.ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChqMtrtyDt,omitempty"`
	FrmsCd      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FrmsCd,omitempty"`
	MemoFld     []common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MemoFld,omitempty"`
	RgnlClrZone common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RgnlClrZone,omitempty"`
	PrtLctn     common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PrtLctn,omitempty"`
	Sgntr       []common.Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Sgntr,omitempty"`
}

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MmbId"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
	PhneNb    common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PrefrdMtd,omitempty"`
}

type CreditTransferMandateData1 struct {
	MndtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MndtId,omitempty"`
	Tp           MandateTypeInformation2   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	DtOfSgntr    common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DtOfSgntr,omitempty"`
	DtOfVrfctn   common.ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DtOfVrfctn,omitempty"`
	ElctrncSgntr common.Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ElctrncSgntr,omitempty"`
	FrstPmtDt    common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FrstPmtDt,omitempty"`
	FnlPmtDt     common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FnlPmtDt,omitempty"`
	Frqcy        Frequency36Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Frqcy,omitempty"`
	Rsn          MandateSetupReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Rsn,omitempty"`
}

type CreditTransferTransaction40 struct {
	PmtId           PaymentIdentification6                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PmtId"`
	PmtTpInf        PaymentTypeInformation26                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PmtTpInf,omitempty"`
	Amt             AmountType4Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt"`
	XchgRateInf     ExchangeRate1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 XchgRateInf,omitempty"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChrgBr,omitempty"`
	MndtRltdInf     CreditTransferMandateData1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MndtRltdInf,omitempty"`
	ChqInstr        Cheque11                                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChqInstr,omitempty"`
	UltmtDbtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 UltmtDbtr,omitempty"`
	IntrmyAgt1      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct  CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct  CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct  CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 IntrmyAgt3Acct,omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtrAgt,omitempty"`
	CdtrAgtAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtrAgtAcct,omitempty"`
	Cdtr            PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cdtr,omitempty"`
	CdtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtrAcct,omitempty"`
	UltmtCdtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 UltmtCdtr,omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent3               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrForCdtrAgt,omitempty"`
	InstrForDbtrAgt InstructionForDebtorAgent1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrForDbtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RgltryRptg,omitempty"`
	Tax             TaxInformation8                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RltdRmtInf,omitempty"`
	RmtInf          RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RmtInf,omitempty"`
	SplmtryData     []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SplmtryData,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Ref common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdOrPrtry"`
	Issr      common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type CustomerCreditTransferInitiationV10 struct {
	GrpHdr      GroupHeader95          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 GrpHdr"`
	PmtInf      []PaymentInstruction34 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PmtInf"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SplmtryData,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dt"`
	DtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Nb     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nb,omitempty"`
	RltdDt common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	Desc common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdOrPrtry"`
	Issr      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CcyOfTrf"`
}

type ExchangeRate1 struct {
	UnitCcy  common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 UnitCcy,omitempty"`
	XchgRate float64                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 XchgRate,omitempty"`
	RateTp   ExchangeRateType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RateTp,omitempty"`
	CtrctId  common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtrctId,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 LEI,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 GrnshmtAdmstr,omitempty"`
	RefNb             common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RefNb,omitempty"`
	Dt                common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdOrPrtry"`
	Issr      common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	Issr    common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr"`
	SchmeNm common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type GroupHeader95 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Authstn,omitempty"`
	NbOfTxs  common.Max15NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 NbOfTxs"`
	CtrlSum  float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtrlSum,omitempty"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InitgPty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FwdgAgt,omitempty"`
	InitnSrc PaymentInitiationSource1                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InitnSrc,omitempty"`
}

type InstructionForCreditorAgent3 struct {
	Cd       ExternalCreditorAgentInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd,omitempty"`
	InstrInf common.Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrInf,omitempty"`
}

type InstructionForDebtorAgent1 struct {
	Cd       ExternalDebtorAgentInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd,omitempty"`
	InstrInf common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrInf,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Clssfctn,omitempty"`
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm"`
	Adr PostalAddress24   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Adr"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
	PstlAdr   PostalAddress24    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PstlAdr,omitempty"`
	Id        Party38Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtryOfRes,omitempty"`
	CtctDtls  Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtctDtls,omitempty"`
}

type PaymentIdentification6 struct {
	InstrId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrId,omitempty"`
	EndToEndId common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 EndToEndId"`
	UETR       common.UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 UETR,omitempty"`
}

type PaymentInitiationSource1 struct {
	Nm    common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm"`
	Prvdr common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prvdr,omitempty"`
	Vrsn  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Vrsn,omitempty"`
}

type PaymentInstruction34 struct {
	PmtInfId        common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PmtInfId"`
	PmtMtd          PaymentMethod3Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PmtMtd"`
	ReqdAdvcTp      AdviceType1                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ReqdAdvcTp,omitempty"`
	BtchBookg       bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 BtchBookg,omitempty"`
	NbOfTxs         common.Max15NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 NbOfTxs,omitempty"`
	CtrlSum         float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtrlSum,omitempty"`
	PmtTpInf        PaymentTypeInformation26                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PmtTpInf,omitempty"`
	ReqdExctnDt     DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ReqdExctnDt"`
	PoolgAdjstmntDt common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PoolgAdjstmntDt,omitempty"`
	Dbtr            PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dbtr"`
	DbtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DbtrAcct"`
	DbtrAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DbtrAgt"`
	DbtrAgtAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DbtrAgtAcct,omitempty"`
	InstrForDbtrAgt common.Max140Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrForDbtrAgt,omitempty"`
	UltmtDbtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 UltmtDbtr,omitempty"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChrgBr,omitempty"`
	ChrgsAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChrgsAcct,omitempty"`
	ChrgsAcctAgt    BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ChrgsAcctAgt,omitempty"`
	CdtTrfTxInf     []CreditTransferTransaction40                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtTrfTxInf"`
}

type PaymentTypeInformation26 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 InstrPrty,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Id common.Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Nb       common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
	Ctry common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dtls,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PstlAdr,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Dt   common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dt,omitempty"`
	Ctry common.CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ctry,omitempty"`
	Cd   common.Max10Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt,omitempty"`
	Inf  []common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Inf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 GrnshmtRmt,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Titl,omitempty"`
	Nm   common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 UltmtDbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Rcrd,omitempty"`
}

type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxId,omitempty"`
	RegnId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxId,omitempty"`
	RegnId  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 RegnId,omitempty"`
	TaxTp   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 FrmsCd,omitempty"`
	Prd      TaxPeriod2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prd,omitempty"`
	TaxAmt   TaxAmount2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 Amt"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlRsn,omitempty"`
	OrgnlTrckgDays   Exact2NumericText                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlTrckgDays,omitempty"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Prtry"`
}

type CustomerPaymentReversalV10 struct {
	GrpHdr             GroupHeader88                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 GrpHdr"`
	OrgnlGrpInf        OriginalGroupHeader16          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlGrpInf"`
	OrgnlPmtInfAndRvsl []OriginalPaymentInstruction37 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlPmtInfAndRvsl,omitempty"`
	SplmtryData        []SupplementaryData1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SplmtryData,omitempty"`
}

type GroupHeader88 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Authstn,omitempty"`
	NbOfTxs  common.Max15NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 NbOfTxs"`
	CtrlSum  float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CtrlSum,omitempty"`
	GrpRvsl  bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 GrpRvsl,omitempty"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 InitgPty,omitempty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 FwdgAgt,omitempty"`
	DbtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 DbtrAgt,omitempty"`
	CdtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CdtrAgt,omitempty"`
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt MandateRelatedInformation14 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 DrctDbtMndt,omitempty"`
	CdtTrfMndt  CreditTransferMandateData1  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CdtTrfMndt,omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 MndtId,omitempty"`
	DtOfSgntr     common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 DtOfSgntr,omitempty"`
	AmdmntInd     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ElctrncSgntr,omitempty"`
	FrstColltnDt  common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 FrstColltnDt,omitempty"`
	FnlColltnDt   common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Rsn,omitempty"`
	TrckgDays     Exact2NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 TrckgDays,omitempty"`
}

type OriginalGroupHeader16 struct {
	OrgnlMsgId   common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlMsgNmId"`
	OrgnlCreDtTm common.ISODateTime       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlCreDtTm,omitempty"`
	RvslRsnInf   []PaymentReversalReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RvslRsnInf,omitempty"`
}

type OriginalPaymentInstruction37 struct {
	RvslPmtInfId  common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RvslPmtInfId,omitempty"`
	OrgnlPmtInfId common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlPmtInfId"`
	OrgnlNbOfTxs  common.Max15NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlCtrlSum,omitempty"`
	BtchBookg     bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 BtchBookg,omitempty"`
	PmtInfRvsl    bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 PmtInfRvsl,omitempty"`
	RvslRsnInf    []PaymentReversalReason9 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RvslRsnInf,omitempty"`
	TxInf         []PaymentTransaction125  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 TxInf,omitempty"`
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Amt,omitempty"`
	IntrBkSttlmDt  common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedData1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Purp,omitempty"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Agt"`
}

type PaymentReversalReason9 struct {
	Orgtr    PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Orgtr,omitempty"`
	Rsn      ReversalReason4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Rsn,omitempty"`
	AddtlInf []common.Max105Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 AddtlInf,omitempty"`
}

type PaymentTransaction125 struct {
	RvslId          common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RvslId,omitempty"`
	OrgnlInstrId    common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlEndToEndId,omitempty"`
	OrgnlUETR       common.UUIDv4Identifier           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlUETR,omitempty"`
	OrgnlInstdAmt   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlInstdAmt,omitempty"`
	RvsdInstdAmt    ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RvsdInstdAmt,omitempty"`
	ChrgBr          ChargeBearerType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ChrgBr,omitempty"`
	RvslRsnInf      []PaymentReversalReason9          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 RvslRsnInf,omitempty"`
	OrgnlTxRef      OriginalTransactionReference31    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 OrgnlTxRef,omitempty"`
	SplmtryData     []SupplementaryData1              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SplmtryData,omitempty"`
}

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CtgyPurp,omitempty"`
}

type ReversalReason4Choice struct {
	Cd    ExternalReversalReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Cd"`
	Prtry common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 Prtry"`
}

type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SttlmMtd"`
	SttlmAcct            CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 ThrdRmbrsmntAgtAcct,omitempty"`
}
