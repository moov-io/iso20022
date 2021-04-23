// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type ErrorHandling3Choice struct {
	Cd    ExternalSystemErrorHandling1Code `xml:"Cd"`
	Prtry common.Max35Text                 `xml:"Prtry"`
}

func (r ErrorHandling3Choice) Validate() error {
	return utils.Validate(&r)
}

type ErrorHandling5 struct {
	Err  ErrorHandling3Choice `xml:"Err"`
	Desc *common.Max140Text   `xml:"Desc,omitempty" json:",omitempty"`
}

func (r ErrorHandling5) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessInformation1 struct {
	Qlfr     *InformationQualifierType1 `xml:"Qlfr,omitempty" json:",omitempty"`
	Sbjt     *common.Max35Text          `xml:"Sbjt,omitempty" json:",omitempty"`
	SbjtDtls *common.Max350Text         `xml:"SbjtDtls,omitempty" json:",omitempty"`
}

func (r GeneralBusinessInformation1) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessOrError7Choice struct {
	OprlErr []ErrorHandling5         `xml:"OprlErr" json:",omitempty"`
	BizRpt  []GeneralBusinessReport6 `xml:"BizRpt" json:",omitempty"`
}

func (r GeneralBusinessOrError7Choice) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessOrError8Choice struct {
	BizErr []ErrorHandling5            `xml:"BizErr" json:",omitempty"`
	GnlBiz GeneralBusinessInformation1 `xml:"GnlBiz" json:",omitempty"`
}

func (r GeneralBusinessOrError8Choice) Validate() error {
	return utils.Validate(&r)
}

type GeneralBusinessReport6 struct {
	BizInfRef   common.Max35Text              `xml:"BizInfRef"`
	GnlBizOrErr GeneralBusinessOrError8Choice `xml:"GnlBizOrErr"`
}

func (r GeneralBusinessReport6) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification1 struct {
	Id      common.Max35Text  `xml:"Id"`
	SchmeNm *common.Max35Text `xml:"SchmeNm,omitempty" json:",omitempty"`
	Issr    *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericIdentification1) Validate() error {
	return utils.Validate(&r)
}

type InformationQualifierType1 struct {
	IsFrmtd bool           `xml:"IsFrmtd,omitempty" json:",omitempty"`
	Prty    *Priority1Code `xml:"Prty,omitempty" json:",omitempty"`
}

func (r InformationQualifierType1) Validate() error {
	return utils.Validate(&r)
}

type MessageHeader7 struct {
	MsgId       common.Max35Text        `xml:"MsgId"`
	CreDtTm     *common.ISODateTime     `xml:"CreDtTm,omitempty" json:",omitempty"`
	ReqTp       *RequestType4Choice     `xml:"ReqTp,omitempty" json:",omitempty"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty" json:",omitempty"`
	QryNm       *common.Max35Text       `xml:"QryNm,omitempty" json:",omitempty"`
}

func (r MessageHeader7) Validate() error {
	return utils.Validate(&r)
}

type OriginalBusinessQuery1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId *common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalBusinessQuery1) Validate() error {
	return utils.Validate(&r)
}

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"Enqry"`
	Prtry   GenericIdentification1                 `xml:"Prtry"`
}

func (r RequestType4Choice) Validate() error {
	return utils.Validate(&r)
}

type ReturnGeneralBusinessInformationV06 struct {
	XMLName     *xml.Name                     `json:",omitempty"`
	Attr        *utils.Attr                   `xml:",attr,omitempty" json:",omitempty"`
	MsgHdr      MessageHeader7                `xml:"MsgHdr"`
	RptOrErr    GeneralBusinessOrError7Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReturnGeneralBusinessInformationV06) Validate() error {
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

type MessageHeader1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r MessageHeader1) Validate() error {
	return utils.Validate(&r)
}

type ModifyStandingOrderV06 struct {
	XMLName          *xml.Name                    `json:",omitempty"`
	Attr             *utils.Attr                  `xml:",attr,omitempty" json:",omitempty"`
	MsgHdr           MessageHeader1               `xml:"MsgHdr"`
	StgOrdrId        StandingOrderIdentification4 `xml:"StgOrdrId"`
	NewStgOrdrValSet StandingOrder7               `xml:"NewStgOrdrValSet"`
	SplmtryData      []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ModifyStandingOrderV06) Validate() error {
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

type Case5 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party40Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

func (r Case5) Validate() error {
	return utils.Validate(&r)
}

type CaseAssignment5 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party40Choice      `xml:"Assgnr"`
	Assgne  Party40Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r CaseAssignment5) Validate() error {
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

type InvestigationRejectionJustification1 struct {
	RjctnRsn InvestigationRejection1Code `xml:"RjctnRsn"`
}

func (r InvestigationRejectionJustification1) Validate() error {
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

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

func (r Party40Choice) Validate() error {
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

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type RejectInvestigationV06 struct {
	XMLName     *xml.Name                            `json:",omitempty"`
	Attr        *utils.Attr                          `xml:",attr,omitempty" json:",omitempty"`
	Assgnmt     CaseAssignment5                      `xml:"Assgnmt"`
	Case        *Case5                               `xml:"Case,omitempty" json:",omitempty"`
	Justfn      InvestigationRejectionJustification1 `xml:"Justfn"`
	SplmtryData []SupplementaryData1                 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RejectInvestigationV06) Validate() error {
	return utils.Validate(&r)
}

type RequestForDuplicateV06 struct {
	XMLName     *xml.Name            `json:",omitempty"`
	Attr        *utils.Attr          `xml:",attr,omitempty" json:",omitempty"`
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        *Case5               `xml:"Case,omitempty" json:",omitempty"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestForDuplicateV06) Validate() error {
	return utils.Validate(&r)
}

