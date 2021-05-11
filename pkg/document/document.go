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
}

// data representation attributes
var availableDocuments = map[string]Iso20022Document{
	utils.DocumentAcmt03600101NameSpace: &acmt_v01.DocumentAcmt03600101{},
	utils.DocumentAcmt02200102NameSpace: &acmt_v02.DocumentAcmt02200102{},
	utils.DocumentAcmt02300102NameSpace: &acmt_v02.DocumentAcmt02300102{},
	utils.DocumentAcmt02400102NameSpace: &acmt_v02.DocumentAcmt02400102{},
	utils.DocumentAcmt03000102NameSpace: &acmt_v02.DocumentAcmt03000102{},
	utils.DocumentAcmt03300102NameSpace: &acmt_v02.DocumentAcmt03300102{},
	utils.DocumentAcmt03500102NameSpace: &acmt_v02.DocumentAcmt03500102{},
	utils.DocumentAcmt03700102NameSpace: &acmt_v02.DocumentAcmt03700102{},
	utils.DocumentAcmt00700103NameSpace: &acmt_v03.DocumentAcmt00700103{},
	utils.DocumentAcmt00800103NameSpace: &acmt_v03.DocumentAcmt00800103{},
	utils.DocumentAcmt00900103NameSpace: &acmt_v03.DocumentAcmt00900103{},
	utils.DocumentAcmt01000103NameSpace: &acmt_v03.DocumentAcmt01000103{},
	utils.DocumentAcmt01100103NameSpace: &acmt_v03.DocumentAcmt01100103{},
	utils.DocumentAcmt01200103NameSpace: &acmt_v03.DocumentAcmt01200103{},
	utils.DocumentAcmt01300103NameSpace: &acmt_v03.DocumentAcmt01300103{},
	utils.DocumentAcmt01400103NameSpace: &acmt_v03.DocumentAcmt01400103{},
	utils.DocumentAcmt01500103NameSpace: &acmt_v03.DocumentAcmt01500103{},
	utils.DocumentAcmt01600103NameSpace: &acmt_v03.DocumentAcmt01600103{},
	utils.DocumentAcmt01700103NameSpace: &acmt_v03.DocumentAcmt01700103{},
	utils.DocumentAcmt01800103NameSpace: &acmt_v03.DocumentAcmt01800103{},
	utils.DocumentAcmt01900103NameSpace: &acmt_v03.DocumentAcmt01900103{},
	utils.DocumentAcmt02000103NameSpace: &acmt_v03.DocumentAcmt02000103{},
	utils.DocumentAcmt02100103NameSpace: &acmt_v03.DocumentAcmt02100103{},
	utils.DocumentAcmt02700103NameSpace: &acmt_v03.DocumentAcmt02700103{},
	utils.DocumentAcmt02800103NameSpace: &acmt_v03.DocumentAcmt02800103{},
	utils.DocumentAcmt02900103NameSpace: &acmt_v03.DocumentAcmt02900103{},
	utils.DocumentAcmt03100103NameSpace: &acmt_v03.DocumentAcmt03100103{},
	utils.DocumentAcmt03200103NameSpace: &acmt_v03.DocumentAcmt03200103{},
	utils.DocumentAcmt03400103NameSpace: &acmt_v03.DocumentAcmt03400103{},
	utils.DocumentAdmi00200101NameSpace: &admi_v01.DocumentAdmi00200101{},
	utils.DocumentAdmi00400101NameSpace: &admi_v01.DocumentAdmi00400101{},
	utils.DocumentAdmi00500101NameSpace: &admi_v01.DocumentAdmi00500101{},
	utils.DocumentAdmi00600101NameSpace: &admi_v01.DocumentAdmi00600101{},
	utils.DocumentAdmi00700101NameSpace: &admi_v01.DocumentAdmi00700101{},
	utils.DocumentAdmi01100101NameSpace: &admi_v01.DocumentAdmi01100101{},
	utils.DocumentAdmi01700101NameSpace: &admi_v01.DocumentAdmi01700101{},
	utils.DocumentAdmi00400102NameSpace: &admi_v02.DocumentAdmi00400102{},
	utils.DocumentAdmi00900102NameSpace: &admi_v02.DocumentAdmi00900102{},
	utils.DocumentAdmi01000102NameSpace: &admi_v02.DocumentAdmi01000102{},
	utils.DocumentAuth00100101NameSpace: &auth_v01.DocumentAuth00100101{},
	utils.DocumentAuth00200101NameSpace: &auth_v01.DocumentAuth00200101{},
	utils.DocumentAuth00300101NameSpace: &auth_v01.DocumentAuth00300101{},
	utils.DocumentAuth01800102NameSpace: &auth_v02.DocumentAuth01800102{},
	utils.DocumentAuth01900102NameSpace: &auth_v02.DocumentAuth01900102{},
	utils.DocumentAuth02000102NameSpace: &auth_v02.DocumentAuth02000102{},
	utils.DocumentAuth02100102NameSpace: &auth_v02.DocumentAuth02100102{},
	utils.DocumentAuth02200102NameSpace: &auth_v02.DocumentAuth02200102{},
	utils.DocumentAuth02300102NameSpace: &auth_v02.DocumentAuth02300102{},
	utils.DocumentAuth02400102NameSpace: &auth_v02.DocumentAuth02400102{},
	utils.DocumentAuth02500102NameSpace: &auth_v02.DocumentAuth02500102{},
	utils.DocumentAuth02600102NameSpace: &auth_v02.DocumentAuth02600102{},
	utils.DocumentAuth02700102NameSpace: &auth_v02.DocumentAuth02700102{},
	utils.DocumentCamt10100101NameSpace: &camt_v01.DocumentCamt10100101{},
	utils.DocumentCamt10200101NameSpace: &camt_v01.DocumentCamt10200101{},
	utils.DocumentCamt10300101NameSpace: &camt_v01.DocumentCamt10300101{},
	utils.DocumentCamt10400101NameSpace: &camt_v01.DocumentCamt10400101{},
	utils.DocumentCamt03500103NameSpace: &camt_v03.DocumentCamt03500103{},
	utils.DocumentCamt06900103NameSpace: &camt_v03.DocumentCamt06900103{},
	utils.DocumentCamt07100103NameSpace: &camt_v03.DocumentCamt07100103{},
	utils.DocumentCamt08600103NameSpace: &camt_v03.DocumentCamt08600103{},
	utils.DocumentCamt01300104NameSpace: &camt_v04.DocumentCamt01300104{},
	utils.DocumentCamt01400104NameSpace: &camt_v04.DocumentCamt01400104{},
	utils.DocumentCamt01500104NameSpace: &camt_v04.DocumentCamt01500104{},
	utils.DocumentCamt01600104NameSpace: &camt_v04.DocumentCamt01600104{},
	utils.DocumentCamt01700104NameSpace: &camt_v04.DocumentCamt01700104{},
	utils.DocumentCamt02000104NameSpace: &camt_v04.DocumentCamt02000104{},
	utils.DocumentCamt03200104NameSpace: &camt_v04.DocumentCamt03200104{},
	utils.DocumentCamt03800104NameSpace: &camt_v04.DocumentCamt03800104{},
	utils.DocumentCamt07000104NameSpace: &camt_v04.DocumentCamt07000104{},
	utils.DocumentCamt01800105NameSpace: &camt_v05.DocumentCamt01800105{},
	utils.DocumentCamt02500105NameSpace: &camt_v05.DocumentCamt02500105{},
	utils.DocumentCamt02600105NameSpace: &camt_v05.DocumentCamt02600105{},
	utils.DocumentCamt02800105NameSpace: &camt_v05.DocumentCamt02800105{},
	utils.DocumentCamt03000105NameSpace: &camt_v05.DocumentCamt03000105{},
	utils.DocumentCamt03500105NameSpace: &camt_v05.DocumentCamt03500105{},
	utils.DocumentCamt03600105NameSpace: &camt_v05.DocumentCamt03600105{},
	utils.DocumentCamt03900105NameSpace: &camt_v05.DocumentCamt03900105{},
	utils.DocumentCamt04600105NameSpace: &camt_v05.DocumentCamt04600105{},
	utils.DocumentCamt04800105NameSpace: &camt_v05.DocumentCamt04800105{},
	utils.DocumentCamt04900105NameSpace: &camt_v05.DocumentCamt04900105{},
	utils.DocumentCamt05000105NameSpace: &camt_v05.DocumentCamt05000105{},
	utils.DocumentCamt05100105NameSpace: &camt_v05.DocumentCamt05100105{},
	utils.DocumentCamt05600105NameSpace: &camt_v05.DocumentCamt05600105{},
	utils.DocumentCamt06000105NameSpace: &camt_v05.DocumentCamt06000105{},
	utils.DocumentCamt02100106NameSpace: &camt_v06.DocumentCamt02100106{},
	utils.DocumentCamt02400106NameSpace: &camt_v06.DocumentCamt02400106{},
	utils.DocumentCamt02900106NameSpace: &camt_v06.DocumentCamt02900106{},
	utils.DocumentCamt03100106NameSpace: &camt_v06.DocumentCamt03100106{},
	utils.DocumentCamt03300106NameSpace: &camt_v06.DocumentCamt03300106{},
	utils.DocumentCamt03400106NameSpace: &camt_v06.DocumentCamt03400106{},
	utils.DocumentCamt04700106NameSpace: &camt_v06.DocumentCamt04700106{},
	utils.DocumentCamt05700106NameSpace: &camt_v06.DocumentCamt05700106{},
	utils.DocumentCamt05800106NameSpace: &camt_v06.DocumentCamt05800106{},
	utils.DocumentCamt05900106NameSpace: &camt_v06.DocumentCamt05900106{},
	utils.DocumentCamt00300107NameSpace: &camt_v07.DocumentCamt00300107{},
	utils.DocumentCamt00900107NameSpace: &camt_v07.DocumentCamt00900107{},
	utils.DocumentCamt01100107NameSpace: &camt_v07.DocumentCamt01100107{},
	utils.DocumentCamt01200107NameSpace: &camt_v07.DocumentCamt01200107{},
	utils.DocumentCamt01900107NameSpace: &camt_v07.DocumentCamt01900107{},
	utils.DocumentCamt02300107NameSpace: &camt_v07.DocumentCamt02300107{},
	utils.DocumentCamt08700107NameSpace: &camt_v07.DocumentCamt08700107{},
	utils.DocumentCamt00400108NameSpace: &camt_v08.DocumentCamt00400108{},
	utils.DocumentCamt00500108NameSpace: &camt_v08.DocumentCamt00500108{},
	utils.DocumentCamt00600108NameSpace: &camt_v08.DocumentCamt00600108{},
	utils.DocumentCamt00700108NameSpace: &camt_v08.DocumentCamt00700108{},
	utils.DocumentCamt00800108NameSpace: &camt_v08.DocumentCamt00800108{},
	utils.DocumentCamt01000108NameSpace: &camt_v08.DocumentCamt01000108{},
	utils.DocumentCamt02600108NameSpace: &camt_v08.DocumentCamt02600108{},
	utils.DocumentCamt02700108NameSpace: &camt_v08.DocumentCamt02700108{},
	utils.DocumentCamt03700108NameSpace: &camt_v08.DocumentCamt03700108{},
	utils.DocumentCamt05200108NameSpace: &camt_v08.DocumentCamt05200108{},
	utils.DocumentCamt05300108NameSpace: &camt_v08.DocumentCamt05300108{},
	utils.DocumentCamt05400108NameSpace: &camt_v08.DocumentCamt05400108{},
	utils.DocumentCamt05500109NameSpace: &camt_v09.DocumentCamt05500109{},
	utils.DocumentCamt05600109NameSpace: &camt_v09.DocumentCamt05600109{},
	utils.DocumentCamt02800110NameSpace: &camt_v10.DocumentCamt02800110{},
	utils.DocumentCamt02900110NameSpace: &camt_v10.DocumentCamt02900110{},
	utils.DocumentHead00100101NameSpace: &head_v01.BusinessApplicationHeaderV01{},
	utils.DocumentHead00100102NameSpace: &head_v02.BusinessApplicationHeaderV02{},
	utils.DocumentPacs01000104NameSpace: &pacs_v04.DocumentPacs01000104{},
	utils.DocumentPacs02800104NameSpace: &pacs_v04.DocumentPacs02800104{},
	utils.DocumentPacs00200107NameSpace: &pacs_v07.DocumentPacs00200107{},
	utils.DocumentPacs00300108NameSpace: &pacs_v08.DocumentPacs00300108{},
	utils.DocumentPacs00800109NameSpace: &pacs_v09.DocumentPacs00800109{},
	utils.DocumentPacs00900109NameSpace: &pacs_v09.DocumentPacs00900109{},
	utils.DocumentPacs00200110NameSpace: &pacs_v10.DocumentPacs00200110{},
	utils.DocumentPacs00400110NameSpace: &pacs_v10.DocumentPacs00400110{},
	utils.DocumentPacs00700110NameSpace: &pacs_v10.DocumentPacs00700110{},
	utils.DocumentPacs00200111NameSpace: &pacs_v11.DocumentPacs00200111{},
	utils.DocumentPain00700101NameSpace: &pain_v01.DocumentPain00700101{},
	utils.DocumentPain01800101NameSpace: &pain_v01.DocumentPain01800101{},
	utils.DocumentPain00900105NameSpace: &pain_v05.DocumentPain00900105{},
	utils.DocumentPain01000105NameSpace: &pain_v05.DocumentPain01000105{},
	utils.DocumentPain01100105NameSpace: &pain_v05.DocumentPain01100105{},
	utils.DocumentPain01200105NameSpace: &pain_v05.DocumentPain01200105{},
	utils.DocumentPain01300105NameSpace: &pain_v05.DocumentPain01300105{},
	utils.DocumentPain01400105NameSpace: &pain_v05.DocumentPain01400105{},
	utils.DocumentPain01400108NameSpace: &pain_v08.DocumentPain01400108{},
	utils.DocumentPain01300108NameSpace: &pain_v08.DocumentPain01300108{},
	utils.DocumentPain00800109NameSpace: &pain_v09.DocumentPain00800109{},
	utils.DocumentPain00100110NameSpace: &pain_v10.DocumentPain00100110{},
	utils.DocumentPain00700110NameSpace: &pain_v10.DocumentPain00700110{},
	utils.DocumentPain00200111NameSpace: &pain_v11.DocumentPain00200111{},
	utils.DocumentReda06600101NameSpace: &reda_v01.DocumentReda06600101{},
	utils.DocumentReda06700101NameSpace: &reda_v01.DocumentReda06700101{},
	utils.DocumentReda06800101NameSpace: &reda_v01.DocumentReda06800101{},
	utils.DocumentReda06900101NameSpace: &reda_v01.DocumentReda06900101{},
	utils.DocumentReda07000101NameSpace: &reda_v01.DocumentReda07000101{},
	utils.DocumentReda07100101NameSpace: &reda_v01.DocumentReda07100101{},
	utils.DocumentReda07200101NameSpace: &reda_v01.DocumentReda07200101{},
	utils.DocumentReda07300101NameSpace: &reda_v01.DocumentReda07300101{},
	utils.DocumentRemt00100102NameSpace: &remt_v02.DocumentRemt00100102{},
	utils.DocumentRemt00200102NameSpace: &remt_v02.DocumentRemt00200102{},
	utils.DocumentRemt00100104NameSpace: &remt_v04.DocumentRemt00100104{},
}

