// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
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

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                     `xml:"MmbId"`
}

type CreateLimitV01 struct {
	MsgHdr      MessageHeader1       `xml:"MsgHdr"`
	LmtData     []LimitStructure4    `xml:"LmtData"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          *common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
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

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text       `xml:"Issr"`
	SchmeNm *common.Max35Text      `xml:"SchmeNm,omitempty" json:",omitempty"`
}

type LimitIdentification5 struct {
	SysId          *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	BilLmtCtrPtyId *BranchAndFinancialInstitutionIdentification6 `xml:"BilLmtCtrPtyId,omitempty" json:",omitempty"`
	Tp             LimitType1Choice                              `xml:"Tp"`
	AcctOwnr       *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId         *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type LimitStructure4 struct {
	LmtId     LimitIdentification5    `xml:"LmtId"`
	StartDtTm *DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice           `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

type LimitType1Choice struct {
	Cd    LimitType3Code   `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
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

type SupplementaryData1 struct {
	PlcAndNm *common.Max350Text         `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice         `xml:"Id"`
	Tp   *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy *ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type CreateStandingOrderV01 struct {
	MsgHdr      MessageHeader1               `xml:"MsgHdr"`
	StgOrdrId   StandingOrderIdentification4 `xml:"StgOrdrId"`
	ValSet      StandingOrder7               `xml:"ValSet"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

type DatePeriod2Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type ExecutionType1Choice struct {
	Tm  common.ISOTime   `xml:"Tm"`
	Evt EventType1Choice `xml:"Evt"`
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text       `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type StandingOrder7 struct {
	Amt          *Amount2Choice                                `xml:"Amt,omitempty" json:",omitempty"`
	Cdtr         *BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct     *CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	Dbtr         *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct     *CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	ExctnTp      *ExecutionType1Choice                         `xml:"ExctnTp,omitempty" json:",omitempty"`
	Frqcy        *Frequency2Code                               `xml:"Frqcy,omitempty" json:",omitempty"`
	VldtyPrd     *DatePeriod2Choice                            `xml:"VldtyPrd,omitempty" json:",omitempty"`
	ZeroSweepInd bool                                          `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
}

type StandingOrderIdentification4 struct {
	Id       *common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	Acct     CashAccount38                                 `xml:"Acct"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

type CreateReservationV01 struct {
	MsgHdr      MessageHeader1             `xml:"MsgHdr"`
	RsvatnId    ReservationIdentification2 `xml:"RsvatnId"`
	ValSet      Reservation4               `xml:"ValSet"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type Reservation4 struct {
	StartDtTm *DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice           `xml:"Amt"`
}

type ReservationIdentification2 struct {
	RsvatnId *common.Max35Text                             `xml:"RsvatnId,omitempty" json:",omitempty"`
	SysId    *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	Tp       ReservationType1Choice                        `xml:"Tp"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId   *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type ReservationType1Choice struct {
	Cd    ReservationType2Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type CommunicationAddress8 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"PstlAdr"`
	PhneNb   common.PhoneNumber       `xml:"PhneNb"`
	FaxNb    *common.PhoneNumber      `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max256Text       `xml:"EmailAdr,omitempty" json:",omitempty"`
}

type ContactIdentificationAndAddress1 struct {
	Nm     *common.Max35Text     `xml:"Nm,omitempty" json:",omitempty"`
	Role   PaymentRole1Code      `xml:"Role"`
	ComAdr CommunicationAddress8 `xml:"ComAdr"`
}

type CreateMemberV01 struct {
	MsgHdr      MessageHeader1              `xml:"MsgHdr"`
	MmbId       MemberIdentification3Choice `xml:"MmbId"`
	ValSet      Member6                     `xml:"ValSet"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type LongPostalAddress1Choice struct {
	Ustrd common.Max140Text            `xml:"Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"Strd"`
}

type Member6 struct {
	MmbRtrAdr []MemberIdentification3Choice      `xml:"MmbRtrAdr,omitempty" json:",omitempty"`
	CtctRef   []ContactIdentificationAndAddress1 `xml:"CtctRef,omitempty" json:",omitempty"`
	ComAdr    *CommunicationAddress8             `xml:"ComAdr,omitempty" json:",omitempty"`
}

type MemberIdentification3Choice struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Othr        GenericFinancialIdentification1     `xml:"Othr"`
}

type StructuredLongPostalAddress1 struct {
	BldgNm     *common.Max35Text  `xml:"BldgNm,omitempty" json:",omitempty"`
	StrtNm     *common.Max35Text  `xml:"StrtNm,omitempty" json:",omitempty"`
	StrtBldgId *common.Max35Text  `xml:"StrtBldgId,omitempty" json:",omitempty"`
	Flr        *common.Max16Text  `xml:"Flr,omitempty" json:",omitempty"`
	TwnNm      common.Max35Text   `xml:"TwnNm"`
	DstrctNm   *common.Max35Text  `xml:"DstrctNm,omitempty" json:",omitempty"`
	RgnId      *common.Max35Text  `xml:"RgnId,omitempty" json:",omitempty"`
	Stat       *common.Max35Text  `xml:"Stat,omitempty" json:",omitempty"`
	CtyId      *common.Max35Text  `xml:"CtyId,omitempty" json:",omitempty"`
	Ctry       common.CountryCode `xml:"Ctry"`
	PstCdId    common.Max16Text   `xml:"PstCdId"`
	POB        *common.Max16Text  `xml:"POB,omitempty" json:",omitempty"`
}
