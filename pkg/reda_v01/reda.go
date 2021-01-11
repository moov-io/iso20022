// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import "github.com/moov-io/iso20022/pkg/common"

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Prtry"`
}

type Contact4 struct {
	NmPrfx    common.NamePrefix2Code      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 NmPrfx,omitempty"`
	Nm        common.Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Nm,omitempty"`
	PhneNb    common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PhneNb,omitempty"`
	MobNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 MobNb,omitempty"`
	FaxNb     common.PhoneNumber          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 FaxNb,omitempty"`
	EmailAdr  common.Max2048Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EmailAdr,omitempty"`
	EmailPurp common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EmailPurp,omitempty"`
	JobTitl   common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 JobTitl,omitempty"`
	Rspnsblty common.Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Rspnsblty,omitempty"`
	Dept      common.Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PrefrdMtd,omitempty"`
}

type CreditorEnrolment3 struct {
	Enrlmnt      CreditorServiceEnrolment1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Enrlmnt"`
	CdtrTradgNm  common.Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CdtrTradgNm,omitempty"`
	Cdtr         RTPPartyIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Cdtr"`
	UltmtCdtr    RTPPartyIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 UltmtCdtr,omitempty"`
	MrchntCtgyCd common.MerchantCategoryCodeIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 MrchntCtgyCd"`
	CdtrLogo     common.Max10KBinary                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CdtrLogo,omitempty"`
}

type CreditorInvoice3 struct {
	LtdPresntmntInd   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 LtdPresntmntInd"`
	CstmrIdTp         CustomerTypeRequest2    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CstmrIdTp,omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CtrctFrmtTp,omitempty"`
	CtrctRefTp        []DocumentType1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CtrctRefTp,omitempty"`
	CdtrInstr         common.Max500Text       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CdtrInstr,omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 ActvtnReqDlvryPty"`
}

type CreditorServiceEnrolment1 struct {
	EnrlmntStartDt  DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EnrlmntStartDt,omitempty"`
	EnrlmntEndDt    DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EnrlmntEndDt,omitempty"`
	Vsblty          Visibilty1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Vsblty,omitempty"`
	SvcActvtnAllwd  bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SvcActvtnAllwd"`
	SvcDescLk       common.Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SvcDescLk,omitempty"`
	CdtrSvcActvtnLk common.Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CdtrSvcActvtnLk,omitempty"`
}

type CustomerTypeRequest2 struct {
	Reqd   bool              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Reqd"`
	OrgTp  OrganisationType2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 OrgTp,omitempty"`
	PrvtTp PersonType2       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PrvtTp,omitempty"`
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Dt"`
	DtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 DtTm"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     common.ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 BirthDt"`
	PrvcOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CityOfBirth"`
	CtryOfBirth common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CtryOfBirth"`
}

type DocumentFormat2Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Cd"`
	Prtry GenericIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Prtry"`
}

type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Cd"`
	Prtry GenericIdentification1    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Prtry"`
}

type EnrolmentHeader2 struct {
	MsgId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 MsgId"`
	CreDtTm  common.ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CreDtTm"`
	MsgOrgtr RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 MsgRcpt,omitempty"`
	InitgPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 InitgPty"`
}

type GenericIdentification1 struct {
	Id      common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Id"`
	SchmeNm common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Id"`
	Issr    common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Issr"`
	SchmeNm common.Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Issr,omitempty"`
}

type GenericOrganisationType1 struct {
	Reqd    bool                                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Reqd"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SchmeNm"`
}

type GenericPersonIdentification1 struct {
	Id      common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SchmeNm,omitempty"`
	Issr    common.Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Issr,omitempty"`
}

type GenericPersonType1 struct {
	Reqd    bool                                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Reqd"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SchmeNm"`
}

type OrganisationIdentification37 struct {
	AnyBIC   common.AnyBICDec2014Identifier       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 AnyBIC,omitempty"`
	LEI      common.LEIIdentifier                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 LEI,omitempty"`
	EmailAdr common.Max256Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EmailAdr,omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Cd"`
	Prtry common.Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Prtry"`
}

