// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

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
	BrnchId    *BranchData2                        `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification5) Validate() error {
	return utils.Validate(&r)
}

type BranchData2 struct {
	Id      *common.Max35Text  `xml:"Id,omitempty" json:",omitempty"`
	Nm      *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress6    `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData2) Validate() error {
	return utils.Validate(&r)
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice         `xml:"Id"`
	Tp  *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm  *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
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
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return utils.Validate(&r)
}

type ContactDetails2 struct {
	NmPrfx   *common.NamePrefix1Code `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm       *common.Max140Text      `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb   *common.PhoneNumber     `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb    *common.PhoneNumber     `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb    *common.PhoneNumber     `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max2048Text     `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr     *common.Max35Text       `xml:"Othr,omitempty" json:",omitempty"`
}

func (r ContactDetails2) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate  `xml:"FrDt"`
	ToDt *common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
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
	BICFI       *common.BICFIIdentifier              `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
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

type GroupHeader47 struct {
	MsgId    common.Max35Text                              `xml:"MsgId"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm"`
	Authstn  []Authorisation1Choice                        `xml:"Authstn,omitempty" json:",omitempty"`
	InitgPty *PartyIdentification43                        `xml:"InitgPty,omitempty" json:",omitempty"`
	InstgAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty" json:",omitempty"`
	InstdAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty" json:",omitempty"`
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

type Mandate10 struct {
	MndtId        []common.Max35Text                            `xml:"MndtId,omitempty" json:",omitempty"`
	MndtReqId     common.Max35Text                              `xml:"MndtReqId"`
	Authntcn      *MandateAuthentication1                       `xml:"Authntcn,omitempty" json:",omitempty"`
	Tp            *MandateTypeInformation2                      `xml:"Tp,omitempty" json:",omitempty"`
	Ocrncs        *MandateOccurrences4                          `xml:"Ocrncs,omitempty" json:",omitempty"`
	TrckgInd      bool                                          `xml:"TrckgInd"`
	FrstColltnAmt *ActiveCurrencyAndAmount                      `xml:"FrstColltnAmt,omitempty" json:",omitempty"`
	ColltnAmt     *ActiveCurrencyAndAmount                      `xml:"ColltnAmt,omitempty" json:",omitempty"`
	MaxAmt        *ActiveCurrencyAndAmount                      `xml:"MaxAmt,omitempty" json:",omitempty"`
	Adjstmnt      *MandateAdjustment1                           `xml:"Adjstmnt,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice                    `xml:"Rsn,omitempty" json:",omitempty"`
	CdtrSchmeId   *PartyIdentification43                        `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	Cdtr          PartyIdentification43                         `xml:"Cdtr"`
	CdtrAcct      *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	CdtrAgt       *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	UltmtCdtr     *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Dbtr          PartyIdentification43                         `xml:"Dbtr"`
	DbtrAcct      *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5  `xml:"DbtrAgt"`
	UltmtDbtr     *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	MndtRef       *common.Max35Text                             `xml:"MndtRef,omitempty" json:",omitempty"`
	RfrdDoc       []ReferredMandateDocument1                    `xml:"RfrdDoc,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r Mandate10) Validate() error {
	return utils.Validate(&r)
}

type MandateAdjustment1 struct {
	DtAdjstmntRuleInd bool                     `xml:"DtAdjstmntRuleInd"`
	Ctgy              *Frequency37Choice       `xml:"Ctgy,omitempty" json:",omitempty"`
	Amt               *ActiveCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Rate              float64                  `xml:"Rate,omitempty" json:",omitempty"`
}

func (r MandateAdjustment1) Validate() error {
	return utils.Validate(&r)
}

type MandateAuthentication1 struct {
	MsgAuthntcnCd *common.Max16Text             `xml:"MsgAuthntcnCd,omitempty" json:",omitempty"`
	Dt            *common.ISODate               `xml:"Dt,omitempty" json:",omitempty"`
	Chanl         *AuthenticationChannel1Choice `xml:"Chanl,omitempty" json:",omitempty"`
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

type MandateInitiationRequestV05 struct {
	Attr        []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      GroupHeader47        `xml:"GrpHdr"`
	Mndt        []Mandate10          `xml:"Mndt" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateInitiationRequestV05) Validate() error {
	return utils.Validate(&r)
}

