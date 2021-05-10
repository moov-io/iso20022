// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type InformationRequestOpeningV01 struct {
	Attr        []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	InvstgtnId  common.Max35Text           `xml:"InvstgtnId"`
	LglMndtBsis LegalMandate1              `xml:"LglMndtBsis"`
	CnfdtltySts bool                       `xml:"CnfdtltySts"`
	DueDt       *DueDate1                  `xml:"DueDt,omitempty" json:",omitempty"`
	InvstgtnPrd DateOrDateTimePeriodChoice `xml:"InvstgtnPrd"`
	SchCrit     SearchCriteria1Choice      `xml:"SchCrit"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r InformationRequestOpeningV01) Validate() error {
	return utils.Validate(&r)
}

type InformationRequestResponseV01 struct {
	Attr        []utils.Attr          `xml:",any,attr,omitempty" json:",omitempty"`
	RspnId      common.Max35Text      `xml:"RspnId"`
	InvstgtnId  common.Max35Text      `xml:"InvstgtnId"`
	RspnSts     StatusResponse1Code   `xml:"RspnSts"`
	SchCrit     SearchCriteria1Choice `xml:"SchCrit"`
	RtrInd      []ReturnIndicator1    `xml:"RtrInd" json:",omitempty"`
	SplmtryData []SupplementaryData1  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r InformationRequestResponseV01) Validate() error {
	return utils.Validate(&r)
}

type InformationRequestStatusChangeNotificationV01 struct {
	Attr        []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	OrgnlBizQry common.Max35Text     `xml:"OrgnlBizQry"`
	CnfdtltySts bool                 `xml:"CnfdtltySts"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r InformationRequestStatusChangeNotificationV01) Validate() error {
	return utils.Validate(&r)
}

type AccountAndParties1 struct {
	Id            CashAccount25              `xml:"Id"`
	InvstgtdPties InvestigatedParties1Choice `xml:"InvstgtdPties"`
	AuthrtyReqTp  []AuthorityRequestType1    `xml:"AuthrtyReqTp" json:",omitempty"`
}

func (r AccountAndParties1) Validate() error {
	return utils.Validate(&r)
}

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

type AuthorityInvestigation2 struct {
	Tp                 AuthorityRequestType1       `xml:"Tp"`
	InvstgtdRoles      InvestigatedParties1Choice  `xml:"InvstgtdRoles"`
	AddtlInvstgtdPties *InvestigatedParties1Choice `xml:"AddtlInvstgtdPties,omitempty" json:",omitempty"`
	AddtlInf           *common.Max500Text          `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r AuthorityInvestigation2) Validate() error {
	return utils.Validate(&r)
}

type AuthorityRequestType1 struct {
	MsgNmId common.Max35Text   `xml:"MsgNmId"`
	MsgNm   *common.Max140Text `xml:"MsgNm,omitempty" json:",omitempty"`
}

func (r AuthorityRequestType1) Validate() error {
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

type CashAccount25 struct {
	Id   AccountIdentification4Choice                  `xml:"Id"`
	Tp   *CashAccountType2Choice                       `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode          `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                             `xml:"Nm,omitempty" json:",omitempty"`
	Ownr *PartyIdentification43                        `xml:"Ownr,omitempty" json:",omitempty"`
	Svcr *BranchAndFinancialInstitutionIdentification5 `xml:"Svcr,omitempty" json:",omitempty"`
}

func (r CashAccount25) Validate() error {
	return utils.Validate(&r)
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CashAccountType2Choice) Validate() error {
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

type CustomerIdentification1 struct {
	Pty        PartyIdentification43     `xml:"Pty"`
	AuthrtyReq []AuthorityInvestigation2 `xml:"AuthrtyReq" json:",omitempty"`
}

func (r CustomerIdentification1) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth *common.Max35Text  `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
	return utils.Validate(&r)
}

type DateOrDateTimePeriodChoice struct {
	Dt   DatePeriodDetails     `xml:"Dt"`
	DtTm DateTimePeriodDetails `xml:"DtTm"`
}

func (r DateOrDateTimePeriodChoice) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriodDetails) Validate() error {
	return utils.Validate(&r)
}

type DateTimePeriodDetails struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
}

func (r DateTimePeriodDetails) Validate() error {
	return utils.Validate(&r)
}

type DueDate1 struct {
	DueDt    *common.ISODate    `xml:"DueDt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DueDate1) Validate() error {
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

type InvestigatedParties1Choice struct {
	Cd    InvestigatedParties1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r InvestigatedParties1Choice) Validate() error {
	return utils.Validate(&r)
}

type LegalMandate1 struct {
	Prgrph common.Max35Text   `xml:"Prgrph"`
	Dsclmr *common.Max350Text `xml:"Dsclmr,omitempty" json:",omitempty"`
}

func (r LegalMandate1) Validate() error {
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

type PaymentInstrumentType1 struct {
	CardNb       Min8Max28NumericText    `xml:"CardNb"`
	AuthrtyReqTp []AuthorityRequestType1 `xml:"AuthrtyReqTp" json:",omitempty"`
	AddtlInf     *common.Max500Text      `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r PaymentInstrumentType1) Validate() error {
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

type RequestType1 struct {
	Nb       common.Max35Text              `xml:"Nb"`
	Tp       []TransactionRequestType1Code `xml:"Tp" json:",omitempty"`
	AddtlInf *common.Max500Text            `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r RequestType1) Validate() error {
	return utils.Validate(&r)
}

type SearchCriteria1Choice struct {
	Acct      AccountAndParties1      `xml:"Acct"`
	CstmrId   CustomerIdentification1 `xml:"CstmrId"`
	PmtInstrm PaymentInstrumentType1  `xml:"PmtInstrm"`
	OrgnlTxNb []RequestType1          `xml:"OrgnlTxNb" json:",omitempty"`
}

func (r SearchCriteria1Choice) Validate() error {
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

type InvestigationResult1Choice struct {
	Rslt        SupplementaryDataEnvelope1 `xml:"Rslt"`
	InvstgtnSts InvestigationStatus1Code   `xml:"InvstgtnSts"`
}

func (r InvestigationResult1Choice) Validate() error {
	return utils.Validate(&r)
}

type ReturnIndicator1 struct {
	RspnPrd      *DateOrDateTimePeriodChoice `xml:"RspnPrd,omitempty" json:",omitempty"`
	AuthrtyReqTp *AuthorityRequestType1      `xml:"AuthrtyReqTp"`
	InvstgtnRslt InvestigationResult1Choice  `xml:"InvstgtnRslt"`
	AddtlInf     *common.Max500Text          `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r ReturnIndicator1) Validate() error {
	return utils.Validate(&r)
}
