// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v11

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
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
	Cd    common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlRsn,omitempty"`
	OrgnlTrckgDays   common.Exact2NumericText                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlTrckgDays,omitempty"`
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 EqvtAmt"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 LEI,omitempty"`
	Nm      common.Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nm,omitempty"`
	PstlAdr PostalAddress24      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	Tp   CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nm,omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Agt"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MmbId"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nm,omitempty"`
	PhneNb    common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PrefrdMtd,omitempty"`
}

type CreditTransferMandateData1 struct {
	MndtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MndtId,omitempty"`
	Tp           MandateTypeInformation2   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	DtOfSgntr    common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtOfSgntr,omitempty"`
	DtOfVrfctn   common.ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtOfVrfctn,omitempty"`
	ElctrncSgntr common.Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ElctrncSgntr,omitempty"`
	FrstPmtDt    common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FrstPmtDt,omitempty"`
	FnlPmtDt     common.ISODate            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FnlPmtDt,omitempty"`
	Frqcy        Frequency36Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Frqcy,omitempty"`
	Rsn          MandateSetupReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rsn,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Ref common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdOrPrtry"`
	Issr      common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type CurrencyExchange13 struct {
	SrcCcy   common.ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SrcCcy"`
	TrgtCcy  common.ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TrgtCcy"`
	XchgRate float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 XchgRate"`
	UnitCcy  common.ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 UnitCcy,omitempty"`
}

type CustomerPaymentStatusReportV11 struct {
	GrpHdr            GroupHeader86                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupHeader17          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction38 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlPmtInfAndSts,omitempty"`
	SplmtryData       []SupplementaryData1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SplmtryData,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dt"`
	DtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Nb     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nb,omitempty"`
	RltdDt common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	Desc common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdOrPrtry"`
	Issr      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CcyOfTrf"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 LEI,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp"`
	PtInTm common.Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 GrnshmtAdmstr,omitempty"`
	RefNb             common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RefNb,omitempty"`
	Dt                common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdOrPrtry"`
	Issr      common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	Issr    common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr"`
	SchmeNm common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type GroupHeader86 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CreDtTm"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InitgPty,omitempty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FwdgAgt,omitempty"`
	DbtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DbtrAgt,omitempty"`
	CdtrAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtrAgt,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type MandateRelatedData1Choice struct {
	DrctDbtMndt MandateRelatedInformation14 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DrctDbtMndt,omitempty"`
	CdtTrfMndt  CreditTransferMandateData1  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtTrfMndt,omitempty"`
}

type MandateRelatedInformation14 struct {
	MndtId        common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MndtId,omitempty"`
	DtOfSgntr     common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtOfSgntr,omitempty"`
	AmdmntInd     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ElctrncSgntr,omitempty"`
	FrstColltnDt  common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FrstColltnDt,omitempty"`
	FnlColltnDt   common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rsn,omitempty"`
	TrckgDays     common.Exact2NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TrckgDays,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Clssfctn,omitempty"`
}

type NumberOfTransactionsPerStatus5 struct {
	DtldNbOfTxs common.Max15NumericText               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtldNbOfTxs"`
	DtldSts     ExternalPaymentTransactionStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtldSts"`
	DtldCtrlSum float64                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtldCtrlSum,omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type OriginalGroupHeader17 struct {
	OrgnlMsgId    common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlMsgNmId"`
	OrgnlCreDtTm  common.ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlCreDtTm,omitempty"`
	OrgnlNbOfTxs  common.Max15NumericText          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlCtrlSum,omitempty"`
	GrpSts        ExternalPaymentGroupStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 GrpSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 NbOfTxsPerSts,omitempty"`
}

type OriginalPaymentInstruction38 struct {
	OrgnlPmtInfId common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlPmtInfId"`
	OrgnlNbOfTxs  common.Max15NumericText          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlNbOfTxs,omitempty"`
	OrgnlCtrlSum  float64                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlCtrlSum,omitempty"`
	PmtInfSts     ExternalPaymentGroupStatus1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PmtInfSts,omitempty"`
	StsRsnInf     []StatusReasonInformation12      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 StsRsnInf,omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 NbOfTxsPerSts,omitempty"`
	TxInfAndSts   []PaymentTransaction126          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TxInfAndSts,omitempty"`
}

type OriginalTransactionReference31 struct {
	IntrBkSttlmAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 IntrBkSttlmAmt,omitempty"`
	Amt            AmountType4Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt,omitempty"`
	IntrBkSttlmDt  common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 IntrBkSttlmDt,omitempty"`
	ReqdColltnDt   common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ReqdColltnDt,omitempty"`
	ReqdExctnDt    DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ReqdExctnDt,omitempty"`
	CdtrSchmeId    PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtrSchmeId,omitempty"`
	SttlmInf       SettlementInstruction7                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SttlmInf,omitempty"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PmtTpInf,omitempty"`
	PmtMtd         PaymentMethod4Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PmtMtd,omitempty"`
	MndtRltdInf    MandateRelatedData1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 MndtRltdInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RmtInf,omitempty"`
	UltmtDbtr      Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 UltmtDbtr,omitempty"`
	Dbtr           Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dbtr,omitempty"`
	DbtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DbtrAcct,omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DbtrAgt,omitempty"`
	DbtrAgtAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DbtrAgtAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtrAgt,omitempty"`
	CdtrAgtAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtrAgtAcct,omitempty"`
	Cdtr           Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cdtr,omitempty"`
	CdtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtrAcct,omitempty"`
	UltmtCdtr      Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 UltmtCdtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Purp,omitempty"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Agt"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nm,omitempty"`
	PstlAdr   PostalAddress24    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PstlAdr,omitempty"`
	Id        Party38Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtryOfRes,omitempty"`
	CtctDtls  Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtctDtls,omitempty"`
}

type PaymentTransaction126 struct {
	StsId           common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 StsId,omitempty"`
	OrgnlInstrId    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlEndToEndId,omitempty"`
	OrgnlUETR       common.UUIDv4Identifier               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlUETR,omitempty"`
	TxSts           ExternalPaymentTransactionStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TxSts,omitempty"`
	StsRsnInf       []StatusReasonInformation12           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 StsRsnInf,omitempty"`
	ChrgsInf        []Charges7                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ChrgsInf,omitempty"`
	TrckrData       TrackerData1                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TrckrData,omitempty"`
	AccptncDtTm     common.ISODateTime                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AccptncDtTm,omitempty"`
	AcctSvcrRef     common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AcctSvcrRef,omitempty"`
	ClrSysRef       common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ClrSysRef,omitempty"`
	OrgnlTxRef      OriginalTransactionReference31        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 OrgnlTxRef,omitempty"`
	SplmtryData     []SupplementaryData1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SplmtryData,omitempty"`
}

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Id common.Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Nb       common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Issr,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Strd,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type SettlementInstruction7 struct {
	SttlmMtd             SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SttlmMtd"`
	SttlmAcct            CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SttlmAcct,omitempty"`
	ClrSys               ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ClrSys,omitempty"`
	InstgRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ThrdRmbrsmntAgtAcct,omitempty"`
}

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type StatusReasonInformation12 struct {
	Orgtr    PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Orgtr,omitempty"`
	Rsn      StatusReason6Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rsn,omitempty"`
	AddtlInf []common.Max105Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AddtlInf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 GrnshmtRmt,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Titl,omitempty"`
	Nm   common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 UltmtDbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxId,omitempty"`
	RegnId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxId,omitempty"`
	RegnId  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 RegnId,omitempty"`
	TaxTp   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 FrmsCd,omitempty"`
	Prd      TaxPeriod2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prd,omitempty"`
	TaxAmt   TaxAmount2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Amt"`
}

type TrackerData1 struct {
	ConfdDt   DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ConfdDt"`
	ConfdAmt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ConfdAmt"`
	TrckrRcrd []TrackerRecord1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 TrckrRcrd"`
}

type TrackerRecord1 struct {
	Agt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 Agt"`
	ChrgBr       ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ChrgBr,omitempty"`
	ChrgsAmt     ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 ChrgsAmt,omitempty"`
	XchgRateData CurrencyExchange13                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.002.001.11 XchgRateData,omitempty"`
}
