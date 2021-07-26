// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import (
	"encoding/xml"

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

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

func (r Amount2Choice) Validate() error {
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

type CreateLimitV01 struct {
	XMLName     xml.Name             `xml:"CretLmt"`
	MsgHdr      MessageHeader1       `xml:"MsgHdr"`
	LmtData     []LimitStructure4    `xml:"LmtData" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreateLimitV01) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

func (r DateAndDateTime2Choice) Validate() error {
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

type LimitIdentification5 struct {
	SysId          *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	BilLmtCtrPtyId *BranchAndFinancialInstitutionIdentification6 `xml:"BilLmtCtrPtyId,omitempty" json:",omitempty"`
	Tp             LimitType1Choice                              `xml:"Tp"`
	AcctOwnr       *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId         *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

func (r LimitIdentification5) Validate() error {
	return utils.Validate(&r)
}

type LimitStructure4 struct {
	LmtId     LimitIdentification5    `xml:"LmtId"`
	StartDtTm *DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice           `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode `xml:"CdtDbtInd,omitempty" json:",omitempty"`
}

func (r LimitStructure4) Validate() error {
	return utils.Validate(&r)
}

type LimitType1Choice struct {
	Cd    LimitType3Code   `xml:"Cd"`
	Prtry common.Max35Text `xml:"Prtry"`
}

func (r LimitType1Choice) Validate() error {
	return utils.Validate(&r)
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r MarketInfrastructureIdentification1Choice) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r MessageHeader1) Validate() error {
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

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

func (r SystemIdentification2Choice) Validate() error {
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

type CreateStandingOrderV01 struct {
	XMLName     xml.Name                     `xml:"CretStgOrdr"`
	MsgHdr      MessageHeader1               `xml:"MsgHdr"`
	StgOrdrId   StandingOrderIdentification4 `xml:"StgOrdrId"`
	ValSet      StandingOrder7               `xml:"ValSet"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreateStandingOrderV01) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriod2) Validate() error {
	return utils.Validate(&r)
}

type DatePeriod2Choice struct {
	FrDt   common.ISODate `xml:"FrDt"`
	ToDt   common.ISODate `xml:"ToDt"`
	FrToDt DatePeriod2    `xml:"FrToDt"`
}

func (r DatePeriod2Choice) Validate() error {
	return utils.Validate(&r)
}

type EventType1Choice struct {
	Cd    ExternalSystemEventType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r EventType1Choice) Validate() error {
	return utils.Validate(&r)
}

type ExecutionType1Choice struct {
	Tm  common.ISOTime   `xml:"Tm"`
	Evt EventType1Choice `xml:"Evt"`
}

func (r ExecutionType1Choice) Validate() error {
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

func (r StandingOrder7) Validate() error {
	return utils.Validate(&r)
}

type StandingOrderIdentification4 struct {
	Id       *common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	Acct     CashAccount38                                 `xml:"Acct"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
}

func (r StandingOrderIdentification4) Validate() error {
	return utils.Validate(&r)
}

type CreateReservationV01 struct {
	XMLName     xml.Name                   `xml:"CretRsvatn"`
	MsgHdr      MessageHeader1             `xml:"MsgHdr"`
	RsvatnId    ReservationIdentification2 `xml:"RsvatnId"`
	ValSet      Reservation4               `xml:"ValSet"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreateReservationV01) Validate() error {
	return utils.Validate(&r)
}

type Reservation4 struct {
	StartDtTm *DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice           `xml:"Amt"`
}

func (r Reservation4) Validate() error {
	return utils.Validate(&r)
}

type ReservationIdentification2 struct {
	RsvatnId *common.Max35Text                             `xml:"RsvatnId,omitempty" json:",omitempty"`
	SysId    *SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	Tp       ReservationType1Choice                        `xml:"Tp"`
	AcctOwnr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId   *AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

func (r ReservationIdentification2) Validate() error {
	return utils.Validate(&r)
}

type ReservationType1Choice struct {
	Cd    ReservationType2Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r ReservationType1Choice) Validate() error {
	return utils.Validate(&r)
}

type CommunicationAddress8 struct {
	PstlAdr  LongPostalAddress1Choice `xml:"PstlAdr"`
	PhneNb   common.PhoneNumber       `xml:"PhneNb"`
	FaxNb    *common.PhoneNumber      `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr *common.Max256Text       `xml:"EmailAdr,omitempty" json:",omitempty"`
}

func (r CommunicationAddress8) Validate() error {
	return utils.Validate(&r)
}

type ContactIdentificationAndAddress1 struct {
	Nm     *common.Max35Text     `xml:"Nm,omitempty" json:",omitempty"`
	Role   PaymentRole1Code      `xml:"Role"`
	ComAdr CommunicationAddress8 `xml:"ComAdr"`
}

func (r ContactIdentificationAndAddress1) Validate() error {
	return utils.Validate(&r)
}

type CreateMemberV01 struct {
	XMLName     xml.Name                    `xml:"CretMmb"`
	MsgHdr      MessageHeader1              `xml:"MsgHdr"`
	MmbId       MemberIdentification3Choice `xml:"MmbId"`
	ValSet      Member6                     `xml:"ValSet"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreateMemberV01) Validate() error {
	return utils.Validate(&r)
}

type LongPostalAddress1Choice struct {
	Ustrd common.Max140Text            `xml:"Ustrd"`
	Strd  StructuredLongPostalAddress1 `xml:"Strd"`
}

func (r LongPostalAddress1Choice) Validate() error {
	return utils.Validate(&r)
}

type Member6 struct {
	MmbRtrAdr []MemberIdentification3Choice      `xml:"MmbRtrAdr,omitempty" json:",omitempty"`
	CtctRef   []ContactIdentificationAndAddress1 `xml:"CtctRef,omitempty" json:",omitempty"`
	ComAdr    *CommunicationAddress8             `xml:"ComAdr,omitempty" json:",omitempty"`
}

func (r Member6) Validate() error {
	return utils.Validate(&r)
}

type MemberIdentification3Choice struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	Othr        GenericFinancialIdentification1     `xml:"Othr"`
}

func (r MemberIdentification3Choice) Validate() error {
	return utils.Validate(&r)
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

func (r StructuredLongPostalAddress1) Validate() error {
	return utils.Validate(&r)
}
