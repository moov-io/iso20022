// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import "github.com/moov-io/iso20022/pkg/common"

type BusinessDayCriteria2 struct {
	NewQryNm common.Max35Text             `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []BusinessDaySearchCriteria2 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  BusinessDayReturnCriteria2   `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type BusinessDayCriteria3Choice struct {
	QryNm   common.Max35Text     `xml:"QryNm"`
	NewCrit BusinessDayCriteria2 `xml:"NewCrit"`
}

type BusinessDayQuery2 struct {
	QryTp QueryType2Code             `xml:"QryTp,omitempty" json:",omitempty"`
	Crit  BusinessDayCriteria3Choice `xml:"Crit,omitempty" json:",omitempty"`
}

type BusinessDayReturnCriteria2 struct {
	SysDtInd   bool `xml:"SysDtInd,omitempty" json:",omitempty"`
	SysStsInd  bool `xml:"SysStsInd,omitempty" json:",omitempty"`
	SysCcyInd  bool `xml:"SysCcyInd,omitempty" json:",omitempty"`
	ClsrPrdInd bool `xml:"ClsrPrdInd,omitempty" json:",omitempty"`
	EvtInd     bool `xml:"EvtInd,omitempty" json:",omitempty"`
	SsnPrdInd  bool `xml:"SsnPrdInd,omitempty" json:",omitempty"`
	EvtTpInd   bool `xml:"EvtTpInd,omitempty" json:",omitempty"`
}

type BusinessDaySearchCriteria2 struct {
	SysDt   common.ISODate                `xml:"SysDt,omitempty" json:",omitempty"`
	SysId   []SystemIdentification2Choice `xml:"SysId,omitempty" json:",omitempty"`
	SysCcy  []common.ActiveCurrencyCode   `xml:"SysCcy,omitempty" json:",omitempty"`
	EvtTp   SystemEventType2Choice        `xml:"EvtTp,omitempty" json:",omitempty"`
	ClsrPrd DateTimePeriod1Choice         `xml:"ClsrPrd,omitempty" json:",omitempty"`
}

type DateTimePeriod1 struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
}

