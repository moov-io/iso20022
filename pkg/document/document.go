// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package document

import (
	"encoding/json"
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/acmt_v01"
	"github.com/moov-io/iso20022/pkg/acmt_v02"
	"github.com/moov-io/iso20022/pkg/acmt_v03"
	"github.com/moov-io/iso20022/pkg/admi_v01"
	"github.com/moov-io/iso20022/pkg/admi_v02"
	"github.com/moov-io/iso20022/pkg/auth_v01"
	"github.com/moov-io/iso20022/pkg/auth_v02"
	"github.com/moov-io/iso20022/pkg/camt_v01"
	"github.com/moov-io/iso20022/pkg/camt_v03"
	"github.com/moov-io/iso20022/pkg/camt_v04"
	"github.com/moov-io/iso20022/pkg/camt_v05"
	"github.com/moov-io/iso20022/pkg/camt_v06"
	"github.com/moov-io/iso20022/pkg/camt_v07"
	"github.com/moov-io/iso20022/pkg/camt_v08"
	"github.com/moov-io/iso20022/pkg/camt_v09"
	"github.com/moov-io/iso20022/pkg/camt_v10"
	"github.com/moov-io/iso20022/pkg/head_v01"
	"github.com/moov-io/iso20022/pkg/head_v02"
	"github.com/moov-io/iso20022/pkg/pacs_v04"
	"github.com/moov-io/iso20022/pkg/pacs_v06"
	"github.com/moov-io/iso20022/pkg/pacs_v07"
	"github.com/moov-io/iso20022/pkg/pacs_v08"
	"github.com/moov-io/iso20022/pkg/pacs_v09"
	"github.com/moov-io/iso20022/pkg/pacs_v10"
	"github.com/moov-io/iso20022/pkg/pacs_v11"
	"github.com/moov-io/iso20022/pkg/pain_v01"
	"github.com/moov-io/iso20022/pkg/pain_v05"
	"github.com/moov-io/iso20022/pkg/pain_v08"
	"github.com/moov-io/iso20022/pkg/pain_v09"
	"github.com/moov-io/iso20022/pkg/pain_v10"
	"github.com/moov-io/iso20022/pkg/pain_v11"
	"github.com/moov-io/iso20022/pkg/reda_v01"
	"github.com/moov-io/iso20022/pkg/remt_v02"
	"github.com/moov-io/iso20022/pkg/remt_v04"
	"github.com/moov-io/iso20022/pkg/utils"
)

// Document interface for ISO 20022
type Iso20022Document interface {
	// Validate will be process validation check of document
	Validate() error

	// NameSpace check will return xmlns of document
	NameSpace() string

	// GetXmlName returns xml name of document
	GetXmlName() *xml.Name

	// GetAttrs returns attributes of document
	GetAttrs() []xml.Attr

	// InspectMessage returns message
	InspectMessage() Iso20022Message
}

// Element interface for ISO 20022
type Iso20022Message interface {
	// Validate will be process validation check of document
	Validate() error
}