type OrganisationType2 struct {
	AnyBIC   bool                       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 AnyBIC,omitempty"`
	LEI      bool                       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 LEI,omitempty"`
	EmailAdr bool                       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EmailAdr,omitempty"`
	Othr     []GenericOrganisationType1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Othr,omitempty"`
}

type OtherContact1 struct {
	ChanlTp common.Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 ChanlTp"`
	Id      common.Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Id,omitempty"`
}

type Party49Choice struct {
	OrgId  OrganisationIdentification37 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 OrgId"`
	PrvtId PersonIdentification17       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PrvtId"`
}

type PersonIdentification17 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 DtAndPlcOfBirth,omitempty"`
	EmailAdr        common.Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EmailAdr,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Cd"`
	Prtry common.Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Prtry"`
}

type PersonType2 struct {
	DtAndPlcOfBirth bool                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 DtAndPlcOfBirth,omitempty"`
	EmailAdr        bool                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EmailAdr,omitempty"`
	Othr            []GenericPersonType1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Othr,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 AdrTp,omitempty"`
	Dept        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Dept,omitempty"`
	SubDept     common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SubDept,omitempty"`
	StrtNm      common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 StrtNm,omitempty"`
	BldgNb      common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 BldgNb,omitempty"`
	BldgNm      common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 BldgNm,omitempty"`
	Flr         common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Flr,omitempty"`
	PstBx       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PstBx,omitempty"`
	Room        common.Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Room,omitempty"`
	PstCd       common.Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PstCd,omitempty"`
	TwnNm       common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 TwnNm,omitempty"`
	TwnLctnNm   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 TwnLctnNm,omitempty"`
	DstrctNm    common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 DstrctNm,omitempty"`
	CtrySubDvsn common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CtrySubDvsn,omitempty"`
	Ctry        common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Ctry,omitempty"`
	AdrLine     []common.Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 AdrLine,omitempty"`
}

type RTPPartyIdentification1 struct {
	Nm        common.Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress24    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PstlAdr,omitempty"`
	Id        Party49Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Id,omitempty"`
	CtryOfRes common.CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CtryOfRes,omitempty"`
	CtctDtls  Contact4           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CtctDtls,omitempty"`
}

type RequestToPayCreditorEnrolmentRequestV01 struct {
	Hdr         EnrolmentHeader2     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Hdr"`
	CdtrEnrlmnt []CreditorEnrolment3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 CdtrEnrlmnt"`
	ActvtnData  CreditorInvoice3     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 ActvtnData"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm common.Max350Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Visibilty1 struct {
	StartDt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 StartDt,omitempty"`
	EndDt     DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 EndDt,omitempty"`
	LtdVsblty bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.066.001.01 LtdVsblty,omitempty"`
}

type CreditorEnrolment4 struct {
	Enrlmnt      CreditorServiceEnrolment1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Enrlmnt,omitempty"`
	CdtrTradgNm  common.Max140Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CdtrTradgNm,omitempty"`
	Cdtr         RTPPartyIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Cdtr"`
	UltmtCdtr    RTPPartyIdentification1               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 UltmtCdtr,omitempty"`
	MrchntCtgyCd common.MerchantCategoryCodeIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 MrchntCtgyCd,omitempty"`
	CdtrLogo     common.Max10KBinary                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CdtrLogo,omitempty"`
}

type CreditorEnrolmentAmendment3 struct {
	OrgnlBizInstr OriginalBusinessInstruction1      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 OrgnlBizInstr,omitempty"`
	AmdmntRsn     CreditorEnrolmentAmendmentReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 AmdmntRsn,omitempty"`
	Amdmnt        CreditorEnrolmentAmendment4       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Amdmnt"`
	OrgnlEnrlmnt  OriginalEnrolment2Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 OrgnlEnrlmnt"`
	SplmtryData   []SupplementaryData1              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 SplmtryData,omitempty"`
}

