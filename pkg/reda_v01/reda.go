// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import (
	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
)

type AddressType3Choice struct {
	Cd    common.AddressType2Code `xml:"Cd"`
	Prtry GenericIdentification30 `xml:"Prtry"`
}

func (r AddressType3Choice) Validate() error {
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

type CreditorEnrolment3 struct {
	Enrlmnt      CreditorServiceEnrolment1             `xml:"Enrlmnt"`
	CdtrTradgNm  *common.Max140Text                    `xml:"CdtrTradgNm,omitempty" json:",omitempty"`
	Cdtr         RTPPartyIdentification1               `xml:"Cdtr"`
	UltmtCdtr    *RTPPartyIdentification1              `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	MrchntCtgyCd common.MerchantCategoryCodeIdentifier `xml:"MrchntCtgyCd"`
	CdtrLogo     *common.Max10KBinary                  `xml:"CdtrLogo,omitempty" json:",omitempty"`
}

func (r CreditorEnrolment3) Validate() error {
	return utils.Validate(&r)
}

type CreditorInvoice3 struct {
	LtdPresntmntInd   bool                    `xml:"LtdPresntmntInd"`
	CstmrIdTp         *CustomerTypeRequest2   `xml:"CstmrIdTp,omitempty" json:",omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice `xml:"CtrctFrmtTp,omitempty" json:",omitempty"`
	CtrctRefTp        []DocumentType1Choice   `xml:"CtrctRefTp,omitempty" json:",omitempty"`
	CdtrInstr         *common.Max500Text      `xml:"CdtrInstr,omitempty" json:",omitempty"`
	ActvtnReqDlvryPty RTPPartyIdentification1 `xml:"ActvtnReqDlvryPty"`
}

func (r CreditorInvoice3) Validate() error {
	return utils.Validate(&r)
}

type CreditorServiceEnrolment1 struct {
	EnrlmntStartDt  *DateAndDateTime2Choice `xml:"EnrlmntStartDt,omitempty" json:",omitempty"`
	EnrlmntEndDt    *DateAndDateTime2Choice `xml:"EnrlmntEndDt,omitempty" json:",omitempty"`
	Vsblty          *Visibilty1             `xml:"Vsblty,omitempty" json:",omitempty"`
	SvcActvtnAllwd  bool                    `xml:"SvcActvtnAllwd"`
	SvcDescLk       *common.Max2048Text     `xml:"SvcDescLk,omitempty" json:",omitempty"`
	CdtrSvcActvtnLk *common.Max2048Text     `xml:"CdtrSvcActvtnLk,omitempty" json:",omitempty"`
}

func (r CreditorServiceEnrolment1) Validate() error {
	return utils.Validate(&r)
}

type CustomerTypeRequest2 struct {
	Reqd   bool               `xml:"Reqd"`
	OrgTp  *OrganisationType2 `xml:"OrgTp,omitempty" json:",omitempty"`
	PrvtTp *PersonType2       `xml:"PrvtTp,omitempty" json:",omitempty"`
}

func (r CustomerTypeRequest2) Validate() error {
	return utils.Validate(&r)
}

type DateAndDateTime2Choice struct {
	Dt   common.ISODate     `xml:"Dt"`
	DtTm common.ISODateTime `xml:"DtTm"`
}

func (r DateAndDateTime2Choice) Validate() error {
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

type DocumentFormat2Choice struct {
	Cd    ExternalDocumentFormat1Code `xml:"Cd"`
	Prtry GenericIdentification1      `xml:"Prtry"`
}

func (r DocumentFormat2Choice) Validate() error {
	return utils.Validate(&r)
}

type DocumentType1Choice struct {
	Cd    ExternalDocumentType1Code `xml:"Cd"`
	Prtry GenericIdentification1    `xml:"Prtry"`
}

func (r DocumentType1Choice) Validate() error {
	return utils.Validate(&r)
}

type EnrolmentHeader2 struct {
	MsgId    common.Max35Text         `xml:"MsgId"`
	CreDtTm  common.ISODateTime       `xml:"CreDtTm"`
	MsgOrgtr *RTPPartyIdentification1 `xml:"MsgOrgtr,omitempty" json:",omitempty"`
	MsgRcpt  *RTPPartyIdentification1 `xml:"MsgRcpt,omitempty" json:",omitempty"`
	InitgPty RTPPartyIdentification1  `xml:"InitgPty"`
}

func (r EnrolmentHeader2) Validate() error {
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

type GenericIdentification30 struct {
	Id      common.Exact4AlphaNumericText `xml:"Id"`
	Issr    common.Max35Text              `xml:"Issr"`
	SchmeNm *common.Max35Text             `xml:"SchmeNm,omitempty" json:",omitempty"`
}

func (r GenericIdentification30) Validate() error {
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

type GenericOrganisationType1 struct {
	Reqd    bool                                        `xml:"Reqd"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm"`
}

func (r GenericOrganisationType1) Validate() error {
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

type GenericPersonType1 struct {
	Reqd    bool                                  `xml:"Reqd"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"SchmeNm"`
}

func (r GenericPersonType1) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentification37 struct {
	AnyBIC   *common.AnyBICDec2014Identifier      `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI      *common.LEIIdentifier                `xml:"LEI,omitempty" json:",omitempty"`
	EmailAdr *common.Max256Text                   `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr     []GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationIdentification37) Validate() error {
	return utils.Validate(&r)
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                        `xml:"Prtry"`
}

func (r OrganisationIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type OrganisationType2 struct {
	AnyBIC   bool                       `xml:"AnyBIC,omitempty" json:",omitempty"`
	LEI      bool                       `xml:"LEI,omitempty" json:",omitempty"`
	EmailAdr bool                       `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr     []GenericOrganisationType1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r OrganisationType2) Validate() error {
	return utils.Validate(&r)
}

type OtherContact1 struct {
	ChanlTp common.Max4Text    `xml:"ChanlTp"`
	Id      *common.Max128Text `xml:"Id,omitempty" json:",omitempty"`
}

func (r OtherContact1) Validate() error {
	return utils.Validate(&r)
}

type Party49Choice struct {
	OrgId  OrganisationIdentification37 `xml:"OrgId"`
	PrvtId PersonIdentification17       `xml:"PrvtId"`
}

func (r Party49Choice) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentification17 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1          `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	EmailAdr        *common.Max256Text             `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonIdentification17) Validate() error {
	return utils.Validate(&r)
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"Cd"`
	Prtry common.Max35Text                  `xml:"Prtry"`
}

func (r PersonIdentificationSchemeName1Choice) Validate() error {
	return utils.Validate(&r)
}

type PersonType2 struct {
	DtAndPlcOfBirth bool                 `xml:"DtAndPlcOfBirth,omitempty" json:",omitempty"`
	EmailAdr        bool                 `xml:"EmailAdr,omitempty" json:",omitempty"`
	Othr            []GenericPersonType1 `xml:"Othr,omitempty" json:",omitempty"`
}

func (r PersonType2) Validate() error {
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

type RTPPartyIdentification1 struct {
	Nm        *common.Max140Text  `xml:"Nm,omitempty" json:",omitempty"`
	PstlAdr   *PostalAddress24    `xml:"PstlAdr,omitempty" json:",omitempty"`
	Id        *Party49Choice      `xml:"Id,omitempty" json:",omitempty"`
	CtryOfRes *common.CountryCode `xml:"CtryOfRes,omitempty" json:",omitempty"`
	CtctDtls  *Contact4           `xml:"CtctDtls,omitempty" json:",omitempty"`
}

func (r RTPPartyIdentification1) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayCreditorEnrolmentRequestV01 struct {
	Attr        []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr         EnrolmentHeader2     `xml:"Hdr"`
	CdtrEnrlmnt []CreditorEnrolment3 `xml:"CdtrEnrlmnt" json:",omitempty"`
	ActvtnData  CreditorInvoice3     `xml:"ActvtnData"`
	SplmtryData []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayCreditorEnrolmentRequestV01) Validate() error {
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

type Visibilty1 struct {
	StartDt   *DateAndDateTime2Choice `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt     *DateAndDateTime2Choice `xml:"EndDt,omitempty" json:",omitempty"`
	LtdVsblty bool                    `xml:"LtdVsblty,omitempty" json:",omitempty"`
}

func (r Visibilty1) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolment4 struct {
	Enrlmnt      *CreditorServiceEnrolment1             `xml:"Enrlmnt,omitempty" json:",omitempty"`
	CdtrTradgNm  *common.Max140Text                     `xml:"CdtrTradgNm,omitempty" json:",omitempty"`
	Cdtr         *RTPPartyIdentification1               `xml:"Cdtr"`
	UltmtCdtr    *RTPPartyIdentification1               `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	MrchntCtgyCd *common.MerchantCategoryCodeIdentifier `xml:"MrchntCtgyCd,omitempty" json:",omitempty"`
	CdtrLogo     *common.Max10KBinary                   `xml:"CdtrLogo,omitempty" json:",omitempty"`
}

func (r CreditorEnrolment4) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentAmendment3 struct {
	OrgnlBizInstr *OriginalBusinessInstruction1      `xml:"OrgnlBizInstr,omitempty" json:",omitempty"`
	AmdmntRsn     *CreditorEnrolmentAmendmentReason2 `xml:"AmdmntRsn,omitempty" json:",omitempty"`
	Amdmnt        CreditorEnrolmentAmendment4        `xml:"Amdmnt"`
	OrgnlEnrlmnt  OriginalEnrolment2Choice           `xml:"OrgnlEnrlmnt"`
	SplmtryData   []SupplementaryData1               `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorEnrolmentAmendment3) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentAmendment4 struct {
	CdtrEnrlmnt *CreditorEnrolment4 `xml:"CdtrEnrlmnt,omitempty" json:",omitempty"`
	ActvtnData  *CreditorInvoice4   `xml:"ActvtnData,omitempty" json:",omitempty"`
}

func (r CreditorEnrolmentAmendment4) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentAmendmentReason1Choice struct {
	Cd    ExternalCreditorEnrolmentAmendmentReason1Code `xml:"Cd"`
	Prtry common.Max35Text                              `xml:"Prtry"`
}

func (r CreditorEnrolmentAmendmentReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentAmendmentReason2 struct {
	Orgtr    *RTPPartyIdentification1                `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      CreditorEnrolmentAmendmentReason1Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text                     `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r CreditorEnrolmentAmendmentReason2) Validate() error {
	return utils.Validate(&r)
}

type CreditorInvoice4 struct {
	LtdPresntmntInd   bool                     `xml:"LtdPresntmntInd,omitempty" json:",omitempty"`
	CstmrIdTp         *CustomerTypeRequest2    `xml:"CstmrIdTp,omitempty" json:",omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice  `xml:"CtrctFrmtTp,omitempty" json:",omitempty"`
	CtrctRefTp        []DocumentType1Choice    `xml:"CtrctRefTp,omitempty" json:",omitempty"`
	CdtrInstr         *common.Max500Text       `xml:"CdtrInstr,omitempty" json:",omitempty"`
	ActvtnReqDlvryPty *RTPPartyIdentification1 `xml:"ActvtnReqDlvryPty,omitempty" json:",omitempty"`
}

func (r CreditorInvoice4) Validate() error {
	return utils.Validate(&r)
}

type OriginalBusinessInstruction1 struct {
	MsgId   common.Max35Text    `xml:"MsgId"`
	MsgNmId *common.Max35Text   `xml:"MsgNmId,omitempty" json:",omitempty"`
	CreDtTm *common.ISODateTime `xml:"CreDtTm,omitempty" json:",omitempty"`
}

func (r OriginalBusinessInstruction1) Validate() error {
	return utils.Validate(&r)
}

type OriginalEnrolment2Choice struct {
	OrgnlCdtrId      Party49Choice      `xml:"OrgnlCdtrId"`
	OrgnlEnrlmntData CreditorEnrolment3 `xml:"OrgnlEnrlmntData"`
}

func (r OriginalEnrolment2Choice) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayCreditorEnrolmentAmendmentRequestV01 struct {
	Attr        []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr         EnrolmentHeader2              `xml:"Hdr"`
	AmdmntData  []CreditorEnrolmentAmendment3 `xml:"AmdmntData" json:",omitempty"`
	SplmtryData []SupplementaryData1          `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayCreditorEnrolmentAmendmentRequestV01) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentCancellation2 struct {
	OrgnlBizInstr *OriginalBusinessInstruction1         `xml:"OrgnlBizInstr,omitempty" json:",omitempty"`
	CxlRsn        *CreditorEnrolmentCancellationReason2 `xml:"CxlRsn,omitempty" json:",omitempty"`
	OrgnlEnrlmnt  OriginalEnrolment2Choice              `xml:"OrgnlEnrlmnt"`
	SplmtryData   []SupplementaryData1                  `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r CreditorEnrolmentCancellation2) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentCancellationReason1Choice struct {
	Cd    ExternalCreditorEnrolmentCancellationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                                 `xml:"Prtry"`
}

func (r CreditorEnrolmentCancellationReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentCancellationReason2 struct {
	Orgtr    *RTPPartyIdentification1                   `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      CreditorEnrolmentCancellationReason1Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text                        `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r CreditorEnrolmentCancellationReason2) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayCreditorEnrolmentCancellationRequestV01 struct {
	Attr        []utils.Attr                     `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr         EnrolmentHeader2                 `xml:"Hdr"`
	CxlData     []CreditorEnrolmentCancellation2 `xml:"CxlData" json:",omitempty"`
	SplmtryData []SupplementaryData1             `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayCreditorEnrolmentCancellationRequestV01) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentStatusReason2 struct {
	Orgtr    *RTPPartyIdentification1             `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      CreditorEnrolmentStatusReason2Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text                  `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r CreditorEnrolmentStatusReason2) Validate() error {
	return utils.Validate(&r)
}

type CreditorEnrolmentStatusReason2Choice struct {
	Cd    ExternalCreditorEnrolmentStatusReason1Code `xml:"Cd"`
	Prtry common.Max35Text                           `xml:"Prtry"`
}

func (r CreditorEnrolmentStatusReason2Choice) Validate() error {
	return utils.Validate(&r)
}

type EnrolmentStatus2 struct {
	OrgnlBizInstr   *OriginalBusinessInstruction1   `xml:"OrgnlBizInstr,omitempty" json:",omitempty"`
	Sts             ServiceStatus1Choice            `xml:"Sts"`
	StsRsn          *CreditorEnrolmentStatusReason2 `xml:"StsRsn,omitempty" json:",omitempty"`
	OrgnlEnrlmntRef *OriginalEnrolment2Choice       `xml:"OrgnlEnrlmntRef,omitempty" json:",omitempty"`
	FctvEnrlmntDt   *DateAndDateTime2Choice         `xml:"FctvEnrlmntDt,omitempty" json:",omitempty"`
	SplmtryData     []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r EnrolmentStatus2) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayCreditorEnrolmentStatusReportV01 struct {
	Attr               []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr                EnrolmentHeader2     `xml:"Hdr"`
	OrgnlEnrlmntAndSts []EnrolmentStatus2   `xml:"OrgnlEnrlmntAndSts" json:",omitempty"`
	SplmtryData        []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayCreditorEnrolmentStatusReportV01) Validate() error {
	return utils.Validate(&r)
}

type ServiceStatus1Choice struct {
	Cd    ServiceRequestStatus1Code `xml:"Cd"`
	Prtry common.Max35Text          `xml:"Prtry"`
}

func (r ServiceStatus1Choice) Validate() error {
	return utils.Validate(&r)
}

type ActivationHeader2 struct {
	MsgId    common.Max35Text         `xml:"MsgId"`
	CreDtTm  common.ISODateTime       `xml:"CreDtTm"`
	MsgOrgtr *RTPPartyIdentification1 `xml:"MsgOrgtr,omitempty" json:",omitempty"`
	MsgRcpt  *RTPPartyIdentification1 `xml:"MsgRcpt,omitempty" json:",omitempty"`
	InitgPty RTPPartyIdentification1  `xml:"InitgPty"`
}

func (r ActivationHeader2) Validate() error {
	return utils.Validate(&r)
}

type ContractReference1 struct {
	Tp  *DocumentType1Choice `xml:"Tp,omitempty" json:",omitempty"`
	Ref common.Max500Text    `xml:"Ref"`
}

func (r ContractReference1) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivation3 struct {
	DbtrActvtnId      *common.Max35Text        `xml:"DbtrActvtnId,omitempty" json:",omitempty"`
	DispNm            *common.Max140Text       `xml:"DispNm,omitempty" json:",omitempty"`
	UltmtDbtr         *RTPPartyIdentification1 `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr              RTPPartyIdentification1  `xml:"Dbtr"`
	DbtrSolPrvdr      RTPPartyIdentification1  `xml:"DbtrSolPrvdr"`
	CstmrId           []Party49Choice          `xml:"CstmrId,omitempty" json:",omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice  `xml:"CtrctFrmtTp,omitempty" json:",omitempty"`
	CtrctRef          []ContractReference1     `xml:"CtrctRef,omitempty" json:",omitempty"`
	Cdtr              RTPPartyIdentification1  `xml:"Cdtr"`
	UltmtCdtr         *RTPPartyIdentification1 `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	ActvtnReqDlvryPty *RTPPartyIdentification1 `xml:"ActvtnReqDlvryPty,omitempty" json:",omitempty"`
	StartDt           *DateAndDateTime2Choice  `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt             *DateAndDateTime2Choice  `xml:"EndDt,omitempty" json:",omitempty"`
	DdctdActvtnCd     *common.Max35Text        `xml:"DdctdActvtnCd,omitempty" json:",omitempty"`
}

func (r DebtorActivation3) Validate() error {
	return utils.Validate(&r)
}

type ElectronicInvoice1 struct {
	PresntmntTp PresentmentType1Code `xml:"PresntmntTp"`
}

func (r ElectronicInvoice1) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayDebtorActivationRequestV01 struct {
	Attr            []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr             ActivationHeader2    `xml:"Hdr"`
	DbtrActvtn      []DebtorActivation3  `xml:"DbtrActvtn" json:",omitempty"`
	ElctrncInvcData ElectronicInvoice1   `xml:"ElctrncInvcData"`
	SplmtryData     []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayDebtorActivationRequestV01) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivation4 struct {
	DbtrActvtnId      *common.Max35Text        `xml:"DbtrActvtnId,omitempty" json:",omitempty"`
	DispNm            *common.Max140Text       `xml:"DispNm,omitempty" json:",omitempty"`
	UltmtDbtr         *RTPPartyIdentification1 `xml:"UltmtDbtr,omitempty" json:",omitempty"`
	Dbtr              *RTPPartyIdentification1 `xml:"Dbtr,omitempty" json:",omitempty"`
	DbtrSolPrvdr      *RTPPartyIdentification1 `xml:"DbtrSolPrvdr,omitempty" json:",omitempty"`
	CstmrId           []Party49Choice          `xml:"CstmrId,omitempty" json:",omitempty"`
	CtrctFrmtTp       []DocumentFormat2Choice  `xml:"CtrctFrmtTp,omitempty" json:",omitempty"`
	CtrctRef          []ContractReference1     `xml:"CtrctRef,omitempty" json:",omitempty"`
	Cdtr              *RTPPartyIdentification1 `xml:"Cdtr,omitempty" json:",omitempty"`
	UltmtCdtr         *RTPPartyIdentification1 `xml:"UltmtCdtr,omitempty" json:",omitempty"`
	ActvtnReqDlvryPty *RTPPartyIdentification1 `xml:"ActvtnReqDlvryPty,omitempty" json:",omitempty"`
	StartDt           *DateAndDateTime2Choice  `xml:"StartDt,omitempty" json:",omitempty"`
	EndDt             *DateAndDateTime2Choice  `xml:"EndDt,omitempty" json:",omitempty"`
	DdctdActvtnCd     *common.Max35Text        `xml:"DdctdActvtnCd,omitempty" json:",omitempty"`
}

func (r DebtorActivation4) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationAmendment3 struct {
	OrgnlBizInstr *OriginalBusinessInstruction1     `xml:"OrgnlBizInstr,omitempty" json:",omitempty"`
	AmdmntRsn     *DebtorActivationAmendmentReason2 `xml:"AmdmntRsn,omitempty" json:",omitempty"`
	Amdmnt        DebtorActivationAmendment4        `xml:"Amdmnt"`
	OrgnlActvtn   OriginalActivation2Choice         `xml:"OrgnlActvtn"`
	SplmtryData   []SupplementaryData1              `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DebtorActivationAmendment3) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationAmendment4 struct {
	DbtrActvtn      *DebtorActivation4  `xml:"DbtrActvtn,omitempty" json:",omitempty"`
	ElctrncInvcData *ElectronicInvoice1 `xml:"ElctrncInvcData,omitempty" json:",omitempty"`
}

func (r DebtorActivationAmendment4) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationAmendmentReason1Choice struct {
	Cd    ExternalDebtorActivationAmendmentReason1Code `xml:"Cd"`
	Prtry common.Max35Text                             `xml:"Prtry"`
}

func (r DebtorActivationAmendmentReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationAmendmentReason2 struct {
	Orgtr    *RTPPartyIdentification1               `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      DebtorActivationAmendmentReason1Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text                    `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DebtorActivationAmendmentReason2) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayDebtorActivationAmendmentRequestV01 struct {
	Attr        []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr         ActivationHeader2            `xml:"Hdr"`
	AmdmntData  []DebtorActivationAmendment3 `xml:"AmdmntData" json:",omitempty"`
	SplmtryData []SupplementaryData1         `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayDebtorActivationAmendmentRequestV01) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationCancellation2 struct {
	OrgnlBizInstr *OriginalBusinessInstruction1        `xml:"OrgnlBizInstr,omitempty" json:",omitempty"`
	CxlRsn        *DebtorActivationCancellationReason2 `xml:"CxlRsn,omitempty" json:",omitempty"`
	OrgnlActvtn   OriginalActivation2Choice            `xml:"OrgnlActvtn"`
	SplmtryData   []SupplementaryData1                 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r DebtorActivationCancellation2) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationCancellationReason1Choice struct {
	Cd    ExternalDebtorActivationCancellationReason1Code `xml:"Cd"`
	Prtry common.Max35Text                                `xml:"Prtry"`
}

func (r DebtorActivationCancellationReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationCancellationReason2 struct {
	Orgtr    *RTPPartyIdentification1                  `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      DebtorActivationCancellationReason1Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text                       `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DebtorActivationCancellationReason2) Validate() error {
	return utils.Validate(&r)
}

type OriginalActivation2Choice struct {
	OrgnlDbtrId     Party49Choice     `xml:"OrgnlDbtrId"`
	OrgnlActvtnData DebtorActivation3 `xml:"OrgnlActvtnData"`
}

func (r OriginalActivation2Choice) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayDebtorActivationCancellationRequestV01 struct {
	Attr        []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr         ActivationHeader2               `xml:"Hdr"`
	CxlData     []DebtorActivationCancellation2 `xml:"CxlData" json:",omitempty"`
	SplmtryData []SupplementaryData1            `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayDebtorActivationCancellationRequestV01) Validate() error {
	return utils.Validate(&r)
}

type ActivationStatus2 struct {
	OrgnlBizInstr  *OriginalBusinessInstruction1  `xml:"OrgnlBizInstr,omitempty" json:",omitempty"`
	Sts            ServiceStatus1Choice           `xml:"Sts"`
	StsRsn         *DebtorActivationStatusReason2 `xml:"StsRsn,omitempty" json:",omitempty"`
	OrgnlActvtnRef *OriginalActivation2Choice     `xml:"OrgnlActvtnRef,omitempty" json:",omitempty"`
	FctvActvtnDt   *DateAndDateTime2Choice        `xml:"FctvActvtnDt,omitempty" json:",omitempty"`
	SplmtryData    []SupplementaryData1           `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r ActivationStatus2) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationStatusReason1Choice struct {
	Cd    ExternalDebtorActivationStatusReason1Code `xml:"Cd"`
	Prtry common.Max35Text                          `xml:"Prtry"`
}

func (r DebtorActivationStatusReason1Choice) Validate() error {
	return utils.Validate(&r)
}

type DebtorActivationStatusReason2 struct {
	Orgtr    *RTPPartyIdentification1            `xml:"Orgtr,omitempty" json:",omitempty"`
	Rsn      DebtorActivationStatusReason1Choice `xml:"Rsn"`
	AddtlInf []common.Max105Text                 `xml:"AddtlInf,omitempty" json:",omitempty"`
}

func (r DebtorActivationStatusReason2) Validate() error {
	return utils.Validate(&r)
}

type RequestToPayDebtorActivationStatusReportV01 struct {
	Attr              []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	Hdr               ActivationHeader2    `xml:"Hdr"`
	OrgnlActvtnAndSts []ActivationStatus2  `xml:"OrgnlActvtnAndSts" json:",omitempty"`
	SplmtryData       []SupplementaryData1 `xml:"SplmtryData,omitempty" json:",omitempty"`
}

func (r RequestToPayDebtorActivationStatusReportV01) Validate() error {
	return utils.Validate(&r)
}
