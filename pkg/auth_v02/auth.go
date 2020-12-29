// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Othr"`
}

func (r AccountIdentification4Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
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

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r AddressType3Choice) Validate() error {
	return utils.Validate(&r)
}

type BenchmarkCurveName4Choice struct {
	ISIN ISINOct2015Identifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ISIN"`
	Indx BenchmarkCurveName2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Indx"`
	Nm   common.Max25Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Nm"`
}

func (r BenchmarkCurveName4Choice) Validate() error {
	return utils.Validate(&r)
}

type BinaryFile1 struct {
	MIMETp         common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MIMETp,omitempty"`
	NcodgTp        common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 NcodgTp,omitempty"`
	CharSet        common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CharSet,omitempty"`
	InclBinryObjct Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 InclBinryObjct,omitempty"`
}

func (r BinaryFile1) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 BrnchId,omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification6) Validate() error {
	return utils.Validate(&r)
}

type BranchData3 struct {
	Id      common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id,omitempty"`
	LEI     LEIIdentifier     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LEI,omitempty"`
	Nm      common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Nm,omitempty"`
	PstlAdr PostalAddress24   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PstlAdr,omitempty"`
}

func (r BranchData3) Validate() error {
	return utils.Validate(&r)
}

type CashCollateral5 struct {
	CollId    common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CollId,omitempty"`
	CshAcctId AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CshAcctId,omitempty"`
	AsstNb    common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AsstNb,omitempty"`
	DpstAmt   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DpstAmt,omitempty"`
	DpstTp    DepositType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DpstTp,omitempty"`
	MtrtyDt   common.ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MtrtyDt,omitempty"`
	ValDt     common.ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ValDt,omitempty"`
	XchgRate  float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XchgRate,omitempty"`
	CollVal   ActiveCurrencyAndAmount      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CollVal"`
	Hrcut     float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Hrcut,omitempty"`
}

func (r CashCollateral5) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r ClearingSystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MmbId"`
}

func (r ClearingSystemMemberIdentification2) Validate() error {
	return utils.Validate(&r)
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Nm,omitempty"`
	PhneNb    common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PrefrdMtd,omitempty"`
}

func (r Contact4) Validate() error {
	return utils.Validate(&r)
}

type ContractBalance1 struct {
	Tp        ContractBalanceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Tp"`
	Amt       ActiveCurrencyAndAmount    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt"`
	CdtDbtInd CreditDebit3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CdtDbtInd"`
}

func (r ContractBalance1) Validate() error {
	return utils.Validate(&r)
}

