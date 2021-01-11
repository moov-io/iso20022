// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v09

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type AdviceType1 struct {
	CdtAdvc AdviceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtAdvc,omitempty"`
	DbtAdvc AdviceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DbtAdvc,omitempty"`
}

type AdviceType1Choice struct {
	Cd    AdviceType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlRsn,omitempty"`
	OrgnlTrckgDays   Exact2NumericText                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgnlTrckgDays,omitempty"`
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 LEI,omitempty"`
	Nm      common.Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
	PstlAdr PostalAddress24      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	Tp   CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.ax35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 MmbId"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
	PhneNb    common.honeNumber           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PrefrdMtd,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Ref common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdOrPrtry"`
	Issr      common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type CustomerDirectDebitInitiationV09 struct {
	GrpHdr      GroupHeader83          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 GrpHdr"`
	PmtInf      []PaymentInstruction37 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PmtInf"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SplmtryData,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ToDt"`
}

type DirectDebitTransaction10 struct {
	MndtRltdInf MandateRelatedInformation14 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 MndtRltdInf,omitempty"`
	CdtrSchmeId PartyIdentification135      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtrSchmeId,omitempty"`
	PreNtfctnId common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PreNtfctnId,omitempty"`
	PreNtfctnDt common.ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PreNtfctnDt,omitempty"`
}

type DirectDebitTransactionInformation23 struct {
	PmtId           PaymentIdentification6                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PmtId"`
	PmtTpInf        PaymentTypeInformation29                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PmtTpInf,omitempty"`
	InstdAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 InstdAmt"`
	ChrgBr          ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ChrgBr,omitempty"`
	DrctDbtTx       DirectDebitTransaction10                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DrctDbtTx,omitempty"`
	UltmtCdtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 UltmtCdtr,omitempty"`
	DbtrAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DbtrAgt"`
	DbtrAgtAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DbtrAgtAcct,omitempty"`
	Dbtr            PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dbtr"`
	DbtrAcct        CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DbtrAcct"`
	UltmtDbtr       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 UltmtDbtr,omitempty"`
	InstrForCdtrAgt common.Max140Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 InstrForCdtrAgt,omitempty"`
	Purp            Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Purp,omitempty"`
	RgltryRptg      []RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RgltryRptg,omitempty"`
	Tax             TaxInformation8                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tax,omitempty"`
	RltdRmtInf      []RemittanceLocation7                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RltdRmtInf,omitempty"`
	RmtInf          RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RmtInf,omitempty"`
	SplmtryData     []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SplmtryData,omitempty"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Nb     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nb,omitempty"`
	RltdDt common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	Desc common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdOrPrtry"`
	Issr      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 LEI,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 GrnshmtAdmstr,omitempty"`
	RefNb             common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RefNb,omitempty"`
	Dt                common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdOrPrtry"`
	Issr      common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	Issr    common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr"`
	SchmeNm common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type GroupHeader83 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Authstn,omitempty"`
	NbOfTxs  common.Max15NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 NbOfTxs"`
	CtrlSum  float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtrlSum,omitempty"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 InitgPty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FwdgAgt,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type MandateRelatedInformation14 struct {
	MndtId        common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 MndtId,omitempty"`
	DtOfSgntr     common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DtOfSgntr,omitempty"`
	AmdmntInd     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ElctrncSgntr,omitempty"`
	FrstColltnDt  common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FrstColltnDt,omitempty"`
	FnlColltnDt   common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Rsn,omitempty"`
	TrckgDays     Exact2NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TrckgDays,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm"`
	Adr PostalAddress24   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Adr"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
	PstlAdr   PostalAddress24    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PstlAdr,omitempty"`
	Id        Party38Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtryOfRes,omitempty"`
	CtctDtls  Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtctDtls,omitempty"`
}

type PaymentIdentification6 struct {
	InstrId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 InstrId,omitempty"`
	EndToEndId common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 EndToEndId"`
	UETR       common.UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 UETR,omitempty"`
}

type PaymentInstruction37 struct {
	PmtInfId     common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PmtInfId"`
	PmtMtd       PaymentMethod2Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PmtMtd"`
	ReqdAdvcTp   AdviceType1                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ReqdAdvcTp,omitempty"`
	BtchBookg    bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 BtchBookg,omitempty"`
	NbOfTxs      common.Max15NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 NbOfTxs,omitempty"`
	CtrlSum      float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtrlSum,omitempty"`
	PmtTpInf     PaymentTypeInformation29                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PmtTpInf,omitempty"`
	ReqdColltnDt common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ReqdColltnDt"`
	Cdtr         PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cdtr"`
	CdtrAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtrAcct"`
	CdtrAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtrAgt"`
	CdtrAgtAcct  CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtrAgtAcct,omitempty"`
	UltmtCdtr    PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 UltmtCdtr,omitempty"`
	ChrgBr       ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ChrgBr,omitempty"`
	ChrgsAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ChrgsAcct,omitempty"`
	ChrgsAcctAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ChrgsAcctAgt,omitempty"`
	CdtrSchmeId  PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtrSchmeId,omitempty"`
	DrctDbtTxInf []DirectDebitTransactionInformation23        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DrctDbtTxInf"`
}

type PaymentTypeInformation29 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 InstrPrty,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Id common.Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Nb       common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
	Ctry common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dtls,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PstlAdr,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Dt   common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dt,omitempty"`
	Ctry common.CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ctry,omitempty"`
	Cd   common.Max10Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Amt,omitempty"`
	Inf  []common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Inf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 GrnshmtRmt,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Titl,omitempty"`
	Nm   common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 UltmtDbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Rcrd,omitempty"`
}

type TaxInformation8 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxId,omitempty"`
	RegnId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxId,omitempty"`
	RegnId  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 RegnId,omitempty"`
	TaxTp   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 FrmsCd,omitempty"`
	Prd      TaxPeriod2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prd,omitempty"`
	TaxAmt   TaxAmount2        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09 Amt"`
}