type MandateOccurrences4 struct {
	SeqTp        SequenceType2Code   `xml:"SeqTp"`
	Frqcy        *Frequency36Choice  `xml:"Frqcy,omitempty" json:",omitempty"`
	Drtn         *DatePeriodDetails1 `xml:"Drtn,omitempty" json:",omitempty"`
	FrstColltnDt *common.ISODate     `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt  *common.ISODate     `xml:"FnlColltnDt,omitempty" json:",omitempty"`
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

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"CtgyPurp,omitempty" json:",omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"Clssfctn,omitempty" json:",omitempty"`
}

func (r MandateTypeInformation2) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification8 struct {
	AnyBIC *common.AnyBICIdentifier             `xml:"AnyBIC,omitempty" json:",omitempty"`
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

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"OrgId"`
	PrvtId PersonIdentification5       `xml:"PrvtId"`
}

func (r Party11Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification43 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress6     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party11Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *ContactDetails2    `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification43) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
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
	AdrTp       *common.AddressType2Code `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text        `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text        `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text        `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text        `xml:"BldgNb,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text        `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text        `xml:"TwnNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text        `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode      `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text       `xml:"AdrLine,omitempty" json:",omitempty"`
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
	Issr      *common.Max35Text           `xml:"Issr,omitempty" json:",omitempty"`
}

func (r ReferredDocumentType4) Validate() error {
	return utils.Validate(&r)
}

type ReferredMandateDocument1 struct {
	Tp      *ReferredDocumentType4 `xml:"Tp,omitempty" json:",omitempty"`
	Nb      *common.Max35Text      `xml:"Nb,omitempty" json:",omitempty"`
	CdtrRef *common.Max35Text      `xml:"CdtrRef,omitempty" json:",omitempty"`
	RltdDt  *common.ISODate        `xml:"RltdDt,omitempty" json:",omitempty"`
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

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type Mandate8 struct {
	MndtId        common.Max35Text                              `xml:"MndtId"`
	MndtReqId     *common.Max35Text                             `xml:"MndtReqId,omitempty" json:",omitempty"`
	Authntcn      *MandateAuthentication1                       `xml:"Authntcn,omitempty" json:",omitempty"`
	Tp            *MandateTypeInformation2                      `xml:"Tp,omitempty" json:",omitempty"`
	Ocrncs        *MandateOccurrences4                          `xml:"Ocrncs,omitempty" json:",omitempty"`
	TrckgInd      bool                                          `xml:"TrckgInd"`
	FrstColltnAmt *ActiveCurrencyAndAmount                      `xml:"FrstColltnAmt,omitempty" json:",omitempty"`
	ColltnAmt     *ActiveCurrencyAndAmount                      `xml:"ColltnAmt,omitempty" json:",omitempty"`
	MaxAmt        *ActiveCurrencyAndAmount                      `xml:"MaxAmt,omitempty" json:",omitempty"`
	Adjstmnt      *MandateAdjustment1                           `xml:"Adjstmnt,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice                    `xml:"Rsn,omitempty" json:",omitempty"`
	CdtrSchmeId   *PartyIdentification43                        `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	Cdtr          *PartyIdentification43                        `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct      *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	CdtrAgt       *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	UltmtCdtr     *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Dbtr          *PartyIdentification43                        `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct      *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt       *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	UltmtDbtr     *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	MndtRef       *common.Max35Text                             `xml:"MndtRef,omitempty" json:",omitempty"`
	RfrdDoc       []ReferredMandateDocument1                    `xml:"RfrdDoc,omitempty" json:",omitempty"`
}

func (r Mandate8) Validate() error {
	return utils.Validate(&r)
}

type Mandate9 struct {
	MndtId        common.Max35Text                              `xml:"MndtId"`
	MndtReqId     *common.Max35Text                             `xml:"MndtReqId,omitempty" json:",omitempty"`
	Authntcn      *MandateAuthentication1                       `xml:"Authntcn,omitempty" json:",omitempty"`
	Tp            *MandateTypeInformation2                      `xml:"Tp,omitempty" json:",omitempty"`
	Ocrncs        *MandateOccurrences4                          `xml:"Ocrncs,omitempty" json:",omitempty"`
	TrckgInd      bool                                          `xml:"TrckgInd"`
	FrstColltnAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"FrstColltnAmt,omitempty" json:",omitempty"`
	ColltnAmt     *ActiveOrHistoricCurrencyAndAmount            `xml:"ColltnAmt,omitempty" json:",omitempty"`
	MaxAmt        *ActiveOrHistoricCurrencyAndAmount            `xml:"MaxAmt,omitempty" json:",omitempty"`
	Adjstmnt      *MandateAdjustment1                           `xml:"Adjstmnt,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice                    `xml:"Rsn,omitempty" json:",omitempty"`
	CdtrSchmeId   *PartyIdentification43                        `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	Cdtr          PartyIdentification43                         `xml:"Cdtr"`
	CdtrAcct      *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	CdtrAgt       *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	UltmtCdtr     *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Dbtr          PartyIdentification43                         `xml:"Dbtr"`
	DbtrAcct      *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5  `xml:"DbtrAgt"`
	UltmtDbtr     *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	MndtRef       *common.Max35Text                             `xml:"MndtRef,omitempty" json:",omitempty"`
	RfrdDoc       []ReferredMandateDocument1                    `xml:"RfrdDoc,omitempty" json:",omitempty"`
}

