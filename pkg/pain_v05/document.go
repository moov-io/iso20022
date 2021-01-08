package pain_v05

type DocumentPain00900105 struct {
	MndtInitnReq MandateInitiationRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.05 MndtInitnReq"`
}

type DocumentPain01000105 struct {
	MndtAmdmntReq MandateAmendmentRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.010.001.05 MndtAmdmntReq"`
}

type DocumentPain01200105 struct {
	MndtAccptncRpt MandateAcceptanceReportV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.05 MndtAccptncRpt"`
}

type DocumentPain01100105 struct {
	MndtCxlReq MandateCancellationRequestV05 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.05 MndtCxlReq"`
}