type DuplicateV06 struct {
	XMLName     *xml.Name            `json:",omitempty"`
	Attr        *utils.Attr          `xml:",attr,omitempty" json:",omitempty"`
	Assgnmt     CaseAssignment5      `xml:"Assgnmt"`
	Case        *Case5               `xml:"Case,omitempty" json:",omitempty"`
	Dplct       ProprietaryData7     `xml:"Dplct"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DuplicateV06) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryData6 struct {
	Any SkipPayload `xml:"Any"`
}

func (r ProprietaryData6) Validate() error {
	return utils.Validate(&r)
}

type ProprietaryData7 struct {
	Tp   common.Max35Text `xml:"Tp"`
	Data ProprietaryData6 `xml:"Data"`
}

func (r ProprietaryData7) Validate() error {
	return utils.Validate(&r)
}

type SkipPayload struct {
	Item string `xml:",any"`
}

func (r SkipPayload) Validate() error {
	return utils.Validate(&r)
}

type CurrentAndDefaultReservation4 struct {
	CurRsvatn  []ReservationReport6 `xml:"CurRsvatn,omitempty" json:",omitempty"`
	DfltRsvatn []ReservationReport6 `xml:"DfltRsvatn,omitempty" json:",omitempty"`
}

func (r CurrentAndDefaultReservation4) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

func (r DateAndDateTime2Choice) Validate() error {
	return utils.Validate(&r)
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r MarketInfrastructureIdentification1Choice) Validate() error {
	return utils.Validate(&r)
}

type Reservation3 struct {
	Amt       Amount2Choice             `xml:"Amt"`
	Sts       *ReservationStatus1Choice `xml:"Sts,omitempty" json:",omitempty"`
	StartDtTm *DateAndDateTime2Choice   `xml:"StartDtTm,omitempty" json:",omitempty"`
}

func (r Reservation3) Validate() error {
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

type ReservationOrError8Choice struct {
	BizRpt  CurrentAndDefaultReservation4 `xml:"BizRpt"`
	OprlErr []ErrorHandling5              `xml:"OprlErr" json:",omitempty"`
}

func (r ReservationOrError8Choice) Validate() error {
	return utils.Validate(&r)
}

type ReservationOrError9Choice struct {
	Rsvatn Reservation3     `xml:"Rsvatn"`
	BizErr []ErrorHandling5 `xml:"BizErr" json:",omitempty"`
}

func (r ReservationOrError9Choice) Validate() error {
	return utils.Validate(&r)
}

type ReservationReport6 struct {
	RsvatnId    ReservationIdentification2 `xml:"RsvatnId"`
	RsvatnOrErr ReservationOrError9Choice  `xml:"RsvatnOrErr"`
}

func (r ReservationReport6) Validate() error {
	return utils.Validate(&r)
}

type ReservationStatus1Choice struct {
	Cd    ReservationStatus1Code `xml:"Cd"`
	Prtry common.Max35Text       `xml:"Prtry"`
}

func (r ReservationStatus1Choice) Validate() error {
	return utils.Validate(&r)
}

type ReservationType1Choice struct {
	Cd    ReservationType2Code `xml:"Cd"`
	Prtry common.Max35Text     `xml:"Prtry"`
}

func (r ReservationType1Choice) Validate() error {
	return utils.Validate(&r)
}

type ReturnReservationV06 struct {
	XMLName     *xml.Name                 `json:",omitempty"`
	Attr        *utils.Attr               `xml:",attr,omitempty" json:",omitempty"`
	MsgHdr      MessageHeader7            `xml:"MsgHdr"`
	RptOrErr    ReservationOrError8Choice `xml:"RptOrErr"`
	SplmtryData []SupplementaryData1      `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ReturnReservationV06) Validate() error {
	return utils.Validate(&r)
}

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"MktInfrstrctrId"`
	Ctry            common.CountryCode                        `xml:"Ctry"`
}

func (r SystemIdentification2Choice) Validate() error {
	return utils.Validate(&r)
}

type AccountNotification16 struct {
	Id         common.Max35Text                              `xml:"Id"`
	Acct       *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   *Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   *CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	TtlAmt     *ActiveOrHistoricCurrencyAndAmount            `xml:"TtlAmt,omitempty" json:",omitempty"`
	XpctdValDt *common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr       *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	Itm        []NotificationItem7                           `xml:"Itm" json:",omitempty"`
}

func (r AccountNotification16) Validate() error {
	return utils.Validate(&r)
}

type NotificationItem7 struct {
	Id         common.Max35Text                              `xml:"Id"`
	EndToEndId *common.Max35Text                             `xml:"EndToEndId,omitempty" json:",omitempty"`
	UETR       *common.UUIDv4Identifier                      `xml:"UETR,omitempty" json:",omitempty"`
	Acct       *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   *Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   *CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	Amt        *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt"`
	XpctdValDt *common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr       *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	Purp       *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RltdRmtInf *RemittanceLocation7                          `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf     *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
}

func (r NotificationItem7) Validate() error {
	return utils.Validate(&r)
}

type NotificationToReceiveV06 struct {
	XMLName     *xml.Name             `json:",omitempty"`
	Attr        *utils.Attr           `xml:",attr,omitempty" json:",omitempty"`
	GrpHdr      GroupHeader77         `xml:"GrpHdr"`
	Ntfctn      AccountNotification16 `xml:"Ntfctn"`
	SplmtryData []SupplementaryData1  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r NotificationToReceiveV06) Validate() error {
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

type RemittanceLocation7 struct {
	RmtId       *common.Max35Text         `xml:"RmtId,omitempty" json:",omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"RmtLctnDtls,omitempty" json:",omitempty"`
}