func (r Mandate9) Validate() error {
	return utils.Validate(&r)
}

type MandateAmendment5 struct {
	OrgnlMsgInf *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty" json:",omitempty"`
	AmdmntRsn   MandateAmendmentReason1      `xml:"AmdmntRsn"`
	Mndt        Mandate8                     `xml:"Mndt"`
	OrgnlMndt   OriginalMandate4Choice       `xml:"OrgnlMndt"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateAmendment5) Validate() error {
	return utils.Validate(&r)
}

type MandateAmendmentReason1 struct {
	Orgtr    *PartyIdentification43 `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      MandateReason1Choice   `xml:"Rsn"`
	AddtlInf []common.Max105Text    `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r MandateAmendmentReason1) Validate() error {
	return utils.Validate(&r)
}

type MandateAmendmentRequestV05 struct {
	Attr              []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr            GroupHeader47        `xml:"GrpHdr"`
	UndrlygAmdmntDtls []MandateAmendment5  `xml:"UndrlygAmdmntDtls" json:",omitempty"`
	SplmtryData       []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateAmendmentRequestV05) Validate() error {
	return utils.Validate(&r)
}

type MandateReason1Choice struct {
	Cd    ExternalMandateReason1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

func (r MandateReason1Choice) Validate() error {
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
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId common.Max35Text    `xml:"MsgNmId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalMessageInformation1) Validate() error {
	return utils.Validate(&r)
}

type MandateCancellation5 struct {
	OrgnlMsgInf *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty" json:",omitempty"`
	CxlRsn      PaymentCancellationReason1   `xml:"CxlRsn"`
	OrgnlMndt   OriginalMandate4Choice       `xml:"OrgnlMndt"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateCancellation5) Validate() error {
	return utils.Validate(&r)
}

type MandateCancellationRequestV05 struct {
	Attr           []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr         GroupHeader47          `xml:"GrpHdr"`
	UndrlygCxlDtls []MandateCancellation5 `xml:"UndrlygCxlDtls" json:",omitempty"`
	SplmtryData    []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateCancellationRequestV05) Validate() error {
	return utils.Validate(&r)
}

type PaymentCancellationReason1 struct {
	Orgtr    *PartyIdentification43 `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      MandateReason1Choice   `xml:"Rsn"`
	AddtlInf []common.Max105Text    `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r PaymentCancellationReason1) Validate() error {
	return utils.Validate(&r)
}

type AcceptanceResult6 struct {
	Accptd          bool                  `xml:"Accptd"`
	RjctRsn         *MandateReason1Choice `xml:"RjctRsn,omitempty" json:",omitempty"`
	AddtlRjctRsnInf []common.Max105Text   `xml:"AddtlRjctRsnInf,omitempty" json:",omitempty"`
}

func (r AcceptanceResult6) Validate() error {
	return utils.Validate(&r)
}

type Mandate11 struct {
	MndtId        *common.Max35Text                             `xml:"MndtId,omitempty" json:",omitempty"`
	MndtReqId     *common.Max35Text                             `xml:"MndtReqId,omitempty" json:",omitempty"`
	Authntcn      *MandateAuthentication1                       `xml:"Authntcn,omitempty" json:",omitempty"`
	Tp            *MandateTypeInformation2                      `xml:"Tp,omitempty" json:",omitempty"`
	Ocrncs        *MandateOccurrences4                          `xml:"Ocrncs,omitempty" json:",omitempty"`
	TrckgInd      bool                                          `xml:"TrckgInd"`
	FrstColltnAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"FrstColltnAmt,omitempty" json:",omitempty"`
	ColltnAmt     *ActiveOrHistoricCurrencyAndAmount            `xml:"ColltnAmt,omitempty" json:",omitempty"`
	MaxAmt        *ActiveOrHistoricCurrencyAndAmount            `xml:"MaxAmt,omitempty" json:",omitempty"`
	Adjstmnt      *MandateAdjustment1                           `xml:"Adjstmnt,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice                    `xml:"Rsn,omitempty" json:",omitempty"`
	CdtrSchmeId   PartyIdentification43                         `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	Cdtr          PartyIdentification43                         `xml:"Cdtr"`
	CdtrAcct      *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	CdtrAgt       *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	UltmtCdtr     *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	Dbtr          PartyIdentification43                         `xml:"Dbtr"`
	DbtrAcct      *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt       BranchAndFinancialInstitutionIdentification5  `xml:"DbtrAgt"`
	UltmtDbtr     *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	MndtRef       *common.Max35Text                             `xml:"MndtRef,omitempty" json:",omitempty"`
	RfrdDoc       []ReferredMandateDocument1                    `xml:"RfrdDoc,omitempty" json:",omitempty"`
}

