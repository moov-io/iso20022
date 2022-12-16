// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	testFileName        = filepath.Join("..", "..", "test", "testdata", "valid_acmt_v03.json")
	testInvalidFileName = filepath.Join("..", "..", "test", "testdata", "invalid_file1")
	testErrorFileName   = filepath.Join("..", "..", "test", "testdata", "invalid_pain_v11.json")
	testJsonFileName    = filepath.Join("..", "..", "test", "testdata", "valid_pacs_v11.json")
	testXmlFileName     = filepath.Join("..", "..", "test", "testdata", "valid_pain_v11.xml")
)

func TestMain(m *testing.M) {
	initRootCmd()
	os.Exit(m.Run())
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func deleteFile() {
	// delete file
	os.Remove("output")
}

func TestConvertWithoutInput(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--format", string(utils.DocumentTypeJson))
	if err == nil {
		t.Errorf("invalid input file")
	}
	deleteFile()
}

func TestConvertWithInvalidParam(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "--input", testFileName, "--format", string(utils.DocumentTypeJson))
	if err == nil {
		t.Errorf("requires output argument")
	}
	deleteFile()
}

func TestConvertJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testFileName, "--format", string(utils.DocumentTypeJson))
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestConvertXml(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testFileName, "--format", string(utils.DocumentTypeXml))
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestConvertUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testFileName, "--format", "unknown")
	if err == nil {
		t.Errorf("don't support the format")
	}
	deleteFile()
}

func TestPrintJson(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testFileName, "--format", string(utils.DocumentTypeJson))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintXml(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testFileName, "--format", string(utils.DocumentTypeXml))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testFileName, "--format", "unknown")
	if err == nil {
		t.Errorf("don't support the format")
	}
}

func TestValidator(t *testing.T) {
	_, err := executeCommand(rootCmd, "validator", "--input", testFileName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUnknown(t *testing.T) {
	_, err := executeCommand(rootCmd, "unknown")
	if err == nil {
		t.Errorf("don't support unknown")
	}
}

func TestPrintWithInvalidData(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testInvalidFileName, "--format", string(utils.DocumentTypeJson))
	if err == nil {
		t.Errorf("invalid file data")
	}
}

func TestConvertWithInvalidData(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testInvalidFileName, "--format", string(utils.DocumentTypeJson))
	if err == nil {
		t.Errorf("invalid file data")
	}
	deleteFile()
}

func TestValidatorWithInvalidData(t *testing.T) {
	_, err := executeCommand(rootCmd, "validator", "--input", testInvalidFileName)
	if err == nil {
		t.Errorf("error data")
	}
}

func TestPrintWithErrorData(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testErrorFileName, "--format", string(utils.DocumentTypeJson))
	if err == nil {
		t.Errorf("error data")
	}
}

func TestConvertWithErrorData(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testErrorFileName, "--format", string(utils.DocumentTypeJson))
	if err == nil {
		t.Errorf("error data")
	}
	deleteFile()
}

func TestValidatorWithErrorData(t *testing.T) {
	_, err := executeCommand(rootCmd, "validator", "--input", testErrorFileName)
	if err == nil {
		t.Errorf("error data")
	}
}

func TestPrintWithJsonData(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testJsonFileName, "--format", string(utils.DocumentTypeJson))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestConvertWithJsonData(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testJsonFileName, "--format", string(utils.DocumentTypeJson))
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestValidatorWithJsonData(t *testing.T) {
	_, err := executeCommand(rootCmd, "validator", "--input", testJsonFileName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPrintWithXmlData(t *testing.T) {
	_, err := executeCommand(rootCmd, "print", "--input", testXmlFileName, "--format", string(utils.DocumentTypeXml))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestConvertWithXmlData(t *testing.T) {
	_, err := executeCommand(rootCmd, "convert", "output", "--input", testXmlFileName, "--format", string(utils.DocumentTypeXml))
	if err != nil {
		t.Errorf(err.Error())
	}
	deleteFile()
}

func TestValidatorWithXmlData(t *testing.T) {
	_, err := executeCommand(rootCmd, "validator", "--input", testXmlFileName)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestWebTest(t *testing.T) {
	_, err := executeCommand(rootCmd, "web", "--test=true")
	if err != nil {
		t.Errorf(err.Error())
	}
}