func (r RemittanceLocation7) Validate() error {
	return utils.Validate(&r)
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"Mtd"`
	ElctrncAdr *common.Max2048Text           `xml:"ElctrncAdr,omitempty" json:",omitempty"`
	PstlAdr    *NameAndAddress16             `xml:"PstlAdr,omitempty" json:",omitempty"`
}

func (r RemittanceLocationData1) Validate() error {
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

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

func (r ActiveOrHistoricCurrencyAndAmount) Validate() error {
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

type GroupHeader77 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
	MsgSndr *Party40Choice     `xml:"MsgSndr,omitempty" json:",omitempty"`
}

func (r GroupHeader77) Validate() error {
	return utils.Validate(&r)
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"Nm"`
	Adr PostalAddress24   `xml:"Adr"`
}

func (r NameAndAddress16) Validate() error {
	return utils.Validate(&r)
}

type NotificationToReceiveCancellationAdviceV06 struct {
	XMLName     *xml.Name              `json:",omitempty"`
	Attr        *utils.Attr            `xml:",attr,omitempty" json:",omitempty"`
	GrpHdr      GroupHeader77          `xml:"GrpHdr"`
	OrgnlNtfctn OriginalNotification12 `xml:"OrgnlNtfctn"`
	SplmtryData []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r NotificationToReceiveCancellationAdviceV06) Validate() error {
	return utils.Validate(&r)
}

type OriginalItem6 struct {
	OrgnlItmId      common.Max35Text                  `xml:"OrgnlItmId"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	UETR            *common.UUIDv4Identifier          `xml:"UETR,omitempty" json:",omitempty"`
	Amt             ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	XpctdValDt      *common.ISODate                   `xml:"XpctdValDt,omitempty" json:",omitempty"`
	OrgnlItmRef     *OriginalItemReference5           `xml:"OrgnlItmRef,omitempty" json:",omitempty"`
}

func (r OriginalItem6) Validate() error {
	return utils.Validate(&r)
}

type OriginalItemReference5 struct {
	Acct       *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   *Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   *CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	Dbtr       *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	Purp       *Purpose2Choice                               `xml:"Purp,omitempty" json:",omitempty"`
	RltdRmtInf *RemittanceLocation7                          `xml:"RltdRmtInf,omitempty" json:",omitempty"`
	RmtInf     *RemittanceInformation16                      `xml:"RmtInf,omitempty" json:",omitempty"`
}

func (r OriginalItemReference5) Validate() error {
	return utils.Validate(&r)
}

type OriginalNotification12 struct {
	OrgnlMsgId     common.Max35Text                  `xml:"OrgnlMsgId"`
	OrgnlCreDtTm   *common.ISODateTime               `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNtfctnId  common.Max35Text                  `xml:"OrgnlNtfctnId"`
	NtfctnCxl      bool                              `xml:"NtfctnCxl,omitempty" json:",omitempty"`
	OrgnlNtfctnRef []OriginalNotificationReference10 `xml:"OrgnlNtfctnRef,omitempty" json:",omitempty"`
}

func (r OriginalNotification12) Validate() error {
	return utils.Validate(&r)
}

type OriginalNotificationReference10 struct {
	Acct       *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr   *Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr   *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct   *CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	TtlAmt     *ActiveOrHistoricCurrencyAndAmount            `xml:"TtlAmt,omitempty" json:",omitempty"`
	XpctdValDt *common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr       *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt    *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt  *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	OrgnlItm   []OriginalItem6                               `xml:"OrgnlItm" json:",omitempty"`
}

func (r OriginalNotificationReference10) Validate() error {
	return utils.Validate(&r)
}

type GroupHeader84 struct {
	MsgId   common.Max35Text   `xml:"MsgId"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
	MsgRcpt *Party40Choice     `xml:"MsgRcpt,omitempty" json:",omitempty"`
}

func (r GroupHeader84) Validate() error {
	return utils.Validate(&r)
}

type NotificationToReceiveStatusReportV06 struct {
	XMLName           *xml.Name              `json:",omitempty"`
	Attr              *utils.Attr            `xml:",attr,omitempty" json:",omitempty"`
	GrpHdr            GroupHeader84          `xml:"GrpHdr"`
	OrgnlNtfctnAndSts OriginalNotification11 `xml:"OrgnlNtfctnAndSts"`
	SplmtryData       []SupplementaryData1   `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r NotificationToReceiveStatusReportV06) Validate() error {
	return utils.Validate(&r)
}

