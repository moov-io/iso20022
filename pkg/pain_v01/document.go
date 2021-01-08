package pain_v01

type DocumentPain00700101 struct {
	MndtCpyReq MandateCopyRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 MndtCpyReq"`
}

type DocumentPain01800101 struct {
	MndtSspnsnReq MandateSuspensionRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:pain.018.001.01 MndtSspnsnReq"`
}
