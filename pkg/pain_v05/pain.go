// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type AuthenticationChannel1Choice struct {
	Cd    ExternalAuthenticationChannel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id,omitempty"`
	Nm      common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PstlAdr PostalAddress6    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	Tp  CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp,omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ccy,omitempty"`
	Nm  common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   common.NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 NmPrfx,omitempty"`
	Nm       common.Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PhneNb   common.PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PhneNb,omitempty"`
	MobNb    common.PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MobNb,omitempty"`
	FaxNb    common.PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FaxNb,omitempty"`
	EmailAdr common.Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 EmailAdr,omitempty"`
	Othr     common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ToDt,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       common.BICFIIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ClrSysMmbId,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PtInTm"`
}

type Frequency37Choice struct {
	Cd    Frequency10Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp"`
	PtInTm Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CntPerPrd"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type GroupHeader47 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Authstn,omitempty"`
	InitgPty PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 InitgPty,omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 InstgAgt,omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 InstdAgt,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type Mandate10 struct {
	MndtId        []common.Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtId,omitempty"`
	MndtReqId     common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtReqId"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp,omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 TrckgInd"`
	FrstColltnAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 ColltnAmt,omitempty"`
	MaxAmt        ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cdtr"`
	CdtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Dbtr"`
	DbtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DbtrAgt"`
	UltmtDbtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 UltmtDbtr,omitempty"`
	MndtRef       common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 RfrdDoc,omitempty"`
	SplmtryData   []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SplmtryData,omitempty"`
}

type MandateAdjustment1 struct {
	DtAdjstmntRuleInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DtAdjstmntRuleInd"`
	Ctgy              Frequency37Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ctgy,omitempty"`
	Amt               ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Amt,omitempty"`
	Rate              float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Rate,omitempty"`
}