var (
	messageConstructor = map[string]Iso20022Message{
		utils.DocumentAcmt03600101NameSpace: &acmt_v01.AccountSwitchTerminationSwitchV01{},
		utils.DocumentAcmt02200102NameSpace: &acmt_v02.IdentificationModificationAdviceV02{},
		utils.DocumentAcmt02300102NameSpace: &acmt_v02.IdentificationVerificationRequestV02{},
		utils.DocumentAcmt02400102NameSpace: &acmt_v02.IdentificationVerificationReportV02{},
		utils.DocumentAcmt03000102NameSpace: &acmt_v02.AccountSwitchRequestRedirectionV02{},
		utils.DocumentAcmt03300102NameSpace: &acmt_v02.AccountSwitchNotifyAccountSwitchCompleteV02{},
		utils.DocumentAcmt03500102NameSpace: &acmt_v02.AccountSwitchPaymentResponseV02{},
		utils.DocumentAcmt03700102NameSpace: &acmt_v02.AccountSwitchTechnicalRejectionV02{},
		utils.DocumentAcmt00700103NameSpace: &acmt_v03.AccountOpeningRequestV03{},
		utils.DocumentAcmt00800103NameSpace: &acmt_v03.AccountOpeningAmendmentRequestV03{},
		utils.DocumentAcmt00900103NameSpace: &acmt_v03.AccountOpeningAdditionalInformationRequestV03{},
		utils.DocumentAcmt01000103NameSpace: &acmt_v03.AccountRequestAcknowledgementV03{},
		utils.DocumentAcmt01100103NameSpace: &acmt_v03.AccountRequestRejectionV03{},
		utils.DocumentAcmt01200103NameSpace: &acmt_v03.AccountAdditionalInformationRequestV03{},
		utils.DocumentAcmt01300103NameSpace: &acmt_v03.AccountReportRequestV03{},
		utils.DocumentAcmt01400103NameSpace: &acmt_v03.AccountReportV03{},
		utils.DocumentAcmt01500103NameSpace: &acmt_v03.AccountExcludedMandateMaintenanceRequestV03{},
		utils.DocumentAcmt01600103NameSpace: &acmt_v03.AccountExcludedMandateMaintenanceAmendmentRequestV03{},
		utils.DocumentAcmt01700103NameSpace: &acmt_v03.AccountMandateMaintenanceRequestV03{},
		utils.DocumentAcmt01800103NameSpace: &acmt_v03.AccountMandateMaintenanceAmendmentRequestV03{},
		utils.DocumentAcmt01900103NameSpace: &acmt_v03.AccountClosingRequestV03{},
		utils.DocumentAcmt02000103NameSpace: &acmt_v03.AccountClosingAmendmentRequestV03{},
		utils.DocumentAcmt02100103NameSpace: &acmt_v03.AccountClosingAdditionalInformationRequestV03{},
		utils.DocumentAcmt02700103NameSpace: &acmt_v03.AccountSwitchInformationRequestV03{},
		utils.DocumentAcmt02800103NameSpace: &acmt_v03.AccountSwitchInformationResponseV03{},
		utils.DocumentAcmt02900103NameSpace: &acmt_v03.AccountSwitchCancelExistingPaymentV03{},
		utils.DocumentAcmt03100103NameSpace: &acmt_v03.AccountSwitchRequestBalanceTransferV03{},
		utils.DocumentAcmt03200103NameSpace: &acmt_v03.AccountSwitchBalanceTransferAcknowledgementV03{},
		utils.DocumentAcmt03400103NameSpace: &acmt_v03.AccountSwitchRequestPaymentV03{},
		utils.DocumentAdmi00200101NameSpace: &admi_v01.Admi00200101{},
		utils.DocumentAdmi00400101NameSpace: &admi_v01.Admi00400101{},
		utils.DocumentAdmi00500101NameSpace: &admi_v01.ReportQueryRequestV01{},
		utils.DocumentAdmi00600101NameSpace: &admi_v01.ResendRequestV01{},
		utils.DocumentAdmi00700101NameSpace: &admi_v01.ReceiptAcknowledgementV01{},
		utils.DocumentAdmi01100101NameSpace: &admi_v01.SystemEventAcknowledgementV01{},
		utils.DocumentAdmi01700101NameSpace: &admi_v01.ProcessingRequestV01{},
		utils.DocumentAdmi00400102NameSpace: &admi_v02.SystemEventNotificationV02{},
		utils.DocumentAdmi00900102NameSpace: &admi_v02.StaticDataRequestV02{},
		utils.DocumentAdmi01000102NameSpace: &admi_v02.StaticDataReportV02{},
		utils.DocumentAuth00100101NameSpace: &auth_v01.InformationRequestOpeningV01{},
		utils.DocumentAuth00200101NameSpace: &auth_v01.InformationRequestResponseV01{},
		utils.DocumentAuth00300101NameSpace: &auth_v01.InformationRequestStatusChangeNotificationV01{},
		utils.DocumentAuth01800102NameSpace: &auth_v02.ContractRegistrationRequestV02{},
		utils.DocumentAuth01900102NameSpace: &auth_v02.ContractRegistrationConfirmationV02{},
		utils.DocumentAuth02000102NameSpace: &auth_v02.ContractRegistrationClosureRequestV02{},
		utils.DocumentAuth02100102NameSpace: &auth_v02.ContractRegistrationAmendmentRequestV02{},
		utils.DocumentAuth02200102NameSpace: &auth_v02.ContractRegistrationStatementV02{},
		utils.DocumentAuth02300102NameSpace: &auth_v02.ContractRegistrationStatementRequestV02{},
		utils.DocumentAuth02400102NameSpace: &auth_v02.PaymentRegulatoryInformationNotificationV02{},
		utils.DocumentAuth02500102NameSpace: &auth_v02.CurrencyControlSupportingDocumentDeliveryV02{},
		utils.DocumentAuth02600102NameSpace: &auth_v02.CurrencyControlRequestOrLetterV02{},
		utils.DocumentAuth02700102NameSpace: &auth_v02.CurrencyControlStatusAdviceV02{},
		utils.DocumentCamt10100101NameSpace: &camt_v01.CreateLimitV01{},
		utils.DocumentCamt10200101NameSpace: &camt_v01.CreateStandingOrderV01{},
		utils.DocumentCamt10300101NameSpace: &camt_v01.CreateReservationV01{},
		utils.DocumentCamt10400101NameSpace: &camt_v01.CreateMemberV01{},
		utils.DocumentCamt03500103NameSpace: &camt_v03.ProprietaryFormatInvestigationV03{},
		utils.DocumentCamt06900103NameSpace: &camt_v03.GetStandingOrderV03{},
		utils.DocumentCamt07100103NameSpace: &camt_v03.DeleteStandingOrderV03{},
		utils.DocumentCamt08600103NameSpace: &camt_v03.BankServicesBillingStatementV03{},
		utils.DocumentCamt01300104NameSpace: &camt_v04.GetMemberV04{},
		utils.DocumentCamt01400104NameSpace: &camt_v04.ReturnMemberV04{},
		utils.DocumentCamt01500104NameSpace: &camt_v04.ModifyMemberV04{},
		utils.DocumentCamt01600104NameSpace: &camt_v04.GetCurrencyExchangeRateV04{},
		utils.DocumentCamt01700104NameSpace: &camt_v04.ReturnCurrencyExchangeRateV04{},
		utils.DocumentCamt02000104NameSpace: &camt_v04.GetGeneralBusinessInformationV04{},
		utils.DocumentCamt03200104NameSpace: &camt_v04.CancelCaseAssignmentV04{},
		utils.DocumentCamt03800104NameSpace: &camt_v04.CaseStatusReportRequestV04{},
		utils.DocumentCamt07000104NameSpace: &camt_v04.ReturnStandingOrderV04{},
		utils.DocumentCamt01800105NameSpace: &camt_v05.GetBusinessDayInformationV05{},
		utils.DocumentCamt02500105NameSpace: &camt_v05.ReceiptV05{},
		utils.DocumentCamt02600105NameSpace: &camt_v05.UnableToApplyV05{},
		utils.DocumentCamt02800105NameSpace: &camt_v05.AdditionalPaymentInformationV05{},
		utils.DocumentCamt03000105NameSpace: &camt_v05.NotificationOfCaseAssignmentV05{},
		utils.DocumentCamt03500105NameSpace: &camt_v05.ProprietaryFormatInvestigationV05{},
		utils.DocumentCamt03600105NameSpace: &camt_v05.DebitAuthorisationResponseV05{},
		utils.DocumentCamt03900105NameSpace: &camt_v05.CaseStatusReportV05{},
		utils.DocumentCamt04600105NameSpace: &camt_v05.GetReservationV05{},
		utils.DocumentCamt04800105NameSpace: &camt_v05.ModifyReservationV05{},
		utils.DocumentCamt04900105NameSpace: &camt_v05.DeleteReservationV05{},
		utils.DocumentCamt05000105NameSpace: &camt_v05.LiquidityCreditTransferV05{},
		utils.DocumentCamt05100105NameSpace: &camt_v05.LiquidityDebitTransferV05{},
		utils.DocumentCamt05600105NameSpace: &camt_v05.FIToFIPaymentCancellationRequestV05{},
		utils.DocumentCamt06000105NameSpace: &camt_v05.AccountReportingRequestV05{},
		utils.DocumentCamt02100106NameSpace: &camt_v06.ReturnGeneralBusinessInformationV06{},
		utils.DocumentCamt02400106NameSpace: &camt_v06.ModifyStandingOrderV06{},
		utils.DocumentCamt02900106NameSpace: &camt_v06.ResolutionOfInvestigationV06{},
		utils.DocumentCamt03100106NameSpace: &camt_v06.RejectInvestigationV06{},
		utils.DocumentCamt03300106NameSpace: &camt_v06.RequestForDuplicateV06{},
		utils.DocumentCamt03400106NameSpace: &camt_v06.DuplicateV06{},
		utils.DocumentCamt04700106NameSpace: &camt_v06.ReturnReservationV06{},
		utils.DocumentCamt05700106NameSpace: &camt_v06.NotificationToReceiveV06{},
		utils.DocumentCamt05800106NameSpace: &camt_v06.NotificationToReceiveCancellationAdviceV06{},
		utils.DocumentCamt05900106NameSpace: &camt_v06.NotificationToReceiveStatusReportV06{},
		utils.DocumentCamt00300107NameSpace: &camt_v07.GetAccountV07{},
		utils.DocumentCamt00900107NameSpace: &camt_v07.GetLimitV07{},
		utils.DocumentCamt01100107NameSpace: &camt_v07.ModifyLimitV07{},
		utils.DocumentCamt01200107NameSpace: &camt_v07.DeleteLimitV07{},
		utils.DocumentCamt01900107NameSpace: &camt_v07.ReturnBusinessDayInformationV07{},
		utils.DocumentCamt02300107NameSpace: &camt_v07.BackupPaymentV07{},
		utils.DocumentCamt08700107NameSpace: &camt_v07.RequestToModifyPaymentV07{},
		utils.DocumentCamt00400108NameSpace: &camt_v08.ReturnAccountV08{},
		utils.DocumentCamt00500108NameSpace: &camt_v08.GetTransactionV08{},
		utils.DocumentCamt00600108NameSpace: &camt_v08.ReturnTransactionV08{},
		utils.DocumentCamt00700108NameSpace: &camt_v08.ModifyTransactionV08{},
		utils.DocumentCamt00800108NameSpace: &camt_v08.CancelTransactionV08{},
		utils.DocumentCamt01000108NameSpace: &camt_v08.ReturnLimitV08{},
		utils.DocumentCamt02600108NameSpace: &camt_v08.UnableToApplyV08{},
		utils.DocumentCamt02700108NameSpace: &camt_v08.ClaimNonReceiptV08{},
		utils.DocumentCamt03700108NameSpace: &camt_v08.DebitAuthorisationRequestV08{},
		utils.DocumentCamt05200108NameSpace: &camt_v08.BankToCustomerAccountReportV08{},
		utils.DocumentCamt05300108NameSpace: &camt_v08.BankToCustomerStatementV08{},
		utils.DocumentCamt05400108NameSpace: &camt_v08.BankToCustomerDebitCreditNotificationV08{},
		utils.DocumentCamt05500109NameSpace: &camt_v09.CustomerPaymentCancellationRequestV09{},
		utils.DocumentCamt05600109NameSpace: &camt_v09.FIToFIPaymentCancellationRequestV09{},
		utils.DocumentCamt02800110NameSpace: &camt_v10.AdditionalPaymentInformationV10{},
		utils.DocumentCamt02900110NameSpace: &camt_v10.ResolutionOfInvestigationV10{},
		utils.DocumentHead00100101NameSpace: &head_v01.BusinessApplicationHeaderV01{},
		utils.DocumentHead00100102NameSpace: &head_v02.BusinessApplicationHeaderV02{},
		utils.DocumentPacs01000104NameSpace: &pacs_v04.FinancialInstitutionDirectDebitV04{},
		utils.DocumentPacs02800104NameSpace: &pacs_v04.FIToFIPaymentStatusRequestV04{},
		utils.DocumentPacs00800106NameSpace: &pacs_v06.FIToFICustomerCreditTransferV06{},
		utils.DocumentPacs00200107NameSpace: &pacs_v07.FIToFIPaymentStatusReportV07{},
		utils.DocumentPacs00300108NameSpace: &pacs_v08.FIToFICustomerDirectDebitV08{},
		utils.DocumentPacs00800109NameSpace: &pacs_v09.FIToFICustomerCreditTransferV09{},
		utils.DocumentPacs00900109NameSpace: &pacs_v09.FinancialInstitutionCreditTransferV09{},
		utils.DocumentPacs00200110NameSpace: &pacs_v10.FIToFIPaymentStatusReportV10{},
		utils.DocumentPacs00400110NameSpace: &pacs_v10.PaymentReturnV10{},
		utils.DocumentPacs00700110NameSpace: &pacs_v10.FIToFIPaymentReversalV10{},
		utils.DocumentPacs00200111NameSpace: &pacs_v11.FIToFIPaymentStatusReportV11{},
		utils.DocumentPain00700101NameSpace: &pain_v01.MandateCopyRequestV01{},
		utils.DocumentPain01800101NameSpace: &pain_v01.MandateSuspensionRequestV01{},
		utils.DocumentPain00900105NameSpace: &pain_v05.MandateInitiationRequestV05{},
		utils.DocumentPain01000105NameSpace: &pain_v05.MandateAmendmentRequestV05{},
		utils.DocumentPain01100105NameSpace: &pain_v05.MandateCancellationRequestV05{},
		utils.DocumentPain01200105NameSpace: &pain_v05.MandateAcceptanceReportV05{},
		utils.DocumentPain01300105NameSpace: &pain_v05.CreditorPaymentActivationRequestV05{},
		utils.DocumentPain01400105NameSpace: &pain_v05.CreditorPaymentActivationRequestStatusReportV05{},
		utils.DocumentPain01400108NameSpace: &pain_v08.CreditorPaymentActivationRequestStatusReportV08{},
		utils.DocumentPain01300108NameSpace: &pain_v08.CreditorPaymentActivationRequestV08{},
		utils.DocumentPain00800109NameSpace: &pain_v09.CustomerDirectDebitInitiationV09{},
		utils.DocumentPain00100110NameSpace: &pain_v10.CustomerCreditTransferInitiationV10{},
		utils.DocumentPain00700110NameSpace: &pain_v10.CustomerPaymentReversalV10{},
		utils.DocumentPain00200111NameSpace: &pain_v11.CustomerPaymentStatusReportV11{},
		utils.DocumentReda06600101NameSpace: &reda_v01.RequestToPayCreditorEnrolmentRequestV01{},
		utils.DocumentReda06700101NameSpace: &reda_v01.RequestToPayCreditorEnrolmentAmendmentRequestV01{},
		utils.DocumentReda06800101NameSpace: &reda_v01.RequestToPayCreditorEnrolmentCancellationRequestV01{},
		utils.DocumentReda06900101NameSpace: &reda_v01.RequestToPayCreditorEnrolmentStatusReportV01{},
		utils.DocumentReda07000101NameSpace: &reda_v01.RequestToPayDebtorActivationRequestV01{},
		utils.DocumentReda07100101NameSpace: &reda_v01.RequestToPayDebtorActivationAmendmentRequestV01{},
		utils.DocumentReda07200101NameSpace: &reda_v01.RequestToPayDebtorActivationCancellationRequestV01{},
		utils.DocumentReda07300101NameSpace: &reda_v01.RequestToPayDebtorActivationStatusReportV01{},
		utils.DocumentRemt00100102NameSpace: &remt_v02.RemittanceAdviceV02{},
		utils.DocumentRemt00200102NameSpace: &remt_v02.RemittanceLocationAdviceV02{},
		utils.DocumentRemt00100104NameSpace: &remt_v04.RemittanceAdviceV04{},
	}
)

