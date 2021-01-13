// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v08

import "github.com/moov-io/iso20022/pkg/common"

type AccountIdentification4Choice struct {
	IBAN common.IBAN2007Identifier     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64                   `xml:",chardata"`
	Ccy   common.ActiveCurrencyCode `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                             `xml:",chardata"`
	Ccy   common.ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type AmendmentInformationDetails13 struct {
	OrgnlMndtId      common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlMndtId,omitempty"`
	OrgnlCdtrSchmeId PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlCdtrSchmeId,omitempty"`
	OrgnlCdtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlCdtrAgt,omitempty"`
	OrgnlCdtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlCdtrAgtAcct,omitempty"`
	OrgnlDbtr        PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlDbtr,omitempty"`
	OrgnlDbtrAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlDbtrAcct,omitempty"`
	OrgnlDbtrAgt     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlDbtrAgt,omitempty"`
	OrgnlDbtrAgtAcct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlDbtrAgtAcct,omitempty"`
	OrgnlFnlColltnDt common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlFnlColltnDt,omitempty"`
	OrgnlFrqcy       Frequency36Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlFrqcy,omitempty"`
	OrgnlRsn         MandateSetupReason1Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlRsn,omitempty"`
	OrgnlTrckgDays   common.Exact2NumericText                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgnlTrckgDays,omitempty"`
}

