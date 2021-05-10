// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admi_v01

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentAdmi00200101 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	Admi00200101 Admi00200101 `xml:"admi.002.001.01"`
}

func (doc DocumentAdmi00200101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00200101) NameSpace() string {
	return utils.DocumentAdmi00200101NameSpace
}

func (doc DocumentAdmi00200101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
		Admi00200101 Admi00200101 `xml:"admi.002.001.01"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi00400101 struct {
	XMLName      xml.Name
	Attrs        []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	Admi00400101 Admi00400101 `xml:"admi.004.001.01"`
}

func (doc DocumentAdmi00400101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00400101) NameSpace() string {
	return utils.DocumentAdmi00400101NameSpace
}

func (doc DocumentAdmi00400101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName      xml.Name
		Attrs        []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
		Admi00400101 Admi00400101 `xml:"admi.004.001.01"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi00500101 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr          `xml:",any,attr,omitempty" json:",omitempty"`
	RptQryReq ReportQueryRequestV01 `xml:"RptQryReq"`
}

func (doc DocumentAdmi00500101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00500101) NameSpace() string {
	return utils.DocumentAdmi00500101NameSpace
}

func (doc DocumentAdmi00500101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr          `xml:",any,attr,omitempty" json:",omitempty"`
		RptQryReq ReportQueryRequestV01 `xml:"RptQryReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi00600101 struct {
	XMLName xml.Name
	Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	RsndReq ResendRequestV01 `xml:"RsndReq"`
}

func (doc DocumentAdmi00600101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00600101) NameSpace() string {
	return utils.DocumentAdmi00600101NameSpace
}

func (doc DocumentAdmi00600101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
		RsndReq ResendRequestV01 `xml:"RsndReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi00700101 struct {
	XMLName xml.Name
	Attrs   []utils.Attr              `xml:",any,attr,omitempty" json:",omitempty"`
	RctAck  ReceiptAcknowledgementV01 `xml:"RctAck"`
}

func (doc DocumentAdmi00700101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi00700101) NameSpace() string {
	return utils.DocumentAdmi00700101NameSpace
}

func (doc DocumentAdmi00700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr              `xml:",any,attr,omitempty" json:",omitempty"`
		RctAck  ReceiptAcknowledgementV01 `xml:"RctAck"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi01100101 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
	SysEvtAck SystemEventAcknowledgementV01 `xml:"SysEvtAck"`
}

func (doc DocumentAdmi01100101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi01100101) NameSpace() string {
	return utils.DocumentAdmi01100101NameSpace
}

func (doc DocumentAdmi01100101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr                  `xml:",any,attr,omitempty" json:",omitempty"`
		SysEvtAck SystemEventAcknowledgementV01 `xml:"SysEvtAck"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}

type DocumentAdmi01700101 struct {
	XMLName xml.Name
	Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	PrcgReq ProcessingRequestV01 `xml:"PrcgReq"`
}

func (doc DocumentAdmi01700101) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentAdmi01700101) NameSpace() string {
	return utils.DocumentAdmi01700101NameSpace
}

func (doc DocumentAdmi01700101) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		PrcgReq ProcessingRequestV01 `xml:"PrcgReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name = doc.XMLName
	}
	return e.EncodeElement(&α, start)
}