type OriginalItemAndStatus6 struct {
	OrgnlItmId      common.Max35Text                  `xml:"OrgnlItmId"`
	OrgnlEndToEndId *common.Max35Text                 `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlUETR       *common.UUIDv4Identifier          `xml:"OrgnlUETR,omitempty" json:",omitempty"`
	Amt             ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
	XpctdValDt      *common.ISODate                   `xml:"XpctdValDt,omitempty" json:",omitempty"`
	ItmSts          NotificationStatus3Code           `xml:"ItmSts"`
	AddtlStsInf     *common.Max105Text                `xml:"AddtlStsInf,omitempty" json:",omitempty"`
	OrgnlItmRef     *OriginalItemReference5           `xml:"OrgnlItmRef,omitempty" json:",omitempty"`
}

func (r OriginalItemAndStatus6) Validate() error {
	return utils.Validate(&r)
}

type OriginalNotification11 struct {
	OrgnlMsgId     common.Max35Text                 `xml:"OrgnlMsgId"`
	OrgnlCreDtTm   *common.ISODateTime              `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNtfctnId  common.Max35Text                 `xml:"OrgnlNtfctnId"`
	NtfctnSts      *NotificationStatus3Code         `xml:"NtfctnSts,omitempty" json:",omitempty"`
	AddtlStsInf    *common.Max140Text               `xml:"AddtlStsInf,omitempty" json:",omitempty"`
	OrgnlNtfctnRef []OriginalNotificationReference9 `xml:"OrgnlNtfctnRef,omitempty" json:",omitempty"`
}

func (r OriginalNotification11) Validate() error {
	return utils.Validate(&r)
}

type OriginalNotificationReference9 struct {
	Acct           *CashAccount38                                `xml:"Acct,omitempty" json:",omitempty"`
	AcctOwnr       *Party40Choice                                `xml:"AcctOwnr,omitempty" json:",omitempty"`
	AcctSvcr       *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr,omitempty" json:",omitempty"`
	RltdAcct       *CashAccount38                                `xml:"RltdAcct,omitempty" json:",omitempty"`
	TtlAmt         *ActiveOrHistoricCurrencyAndAmount            `xml:"TtlAmt,omitempty" json:",omitempty"`
	XpctdValDt     *common.ISODate                               `xml:"XpctdValDt,omitempty" json:",omitempty"`
	Dbtr           *Party40Choice                                `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	IntrmyAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty" json:",omitempty"`
	OrgnlItmAndSts []OriginalItemAndStatus6                      `xml:"OrgnlItmAndSts" json:",omitempty"`
}

func (r OriginalNotificationReference9) Validate() error {
	return utils.Validate(&r)
}

type AmendmentInformationDetails10 struct {
	OrgnlMndtId      *common.Max35Text                             `xml:"OrgnlMndtId,omitempty" json:",omitempty"`
	OrgnlCdtrSchmeId *PartyIdentification43                        `xml:"OrgnlCdtrSchmeId,omitempty" json:",omitempty"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty" json:",omitempty"`
	OrgnlCdtrAgtAcct *CashAccount24                                `xml:"OrgnlCdtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlDbtr        *PartyIdentification43                        `xml:"OrgnlDbtr,omitempty" json:",omitempty"`
	OrgnlDbtrAcct    *CashAccount24                                `xml:"OrgnlDbtrAcct,omitempty" json:",omitempty"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty" json:",omitempty"`
	OrgnlDbtrAgtAcct *CashAccount24                                `xml:"OrgnlDbtrAgtAcct,omitempty" json:",omitempty"`
	OrgnlFnlColltnDt *common.ISODate                               `xml:"OrgnlFnlColltnDt,omitempty" json:",omitempty"`
	OrgnlFrqcy       *Frequency21Choice                            `xml:"OrgnlFrqcy,omitempty" json:",omitempty"`
	OrgnlRsn         *MandateSetupReason1Choice                    `xml:"OrgnlRsn,omitempty" json:",omitempty"`
}

func (r AmendmentInformationDetails10) Validate() error {
	return utils.Validate(&r)
}

type AmountType4Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"EqvtAmt"`
}

func (r AmountType4Choice) Validate() error {
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

type CancellationStatusReason2 struct {
	Orgtr    *PartyIdentification43           `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      *CancellationStatusReason2Choice `xml:"Rsn,omitempty" json:",omitempty"`
	AddtlInf []common.Max105Text              `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r CancellationStatusReason2) Validate() error {
	return utils.Validate(&r)
}

type CancellationStatusReason2Choice struct {
	Cd    PaymentCancellationRejection2Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r CancellationStatusReason2Choice) Validate() error {
	return utils.Validate(&r)
}

type Case3 struct {
	Id             common.Max35Text `xml:"Id"`
	Cretr          Party12Choice    `xml:"Cretr"`
	ReopCaseIndctn bool             `xml:"ReopCaseIndctn,omitempty" json:",omitempty"`
}

func (r Case3) Validate() error {
	return utils.Validate(&r)
}

type CaseAssignment3 struct {
	Id      common.Max35Text   `xml:"Id"`
	Assgnr  Party12Choice      `xml:"Assgnr"`
	Assgne  Party12Choice      `xml:"Assgne"`
	CreDtTm common.ISODateTime `xml:"CreDtTm"`
}

func (r CaseAssignment3) Validate() error {
	return utils.Validate(&r)
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice         `xml:"Id"`
	Tp  *CashAccountType2Choice              `xml:"Tp,omitempty" json:",omitempty"`
	Ccy *common.ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:",omitempty"`
	Nm  *common.Max70Text                    `xml:"Nm,omitempty" json:",omitempty"`
}

func (r CashAccount24) Validate() error {
	return utils.Validate(&r)
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r CategoryPurpose1Choice) Validate() error {
	return utils.Validate(&r)
}

type ChargeType3Choice struct {
	Cd    ExternalChargeType1Code `xml:"Cd"`
	Prtry GenericIdentification3  `xml:"Prtry"`
}

func (r ChargeType3Choice) Validate() error {
	return utils.Validate(&r)
}

