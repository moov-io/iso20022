package pain_v10

type DocumentPain00700110 struct {
	CstmrPmtRvsl CustomerPaymentReversalV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.10 CstmrPmtRvsl"`
}

type DocumentPain00100110 struct {
	CstmrCdtTrfInitn CustomerCreditTransferInitiationV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.10 CstmrCdtTrfInitn"`
}