type CreditorEnrolmentAmendment4 struct {
	CdtrEnrlmnt CreditorEnrolment4 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CdtrEnrlmnt,omitempty"`
	ActvtnData  CreditorInvoice4   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 ActvtnData,omitempty"`
}

type CreditorEnrolmentAmendmentReason1Choice struct {
	Cd    ExternalCreditorEnrolmentAmendmentReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Cd"`
	Prtry common.Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Prtry"`
}

type CreditorEnrolmentAmendmentReason2 struct {
	Orgtr    RTPPartyIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Orgtr,omitempty"`
	Rsn      CreditorEnrolmentAmendmentReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Rsn"`
	AddtlInf []common.Max105Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 AddtlInf,omitempty"`
}

type CreditorInvoice4 struct {
	LtdPresntmntInd   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 LtdPresntmntInd,omitempty"`
	CstmrIdTp         CustomerTypeRequest2    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CstmrIdTp,omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CtrctFrmtTp,omitempty"`
	CtrctRefTp        []DocumentType1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CtrctRefTp,omitempty"`
	CdtrInstr         common.Max500Text       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CdtrInstr,omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 ActvtnReqDlvryPty,omitempty"`
}

type OriginalBusinessInstruction1 struct {
	MsgId   common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 MsgId"`
	MsgNmId common.Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 MsgNmId,omitempty"`
	CreDtTm common.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 CreDtTm,omitempty"`
}

type OriginalEnrolment2Choice struct {
	OrgnlCdtrId      Party49Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 OrgnlCdtrId"`
	OrgnlEnrlmntData CreditorEnrolment3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 OrgnlEnrlmntData"`
}

type RequestToPayCreditorEnrolmentAmendmentRequestV01 struct {
	Hdr         EnrolmentHeader2              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 Hdr"`
	AmdmntData  []CreditorEnrolmentAmendment3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 AmdmntData"`
	SplmtryData []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.067.001.01 SplmtryData,omitempty"`
}

type CreditorEnrolmentCancellation2 struct {
	OrgnlBizInstr OriginalBusinessInstruction1         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgnlBizInstr,omitempty"`
	CxlRsn        CreditorEnrolmentCancellationReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CxlRsn,omitempty"`
	OrgnlEnrlmnt  OriginalEnrolment2Choice             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 OrgnlEnrlmnt"`
	SplmtryData   []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SplmtryData,omitempty"`
}

type CreditorEnrolmentCancellationReason1Choice struct {
	Cd    ExternalCreditorEnrolmentCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Cd"`
	Prtry common.Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Prtry"`
}

type CreditorEnrolmentCancellationReason2 struct {
	Orgtr    RTPPartyIdentification1                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Orgtr,omitempty"`
	Rsn      CreditorEnrolmentCancellationReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Rsn"`
	AddtlInf []common.Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 AddtlInf,omitempty"`
}

type RequestToPayCreditorEnrolmentCancellationRequestV01 struct {
	Hdr         EnrolmentHeader2                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 Hdr"`
	CxlData     []CreditorEnrolmentCancellation2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 CxlData"`
	SplmtryData []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.068.001.01 SplmtryData,omitempty"`
}

type CreditorEnrolmentStatusReason2 struct {
	Orgtr    RTPPartyIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Orgtr,omitempty"`
	Rsn      CreditorEnrolmentStatusReason2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Rsn"`
	AddtlInf []common.Max105Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 AddtlInf,omitempty"`
}

type CreditorEnrolmentStatusReason2Choice struct {
	Cd    ExternalCreditorEnrolmentStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Cd"`
	Prtry common.Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Prtry"`
}

