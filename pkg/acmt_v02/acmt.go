// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"FinInstnId"`
	BrnchId    *BranchData2                        `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData2 struct {
	Id      *common.Max35Text  `xml:"Id,omitempty" json:",omitempty"`
	Nm      *common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress6    `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
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

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *common.BICFIIdentifier              `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text          `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                          `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                             `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                       `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

type IdentificationAssignment2 struct {
	MsgId   common.Max35Text                              `xml:"MsgId"`
	CreDtTm common.ISODateTime                            `xml:"CreDtTm"`
	Cretr   *Party12Choice                                `xml:"Cretr,omitempty" json:",omitempty"`
	FrstAgt *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty" json:",omitempty"`
	Assgnr  Party12Choice                                 `xml:"Assgnr"`
	Assgne  Party12Choice                                 `xml:"Assgne"`
}

type IdentificationInformation2 struct {
	Pty  *PartyIdentification43                        `xml:"Pty,omitempty" json:",omitempty"`
	Acct *AccountIdentification4Choice                 `xml:"Acct,omitempty" json:",omitempty"`
	Agt  *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty" json:",omitempty"`
}

type IdentificationModification2 struct {
	Id                common.Max35Text            `xml:"Id"`
	OrgnlPtyAndAcctId *IdentificationInformation2 `xml:"OrgnlPtyAndAcctId,omitempty" json:",omitempty"`
	UpdtdPtyAndAcctId IdentificationInformation2  `xml:"UpdtdPtyAndAcctId"`
	AddtlInf          *common.Max140Text          `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type IdentificationModificationAdviceV02 struct {
	Assgnmt     IdentificationAssignment2       `xml:"Assgnmt"`
	OrgnlTxRef  *OriginalTransactionReference18 `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
	Mod         []IdentificationModification2   `xml:"Mod"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type OrganisationIdentification8 struct {
	AnyBIC *common.AnyBICIdentifier             `xml:"AnyBIC,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

type OriginalTransactionReference18 struct {
	MsgId   *common.Max35Text        `xml:"MsgId,omitempty" json:",omitempty"`
	MsgNmId *common.Max35Text        `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime      `xml:"CreDtTm,omitempty" json:",omitempty"`
	OrgnlTx []PaymentIdentification4 `xml:"OrgnlTx,omitempty" json:",omitempty"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"OrgId"`
	PrvtId PersonIdentification5       `xml:"PrvtId"`
}

type Party12Choice struct {
	Pty PartyIdentification43                        `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

type PartyIdentification43 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress6     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party11Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *ContactDetails2    `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type PaymentIdentification4 struct {
	InstrId    *common.Max35Text                             `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text                              `xml:"EndToEndId"`
	TxId       common.Max35Text                              `xml:"TxId"`
	ClrSysRef  *common.Max35Text                             `xml:"ClrSysRef,omitempty" json:",omitempty"`
	FrstAgt    *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty" json:",omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
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

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type IdentificationVerification2 struct {
	Id           common.Max35Text           `xml:"Id"`
	PtyAndAcctId IdentificationInformation2 `xml:"PtyAndAcctId"`
}

type IdentificationVerificationRequestV02 struct {
	Assgnmt     IdentificationAssignment2     `xml:"Assgnmt"`
	Vrfctn      []IdentificationVerification2 `xml:"Vrfctn"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type IdentificationVerificationReportV02 struct {
	Assgnmt      IdentificationAssignment2 `xml:"Assgnmt"`
	OrgnlAssgnmt *MessageIdentification5   `xml:"OrgnlAssgnmt,omitempty" json:",omitempty"`
	Rpt          []VerificationReport2     `xml:"Rpt"`
	SplmtryData  []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MessageIdentification5 struct {
	MsgId   *common.Max35Text                             `xml:"MsgId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime                           `xml:"CreDtTm,omitempty" json:",omitempty"`
	FrstAgt *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty" json:",omitempty"`
}

type VerificationReason1Choice struct {
	Cd    ExternalVerificationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type VerificationReport2 struct {
	OrgnlId           common.Max35Text            `xml:"OrgnlId"`
	Vrfctn            bool                        `xml:"Vrfctn"`
	Rsn               *VerificationReason1Choice  `xml:"Rsn,omitempty" json:",omitempty"`
	OrgnlPtyAndAcctId *IdentificationInformation2 `xml:"OrgnlPtyAndAcctId,omitempty" json:",omitempty"`
	UpdtdPtyAndAcctId *IdentificationInformation2 `xml:"UpdtdPtyAndAcctId,omitempty" json:",omitempty"`
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

type AccountSwitchRequestRedirectionV02 struct {
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	NewAcct       CashAccount39          `xml:"NewAcct"`
	OdAcct        CashAccount39          `xml:"OdAcct"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                         `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData3 struct {
	Id      *common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     *common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      *common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr *PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
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

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
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

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

type MessageIdentification1 struct {
	Id      common.Max35Text   `xml:"Id"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type OrganisationIdentification29 struct {
	AnyBIC *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `xml:"ChanlTp"`
	Id      *common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

type PartyIdentification135 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
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

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type AccountSwitchNotifyAccountSwitchCompleteV02 struct {
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type AccountSwitchPaymentResponseV02 struct {
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type AccountSwitchTechnicalRejectionV02 struct {
	MsgId         MessageIdentification1 `xml:"MsgId"`
	AcctSwtchDtls AccountSwitchDetails1  `xml:"AcctSwtchDtls"`
	SplmtryData   []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type ResponseDetails1 struct {
	RspnCd    common.Max35Text   `xml:"RspnCd"`
	AddtlDtls *common.Max350Text `xml:"AddtlDtls,omitempty" json:",omitempty"`
}
