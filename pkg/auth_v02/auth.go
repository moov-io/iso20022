// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

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

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
	return utils.Validate(&r)
}

type BenchmarkCurveName4Choice struct {
	ISIN ISINOct2015Identifier   `xml:"ISIN"`
	Indx BenchmarkCurveName2Code `xml:"Indx"`
	Nm   common.Max25Text        `xml:"Nm"`
}

func (r BenchmarkCurveName4Choice) Validate() error {
	return utils.Validate(&r)
}

type BinaryFile1 struct {
	MIMETp         *common.Max35Text     `xml:"MIMETp,omitempty" json:",omitempty"`
	NcodgTp        *common.Max35Text     `xml:"NcodgTp,omitempty" json:",omitempty"`
	CharSet        *common.Max35Text     `xml:"CharSet,omitempty" json:",omitempty"`
	InclBinryObjct *common.Max100KBinary `xml:"InclBinryObjct,omitempty" json:",omitempty"`
}

func (r BinaryFile1) Validate() error {
	return utils.Validate(&r)
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

func (r BranchAndFinancialInstitutionIdentification6) Validate() error {
	return utils.Validate(&r)
}

type BranchData3 struct {
	Id      *common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r BranchData3) Validate() error {
	return utils.Validate(&r)
}

type CashCollateral5 struct {
	CollId    *common.Max35Text             `xml:"CollId,omitempty" json:",omitempty"`
	CshAcctId *AccountIdentification4Choice `xml:"CshAcctId,omitempty" json:",omitempty"`
	AsstNb    *common.Max35Text             `xml:"AsstNb,omitempty" json:",omitempty"`
	DpstAmt   *ActiveCurrencyAndAmount      `xml:"DpstAmt,omitempty" json:",omitempty"`
	DpstTp    *DepositType1Code             `xml:"DpstTp,omitempty" json:",omitempty"`
	MtrtyDt   *common.ISODate               `xml:"MtrtyDt,omitempty" json:",omitempty"`
	ValDt     *common.ISODate               `xml:"ValDt,omitempty" json:",omitempty"`
	XchgRate  float64                       `xml:"XchgRate,omitempty" json:",omitempty"`
	CollVal   ActiveCurrencyAndAmount       `xml:"CollVal"`
	Hrcut     float64                       `xml:"Hrcut,omitempty" json:",omitempty"`
}

func (r CashCollateral5) Validate() error {
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

type Contact4 struct {
	NmPrfx    *common.NamePrefix2Code      `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm        *common.Max140Text           `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb    *common.PhoneNumber          `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb     *common.PhoneNumber          `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb     *common.PhoneNumber          `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr  *common.Max2048Text          `xml:"EmailAdr,omitempty" json:",omitempty"`
	EmailPurp *common.Max35Text            `xml:"EmailPurp,omitempty" json:",omitempty"`
	JobTitl   *common.Max35Text            `xml:"JobTitl,omitempty" json:",omitempty"`
	Rspnsblty *common.Max35Text            `xml:"Rspnsblty,omitempty" json:",omitempty"`
	Dept      *common.Max70Text            `xml:"Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1              `xml:"Othr,omitempty" json:",omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:",omitempty"`
}

func (r Contact4) Validate() error {
	return utils.Validate(&r)
}

type ContractBalance1 struct {
	Tp        ContractBalanceType1Choice `xml:"Tp"`
	Amt       ActiveCurrencyAndAmount    `xml:"Amt"`
	CdtDbtInd CreditDebit3Code           `xml:"CdtDbtInd"`
}

func (r ContractBalance1) Validate() error {
	return utils.Validate(&r)
}

type ContractBalanceType1Choice struct {
	Cd    ExternalContractBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

func (r ContractBalanceType1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractCollateral1 struct {
	TtlAmt   ActiveCurrencyAndAmount `xml:"TtlAmt"`
	CollDesc []CashCollateral5       `xml:"CollDesc,omitempty" json:",omitempty"`
	AddtlInf *common.Max1025Text     `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r ContractCollateral1) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistration3 struct {
	CtrctRegnId   common.Max35Text                             `xml:"CtrctRegnId"`
	RptgPty       TradeParty5                                  `xml:"RptgPty"`
	RegnAgt       BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	CtrctRegnOpng []ContractRegistration4                      `xml:"CtrctRegnOpng" json:",omitempty"`
	SplmtryData   []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistration3) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistration4 struct {
	CtrctRegnOpngId common.Max35Text              `xml:"CtrctRegnOpngId"`
	Prty            Priority2Code                 `xml:"Prty"`
	Ctrct           UnderlyingContract2Choice     `xml:"Ctrct"`
	CtrctBal        []ContractBalance1            `xml:"CtrctBal,omitempty" json:",omitempty"`
	PmtSchdlTp      *PaymentScheduleType1Choice   `xml:"PmtSchdlTp,omitempty" json:",omitempty"`
	PrvsRegnId      []DocumentIdentification22    `xml:"PrvsRegnId,omitempty" json:",omitempty"`
	AddtlInf        *common.Max1025Text           `xml:"AddtlInf,omitempty" json:",omitempty"`
	Attchmnt        []DocumentGeneralInformation3 `xml:"Attchmnt,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistration4) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationRequestV02 struct {
	Attr        []utils.Attr            `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader4  `xml:"GrpHdr"`
	CtrctRegn   []ContractRegistration3 `xml:"CtrctRegn" json:",omitempty"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationRequestV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlHeader4 struct {
	MsgId    common.Max35Text                              `xml:"MsgId"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm"`
	NbOfItms common.Max15NumericText                       `xml:"NbOfItms"`
	InitgPty PartyIdentification135                        `xml:"InitgPty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"FwdgAgt,omitempty" json:",omitempty"`
}

func (r CurrencyControlHeader4) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
	return utils.Validate(&r)
}

type DocumentGeneralInformation3 struct {
	DocTp           ExternalDocumentType1Code   `xml:"DocTp"`
	DocNb           common.Max35Text            `xml:"DocNb"`
	SndrRcvrSeqId   *common.Max140Text          `xml:"SndrRcvrSeqId,omitempty" json:",omitempty"`
	IsseDt          *common.ISODate             `xml:"IsseDt,omitempty" json:",omitempty"`
	URL             *common.Max256Text          `xml:"URL,omitempty" json:",omitempty"`
	LkFileHash      *SignatureEnvelopeReference `xml:"LkFileHash,omitempty" json:",omitempty"`
	AttchdBinryFile BinaryFile1                 `xml:"AttchdBinryFile"`
}

func (r DocumentGeneralInformation3) Validate() error {
	return utils.Validate(&r)
}

type DocumentIdentification22 struct {
	Id       common.Max35Text `xml:"Id"`
	DtOfIsse *common.ISODate  `xml:"DtOfIsse,omitempty" json:",omitempty"`
}

func (r DocumentIdentification22) Validate() error {
	return utils.Validate(&r)
}

type ExchangeRate1 struct {
	UnitCcy  *common.ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty" json:",omitempty"`
	XchgRate float64                              `xml:"XchgRate,omitempty" json:",omitempty"`
	RateTp   *ExchangeRateType1Code               `xml:"RateTp,omitempty" json:",omitempty"`
	CtrctId  *common.Max35Text                    `xml:"CtrctId,omitempty" json:",omitempty"`
}

func (r ExchangeRate1) Validate() error {
	return utils.Validate(&r)
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

func (r FinancialIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

func (r FinancialInstitutionIdentification18) Validate() error {
	return utils.Validate(&r)
}

type FloatingInterestRate4 struct {
	RefRate    BenchmarkCurveName4Choice `xml:"RefRate"`
	Term       InterestRateContractTerm1 `xml:"Term"`
	BsisPtSprd float64                   `xml:"BsisPtSprd"`
}

func (r FloatingInterestRate4) Validate() error {
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

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
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

type InterestPaymentDateRange1 struct {
	IntrstSchdlId *common.Max35Text `xml:"IntrstSchdlId,omitempty" json:",omitempty"`
	XpctdDt       *common.ISODate   `xml:"XpctdDt,omitempty" json:",omitempty"`
	DueDt         *common.ISODate   `xml:"DueDt,omitempty" json:",omitempty"`
}

func (r InterestPaymentDateRange1) Validate() error {
	return utils.Validate(&r)
}

type InterestPaymentDateRange2 struct {
	IntrstSchdlId *common.Max35Text       `xml:"IntrstSchdlId,omitempty" json:",omitempty"`
	Amt           ActiveCurrencyAndAmount `xml:"Amt"`
	DueDt         common.ISODate          `xml:"DueDt"`
	AddtlInf      *common.Max1025Text     `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r InterestPaymentDateRange2) Validate() error {
	return utils.Validate(&r)
}

type InterestPaymentSchedule1Choice struct {
	DtRg     InterestPaymentDateRange1   `xml:"DtRg"`
	SubSchdl []InterestPaymentDateRange2 `xml:"SubSchdl" json:",omitempty"`
}

func (r InterestPaymentSchedule1Choice) Validate() error {
	return utils.Validate(&r)
}

type InterestRate2Choice struct {
	Fxd  float64               `xml:"Fxd"`
	Fltg FloatingInterestRate4 `xml:"Fltg"`
}

func (r InterestRate2Choice) Validate() error {
	return utils.Validate(&r)
}

type InterestRateContractTerm1 struct {
	Unit RateBasis1Code `xml:"Unit"`
	Val  float64        `xml:"Val"`
}

func (r InterestRateContractTerm1) Validate() error {
	return utils.Validate(&r)
}

type LegalOrganisation2 struct {
	Id           *common.Max35Text  `xml:"Id,omitempty" json:",omitempty"`
	Nm           *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
	EstblishmtDt *common.ISODate    `xml:"EstblishmtDt,omitempty" json:",omitempty"`
	RegnDt       *common.ISODate    `xml:"RegnDt,omitempty" json:",omitempty"`
}

func (r LegalOrganisation2) Validate() error {
	return utils.Validate(&r)
}

type LoanContract2 struct {
	CtrctDocId  DocumentIdentification22       `xml:"CtrctDocId"`
	Buyr        []TradeParty5                  `xml:"Buyr" json:",omitempty"`
	Sellr       []TradeParty5                  `xml:"Sellr" json:",omitempty"`
	Amt         ActiveCurrencyAndAmount        `xml:"Amt"`
	MtrtyDt     common.ISODate                 `xml:"MtrtyDt"`
	PrlngtnFlg  bool                           `xml:"PrlngtnFlg"`
	StartDt     common.ISODate                 `xml:"StartDt"`
	SttlmCcy    common.ActiveCurrencyCode      `xml:"SttlmCcy"`
	SpclConds   *SpecialCondition1             `xml:"SpclConds,omitempty" json:",omitempty"`
	DrtnCd      Exact1NumericText              `xml:"DrtnCd"`
	IntrstRate  InterestRate2Choice            `xml:"IntrstRate"`
	Trch        []LoanContractTranche1         `xml:"Trch,omitempty" json:",omitempty"`
	PmtSchdl    *PaymentSchedule1Choice        `xml:"PmtSchdl,omitempty" json:",omitempty"`
	IntrstSchdl InterestPaymentSchedule1Choice `xml:"IntrstSchdl"`
	IntraCpnyLn bool                           `xml:"IntraCpnyLn"`
	Coll        *ContractCollateral1           `xml:"Coll,omitempty" json:",omitempty"`
	SndctdLn    []SyndicatedLoan2              `xml:"SndctdLn,omitempty" json:",omitempty"`
	Attchmnt    []DocumentGeneralInformation3  `xml:"Attchmnt,omitempty" json:",omitempty"`
}

func (r LoanContract2) Validate() error {
	return utils.Validate(&r)
}

type LoanContractTranche1 struct {
	TrchNb      float64                 `xml:"TrchNb"`
	XpctdDt     common.ISODate          `xml:"XpctdDt"`
	Amt         ActiveCurrencyAndAmount `xml:"Amt"`
	DueDt       *common.ISODate         `xml:"DueDt,omitempty" json:",omitempty"`
	DrtnCd      *Exact1NumericText      `xml:"DrtnCd,omitempty" json:",omitempty"`
	LastTrchInd bool                    `xml:"LastTrchInd,omitempty" json:",omitempty"`
}

func (r LoanContractTranche1) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification29) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `xml:"ChanlTp"`
	Id      *common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

func (r OtherContact1) Validate() error {
	return utils.Validate(&r)
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

func (r Party38Choice) Validate() error {
	return utils.Validate(&r)
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification135) Validate() error {
	return utils.Validate(&r)
}

type PaymentDateRange1 struct {
	PmtSchdlId *common.Max35Text `xml:"PmtSchdlId,omitempty" json:",omitempty"`
	XpctdDt    *common.ISODate   `xml:"XpctdDt,omitempty" json:",omitempty"`
	DueDt      *common.ISODate   `xml:"DueDt,omitempty" json:",omitempty"`
}

func (r PaymentDateRange1) Validate() error {
	return utils.Validate(&r)
}

type PaymentDateRange2 struct {
	PmtSchdlId *common.Max35Text       `xml:"PmtSchdlId,omitempty" json:",omitempty"`
	Amt        ActiveCurrencyAndAmount `xml:"Amt"`
	XpctdDt    *common.ISODate         `xml:"XpctdDt,omitempty" json:",omitempty"`
	DueDt      common.ISODate          `xml:"DueDt"`
}

func (r PaymentDateRange2) Validate() error {
	return utils.Validate(&r)
}

type PaymentSchedule1Choice struct {
	DtRg     PaymentDateRange1   `xml:"DtRg"`
	SubSchdl []PaymentDateRange2 `xml:"SubSchdl" json:",omitempty"`
}

func (r PaymentSchedule1Choice) Validate() error {
	return utils.Validate(&r)
}

type PaymentScheduleType1Choice struct {
	Cd    PaymentScheduleType1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r PaymentScheduleType1Choice) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification13) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        *common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     *common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      *common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      *common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	BldgNm      *common.Max35Text   `xml:"BldgNm,omitempty" json:",omitempty"`
	Flr         *common.Max70Text   `xml:"Flr,omitempty" json:",omitempty"`
	PstBx       *common.Max16Text   `xml:"PstBx,omitempty" json:",omitempty"`
	Room        *common.Max70Text   `xml:"Room,omitempty" json:",omitempty"`
	PstCd       *common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       *common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   *common.Max35Text   `xml:"TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    *common.Max35Text   `xml:"DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn *common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text  `xml:"AdrLine,omitempty" json:",omitempty"`
}

func (r PostalAddress24) Validate() error {
	return utils.Validate(&r)
}

type ShipmentDateRange1 struct {
	EarlstShipmntDt *common.ISODate `xml:"EarlstShipmntDt,omitempty" json:",omitempty"`
	LatstShipmntDt  *common.ISODate `xml:"LatstShipmntDt,omitempty" json:",omitempty"`
}

func (r ShipmentDateRange1) Validate() error {
	return utils.Validate(&r)
}

type ShipmentDateRange2 struct {
	SubQtyVal       float64         `xml:"SubQtyVal"`
	EarlstShipmntDt *common.ISODate `xml:"EarlstShipmntDt,omitempty" json:",omitempty"`
	LatstShipmntDt  *common.ISODate `xml:"LatstShipmntDt,omitempty" json:",omitempty"`
}

func (r ShipmentDateRange2) Validate() error {
	return utils.Validate(&r)
}

type ShipmentSchedule2Choice struct {
	ShipmntDtRg     ShipmentDateRange1   `xml:"ShipmntDtRg"`
	ShipmntSubSchdl []ShipmentDateRange2 `xml:"ShipmntSubSchdl" json:",omitempty"`
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
	IncmgAmt           ActiveCurrencyAndAmount  `xml:"IncmgAmt"`
	OutgngAmt          ActiveCurrencyAndAmount  `xml:"OutgngAmt"`
	IncmgAmtToOthrAcct *ActiveCurrencyAndAmount `xml:"IncmgAmtToOthrAcct,omitempty" json:",omitempty"`
	PmtFrOthrAcct      *ActiveCurrencyAndAmount `xml:"PmtFrOthrAcct,omitempty" json:",omitempty"`
}

func (r SpecialCondition1) Validate() error {
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

type SyndicatedLoan2 struct {
	Brrwr       TradeParty5              `xml:"Brrwr"`
	Lndr        *TradeParty5             `xml:"Lndr,omitempty" json:",omitempty"`
	Amt         *ActiveCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Shr         float64                  `xml:"Shr,omitempty" json:",omitempty"`
	XchgRateInf *ExchangeRate1           `xml:"XchgRateInf,omitempty" json:",omitempty"`
}

func (r SyndicatedLoan2) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptionReasonFormat1Choice struct {
	Ustrd common.Max140Text    `xml:"Ustrd"`
	Strd  TaxExemptReason1Code `xml:"Strd"`
}

func (r TaxExemptionReasonFormat1Choice) Validate() error {
	return utils.Validate(&r)
}

type TaxParty4 struct {
	TaxId       *common.Max35Text                 `xml:"TaxId,omitempty" json:",omitempty"`
	TaxTp       *common.Max35Text                 `xml:"TaxTp,omitempty" json:",omitempty"`
	RegnId      *common.Max35Text                 `xml:"RegnId,omitempty" json:",omitempty"`
	TaxXmptnRsn []TaxExemptionReasonFormat1Choice `xml:"TaxXmptnRsn,omitempty" json:",omitempty"`
}

func (r TaxParty4) Validate() error {
	return utils.Validate(&r)
}

type TradeContract2 struct {
	CtrctDocId   *DocumentIdentification22     `xml:"CtrctDocId,omitempty" json:",omitempty"`
	Amt          ActiveCurrencyAndAmount       `xml:"Amt"`
	Buyr         []TradeParty5                 `xml:"Buyr" json:",omitempty"`
	Sellr        []TradeParty5                 `xml:"Sellr" json:",omitempty"`
	MtrtyDt      common.ISODate                `xml:"MtrtyDt"`
	PrlngtnFlg   bool                          `xml:"PrlngtnFlg"`
	StartDt      common.ISODate                `xml:"StartDt"`
	SttlmCcy     common.ActiveCurrencyCode     `xml:"SttlmCcy"`
	XchgRateInf  *ExchangeRate1                `xml:"XchgRateInf,omitempty" json:",omitempty"`
	PmtSchdl     *InterestPaymentDateRange1    `xml:"PmtSchdl,omitempty" json:",omitempty"`
	ShipmntSchdl *ShipmentSchedule2Choice      `xml:"ShipmntSchdl,omitempty" json:",omitempty"`
	Attchmnt     []DocumentGeneralInformation3 `xml:"Attchmnt,omitempty" json:",omitempty"`
}

func (r TradeContract2) Validate() error {
	return utils.Validate(&r)
}

type TradeParty5 struct {
	PtyId  PartyIdentification135 `xml:"PtyId"`
	LglOrg *LegalOrganisation2    `xml:"LglOrg,omitempty" json:",omitempty"`
	TaxPty []TaxParty4            `xml:"TaxPty,omitempty" json:",omitempty"`
}

func (r TradeParty5) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingContract2Choice struct {
	Ln   LoanContract2  `xml:"Ln"`
	Trad TradeContract2 `xml:"Trad"`
}

func (r UnderlyingContract2Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractClosureReason1Choice struct {
	Cd    ExternalContractClosureReason1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

func (r ContractClosureReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationConfirmationV02 struct {
	Attr        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader6 `xml:"GrpHdr"`
	RegdCtrct   []RegisteredContract7  `xml:"RegdCtrct" json:",omitempty"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationConfirmationV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlHeader6 struct {
	MsgId    common.Max35Text                             `xml:"MsgId"`
	CreDtTm  common.ISODateTime                           `xml:"CreDtTm"`
	NbOfItms common.Max15NumericText                      `xml:"NbOfItms"`
	RcvgPty  PartyIdentification135                       `xml:"RcvgPty"`
	RegnAgt  BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
}

func (r CurrencyControlHeader6) Validate() error {
	return utils.Validate(&r)
}

type DocumentIdentification28 struct {
	Id       *common.Max35Text `xml:"Id,omitempty" json:",omitempty"`
	DtOfIsse common.ISODate    `xml:"DtOfIsse"`
}

func (r DocumentIdentification28) Validate() error {
	return utils.Validate(&r)
}

type DocumentIdentification29 struct {
	Id       common.Max35Text `xml:"Id"`
	DtOfIsse common.ISODate   `xml:"DtOfIsse"`
}

func (r DocumentIdentification29) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract7 struct {
	OrgnlCtrctRegnReq *common.Max35Text                            `xml:"OrgnlCtrctRegnReq,omitempty" json:",omitempty"`
	RptgPty           TradeParty5                                  `xml:"RptgPty"`
	RegnAgt           BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	IssrFI            BranchAndFinancialInstitutionIdentification6 `xml:"IssrFI"`
	Ctrct             UnderlyingContract2Choice                    `xml:"Ctrct"`
	CtrctBal          []ContractBalance1                           `xml:"CtrctBal,omitempty" json:",omitempty"`
	PmtSchdlTp        *PaymentScheduleType1Choice                  `xml:"PmtSchdlTp,omitempty" json:",omitempty"`
	RegdCtrctId       DocumentIdentification29                     `xml:"RegdCtrctId"`
	PrvsRegdCtrctId   *DocumentIdentification22                    `xml:"PrvsRegdCtrctId,omitempty" json:",omitempty"`
	RegdCtrctJrnl     []RegisteredContractJournal2                 `xml:"RegdCtrctJrnl,omitempty" json:",omitempty"`
	Amdmnt            []RegisteredContractAmendment1               `xml:"Amdmnt,omitempty" json:",omitempty"`
	Submissn          RegisteredContractCommunication1             `xml:"Submissn"`
	Dlvry             RegisteredContractCommunication1             `xml:"Dlvry"`
	LnPrncplAmt       *ActiveCurrencyAndAmount                     `xml:"LnPrncplAmt,omitempty" json:",omitempty"`
	EstmtdDtInd       bool                                         `xml:"EstmtdDtInd"`
	IntrCpnyLn        bool                                         `xml:"IntrCpnyLn"`
	AddtlInf          *common.Max1025Text                          `xml:"AddtlInf,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RegisteredContract7) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContractAmendment1 struct {
	AmdmntDt  common.ISODate           `xml:"AmdmntDt"`
	Doc       DocumentIdentification28 `xml:"Doc"`
	StartDt   *common.ISODate          `xml:"StartDt,omitempty" json:",omitempty"`
	AmdmntRsn *common.Max35Text        `xml:"AmdmntRsn,omitempty" json:",omitempty"`
	AddtlInf  *common.Max1025Text      `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r RegisteredContractAmendment1) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContractCommunication1 struct {
	Mtd CommunicationMethod4Code `xml:"Mtd"`
	Dt  common.ISODate           `xml:"Dt"`
}

func (r RegisteredContractCommunication1) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContractJournal2 struct {
	RegnAgt BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	UnqId   *DocumentIdentification28                    `xml:"UnqId,omitempty" json:",omitempty"`
	ClsrDt  common.ISODate                               `xml:"ClsrDt"`
	ClsrRsn ContractClosureReason1Choice                 `xml:"ClsrRsn"`
}

func (r RegisteredContractJournal2) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationClosureRequestV02 struct {
	Attr          []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr        CurrencyControlHeader4 `xml:"GrpHdr"`
	RegdCtrctClsr []RegisteredContract6  `xml:"RegdCtrctClsr" json:",omitempty"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationClosureRequestV02) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract6 struct {
	RegdCtrctClsrId common.Max35Text                             `xml:"RegdCtrctClsrId"`
	RptgPty         TradeParty5                                  `xml:"RptgPty"`
	RegnAgt         BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	OrgnlRegdCtrct  DocumentIdentification29                     `xml:"OrgnlRegdCtrct"`
	Prty            Priority2Code                                `xml:"Prty"`
	ClsrRsn         ContractClosureReason1Choice                 `xml:"ClsrRsn"`
	SplmtryData     []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RegisteredContract6) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationAmendmentRequestV02 struct {
	Attr            []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr          CurrencyControlHeader4 `xml:"GrpHdr"`
	CtrctRegnAmdmnt []RegisteredContract9  `xml:"CtrctRegnAmdmnt" json:",omitempty"`
	SplmtryData     []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationAmendmentRequestV02) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract10 struct {
	RegdCtrctAmdmntId common.Max35Text              `xml:"RegdCtrctAmdmntId"`
	OrgnlRegdCtrctId  common.Max35Text              `xml:"OrgnlRegdCtrctId"`
	Prty              Priority2Code                 `xml:"Prty"`
	Ctrct             UnderlyingContract2Choice     `xml:"Ctrct"`
	CtrctBal          []ContractBalance1            `xml:"CtrctBal,omitempty" json:",omitempty"`
	PmtSchdlTp        *PaymentScheduleType1Choice   `xml:"PmtSchdlTp,omitempty" json:",omitempty"`
	AddtlInf          *common.Max1025Text           `xml:"AddtlInf,omitempty" json:",omitempty"`
	Attchmnt          []DocumentGeneralInformation3 `xml:"Attchmnt,omitempty" json:",omitempty"`
	SplmtryData       []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RegisteredContract10) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract9 struct {
	CtrctRegnAmdmntId common.Max35Text                             `xml:"CtrctRegnAmdmntId"`
	RptgPty           TradeParty5                                  `xml:"RptgPty"`
	RegnAgt           BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	RegdCtrctAmdmnt   []RegisteredContract10                       `xml:"RegdCtrctAmdmnt" json:",omitempty"`
	SplmtryData       []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RegisteredContract9) Validate() error {
	return utils.Validate(&r)
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice         `xml:"Id"`
	Tp   *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

func (r CashAccount38) Validate() error {
	return utils.Validate(&r)
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CashAccountType2Choice) Validate() error {
	return utils.Validate(&r)
}

type CertificateIdentification1 struct {
	MsgId       *common.Max35Text      `xml:"MsgId,omitempty" json:",omitempty"`
	AcctSvcrRef *common.Max35Text      `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	PmtInfId    *common.Max35Text      `xml:"PmtInfId,omitempty" json:",omitempty"`
	InstrId     *common.Max35Text      `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId  *common.Max35Text      `xml:"EndToEndId,omitempty" json:",omitempty"`
	Prtry       *ProprietaryReference1 `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r CertificateIdentification1) Validate() error {
	return utils.Validate(&r)
}

type CertificateReference1 struct {
	Id CertificateIdentification1 `xml:"Id"`
	Dt common.ISODate             `xml:"Dt"`
}

func (r CertificateReference1) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationReference1Choice struct {
	RegdCtrctId common.Max35Text         `xml:"RegdCtrctId"`
	Ctrct       DocumentIdentification28 `xml:"Ctrct"`
}

func (r ContractRegistrationReference1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatement2 struct {
	StmtId             *common.Max35Text                            `xml:"StmtId,omitempty" json:",omitempty"`
	RptgPty            TradeParty5                                  `xml:"RptgPty"`
	RegnAgt            BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	RptgPrd            ReportingPeriod4                             `xml:"RptgPrd"`
	RegdCtrct          RegisteredContract8                          `xml:"RegdCtrct"`
	TxJrnl             []TransactionCertificate3                    `xml:"TxJrnl,omitempty" json:",omitempty"`
	SpprtgDocJrnl      []SupportingDocument2                        `xml:"SpprtgDocJrnl,omitempty" json:",omitempty"`
	AddtlSpprtgDocJrnl []SupportingDocument2                        `xml:"AddtlSpprtgDocJrnl,omitempty" json:",omitempty"`
	RgltryRuleVldtn    []GenericValidationRuleIdentification1       `xml:"RgltryRuleVldtn,omitempty" json:",omitempty"`
	TtlCtrctTrnvrSum   ActiveCurrencyAndAmount                      `xml:"TtlCtrctTrnvrSum"`
	SplmtryData        []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationStatement2) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementV02 struct {
	Attr        []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader6           `xml:"GrpHdr"`
	Stmt        []ContractRegistrationStatement2 `xml:"Stmt" json:",omitempty"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationStatementV02) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod3 struct {
	FrDt common.ISODate  `xml:"FrDt"`
	ToDt *common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

func (r DatePeriod3) Validate() error {
	return utils.Validate(&r)
}

type DocumentAmendment1 struct {
	CrrctnId   float64           `xml:"CrrctnId"`
	OrgnlDocId *common.Max35Text `xml:"OrgnlDocId,omitempty" json:",omitempty"`
}

func (r DocumentAmendment1) Validate() error {
	return utils.Validate(&r)
}

type GenericValidationRuleIdentification1 struct {
	Id      common.Max35Text                 `xml:"Id"`
	Desc    *common.Max350Text               `xml:"Desc,omitempty" json:",omitempty"`
	SchmeNm *ValidationRuleSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericValidationRuleIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryReference1 struct {
	Tp  common.Max35Text `xml:"Tp"`
	Ref common.Max35Text `xml:"Ref"`
}

func (r ProprietaryReference1) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

func (r ProxyAccountIdentification1) Validate() error {
	return utils.Validate(&r)
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

func (r ProxyAccountType1Choice) Validate() error {
	return utils.Validate(&r)
}

type RegisteredContract8 struct {
	OrgnlCtrctRegnReq *common.Max35Text                            `xml:"OrgnlCtrctRegnReq,omitempty" json:",omitempty"`
	IssrFI            BranchAndFinancialInstitutionIdentification6 `xml:"IssrFI"`
	Ctrct             UnderlyingContract2Choice                    `xml:"Ctrct"`
	CtrctBal          []ContractBalance1                           `xml:"CtrctBal,omitempty" json:",omitempty"`
	PmtSchdlTp        *PaymentScheduleType1Choice                  `xml:"PmtSchdlTp,omitempty" json:",omitempty"`
	RegdCtrctId       DocumentIdentification29                     `xml:"RegdCtrctId"`
	PrvsRegdCtrctId   *DocumentIdentification22                    `xml:"PrvsRegdCtrctId,omitempty" json:",omitempty"`
	RegdCtrctJrnl     []RegisteredContractJournal2                 `xml:"RegdCtrctJrnl,omitempty" json:",omitempty"`
	Amdmnt            []RegisteredContractAmendment1               `xml:"Amdmnt,omitempty" json:",omitempty"`
	Submissn          RegisteredContractCommunication1             `xml:"Submissn"`
	Dlvry             RegisteredContractCommunication1             `xml:"Dlvry"`
	LnPrncplAmt       *ActiveCurrencyAndAmount                     `xml:"LnPrncplAmt,omitempty" json:",omitempty"`
	EstmtdDtInd       bool                                         `xml:"EstmtdDtInd"`
	IntrCpnyLn        bool                                         `xml:"IntrCpnyLn"`
	AddtlInf          *common.Max1025Text                          `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r RegisteredContract8) Validate() error {
	return utils.Validate(&r)
}

type ReportingPeriod4 struct {
	FrToDt DatePeriod3    `xml:"FrToDt"`
	FrToTm TimePeriod2    `xml:"FrToTm"`
	Tp     QueryType3Code `xml:"Tp"`
}

func (r ReportingPeriod4) Validate() error {
	return utils.Validate(&r)
}

type ShipmentAttribute1 struct {
	Conds         *ExternalShipmentCondition1Code `xml:"Conds,omitempty" json:",omitempty"`
	XpctdDt       *common.ISODate                 `xml:"XpctdDt,omitempty" json:",omitempty"`
	CtryOfCntrPty common.CountryCode              `xml:"CtryOfCntrPty"`
}

func (r ShipmentAttribute1) Validate() error {
	return utils.Validate(&r)
}

type SupportingDocument2 struct {
	SpprtgDocId common.Max35Text                             `xml:"SpprtgDocId"`
	OrgnlReqId  *common.Max35Text                            `xml:"OrgnlReqId,omitempty" json:",omitempty"`
	Cert        DocumentIdentification28                     `xml:"Cert"`
	AcctOwnr    PartyIdentification135                       `xml:"AcctOwnr"`
	AcctSvcr    BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr"`
	Amdmnt      *DocumentAmendment1                          `xml:"Amdmnt,omitempty" json:",omitempty"`
	CtrctRef    ContractRegistrationReference1Choice         `xml:"CtrctRef"`
	Ntry        []SupportingDocumentEntry1                   `xml:"Ntry" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r SupportingDocument2) Validate() error {
	return utils.Validate(&r)
}

type SupportingDocumentEntry1 struct {
	NtryId                      common.Max35Text              `xml:"NtryId"`
	OrgnlDoc                    DocumentIdentification22      `xml:"OrgnlDoc"`
	DocTp                       common.Exact4AlphaNumericText `xml:"DocTp"`
	TtlAmt                      *ActiveCurrencyAndAmount      `xml:"TtlAmt,omitempty" json:",omitempty"`
	TtlAmtAftrShipmnt           *ActiveCurrencyAndAmount      `xml:"TtlAmtAftrShipmnt,omitempty" json:",omitempty"`
	TtlAmtInCtrctCcy            *ActiveCurrencyAndAmount      `xml:"TtlAmtInCtrctCcy,omitempty" json:",omitempty"`
	TtlAmtAftrShipmntInCtrctCcy *ActiveCurrencyAndAmount      `xml:"TtlAmtAftrShipmntInCtrctCcy,omitempty" json:",omitempty"`
	ShipmntAttrbts              ShipmentAttribute1            `xml:"ShipmntAttrbts"`
	AddtlInf                    *common.Max500Text            `xml:"AddtlInf,omitempty" json:",omitempty"`
	Attchmnt                    []DocumentGeneralInformation3 `xml:"Attchmnt,omitempty" json:",omitempty"`
}

func (r SupportingDocumentEntry1) Validate() error {
	return utils.Validate(&r)
}

type TimePeriod2 struct {
	FrTm common.ISOTime  `xml:"FrTm"`
	ToTm *common.ISOTime `xml:"ToTm,omitempty" json:",omitempty"`
}

func (r TimePeriod2) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificate2 struct {
	RfrdDoc   CertificateReference1   `xml:"RfrdDoc"`
	TxDt      common.ISODate          `xml:"TxDt"`
	TxTp      Exact1NumericText       `xml:"TxTp"`
	LclInstrm Exact5NumericText       `xml:"LclInstrm"`
	Amt       ActiveCurrencyAndAmount `xml:"Amt"`
}

func (r TransactionCertificate2) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificate3 struct {
	TxId             common.Max35Text                `xml:"TxId"`
	Cert             DocumentIdentification28        `xml:"Cert"`
	Acct             *CashAccount38                  `xml:"Acct,omitempty" json:",omitempty"`
	BkAcctDmcltnCtry *common.CountryCode             `xml:"BkAcctDmcltnCtry,omitempty" json:",omitempty"`
	Amdmnt           *DocumentAmendment1             `xml:"Amdmnt,omitempty" json:",omitempty"`
	CertRcrd         []TransactionCertificateRecord1 `xml:"CertRcrd" json:",omitempty"`
	SplmtryData      []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r TransactionCertificate3) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificateContract1 struct {
	CtrctRef           *ContractRegistrationReference1Choice `xml:"CtrctRef,omitempty" json:",omitempty"`
	TxAmtInCtrctCcy    *ActiveCurrencyAndAmount              `xml:"TxAmtInCtrctCcy,omitempty" json:",omitempty"`
	XpctdShipmntDt     *common.ISODate                       `xml:"XpctdShipmntDt,omitempty" json:",omitempty"`
	XpctdAdvncPmtRtrDt *common.ISODate                       `xml:"XpctdAdvncPmtRtrDt,omitempty" json:",omitempty"`
	AddtlInf           *common.Max1025Text                   `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TransactionCertificateContract1) Validate() error {
	return utils.Validate(&r)
}

type TransactionCertificateRecord1 struct {
	CertRcrdId common.Max35Text                 `xml:"CertRcrdId"`
	Tx         TransactionCertificate2          `xml:"Tx"`
	Ctrct      *TransactionCertificateContract1 `xml:"Ctrct,omitempty" json:",omitempty"`
	Attchmnt   []DocumentGeneralInformation3    `xml:"Attchmnt,omitempty" json:",omitempty"`
}

func (r TransactionCertificateRecord1) Validate() error {
	return utils.Validate(&r)
}

type ValidationRuleSchemeName1Choice struct {
	Cd    ExternalValidationRuleIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r ValidationRuleSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementCriteria1 struct {
	TxJrnl             bool `xml:"TxJrnl,omitempty" json:",omitempty"`
	SpprtgDocJrnl      bool `xml:"SpprtgDocJrnl,omitempty" json:",omitempty"`
	AddtlSpprtgDocJrnl bool `xml:"AddtlSpprtgDocJrnl,omitempty" json:",omitempty"`
	RgltryRuleVldtn    bool `xml:"RgltryRuleVldtn,omitempty" json:",omitempty"`
}

func (r ContractRegistrationStatementCriteria1) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementRequest2 struct {
	StmtReqId   common.Max35Text                             `xml:"StmtReqId"`
	RptgPrd     ReportingPeriod4                             `xml:"RptgPrd"`
	RptgPty     TradeParty5                                  `xml:"RptgPty"`
	RegnAgt     BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	RegdCtrctId common.Max35Text                             `xml:"RegdCtrctId"`
	RtrCrit     *ContractRegistrationStatementCriteria1      `xml:"RtrCrit,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationStatementRequest2) Validate() error {
	return utils.Validate(&r)
}

type ContractRegistrationStatementRequestV02 struct {
	Attr        []utils.Attr                            `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader4                  `xml:"GrpHdr"`
	StmtReq     []ContractRegistrationStatementRequest2 `xml:"StmtReq" json:",omitempty"`
	SplmtryData []SupplementaryData1                    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ContractRegistrationStatementRequestV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlHeader5 struct {
	MsgId    common.Max35Text                              `xml:"MsgId"`
	CreDtTm  common.ISODateTime                            `xml:"CreDtTm"`
	NbOfItms common.Max15NumericText                       `xml:"NbOfItms"`
	InitgPty Party40Choice                                 `xml:"InitgPty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"FwdgAgt,omitempty" json:",omitempty"`
}

func (r CurrencyControlHeader5) Validate() error {
	return utils.Validate(&r)
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (r Party40Choice) Validate() error {
	return utils.Validate(&r)
}

type PaymentRegulatoryInformationNotificationV02 struct {
	Attr        []utils.Attr                       `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader5             `xml:"GrpHdr"`
	TxNtfctn    []RegulatoryReportingNotification2 `xml:"TxNtfctn" json:",omitempty"`
	SplmtryData []SupplementaryData1               `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r PaymentRegulatoryInformationNotificationV02) Validate() error {
	return utils.Validate(&r)
}

type RegulatoryReportingNotification2 struct {
	TxNtfctnId common.Max35Text                             `xml:"TxNtfctnId"`
	AcctOwnr   PartyIdentification135                       `xml:"AcctOwnr"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr"`
	TxCert     []TransactionCertificate3                    `xml:"TxCert" json:",omitempty"`
}

func (r RegulatoryReportingNotification2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlSupportingDocumentDeliveryV02 struct {
	Attr        []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader5 `xml:"GrpHdr"`
	SpprtgDoc   []SupportingDocument2  `xml:"SpprtgDoc" json:",omitempty"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CurrencyControlSupportingDocumentDeliveryV02) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlRequestOrLetterV02 struct {
	Attr        []utils.Attr                         `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader5               `xml:"GrpHdr"`
	ReqOrLttr   []SupportingDocumentRequestOrLetter2 `xml:"ReqOrLttr" json:",omitempty"`
	SplmtryData []SupplementaryData1                 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CurrencyControlRequestOrLetterV02) Validate() error {
	return utils.Validate(&r)
}

type OriginalMessage4 struct {
	OrgnlSndr    *Party40Choice      `xml:"OrgnlSndr,omitempty" json:",omitempty"`
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlPackgId *common.Max35Text   `xml:"OrgnlPackgId,omitempty" json:",omitempty"`
	OrgnlRcrdId  common.Max35Text    `xml:"OrgnlRcrdId"`
}

func (r OriginalMessage4) Validate() error {
	return utils.Validate(&r)
}

type SupportingDocumentRequestOrLetter2 struct {
	ReqOrLttrId common.Max35Text              `xml:"ReqOrLttrId"`
	Dt          *common.ISODate               `xml:"Dt,omitempty" json:",omitempty"`
	Sndr        *Party40Choice                `xml:"Sndr,omitempty" json:",omitempty"`
	Rcvr        *Party40Choice                `xml:"Rcvr,omitempty" json:",omitempty"`
	OrgnlRefs   []OriginalMessage4            `xml:"OrgnlRefs,omitempty" json:",omitempty"`
	Sbjt        common.Max140Text             `xml:"Sbjt"`
	Tp          SupportDocumentType1Code      `xml:"Tp"`
	Desc        *common.Max1025Text           `xml:"Desc,omitempty" json:",omitempty"`
	RspnReqrd   bool                          `xml:"RspnReqrd"`
	DueDt       *common.ISODate               `xml:"DueDt,omitempty" json:",omitempty"`
	Attchmnt    []DocumentGeneralInformation3 `xml:"Attchmnt,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r SupportingDocumentRequestOrLetter2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlGroupStatus2 struct {
	OrgnlRefs OriginalMessage5                             `xml:"OrgnlRefs"`
	RptgPty   TradeParty5                                  `xml:"RptgPty"`
	RegnAgt   BranchAndFinancialInstitutionIdentification6 `xml:"RegnAgt"`
	RptgPrd   *Period4Choice                               `xml:"RptgPrd,omitempty" json:",omitempty"`
	Sts       *StatisticalReportingStatus1Code             `xml:"Sts,omitempty" json:",omitempty"`
	StsRsn    []ValidationStatusReason2                    `xml:"StsRsn,omitempty" json:",omitempty"`
	StsDtTm   *common.ISODateTime                          `xml:"StsDtTm,omitempty" json:",omitempty"`
}

func (r CurrencyControlGroupStatus2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlPackageStatus2 struct {
	PackgId common.Max35Text                `xml:"PackgId"`
	Sts     StatisticalReportingStatus1Code `xml:"Sts"`
	StsRsn  []ValidationStatusReason2       `xml:"StsRsn,omitempty" json:",omitempty"`
	StsDtTm *common.ISODateTime             `xml:"StsDtTm,omitempty" json:",omitempty"`
	RcrdSts []CurrencyControlRecordStatus2  `xml:"RcrdSts,omitempty" json:",omitempty"`
}

func (r CurrencyControlPackageStatus2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlRecordStatus2 struct {
	RcrdId  common.Max35Text                `xml:"RcrdId"`
	Sts     StatisticalReportingStatus1Code `xml:"Sts"`
	StsRsn  []ValidationStatusReason2       `xml:"StsRsn,omitempty" json:",omitempty"`
	StsDtTm *common.ISODateTime             `xml:"StsDtTm,omitempty" json:",omitempty"`
	DocId   *DocumentIdentification28       `xml:"DocId,omitempty" json:",omitempty"`
}

func (r CurrencyControlRecordStatus2) Validate() error {
	return utils.Validate(&r)
}

type CurrencyControlStatusAdviceV02 struct {
	Attr        []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	GrpHdr      CurrencyControlHeader6          `xml:"GrpHdr"`
	GrpSts      []CurrencyControlGroupStatus2   `xml:"GrpSts" json:",omitempty"`
	PackgSts    []CurrencyControlPackageStatus2 `xml:"PackgSts,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CurrencyControlStatusAdviceV02) Validate() error {
	return utils.Validate(&r)
}

type OriginalMessage5 struct {
	OrgnlSndr    *Party40Choice      `xml:"OrgnlSndr,omitempty" json:",omitempty"`
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalMessage5) Validate() error {
	return utils.Validate(&r)
}

type Period2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r Period2) Validate() error {
	return utils.Validate(&r)
}

type Period4Choice struct {
	Dt       common.ISODate `xml:"Dt"`
	FrDt     common.ISODate `xml:"FrDt"`
	ToDt     common.ISODate `xml:"ToDt"`
	FrDtToDt Period2        `xml:"FrDtToDt"`
}

func (r Period4Choice) Validate() error {
	return utils.Validate(&r)
}

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r StatusReason6Choice) Validate() error {
	return utils.Validate(&r)
}

type ValidationStatusReason2 struct {
	Orgtr     *PartyIdentification135                `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn       *StatusReason6Choice                   `xml:"Rsn,omitempty" json:",omitempty"`
	VldtnRule []GenericValidationRuleIdentification1 `xml:"VldtnRule,omitempty" json:",omitempty"`
	AddtlInf  []common.Max105Text                    `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r ValidationStatusReason2) Validate() error {
	return utils.Validate(&r)
}
