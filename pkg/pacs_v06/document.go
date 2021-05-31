package pacs_v06

import (
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentPacs00800106 struct {
	XMLName           xml.Name
	Attrs             []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
	FIToFICstmrCdtTrf FIToFICustomerCreditTransferV06 `xml:"FIToFICstmrCdtTrf"`
}

func (doc DocumentPacs00800106) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentPacs00800106) NameSpace() string {
	return utils.DocumentPacs00800106NameSpace
}

func (doc DocumentPacs00800106) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	α := struct {
		XMLName           xml.Name
		Attrs             []utils.Attr                    `xml:",any,attr,omitempty" json:",omitempty"`
		FIToFICstmrCdtTrf FIToFICustomerCreditTransferV06 `xml:"FIToFICstmrCdtTrf"`
	}(doc)

	utils.SettingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&α, start)
}

func (doc *DocumentPacs00800106) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *DocumentPacs00800106) GetAttrs() []utils.Attr {
	return doc.Attrs
}