type ContractBalanceType1Choice struct {
	Cd    ExternalContractBalanceType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r ContractBalanceType1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractCollateral1 struct {
	TtlAmt   ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TtlAmt"`
	CollDesc []CashCollateral5       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CollDesc,omitempty"`
	AddtlInf common.Max1025Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AddtlInf,omitempty"`
}

func (r ContractCollateral1) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistration3 struct {
	CtrctRegnId   common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctRegnId"`
	RptgPty       TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 RptgPty"`
	RegnAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 RegnAgt"`
	CtrctRegnOpng []ContractRegistration4                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctRegnOpng"`
	SplmtryData   []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistration3) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistration4 struct {
	CtrctRegnOpngId common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctRegnOpngId"`
	Prty            Priority2Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prty"`
	Ctrct           UnderlyingContract2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Ctrct"`
	CtrctBal        []ContractBalance1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctBal,omitempty"`
	PmtSchdlTp      PaymentScheduleType1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PmtSchdlTp,omitempty"`
	PrvsRegnId      []DocumentIdentification22    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PrvsRegnId,omitempty"`
	AddtlInf        common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AddtlInf,omitempty"`
	Attchmnt        []DocumentGeneralInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Attchmnt,omitempty"`
	SplmtryData     []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistration4) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationRequestV02 struct {
	GrpHdr      CurrencyControlHeader4  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 GrpHdr"`
	CtrctRegn   []ContractRegistration3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctRegn"`
	SplmtryData []SupplementaryData1    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationRequestV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlHeader4 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CreDtTm"`
	NbOfItms Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 NbOfItms"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 InitgPty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 FwdgAgt,omitempty"`
}

func (r CurrencyControlHeader4) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 BirthDt"`
	PrvcOfBirth common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CityOfBirth"`
	CtryOfBirth CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
	return utils.Validate(&r)
}

type DocumentGeneralInformation3 struct {
	DocTp           ExternalDocumentType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DocTp"`
	DocNb           common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DocNb"`
	SndrRcvrSeqId   common.Max140Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SndrRcvrSeqId,omitempty"`
	IsseDt          common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IsseDt,omitempty"`
	URL             common.Max256Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 URL,omitempty"`
	LkFileHash      SignatureEnvelopeReference `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LkFileHash,omitempty"`
	AttchdBinryFile BinaryFile1                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AttchdBinryFile"`
}

func (r DocumentGeneralInformation3) Validate() error {
	return utils.Validate(&r)
}

type DocumentIdentification22 struct {
	Id       common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id"`
	DtOfIsse common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DtOfIsse,omitempty"`
}

func (r DocumentIdentification22) Validate() error {
	return utils.Validate(&r)
}

type ExchangeRate1 struct {
	UnitCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 UnitCcy,omitempty"`
	XchgRate float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XchgRate,omitempty"`
	RateTp   ExchangeRateType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 RateTp,omitempty"`
	CtrctId  common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctId,omitempty"`
}

func (r ExchangeRate1) Validate() error {
	return utils.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LEI,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Othr,omitempty"`
}

func (r FinancialInstitutionIdentification18) Validate() error {
	return utils.Validate(&r)
}

type FloatingInterestRate4 struct {
	RefRate    BenchmarkCurveName4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 RefRate"`
	Term       InterestRateContractTerm1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Term"`
	BsisPtSprd float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 BsisPtSprd"`
}

func (r FloatingInterestRate4) Validate() error {
	return utils.Validate(&r)
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Issr,omitempty"`
}

func (r GenericAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Issr,omitempty"`
}

func (r GenericFinancialIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id"`
	Issr    common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Issr"`
	SchmeNm common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SchmeNm,omitempty"`
}

func (r GenericIdentification30) Validate() error {
	return utils.Validate(&r)
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Issr,omitempty"`
}

func (r GenericOrganisationIdentification1) Validate() error {
	return utils.Validate(&r)
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Issr,omitempty"`
}

func (r GenericPersonIdentification1) Validate() error {
	return utils.Validate(&r)
}

type InterestPaymentDateRange1 struct {
	IntrstSchdlId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IntrstSchdlId,omitempty"`
	XpctdDt       common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XpctdDt,omitempty"`
	DueDt         common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DueDt,omitempty"`
}

func (r InterestPaymentDateRange1) Validate() error {
	return utils.Validate(&r)
}

type InterestPaymentDateRange2 struct {
	IntrstSchdlId common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IntrstSchdlId,omitempty"`
	Amt           ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt"`
	DueDt         common.ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DueDt"`
	AddtlInf      common.Max1025Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AddtlInf,omitempty"`
}

func (r InterestPaymentDateRange2) Validate() error {
	return utils.Validate(&r)
}

type InterestPaymentSchedule1Choice struct {
	DtRg     InterestPaymentDateRange1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DtRg"`
	SubSchdl []InterestPaymentDateRange2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SubSchdl"`
}

func (r InterestPaymentSchedule1Choice) Validate() error {
	return utils.Validate(&r)
}

type InterestRate2Choice struct {
	Fxd  float64               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Fxd"`
	Fltg FloatingInterestRate4 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Fltg"`
}

func (r InterestRate2Choice) Validate() error {
	return utils.Validate(&r)
}

type InterestRateContractTerm1 struct {
	Unit RateBasis1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Unit"`
	Val  float64        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Val"`
}

func (r InterestRateContractTerm1) Validate() error {
	return utils.Validate(&r)
}

type LegalOrganisation2 struct {
	Id           common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id,omitempty"`
	Nm           common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Nm,omitempty"`
	EstblishmtDt common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 EstblishmtDt,omitempty"`
	RegnDt       common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 RegnDt,omitempty"`
}

func (r LegalOrganisation2) Validate() error {
	return utils.Validate(&r)
}

type LoanContract2 struct {
	CtrctDocId  DocumentIdentification22       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctDocId"`
	Buyr        []TradeParty5                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Buyr"`
	Sellr       []TradeParty5                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Sellr"`
	Amt         ActiveCurrencyAndAmount        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt"`
	MtrtyDt     common.ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MtrtyDt"`
	PrlngtnFlg  bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PrlngtnFlg"`
	StartDt     common.ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 StartDt"`
	SttlmCcy    ActiveCurrencyCode             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SttlmCcy"`
	SpclConds   SpecialCondition1              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SpclConds,omitempty"`
	DrtnCd      Exact1NumericText              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DrtnCd"`
	IntrstRate  InterestRate2Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IntrstRate"`
	Trch        []LoanContractTranche1         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Trch,omitempty"`
	PmtSchdl    PaymentSchedule1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PmtSchdl,omitempty"`
	IntrstSchdl InterestPaymentSchedule1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IntrstSchdl"`
	IntraCpnyLn bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IntraCpnyLn"`
	Coll        ContractCollateral1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Coll,omitempty"`
	SndctdLn    []SyndicatedLoan2              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SndctdLn,omitempty"`
	Attchmnt    []DocumentGeneralInformation3  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Attchmnt,omitempty"`
}

func (r LoanContract2) Validate() error {
	return utils.Validate(&r)
}

type LoanContractTranche1 struct {
	TrchNb      float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TrchNb"`
	XpctdDt     common.ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XpctdDt"`
	Amt         ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt"`
	DueDt       common.ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DueDt,omitempty"`
	DrtnCd      Exact1NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DrtnCd,omitempty"`
	LastTrchInd bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LastTrchInd,omitempty"`
}

func (r LoanContractTranche1) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Othr,omitempty"`
}