func (r Mandate11) Validate() error {
	return utils.Validate(&r)
}

type MandateAcceptance5 struct {
	OrgnlMsgInf *OriginalMessageInformation1 `xml:"OrgnlMsgInf,omitempty" json:",omitempty"`
	AccptncRslt AcceptanceResult6            `xml:"AccptncRslt"`
	OrgnlMndt   *OriginalMandate5Choice      `xml:"OrgnlMndt,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateAcceptance5) Validate() error {
	return utils.Validate(&r)
}

type MandateAcceptanceReportV05 struct {
	Attr               []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr             GroupHeader47        `xml:"GrpHdr"`
	UndrlygAccptncDtls []MandateAcceptance5 `xml:"UndrlygAccptncDtls" json:",omitempty"`
	SplmtryData        []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r MandateAcceptanceReportV05) Validate() error {
	return utils.Validate(&r)
}

type OriginalMandate5Choice struct {
	OrgnlMndtId common.Max35Text `xml:"OrgnlMndtId"`
	OrgnlMndt   Mandate11        `xml:"OrgnlMndt"`
}

func (r OriginalMandate5Choice) Validate() error {
	return utils.Validate(&r)
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

func (r AmountType4Choice) Validate() error {
	return utils.Validate(&r)
}

type Cheque7 struct {
	ChqTp       *ChequeType2Code             `xml:"ChqTp,omitempty" json:",omitempty"`
	ChqNb       *common.Max35Text            `xml:"ChqNb,omitempty" json:",omitempty"`
	ChqFr       *NameAndAddress10            `xml:"ChqFr,omitempty" json:",omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty" json:",omitempty"`
	DlvrTo      *NameAndAddress10            `xml:"DlvrTo,omitempty" json:",omitempty"`
	InstrPrty   *Priority2Code               `xml:"InstrPrty,omitempty" json:",omitempty"`
	ChqMtrtyDt  *common.ISODate              `xml:"ChqMtrtyDt,omitempty" json:",omitempty"`
	FrmsCd      *common.Max35Text            `xml:"FrmsCd,omitempty" json:",omitempty"`
	MemoFld     []common.Max35Text           `xml:"MemoFld,omitempty" json:",omitempty"`
	RgnlClrZone *common.Max35Text            `xml:"RgnlClrZone,omitempty" json:",omitempty"`
	PrtLctn     *common.Max35Text            `xml:"PrtLctn,omitempty" json:",omitempty"`
	Sgntr       []common.Max70Text           `xml:"Sgntr,omitempty" json:",omitempty"`
}

func (r Cheque7) Validate() error {
	return utils.Validate(&r)
}

type CreditTransferTransaction22 struct {
	PmtId           PaymentIdentification1                        `xml:"PmtId"`
	PmtTpInf        *PaymentTypeInformation19                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	Amt             AmountType4Choice                             `xml:"Amt"`
	ChrgBr          ChargeBearerType1Code                         `xml:"ChrgBr"`
	ChqInstr        *Cheque7                                      `xml:"ChqInstr,omitempty" json:",omitempty"`
	UltmtDbtr       *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification5  `xml:"CdtrAgt"`
	Cdtr            PartyIdentification43                         `xml:"Cdtr"`
	CdtrAcct        *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr       *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent1                `xml:"InstrForCdtrAgt,omitempty" json:",omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty" json:",omitempty"`
	Tax             *TaxInformation3                              `xml:"Tax,omitempty" json:",omitempty"`
	RltdRmtInf      []RemittanceLocation4                         `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf          *RemittanceInformation11                      `xml:"RmtInf,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditTransferTransaction22) Validate() error {
	return utils.Validate(&r)
}

type CreditorPaymentActivationRequestV05 struct {
	Attr        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      GroupHeader45          `xml:"GrpHdr"`
	PmtInf      []PaymentInstruction19 `xml:"PmtInf" json:",omitempty"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorPaymentActivationRequestV05) Validate() error {
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
	Issr      common.Max35Text             `xml:"Issr,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType2) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriodDetails) Validate() error {
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

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode           `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       *common.Max4Text                  `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max140Text                `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DocumentAdjustment1) Validate() error {
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
	Id   []DocumentLineIdentification1 `xml:"Id"`
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

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

func (r EquivalentAmount2) Validate() error {
	return utils.Validate(&r)
}

type Garnishment1 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification43             `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification43             `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment1) Validate() error {
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

type GroupHeader45 struct {
	MsgId    common.Max35Text        `xml:"MsgId"`
	CreDtTm  common.ISODateTime      `xml:"CreDtTm"`
	NbOfTxs  common.Max15NumericText `xml:"NbOfTxs"`
	CtrlSum  float64                 `xml:"CtrlSum,omitempty" json:",omitempty"`
	InitgPty PartyIdentification43   `xml:"InitgPty"`
}

func (r GroupHeader45) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress10 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress6    `xml:"Adr"`
}

