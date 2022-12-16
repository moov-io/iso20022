// Copyright 2021 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moov-io/iso20022/pkg/document"
	"github.com/moov-io/iso20022/pkg/utils"
)

func outputError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func outputSuccess(w http.ResponseWriter, output string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": output,
	})
}

func parseInputFromRequest(r *http.Request) (document.Iso20022Document, error) {
	inputFile, _, err := r.FormFile("input")
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	var input bytes.Buffer
	if _, err = io.Copy(&input, inputFile); err != nil {
		return nil, err
	}

	return document.ParseIso20022Document(input.Bytes())
}

func messageToBuf(format utils.DocumentType, doc document.Iso20022Document) ([]byte, error) {
	var output []byte
	var err error
	switch format {
	case utils.DocumentTypeJson:
		output, err = json.MarshalIndent(doc, "", "\t")
	case utils.DocumentTypeXml:
		output, err = xml.MarshalIndent(doc, "", "\t")
	case utils.DocumentTypeUnknown:
		err = errors.New("unknown document type")
	}
	return output, err
}

func outputBufferToWriter(w http.ResponseWriter, doc document.Iso20022Document, format utils.DocumentType) {
	w.WriteHeader(http.StatusOK)
	switch format {
	case utils.DocumentTypeJson:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(doc)
	case utils.DocumentTypeXml:
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		xml.NewEncoder(w).Encode(doc)
	case utils.DocumentTypeUnknown:
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.Write([]byte(`{"error": "invalid format"}`))
	}
}

func getFormat(r *http.Request) (utils.DocumentType, error) {
	var format utils.DocumentType
	ff := r.FormValue("format")
	if ff == "" {
		format = utils.DocumentTypeXml
	} else {
		format = utils.DocumentType(ff)
	}
	if format != utils.DocumentTypeXml && format != utils.DocumentTypeJson {
		return format, fmt.Errorf("%s is an invalid format: %v", ff, format)
	}
	return format, nil
}

// validator - validate the file based on publication 1220
func validator(w http.ResponseWriter, r *http.Request) {
	doc, err := parseInputFromRequest(r)
	if err != nil {
		outputError(w, http.StatusBadRequest, err)
		return
	}

	err = doc.Validate()
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	outputSuccess(w, "valid file")
}

// validator - print file with ascii or json format
func print(w http.ResponseWriter, r *http.Request) {
	doc, err := parseInputFromRequest(r)
	if err != nil {
		outputError(w, http.StatusBadRequest, err)
		return
	}

	format, err := getFormat(r)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}
	_, err = messageToBuf(format, doc)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	outputBufferToWriter(w, doc, format)
}

// convert - convert file with ascii or json format
func convert(w http.ResponseWriter, r *http.Request) {
	message, err := parseInputFromRequest(r)
	if err != nil {
		outputError(w, http.StatusBadRequest, err)
		return
	}

	format, err := getFormat(r)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	output, err := messageToBuf(format, message)
	if err != nil {
		outputError(w, http.StatusNotImplemented, err)
		return
	}

	filename := "converted_file"
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

// health - health check
func health(w http.ResponseWriter, r *http.Request) {
	outputSuccess(w, "alive")
}

// configure handlers
func ConfigureHandlers(r *mux.Router) error {
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/print", print).Methods("POST")
	r.HandleFunc("/validator", validator).Methods("POST")
	r.HandleFunc("/convert", convert).Methods("POST")
	return nil
}
