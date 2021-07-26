// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package document

import (
	"encoding/json"
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/utils"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonXmlWithDocumentCamt05500109(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_camt_v09.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_camt_v09.json"))
	assert.Equal(t, nil, err)

	doc, err := NewDocument(utils.DocumentCamt05500109NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentCamt05500109NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))
}

func TestJsonXmlWithDocumentAcmt00700103(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_acmt_v03.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_acmt_v03.json"))
	assert.Equal(t, nil, err)

	doc, err := NewDocument(utils.DocumentAcmt00700103NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentAcmt00700103NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))
}

func TestJsonXmlWithDocumentAuth01800102(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_auth_v02.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_auth_v02.json"))
	assert.Equal(t, nil, err)

	doc, err := NewDocument(utils.DocumentAuth01800102NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentAuth01800102NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))
}

func TestJsonXmlWithDocumentPacs00200111(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_pacs_v11.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_pacs_v11.json"))
	assert.Equal(t, nil, err)

	doc, err := NewDocument(utils.DocumentPacs00200111NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentPacs00200111NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))
}

func TestJsonXmlWithDocumentPain00200111(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_pain_v11.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_pain_v11.json"))
	assert.Equal(t, nil, err)

	doc, err := NewDocument(utils.DocumentPain00200111NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentPain00200111NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))
}

func TestJsonXmlWithDocumentReda06600101(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_reda_v01.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_reda_v01.json"))
	assert.Equal(t, nil, err)

	doc, err := NewDocument(utils.DocumentReda06600101NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentReda06600101NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Equal(t, nil, err)
	assert.Equal(t, expectJson, string(buf))

}

func TestJsonXmlWithDocumentRemt00100104(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_remt_v04.xml"))
	assert.Nil(t, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_remt_v04.json"))
	assert.Nil(t, err)

	doc, err := NewDocument(utils.DocumentRemt00100104NameSpace)
	assert.Equal(t, nil, err)
	err = xml.Unmarshal(inputXml, doc)
	assert.Nil(t, err)
	assert.Nil(t, doc.Validate())

	expectXml := strings.ReplaceAll(string(inputXml), "\r\n", "\n")
	expectJson := strings.ReplaceAll(string(inputJson), "\r\n", "\n")

	buf, err := xml.MarshalIndent(doc, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectJson, string(buf))

	doc, err = NewDocument(utils.DocumentRemt00100104NameSpace)
	assert.Equal(t, nil, err)
	err = json.Unmarshal(inputJson, doc)
	assert.Nil(t, err)
	assert.Nil(t, doc.Validate())

	buf, err = xml.MarshalIndent(doc, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectXml, string(buf))
	buf, err = json.MarshalIndent(doc, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectJson, string(buf))

}

func TestJsonXmlWithDummy(t *testing.T) {
	inputXml, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_remt_v04.xml"))
	assert.Equal(t, nil, err)

	inputJson, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "valid_remt_v04.json"))
	assert.Equal(t, nil, err)

	var docInterface Iso20022Document
	docInterface, err = ParseIso20022Document(inputJson)
	assert.Equal(t, nil, err)
	assert.NotNil(t, docInterface)
	assert.Equal(t, nil, docInterface.Validate())

	docInterface, err = ParseIso20022Document(inputXml)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, docInterface.Validate())
}

func TestJsonXmlWithFiles(t *testing.T) {
	validFileList := []string{
		"valid_acmt_v03.xml",
		"valid_auth_v02.xml",
		"valid_camt_v09.xml",
		"valid_pacs_v11.xml",
		"valid_pain_v11.xml",
		"valid_reda_v01.xml",
		"valid_remt_v04.xml",
		"valid_acmt_v03.json",
		"valid_auth_v02.json",
		"valid_camt_v09.json",
		"valid_pacs_v11.json",
		"valid_pain_v11.json",
		"valid_reda_v01.json",
		"valid_remt_v04.json",
	}

	for _, fileName := range validFileList {
		input, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", fileName))
		assert.Equal(t, nil, err)

		var docInterface Iso20022Document
		docInterface, err = ParseIso20022Document(input)
		assert.Nil(t, err)
		assert.Nil(t, docInterface.Validate())
	}

	unsupportedFileList := []string{
		"FI_camt_052_sample.xml.xml",
		"FI_camt_053_sample.xml.xml",
		"FI_camt_054_sample.xml.xml",
		"200519_camt.054-Debit_P_CH2909000000250094239_1110092692_0_2019042401501580.xml",
		"200519_camt.054-Credit_P_CH2909000000250094239_1110092691_0_2019042421291293.xml",
		"200519_camt.052_P_CH2909000000250094239_1110092686_0_2019042416072347.xml",
		"200519_camt.054_P_CH2909000000250094239_1111091335_0_2020061900081727.xml",
		"200924_camt.054_P_CH2909000000250094239_1111091335_0_2020061900081727.xml",
		"200924_camt.054_P_CH2909000000250094239_1110092703_0_2019042423412214.xml",
		"200519_camt054-ESR-ASR_P_CH2909000000250094239_1110092704_0_2019042500372179.xml",
		"camt.053_P_CH2909000000250094239_1110092698_0_2020112503071366.xml",
		"200519_camt054-chdd_p_ch2909000000250094239_1110097484_0_20190520700381159.xml",
		"pain002-chdd-cor1_p_ch2909000000250094239_1110097483_0_2018031317221082.xml",
		"pain008-musterfile.xml",
		"camt054-returns_p_ch5109000000250092291_1110097606_0_2020112500512470.xml",
		"200519_camt054-returns_p_ch2909000000250094239_1109800798_0_2019052023472022.xml",
		"200519_camt054-epo_p_ch5109000000250092291_1110097605_0_2019052003522556.xml",
		"200519_camt054-epo_p_ch5109000000250092291_1110097605_0_2019052103322231.xml",
		"pain002-epo_p_0_0_0_2018032614401842.xml",
		"pain002-epo_p_0_0_0_2018031510491259.xml",
		"pain002-epo_p_ch5109000000250092291_1110097605_0_2018031511252307.xml",
		"pain002-epo_p_ch2909000000250094239_1109800799_0_2018032612092784.xml",
		"musterfile_pain.001_Nov2020.xml",
		"gistfile1.xml",
		"statement_1.xml",
		"International_payment_RUB_naujas_1.xml",
		"International_payment_USD_naujas_1.xml",
		"sepa_payment_naujas_1.xml",
	}

	for _, fileName := range unsupportedFileList {
		input, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", fileName))
		assert.Nil(t, err)

		_, err = ParseIso20022Document(input)
		assert.NotNil(t, err)
		assert.Equal(t, "The namespace of document is unsupported", err.Error())
	}

	fileListWithoutNamespace := []string{
		"invalid_pain_v11.json",
		"invalid_pain_v11.xml",
	}

	for _, fileName := range fileListWithoutNamespace {
		input, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", fileName))
		assert.Nil(t, err)

		_, err = ParseIso20022Document(input)
		assert.NotNil(t, err)
		assert.Equal(t, "The namespace of document is omitted", err.Error())
	}

	invalidFileList := []string{
		"invalid_file1",
		"invalid_file2",
	}

	for _, fileName := range invalidFileList {
		input, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", fileName))
		assert.Nil(t, err)

		_, err = ParseIso20022Document(input)
		assert.NotNil(t, err)
		assert.Equal(t, "The type of file is invalid", err.Error())
	}
}