type Charges3 struct {
	TtlChrgsAndTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty" json:",omitempty"`
	Rcrd              []ChargesRecord1                   `xml:"Rcrd,omitempty"`
}

func (r Charges3) Validate() error {
	return utils.Validate(&r)
}

type ChargesRecord1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount             `xml:"Amt"`
	CdtDbtInd *common.CreditDebitCode                       `xml:"CdtDbtInd,omitempty" json:",omitempty"`
	Tp        *ChargeType3Choice                            `xml:"Tp,omitempty" json:",omitempty"`
	Rate      float64                                       `xml:"Rate,omitempty" json:",omitempty"`
	Br        *ChargeBearerType1Code                        `xml:"Br,omitempty" json:",omitempty"`
	Agt       *BranchAndFinancialInstitutionIdentification5 `xml:"Agt,omitempty" json:",omitempty"`
	Tax       *TaxCharges2                                  `xml:"Tax,omitempty" json:",omitempty"`
}

func (r ChargesRecord1) Validate() error {
	return utils.Validate(&r)
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"Cd"`
	Prtry common.Max35Text                `xml:"Prtry"`
}

func (r ClearingSystemIdentification3Choice) Validate() error {
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

type CorrectiveGroupInformation1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId common.Max35Text    `xml:"MsgNmId"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r CorrectiveGroupInformation1) Validate() error {
	return utils.Validate(&r)
}

type CorrectiveInterbankTransaction1 struct {
	GrpHdr         *CorrectiveGroupInformation1       `xml:"GrpHdr,omitempty" json:",omitempty"`
	InstrId        *common.Max35Text                  `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId     *common.Max35Text                  `xml:"EndToEndId,omitempty" json:",omitempty"`
	TxId           *common.Max35Text                  `xml:"TxId,omitempty" json:",omitempty"`
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                    `xml:"IntrBkSttlmDt" json:",omitempty"`
}

func (r CorrectiveInterbankTransaction1) Validate() error {
	return utils.Validate(&r)
}

type CorrectivePaymentInitiation1 struct {
	GrpHdr       *CorrectiveGroupInformation1       `xml:"GrpHdr,omitempty" json:",omitempty"`
	PmtInfId     *common.Max35Text                  `xml:"PmtInfId,omitempty" json:",omitempty"`
	InstrId      *common.Max35Text                  `xml:"InstrId,omitempty" json:",omitempty"`
	EndToEndId   *common.Max35Text                  `xml:"EndToEndId,omitempty" json:",omitempty"`
	InstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`
	ReqdExctnDt  *common.ISODate                    `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	ReqdColltnDt *common.ISODate                    `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
}

func (r CorrectivePaymentInitiation1) Validate() error {
	return utils.Validate(&r)
}

type CorrectiveTransaction1Choice struct {
	Initn  CorrectivePaymentInitiation1    `xml:"Initn"`
	IntrBk CorrectiveInterbankTransaction1 `xml:"IntrBk"`
}

func (r CorrectiveTransaction1Choice) Validate() error {
	return utils.Validate(&r)
}

type DateAndPlaceOfBirth struct {
	BirthDt     common.ISODate     `xml:"BirthDt"`
	PrvcOfBirth *common.Max35Text  `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"CtryOfBirth"`
}

func (r DateAndPlaceOfBirth) Validate() error {
	return utils.Validate(&r)
}

type DatePeriodDetails struct {
	FrDt common.ISODate `xml:"FrDt"`
	ToDt common.ISODate `xml:"ToDt"`
}

func (r DatePeriodDetails) Validate() error {
	return utils.Validate(&r)
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount   `xml:"Amt"`
	CcyOfTrf common.ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`
}

func (r EquivalentAmount2) Validate() error {
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

type Frequency21Choice struct {
	Tp  Frequency6Code   `xml:"Tp"`
	Prd FrequencyPeriod1 `xml:"Prd"`
}

func (r Frequency21Choice) Validate() error {
	return utils.Validate(&r)
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"Tp"`
	CntPerPrd float64        `xml:"CntPerPrd"`
}

func (r FrequencyPeriod1) Validate() error {
	return utils.Validate(&r)
}

type Garnishment1 struct {
	Tp                GarnishmentType1                   `xml:"Tp"`
	Grnshee           *PartyIdentification43             `xml:"Grnshee,omitempty" json:",omitempty"`
	GrnshmtAdmstr     *PartyIdentification43             `xml:"GrnshmtAdmstr,omitempty" json:",omitempty"`
	RefNb             *common.Max140Text                 `xml:"RefNb,omitempty" json:",omitempty"`
	Dt                *common.ISODate                    `xml:"Dt,omitempty" json:",omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:",omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"FmlyMdclInsrncInd,omitempty" json:",omitempty"`
	MplyeeTermntnInd  bool                               `xml:"MplyeeTermntnInd,omitempty" json:",omitempty"`
}

func (r Garnishment1) Validate() error {
	return utils.Validate(&r)
}

type GenericIdentification3 struct {
	Id   common.Max35Text  `xml:"Id"`
	Issr *common.Max35Text `xml:"Issr,omitempty" json:",omitempty"`
}

func (r GenericIdentification3) Validate() error {
	return utils.Validate(&r)
}

type InvestigationStatus3Choice struct {
	Conf           InvestigationExecutionConfirmation3Code `xml:"Conf"`
	RjctdMod       []ModificationRejection2Code            `xml:"RjctdMod"`
	DplctOf        Case3                                   `xml:"DplctOf"`
	AssgnmtCxlConf bool                                    `xml:"AssgnmtCxlConf"`
}