func (r NameAndAddress10) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification1 struct {
	InstrId    *common.Max35Text `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text  `xml:"EndToEndId"`
}

func (r PaymentIdentification1) Validate() error {
	return utils.Validate(&r)
}

type PaymentInstruction19 struct {
	PmtInfId    *common.Max35Text                            `xml:"PmtInfId,omitempty" json:",omitempty"`
	PmtMtd      PaymentMethod7Code                           `xml:"PmtMtd"`
	PmtTpInf    *PaymentTypeInformation19                    `xml:"PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt *common.ISODate                              `xml:"ReqdExctnDt"`
	Dbtr        PartyIdentification43                        `xml:"Dbtr"`
	DbtrAcct    *CashAccount24                               `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt     BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`
	UltmtDbtr   *PartyIdentification43                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	ChrgBr      *ChargeBearerType1Code                       `xml:"ChrgBr,omitempty" json:",omitempty"`
	CdtTrfTx    []CreditTransferTransaction22                `xml:"CdtTrfTx"`
}

func (r PaymentInstruction19) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation19 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation19) Validate() error {
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

type RegulatoryAuthority2 struct {
	Nm   *common.Max140Text  `xml:"Nm,omitempty"`
	Ctry *common.CountryCode `xml:"Ctry,omitempty"`
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

type RemittanceInformation11 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation13 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation11) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocation4 struct {
	RmtId       *common.Max35Text            `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationDetails1 `xml:"RmtLctnDtls,omitempty" json:",omitempty"`
}

func (r RemittanceLocation4) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocationDetails1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr *common.Max2048Text           `xml:"ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    *NameAndAddress10             `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r RemittanceLocationDetails1) Validate() error {
	return utils.Validate(&r)
}

type StructuredRegulatoryReporting3 struct {
	Tp   *common.Max35Text                 `xml:"Tp,omitempty" json:",omitempty"`
	Dt   *common.ISODate                   `xml:"Dt,omitempty" json:",omitempty"`
	Ctry *common.CountryCode               `xml:"Ctry,omitempty" json:",omitempty"`
	Cd   *common.Max10Text                 `xml:"Cd,omitempty" json:",omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Inf  []common.Max35Text                `xml:"Inf,omitempty" json:",omitempty"`
}

func (r StructuredRegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation13 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification43         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification43         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation4               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment1                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation13) Validate() error {
	return utils.Validate(&r)
}

type TaxAmount1 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails1                `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount1) Validate() error {
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

type TaxInformation3 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	AdmstnZn        *common.Max35Text                  `xml:"AdmstnZn,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord1                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation3) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation4 struct {
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
	Rcrd            []TaxRecord1                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation4) Validate() error {
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

type TaxPeriod1 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriodDetails    `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod1) Validate() error {
	return utils.Validate(&r)
}

type TaxRecord1 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod1        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount1        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord1) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails1) Validate() error {
	return utils.Validate(&r)
}

type InstructionForCreditorAgent1 struct {
	Cd       *Instruction3Code  `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForCreditorAgent1) Validate() error {
	return utils.Validate(&r)
}

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"Cd"`
	Prtry common.Max35Text    `xml:"Prtry"`
}