type MandateAuthentication1 struct {
	MsgAuthntcnCd common.Max16Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MsgAuthntcnCd,omitempty"`
	Dt            common.ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Dt,omitempty"`
	Chanl         AuthenticationChannel1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Chanl,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type MandateInitiationRequestV05 struct {
	GrpHdr      GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 GrpHdr"`
	Mndt        []Mandate10          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Mndt"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SplmtryData,omitempty"`
}

type MandateOccurrences4 struct {
	SeqTp        SequenceType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SeqTp"`
	Frqcy        Frequency36Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Frqcy,omitempty"`
	Drtn         DatePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Drtn,omitempty"`
	FrstColltnDt common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FrstColltnDt,omitempty"`
	FnlColltnDt  common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 FnlColltnDt,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtgyPurp,omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Clssfctn,omitempty"`
}

type OrganisationIdentification8 struct {
	AnyBIC common.AnyBICIdentifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PrvtId"`
}

type PartyIdentification43 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nm,omitempty"`
	PstlAdr   PostalAddress6     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstlAdr,omitempty"`
	Id        Party11Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type PostalAddress6 struct {
	AdrTp       common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 AdrTp,omitempty"`
	Dept        common.Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Dept,omitempty"`
	SubDept     common.Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 SubDept,omitempty"`
	StrtNm      common.Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 StrtNm,omitempty"`
	BldgNb      common.Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 BldgNb,omitempty"`
	PstCd       common.Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PstCd,omitempty"`
	TwnNm       common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 TwnNm,omitempty"`
	CtrySubDvsn common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Ctry,omitempty"`
	AdrLine     []common.Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 AdrLine,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Issr,omitempty"`
}

type ReferredMandateDocument1 struct {
	Tp      ReferredDocumentType4 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Tp,omitempty"`
	Nb      common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Nb,omitempty"`
	CdtrRef common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 CdtrRef,omitempty"`
	RltdDt  common.ISODate        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 RltdDt,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type Mandate8 struct {
	MndtId        common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtId"`
	MndtReqId     common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtReqId,omitempty"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Tp,omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 TrckgInd"`
	FrstColltnAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 ColltnAmt,omitempty"`
	MaxAmt        ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Cdtr,omitempty"`
	CdtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Dbtr,omitempty"`
	DbtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 DbtrAgt,omitempty"`
	UltmtDbtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 UltmtDbtr,omitempty"`
	MndtRef       common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 RfrdDoc,omitempty"`
}

type Mandate9 struct {
	MndtId        common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtId"`
	MndtReqId     common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtReqId,omitempty"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Tp,omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 TrckgInd"`
	FrstColltnAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 ColltnAmt,omitempty"`
	MaxAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Cdtr"`
	CdtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Dbtr"`
	DbtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 DbtrAgt"`
	UltmtDbtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 UltmtDbtr,omitempty"`
	MndtRef       common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 RfrdDoc,omitempty"`
}

type MandateAmendment5 struct {
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 OrgnlMsgInf,omitempty"`
	AmdmntRsn   MandateAmendmentReason1     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 AmdmntRsn"`
	Mndt        Mandate8                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Mndt"`
	OrgnlMndt   OriginalMandate4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 OrgnlMndt"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 SplmtryData,omitempty"`
}

type MandateAmendmentReason1 struct {
	Orgtr    PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Orgtr,omitempty"`
	Rsn      MandateReason1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Rsn"`
	AddtlInf []common.Max105Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 AddtlInf,omitempty"`
}

type MandateAmendmentRequestV05 struct {
	GrpHdr            GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 GrpHdr"`
	UndrlygAmdmntDtls []MandateAmendment5  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 UndrlygAmdmntDtls"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 SplmtryData,omitempty"`
}

type MandateReason1Choice struct {
	Cd    ExternalMandateReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 Prtry"`
}

type OriginalMandate4Choice struct {
	OrgnlMndtId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 OrgnlMndtId"`
	OrgnlMndt   Mandate9         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 OrgnlMndt"`
}

type OriginalMessageInformation1 struct {
	MsgId   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MsgId"`
	MsgNmId common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MsgNmId"`
	CreDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 CreDtTm,omitempty"`
}

type MandateCancellation5 struct {
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 OrgnlMsgInf,omitempty"`
	CxlRsn      PaymentCancellationReason1  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 CxlRsn"`
	OrgnlMndt   OriginalMandate4Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 OrgnlMndt"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 SplmtryData,omitempty"`
}

type MandateCancellationRequestV05 struct {
	GrpHdr         GroupHeader47          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 GrpHdr"`
	UndrlygCxlDtls []MandateCancellation5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 UndrlygCxlDtls"`
	SplmtryData    []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 SplmtryData,omitempty"`
}

type PaymentCancellationReason1 struct {
	Orgtr    PartyIdentification43 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 Orgtr,omitempty"`
	Rsn      MandateReason1Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 Rsn"`
	AddtlInf []common.Max105Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 AddtlInf,omitempty"`
}

type AcceptanceResult6 struct {
	Accptd          bool                 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Accptd"`
	RjctRsn         MandateReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 RjctRsn,omitempty"`
	AddtlRjctRsnInf []common.Max105Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 AddtlRjctRsnInf,omitempty"`
}

type Mandate11 struct {
	MndtId        common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 MndtId,omitempty"`
	MndtReqId     common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 MndtReqId,omitempty"`
	Authntcn      MandateAuthentication1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Authntcn,omitempty"`
	Tp            MandateTypeInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Tp,omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Ocrncs,omitempty"`
	TrckgInd      bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 TrckgInd"`
	FrstColltnAmt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 FrstColltnAmt,omitempty"`
	ColltnAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 ColltnAmt,omitempty"`
	MaxAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 MaxAmt,omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Adjstmnt,omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Rsn,omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 CdtrSchmeId,omitempty"`
	Cdtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Cdtr"`
	CdtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 CdtrAcct,omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 CdtrAgt,omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 UltmtCdtr,omitempty"`
	Dbtr          PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 Dbtr"`
	DbtrAcct      CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 DbtrAcct,omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 DbtrAgt"`
	UltmtDbtr     PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 UltmtDbtr,omitempty"`
	MndtRef       common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 MndtRef,omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 RfrdDoc,omitempty"`
}

type MandateAcceptance5 struct {
	OrgnlMsgInf OriginalMessageInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 OrgnlMsgInf,omitempty"`
	AccptncRslt AcceptanceResult6           `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 AccptncRslt"`
	OrgnlMndt   OriginalMandate5Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 OrgnlMndt,omitempty"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 SplmtryData,omitempty"`
}

type MandateAcceptanceReportV05 struct {
	GrpHdr             GroupHeader47        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 GrpHdr"`
	UndrlygAccptncDtls []MandateAcceptance5 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 UndrlygAccptncDtls"`
	SplmtryData        []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 SplmtryData,omitempty"`
}

type OriginalMandate5Choice struct {
	OrgnlMndtId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 OrgnlMndtId"`
	OrgnlMndt   Mandate11        `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 OrgnlMndt"`
}