type EnrolmentStatus2 struct {
	OrgnlBizInstr   OriginalBusinessInstruction1   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 OrgnlBizInstr,omitempty"`
	Sts             ServiceStatus1Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Sts"`
	StsRsn          CreditorEnrolmentStatusReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 StsRsn,omitempty"`
	OrgnlEnrlmntRef OriginalEnrolment2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 OrgnlEnrlmntRef,omitempty"`
	FctvEnrlmntDt   DateAndDateTime2Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 FctvEnrlmntDt,omitempty"`
	SplmtryData     []SupplementaryData1           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 SplmtryData,omitempty"`
}

type RequestToPayCreditorEnrolmentStatusReportV01 struct {
	Hdr                EnrolmentHeader2     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Hdr"`
	OrgnlEnrlmntAndSts []EnrolmentStatus2   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 OrgnlEnrlmntAndSts"`
	SplmtryData        []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 SplmtryData,omitempty"`
}

type ServiceStatus1Choice struct {
	Cd    ServiceRequestStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Cd"`
	Prtry common.Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.069.001.01 Prtry"`
}

type ActivationHeader2 struct {
	MsgId    common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 MsgId"`
	CreDtTm  common.ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 CreDtTm"`
	MsgOrgtr RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 MsgOrgtr,omitempty"`
	MsgRcpt  RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 MsgRcpt,omitempty"`
	InitgPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 InitgPty"`
}

type ContractReference1 struct {
	Tp  DocumentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 Tp,omitempty"`
	Ref common.Max500Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 Ref"`
}

type DebtorActivation3 struct {
	DbtrActvtnId      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 DbtrActvtnId,omitempty"`
	DispNm            common.Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 DispNm,omitempty"`
	UltmtDbtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 UltmtDbtr,omitempty"`
	Dbtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 Dbtr"`
	DbtrSolPrvdr      RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 DbtrSolPrvdr"`
	CstmrId           []Party49Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 CstmrId,omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 CtrctFrmtTp,omitempty"`
	CtrctRef          []ContractReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 CtrctRef,omitempty"`
	Cdtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 Cdtr"`
	UltmtCdtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 UltmtCdtr,omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 ActvtnReqDlvryPty,omitempty"`
	StartDt           DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 StartDt,omitempty"`
	EndDt             DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 EndDt,omitempty"`
	DdctdActvtnCd     common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 DdctdActvtnCd,omitempty"`
}

type ElectronicInvoice1 struct {
	PresntmntTp PresentmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 PresntmntTp"`
}

type RequestToPayDebtorActivationRequestV01 struct {
	Hdr             ActivationHeader2    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 Hdr"`
	DbtrActvtn      []DebtorActivation3  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 DbtrActvtn"`
	ElctrncInvcData ElectronicInvoice1   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 ElctrncInvcData"`
	SplmtryData     []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.070.001.01 SplmtryData,omitempty"`
}

type DebtorActivation4 struct {
	DbtrActvtnId      common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrActvtnId,omitempty"`
	DispNm            common.Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DispNm,omitempty"`
	UltmtDbtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 UltmtDbtr,omitempty"`
	Dbtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Dbtr,omitempty"`
	DbtrSolPrvdr      RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrSolPrvdr,omitempty"`
	CstmrId           []Party49Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CstmrId,omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrctFrmtTp,omitempty"`
	CtrctRef          []ContractReference1    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 CtrctRef,omitempty"`
	Cdtr              RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cdtr,omitempty"`
	UltmtCdtr         RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 UltmtCdtr,omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ActvtnReqDlvryPty,omitempty"`
	StartDt           DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 StartDt,omitempty"`
	EndDt             DateAndDateTime2Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 EndDt,omitempty"`
	DdctdActvtnCd     common.Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DdctdActvtnCd,omitempty"`
}

type DebtorActivationAmendment3 struct {
	OrgnlBizInstr OriginalBusinessInstruction1     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgnlBizInstr,omitempty"`
	AmdmntRsn     DebtorActivationAmendmentReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AmdmntRsn,omitempty"`
	Amdmnt        DebtorActivationAmendment4       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Amdmnt"`
	OrgnlActvtn   OriginalActivation2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 OrgnlActvtn"`
	SplmtryData   []SupplementaryData1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SplmtryData,omitempty"`
}

