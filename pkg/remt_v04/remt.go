package remt_v04

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type AmountType3Choice struct {
	InstdAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 InstdAmt"`
	EqvtAmt  EquivalentAmount2                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 EqvtAmt"`
}

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PrefrdMtd,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Ref Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdOrPrtry"`
	Issr      Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type DocumentAdjustment1 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt"`
	CdtDbtInd CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtDbtInd,omitempty"`
	Rsn       Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Rsn,omitempty"`
	AddtlInf  Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Nb     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nb,omitempty"`
	RltdDt ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	Desc Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Desc,omitempty"`
	Amt  RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdOrPrtry"`
	Issr      Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type EquivalentAmount2 struct {
	Amt      ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt"`
	CcyOfTrf ActiveOrHistoricCurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CcyOfTrf"`
}

type ExchangeRate1 struct {
	UnitCcy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 UnitCcy,omitempty"`
	XchgRate float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 XchgRate,omitempty"`
	RateTp   ExchangeRateType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RateTp,omitempty"`
	CtrctId  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtrctId,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Othr,omitempty"`
}

type Garnishment3 struct {
	Tp                GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp"`
	Grnshee           PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Grnshee,omitempty"`
	GrnshmtAdmstr     PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 GrnshmtAdmstr,omitempty"`
	RefNb             Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RefNb,omitempty"`
	Dt                ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dt,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdOrPrtry"`
	Issr      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type GroupHeader79 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Authstn,omitempty"`
	CpyInd   CopyDuplicate1Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CpyInd,omitempty"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 InitgPty"`
	MsgRcpt  PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 MsgRcpt,omitempty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FwdgAgt,omitempty"`
}

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type OriginalPaymentInformation8 struct {
	Refs         TransactionReferences5                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Refs"`
	PmtTpInf     PaymentTypeInformation26                     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PmtTpInf,omitempty"`
	Amt          AmountType3Choice                            `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt,omitempty"`
	XchgRateInf  ExchangeRate1                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 XchgRateInf,omitempty"`
	ReqdExctnDt  DateAndDateTime2Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 ReqdExctnDt,omitempty"`
	ReqdColltnDt ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 ReqdColltnDt,omitempty"`
	Dbtr         PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dbtr,omitempty"`
	DbtrAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DbtrAcct,omitempty"`
	DbtrAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DbtrAgt,omitempty"`
	Cdtr         PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cdtr,omitempty"`
	CdtrAcct     CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtrAcct,omitempty"`
	CdtrAgt      BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtrAgt,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtctDtls,omitempty"`
}

type PaymentTypeInformation26 struct {
	InstrPrty Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 InstrPrty,omitempty"`
	SvcLvl    []ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SvcLvl,omitempty"`
	LclInstrm LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 LclInstrm,omitempty"`
	CtgyPurp  CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type ReferredDocumentInformation7 struct {
	Tp       ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Nb       Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nb,omitempty"`
	RltdDt   ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RltdDt,omitempty"`
	LineDtls []DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdOrPrtry"`
	Issr      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Issr,omitempty"`
}

type RemittanceAdviceV04 struct {
	GrpHdr      GroupHeader79             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 GrpHdr"`
	RmtInf      []RemittanceInformation19 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtInf"`
	SplmtryData []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SplmtryData,omitempty"`
}

type RemittanceAmount2 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DuePyblAmt,omitempty"`
	DscntApldAmt      []DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DscntApldAmt,omitempty"`
	CdtNoteAmt        ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtNoteAmt,omitempty"`
	TaxAmt            []TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtdAmt,omitempty"`
}

type RemittanceInformation19 struct {
	RmtId       Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RmtId,omitempty"`
	Ustrd       []Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Ustrd,omitempty"`
	Strd        []StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Strd,omitempty"`
	OrgnlPmtInf OriginalPaymentInformation8         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 OrgnlPmtInf"`
}

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RfrdDocInf,omitempty"`
	RfrdDocAmt  RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RfrdDocAmt,omitempty"`
	CdtrRefInf  CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtrRefInf,omitempty"`
	Invcr       PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Invcr,omitempty"`
	Invcee      PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Invcee,omitempty"`
	TaxRmt      TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxRmt,omitempty"`
	GrnshmtRmt  Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 GrnshmtRmt,omitempty"`
	AddtlRmtInf []Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TaxAmount2 struct {
	Rate         float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Rate,omitempty"`
	TaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxblBaseAmt,omitempty"`
	TtlAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TtlAmt,omitempty"`
	Dtls         []TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt"`
}

type TaxAmountType1Choice struct {
	Cd    ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cd"`
	Prtry Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prtry"`
}

type TaxAuthorisation1 struct {
	Titl Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Titl,omitempty"`
	Nm   Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Cdtr,omitempty"`
	Dbtr            TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dbtr,omitempty"`
	UltmtDbtr       TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 UltmtDbtr,omitempty"`
	AdmstnZone      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AdmstnZone,omitempty"`
	RefNb           Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RefNb,omitempty"`
	Mtd             Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Mtd,omitempty"`
	TtlTaxblBaseAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TtlTaxAmt,omitempty"`
	Dt              ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Dt,omitempty"`
	SeqNb           float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 SeqNb,omitempty"`
	Rcrd            []TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxId,omitempty"`
	RegnId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RegnId,omitempty"`
	TaxTp  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxId,omitempty"`
	RegnId  Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 RegnId,omitempty"`
	TaxTp   Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxTp,omitempty"`
	Authstn TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Yr,omitempty"`
	Tp     TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	FrToDt DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Tp,omitempty"`
	Ctgy     Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Ctgy,omitempty"`
	CtgyDtls Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CtgyDtls,omitempty"`
	DbtrSts  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 DbtrSts,omitempty"`
	CertId   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CertId,omitempty"`
	FrmsCd   Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 FrmsCd,omitempty"`
	Prd      TaxPeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prd,omitempty"`
	TaxAmt   TaxAmount2 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TaxAmt,omitempty"`
	AddtlInf Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Prd,omitempty"`
	Amt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 Amt"`
}

type TransactionReferences5 struct {
	PmtInfId    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 PmtInfId,omitempty"`
	InstrId     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 InstrId,omitempty"`
	EndToEndId  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 EndToEndId"`
	TxId        Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 TxId,omitempty"`
	UETR        UUIDv4Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 UETR,omitempty"`
	MndtId      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 MndtId,omitempty"`
	CdtrSchmeId PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.04 CdtrSchmeId,omitempty"`
}