func (r InvestigationStatus3Choice) Validate() error {
	return utils.Validate(&r)
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"Cd"`
	Prtry common.Max35Text             `xml:"Prtry"`
}

func (r LocalInstrument2Choice) Validate() error {
	return utils.Validate(&r)
}

type MandateRelatedInformation10 struct {
	MndtId        *common.Max35Text              `xml:"MndtId,omitempty" json:",omitempty"`
	DtOfSgntr     *common.ISODate                `xml:"DtOfSgntr,omitempty" json:",omitempty"`
	AmdmntInd     bool                           `xml:"AmdmntInd,omitempty" json:",omitempty"`
	AmdmntInfDtls *AmendmentInformationDetails10 `xml:"AmdmntInfDtls,omitempty" json:",omitempty"`
	ElctrncSgntr  *common.Max1025Text            `xml:"ElctrncSgntr,omitempty" json:",omitempty"`
	FrstColltnDt  *common.ISODate                `xml:"FrstColltnDt,omitempty" json:",omitempty"`
	FnlColltnDt   *common.ISODate                `xml:"FnlColltnDt,omitempty" json:",omitempty"`
	Frqcy         *Frequency21Choice             `xml:"Frqcy,omitempty" json:",omitempty"`
	Rsn           *MandateSetupReason1Choice     `xml:"Rsn,omitempty" json:",omitempty"`
}

func (r MandateRelatedInformation10) Validate() error {
	return utils.Validate(&r)
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"Cd"`
	Prtry common.Max70Text                `xml:"Prtry"`
}

