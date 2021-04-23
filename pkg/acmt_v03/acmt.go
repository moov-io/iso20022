// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type AccountContract2 struct {
	TrgtGoLiveDt *common.ISODate `xml:"TrgtGoLiveDt,omitempty" json:",omitempty"`
	TrgtClsgDt   *common.ISODate `xml:"TrgtClsgDt,omitempty" json:",omitempty"`
	UrgcyFlg     bool            `xml:"UrgcyFlg,omitempty" json:",omitempty"`
}

func (r AccountContract2) Validate() error {
	return utils.Validate(&r)
}

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

func (r AccountIdentification4Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountOpeningRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References4                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	Acct             CustomerAccount4                             `xml:"Acct"`
	CtrctDts         *AccountContract2                            `xml:"CtrctDts,omitempty" json:",omitempty"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	Org              Organisation33                               `xml:"Org"`
	Mndt             []OperationMandate4                          `xml:"Mndt,omitempty" json:",omitempty"`
	Grp              []Group4                                     `xml:"Grp,omitempty" json:",omitempty"`
	RefAcct          *CashAccount38                               `xml:"RefAcct,omitempty" json:",omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountOpeningRequestV03) Validate() error {
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

type Authorisation2 struct {
	MaxAmtByTx          *FixedAmountOrUnlimited1Choice `xml:"MaxAmtByTx,omitempty" json:",omitempty"`
	MaxAmtByPrd         []MaximumAmountByPeriod1       `xml:"MaxAmtByPrd,omitempty" json:",omitempty"`
	MaxAmtByBlkSubmissn *FixedAmountOrUnlimited1Choice `xml:"MaxAmtByBlkSubmissn,omitempty" json:",omitempty"`
}

func (r Authorisation2) Validate() error {
	return utils.Validate(&r)
}

type BankTransactionCodeStructure4 struct {
	Domn  *BankTransactionCodeStructure5            `xml:"Domn,omitempty" json:",omitempty"`
	Prtry *ProprietaryBankTransactionCodeStructure1 `xml:"Prtry,omitempty" json:",omitempty"`
}

func (r BankTransactionCodeStructure4) Validate() error {
	return utils.Validate(&r)
}

type BankTransactionCodeStructure5 struct {
	Cd   ExternalBankTransactionDomain1Code `xml:"Cd"`
	Fmly BankTransactionCodeStructure6      `xml:"Fmly"`
}

func (r BankTransactionCodeStructure5) Validate() error {
	return utils.Validate(&r)
}

type BankTransactionCodeStructure6 struct {
	Cd        ExternalBankTransactionFamily1Code    `xml:"Cd"`
	SubFmlyCd ExternalBankTransactionSubFamily1Code `xml:"SubFmlyCd"`
}

func (r BankTransactionCodeStructure6) Validate() error {
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

type CashAccount38 struct {
	Id   *AccountIdentification4Choice        `xml:"Id"`
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

type Channel2Choice struct {
	Cd    CommunicationMethod3Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r Channel2Choice) Validate() error {
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

type CodeOrProprietary1Choice struct {
	Cd    common.Max4Text         `xml:"Cd"`
	Prtry GenericIdentification13 `xml:"Prtry"`
}

func (r CodeOrProprietary1Choice) Validate() error {
	return utils.Validate(&r)
}

type CommunicationFormat1Choice struct {
	Cd    ExternalCommunicationFormat1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

func (r CommunicationFormat1Choice) Validate() error {
	return utils.Validate(&r)
}

type CommunicationMethod2Choice struct {
	Cd    CommunicationMethod2Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

func (r CommunicationMethod2Choice) Validate() error {
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

type ContractDocument1 struct {
	Ref      common.Max35Text `xml:"Ref"`
	SgnOffDt *common.ISODate  `xml:"SgnOffDt,omitempty" json:",omitempty"`
	Vrsn     *common.Max6Text `xml:"Vrsn,omitempty" json:",omitempty"`
}

func (r ContractDocument1) Validate() error {
	return utils.Validate(&r)
}

type CustomerAccount4 struct {
	Id               *AccountIdentification4Choice `xml:"Id,omitempty" json:",omitempty"`
	Nm               *common.Max70Text             `xml:"Nm,omitempty" json:",omitempty"`
	Sts              *common.AccountStatus3Code    `xml:"Sts,omitempty" json:",omitempty"`
	Tp               *CashAccountType2Choice       `xml:"Tp,omitempty" json:",omitempty"`
	Ccy              common.ActiveCurrencyCode     `xml:"Ccy"`
	MnthlyPmtVal     float64                       `xml:"MnthlyPmtVal,omitempty" json:",omitempty"`
	MnthlyRcvdVal    float64                       `xml:"MnthlyRcvdVal,omitempty" json:",omitempty"`
	MnthlyTxNb       *common.Max5NumericText       `xml:"MnthlyTxNb,omitempty" json:",omitempty"`
	AvrgBal          float64                       `xml:"AvrgBal,omitempty" json:",omitempty"`
	AcctPurp         *common.Max140Text            `xml:"AcctPurp,omitempty" json:",omitempty"`
	FlrNtfctnAmt     float64                       `xml:"FlrNtfctnAmt,omitempty" json:",omitempty"`
	ClngNtfctnAmt    float64                       `xml:"ClngNtfctnAmt,omitempty" json:",omitempty"`
	StmtFrqcyAndFrmt []StatementFrequencyAndForm1  `xml:"StmtFrqcyAndFrmt,omitempty" json:",omitempty"`
	ClsgDt           *common.ISODate               `xml:"ClsgDt,omitempty" json:",omitempty"`
	Rstrctn          []Restriction1                `xml:"Rstrctn,omitempty" json:",omitempty"`
}

func (r CustomerAccount4) Validate() error {
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

type FixedAmountOrUnlimited1Choice struct {
	Amt    ActiveCurrencyAndAmount `xml:"Amt"`
	NotLtd Unlimited9Text          `xml:"NotLtd"`
}

func (r FixedAmountOrUnlimited1Choice) Validate() error {
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

type GenericIdentification13 struct {
	Id      common.Max4AlphaNumericText `xml:"Id"`
	SchmeNm *common.Max35Text           `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text            `xml:"Issr"`
}

func (r GenericIdentification13) Validate() error {
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

type Group4 struct {
	GrpId common.Max4AlphaNumericText `xml:"GrpId"`
	Pty   []PartyAndCertificate4      `xml:"Pty" json:",omitempty"`
}

func (r Group4) Validate() error {
	return utils.Validate(&r)
}

type MaximumAmountByPeriod1 struct {
	MaxAmt   ActiveCurrencyAndAmount `xml:"MaxAmt"`
	NbOfDays common.Max3NumericText  `xml:"NbOfDays"`
}

func (r MaximumAmountByPeriod1) Validate() error {
	return utils.Validate(&r)
}

type MessageIdentification1 struct {
	Id      common.Max35Text   `xml:"Id"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r MessageIdentification1) Validate() error {
	return utils.Validate(&r)
}

type OperationMandate4 struct {
	Id           common.Max35Text                  `xml:"Id"`
	AplblChanl   []Channel2Choice                  `xml:"AplblChanl" json:",omitempty"`
	ReqrdSgntrNb common.Max15PlusSignedNumericText `xml:"ReqrdSgntrNb"`
	SgntrOrdrInd bool                              `xml:"SgntrOrdrInd"`
	MndtHldr     []PartyAndAuthorisation4          `xml:"MndtHldr,omitempty" json:",omitempty"`
	BkOpr        []BankTransactionCodeStructure4   `xml:"BkOpr" json:",omitempty"`
	StartDt      *common.ISODate                   `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt        *common.ISODate                   `xml:"EndDt,omitempty" json:",omitempty"`
}

func (r OperationMandate4) Validate() error {
	return utils.Validate(&r)
}

type Organisation33 struct {
	FullLglNm    common.Max350Text            `xml:"FullLglNm"`
	TradgNm      *common.Max350Text           `xml:"TradgNm,omitempty" json:",omitempty"`
	CtryOfOpr    common.CountryCode           `xml:"CtryOfOpr"`
	RegnDt       *common.ISODate              `xml:"RegnDt,omitempty" json:",omitempty"`
	OprlAdr      *PostalAddress24             `xml:"OprlAdr,omitempty" json:",omitempty"`
	BizAdr       *PostalAddress24             `xml:"BizAdr,omitempty" json:",omitempty"`
	LglAdr       PostalAddress24              `xml:"LglAdr"`
	BllgAdr      *PostalAddress24             `xml:"BllgAdr,omitempty" json:",omitempty"`
	OrgId        OrganisationIdentification29 `xml:"OrgId"`
	RprtvOffcr   []PartyIdentification137     `xml:"RprtvOffcr,omitempty" json:",omitempty"`
	TrsrMgr      *PartyIdentification137      `xml:"TrsrMgr,omitempty" json:",omitempty"`
	MainMndtHldr []PartyIdentification137     `xml:"MainMndtHldr,omitempty" json:",omitempty"`
	Sndr         []PartyIdentification137     `xml:"Sndr,omitempty" json:",omitempty"`
	LglRprtv     []PartyIdentification137     `xml:"LglRprtv,omitempty" json:",omitempty"`
}

func (r Organisation33) Validate() error {
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

type PartyAndAuthorisation4 struct {
	PtyOrGrp  PartyOrGroup2Choice                `xml:"PtyOrGrp"`
	SgntrOrdr *common.Max15PlusSignedNumericText `xml:"SgntrOrdr,omitempty" json:",omitempty"`
	Authstn   Authorisation2                     `xml:"Authstn"`
}

func (r PartyAndAuthorisation4) Validate() error {
	return utils.Validate(&r)
}

type PartyAndCertificate4 struct {
	Pty  PartyIdentification135 `xml:"Pty"`
	Cert *common.Max10KBinary   `xml:"Cert,omitempty" json:",omitempty"`
}

func (r PartyAndCertificate4) Validate() error {
	return utils.Validate(&r)
}

type PartyAndSignature3 struct {
	Pty   PartyIdentification135 `xml:"Pty"`
	Sgntr SkipPayload            `xml:"Sgntr"`
}

func (r PartyAndSignature3) Validate() error {
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

type PartyIdentification137 struct {
	Nm        *common.Max140Text      `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24        `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *PersonIdentification13 `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode     `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4               `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r PartyIdentification137) Validate() error {
	return utils.Validate(&r)
}

type PartyOrGroup2Choice struct {
	GrpId common.Max4AlphaNumericText `xml:"GrpId"`
	Pty   PartyAndCertificate4        `xml:"Pty"`
}

func (r PartyOrGroup2Choice) Validate() error {
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

type ProprietaryBankTransactionCodeStructure1 struct {
	Cd   common.Max35Text  `xml:"Cd"`
	Issr *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r ProprietaryBankTransactionCodeStructure1) Validate() error {
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

type References4 struct {
	MsgId       MessageIdentification1 `xml:"MsgId"`
	PrcId       MessageIdentification1 `xml:"PrcId"`
	AttchdDocNm []common.Max70Text     `xml:"AttchdDocNm,omitempty" json:",omitempty"`
}

func (r References4) Validate() error {
	return utils.Validate(&r)
}

type Restriction1 struct {
	RstrctnTp CodeOrProprietary1Choice `xml:"RstrctnTp"`
	VldFr     common.ISODateTime       `xml:"VldFr"`
	VldUntil  *common.ISODateTime      `xml:"VldUntil,omitempty" json:",omitempty"`
}

func (r Restriction1) Validate() error {
	return utils.Validate(&r)
}

type SkipPayload struct {
	Item string `xml:",any"`
}

func (r SkipPayload) Validate() error {
	return utils.Validate(&r)
}

type StatementFrequencyAndForm1 struct {
	Frqcy    Frequency7Code             `xml:"Frqcy"`
	ComMtd   CommunicationMethod2Choice `xml:"ComMtd"`
	DlvryAdr common.Max350Text          `xml:"DlvryAdr"`
	Frmt     CommunicationFormat1Choice `xml:"Frmt"`
}

func (r StatementFrequencyAndForm1) Validate() error {
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

type AccountOpeningAmendmentRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References4                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	CtrctDts         *AccountContract2                            `xml:"CtrctDts,omitempty" json:",omitempty"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	Acct             CustomerAccount4                             `xml:"Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	Org              Organisation33                               `xml:"Org"`
	Mndt             []OperationMandate4                          `xml:"Mndt,omitempty" json:",omitempty"`
	Grp              []Group4                                     `xml:"Grp,omitempty" json:",omitempty"`
	RefAcct          *CashAccount38                               `xml:"RefAcct,omitempty" json:",omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountOpeningAmendmentRequestV03) Validate() error {
	return utils.Validate(&r)
}

type References3 struct {
	MsgId           MessageIdentification1 `xml:"MsgId"`
	ReqToBeCmpltdId MessageIdentification1 `xml:"ReqToBeCmpltdId"`
	PrcId           MessageIdentification1 `xml:"PrcId"`
	ReqRsn          []common.Max35Text     `xml:"ReqRsn" json:",omitempty"`
	AttchdDocNm     []common.Max70Text     `xml:"AttchdDocNm,omitempty" json:",omitempty"`
}

func (r References3) Validate() error {
	return utils.Validate(&r)
}

type AccountOpeningAdditionalInformationRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References3                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	OrgId            OrganisationIdentification29                 `xml:"OrgId"`
	Acct             CustomerAccount4                             `xml:"Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountOpeningAdditionalInformationRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountRequestAcknowledgementV03 struct {
	XMLName     *xml.Name                                    `json:",omitempty"`
	Attr        *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs        References5                                  `xml:"Refs"`
	Fr          *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	AcctId      []AccountForAction1                          `xml:"AcctId,omitempty" json:",omitempty"`
	OrgId       OrganisationIdentification29                 `xml:"OrgId"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountRequestAcknowledgementV03) Validate() error {
	return utils.Validate(&r)
}

type AccountRequestRejectionV03 struct {
	XMLName     *xml.Name                                    `json:",omitempty"`
	Attr        *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs        References6                                  `xml:"Refs"`
	Fr          *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	AcctId      []AccountForAction1                          `xml:"AcctId,omitempty" json:",omitempty"`
	OrgId       OrganisationIdentification29                 `xml:"OrgId"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountRequestRejectionV03) Validate() error {
	return utils.Validate(&r)
}

type References6 struct {
	RjctdReqTp  UseCases1Code          `xml:"RjctdReqTp"`
	RjctnRsn    []common.Max350Text    `xml:"RjctnRsn" json:",omitempty"`
	RjctdReqId  MessageIdentification1 `xml:"RjctdReqId"`
	MsgId       MessageIdentification1 `xml:"MsgId"`
	PrcId       MessageIdentification1 `xml:"PrcId"`
	AttchdDocNm []common.Max70Text     `xml:"AttchdDocNm,omitempty" json:",omitempty"`
}

func (r References6) Validate() error {
	return utils.Validate(&r)
}

type AccountAdditionalInformationRequestV03 struct {
	XMLName     *xml.Name                                    `json:",omitempty"`
	Attr        *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs        References3                                  `xml:"Refs"`
	Fr          *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	OrgId       OrganisationIdentification29                 `xml:"OrgId"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	AcctId      []AccountForAction1                          `xml:"AcctId" json:",omitempty"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountAdditionalInformationRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountReportRequestV03 struct {
	XMLName     *xml.Name                                    `json:",omitempty"`
	Attr        *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs        References4                                  `xml:"Refs"`
	Fr          *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	AcctId      []AccountForAction1                          `xml:"AcctId" json:",omitempty"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	OrgId       OrganisationIdentification29                 `xml:"OrgId"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountReportRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountContract3 struct {
	TrgtGoLiveDt *common.ISODate `xml:"TrgtGoLiveDt,omitempty" json:",omitempty"`
	TrgtClsgDt   *common.ISODate `xml:"TrgtClsgDt,omitempty" json:",omitempty"`
	GoLiveDt     *common.ISODate `xml:"GoLiveDt,omitempty" json:",omitempty"`
	ClsgDt       *common.ISODate `xml:"ClsgDt,omitempty" json:",omitempty"`
	UrgcyFlg     bool            `xml:"UrgcyFlg,omitempty" json:",omitempty"`
	RmvlInd      bool            `xml:"RmvlInd,omitempty" json:",omitempty"`
}

func (r AccountContract3) Validate() error {
	return utils.Validate(&r)
}

type AccountReport23 struct {
	Acct             CustomerAccount5                              `xml:"Acct"`
	UndrlygMstrAgrmt *ContractDocument1                            `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	CtrctDts         *AccountContract3                             `xml:"CtrctDts,omitempty" json:",omitempty"`
	Mndt             []OperationMandate4                           `xml:"Mndt,omitempty" json:",omitempty"`
	Grp              []Group4                                      `xml:"Grp,omitempty" json:",omitempty"`
	RefAcct          *CashAccount38                                `xml:"RefAcct,omitempty" json:",omitempty"`
	BalTrfAcct       *AccountForAction1                            `xml:"BalTrfAcct,omitempty" json:",omitempty"`
	TrfAcctSvcrId    *BranchAndFinancialInstitutionIdentification6 `xml:"TrfAcctSvcrId,omitempty" json:",omitempty"`
}

func (r AccountReport23) Validate() error {
	return utils.Validate(&r)
}

type AccountReportV03 struct {
	XMLName     *xml.Name                                    `json:",omitempty"`
	Attr        *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs        References5                                  `xml:"Refs"`
	Fr          *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	AcctSvcrId  BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	Org         Organisation33                               `xml:"Org"`
	Rpt         []AccountReport23                            `xml:"Rpt,omitempty" json:",omitempty"`
	DgtlSgntr   []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountReportV03) Validate() error {
	return utils.Validate(&r)
}

type CustomerAccount5 struct {
	Id               []AccountIdentification4Choice `xml:"Id" json:",omitempty"`
	Nm               *common.Max70Text              `xml:"Nm,omitempty" json:",omitempty"`
	Sts              *common.AccountStatus3Code     `xml:"Sts,omitempty" json:",omitempty"`
	Tp               *CashAccountType2Choice        `xml:"Tp,omitempty" json:",omitempty"`
	Ccy              common.ActiveCurrencyCode      `xml:"Ccy"`
	MnthlyPmtVal     float64                        `xml:"MnthlyPmtVal,omitempty" json:",omitempty"`
	MnthlyRcvdVal    float64                        `xml:"MnthlyRcvdVal,omitempty" json:",omitempty"`
	MnthlyTxNb       *common.Max5NumericText        `xml:"MnthlyTxNb,omitempty" json:",omitempty"`
	AvrgBal          float64                        `xml:"AvrgBal,omitempty" json:",omitempty"`
	AcctPurp         *common.Max140Text             `xml:"AcctPurp,omitempty" json:",omitempty"`
	FlrNtfctnAmt     float64                        `xml:"FlrNtfctnAmt,omitempty" json:",omitempty"`
	ClngNtfctnAmt    float64                        `xml:"ClngNtfctnAmt,omitempty" json:",omitempty"`
	StmtFrqcyAndFrmt []StatementFrequencyAndForm1   `xml:"StmtFrqcyAndFrmt,omitempty" json:",omitempty"`
	ClsgDt           *common.ISODate                `xml:"ClsgDt,omitempty" json:",omitempty"`
	Rstrctn          []Restriction1                 `xml:"Rstrctn,omitempty" json:",omitempty"`
}

func (r CustomerAccount5) Validate() error {
	return utils.Validate(&r)
}

type AccountExcludedMandateMaintenanceRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References4                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	CtrctDts         *AccountContract2                            `xml:"CtrctDts,omitempty" json:",omitempty"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	Acct             CustomerAccountModification1                 `xml:"Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	Org              OrganisationModification2                    `xml:"Org"`
	AddtlMsgInf      *AdditionalInformation5                      `xml:"AddtlMsgInf,omitempty" json:",omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountExcludedMandateMaintenanceRequestV03) Validate() error {
	return utils.Validate(&r)
}

type References5 struct {
	ReqTp       UseCases1Code            `xml:"ReqTp"`
	MsgId       MessageIdentification1   `xml:"MsgId"`
	PrcId       MessageIdentification1   `xml:"PrcId"`
	AckdMsgId   []MessageIdentification1 `xml:"AckdMsgId,omitempty" json:",omitempty"`
	Sts         *common.Max35Text        `xml:"Sts,omitempty" json:",omitempty"`
	AttchdDocNm []common.Max70Text       `xml:"AttchdDocNm,omitempty" json:",omitempty"`
}

func (r References5) Validate() error {
	return utils.Validate(&r)
}

type AccountStatusModification1 struct {
	ModCd *Modification1Code        `xml:"ModCd,omitempty" json:",omitempty"`
	Sts   common.AccountStatus3Code `xml:"Sts"`
}

func (r AccountStatusModification1) Validate() error {
	return utils.Validate(&r)
}

type AdditionalInformation5 struct {
	Inf []common.Max256Text `xml:"Inf" json:",omitempty"`
}

func (r AdditionalInformation5) Validate() error {
	return utils.Validate(&r)
}

type AddressModification2 struct {
	ModCd *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	Adr   PostalAddress24    `xml:"Adr"`
}

func (r AddressModification2) Validate() error {
	return utils.Validate(&r)
}

type AmountModification1 struct {
	ModCd *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	Amt   float64            `xml:"Amt"`
}

func (r AmountModification1) Validate() error {
	return utils.Validate(&r)
}

type DateModification1 struct {
	ModCd *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	Dt    common.ISODate     `xml:"Dt"`
}

func (r DateModification1) Validate() error {
	return utils.Validate(&r)
}

type FullLegalNameModification1 struct {
	ModCd     *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	FullLglNm common.Max350Text  `xml:"FullLglNm"`
}

func (r FullLegalNameModification1) Validate() error {
	return utils.Validate(&r)
}

type NameModification1 struct {
	ModCd *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	Nm    common.Max70Text   `xml:"Nm"`
}

func (r NameModification1) Validate() error {
	return utils.Validate(&r)
}

type NumberModification1 struct {
	ModCd *Modification1Code     `xml:"ModCd,omitempty" json:",omitempty"`
	Nb    common.Max5NumericText `xml:"Nb"`
}

func (r NumberModification1) Validate() error {
	return utils.Validate(&r)
}

type OrganisationModification2 struct {
	FullLglNm    FullLegalNameModification1   `xml:"FullLglNm"`
	TradgNm      *TradingNameModification1    `xml:"TradgNm,omitempty" json:",omitempty"`
	CtryOfOpr    common.CountryCode           `xml:"CtryOfOpr"`
	RegnDt       *common.ISODate              `xml:"RegnDt,omitempty" json:",omitempty"`
	OprlAdr      *AddressModification2        `xml:"OprlAdr,omitempty" json:",omitempty"`
	BizAdr       *AddressModification2        `xml:"BizAdr,omitempty" json:",omitempty"`
	LglAdr       AddressModification2         `xml:"LglAdr"`
	BllgAdr      *AddressModification2        `xml:"BllgAdr,omitempty" json:",omitempty"`
	OrgId        OrganisationIdentification29 `xml:"OrgId"`
	RprtvOffcr   []PartyModification2         `xml:"RprtvOffcr,omitempty" json:",omitempty"`
	TrsrMgr      *PartyModification2          `xml:"TrsrMgr,omitempty" json:",omitempty"`
	MainMndtHldr []PartyModification2         `xml:"MainMndtHldr,omitempty" json:",omitempty"`
	Sndr         []PartyModification2         `xml:"Sndr,omitempty" json:",omitempty"`
	LglRprtv     []PartyModification2         `xml:"LglRprtv,omitempty" json:",omitempty"`
}

func (r OrganisationModification2) Validate() error {
	return utils.Validate(&r)
}

type PartyModification2 struct {
	ModCd *Modification1Code     `xml:"ModCd,omitempty" json:",omitempty"`
	PtyId PartyIdentification137 `xml:"PtyId"`
}

func (r PartyModification2) Validate() error {
	return utils.Validate(&r)
}

type PurposeModification1 struct {
	ModCd *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	Purp  common.Max140Text  `xml:"Purp"`
}

func (r PurposeModification1) Validate() error {
	return utils.Validate(&r)
}

type RestrictionModification1 struct {
	ModCd   *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	Rstrctn Restriction1       `xml:"Rstrctn"`
}

func (r RestrictionModification1) Validate() error {
	return utils.Validate(&r)
}

type StatementFrequencyAndFormModification1 struct {
	ModCd            *Modification1Code         `xml:"ModCd,omitempty" json:",omitempty"`
	StmtFrqcyAndForm StatementFrequencyAndForm1 `xml:"StmtFrqcyAndForm"`
}

func (r StatementFrequencyAndFormModification1) Validate() error {
	return utils.Validate(&r)
}

type TradingNameModification1 struct {
	ModCd   *Modification1Code `xml:"ModCd,omitempty" json:",omitempty"`
	TradgNm common.Max350Text  `xml:"TradgNm"`
}

func (r TradingNameModification1) Validate() error {
	return utils.Validate(&r)
}

type TypeModification1 struct {
	ModCd *Modification1Code     `xml:"ModCd,omitempty" json:",omitempty"`
	Tp    CashAccountType2Choice `xml:"Tp"`
}

func (r TypeModification1) Validate() error {
	return utils.Validate(&r)
}

type AccountExcludedMandateMaintenanceAmendmentRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References4                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	CtrctDts         *AccountContract2                            `xml:"CtrctDts,omitempty" json:",omitempty"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	Acct             CustomerAccountModification1                 `xml:"Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	Org              OrganisationModification2                    `xml:"Org"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountExcludedMandateMaintenanceAmendmentRequestV03) Validate() error {
	return utils.Validate(&r)
}

type CustomerAccountModification1 struct {
	Id               []AccountIdentification4Choice           `xml:"Id" json:",omitempty"`
	Nm               *NameModification1                       `xml:"Nm,omitempty" json:",omitempty"`
	Sts              *AccountStatusModification1              `xml:"Sts,omitempty" json:",omitempty"`
	Tp               *TypeModification1                       `xml:"Tp,omitempty" json:",omitempty"`
	Ccy              common.ActiveCurrencyCode                `xml:"Ccy"`
	MnthlyPmtVal     *AmountModification1                     `xml:"MnthlyPmtVal,omitempty" json:",omitempty"`
	MnthlyRcvdVal    *AmountModification1                     `xml:"MnthlyRcvdVal,omitempty" json:",omitempty"`
	MnthlyTxNb       *NumberModification1                     `xml:"MnthlyTxNb,omitempty" json:",omitempty"`
	AvrgBal          *AmountModification1                     `xml:"AvrgBal,omitempty" json:",omitempty"`
	AcctPurp         *PurposeModification1                    `xml:"AcctPurp,omitempty" json:",omitempty"`
	FlrNtfctnAmt     *AmountModification1                     `xml:"FlrNtfctnAmt,omitempty" json:",omitempty"`
	ClngNtfctnAmt    *AmountModification1                     `xml:"ClngNtfctnAmt,omitempty" json:",omitempty"`
	StmtFrqcyAndFrmt []StatementFrequencyAndFormModification1 `xml:"StmtFrqcyAndFrmt,omitempty" json:",omitempty"`
	ClsgDt           *DateModification1                       `xml:"ClsgDt,omitempty" json:",omitempty"`
	Rstrctn          []RestrictionModification1               `xml:"Rstrctn,omitempty" json:",omitempty"`
}

func (r CustomerAccountModification1) Validate() error {
	return utils.Validate(&r)
}

type AccountMandateMaintenanceRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References4                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	CtrctDts         *AccountContract2                            `xml:"CtrctDts,omitempty" json:",omitempty"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	AcctId           []AccountForAction1                          `xml:"AcctId" json:",omitempty"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	OrgId            Organisation34                               `xml:"OrgId"`
	Mndt             []OperationMandate5                          `xml:"Mndt" json:",omitempty"`
	Grp              []Group3                                     `xml:"Grp,omitempty" json:",omitempty"`
	AddtlMsgInf      *AdditionalInformation5                      `xml:"AddtlMsgInf,omitempty" json:",omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountMandateMaintenanceRequestV03) Validate() error {
	return utils.Validate(&r)
}

type Group3 struct {
	ModCd *Modification1Code          `xml:"ModCd,omitempty" json:",omitempty"`
	GrpId common.Max4AlphaNumericText `xml:"GrpId"`
	Pty   []PartyAndCertificate5      `xml:"Pty" json:",omitempty"`
}

func (r Group3) Validate() error {
	return utils.Validate(&r)
}

type OperationMandate5 struct {
	ModCd        *Modification1Code                `xml:"ModCd,omitempty" json:",omitempty"`
	Id           common.Max35Text                  `xml:"Id"`
	AplblChanl   []Channel2Choice                  `xml:"AplblChanl" json:",omitempty"`
	ReqrdSgntrNb common.Max15PlusSignedNumericText `xml:"ReqrdSgntrNb"`
	SgntrOrdrInd bool                              `xml:"SgntrOrdrInd"`
	MndtHldr     []PartyAndAuthorisation5          `xml:"MndtHldr,omitempty" json:",omitempty"`
	BkOpr        []BankTransactionCodeStructure4   `xml:"BkOpr" json:",omitempty"`
	StartDt      *common.ISODate                   `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt        *common.ISODate                   `xml:"EndDt,omitempty" json:",omitempty"`
}

func (r OperationMandate5) Validate() error {
	return utils.Validate(&r)
}

type Organisation34 struct {
	FullLglNm *common.Max350Text           `xml:"FullLglNm,omitempty" json:",omitempty"`
	OrgId     OrganisationIdentification29 `xml:"OrgId"`
}

func (r Organisation34) Validate() error {
	return utils.Validate(&r)
}

type PartyAndAuthorisation5 struct {
	ModCd     *Modification1Code                 `xml:"ModCd,omitempty" json:",omitempty"`
	PtyOrGrp  PartyOrGroup2Choice                `xml:"PtyOrGrp"`
	SgntrOrdr *common.Max15PlusSignedNumericText `xml:"SgntrOrdr,omitempty" json:",omitempty"`
	Authstn   Authorisation2                     `xml:"Authstn"`
}

func (r PartyAndAuthorisation5) Validate() error {
	return utils.Validate(&r)
}

type PartyAndCertificate5 struct {
	ModCd *Modification1Code     `xml:"ModCd,omitempty" json:",omitempty"`
	Pty   PartyIdentification135 `xml:"Pty"`
	Cert  *common.Max10KBinary   `xml:"Cert,omitempty" json:",omitempty"`
}

func (r PartyAndCertificate5) Validate() error {
	return utils.Validate(&r)
}

type AccountMandateMaintenanceAmendmentRequestV03 struct {
	XMLName          *xml.Name                                    `json:",omitempty"`
	Attr             *utils.Attr                                  `xml:",attr,omitempty" json:",omitempty"`
	Refs             References4                                  `xml:"Refs"`
	Fr               *OrganisationIdentification29                `xml:"Fr,omitempty" json:",omitempty"`
	CtrctDts         *AccountContract2                            `xml:"CtrctDts,omitempty" json:",omitempty"`
	UndrlygMstrAgrmt *ContractDocument1                           `xml:"UndrlygMstrAgrmt,omitempty" json:",omitempty"`
	AcctId           []AccountForAction1                          `xml:"AcctId" json:",omitempty"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcrId"`
	OrgId            OrganisationIdentification29                 `xml:"OrgId" json:",omitempty"`
	Mndt             []OperationMandate5                          `xml:"Mndt,omitempty" json:",omitempty"`
	Grp              []Group3                                     `xml:"Grp,omitempty" json:",omitempty"`
	DgtlSgntr        []PartyAndSignature3                         `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData      []SupplementaryData1                         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountMandateMaintenanceAmendmentRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountClosingRequestV03 struct {
	XMLName       *xml.Name                                     `json:",omitempty"`
	Attr          *utils.Attr                                   `xml:",attr,omitempty" json:",omitempty"`
	Refs          References4                                   `xml:"Refs"`
	Fr            *OrganisationIdentification29                 `xml:"Fr,omitempty" json:",omitempty"`
	AcctId        AccountForAction2                             `xml:"AcctId"`
	AcctSvcrId    BranchAndFinancialInstitutionIdentification6  `xml:"AcctSvcrId"`
	OrgId         Organisation34                                `xml:"OrgId"`
	CtrctDts      *AccountContract4                             `xml:"CtrctDts,omitempty" json:",omitempty"`
	BalTrfAcct    *AccountForAction1                            `xml:"BalTrfAcct,omitempty" json:",omitempty"`
	TrfAcctSvcrId *BranchAndFinancialInstitutionIdentification6 `xml:"TrfAcctSvcrId,omitempty" json:",omitempty"`
	DgtlSgntr     []PartyAndSignature3                          `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountClosingRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountForAction2 struct {
	Id  AccountIdentification4Choice `xml:"Id"`
	Nm  *common.Max70Text            `xml:"Nm,omitempty" json:",omitempty"`
	Ccy common.ActiveCurrencyCode    `xml:"Ccy"`
}

func (r AccountForAction2) Validate() error {
	return utils.Validate(&r)
}

type AccountForAction1 struct {
	Id  AccountIdentification4Choice `xml:"Id"`
	Ccy *common.ActiveCurrencyCode   `xml:"Ccy,omitempty" json:",omitempty"`
}

func (r AccountForAction1) Validate() error {
	return utils.Validate(&r)
}

type AccountClosingAmendmentRequestV03 struct {
	XMLName       *xml.Name                                     `json:",omitempty"`
	Attr          *utils.Attr                                   `xml:",attr,omitempty" json:",omitempty"`
	Refs          References4                                   `xml:"Refs"`
	Fr            *OrganisationIdentification29                 `xml:"Fr,omitempty" json:",omitempty"`
	AcctId        AccountForAction1                             `xml:"AcctId"`
	AcctSvcrId    BranchAndFinancialInstitutionIdentification6  `xml:"AcctSvcrId"`
	OrgId         OrganisationIdentification29                  `xml:"OrgId"`
	CtrctDts      *AccountContract4                             `xml:"CtrctDts,omitempty" json:",omitempty"`
	BalTrfAcct    *AccountForAction1                            `xml:"BalTrfAcct,omitempty" json:",omitempty"`
	TrfAcctSvcrId *BranchAndFinancialInstitutionIdentification6 `xml:"TrfAcctSvcrId,omitempty" json:",omitempty"`
	DgtlSgntr     []PartyAndSignature3                          `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountClosingAmendmentRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountContract4 struct {
	TrgtClsgDt *common.ISODate `xml:"TrgtClsgDt,omitempty" json:",omitempty"`
	UrgcyFlg   bool            `xml:"UrgcyFlg,omitempty" json:",omitempty"`
	RmvlInd    bool            `xml:"RmvlInd,omitempty" json:",omitempty"`
}

func (r AccountContract4) Validate() error {
	return utils.Validate(&r)
}

type AccountClosingAdditionalInformationRequestV03 struct {
	XMLName       *xml.Name                                     `json:",omitempty"`
	Attr          *utils.Attr                                   `xml:",attr,omitempty" json:",omitempty"`
	Refs          References3                                   `xml:"Refs"`
	Fr            *OrganisationIdentification29                 `xml:"Fr,omitempty" json:",omitempty"`
	OrgId         OrganisationIdentification29                  `xml:"OrgId"`
	AcctId        AccountForAction1                             `xml:"AcctId"`
	AcctSvcrId    BranchAndFinancialInstitutionIdentification6  `xml:"AcctSvcrId"`
	BalTrfAcct    *AccountForAction1                            `xml:"BalTrfAcct,omitempty" json:",omitempty"`
	TrfAcctSvcrId *BranchAndFinancialInstitutionIdentification6 `xml:"TrfAcctSvcrId,omitempty" json:",omitempty"`
	DgtlSgntr     []PartyAndSignature3                          `xml:"DgtlSgntr,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1                          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountClosingAdditionalInformationRequestV03) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchInformationRequestV03 struct {
	XMLName       *xml.Name              `json:",omitempty"`
	Attr          *utils.Attr            `xml:",attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	NewAcct       NewAccount2            `xml:"NewAcct"`
	OdAcct        CashAccount39          `xml:"OdAcct"`
	BalTrf        []BalanceTransfer3     `xml:"BalTrf,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchInformationRequestV03) Validate() error {
	return utils.Validate(&r)
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
	return utils.Validate(&r)
}

type CitizenshipInformation1 struct {
	Ntlty   string          `xml:"Ntlty"`
	MnrInd  bool            `xml:"MnrInd,omitempty" json:",omitempty"`
	StartDt *common.ISODate `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt   *common.ISODate `xml:"EndDt,omitempty" json:",omitempty"`
}

func (r CitizenshipInformation1) Validate() error {
	return utils.Validate(&r)
}

type CommunicationAddress3 struct {
	Email  *common.Max256Text  `xml:"Email,omitempty" json:",omitempty"`
	Phne   *common.PhoneNumber `xml:"Phne,omitempty" json:",omitempty"`
	Mob    *common.PhoneNumber `xml:"Mob,omitempty" json:",omitempty"`
	FaxNb  *common.PhoneNumber `xml:"FaxNb,omitempty" json:",omitempty"`
	TlxAdr *common.Max35Text   `xml:"TlxAdr,omitempty" json:",omitempty"`
	URLAdr *common.Max256Text  `xml:"URLAdr,omitempty" json:",omitempty"`
}

func (r CommunicationAddress3) Validate() error {
	return utils.Validate(&r)
}

type CountryAndResidentialStatusType1 struct {
	Ctry      common.CountryCode     `xml:"Ctry"`
	ResdtlSts ResidentialStatus1Code `xml:"ResdtlSts"`
}

func (r CountryAndResidentialStatusType1) Validate() error {
	return utils.Validate(&r)
}

type CreditTransferTransaction41 struct {
	PmtId           PaymentIdentification6                        `xml:"PmtId"`
	PmtTpInf        *PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	TaxRateMrkr     *TaxRateMarker1Code                           `xml:"TaxRateMrkr,omitempty" json:",omitempty"`
	Amt             ActiveCurrencyAndAmount                       `xml:"Amt"`
	ChrgBr          *ChargeBearerType1Code                        `xml:"ChrgBr,omitempty" json:",omitempty"`
	ChqInstr        *Cheque11                                     `xml:"ChqInstr,omitempty" json:",omitempty"`
	Frqcy           *Frequency1                                   `xml:"Frqcy,omitempty" json:",omitempty"`
	TrfInstr        *TransferInstruction1                         `xml:"TrfInstr,omitempty" json:",omitempty"`
	UltmtDbtr       *PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:",omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:",omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:",omitempty"`
	CdtrAgt         BranchAndFinancialInstitutionIdentification6  `xml:"CdtrAgt"`
	Cdtr            *PartyIdentification135                       `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct        *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr       *PartyIdentification135                       `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	InstrForCdtrAgt []InstructionForCreditorAgent3                `xml:"InstrForCdtrAgt,omitempty" json:",omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RgltryRptg      []RegulatoryReporting3                        `xml:"RgltryRptg,omitempty" json:",omitempty"`
	Tax             *TaxInformation8                              `xml:"Tax,omitempty" json:",omitempty"`
	RltdRmtInf      []RemittanceLocation6                         `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf          RemittanceInformation16                       `xml:"RmtInf,omitempty" json:",omitempty"`
}

func (r CreditTransferTransaction41) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification44 struct {
	Id     common.Max35Text           `xml:"Id"`
	Tp     OtherIdentification1Choice `xml:"Tp"`
	Issr   *common.Max35Text          `xml:"Issr,omitempty" json:",omitempty"`
	IsseDt *common.ISODate            `xml:"IsseDt,omitempty" json:",omitempty"`
	XpryDt *common.ISODate            `xml:"XpryDt,omitempty" json:",omitempty"`
}

func (r GenericIdentification44) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification47 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max4AlphaNumericText   `xml:"Issr"`
	SchmeNm *common.Max4AlphaNumericText  `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification47) Validate() error {
	return utils.Validate(&r)
}

type IndividualPerson36 struct {
	CurNm            IndividualPersonNameLong2         `xml:"CurNm"`
	PrvsNm           []IndividualPersonNameLong2       `xml:"PrvsNm,omitempty" json:",omitempty"`
	Gndr             *Gender1Code                      `xml:"Gndr,omitempty" json:",omitempty"`
	Lang             string                            `xml:"Lang,omitempty" json:",omitempty"`
	BirthDt          *common.ISODate                   `xml:"BirthDt,omitempty" json:",omitempty"`
	CtryOfBirth      *common.CountryCode               `xml:"CtryOfBirth,omitempty" json:",omitempty"`
	PrvcOfBirth      *common.Max35Text                 `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth      *common.Max35Text                 `xml:"CityOfBirth,omitempty" json:",omitempty"`
	TaxtnCtry        *common.CountryCode               `xml:"TaxtnCtry,omitempty" json:",omitempty"`
	CtryAndResdtlSts *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty" json:",omitempty"`
	SclSctyNb        *common.Max35Text                 `xml:"SclSctyNb,omitempty" json:",omitempty"`
	PstlAdr          []PostalAddress24                 `xml:"PstlAdr,omitempty" json:",omitempty"`
	CtznshInf        []CitizenshipInformation1         `xml:"CtznshInf,omitempty" json:",omitempty"`
	PmryComAdr       *CommunicationAddress3            `xml:"PmryComAdr,omitempty" json:",omitempty"`
	ScndryComAdr     *CommunicationAddress3            `xml:"ScndryComAdr,omitempty" json:",omitempty"`
	OthrId           []GenericIdentification44         `xml:"OthrId,omitempty" json:",omitempty"`
	OthrDtls         []TransferInstruction1            `xml:"OthrDtls,omitempty" json:",omitempty"`
}

func (r IndividualPerson36) Validate() error {
	return utils.Validate(&r)
}

type IndividualPersonNameLong2 struct {
	NmPrfx  *common.NamePrefix2Code `xml:"NmPrfx,omitempty" json:",omitempty"`
	Srnm    common.Max35Text        `xml:"Srnm"`
	GvnNm   *common.Max35Text       `xml:"GvnNm,omitempty" json:",omitempty"`
	MddlNm  *common.Max35Text       `xml:"MddlNm,omitempty" json:",omitempty"`
	Initls  *common.Max6Text        `xml:"Initls,omitempty" json:",omitempty"`
	NmSfx   *common.Max350Text      `xml:"NmSfx,omitempty" json:",omitempty"`
	Nm      *common.Max350Text      `xml:"Nm,omitempty" json:",omitempty"`
	StartDt *common.ISODate         `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt   *common.ISODate         `xml:"EndDt,omitempty" json:",omitempty"`
}

func (r IndividualPersonNameLong2) Validate() error {
	return utils.Validate(&r)
}

type NewAccount2 struct {
	Acct    CashAccount39        `xml:"Acct"`
	AcctPty []IndividualPerson36 `xml:"AcctPty" json:",omitempty"`
	Org     *Organisation35      `xml:"Org,omitempty" json:",omitempty"`
}

func (r NewAccount2) Validate() error {
	return utils.Validate(&r)
}

type Organisation35 struct {
	FullLglNm    common.Max350Text             `xml:"FullLglNm"`
	TradgNm      *common.Max350Text            `xml:"TradgNm,omitempty" json:",omitempty"`
	OrgLglSts    *OrganisationLegalStatus1Code `xml:"OrgLglSts,omitempty" json:",omitempty"`
	EstblishdDt  *common.ISODate               `xml:"EstblishdDt,omitempty" json:",omitempty"`
	RegnNb       *common.Max70Text             `xml:"RegnNb,omitempty" json:",omitempty"`
	RegnCtry     *common.CountryCode           `xml:"RegnCtry,omitempty" json:",omitempty"`
	RegnDt       *common.ISODate               `xml:"RegnDt,omitempty" json:",omitempty"`
	TaxtnIdNb    *common.Max35Text             `xml:"TaxtnIdNb,omitempty" json:",omitempty"`
	TaxtnCtry    *common.CountryCode           `xml:"TaxtnCtry,omitempty" json:",omitempty"`
	CtryOfOpr    *common.CountryCode           `xml:"CtryOfOpr,omitempty" json:",omitempty"`
	BrdRsltnInd  bool                          `xml:"BrdRsltnInd,omitempty" json:",omitempty"`
	BizAdr       *PostalAddress24              `xml:"BizAdr,omitempty" json:",omitempty"`
	OprlAdr      *PostalAddress24              `xml:"OprlAdr,omitempty" json:",omitempty"`
	LglAdr       *PostalAddress24              `xml:"LglAdr,omitempty" json:",omitempty"`
	RprtvOffcr   []PartyIdentification135      `xml:"RprtvOffcr,omitempty" json:",omitempty"`
	TrsrMgr      *PartyIdentification135       `xml:"TrsrMgr,omitempty" json:",omitempty"`
	MainMndtHldr []PartyIdentification135      `xml:"MainMndtHldr,omitempty" json:",omitempty"`
	Sndr         []PartyIdentification135      `xml:"Sndr,omitempty" json:",omitempty"`
}

func (r Organisation35) Validate() error {
	return utils.Validate(&r)
}

type OtherIdentification1Choice struct {
	Cd    PersonIdentificationType5Code `xml:"Cd"`
	Prtry GenericIdentification47       `xml:"Prtry"`
}

func (r OtherIdentification1Choice) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation7 struct {
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
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation7) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	AdmstnZone      *common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             *common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation8) Validate() error {
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

type TaxPeriod2 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecord2 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod2        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount2        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord2) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails2) Validate() error {
	return utils.Validate(&r)
}

type TransferInstruction1 struct {
	TrfInd    bool                `xml:"TrfInd,omitempty" json:",omitempty"`
	Cd        common.Max35Text    `xml:"Cd"`
	Prtry     *common.Max256Text  `xml:"Prtry,omitempty" json:",omitempty"`
	StartDtTm *common.ISODateTime `xml:"StartDtTm,omitempty" json:",omitempty"`
	StartDt   *common.ISODate     `xml:"StartDt,omitempty" json:",omitempty"`
	Desc      *common.Max350Text  `xml:"Desc,omitempty" json:",omitempty"`
}

func (r TransferInstruction1) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchBalanceTransferAcknowledgementV03 struct {
	XMLName       *xml.Name              `json:",omitempty"`
	Attr          *utils.Attr            `xml:",attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	OdAcct        CashAccount39          `xml:"OdAcct"`
	OdAcctBal     AmountAndDirection5    `xml:"OdAcctBal"`
	BalTrf        []BalanceTransfer3     `xml:"BalTrf,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchBalanceTransferAcknowledgementV03) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchDetails1 struct {
	UnqRefNb      common.Max35Text            `xml:"UnqRefNb"`
	RtgUnqRefNb   common.Max35Text            `xml:"RtgUnqRefNb"`
	SwtchRcvdDtTm *common.ISODateTime         `xml:"SwtchRcvdDtTm,omitempty" json:",omitempty"`
	SwtchDt       *common.ISODate             `xml:"SwtchDt,omitempty" json:",omitempty"`
	SwtchTp       SwitchType1Code             `xml:"SwtchTp"`
	SwtchSts      *SwitchStatus1Code          `xml:"SwtchSts,omitempty" json:",omitempty"`
	BalTrfWndw    *BalanceTransferWindow1Code `xml:"BalTrfWndw,omitempty" json:",omitempty"`
	Rspn          []ResponseDetails1          `xml:"Rspn,omitempty" json:",omitempty"`
}

func (r AccountSwitchDetails1) Validate() error {
	return utils.Validate(&r)
}

type AmountAndDirection5 struct {
	Amt    ActiveCurrencyAndAmount `xml:"Amt"`
	CdtDbt *common.CreditDebitCode `xml:"CdtDbt,omitempty" json:",omitempty"`
}

func (r AmountAndDirection5) Validate() error {
	return utils.Validate(&r)
}

type BalanceTransfer3 struct {
	BalTrfRef     *BalanceTransferReference1    `xml:"BalTrfRef,omitempty" json:",omitempty"`
	BalTrfMtd     *SettlementMethod3Choice      `xml:"BalTrfMtd,omitempty" json:",omitempty"`
	BalTrfFndgLmt *BalanceTransferFundingLimit1 `xml:"BalTrfFndgLmt,omitempty" json:",omitempty"`
}

func (r BalanceTransfer3) Validate() error {
	return utils.Validate(&r)
}

type BalanceTransferFundingLimit1 struct {
	CcyAmt ActiveCurrencyAndAmount `xml:"CcyAmt"`
}

func (r BalanceTransferFundingLimit1) Validate() error {
	return utils.Validate(&r)
}

type BalanceTransferReference1 struct {
	BalTrfRef common.Max35Text `xml:"BalTrfRef"`
}

func (r BalanceTransferReference1) Validate() error {
	return utils.Validate(&r)
}

type CashAccount39 struct {
	Id   AccountIdentification4Choice                  `xml:"Id"`
	Tp   *CashAccountType2Choice                       `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode          `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                             `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1                  `xml:"Prxy,omitempty" json:",omitempty"`
	Ownr *PartyIdentification135                       `xml:"Ownr,omitempty" json:",omitempty"`
	Svcr *BranchAndFinancialInstitutionIdentification6 `xml:"Svcr,omitempty" json:",omitempty"`
}

func (r CashAccount39) Validate() error {
	return utils.Validate(&r)
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return utils.Validate(&r)
}

type Cheque11 struct {
	ChqTp       *ChequeType2Code             `xml:"ChqTp,omitempty" json:",omitempty"`
	ChqNb       *common.Max35Text            `xml:"ChqNb,omitempty" json:",omitempty"`
	ChqFr       *NameAndAddress16            `xml:"ChqFr,omitempty" json:",omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty" json:",omitempty"`
	DlvrTo      *NameAndAddress16            `xml:"DlvrTo,omitempty" json:",omitempty"`
	InstrPrty   *Priority2Code               `xml:"InstrPrty,omitempty" json:",omitempty"`
	ChqMtrtyDt  *common.ISODate              `xml:"ChqMtrtyDt,omitempty" json:",omitempty"`
	FrmsCd      *common.Max35Text            `xml:"FrmsCd,omitempty" json:",omitempty"`
	MemoFld     []common.Max35Text           `xml:"MemoFld,omitempty" json:",omitempty"`
	RgnlClrZone *common.Max35Text            `xml:"RgnlClrZone,omitempty" json:",omitempty"`
	PrtLctn     *common.Max35Text            `xml:"PrtLctn,omitempty" json:",omitempty"`
	Sgntr       []common.Max70Text           `xml:"Sgntr,omitempty" json:",omitempty"`
}

func (r Cheque11) Validate() error {
	return utils.Validate(&r)
}

type ChequeDeliveryMethod1Choice struct {
	Cd    ChequeDelivery1Code `xml:"Cd"`
	Prtry common.Max35Text    `xml:"Prtry"`
}

func (r ChequeDeliveryMethod1Choice) Validate() error {
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
	Issr      *common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

func (r CreditorReferenceType2) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
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
	Id   []DocumentLineIdentification1 `xml:"Id" json:",omitempty"`
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

type EndPoint1Choice struct {
	NbOfPmts  *common.Max35Text `xml:"NbOfPmts,omitempty" json:",omitempty"`
	LastPmtDt *common.ISODate   `xml:"LastPmtDt,omitempty" json:",omitempty"`
}

func (r EndPoint1Choice) Validate() error {
	return utils.Validate(&r)
}

type Frequency1 struct {
	Seq                 *common.Max3NumericText     `xml:"Seq,omitempty" json:",omitempty"`
	StartDt             common.ISODate              `xml:"StartDt"`
	EndPtChc            EndPoint1Choice             `xml:"EndPtChc"`
	ReqdFrqcyPttrn      *Frequency37Choice          `xml:"ReqdFrqcyPttrn,omitempty" json:",omitempty"`
	NonWorkgDayAdjstmnt *BusinessDayConvention1Code `xml:"NonWorkgDayAdjstmnt,omitempty" json:",omitempty"`
}

func (r Frequency1) Validate() error {
	return utils.Validate(&r)
}

type Frequency37Choice struct {
	Cd    Frequency10Code  `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

func (r Frequency37Choice) Validate() error {
	return utils.Validate(&r)
}

type Garnishment3 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment3) Validate() error {
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

type InstructionForCreditorAgent3 struct {
	Cd       *ExternalCreditorAgentInstruction1Code `xml:"Cd,omitempty" json:",omitempty"`
	InstrInf *common.Max140Text                     `xml:"InstrInf,omitempty" json:",omitempty"`
}

func (r InstructionForCreditorAgent3) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress24   `xml:"Adr"`
}

func (r NameAndAddress16) Validate() error {
	return utils.Validate(&r)
}

type PaymentIdentification6 struct {
	InstrId    *common.Max35Text        `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text         `xml:"EndToEndId"`
	UETR       *common.UUIDv4Identifier `xml:"UETR,omitempty" json:",omitempty"`
}

func (r PaymentIdentification6) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation26 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	SvcLvl    []ServiceLevel8Choice   `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation26) Validate() error {
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

type RegulatoryAuthority2 struct {
	Nm   *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	Ctry *common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
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

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocation6 struct {
	RmtId             *common.Max35Text              `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnMtd        *RemittanceLocationMethod2Code `xml:"RmtLctnMtd,omitempty" json:",omitempty"`
	RmtLctnElctrncAdr *common.Max2048Text            `xml:"RmtLctnElctrncAdr,omitempty" json:",omitempty"`
	RmtLctnPstlAdr    *NameAndAddress16              `xml:"RmtLctnPstlAdr,omitempty" json:",omitempty"`
}

func (r RemittanceLocation6) Validate() error {
	return utils.Validate(&r)
}

type ResponseDetails1 struct {
	RspnCd    common.Max35Text   `xml:"RspnCd"`
	AddtlDtls *common.Max350Text `xml:"AddtlDtls,omitempty" json:",omitempty"`
}

func (r ResponseDetails1) Validate() error {
	return utils.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceLevel8Choice) Validate() error {
	return utils.Validate(&r)
}

type SettlementMethod3Choice struct {
	Cdt CreditTransferTransaction41 `xml:"Cdt"`
	Dbt CreditTransferTransaction41 `xml:"Dbt"`
}

func (r SettlementMethod3Choice) Validate() error {
	return utils.Validate(&r)
}

type StructuredRegulatoryReporting3 struct {
	Tp   *common.Max35Text                  `xml:"Tp,omitempty" json:",omitempty"`
	Dt   *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	Ctry *common.CountryCode                `xml:"Ctry,omitempty" json:",omitempty"`
	Cd   *common.Max10Text                  `xml:"Cd,omitempty" json:",omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
	Inf  []common.Max35Text                 `xml:"Inf,omitempty" json:",omitempty"`
}

func (r StructuredRegulatoryReporting3) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification135        `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification135        `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation7               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment3                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation16) Validate() error {
	return utils.Validate(&r)
}

type TaxAmount2 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2                `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount2) Validate() error {
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

type AccountSwitchInformationResponseV03 struct {
	XMLName       *xml.Name                        `json:",omitempty"`
	Attr          *utils.Attr                      `xml:",attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1           `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1            `xml:"AcctSwtchDtls"`
	NewAcct       CashAccount39                    `xml:"NewAcct"`
	OdAcct        CashAccount39                    `xml:"OdAcct"`
	PmtInstr      []PaymentInstruction36           `xml:"PmtInstr,omitempty" json:",omitempty"`
	DrctDbtInstr  []DirectDebitInstructionDetails2 `xml:"DrctDbtInstr,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchInformationResponseV03) Validate() error {
	return utils.Validate(&r)
}

type DirectDebitInstructionDetails2 struct {
	MndtId                 common.Max35Text                   `xml:"MndtId"`
	AutomtdDrctDbtInstrInd bool                               `xml:"AutomtdDrctDbtInstrInd,omitempty" json:",omitempty"`
	DrctDbtTrfblInd        bool                               `xml:"DrctDbtTrfblInd,omitempty" json:",omitempty"`
	Cdtr                   PartyIdentification135             `xml:"Cdtr"`
	LastColltnCcyAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"LastColltnCcyAmt,omitempty" json:",omitempty"`
	LastColltnDt           *common.ISODate                    `xml:"LastColltnDt,omitempty" json:",omitempty"`
	OthrDtls               []TransferInstruction1             `xml:"OthrDtls,omitempty" json:",omitempty"`
}

func (r DirectDebitInstructionDetails2) Validate() error {
	return utils.Validate(&r)
}

type PaymentInstruction36 struct {
	PmtInfId        common.Max35Text                              `xml:"PmtInfId"`
	PmtMtd          PaymentMethod3Code                            `xml:"PmtMtd"`
	BtchBookg       bool                                          `xml:"BtchBookg,omitempty" json:",omitempty"`
	NbOfTxs         *common.Max15NumericText                      `xml:"NbOfTxs,omitempty" json:",omitempty"`
	CtrlSum         float64                                       `xml:"CtrlSum,omitempty" json:",omitempty"`
	PmtTpInf        *PaymentTypeInformation26                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	ReqdExctnDt     common.ISODate                                `xml:"ReqdExctnDt"`
	PoolgAdjstmntDt *common.ISODate                               `xml:"PoolgAdjstmntDt,omitempty" json:",omitempty"`
	Dbtr            PartyIdentification135                        `xml:"Dbtr"`
	DbtrAcct        CashAccount38                                 `xml:"DbtrAcct"`
	DbtrAgt         BranchAndFinancialInstitutionIdentification6  `xml:"DbtrAgt"`
	DbtrAgtAcct     *CashAccount38                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	InstrForDbtrAgt *common.Max140Text                            `xml:"InstrForDbtrAgt,omitempty" json:",omitempty"`
	UltmtDbtr       *PartyIdentification135                       `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	ChrgBr          *ChargeBearerType1Code                        `xml:"ChrgBr,omitempty" json:",omitempty"`
	ChrgsAcct       *CashAccount38                                `xml:"ChrgsAcct,omitempty" json:",omitempty"`
	ChrgsAcctAgt    *BranchAndFinancialInstitutionIdentification6 `xml:"ChrgsAcctAgt,omitempty" json:",omitempty"`
	CdtTrfTxInf     []CreditTransferTransaction41                 `xml:"CdtTrfTxInf" json:",omitempty"`
}

func (r PaymentInstruction36) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchCancelExistingPaymentV03 struct {
	XMLName       *xml.Name                        `json:",omitempty"`
	Attr          *utils.Attr                      `xml:",attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1           `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1            `xml:"AcctSwtchDtls"`
	OdAcct        CashAccount39                    `xml:"OdAcct"`
	PmtInstr      []PaymentInstruction36           `xml:"PmtInstr,omitempty" json:",omitempty"`
	DrctDbtInstr  []DirectDebitInstructionDetails2 `xml:"DrctDbtInstr,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchCancelExistingPaymentV03) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchRequestBalanceTransferV03 struct {
	XMLName       *xml.Name              `json:",omitempty"`
	Attr          *utils.Attr            `xml:",attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	NewAcct       CashAccount39          `xml:"NewAcct"`
	NmntdAcct     *CashAccount39         `xml:"NmntdAcct,omitempty" json:",omitempty"`
	BalTrf        []BalanceTransfer3     `xml:"BalTrf,omitempty" json:",omitempty"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchRequestBalanceTransferV03) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchRequestPaymentV03 struct {
	XMLName       *xml.Name                   `json:",omitempty"`
	Attr          *utils.Attr                 `xml:",attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1      `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1       `xml:"AcctSwtchDtls"`
	OdAcct        CashAccount39               `xml:"OdAcct"`
	CdtInstr      CreditTransferTransaction41 `xml:"CdtInstr"`
	SplmtryData   []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchRequestPaymentV03) Validate() error {
	return utils.Validate(&r)
}
