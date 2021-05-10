// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

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

type IdentificationAssignment2 struct {
	MsgId   common.Max35Text                              `xml:"MsgId"`
	CreDtTm common.ISODateTime                            `xml:"CreDtTm"`
	Cretr   *Party12Choice                                `xml:"Cretr,omitempty" json:",omitempty"`
	FrstAgt *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty" json:",omitempty"`
	Assgnr  Party12Choice                                 `xml:"Assgnr"`
	Assgne  Party12Choice                                 `xml:"Assgne"`
}

func (r IdentificationAssignment2) Validate() error {
	return utils.Validate(&r)
}

type IdentificationInformation2 struct {
	Pty  *PartyIdentification43                        `xml:"Pty,omitempty" json:",omitempty"`
	Acct *AccountIdentification4Choice                 `xml:"Acct,omitempty" json:",omitempty"`
	Agt  *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty" json:",omitempty"`
}

func (r IdentificationInformation2) Validate() error {
	return utils.Validate(&r)
}

type IdentificationModification2 struct {
	Id                common.Max35Text            `xml:"Id"`
	OrgnlPtyAndAcctId *IdentificationInformation2 `xml:"OrgnlPtyAndAcctId,omitempty" json:",omitempty"`
	UpdtdPtyAndAcctId IdentificationInformation2  `xml:"UpdtdPtyAndAcctId"`
	AddtlInf          *common.Max140Text          `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r IdentificationModification2) Validate() error {
	return utils.Validate(&r)
}

type IdentificationModificationAdviceV02 struct {
	Attr        []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	Assgnmt     IdentificationAssignment2       `xml:"Assgnmt"`
	OrgnlTxRef  *OriginalTransactionReference18 `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	Mod         []IdentificationModification2   `xml:"Mod" json:",omitempty"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r IdentificationModificationAdviceV02) Validate() error {
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

type OriginalTransactionReference18 struct {
	MsgId   *common.Max35Text        `xml:"MsgId,omitempty" json:",omitempty"`
	MsgNmId *common.Max35Text        `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime      `xml:"CreDtTm,omitempty" json:",omitempty"`
	OrgnlTx []PaymentIdentification4 `xml:"OrgnlTx,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference18) Validate() error {
	return utils.Validate(&r)
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"OrgId"`
	PrvtId PersonIdentification5       `xml:"PrvtId"`
}

func (r Party11Choice) Validate() error {
	return utils.Validate(&r)
}

type Party12Choice struct {
	Pty PartyIdentification43                        `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (r Party12Choice) Validate() error {
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

type PaymentIdentification4 struct {
	InstrId    *common.Max35Text                             `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text                              `xml:"EndToEndId"`
	TxId       common.Max35Text                              `xml:"TxId"`
	ClrSysRef  *common.Max35Text                             `xml:"ClrSysRef,omitempty" json:",omitempty"`
	FrstAgt    *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty" json:",omitempty"`
}

func (r PaymentIdentification4) Validate() error {
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

type IdentificationVerification2 struct {
	Id           common.Max35Text           `xml:"Id"`
	PtyAndAcctId IdentificationInformation2 `xml:"PtyAndAcctId"`
}

func (r IdentificationVerification2) Validate() error {
	return utils.Validate(&r)
}

type IdentificationVerificationRequestV02 struct {
	Attr        []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	Assgnmt     IdentificationAssignment2     `xml:"Assgnmt"`
	Vrfctn      []IdentificationVerification2 `xml:"Vrfctn" json:",omitempty"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r IdentificationVerificationRequestV02) Validate() error {
	return utils.Validate(&r)
}

type IdentificationVerificationReportV02 struct {
	Attr         []utils.Attr              `xml:",any,attr,omitempty" json:",omitempty"`
	Assgnmt      IdentificationAssignment2 `xml:"Assgnmt"`
	OrgnlAssgnmt *MessageIdentification5   `xml:"OrgnlAssgnmt,omitempty" json:",omitempty"`
	Rpt          []VerificationReport2     `xml:"Rpt" json:",omitempty"`
	SplmtryData  []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r IdentificationVerificationReportV02) Validate() error {
	return utils.Validate(&r)
}

type MessageIdentification5 struct {
	MsgId   *common.Max35Text                             `xml:"MsgId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime                           `xml:"CreDtTm,omitempty" json:",omitempty"`
	FrstAgt *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty" json:",omitempty"`
}

func (r MessageIdentification5) Validate() error {
	return utils.Validate(&r)
}

type VerificationReason1Choice struct {
	Cd    ExternalVerificationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r VerificationReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type VerificationReport2 struct {
	OrgnlId           common.Max35Text            `xml:"OrgnlId"`
	Vrfctn            bool                        `xml:"Vrfctn"`
	Rsn               *VerificationReason1Choice  `xml:"Rsn,omitempty" json:",omitempty"`
	OrgnlPtyAndAcctId *IdentificationInformation2 `xml:"OrgnlPtyAndAcctId,omitempty" json:",omitempty"`
	UpdtdPtyAndAcctId *IdentificationInformation2 `xml:"UpdtdPtyAndAcctId,omitempty" json:",omitempty"`
}

func (r VerificationReport2) Validate() error {
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

type AccountSwitchRequestRedirectionV02 struct {
	Attr          []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	NewAcct       CashAccount39          `xml:"NewAcct"`
	OdAcct        CashAccount39          `xml:"OdAcct"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchRequestRedirectionV02) Validate() error {
	return utils.Validate(&r)
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
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

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CashAccountType2Choice) Validate() error {
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

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth1) Validate() error {
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

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
	return utils.Validate(&r)
}

type MessageIdentification1 struct {
	Id      common.Max35Text   `xml:"Id"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r MessageIdentification1) Validate() error {
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

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification13) Validate() error {
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

type AccountSwitchNotifyAccountSwitchCompleteV02 struct {
	Attr          []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchNotifyAccountSwitchCompleteV02) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchPaymentResponseV02 struct {
	Attr          []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchPaymentResponseV02) Validate() error {
	return utils.Validate(&r)
}

type AccountSwitchTechnicalRejectionV02 struct {
	Attr          []utils.Attr           `xml:",any,attr,omitempty" json:",omitempty"`
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r AccountSwitchTechnicalRejectionV02) Validate() error {
	return utils.Validate(&r)
}

type ResponseDetails1 struct {
	RspnCd    common.Max35Text   `xml:"RspnCd"`
	AddtlDtls *common.Max350Text `xml:"AddtlDtls,omitempty" json:",omitempty"`
}

func (r ResponseDetails1) Validate() error {
	return utils.Validate(&r)
}