func (r OrganisationIdentification29) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id,omitempty"`
}

func (r OtherContact1) Validate() error {
	return utils.Validate(&r)
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PrvtId"`
}

func (r Party38Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification135 struct {
	Nm        common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress24   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PstlAdr,omitempty"`
	Id        Party38Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Id,omitempty"`
	CtryOfRes CountryCode       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtryOfRes,omitempty"`
	CtctDtls  Contact4          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtctDtls,omitempty"`
}

func (r PartyIdentification135) Validate() error {
	return utils.Validate(&r)
}

type PaymentDateRange1 struct {
	PmtSchdlId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PmtSchdlId,omitempty"`
	XpctdDt    common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XpctdDt,omitempty"`
	DueDt      common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DueDt,omitempty"`
}

func (r PaymentDateRange1) Validate() error {
	return utils.Validate(&r)
}

type PaymentDateRange2 struct {
	PmtSchdlId common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PmtSchdlId,omitempty"`
	Amt        ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt"`
	XpctdDt    common.ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XpctdDt,omitempty"`
	DueDt      common.ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DueDt"`
}

func (r PaymentDateRange2) Validate() error {
	return utils.Validate(&r)
}

type PaymentSchedule1Choice struct {
	DtRg     PaymentDateRange1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DtRg"`
	SubSchdl []PaymentDateRange2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SubSchdl"`
}

func (r PaymentSchedule1Choice) Validate() error {
	return utils.Validate(&r)
}

type PaymentScheduleType1Choice struct {
	Cd    PaymentScheduleType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r PaymentScheduleType1Choice) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Othr,omitempty"`
}

func (r PersonIdentification13) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 AdrLine,omitempty"`
}

func (r PostalAddress24) Validate() error {
	return utils.Validate(&r)
}

type ShipmentDateRange1 struct {
	EarlstShipmntDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 EarlstShipmntDt,omitempty"`
	LatstShipmntDt  common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LatstShipmntDt,omitempty"`
}

func (r ShipmentDateRange1) Validate() error {
	return utils.Validate(&r)
}

type ShipmentDateRange2 struct {
	SubQtyVal       float64        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SubQtyVal"`
	EarlstShipmntDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 EarlstShipmntDt,omitempty"`
	LatstShipmntDt  common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LatstShipmntDt,omitempty"`
}

func (r ShipmentDateRange2) Validate() error {
	return utils.Validate(&r)
}

type ShipmentSchedule2Choice struct {
	ShipmntDtRg     ShipmentDateRange1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ShipmntDtRg"`
	ShipmntSubSchdl []ShipmentDateRange2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ShipmntSubSchdl"`
}

func (r ShipmentSchedule2Choice) Validate() error {
	return utils.Validate(&r)
}

type SignatureEnvelopeReference struct {
	Item string `xml:",any"`
}

func (r SignatureEnvelopeReference) Validate() error {
	return utils.Validate(&r)
}

type SpecialCondition1 struct {
	IncmgAmt           ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IncmgAmt"`
	OutgngAmt          ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 OutgngAmt"`
	IncmgAmtToOthrAcct ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 IncmgAmtToOthrAcct,omitempty"`
	PmtFrOthrAcct      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PmtFrOthrAcct,omitempty"`
}

