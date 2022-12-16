// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	baseLog "github.com/moov-io/base/log"
	"github.com/moov-io/iso20022/pkg/document"
	"github.com/moov-io/iso20022/pkg/server"
	"github.com/moov-io/iso20022/pkg/utils"
)

var (
	documentFileName string
	documentBuffer   []byte
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Launches web server",
	Long:  "Launches web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		env := &server.Environment{
			Logger: baseLog.NewDefaultLogger(),
		}

		env, err := server.NewEnvironment(env)
		if err != nil {
			env.Logger.Fatal().LogErrorf("Error loading up environment.", err).Err()
			os.Exit(1)
		}
		defer env.Shutdown()

		env.Logger.Info().Log("Starting services")
		test, _ := cmd.Flags().GetBool("test")
		if !test {
			shutdown := env.RunServers(true)
			defer shutdown()
		}
		return nil
	},
}

var Validate = &cobra.Command{
	Use:   "validator",
	Short: "Validate iso20022 message",
	Long:  "Validate an incoming iso20022 message",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		doc, err := document.ParseIso20022Document(documentBuffer)
		if err != nil {
			return err
		}

		err = doc.Validate()
		if err != nil {
			return err
		}

		fmt.Println("the iso20022 (" + doc.NameSpace() + ") message is valid")
		return nil
	},
}

var Print = &cobra.Command{
	Use:   "print",
	Short: "Print iso20022 message",
	Long:  "Print an incoming iso20022 message with special format (options: json, xml)",
	RunE: func(cmd *cobra.Command, args []string) error {
		var format utils.DocumentType
		if ff, err := cmd.Flags().GetString("format"); err != nil {
			return err
		} else {
			format = utils.DocumentType(ff)
		}
		if string(format) != "" && (format != utils.DocumentTypeJson && format != utils.DocumentTypeXml) {
			return errors.New("don't support the format")
		} else {
			format = utils.DocumentTypeXml
		}

		doc, err := document.ParseIso20022Document(documentBuffer)
		if err != nil {
			return err
		}

		var output []byte
		switch format {
		case utils.DocumentTypeJson:
			output, err = json.MarshalIndent(doc, "", "\t")
		case utils.DocumentTypeXml:
			output, err = xml.MarshalIndent(doc, "", "\t")
		case utils.DocumentTypeUnknown:
			err = errors.New("invalid format")
		}
		if err != nil {
			return err
		}

		fmt.Println(string(output))
		return nil
	},
}

var Convert = &cobra.Command{
	Use:   "convert [output]",
	Short: "Convert iso20022 document file format",
	Long:  "Convert an incoming iso20022 document format into another format (options: json, xml)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires output argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var format utils.DocumentType
		if ff, err := cmd.Flags().GetString("format"); err != nil {
			return err
		} else {
			format = utils.DocumentType(ff)
		}
		if string(format) != "" && (format != utils.DocumentTypeJson && format != utils.DocumentTypeXml) {
			return errors.New("don't support the format")
		} else {
			format = utils.DocumentTypeXml
		}

		doc, err := document.ParseIso20022Document(documentBuffer)
		if err != nil {
			return err
		}

		var output []byte
		switch format {
		case utils.DocumentTypeJson:
			output, err = json.MarshalIndent(doc, "", "\t")
		case utils.DocumentTypeXml:
			output, err = xml.MarshalIndent(doc, "", "\t")
		case utils.DocumentTypeUnknown:
			err = errors.New("invalid format")
		}
		if err != nil {
			return err
		}

		wFile, err := os.Create(args[0])
		if err != nil {
			return err
		}
		_, err = wFile.Write(output)
		wFile.Close()

		return err
	},
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		isWeb := false
		cmdNames := make([]string, 0)
		getName := func(c *cobra.Command) {}
		getName = func(c *cobra.Command) {
			if c == nil {
				return
			}
			cmdNames = append([]string{c.Name()}, cmdNames...)
			if c.Name() == "web" {
				isWeb = true
			}
			getName(c.Parent())
		}
		getName(cmd)

		if !isWeb {
			if documentFileName == "" {
				path, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				documentFileName = filepath.Join(path, "iso20022_document.xml")
			}

			_, err := os.Stat(documentFileName)
			if os.IsNotExist(err) {
				return errors.New("invalid input file")
			}

			documentBuffer, err = os.ReadFile(documentFileName)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func initRootCmd() {
	WebCmd.Flags().BoolP("test", "t", false, "test server")
	Convert.Flags().String("format", "xml", "format of document file")
	Print.Flags().String("format", "xml", "print format")

	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&documentFileName, "input", "", "iso20022 document (valid types are xml, json. default is $PWD/iso20022_document.xml)")
	rootCmd.AddCommand(WebCmd)
	rootCmd.AddCommand(Convert)
	rootCmd.AddCommand(Print)
	rootCmd.AddCommand(Validate)
}

func main() {
	initRootCmd()

	rootCmd.Execute()
}