type documentDummy struct {
	XMLName xml.Name
	Attrs   []utils.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

// ParseIso20022Document will return a interface of ISO 20022 document after pass buffer
func ParseIso20022Document(buf []byte) (doc Iso20022Document, err error) {
	bType := utils.GetBufferFormat(buf)
	if bType == utils.DocumentTypeUnknown {
		return nil, utils.NewErrInvalidFileType()
	}

	var exit bool

	var dummy documentDummy
	if bType == utils.DocumentTypeJson {
		err = json.Unmarshal(buf, &dummy)
	} else if bType == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, &dummy)
	}
	if err != nil {
		return nil, err
	}

	var namespace string
	if len(dummy.XMLName.Space) > 0 {
		namespace = dummy.XMLName.Space
	}

	if len(namespace) == 0 {
		for _, attr := range dummy.Attrs {
			if attr.Name.Local == utils.XmlDefaultNamespace {
				namespace = attr.Value
			}
		}
	}

	if namespace == "" {
		return nil, utils.NewErrOmittedNameSpace()
	}

	doc, exit = availableDocuments[namespace]
	if !exit {
		return nil, utils.NewErrUnsupportedNameSpace()
	}

	if bType == utils.DocumentTypeJson {
		err = json.Unmarshal(buf, doc)
	} else if bType == utils.DocumentTypeXml {
		err = xml.Unmarshal(buf, doc)
	}

	if err != nil {
		return nil, err
	}

	return
}
