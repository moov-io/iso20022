package remt_v02

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Prtry"`
}

type Authorisation1Choice struct {
	Cd    Authorisation1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Cd"`
	Prtry Max128Text         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Prtry"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PrefrdMtd,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CtryOfBirth"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Issr,omitempty"`
}

type GroupHeader79 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CreDtTm"`
	Authstn  []Authorisation1Choice                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Authstn,omitempty"`
	CpyInd   CopyDuplicate1Code                           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CpyInd,omitempty"`
	InitgPty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 InitgPty"`
	MsgRcpt  PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 MsgRcpt,omitempty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 FwdgAgt,omitempty"`
}

type NameAndAddress16 struct {
	Nm  Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Nm"`
	Adr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Adr"`
}

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PrvtId"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Prtry"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 AdrLine,omitempty"`
}

type RemittanceLocation5 struct {
	RmtId       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 RmtId,omitempty"`
	RmtLctnDtls []RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 RmtLctnDtls"`
	Refs        TransactionReferences5    `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Refs"`
}

type RemittanceLocationAdviceV02 struct {
	GrpHdr      GroupHeader79         `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 GrpHdr"`
	RmtLctn     []RemittanceLocation5 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 RmtLctn"`
	SplmtryData []SupplementaryData1  `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 SplmtryData,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Mtd"`
	ElctrncAdr Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 ElctrncAdr,omitempty"`
	PstlAdr    NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PstlAdr,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TransactionReferences5 struct {
	PmtInfId    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 PmtInfId,omitempty"`
	InstrId     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 InstrId,omitempty"`
	EndToEndId  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 EndToEndId"`
	TxId        Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 TxId,omitempty"`
	UETR        UUIDv4Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 UETR,omitempty"`
	MndtId      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 MndtId,omitempty"`
	CdtrSchmeId PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.02 CdtrSchmeId,omitempty"`
}