func (r MandateSetupReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type NumberOfCancellationsPerStatus1 struct {
	DtldNbOfTxs common.Max15NumericText           `xml:"DtldNbOfTxs"`
	DtldSts     CancellationIndividualStatus1Code `xml:"DtldSts"`
	DtldCtrlSum float64                           `xml:"DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfCancellationsPerStatus1) Validate() error {
	return utils.Validate(&r)
}

type NumberOfTransactionsPerStatus1 struct {
	DtldNbOfTxs common.Max15NumericText          `xml:"DtldNbOfTxs"`
	DtldSts     TransactionIndividualStatus1Code `xml:"DtldSts"`
	DtldCtrlSum float64                          `xml:"DtldCtrlSum,omitempty" json:",omitempty"`
}

func (r NumberOfTransactionsPerStatus1) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification8 struct {
	AnyBIC *common.AnyBICIdentifier             `xml:"AnyBIC,omitempty" json:",omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification8) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupHeader5 struct {
	OrgnlGrpCxlId    *common.Max35Text                `xml:"OrgnlGrpCxlId,omitempty" json:",omitempty"`
	RslvdCase        *Case3                           `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlMsgId       common.Max35Text                 `xml:"OrgnlMsgId"`
	OrgnlMsgNmId     common.Max35Text                 `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm     *common.ISODateTime              `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
	OrgnlNbOfTxs     *common.Max15NumericText         `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum     float64                          `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	GrpCxlSts        *GroupCancellationStatus1Code    `xml:"GrpCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf     []CancellationStatusReason2      `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerCxlSts []NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty" json:",omitempty"`
}

func (r OriginalGroupHeader5) Validate() error {
	return utils.Validate(&r)
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   common.Max35Text    `xml:"OrgnlMsgId"`
	OrgnlMsgNmId common.Max35Text    `xml:"OrgnlMsgNmId"`
	OrgnlCreDtTm *common.ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalGroupInformation3) Validate() error {
	return utils.Validate(&r)
}

type OriginalPaymentInstruction17 struct {
	OrgnlPmtInfCxlId *common.Max35Text                 `xml:"OrgnlPmtInfCxlId,omitempty" json:",omitempty"`
	RslvdCase        *Case3                            `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlPmtInfId    common.Max35Text                  `xml:"OrgnlPmtInfId"`
	OrgnlGrpInf      *OriginalGroupInformation3        `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlNbOfTxs     *common.Max15NumericText          `xml:"OrgnlNbOfTxs,omitempty" json:",omitempty"`
	OrgnlCtrlSum     float64                           `xml:"OrgnlCtrlSum,omitempty" json:",omitempty"`
	PmtInfCxlSts     *GroupCancellationStatus1Code     `xml:"PmtInfCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf     []CancellationStatusReason2       `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	NbOfTxsPerCxlSts []NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty" json:",omitempty"`
	TxInfAndSts      []PaymentTransaction66            `xml:"TxInfAndSts,omitempty" json:",omitempty"`
}

func (r OriginalPaymentInstruction17) Validate() error {
	return utils.Validate(&r)
}

type OriginalTransactionReference22 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	Amt            *AmountType4Choice                            `xml:"Amt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                               `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ReqdColltnDt   *common.ISODate                               `xml:"ReqdColltnDt,omitempty" json:",omitempty"`
	ReqdExctnDt    *common.ISODate                               `xml:"ReqdExctnDt,omitempty" json:",omitempty"`
	CdtrSchmeId    *PartyIdentification43                        `xml:"CdtrSchmeId,omitempty" json:",omitempty"`
	SttlmInf       *SettlementInstruction4                       `xml:"SttlmInf,omitempty" json:",omitempty"`
	PmtTpInf       *PaymentTypeInformation25                     `xml:"PmtTpInf,omitempty" json:",omitempty"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:",omitempty"`
	MndtRltdInf    *MandateRelatedInformation10                  `xml:"MndtRltdInf,omitempty" json:",omitempty"`
	RmtInf         *RemittanceInformation11                      `xml:"RmtInf,omitempty" json:",omitempty"`
	UltmtDbtr      *PartyIdentification43                        `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr           *PartyIdentification43                        `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrAcct       *CashAccount24                                `xml:"DbtrAcct,omitempty" json:",omitempty"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:",omitempty"`
	DbtrAgtAcct    *CashAccount24                                `xml:"DbtrAgtAcct,omitempty" json:",omitempty"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:",omitempty"`
	CdtrAgtAcct    *CashAccount24                                `xml:"CdtrAgtAcct,omitempty" json:",omitempty"`
	Cdtr           *PartyIdentification43                        `xml:"Cdtr,omitempty" json:",omitempty"`
	CdtrAcct       *CashAccount24                                `xml:"CdtrAcct,omitempty" json:",omitempty"`
	UltmtCdtr      *PartyIdentification43                        `xml:"UltmtCdtr,omitempty" json:",omitempty"`
}

func (r OriginalTransactionReference22) Validate() error {
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

type PaymentTransaction66 struct {
	CxlStsId          *common.Max35Text                  `xml:"CxlStsId,omitempty" json:",omitempty"`
	RslvdCase         *Case3                             `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlInstrId      *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId   *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	TxCxlSts          *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf      []CancellationStatusReason2        `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	OrgnlInstdAmt     *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty" json:",omitempty"`
	OrgnlReqdExctnDt  *common.ISODate                    `xml:"OrgnlReqdExctnDt,omitempty" json:",omitempty"`
	OrgnlReqdColltnDt *common.ISODate                    `xml:"OrgnlReqdColltnDt,omitempty" json:",omitempty"`
	OrgnlTxRef        *OriginalTransactionReference22    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r PaymentTransaction66) Validate() error {
	return utils.Validate(&r)
}

type PaymentTransaction67 struct {
	CxlStsId            *common.Max35Text                  `xml:"CxlStsId,omitempty" json:",omitempty"`
	RslvdCase           *Case3                             `xml:"RslvdCase,omitempty" json:",omitempty"`
	OrgnlGrpInf         *OriginalGroupInformation3         `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlInstrId        *common.Max35Text                  `xml:"OrgnlInstrId,omitempty" json:",omitempty"`
	OrgnlEndToEndId     *common.Max35Text                  `xml:"OrgnlEndToEndId,omitempty" json:",omitempty"`
	OrgnlTxId           *common.Max35Text                  `xml:"OrgnlTxId,omitempty" json:",omitempty"`
	OrgnlClrSysRef      *common.Max35Text                  `xml:"OrgnlClrSysRef,omitempty" json:",omitempty"`
	TxCxlSts            *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty" json:",omitempty"`
	CxlStsRsnInf        []CancellationStatusReason2        `xml:"CxlStsRsnInf,omitempty" json:",omitempty"`
	RsltnRltdInf        *ResolutionInformation1            `xml:"RsltnRltdInf,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:",omitempty"`
	OrgnlIntrBkSttlmDt  *common.ISODate                    `xml:"OrgnlIntrBkSttlmDt,omitempty" json:",omitempty"`
	Assgnr              *Party12Choice                     `xml:"Assgnr,omitempty" json:",omitempty"`
	Assgne              *Party12Choice                     `xml:"Assgne,omitempty" json:",omitempty"`
	OrgnlTxRef          *OriginalTransactionReference22    `xml:"OrgnlTxRef,omitempty" json:",omitempty"`
}

func (r PaymentTransaction67) Validate() error {
	return utils.Validate(&r)
}

type PaymentTypeInformation25 struct {
	InstrPrty *Priority2Code          `xml:"InstrPrty,omitempty" json:",omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"ClrChanl,omitempty" json:",omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl,omitempty" json:",omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:",omitempty"`
	SeqTp     *SequenceType3Code      `xml:"SeqTp,omitempty" json:",omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:",omitempty"`
}

func (r PaymentTypeInformation25) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth           `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification5) Validate() error {
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

type RemittanceInformation11 struct {
	Ustrd []common.Max140Text                 `xml:"Ustrd,omitempty" json:",omitempty"`
	Strd  []StructuredRemittanceInformation13 `xml:"Strd,omitempty" json:",omitempty"`
}

func (r RemittanceInformation11) Validate() error {
	return utils.Validate(&r)
}

type ResolutionInformation1 struct {
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:",omitempty"`
	IntrBkSttlmDt  *common.ISODate                    `xml:"IntrBkSttlmDt,omitempty" json:",omitempty"`
	ClrChanl       *ClearingChannel2Code              `xml:"ClrChanl,omitempty" json:",omitempty"`
}

func (r ResolutionInformation1) Validate() error {
	return utils.Validate(&r)
}

type ResolutionOfInvestigationV06 struct {
	XMLName      *xml.Name                     `json:",omitempty"`
	Attr         *utils.Attr                   `xml:",attr,omitempty" json:",omitempty"`
	Assgnmt      CaseAssignment3               `xml:"Assgnmt"`
	RslvdCase    *Case3                        `xml:"RslvdCase,omitempty" json:",omitempty"`
	Sts          InvestigationStatus3Choice    `xml:"Sts"`
	CxlDtls      []UnderlyingTransaction14     `xml:"CxlDtls,omitempty" json:",omitempty"`
	StmtDtls     *StatementResolutionEntry2    `xml:"StmtDtls,omitempty" json:",omitempty"`
	CrrctnTx     *CorrectiveTransaction1Choice `xml:"CrrctnTx,omitempty" json:",omitempty"`
	RsltnRltdInf *ResolutionInformation1       `xml:"RsltnRltdInf,omitempty" json:",omitempty"`
	SplmtryData  []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ResolutionOfInvestigationV06) Validate() error {
	return utils.Validate(&r)
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceLevel8Choice) Validate() error {
	return utils.Validate(&r)
}

type SettlementInstruction4 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"SttlmMtd"`
	SttlmAcct            *CashAccount24                                `xml:"SttlmAcct,omitempty" json:",omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty" json:",omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty" json:",omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount24                                `xml:"InstgRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty" json:",omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount24                                `xml:"InstdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty" json:",omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount24                                `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:",omitempty"`
}

func (r SettlementInstruction4) Validate() error {
	return utils.Validate(&r)
}

type StatementResolutionEntry2 struct {
	OrgnlGrpInf *OriginalGroupInformation3         `xml:"OrgnlGrpInf,omitempty" json:",omitempty"`
	OrgnlStmtId *common.Max35Text                  `xml:"OrgnlStmtId,omitempty" json:",omitempty"`
	AcctSvcrRef *common.Max35Text                  `xml:"AcctSvcrRef,omitempty" json:",omitempty"`
	CrrctdAmt   *ActiveOrHistoricCurrencyAndAmount `xml:"CrrctdAmt,omitempty" json:",omitempty"`
	Chrgs       []Charges3                         `xml:"Chrgs,omitempty" json:",omitempty"`
	Purp        *Purpose2Choice                    `xml:"Purp,omitempty" json:",omitempty"`
}

func (r StatementResolutionEntry2) Validate() error {
	return utils.Validate(&r)
}

type StructuredRemittanceInformation13 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:",omitempty"`
	RfrdDocAmt  *RemittanceAmount2             `xml:"RfrdDocAmt,omitempty" json:",omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:",omitempty"`
	Invcr       *PartyIdentification43         `xml:"Invcr,omitempty" json:",omitempty"`
	Invcee      *PartyIdentification43         `xml:"Invcee,omitempty" json:",omitempty"`
	TaxRmt      *TaxInformation4               `xml:"TaxRmt,omitempty" json:",omitempty"`
	GrnshmtRmt  *Garnishment1                  `xml:"GrnshmtRmt,omitempty" json:",omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"AddtlRmtInf,omitempty" json:",omitempty"`
}