type DateTimePeriod1Choice struct {
	FrDtTm common.ISODateTime `xml:"FrDtTm"`
	ToDtTm common.ISODateTime `xml:"ToDtTm"`
	DtTmRg DateTimePeriod1    `xml:"DtTmRg"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text `xml:"Id"`
	SchmeNm common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type GetBusinessDayInformationV05 struct {
	MsgHdr          MessageHeader9       `xml:"MsgHdr"`
	BizDayInfQryDef BusinessDayQuery2    `xml:"BizDayInfQryDef,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type MessageHeader9 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp   RequestType4Choice `xml:"ReqTp,omitempty" json:",omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemEventType2Choice struct {
	Cd    SystemEventType2Code   `xml:"Cd"`
	Prtry GenericIdentification1 `xml:"Prtry"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    BranchData3                          `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                    `xml:"MmbId"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"BICFI,omitempty" json:",omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:",omitempty"`
	LEI         common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	Nm          common.Max140Text                   `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr     PostalAddress24                     `xml:"PstlAdr,omitempty" json:",omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"Othr,omitempty" json:",omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                         `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm common.Max35Text              `xml:"SchmeNm,omitempty" json:",omitempty"`
}

type LongPaymentIdentification2 struct {
	TxId           common.Max35Text                             `xml:"TxId,omitempty" json:",omitempty"`
	UETR           common.UUIDv4Identifier                      `xml:"UETR,omitempty" json:",omitempty"`
	IntrBkSttlmAmt float64                                      `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                               `xml:"IntrBkSttlmDt"`
	PmtMtd         PaymentOrigin1Choice                         `xml:"PmtMtd,omitempty" json:",omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
	NtryTp         EntryTypeIdentifier                          `xml:"NtryTp,omitempty" json:",omitempty"`
	EndToEndId     common.Max35Text                             `xml:"EndToEndId,omitempty" json:",omitempty"`
}

type OriginalMessageAndIssuer1 struct {
	MsgId   common.Max35Text `xml:"MsgId"`
	MsgNmId common.Max35Text `xml:"MsgNmId,omitempty" json:",omitempty"`
	OrgtrNm common.Max70Text `xml:"OrgtrNm,omitempty" json:",omitempty"`
}

type PaymentIdentification6Choice struct {
	TxId      common.Max35Text                `xml:"TxId"`
	QId       QueueTransactionIdentification1 `xml:"QId"`
	LngBizId  LongPaymentIdentification2      `xml:"LngBizId"`
	ShrtBizId ShortPaymentIdentification2     `xml:"ShrtBizId"`
	PrtryId   common.Max70Text                `xml:"PrtryId"`
}

type PaymentOrigin1Choice struct {
	FINMT    common.Max3NumericText `xml:"FINMT"`
	XMLMsgNm common.Max35Text       `xml:"XMLMsgNm"`
	Prtry    common.Max35Text       `xml:"Prtry"`
	Instrm   PaymentInstrument1Code `xml:"Instrm"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"AdrTp,omitempty" json:",omitempty"`
	Dept        common.Max70Text   `xml:"Dept,omitempty" json:",omitempty"`
	SubDept     common.Max70Text   `xml:"SubDept,omitempty" json:",omitempty"`
	StrtNm      common.Max70Text   `xml:"StrtNm,omitempty" json:",omitempty"`
	BldgNb      common.Max16Text   `xml:"BldgNb,omitempty" json:",omitempty"`
	BldgNm      common.Max35Text   `xml:"BldgNm,omitempty" json:",omitempty"`
	Flr         common.Max70Text   `xml:"Flr,omitempty" json:",omitempty"`
	PstBx       common.Max16Text   `xml:"PstBx,omitempty" json:",omitempty"`
	Room        common.Max70Text   `xml:"Room,omitempty" json:",omitempty"`
	PstCd       common.Max16Text   `xml:"PstCd,omitempty" json:",omitempty"`
	TwnNm       common.Max35Text   `xml:"TwnNm,omitempty" json:",omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"TwnLctnNm,omitempty" json:",omitempty"`
	DstrctNm    common.Max35Text   `xml:"DstrctNm,omitempty" json:",omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"CtrySubDvsn,omitempty" json:",omitempty"`
	Ctry        common.CountryCode `xml:"Ctry,omitempty" json:",omitempty"`
	AdrLine     []common.Max70Text `xml:"AdrLine,omitempty" json:",omitempty"`
}

type QueueTransactionIdentification1 struct {
	QId    common.Max16Text `xml:"QId"`
	PosInQ common.Max16Text `xml:"PosInQ"`
}

type Receipt3 struct {
	OrgnlMsgId OriginalMessageAndIssuer1    `xml:"OrgnlMsgId"`
	OrgnlPmtId PaymentIdentification6Choice `xml:"OrgnlPmtId,omitempty" json:",omitempty"`
	ReqHdlg    []RequestHandling1           `xml:"ReqHdlg,omitempty" json:",omitempty"`
}

type ReceiptV05 struct {
	MsgHdr      MessageHeader9       `xml:"MsgHdr"`
	RctDtls     []Receipt3           `xml:"RctDtls"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type RequestHandling1 struct {
	StsCd common.Max4AlphaNumericText `xml:"StsCd"`
	Desc  common.Max140Text           `xml:"Desc,omitempty" json:",omitempty"`
}

type ShortPaymentIdentification2 struct {
	TxId          common.Max35Text                             `xml:"TxId"`
	IntrBkSttlmDt common.ISODate                               `xml:"IntrBkSttlmDt"`
	InstgAgt      BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
}

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type CaseForwardingNotification3 struct {
	Justfn CaseForwardingNotification3Code `xml:"Justfn"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"NmPrfx,omitempty" json:",omitempty"`
	Nm        common.Max140Text           `xml:"Nm,omitempty" json:",omitempty"`
	PhneNb    common.PhoneNumber          `xml:"PhneNb,omitempty" json:",omitempty"`
	MobNb     common.PhoneNumber          `xml:"MobNb,omitempty" json:",omitempty"`
	FaxNb     common.PhoneNumber          `xml:"FaxNb,omitempty" json:",omitempty"`
	EmailAdr  common.Max2048Text          `xml:"EmailAdr,omitempty" json:",omitempty"`
	EmailPurp common.Max35Text            `xml:"EmailPurp,omitempty" json:",omitempty"`
	JobTitl   common.Max35Text            `xml:"JobTitl,omitempty" json:",omitempty"`
	Rspnsblty common.Max35Text            `xml:"Rspnsblty,omitempty" json:",omitempty"`
	Dept      common.Max70Text            `xml:"Dept,omitempty" json:",omitempty"`
	Othr      []OtherContact1             `xml:"Othr,omitempty" json:",omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:",omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"PrvcOfBirth,omitempty" json:",omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                            `xml:"Issr,omitempty" json:",omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text                      `xml:"Issr,omitempty" json:",omitempty"`
}

type NotificationOfCaseAssignmentV05 struct {
	Hdr         ReportHeader5               `xml:"Hdr"`
	Case        Case5                       `xml:"Case"`
	Assgnmt     CaseAssignment5             `xml:"Assgnmt"`
	Ntfctn      CaseForwardingNotification3 `xml:"Ntfctn"`
	SplmtryData []SupplementaryData1        `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI    common.LEIIdentifier                 `xml:"LEI,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"ChanlTp"`
	Id      common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"OrgId"`
	PrvtId PersonIdentification13       `xml:"PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        Party38Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type ReportHeader5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Fr      Party40Choice      `xml:"Fr"`
	To      Party40Choice      `xml:"To"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

type ProprietaryData6 struct {
	Any SkipPayload `xml:"Any"`
}

type ProprietaryData7 struct {
	Tp   common.Max35Text `xml:"Tp"`
	Data ProprietaryData6 `xml:"Data"`
}

type ProprietaryFormatInvestigationV05 struct {
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        Case5                `xml:"Case,omitempty" json:",omitempty"`
	PrtryData   ProprietaryData7     `xml:"PrtryData"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type SkipPayload struct {
	Item string `xml:",any"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type DebitAuthorisationConfirmation2 struct {
	DbtAuthstn bool                    `xml:"DbtAuthstn"`
	AmtToDbt   ActiveCurrencyAndAmount `xml:"AmtToDbt,omitempty" json:",omitempty"`
	ValDtToDbt common.ISODate          `xml:"ValDtToDbt,omitempty" json:",omitempty"`
	Rsn        common.Max140Text       `xml:"Rsn,omitempty" json:",omitempty"`
}

type DebitAuthorisationResponseV05 struct {
	Assgnmt     CaseAssignment5                 `xml:"Assgnmt"`
	Case        Case5                           `xml:"Case,omitempty" json:",omitempty"`
	Conf        DebitAuthorisationConfirmation2 `xml:"Conf"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type CaseStatus2 struct {
	DtTm    common.ISODateTime `xml:"DtTm"`
	CaseSts CaseStatus2Code    `xml:"CaseSts"`
	Rsn     common.Max140Text  `xml:"Rsn,omitempty" json:",omitempty"`
}

type CaseStatusReportV05 struct {
	Hdr         ReportHeader5        `xml:"Hdr"`
	Case        Case5                `xml:"Case"`
	Sts         CaseStatus2          `xml:"Sts"`
	NewAssgnmt  CaseAssignment5      `xml:"NewAssgnmt,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"IBAN"`
	Othr GenericAccountIdentification1 `xml:"Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                   `xml:"Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"Id"`
	SchmeNm AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
}

type GetReservationV05 struct {
	MsgHdr       MessageHeader9       `xml:"MsgHdr"`
	RsvatnQryDef ReservationQuery3    `xml:"RsvatnQryDef,omitempty" json:",omitempty"`
	SplmtryData  []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type ReservationCriteria3Choice struct {
	QryNm   common.Max35Text     `xml:"QryNm"`
	NewCrit ReservationCriteria4 `xml:"NewCrit"`
}

type ReservationCriteria4 struct {
	NewQryNm common.Max35Text             `xml:"NewQryNm,omitempty" json:",omitempty"`
	SchCrit  []ReservationSearchCriteria3 `xml:"SchCrit,omitempty" json:",omitempty"`
	RtrCrit  ReservationReturnCriteria1   `xml:"RtrCrit,omitempty" json:",omitempty"`
}

type ReservationQuery3 struct {
	QryTp      QueryType2Code             `xml:"QryTp,omitempty" json:",omitempty"`
	RsvatnCrit ReservationCriteria3Choice `xml:"RsvatnCrit,omitempty" json:",omitempty"`
}

type ReservationReturnCriteria1 struct {
	StartDtTmInd bool `xml:"StartDtTmInd,omitempty" json:",omitempty"`
	StsInd       bool `xml:"StsInd,omitempty" json:",omitempty"`
}

type ReservationSearchCriteria3 struct {
	SysId        SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	DfltRsvatnTp []ReservationType1Code                       `xml:"DfltRsvatnTp,omitempty" json:",omitempty"`
	CurRsvatnTp  []ReservationType1Code                       `xml:"CurRsvatnTp,omitempty" json:",omitempty"`
	AcctOwnr     BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId       AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

type CurrentOrDefaultReservation2Choice struct {
	Cur  ReservationIdentification2 `xml:"Cur"`
	Dflt ReservationIdentification2 `xml:"Dflt"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

type ModifyReservationV05 struct {
	MsgHdr          MessageHeader1                     `xml:"MsgHdr"`
	RsvatnId        CurrentOrDefaultReservation2Choice `xml:"RsvatnId"`
	NewRsvatnValSet Reservation4                       `xml:"NewRsvatnValSet"`
	SplmtryData     []SupplementaryData1               `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type Reservation4 struct {
	StartDtTm DateAndDateTime2Choice `xml:"StartDtTm,omitempty" json:",omitempty"`
	Amt       Amount2Choice          `xml:"Amt"`
}

type DeleteReservationV05 struct {
	MsgHdr      MessageHeader1             `xml:"MsgHdr"`
	CurRsvatn   ReservationIdentification2 `xml:"CurRsvatn,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1       `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type MessageHeader1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type ReservationIdentification2 struct {
	RsvatnId common.Max35Text                             `xml:"RsvatnId,omitempty" json:",omitempty"`
	SysId    SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	Tp       ReservationType1Choice                       `xml:"Tp"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId   AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type ReservationType1Choice struct {
	Cd    ReservationType2Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"Id"`
	Tp   CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm   common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"Prxy,omitempty" json:",omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type LiquidityCreditTransfer2 struct {
	LqdtyTrfId PaymentIdentification8                       `xml:"LqdtyTrfId,omitempty" json:",omitempty"`
	Cdtr       BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct   CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	TrfdAmt    Amount2Choice                                `xml:"TrfdAmt"`
	Dbtr       BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct   CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	SttlmDt    common.ISODate                               `xml:"SttlmDt,omitempty" json:",omitempty"`
}

type LiquidityCreditTransferV05 struct {
	MsgHdr      MessageHeader1           `xml:"MsgHdr"`
	LqdtyCdtTrf LiquidityCreditTransfer2 `xml:"LqdtyCdtTrf"`
	SplmtryData []SupplementaryData1     `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type PaymentIdentification8 struct {
	InstrId    common.Max35Text        `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId common.Max35Text        `xml:"EndToEndId"`
	TxId       common.Max35Text        `xml:"TxId,omitempty" json:",omitempty"`
	UETR       common.UUIDv4Identifier `xml:"UETR,omitempty" json:",omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text      `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type LiquidityDebitTransfer2 struct {
	LqdtyTrfId PaymentIdentification8                       `xml:"LqdtyTrfId,omitempty" json:",omitempty"`
	Cdtr       BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct   CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	TrfdAmt    Amount2Choice                                `xml:"TrfdAmt"`
	Dbtr       BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct   CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	SttlmDt    common.ISODate                               `xml:"SttlmDt,omitempty" json:",omitempty"`
}

type LiquidityDebitTransferV05 struct {
	MsgHdr      MessageHeader1          `xml:"MsgHdr"`
	LqdtyDbtTrf LiquidityDebitTransfer2 `xml:"LqdtyDbtTrf"`
	SplmtryData []SupplementaryData1    `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type AccountReportingRequestV05 struct {
	GrpHdr      GroupHeader77        `xml:"GrpHdr"`
	RptgReq     []ReportingRequest5  `xml:"RptgReq"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type BalanceSubType1Choice struct {
	Cd    ExternalBalanceSubType1Code `xml:"Cd"`
	Prtry common.Max35Text            `xml:"Prtry"`
}

type BalanceType10Choice struct {
	Cd    ExternalBalanceType1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type BalanceType13 struct {
	CdOrPrtry BalanceType10Choice   `xml:"CdOrPrtry"`
	SubTp     BalanceSubType1Choice `xml:"SubTp,omitempty" json:",omitempty"`
}

type DatePeriodDetails1 struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt,omitempty" json:",omitempty"`
}

type EntryStatus1Choice struct {
	Cd    ExternalEntryStatus1Code `xml:"Cd"`
	Prtry common.Max35Text         `xml:"Prtry"`
}

type GroupHeader77 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
	MsgSndr Party40Choice      `xml:"MsgSndr,omitempty" json:",omitempty"`
}

type Limit2 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd FloorLimitType1Code               `xml:"CdtDbtInd"`
}

type ReportingPeriod2 struct {
	FrToDt DatePeriodDetails1 `xml:"FrToDt"`
	FrToTm TimePeriodDetails1 `xml:"FrToTm,omitempty" json:",omitempty"`
	Tp     QueryType3Code     `xml:"Tp"`
}

type ReportingRequest5 struct {
	Id          common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	ReqdMsgNmId common.Max35Text                             `xml:"ReqdMsgNmId"`
	Acct        CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr    Party40Choice                                `xml:"AcctOwnr"`
	AcctSvcr    BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RptgPrd     ReportingPeriod2                             `xml:"RptgPrd,omitempty" json:",omitempty"`
	RptgSeq     SequenceRange1Choice                         `xml:"RptgSeq,omitempty" json:",omitempty"`
	ReqdTxTp    TransactionType2                             `xml:"ReqdTxTp,omitempty" json:",omitempty"`
	ReqdBalTp   []BalanceType13                              `xml:"ReqdBalTp,omitempty" json:",omitempty"`
}

type SequenceRange1 struct {
	FrSeq common.Max35Text `xml:"FrSeq"`
	ToSeq common.Max35Text `xml:"ToSeq"`
}

type SequenceRange1Choice struct {
	FrSeq   common.Max35Text   `xml:"FrSeq"`
	ToSeq   common.Max35Text   `xml:"ToSeq"`
	FrToSeq []SequenceRange1   `xml:"FrToSeq"`
	EQSeq   []common.Max35Text `xml:"EQSeq"`
	NEQSeq  []common.Max35Text `xml:"NEQSeq"`
}

type TimePeriodDetails1 struct {
	FrTm common.ISOTime `xml:"FrTm"`
	ToTm common.ISOTime `xml:"ToTm,omitempty" json:",omitempty"`
}

type TransactionType2 struct {
	Sts       EntryStatus1Choice     `xml:"Sts"`
	CdtDbtInd common.CreditDebitCode `xml:"CdtDbtInd"`
	FlrLmt    []Limit2               `xml:"FlrLmt,omitempty" json:",omitempty"`
}