func (r ChequeDeliveryMethod1Choice) Validate() error {
	return utils.Validate(&r)
}

type Charges2 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	Agt BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (r Charges2) Validate() error {
	return utils.Validate(&r)
}

type CreditorPaymentActivationRequestStatusReportV05 struct {
	Attr              []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr            GroupHeader46                  `xml:"GrpHdr"`
	OrgnlGrpInfAndSts OriginalGroupInformation25     `xml:"OrgnlGrpInfAndSts"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction19 `xml:"OrgnlPmtInfAndSts,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1           `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorPaymentActivationRequestStatusReportV05) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader46 struct {
	MsgId    common.Max35Text                              `xml:"MsgId"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm"`
	InitgPty PartyIdentification43                         `xml:"InitgPty"`
	DbtrAgt  *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt  *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
}

func (r GroupHeader46) Validate() error {
	return utils.Validate(&r)
}

type NumberOfTransactionsPerStatus3 struct {
	DtldNbOfTxs common.Max15NumericText          `xml:"DtldNbOfTxs"`
	DtldSts     TransactionIndividualStatus3Code `xml:"DtldSts"`
	DtldCtrlSum float64                          `xml:"DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfTransactionsPerStatus3) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupInformation25 struct {
	OrgnlMsgId    common.Max35Text                 `xml:"OrgnlMsgId"`
	OrgnlMsgNmId  common.Max35Text                 `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm  *common.ISODateTime              `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNbOfTxs  *common.Max15NumericText         `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum  float64                          `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	GrpSts        *TransactionGroupStatus3Code     `xml:"GrpSts,omitempty" json:",omitempty"`
	StsRsnInf     []StatusReasonInformation9       `xml:"StsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty" json:",omitempty"`
}

func (r OriginalGroupInformation25) Validate() error {
	return utils.Validate(&r)
}

type OriginalPaymentInstruction19 struct {
	OrgnlPmtInfId *common.Max35Text                `xml:"OrgnlPmtInfId"`
	OrgnlNbOfTxs  *common.Max15NumericText         `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum  float64                          `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	PmtInfSts     *TransactionGroupStatus3Code     `xml:"PmtInfSts,omitempty" json:",omitempty"`
	StsRsnInf     []StatusReasonInformation9       `xml:"StsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerSts []NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty" json:",omitempty"`
	TxInfAndSts   []PaymentTransaction69           `xml:"TxInfAndSts,omitempty" json:",omitempty"`
}

func (r OriginalPaymentInstruction19) Validate() error {
	return utils.Validate(&r)
}

type OriginalTransactionReference23 struct {
	Amt         *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	ReqdExctnDt *common.ISODate                               `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	PmtTpInf    *PaymentTypeInformation19                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd      *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	RmtInf      *RemittanceInformation11                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr   *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr        *PartyIdentification43                        `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct    *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	CdtrAgt     BranchAndFinancialInstitutionIdentification5  `xml:"CdtrAgt"`
	Cdtr        PartyIdentification43                         `xml:"Cdtr"`
	CdtrAcct    *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr   *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference23) Validate() error {
	return utils.Validate(&r)
}

type PaymentTransaction69 struct {
	StsId           *common.Max35Text                 `xml:"StsId,omitempty" json:",omitempty"`
	OrgnlInstrId    *common.Max35Text                 `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	TxSts           *TransactionIndividualStatus3Code `xml:"TxSts,omitempty" json:",omitempty"`
	StsRsnInf       []StatusReasonInformation9        `xml:"StsRsnInf,omitempty" json:",omitempty"`
	ChrgsInf        []Charges2                        `xml:"ChrgsInf,omitempty" json:",omitempty"`
	AccptncDtTm     *common.ISODateTime               `xml:"AccptncDtTm,omitempty" json:",omitempty"`
	AcctSvcrRef     *common.Max35Text                 `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	ClrSysRef       *common.Max35Text                 `xml:"ClrSysRef,omitempty" json:",omitempty"`
	OrgnlTxRef      *OriginalTransactionReference23   `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1              `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentTransaction69) Validate() error {
	return utils.Validate(&r)
}

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r StatusReason6Choice) Validate() error {
	return utils.Validate(&r)
}

type StatusReasonInformation9 struct {
	Orgtr    *PartyIdentification43 `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *StatusReason6Choice   `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text    `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r StatusReasonInformation9) Validate() error {
	return utils.Validate(&r)
}
