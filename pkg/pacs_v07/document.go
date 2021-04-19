package pacs_v07

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00200107 struct {
	Xmlns           string                       `xml:"xmlns,attr"`
	FIToFIPmtStsRpt FIToFIPaymentStatusReportV07 `xml:"FIToFIPmtStsRpt"`
}

func (doc DocumentPacs00200107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00200107) NameSpace() string {
	return utils.DocumentPacs00200107NameSpace
}

func (doc DocumentPacs00200107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtStsRpt FIToFIPaymentStatusReportV07 `xml:"FIToFIPmtStsRpt"`
	}
	output.FIToFIPmtStsRpt = doc.FIToFIPmtStsRpt
	utils.XmlElement(&start, doc.NameSpace())
	return e.EncodeElement(&output, start)
}