func (r SpecialCondition1) Validate() error {
	return utils.Validate(&r)
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Envlp"`
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

type SyndicatedLoan2 struct {
	Brrwr       TradeParty5             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Brrwr"`
	Lndr        TradeParty5             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Lndr,omitempty"`
	Amt         ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt,omitempty"`
	Shr         float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Shr,omitempty"`
	XchgRateInf ExchangeRate1           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XchgRateInf,omitempty"`
}

func (r SyndicatedLoan2) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptionReasonFormat1Choice struct {
	Ustrd common.Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Ustrd"`
	Strd  TaxExemptReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Strd"`
}

func (r TaxExemptionReasonFormat1Choice) Validate() error {
	return utils.Validate(&r)
}

type TaxParty4 struct {
	TaxId       common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TaxId,omitempty"`
	TaxTp       common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TaxTp,omitempty"`
	RegnId      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 RegnId,omitempty"`
	TaxXmptnRsn []TaxExemptionReasonFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TaxXmptnRsn,omitempty"`
}

func (r TaxParty4) Validate() error {
	return utils.Validate(&r)
}

type TradeContract2 struct {
	CtrctDocId   DocumentIdentification22      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 CtrctDocId,omitempty"`
	Amt          ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Amt"`
	Buyr         []TradeParty5                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Buyr"`
	Sellr        []TradeParty5                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Sellr"`
	MtrtyDt      common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 MtrtyDt"`
	PrlngtnFlg   bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PrlngtnFlg"`
	StartDt      common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 StartDt"`
	SttlmCcy     ActiveCurrencyCode            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 SttlmCcy"`
	XchgRateInf  ExchangeRate1                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 XchgRateInf,omitempty"`
	PmtSchdl     InterestPaymentDateRange1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PmtSchdl,omitempty"`
	ShipmntSchdl ShipmentSchedule2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 ShipmntSchdl,omitempty"`
	Attchmnt     []DocumentGeneralInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Attchmnt,omitempty"`
}

func (r TradeContract2) Validate() error {
	return utils.Validate(&r)
}

type TradeParty5 struct {
	PtyId  PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 PtyId"`
	LglOrg LegalOrganisation2     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 LglOrg,omitempty"`
	TaxPty []TaxParty4            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 TaxPty,omitempty"`
}

func (r TradeParty5) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingContract2Choice struct {
	Ln   LoanContract2  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Ln"`
	Trad TradeContract2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.02 Trad"`
}

func (r UnderlyingContract2Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractClosureReason1Choice struct {
	Cd    ExternalContractClosureReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Prtry"`
}

func (r ContractClosureReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationConfirmationV02 struct {
	GrpHdr      CurrencyControlHeader6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 GrpHdr"`
	RegdCtrct   []RegisteredContract7  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RegdCtrct"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationConfirmationV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlHeader6 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 CreDtTm"`
	NbOfItms Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 NbOfItms"`
	RcvgPty  PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RcvgPty"`
	RegnAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RegnAgt"`
}

func (r CurrencyControlHeader6) Validate() error {
	return utils.Validate(&r)
}

type DocumentIdentification28 struct {
	Id       common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Id,omitempty"`
	DtOfIsse common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 DtOfIsse"`
}

func (r DocumentIdentification28) Validate() error {
	return utils.Validate(&r)
}

type DocumentIdentification29 struct {
	Id       common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Id"`
	DtOfIsse common.ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 DtOfIsse"`
}

func (r DocumentIdentification29) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract7 struct {
	OrgnlCtrctRegnReq common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 OrgnlCtrctRegnReq,omitempty"`
	RptgPty           TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RptgPty"`
	RegnAgt           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RegnAgt"`
	IssrFI            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 IssrFI"`
	Ctrct             UnderlyingContract2Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Ctrct"`
	CtrctBal          []ContractBalance1                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 CtrctBal,omitempty"`
	PmtSchdlTp        PaymentScheduleType1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 PmtSchdlTp,omitempty"`
	RegdCtrctId       DocumentIdentification29                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RegdCtrctId"`
	PrvsRegdCtrctId   DocumentIdentification22                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 PrvsRegdCtrctId,omitempty"`
	RegdCtrctJrnl     []RegisteredContractJournal2                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RegdCtrctJrnl,omitempty"`
	Amdmnt            []RegisteredContractAmendment1               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Amdmnt,omitempty"`
	Submissn          RegisteredContractCommunication1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Submissn"`
	Dlvry             RegisteredContractCommunication1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Dlvry"`
	LnPrncplAmt       ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 LnPrncplAmt,omitempty"`
	EstmtdDtInd       bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 EstmtdDtInd"`
	IntrCpnyLn        bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 IntrCpnyLn"`
	AddtlInf          common.Max1025Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 AddtlInf,omitempty"`
	SplmtryData       []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 SplmtryData,omitempty"`
}

func (r RegisteredContract7) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContractAmendment1 struct {
	AmdmntDt  common.ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 AmdmntDt"`
	Doc       DocumentIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Doc"`
	StartDt   common.ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 StartDt,omitempty"`
	AmdmntRsn common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 AmdmntRsn,omitempty"`
	AddtlInf  common.Max1025Text       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 AddtlInf,omitempty"`
}

func (r RegisteredContractAmendment1) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContractCommunication1 struct {
	Mtd CommunicationMethod4Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Mtd"`
	Dt  common.ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 Dt"`
}

func (r RegisteredContractCommunication1) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContractJournal2 struct {
	RegnAgt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 RegnAgt"`
	UnqId   DocumentIdentification28                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 UnqId,omitempty"`
	ClsrDt  common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 ClsrDt"`
	ClsrRsn ContractClosureReason1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.02 ClsrRsn"`
}

func (r RegisteredContractJournal2) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationClosureRequestV02 struct {
	GrpHdr        CurrencyControlHeader4 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 GrpHdr"`
	RegdCtrctClsr []RegisteredContract6  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 RegdCtrctClsr"`
	SplmtryData   []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationClosureRequestV02) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract6 struct {
	RegdCtrctClsrId common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 RegdCtrctClsrId"`
	RptgPty         TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 RptgPty"`
	RegnAgt         BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 RegnAgt"`
	OrgnlRegdCtrct  DocumentIdentification29                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 OrgnlRegdCtrct"`
	Prty            Priority2Code                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 Prty"`
	ClsrRsn         ContractClosureReason1Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 ClsrRsn"`
	SplmtryData     []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.020.001.02 SplmtryData,omitempty"`
}

func (r RegisteredContract6) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationAmendmentRequestV02 struct {
	GrpHdr          CurrencyControlHeader4 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 GrpHdr"`
	CtrctRegnAmdmnt []RegisteredContract9  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 CtrctRegnAmdmnt"`
	SplmtryData     []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationAmendmentRequestV02) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract10 struct {
	RegdCtrctAmdmntId common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 RegdCtrctAmdmntId"`
	OrgnlRegdCtrctId  common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 OrgnlRegdCtrctId"`
	Prty              Priority2Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 Prty"`
	Ctrct             UnderlyingContract2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 Ctrct"`
	CtrctBal          []ContractBalance1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 CtrctBal,omitempty"`
	PmtSchdlTp        PaymentScheduleType1Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 PmtSchdlTp,omitempty"`
	AddtlInf          common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 AddtlInf,omitempty"`
	Attchmnt          []DocumentGeneralInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 Attchmnt,omitempty"`
	SplmtryData       []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 SplmtryData,omitempty"`
}

func (r RegisteredContract10) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract9 struct {
	CtrctRegnAmdmntId common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 CtrctRegnAmdmntId"`
	RptgPty           TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 RptgPty"`
	RegnAgt           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 RegnAgt"`
	RegdCtrctAmdmnt   []RegisteredContract10                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 RegdCtrctAmdmnt"`
	SplmtryData       []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.02 SplmtryData,omitempty"`
}

func (r RegisteredContract9) Validate() error {
	return utils.Validate(&r)
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Ccy,omitempty"`
	Nm   common.Max70Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Prxy,omitempty"`
}

func (r CashAccount38) Validate() error {
	return utils.Validate(&r)
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Prtry"`
}

func (r CashAccountType2Choice) Validate() error {
	return utils.Validate(&r)
}

type CertificateIdentification1 struct {
	MsgId       common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 MsgId,omitempty"`
	AcctSvcrRef common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AcctSvcrRef,omitempty"`
	PmtInfId    common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 PmtInfId,omitempty"`
	InstrId     common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 InstrId,omitempty"`
	EndToEndId  common.Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 EndToEndId,omitempty"`
	Prtry       ProprietaryReference1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Prtry,omitempty"`
}

func (r CertificateIdentification1) Validate() error {
	return utils.Validate(&r)
}

type CertificateReference1 struct {
	Id CertificateIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Id"`
	Dt common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Dt"`
}

func (r CertificateReference1) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationReference1Choice struct {
	RegdCtrctId common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RegdCtrctId"`
	Ctrct       DocumentIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Ctrct"`
}

func (r ContractRegistrationReference1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatement2 struct {
	StmtId             common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 StmtId,omitempty"`
	RptgPty            TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RptgPty"`
	RegnAgt            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RegnAgt"`
	RptgPrd            ReportingPeriod4                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RptgPrd"`
	RegdCtrct          RegisteredContract8                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RegdCtrct"`
	TxJrnl             []TransactionCertificate3                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TxJrnl,omitempty"`
	SpprtgDocJrnl      []SupportingDocument2                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SpprtgDocJrnl,omitempty"`
	AddtlSpprtgDocJrnl []SupportingDocument2                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AddtlSpprtgDocJrnl,omitempty"`
	RgltryRuleVldtn    []GenericValidationRuleIdentification1       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RgltryRuleVldtn,omitempty"`
	TtlCtrctTrnvrSum   ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TtlCtrctTrnvrSum"`
	SplmtryData        []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationStatement2) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementV02 struct {
	GrpHdr      CurrencyControlHeader6           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 GrpHdr"`
	Stmt        []ContractRegistrationStatement2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Stmt"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationStatementV02) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod3 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 ToDt,omitempty"`
}

func (r DatePeriod3) Validate() error {
	return utils.Validate(&r)
}

type DocumentAmendment1 struct {
	CrrctnId   float64          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CrrctnId"`
	OrgnlDocId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 OrgnlDocId,omitempty"`
}

func (r DocumentAmendment1) Validate() error {
	return utils.Validate(&r)
}

type GenericValidationRuleIdentification1 struct {
	Id      common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Id"`
	Desc    common.Max350Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Desc,omitempty"`
	SchmeNm ValidationRuleSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SchmeNm,omitempty"`
	Issr    common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Issr,omitempty"`
}

func (r GenericValidationRuleIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryReference1 struct {
	Tp  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Tp"`
	Ref common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Ref"`
}

func (r ProprietaryReference1) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Tp,omitempty"`
	Id common.Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Id"`
}

func (r ProxyAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Prtry"`
}

func (r ProxyAccountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract8 struct {
	OrgnlCtrctRegnReq common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 OrgnlCtrctRegnReq,omitempty"`
	IssrFI            BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 IssrFI"`
	Ctrct             UnderlyingContract2Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Ctrct"`
	CtrctBal          []ContractBalance1                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctBal,omitempty"`
	PmtSchdlTp        PaymentScheduleType1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 PmtSchdlTp,omitempty"`
	RegdCtrctId       DocumentIdentification29                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RegdCtrctId"`
	PrvsRegdCtrctId   DocumentIdentification22                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 PrvsRegdCtrctId,omitempty"`
	RegdCtrctJrnl     []RegisteredContractJournal2                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RegdCtrctJrnl,omitempty"`
	Amdmnt            []RegisteredContractAmendment1               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Amdmnt,omitempty"`
	Submissn          RegisteredContractCommunication1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Submissn"`
	Dlvry             RegisteredContractCommunication1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Dlvry"`
	LnPrncplAmt       ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 LnPrncplAmt,omitempty"`
	EstmtdDtInd       bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 EstmtdDtInd"`
	IntrCpnyLn        bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 IntrCpnyLn"`
	AddtlInf          common.Max1025Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AddtlInf,omitempty"`
}

func (r RegisteredContract8) Validate() error {
	return utils.Validate(&r)
}

type ReportingPeriod4 struct {
	FrToDt DatePeriod3    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 FrToDt"`
	FrToTm TimePeriod2    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 FrToTm"`
	Tp     QueryType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Tp"`
}

func (r ReportingPeriod4) Validate() error {
	return utils.Validate(&r)
}

type ShipmentAttribute1 struct {
	Conds         ExternalShipmentCondition1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Conds,omitempty"`
	XpctdDt       common.ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 XpctdDt,omitempty"`
	CtryOfCntrPty CountryCode                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtryOfCntrPty"`
}

func (r ShipmentAttribute1) Validate() error {
	return utils.Validate(&r)
}

type SupportingDocument2 struct {
	SpprtgDocId common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SpprtgDocId"`
	OrgnlReqId  common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 OrgnlReqId,omitempty"`
	Cert        DocumentIdentification28                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Cert"`
	AcctOwnr    PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AcctOwnr"`
	AcctSvcr    BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AcctSvcr"`
	Amdmnt      DocumentAmendment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Amdmnt,omitempty"`
	CtrctRef    ContractRegistrationReference1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctRef"`
	Ntry        []SupportingDocumentEntry1                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Ntry"`
	SplmtryData []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SplmtryData,omitempty"`
}

func (r SupportingDocument2) Validate() error {
	return utils.Validate(&r)
}

type SupportingDocumentEntry1 struct {
	NtryId                      common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 NtryId"`
	OrgnlDoc                    DocumentIdentification22      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 OrgnlDoc"`
	DocTp                       Exact4AlphaNumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 DocTp"`
	TtlAmt                      ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TtlAmt,omitempty"`
	TtlAmtAftrShipmnt           ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TtlAmtAftrShipmnt,omitempty"`
	TtlAmtInCtrctCcy            ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TtlAmtInCtrctCcy,omitempty"`
	TtlAmtAftrShipmntInCtrctCcy ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TtlAmtAftrShipmntInCtrctCcy,omitempty"`
	ShipmntAttrbts              ShipmentAttribute1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 ShipmntAttrbts"`
	AddtlInf                    common.Max500Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AddtlInf,omitempty"`
	Attchmnt                    []DocumentGeneralInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Attchmnt,omitempty"`
}

func (r SupportingDocumentEntry1) Validate() error {
	return utils.Validate(&r)
}

type TimePeriod2 struct {
	FrTm common.ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 FrTm"`
	ToTm common.ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 ToTm,omitempty"`
}

func (r TimePeriod2) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificate2 struct {
	RfrdDoc   CertificateReference1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 RfrdDoc"`
	TxDt      common.ISODate          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TxDt"`
	TxTp      Exact1NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TxTp"`
	LclInstrm Exact5NumericText       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 LclInstrm"`
	Amt       ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Amt"`
}

func (r TransactionCertificate2) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificate3 struct {
	TxId             common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TxId"`
	Cert             DocumentIdentification28        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Cert"`
	Acct             CashAccount38                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Acct,omitempty"`
	BkAcctDmcltnCtry CountryCode                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 BkAcctDmcltnCtry,omitempty"`
	Amdmnt           DocumentAmendment1              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Amdmnt,omitempty"`
	CertRcrd         []TransactionCertificateRecord1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CertRcrd"`
	SplmtryData      []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 SplmtryData,omitempty"`
}

func (r TransactionCertificate3) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificateContract1 struct {
	CtrctRef           ContractRegistrationReference1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CtrctRef,omitempty"`
	TxAmtInCtrctCcy    ActiveCurrencyAndAmount              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 TxAmtInCtrctCcy,omitempty"`
	XpctdShipmntDt     common.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 XpctdShipmntDt,omitempty"`
	XpctdAdvncPmtRtrDt common.ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 XpctdAdvncPmtRtrDt,omitempty"`
	AddtlInf           common.Max1025Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 AddtlInf,omitempty"`
}

func (r TransactionCertificateContract1) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificateRecord1 struct {
	CertRcrdId common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 CertRcrdId"`
	Tx         TransactionCertificate2         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Tx"`
	Ctrct      TransactionCertificateContract1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Ctrct,omitempty"`
	Attchmnt   []DocumentGeneralInformation3   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Attchmnt,omitempty"`
}

func (r TransactionCertificateRecord1) Validate() error {
	return utils.Validate(&r)
}

type ValidationRuleSchemeName1Choice struct {
	Cd    ExternalValidationRuleIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.02 Prtry"`
}

func (r ValidationRuleSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementCriteria1 struct {
	TxJrnl             bool `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 TxJrnl,omitempty"`
	SpprtgDocJrnl      bool `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 SpprtgDocJrnl,omitempty"`
	AddtlSpprtgDocJrnl bool `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 AddtlSpprtgDocJrnl,omitempty"`
	RgltryRuleVldtn    bool `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 RgltryRuleVldtn,omitempty"`
}

func (r ContractRegistrationStatementCriteria1) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementRequest2 struct {
	StmtReqId   common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 StmtReqId"`
	RptgPrd     ReportingPeriod4                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 RptgPrd"`
	RptgPty     TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 RptgPty"`
	RegnAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 RegnAgt"`
	RegdCtrctId common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 RegdCtrctId"`
	RtrCrit     ContractRegistrationStatementCriteria1       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 RtrCrit,omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationStatementRequest2) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementRequestV02 struct {
	GrpHdr      CurrencyControlHeader4                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 GrpHdr"`
	StmtReq     []ContractRegistrationStatementRequest2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 StmtReq"`
	SplmtryData []SupplementaryData1                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.023.001.02 SplmtryData,omitempty"`
}

func (r ContractRegistrationStatementRequestV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlHeader5 struct {
	MsgId    common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 CreDtTm"`
	NbOfItms Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 NbOfItms"`
	InitgPty Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 InitgPty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 FwdgAgt,omitempty"`
}

func (r CurrencyControlHeader5) Validate() error {
	return utils.Validate(&r)
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 Agt"`
}

func (r Party40Choice) Validate() error {
	return utils.Validate(&r)
}

type PaymentRegulatoryInformationNotificationV02 struct {
	GrpHdr      CurrencyControlHeader5             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 GrpHdr"`
	TxNtfctn    []RegulatoryReportingNotification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 TxNtfctn"`
	SplmtryData []SupplementaryData1               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 SplmtryData,omitempty"`
}

func (r PaymentRegulatoryInformationNotificationV02) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryReportingNotification2 struct {
	TxNtfctnId common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 TxNtfctnId"`
	AcctOwnr   PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 AcctOwnr"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 AcctSvcr"`
	TxCert     []TransactionCertificate3                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.024.001.02 TxCert"`
}

func (r RegulatoryReportingNotification2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlSupportingDocumentDeliveryV02 struct {
	GrpHdr      CurrencyControlHeader5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 GrpHdr"`
	SpprtgDoc   []SupportingDocument2  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SpprtgDoc"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SplmtryData,omitempty"`
}

func (r CurrencyControlSupportingDocumentDeliveryV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlRequestOrLetterV02 struct {
	GrpHdr      CurrencyControlHeader5               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 GrpHdr"`
	ReqOrLttr   []SupportingDocumentRequestOrLetter2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 ReqOrLttr"`
	SplmtryData []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 SplmtryData,omitempty"`
}

func (r CurrencyControlRequestOrLetterV02) Validate() error {
	return utils.Validate(&r)
}

type OriginalMessage4 struct {
	OrgnlSndr    Party40Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlSndr,omitempty"`
	OrgnlMsgId   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlMsgNmId"`
	OrgnlCreDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlCreDtTm,omitempty"`
	OrgnlPackgId common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlPackgId,omitempty"`
	OrgnlRcrdId  common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlRcrdId"`
}

func (r OriginalMessage4) Validate() error {
	return utils.Validate(&r)
}

type SupportingDocumentRequestOrLetter2 struct {
	ReqOrLttrId common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 ReqOrLttrId"`
	Dt          common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Dt,omitempty"`
	Sndr        Party40Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Sndr,omitempty"`
	Rcvr        Party40Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Rcvr,omitempty"`
	OrgnlRefs   []OriginalMessage4            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 OrgnlRefs,omitempty"`
	Sbjt        common.Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Sbjt"`
	Tp          SupportDocumentType1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Tp"`
	Desc        common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Desc,omitempty"`
	RspnReqrd   bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 RspnReqrd"`
	DueDt       common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 DueDt,omitempty"`
	Attchmnt    []DocumentGeneralInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 Attchmnt,omitempty"`
	SplmtryData []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.02 SplmtryData,omitempty"`
}

func (r SupportingDocumentRequestOrLetter2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlGroupStatus2 struct {
	OrgnlRefs OriginalMessage5                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlRefs"`
	RptgPty   TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RptgPty"`
	RegnAgt   BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RegnAgt"`
	RptgPrd   Period4Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RptgPrd,omitempty"`
	Sts       StatisticalReportingStatus1Code              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Sts,omitempty"`
	StsRsn    []ValidationStatusReason2                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsRsn,omitempty"`
	StsDtTm   common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsDtTm,omitempty"`
}

func (r CurrencyControlGroupStatus2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlPackageStatus2 struct {
	PackgId common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PackgId"`
	Sts     StatisticalReportingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Sts"`
	StsRsn  []ValidationStatusReason2       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsRsn,omitempty"`
	StsDtTm common.ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsDtTm,omitempty"`
	RcrdSts []CurrencyControlRecordStatus2  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RcrdSts,omitempty"`
}

func (r CurrencyControlPackageStatus2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlRecordStatus2 struct {
	RcrdId  common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RcrdId"`
	Sts     StatisticalReportingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Sts"`
	StsRsn  []ValidationStatusReason2       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsRsn,omitempty"`
	StsDtTm common.ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsDtTm,omitempty"`
	DocId   DocumentIdentification28        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 DocId,omitempty"`
}

func (r CurrencyControlRecordStatus2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlStatusAdviceV02 struct {
	GrpHdr      CurrencyControlHeader6          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 GrpHdr"`
	GrpSts      []CurrencyControlGroupStatus2   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 GrpSts"`
	PackgSts    []CurrencyControlPackageStatus2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PackgSts,omitempty"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SplmtryData,omitempty"`
}

func (r CurrencyControlStatusAdviceV02) Validate() error {
	return utils.Validate(&r)
}

type OriginalMessage5 struct {
	OrgnlSndr    Party40Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlSndr,omitempty"`
	OrgnlMsgId   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlMsgNmId"`
	OrgnlCreDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlCreDtTm,omitempty"`
}

func (r OriginalMessage5) Validate() error {
	return utils.Validate(&r)
}

type Period2 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ToDt"`
}

func (r Period2) Validate() error {
	return utils.Validate(&r)
}

type Period4Choice struct {
	Dt       common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Dt"`
	FrDt     common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FrDt"`
	ToDt     common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ToDt"`
	FrDtToDt Period2        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FrDtToDt"`
}

func (r Period4Choice) Validate() error {
	return utils.Validate(&r)
}

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

func (r StatusReason6Choice) Validate() error {
	return utils.Validate(&r)
}

type ValidationStatusReason2 struct {
	Orgtr     PartyIdentification135                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Orgtr,omitempty"`
	Rsn       StatusReason6Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Rsn,omitempty"`
	VldtnRule []GenericValidationRuleIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 VldtnRule,omitempty"`
	AddtlInf  []common.Max105Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 AddtlInf,omitempty"`
}

func (r ValidationStatusReason2) Validate() error {
	return utils.Validate(&r)
}
