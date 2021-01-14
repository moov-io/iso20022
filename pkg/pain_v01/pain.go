// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v01

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

type AuthenticationChannel1Choice struct {
	Cd    ExternalAuthenticationChannel1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

func (r AuthenticationChannel1Choice) Validate() error {
	return utils.Validate(&r)
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"Cd"`
	Prtry common.Max128Text         `xml:"Prtry"`
}

func (r Authorisation1Choice) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"FinInstnId"`
	BrnchId    BranchData2                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification5) Validate() error {
	return utils.Validate(&r)
}

type BranchData2 struct {
	Id      common.Max35Text  `xml:"Id,omitempty" json:",omitempty"`
	Nm      common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr PostalAddress6    `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData2) Validate() error {
	return utils.Validate(&r)
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice        `xml:"Id"`
	Tp  CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm  common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
}

func (r CashAccount24) Validate() error {
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

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                    `xml:"MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return utils.Validate(&r)
}

type ContactDetails2 struct {
	NmPrfx   common.NamePrefix1Code `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm       common.Max140Text      `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb   common.PhoneNumber     `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb    common.PhoneNumber     `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb    common.PhoneNumber     `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr common.Max2048Text     `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr     common.Max35Text       `xml:"Othr,omitempty" json:",omitempty"`
}

func (r ContactDetails2) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

func (r DatePeriodDetails1) Validate() error {
	return utils.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification8 struct {
	BICFI       common.BICFIIdentifier              `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	Nm          common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     PostalAddress6                      `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification8) Validate() error {
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

type Frequency37Choice struct {
	Cd    Frequency10Code  `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

func (r Frequency37Choice) Validate() error {
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

type GroupHeader47 struct {
	MsgId    common.Max35Text                             `xml:"MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"Authstn,omitempty" json:",omitempty"`
	InitgPty PartyIdentification43                        `xml:"InitgPty,omitempty" json:",omitempty"`
	InstgAgt BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty" json:",omitempty"`
	InstdAgt BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty" json:",omitempty"`
}

func (r GroupHeader47) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
}

type Mandate9 struct {
	MndtId        common.Max35Text                             `xml:"MndtId"`
	MndtReqId     common.Max35Text                             `xml:"MndtReqId,omitempty" json:",omitempty"`
	Authntcn      MandateAuthentication1                       `xml:"Authntcn,omitempty" json:",omitempty"`
	Tp            MandateTypeInformation2                      `xml:"Tp,omitempty" json:",omitempty"`
	Ocrncs        MandateOccurrences4                          `xml:"Ocrncs,omitempty" json:",omitempty"`
	TrckgInd      bool                                         `xml:"TrckgInd"`
	FrstColltnAmt ActiveOrHistoricCurrencyAndAmount            `xml:"FrstColltnAmt,omitempty" json:",omitempty"`
	ColltnAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"ColltnAmt,omitempty" json:",omitempty"`
	MaxAmt        ActiveOrHistoricCurrencyAndAmount            `xml:"MaxAmt,omitempty" json:",omitempty"`
	Adjstmnt      MandateAdjustment1                           `xml:"Adjstmnt,omitempty" json:",omitempty"`
	Rsn           MandateSetupReason1Choice                    `xml:"Rsn,omitempty" json:",omitempty"`
	CdtrSchmeId   PartyIdentification43                        `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	Cdtr          PartyIdentification43                        `xml:"Cdtr"`
	CdtrAcct      CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	CdtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	UltmtCdtr     PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Dbtr          PartyIdentification43                        `xml:"Dbtr"`
	DbtrAcct      CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`
	UltmtDbtr     PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	MndtRef       common.Max35Text                             `xml:"MndtRef,omitempty" json:",omitempty"`
	RfrdDoc       []ReferredMandateDocument1                   `xml:"RfrdDoc,omitempty" json:",omitempty"`
}

func (r Mandate9) Validate() error {
	return utils.Validate(&r)
}

type MandateAdjustment1 struct {
	DtAdjstmntRuleInd bool                    `xml:"DtAdjstmntRuleInd"`
	Ctgy              Frequency37Choice       `xml:"Ctgy,omitempty" json:",omitempty"`
	Amt               ActiveCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Rate              float64                 `xml:"Rate,omitempty" json:",omitempty"`
}

func (r MandateAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type MandateAuthentication1 struct {
	MsgAuthntcnCd common.Max16Text             `xml:"MsgAuthntcnCd,omitempty" json:",omitempty"`
	Dt            common.ISODate               `xml:"Dt,omitempty" json:",omitempty"`
	Chanl         AuthenticationChannel1Choice `xml:"Chanl,omitempty" json:",omitempty"`
}

func (r MandateAuthentication1) Validate() error {
	return utils.Validate(&r)
}

type MandateClassification1Choice struct {
	Cd    common.MandateClassification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r MandateClassification1Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateCopy1 struct {
	OrgnlMsgInf OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty" json:",omitempty"`
	OrgnlMndt   OriginalMandate4Choice      `xml:"OrgnlMndt"`
	MndtSts     MandateStatus1Choice        `xml:"MndtSts,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateCopy1) Validate() error {
	return utils.Validate(&r)
}

type MandateCopyRequestV01 struct {
	GrpHdr            GroupHeader47        `xml:"GrpHdr"`
	UndrlygCpyReqDtls []MandateCopy1       `xml:"UndrlygCpyReqDtls"`
	SplmtryData       []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateCopyRequestV01) Validate() error {
	return utils.Validate(&r)
}

type MandateOccurrences4 struct {
	SeqTp        SequenceType2Code  `xml:"SeqTp"`
	Frqcy        Frequency36Choice  `xml:"Frqcy,omitempty" json:",omitempty"`
	Drtn         DatePeriodDetails1 `xml:"Drtn,omitempty" json:",omitempty"`
	FrstColltnDt common.ISODate     `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt  common.ISODate     `xml:"FnlColltnDt,omitempty" json:",omitempty"`
}

func (r MandateOccurrences4) Validate() error {
	return utils.Validate(&r)
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

func (r MandateSetupReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateStatus1Choice struct {
	Cd    ExternalMandateStatus1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

func (r MandateStatus1Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateTypeInformation2 struct {
	SvcLvl    ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
}

func (r MandateTypeInformation2) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification8 struct {
	AnyBIC common.AnyBICIdentifier              `xml:"AnyBIC,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification8) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type OriginalMandate4Choice struct {
	OrgnlMndtId common.Max35Text `xml:"OrgnlMndtId"`
	OrgnlMndt   Mandate9         `xml:"OrgnlMndt"`
}

func (r OriginalMandate4Choice) Validate() error {
	return utils.Validate(&r)
}

type OriginalMessageInformation1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	MsgNmId common.Max35Text   `xml:"MsgNmId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalMessageInformation1) Validate() error {
	return utils.Validate(&r)
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"OrgId"`
	PrvtId PersonIdentification5       `xml:"PrvtId"`
}

func (r Party11Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification43 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   PostalAddress6     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        Party11Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  ContactDetails2    `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification43) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification5) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress6 struct {
	AdrTp       common.AddressType2Code `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        common.Max70Text        `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     common.Max70Text        `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      common.Max70Text        `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      common.Max16Text        `xml:"BldgNb,omitempty" json:",omitempty"`
	PstCd       common.Max16Text        `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       common.Max35Text        `xml:"TwnNm,omitempty" json:",omitempty"`
	CtrySubDvsn common.Max35Text        `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        common.CountryCode      `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text      `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress6) Validate() error {
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

type ReferredMandateDocument1 struct {
	Tp      ReferredDocumentType4 `xml:"Tp,omitempty" json:",omitempty"`
	Nb      common.Max35Text      `xml:"Nb,omitempty" json:",omitempty"`
	CdtrRef common.Max35Text      `xml:"CdtrRef,omitempty" json:",omitempty"`
	RltdDt  common.ISODate        `xml:"RltdDt,omitempty" json:",omitempty"`
}

func (r ReferredMandateDocument1) Validate() error {
	return utils.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceLevel8Choice) Validate() error {
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

type MandateSuspension1 struct {
	SspnsnReqId common.Max35Text            `xml:"SspnsnReqId"`
	OrgnlMsgInf OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty" json:",omitempty"`
	SspnsnRsn   MandateSuspensionReason1    `xml:"SspnsnRsn"`
	OrgnlMndt   OriginalMandate4Choice      `xml:"OrgnlMndt"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateSuspension1) Validate() error {
	return utils.Validate(&r)
}

type MandateSuspensionReason1 struct {
	Orgtr    PartyIdentification43          `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      MandateSuspensionReason1Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text            `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r MandateSuspensionReason1) Validate() error {
	return utils.Validate(&r)
}

type MandateSuspensionReason1Choice struct {
	Cd    ExternalMandateSuspensionReason1Code `xml:"Cd"`
	Prtry common.Max35Text                     `xml:"Prtry"`
}

func (r MandateSuspensionReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateSuspensionRequestV01 struct {
	GrpHdr            GroupHeader47        `xml:"GrpHdr"`
	UndrlygSspnsnDtls []MandateSuspension1 `xml:"UndrlygSspnsnDtls"`
	SplmtryData       []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateSuspensionRequestV01) Validate() error {
	return utils.Validate(&r)
}