type DebtorActivationAmendment4 struct {
	DbtrActvtn      DebtorActivation4  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 DbtrActvtn,omitempty"`
	ElctrncInvcData ElectronicInvoice1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 ElctrncInvcData,omitempty"`
}

type DebtorActivationAmendmentReason1Choice struct {
	Cd    ExternalDebtorActivationAmendmentReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Cd"`
	Prtry common.Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Prtry"`
}

type DebtorActivationAmendmentReason2 struct {
	Orgtr    RTPPartyIdentification1                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Orgtr,omitempty"`
	Rsn      DebtorActivationAmendmentReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Rsn"`
	AddtlInf []common.Max105Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AddtlInf,omitempty"`
}

type RequestToPayDebtorActivationAmendmentRequestV01 struct {
	Hdr         ActivationHeader2            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 Hdr"`
	AmdmntData  []DebtorActivationAmendment3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 AmdmntData"`
	SplmtryData []SupplementaryData1         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.071.001.01 SplmtryData,omitempty"`
}

type DebtorActivationCancellation2 struct {
	OrgnlBizInstr OriginalBusinessInstruction1        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 OrgnlBizInstr,omitempty"`
	CxlRsn        DebtorActivationCancellationReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 CxlRsn,omitempty"`
	OrgnlActvtn   OriginalActivation2Choice           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 OrgnlActvtn"`
	SplmtryData   []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 SplmtryData,omitempty"`
}

type DebtorActivationCancellationReason1Choice struct {
	Cd    ExternalDebtorActivationCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 Cd"`
	Prtry common.Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 Prtry"`
}

type DebtorActivationCancellationReason2 struct {
	Orgtr    RTPPartyIdentification1                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 Orgtr,omitempty"`
	Rsn      DebtorActivationCancellationReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 Rsn"`
	AddtlInf []common.Max105Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 AddtlInf,omitempty"`
}

type OriginalActivation2Choice struct {
	OrgnlDbtrId     Party49Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 OrgnlDbtrId"`
	OrgnlActvtnData DebtorActivation3 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 OrgnlActvtnData"`
}

type RequestToPayDebtorActivationCancellationRequestV01 struct {
	Hdr         ActivationHeader2               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 Hdr"`
	CxlData     []DebtorActivationCancellation2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 CxlData"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.072.001.01 SplmtryData,omitempty"`
}

type ActivationStatus2 struct {
	OrgnlBizInstr  OriginalBusinessInstruction1  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 OrgnlBizInstr,omitempty"`
	Sts            ServiceStatus1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 Sts"`
	StsRsn         DebtorActivationStatusReason2 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 StsRsn,omitempty"`
	OrgnlActvtnRef OriginalActivation2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 OrgnlActvtnRef,omitempty"`
	FctvActvtnDt   DateAndDateTime2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 FctvActvtnDt,omitempty"`
	SplmtryData    []SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 SplmtryData,omitempty"`
}

type DebtorActivationStatusReason1Choice struct {
	Cd    ExternalDebtorActivationStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 Cd"`
	Prtry common.Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 Prtry"`
}

type DebtorActivationStatusReason2 struct {
	Orgtr    RTPPartyIdentification1             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 Orgtr,omitempty"`
	Rsn      DebtorActivationStatusReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 Rsn"`
	AddtlInf []common.Max105Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 AddtlInf,omitempty"`
}

type RequestToPayDebtorActivationStatusReportV01 struct {
	Hdr               ActivationHeader2    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 Hdr"`
	OrgnlActvtnAndSts []ActivationStatus2  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 OrgnlActvtnAndSts"`
	SplmtryData       []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.073.001.01 SplmtryData,omitempty"`
}
