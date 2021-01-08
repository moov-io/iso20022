// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import "github.com/moov-io/iso20022/pkg/common"

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"Err"`
	Desc common.Max140Text    `xml:"Desc,omitempty" json:",omitempty"`
}

type GeneralBusinessInformation1 struct {
	Qlfr     InformationQualifierType1 `xml:"Qlfr,omitempty" json:",omitempty"`
	Sbjt     common.Max35Text          `xml:"Sbjt,omitempty" json:",omitempty"`
	SbjtDtls common.Max350Text         `xml:"SbjtDtls,omitempty" json:",omitempty"`
}

type GeneralBusinessOrError7Choice struct {
	OprlErr []ErrorHandling5         `xml:"OprlErr"`
	BizRpt  []GeneralBusinessReport6 `xml:"BizRpt"`
}

type GeneralBusinessOrError8Choice struct {
	BizErr []ErrorHandling5            `xml:"BizErr"`
	GnlBiz GeneralBusinessInformation1 `xml:"GnlBiz"`
}

type GeneralBusinessReport6 struct {
	BizInfRef   common.Max35Text              `xml:"BizInfRef"`
	GnlBizOrErr GeneralBusinessOrError8Choice `xml:"GnlBizOrErr"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text `xml:"Id"`
	SchmeNm common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

type InformationQualifierType1 struct {
	IsFrmtd bool          `xml:"IsFrmtd,omitempty" json:",omitempty"`
	Prty    Priority1Code `xml:"Prty,omitempty" json:",omitempty"`
}

