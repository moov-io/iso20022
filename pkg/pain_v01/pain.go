// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v01

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AuthenticationChannel1Choice struct {
	Cd    ExternalAuthenticationChannel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id,omitempty"`
	Nm      common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id"`
	Tp  CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Tp,omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Ccy,omitempty"`
	Nm  common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   common.NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 NmPrfx,omitempty"`
	Nm       common.Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Nm,omitempty"`
	PhneNb   common.PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PhneNb,omitempty"`
	MobNb    common.PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MobNb,omitempty"`
	FaxNb    common.PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 FaxNb,omitempty"`
	EmailAdr common.Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 EmailAdr,omitempty"`
	Othr     common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Othr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 ToDt,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       common.BICFIIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 ClrSysMmbId,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PtInTm"`
}

type Frequency37Choice struct {
	Cd    Frequency10Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CntPerPrd"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Issr,omitempty"`
}

type GroupHeader47 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Authstn,omitempty"`
	InitgPty PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 InitgPty,omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 InstdAgt,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type Mandate9 struct {
	MndtId        common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtId"`
	MndtReqId     common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtReqId,omitempty"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Tp,omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 TrckgInd"`
	FrstColltnAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 ColltnAmt,omitempty"`
	MaxAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cdtr"`
	CdtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Dbtr"`
	DbtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 DbtrAgt"`
	UltmtDbtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 UltmtDbtr,omitempty"`
	MndtRef       common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 RfrdDoc,omitempty"`
}

type MandateAdjustment1 struct {
	DtAdjstmntRuleInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 DtAdjstmntRuleInd"`
	Ctgy              Frequency37Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Ctgy,omitempty"`
	Amt               ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Amt,omitempty"`
	Rate              float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Rate,omitempty"`
}

type MandateAuthentication1 struct {
	MsgAuthntcnCd common.Max16Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MsgAuthntcnCd,omitempty"`
	Dt            common.ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Dt,omitempty"`
	Chanl         AuthenticationChannel1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Chanl,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type MandateCopy1 struct {
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 OrgnlMsgInf,omitempty"`
	OrgnlMndt   OriginalMandate4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 OrgnlMndt"`
	MndtSts     MandateStatus1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtSts,omitempty"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SplmtryData,omitempty"`
}

type MandateCopyRequestV01 struct {
	GrpHdr            GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 GrpHdr"`
	UndrlygCpyReqDtls []MandateCopy1       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 UndrlygCpyReqDtls"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SplmtryData,omitempty"`
}

type MandateOccurrences4 struct {
	SeqTp        SequenceType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SeqTp"`
	Frqcy        Frequency36Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Frqcy,omitempty"`
	Drtn         DatePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Drtn,omitempty"`
	FrstColltnDt common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 FrstColltnDt,omitempty"`
	FnlColltnDt  common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 FnlColltnDt,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type MandateStatus1Choice struct {
	Cd    ExternalMandateStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Clssfctn,omitempty"`
}

type OrganisationIdentification8 struct {
	AnyBIC common.AnyBICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type OriginalMandate4Choice struct {
	OrgnlMndtId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 OrgnlMndtId"`
	OrgnlMndt   Mandate9         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 OrgnlMndt"`
}

type OriginalMessageInformation1 struct {
	MsgId   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MsgId"`
	MsgNmId common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MsgNmId"`
	CreDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CreDtTm,omitempty"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PrvtId"`
}

type PartyIdentification43 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PstlAdr,omitempty"`
	Id        Party11Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type PostalAddress6 struct {
	AdrTp       common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 AdrTp,omitempty"`
	Dept        common.Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Dept,omitempty"`
	SubDept     common.Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 SubDept,omitempty"`
	StrtNm      common.Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 StrtNm,omitempty"`
	BldgNb      common.Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 BldgNb,omitempty"`
	PstCd       common.Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PstCd,omitempty"`
	TwnNm       common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 TwnNm,omitempty"`
	CtrySubDvsn common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Ctry,omitempty"`
	AdrLine     []common.Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 AdrLine,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Issr,omitempty"`
}

type ReferredMandateDocument1 struct {
	Tp      ReferredDocumentType4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Tp,omitempty"`
	Nb      common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Nb,omitempty"`
	CdtrRef common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 CdtrRef,omitempty"`
	RltdDt  common.ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 RltdDt,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type MandateSuspension1 struct {
	SspnsnReqId common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 SspnsnReqId"`
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 OrgnlMsgInf,omitempty"`
	SspnsnRsn   MandateSuspensionReason1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 SspnsnRsn"`
	OrgnlMndt   OriginalMandate4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 OrgnlMndt"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 SplmtryData,omitempty"`
}

type MandateSuspensionReason1 struct {
	Orgtr    PartyIdentification43          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 Orgtr,omitempty"`
	Rsn      MandateSuspensionReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 Rsn"`
	AddtlInf []common.Max105Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 AddtlInf,omitempty"`
}

type MandateSuspensionReason1Choice struct {
	Cd    ExternalMandateSuspensionReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 Cd"`
	Prtry common.Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 Prtry"`
}

type MandateSuspensionRequestV01 struct {
	GrpHdr            GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 GrpHdr"`
	UndrlygSspnsnDtls []MandateSuspension1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 UndrlygSspnsnDtls"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 SplmtryData,omitempty"`
}