type documentDummy struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

func (dummy documentDummy) NameSpace() string {
	for _, attr := range dummy.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func NewDocument(space string) (doc Iso20022Document, err error) {
	msg := messageConstructor[space]
	if msg == nil {
		return nil, utils.NewErrUnsupportedNameSpace()
	}

	return &Iso20022DocumentObject{
		Message: msg,
	}, nil
}

// ParseIso20022Document will return a interface of ISO 20022 document after pass buffer
func ParseIso20022Document(buf []byte) (Iso20022Document, error) {

	bType := utils.GetBufferFormat(buf)
	if bType == utils.DocumentTypeUnknown {
		return nil, utils.NewErrInvalidFileType()
	}

	var dummy documentDummy
	var err error

	if bType == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, &dummy)
	} else {
		err = json.Unmarshal(buf, &dummy)
	}
	if err != nil {
		return nil, err
	}

	namespace := dummy.NameSpace()
	if namespace == "" {
		return nil, utils.NewErrOmittedNameSpace()
	}

	msg := messageConstructor[namespace]
	if msg == nil {
		return nil, utils.NewErrUnsupportedNameSpace()
	}

	doc := &Iso20022DocumentObject{
		Message: msg,
	}

	if bType == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, doc)
	} else {
		err = json.Unmarshal(buf, doc)
	}
	if err != nil {
		return nil, err
	}

	return doc, nil
}

type Iso20022DocumentObject struct {
	XMLName xml.Name
	Attrs   []xml.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
	Message Iso20022Message `xml:",any"`
}

func (doc Iso20022DocumentObject) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc Iso20022DocumentObject) NameSpace() string {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (doc *Iso20022DocumentObject) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *Iso20022DocumentObject) GetAttrs() []xml.Attr {
	return doc.Attrs
}

func (doc *Iso20022DocumentObject) InspectMessage() Iso20022Message {
	return doc.Message
}

func (doc Iso20022DocumentObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName xml.Name
		Attrs   []xml.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
		Message Iso20022Message `xml:",any"`
	}(doc)

	updatingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func updatingStartElement(start *xml.StartElement, attrs []xml.Attr, name xml.Name) {
	for _, attr := range attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			name.Space = ""
		}
	}
	if len(name.Local) > 0 {
		start.Name.Local = name.Local
	}
	start.Name.Space = name.Space
}