type Authorisation1Choice struct {
	Cd    common.Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id,omitempty"`
	LEI     common.LEIIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 LEI,omitempty"`
	Nm      common.Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
	PstlAdr PostalAddress24      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	Tp   CashAccountType2Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Ccy  common.ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ccy,omitempty"`
	Nm   common.Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
	Prxy ProxyAccountIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type Charges7 struct {
	Amt ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Agt"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ClrSysId,omitempty"`
	MmbId    common.Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 MmbId"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
	PhneNb    common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PrefrdMtd,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Ref common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdOrPrtry"`
	Issr      common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FrDt"`
	ToDt common.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ToDt"`
}

type DirectDebitTransaction10 struct {
	MndtRltdInf MandateRelatedInformation14 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 MndtRltdInf,omitempty"`
	CdtrSchmeId PartyIdentification135      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtrSchmeId,omitempty"`
	PreNtfctnId common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PreNtfctnId,omitempty"`
	PreNtfctnDt common.ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PreNtfctnDt,omitempty"`
}

type DirectDebitTransactionInformation24 struct {
	PmtId          PaymentIdentification7                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PmtId"`
	PmtTpInf       PaymentTypeInformation27                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PmtTpInf,omitempty"`
	IntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrBkSttlmAmt"`
	IntrBkSttlmDt  common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrBkSttlmDt,omitempty"`
	SttlmPrty      Priority3Code                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SttlmPrty,omitempty"`
	SttlmTmIndctn  SettlementDateTimeIndication1                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SttlmTmIndctn,omitempty"`
	InstdAmt       ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstdAmt,omitempty"`
	XchgRate       float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 XchgRate,omitempty"`
	ChrgBr         ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ChrgBr"`
	ChrgsInf       []Charges7                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ChrgsInf,omitempty"`
	ReqdColltnDt   common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ReqdColltnDt,omitempty"`
	DrctDbtTx      DirectDebitTransaction10                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DrctDbtTx,omitempty"`
	Cdtr           PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cdtr"`
	CdtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtrAcct,omitempty"`
	CdtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtrAgt"`
	CdtrAgtAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtrAgtAcct,omitempty"`
	UltmtCdtr      PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 UltmtCdtr,omitempty"`
	InitgPty       PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InitgPty,omitempty"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstgAgt,omitempty"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstdAgt,omitempty"`
	IntrmyAgt1     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3     BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrmyAgt3Acct,omitempty"`
	Dbtr           PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dbtr"`
	DbtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DbtrAcct"`
	DbtrAgt        BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DbtrAgt"`
	DbtrAgtAcct    CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DbtrAgtAcct,omitempty"`
	UltmtDbtr      PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 UltmtDbtr,omitempty"`
	Purp           Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Purp,omitempty"`
	RgltryRptg     []RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RgltryRptg,omitempty"`
	RltdRmtInf     []RemittanceLocation7                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RltdRmtInf,omitempty"`
	RmtInf         RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RmtInf,omitempty"`
	SplmtryData    []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SplmtryData,omitempty"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt"`
	CdtDbtInd common.CreditDebitCode            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtDbtInd,omitempty"`
	Rsn       common.Max4Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Rsn,omitempty"`
	AddtlInf  common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Nb     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nb,omitempty"`
	RltdDt common.ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	Desc common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdOrPrtry"`
	Issr      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type FIToFICustomerDirectDebitV08 struct {
	GrpHdr       GroupHeader94                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 GrpHdr"`
	DrctDbtTxInf []DirectDebitTransactionInformation24 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DrctDbtTxInf"`
	SplmtryData  []SupplementaryData1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SplmtryData,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       common.BICFIDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ClrSysMmbId,omitempty"`
	LEI         common.LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 LEI,omitempty"`
	Nm          common.Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp"`
	Prd    FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prd"`
	PtInTm FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PtInTm"`
}

type FrequencyAndMoment1 struct {
	Tp     Frequency6Code           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp"`
	PtInTm common.Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp"`
	CntPerPrd float64        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CntPerPrd"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 GrnshmtAdmstr,omitempty"`
	RefNb             common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RefNb,omitempty"`
	Dt                common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdOrPrtry"`
	Issr      common.Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      common.Max34Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	Issr    common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr"`
	SchmeNm common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type GroupHeader94 struct {
	MsgId             common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 MsgId"`
	CreDtTm           common.ISODateTime                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CreDtTm"`
	Authstn           []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Authstn,omitempty"`
	BtchBookg         bool                                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 BtchBookg,omitempty"`
	NbOfTxs           common.Max15NumericText                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 NbOfTxs"`
	CtrlSum           float64                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtrlSum,omitempty"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TtlIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt     common.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 IntrBkSttlmDt,omitempty"`
	SttlmInf          SettlementInstruction8                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SttlmInf"`
	PmtTpInf          PaymentTypeInformation27                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PmtTpInf,omitempty"`
	InstgAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstgAgt,omitempty"`
	InstdAgt          BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstdAgt,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type MandateRelatedInformation14 struct {
	MndtId        common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 MndtId,omitempty"`
	DtOfSgntr     common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DtOfSgntr,omitempty"`
	AmdmntInd     bool                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AmdmntInd,omitempty"`
	AmdmntInfDtls AmendmentInformationDetails13 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AmdmntInfDtls,omitempty"`
	ElctrncSgntr  common.Max1025Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ElctrncSgntr,omitempty"`
	FrstColltnDt  common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FrstColltnDt,omitempty"`
	FnlColltnDt   common.ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FnlColltnDt,omitempty"`
	Frqcy         Frequency36Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Frqcy,omitempty"`
	Rsn           MandateSetupReason1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Rsn,omitempty"`
	TrckgDays     common.Exact2NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TrckgDays,omitempty"`
}

type MandateSetupReason1Choice struct {
	Cd    ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type NameAndAddress16 struct {
	Nm  common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm"`
	Adr PostalAddress24   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Adr"`
}

type OrganisationIdentification29 struct {
	AnyBIC common.AnyBICDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AnyBIC,omitempty"`
	LEI    common.LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
	PstlAdr   PostalAddress24    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PstlAdr,omitempty"`
	Id        Party38Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtryOfRes,omitempty"`
	CtctDtls  Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtctDtls,omitempty"`
}

type PaymentIdentification7 struct {
	InstrId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstrId,omitempty"`
	EndToEndId common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 EndToEndId"`
	TxId       common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TxId,omitempty"`
	UETR       common.UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 UETR,omitempty"`
	ClrSysRef  common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ClrSysRef,omitempty"`
}

type PaymentTypeInformation27 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 InstrPrty,omitempty"`
	ClrChanl  ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ClrChanl,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 LclInstrm,omitempty"`
	SeqTp     SequenceType3Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SeqTp,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Id common.Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type Purpose2Choice struct {
	Cd    ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Nb       common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nb,omitempty"`
	RltdDt   common.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdOrPrtry"`
	Issr      common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
	Ctry common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DbtCdtRptgInd,omitempty"`
	Authrty       RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Authrty,omitempty"`
	Dtls          []StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dtls,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ustrd,omitempty"`
	Strd  []StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Mtd"`
	ElctrncAdr common.Max2048Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PstlAdr,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type SettlementDateTimeIndication1 struct {
	DbtDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DbtDtTm,omitempty"`
	CdtDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtDtTm,omitempty"`
}

type SettlementInstruction8 struct {
	SttlmMtd  SettlementMethod2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SttlmMtd"`
	SttlmAcct CashAccount38                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SttlmAcct,omitempty"`
	ClrSys    ClearingSystemIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 ClrSys,omitempty"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Dt   common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dt,omitempty"`
	Ctry common.CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ctry,omitempty"`
	Cd   common.Max10Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd,omitempty"`
	Amt  ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt,omitempty"`
	Inf  []common.Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Inf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 GrnshmtRmt,omitempty"`
	AddtlRmtInf []common.Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cd"`
	Prtry common.Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Titl,omitempty"`
	Nm   common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 UltmtDbtr,omitempty"`
	AdmstnZone      common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AdmstnZone,omitempty"`
	RefNb           common.Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RefNb,omitempty"`
	Mtd             common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TtlTaxAmt,omitempty"`
	Dt              common.ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxId,omitempty"`
	RegnId common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RegnId,omitempty"`
	TaxTp  common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxId,omitempty"`
	RegnId  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 RegnId,omitempty"`
	TaxTp   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     common.ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Tp,omitempty"`
	Ctgy     common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Ctgy,omitempty"`
	CtgyDtls common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CtgyDtls,omitempty"`
	DbtrSts  common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 DbtrSts,omitempty"`
	CertId   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 CertId,omitempty"`
	FrmsCd   common.Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 FrmsCd,omitempty"`
	Prd      TaxPeriod2        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prd,omitempty"`
	TaxAmt   TaxAmount2        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 TaxAmt,omitempty"`
	AddtlInf common.Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.003.001.08 Amt"`
}