func (r StructuredRemittanceInformation13) Validate() error {
	return utils.Validate(&r)
}

type TaxAmount1 struct {
	Rate         float64                            `xml:"Rate,omitempty" json:",omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:",omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:",omitempty"`
	Dtls         []TaxRecordDetails1                `xml:"Dtls,omitempty" json:",omitempty"`
}

func (r TaxAmount1) Validate() error {
	return utils.Validate(&r)
}

type TaxCharges2 struct {
	Id   *common.Max35Text                  `xml:"Id,omitempty" json:",omitempty"`
	Rate float64                            `xml:"Rate,omitempty" json:",omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:",omitempty"`
}

func (r TaxCharges2) Validate() error {
	return utils.Validate(&r)
}

type TaxInformation4 struct {
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
	Rcrd            []TaxRecord1                       `xml:"Rcrd,omitempty" json:",omitempty"`
}

func (r TaxInformation4) Validate() error {
	return utils.Validate(&r)
}

type TaxPeriod1 struct {
	Yr     *common.ISODate       `xml:"Yr,omitempty" json:",omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:",omitempty"`
	FrToDt *DatePeriodDetails    `xml:"FrToDt,omitempty" json:",omitempty"`
}

func (r TaxPeriod1) Validate() error {
	return utils.Validate(&r)
}

type TaxRecord1 struct {
	Tp       *common.Max35Text  `xml:"Tp,omitempty" json:",omitempty"`
	Ctgy     *common.Max35Text  `xml:"Ctgy,omitempty" json:",omitempty"`
	CtgyDtls *common.Max35Text  `xml:"CtgyDtls,omitempty" json:",omitempty"`
	DbtrSts  *common.Max35Text  `xml:"DbtrSts,omitempty" json:",omitempty"`
	CertId   *common.Max35Text  `xml:"CertId,omitempty" json:",omitempty"`
	FrmsCd   *common.Max35Text  `xml:"FrmsCd,omitempty" json:",omitempty"`
	Prd      *TaxPeriod1        `xml:"Prd,omitempty" json:",omitempty"`
	TaxAmt   *TaxAmount1        `xml:"TaxAmt,omitempty" json:",omitempty"`
	AddtlInf *common.Max140Text `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r TaxRecord1) Validate() error {
	return utils.Validate(&r)
}

type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                       `xml:"Prd,omitempty" json:",omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`
}

func (r TaxRecordDetails1) Validate() error {
	return utils.Validate(&r)
}

type UnderlyingTransaction14 struct {
	OrgnlGrpInfAndSts *OriginalGroupHeader5          `xml:"OrgnlGrpInfAndSts,omitempty" json:",omitempty"`
	OrgnlPmtInfAndSts []OriginalPaymentInstruction17 `xml:"OrgnlPmtInfAndSts,omitempty" json:",omitempty"`
	TxInfAndSts       []PaymentTransaction67         `xml:"TxInfAndSts,omitempty" json:",omitempty"`
}

func (r UnderlyingTransaction14) Validate() error {
	return utils.Validate(&r)
}