type MessageHeader7 struct {
	MsgId       common.Max35Text       `xml:"MsgId"`
	CreDtTm     common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp       RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	OrgnlBizQry OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	MsgNmId common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

type ReturnGeneralBusinessInformationV06 struct {
	MsgHdr      MessageHeader7                `xml:"MsgHdr"`
	RptOrErr    GeneralBusinessOrError7Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"PlcAndNm,omitempty" json:",omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

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
	BrnchId    BranchData3                          `xml:"BrnchId,omitempty" json:",omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"Id,omitempty" json:",omitempty"`
	LEI     common.LEIIdentifier `xml:"LEI,omitempty" json:",omitempty"`
	Nm      common.Max140Text    `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr PostalAddress24      `xml:"PstlAdr,omitempty" json:",omitempty"`
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

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:",omitempty"`
	MmbId    common.Max35Text                    `xml:"MmbId"`
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

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"Id"`
	SchmeNm AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    common.Max35Text         `xml:"Issr,omitempty" json:",omitempty"`
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

type MessageHeader1 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

type ModifyStandingOrderV06 struct {
	MsgHdr           MessageHeader1               `xml:"MsgHdr"`
	StgOrdrId        StandingOrderIdentification4 `xml:"StgOrdrId"`
	NewStgOrdrValSet StandingOrder7               `xml:"NewStgOrdrValSet"`
	SplmtryData      []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
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

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Id common.Max2048Text      `xml:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type StandingOrder7 struct {
	Amt          Amount2Choice                                `xml:"Amt,omitempty" json:",omitempty"`
	Cdtr         BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct     CashAccount38                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	Dbtr         BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct     CashAccount38                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	ExctnTp      ExecutionType1Choice                         `xml:"ExctnTp,omitempty" json:",omitempty"`
	Frqcy        Frequency2Code                               `xml:"Frqcy,omitempty" json:",omitempty"`
	VldtyPrd     DatePeriod2Choice                            `xml:"VldtyPrd,omitempty" json:",omitempty"`
	ZeroSweepInd bool                                         `xml:"ZeroSweepInd,omitempty" json:",omitempty"`
}

type StandingOrderIdentification4 struct {
	Id       common.Max35Text                             `xml:"Id,omitempty" json:",omitempty"`
	Acct     CashAccount38                                `xml:"Acct"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
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

type InvestigationRejectionJustification1 struct {
	RjctnRsn InvestigationRejection1Code `xml:"RjctnRsn"`
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

type RejectInvestigationV06 struct {
	Assgnmt     CaseAssignment5                      `xml:"Assgnmt"`
	Case        Case5                                `xml:"Case,omitempty" json:",omitempty"`
	Justfn      InvestigationRejectionJustification1 `xml:"Justfn"`
	SplmtryData []SupplementaryData1                 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type RequestForDuplicateV06 struct {
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        Case5                `xml:"Case,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type DuplicateV06 struct {
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        Case5                `xml:"Case,omitempty" json:",omitempty"`
	Dplct       ProprietaryData7     `xml:"Dplct"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type ProprietaryData6 struct {
	Any SkipPayload `xml:"Any"`
}

type ProprietaryData7 struct {
	Tp   common.Max35Text `xml:"Tp"`
	Data ProprietaryData6 `xml:"Data"`
}

type SkipPayload struct {
	Item string `xml:",any"`
}

type CurrentAndDefaultReservation4 struct {
	CurRsvatn  []ReservationReport6 `xml:"CurRsvatn,omitempty" json:",omitempty"`
	DfltRsvatn []ReservationReport6 `xml:"DfltRsvatn,omitempty" json:",omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

type Reservation3 struct {
	Amt       Amount2Choice            `xml:"Amt"`
	Sts       ReservationStatus1Choice `xml:"Sts,omitempty" json:",omitempty"`
	StartDtTm DateAndDateTime2Choice   `xml:"StartDtTm,omitempty" json:",omitempty"`
}

type ReservationIdentification2 struct {
	RsvatnId common.Max35Text                             `xml:"RsvatnId,omitempty" json:",omitempty"`
	SysId    SystemIdentification2Choice                  `xml:"SysId,omitempty" json:",omitempty"`
	Tp       ReservationType1Choice                       `xml:"Tp"`
	AcctOwnr BranchAndFinancialInstitutionIdentification6 `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctId   AccountIdentification4Choice                 `xml:"AcctId,omitempty" json:",omitempty"`
}

type ReservationOrError8Choice struct {
	BizRpt  CurrentAndDefaultReservation4 `xml:"BizRpt"`
	OprlErr []ErrorHandling5              `xml:"OprlErr"`
}

type ReservationOrError9Choice struct {
	Rsvatn Reservation3     `xml:"Rsvatn"`
	BizErr []ErrorHandling5 `xml:"BizErr"`
}

type ReservationReport6 struct {
	RsvatnId    ReservationIdentification2 `xml:"RsvatnId"`
	RsvatnOrErr ReservationOrError9Choice  `xml:"RsvatnOrErr"`
}

type ReservationStatus1Choice struct {
	Cd    ReservationStatus1Code `xml:"Cd"`
	Prtry common.Max35Text       `xml:"Prtry"`
}

type ReservationType1Choice struct {
	Cd    ReservationType2Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type ReturnReservationV06 struct {
	MsgHdr      MessageHeader7            `xml:"MsgHdr"`
	RptOrErr    ReservationOrError8Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

type AccountNotification16 struct {
	Id         common.Max35Text                             `xml:"Id"`
	Acct       CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	TtlAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"TtlAmt,omitempty" json:",omitempty"`
	XpctdValDt common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr       Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	Itm        []NotificationItem7                          `xml:"Itm"`
}

type NotificationItem7 struct {
	Id         common.Max35Text                             `xml:"Id"`
	EndToEndId common.Max35Text                             `xml:"EndToEndId,omitempty" json:",omitempty"`
	UETR       common.UUIDv4Identifier                      `xml:"UETR,omitempty" json:",omitempty"`
	Acct       CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	Amt        ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	XpctdValDt common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr       Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	Purp       Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RltdRmtInf RemittanceLocation7                          `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf     RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
}

type NotificationToReceiveV06 struct {
	GrpHdr      GroupHeader77         `xml:"GrpHdr"`
	Ntfctn      AccountNotification16 `xml:"Ntfctn"`
	SplmtryData []SupplementaryData1  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"Tp,omitempty" json:",omitempty"`
	Nb       common.Max35Text           `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt   common.ISODate             `xml:"RltdDt,omitempty" json:",omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"LineDtls,omitempty" json:",omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text            `xml:"Issr,omitempty" json:",omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:",omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"DscntApldAmt,omitempty" json:",omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:",omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"TaxAmt,omitempty" json:",omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:",omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:",omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    NameAndAddress16              `xml:"PstlAdr,omitempty" json:",omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       PartyIdentification135         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      PartyIdentification135         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      TaxInformation7                `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"Dtls,omitempty" json:",omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text           `xml:"Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"Titl,omitempty" json:",omitempty"`
	Nm   common.Max140Text `xml:"Nm,omitempty" json:",omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"Cdtr,omitempty" json:",omitempty"`
	Dbtr            TaxParty2                         `xml:"Dbtr,omitempty" json:",omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"AdmstnZone,omitempty" json:",omitempty"`
	RefNb           common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Mtd             common.Max35Text                  `xml:"Mtd,omitempty" json:",omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:",omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:",omitempty"`
	Dt              common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	SeqNb           float64                           `xml:"SeqNb,omitempty" json:",omitempty"`
	Rcrd            []TaxRecord2                      `xml:"Rcrd,omitempty" json:",omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId common.Max35Text `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp  common.Max35Text `xml:"TaxTp,omitempty" json:",omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"TaxId,omitempty" json:",omitempty"`
	RegnId  common.Max35Text  `xml:"RegnId,omitempty" json:",omitempty"`
	TaxTp   common.Max35Text  `xml:"TaxTp,omitempty" json:",omitempty"`
	Authstn TaxAuthorisation1 `xml:"Authstn,omitempty" json:",omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt DatePeriod2          `xml:"FrToDt,omitempty" json:",omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      TaxPeriod2        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   TaxAmount2        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"Tp,omitempty" json:",omitempty"`
	Ref common.Max35Text       `xml:"Ref,omitempty" json:",omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"Cd"`
	Prtry common.Max35Text  `xml:"Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text             `xml:"Issr,omitempty" json:",omitempty"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"Tp,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Rsn       common.Max4Text                   `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf  common.Max140Text                 `xml:"AddtlInf,omitempty" json:",omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"Tp,omitempty" json:",omitempty"`
	Nb     common.Max35Text  `xml:"Nb,omitempty" json:",omitempty"`
	RltdDt common.ISODate    `xml:"RltdDt,omitempty" json:",omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"Id"`
	Desc common.Max2048Text            `xml:"Desc,omitempty" json:",omitempty"`
	Amt  RemittanceAmount3             `xml:"Amt,omitempty" json:",omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text        `xml:"Issr,omitempty" json:",omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"Cd"`
	Prtry common.Max35Text              `xml:"Prtry"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"Tp"`
	Grnshee           PartyIdentification135            `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                              `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`
	Issr      common.Max35Text       `xml:"Issr,omitempty" json:",omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

type GroupHeader77 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
	MsgSndr Party40Choice      `xml:"MsgSndr,omitempty" json:",omitempty"`
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress24   `xml:"Adr"`
}

type NotificationToReceiveCancellationAdviceV06 struct {
	GrpHdr      GroupHeader77          `xml:"GrpHdr"`
	OrgnlNtfctn OriginalNotification12 `xml:"OrgnlNtfctn"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type OriginalItem6 struct {
	OrgnlItmId      common.Max35Text                  `xml:"OrgnlItmId"`
	OrgnlEndToEndId common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	UETR            common.UUIDv4Identifier           `xml:"UETR,omitempty" json:",omitempty"`
	Amt             ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	XpctdValDt      common.ISODate                    `xml:"XpctdValDt,omitempty" json:",omitempty"`
	OrgnlItmRef     OriginalItemReference5            `xml:"OrgnlItmRef,omitempty" json:",omitempty"`
}

type OriginalItemReference5 struct {
	Acct       CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	Dbtr       Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	Purp       Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RltdRmtInf RemittanceLocation7                          `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf     RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
}

type OriginalNotification12 struct {
	OrgnlMsgId     common.Max35Text                  `xml:"OrgnlMsgId"`
	OrgnlCreDtTm   common.ISODateTime                `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNtfctnId  common.Max35Text                  `xml:"OrgnlNtfctnId"`
	NtfctnCxl      bool                              `xml:"NtfctnCxl,omitempty" json:",omitempty"`
	OrgnlNtfctnRef []OriginalNotificationReference10 `xml:"OrgnlNtfctnRef,omitempty" json:",omitempty"`
}

type OriginalNotificationReference10 struct {
	Acct       CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	TtlAmt     ActiveOrHistoricCurrencyAndAmount            `xml:"TtlAmt,omitempty" json:",omitempty"`
	XpctdValDt common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr       Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	OrgnlItm   []OriginalItem6                              `xml:"OrgnlItm"`
}

type GroupHeader84 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
	MsgRcpt Party40Choice      `xml:"MsgRcpt,omitempty" json:",omitempty"`
}

type NotificationToReceiveStatusReportV06 struct {
	GrpHdr            GroupHeader84          `xml:"GrpHdr"`
	OrgnlNtfctnAndSts OriginalNotification11 `xml:"OrgnlNtfctnAndSts"`
	SplmtryData       []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

type OriginalItemAndStatus6 struct {
	OrgnlItmId      common.Max35Text                  `xml:"OrgnlItmId"`
	OrgnlEndToEndId common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       common.UUIDv4Identifier           `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	Amt             ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	XpctdValDt      common.ISODate                    `xml:"XpctdValDt,omitempty" json:",omitempty"`
	ItmSts          NotificationStatus3Code           `xml:"ItmSts"`
	AddtlStsInf     common.Max105Text                 `xml:"AddtlStsInf,omitempty" json:",omitempty"`
	OrgnlItmRef     OriginalItemReference5            `xml:"OrgnlItmRef,omitempty" json:",omitempty"`
}

type OriginalNotification11 struct {
	OrgnlMsgId     common.Max35Text                 `xml:"OrgnlMsgId"`
	OrgnlCreDtTm   common.ISODateTime               `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNtfctnId  common.Max35Text                 `xml:"OrgnlNtfctnId"`
	NtfctnSts      NotificationStatus3Code          `xml:"NtfctnSts,omitempty" json:",omitempty"`
	AddtlStsInf    common.Max140Text                `xml:"AddtlStsInf,omitempty" json:",omitempty"`
	OrgnlNtfctnRef []OriginalNotificationReference9 `xml:"OrgnlNtfctnRef,omitempty" json:",omitempty"`
}

type OriginalNotificationReference9 struct {
	Acct           CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr       Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr       BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct       CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	TtlAmt         ActiveOrHistoricCurrencyAndAmount            `xml:"TtlAmt,omitempty" json:",omitempty"`
	XpctdValDt     common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr           Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt      BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	OrgnlItmAndSts []OriginalItemAndStatus6                     `xml:"OrgnlItmAndSts"`
}
