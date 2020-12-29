// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v02

import (
	"bytes"
	"encoding/base64"
	"github.com/moov-io/iso20022/pkg/utils"
	"reflect"
	"regexp"
)

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

func (r ActiveCurrencyCode) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{3,3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ActiveCurrencyCode")
	}
	return nil
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

func (r ActiveOrHistoricCurrencyCode) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{3,3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ActiveOrHistoricCurrencyCode")
	}
	return nil
}

func (r AddressType2Code) Validate() error {
	for _, vv := range []string{
		"ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AddressType2Code")
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

func (r AnyBICDec2014Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("AnyBICDec2014Identifier")
	}
	return nil
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

func (r BICFIDec2014Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("BICFIDec2014Identifier")
	}
	return nil
}

// May be one of WIBO, TREA, TIBO, TLBO, SWAP, STBO, PRBO, PFAN, NIBO, MAAA, MOSP, LIBO, LIBI, JIBA, ISDA, GCFR, FUSW, EUCH, EUUS, EURI, EONS, EONA, CIBO, CDOR, BUBO, BBSW
type BenchmarkCurveName2Code string

func (r BenchmarkCurveName2Code) Validate() error {
	for _, vv := range []string{
		"WIBO", "TREA", "TIBO", "TLBO", "SWAP", "STBO", "PRBO", "PFAN", "NIBO", "MAAA", "MOSP", "LIBO", "LIBI",
		"JIBA", "ISDA", "GCFR", "FUSW", "EUCH", "EUUS", "EURI", "EONS", "EONA", "CIBO", "CDOR", "BUBO", "BBSW",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BenchmarkCurveName2Code")
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

func (r CountryCode) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("CountryCode")
	}
	return nil
}

// May be one of CRDT, DBIT
type CreditDebit3Code string

func (r CreditDebit3Code) Validate() error {
	for _, vv := range []string{
		"CRDT", "DBIT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CreditDebit3Code")
}

// May be one of FITE, CALL
type DepositType1Code string

func (r DepositType1Code) Validate() error {
	for _, vv := range []string{
		"FITE", "CALL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("DepositType1Code")
}

// Must match the pattern [0-9]
type Exact1NumericText string

func (r Exact1NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Exact1NumericText")
	}
	return nil
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

func (r Exact4AlphaNumericText) Validate() error {
	reg := regexp.MustCompile(`[a-zA-Z0-9]{4}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Exact4AlphaNumericText")
	}
	return nil
}

// Must match the pattern [0-9]{5}
type Exact5NumericText string

func (r Exact5NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{5}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Exact5NumericText")
	}
	return nil
}

// May be one of SPOT, SALE, AGRD
type ExchangeRateType1Code string

func (r ExchangeRateType1Code) Validate() error {
	for _, vv := range []string{
		"SPOT", "SALE", "AGRD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ExchangeRateType1Code")
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

func (r ExternalAccountIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalAccountIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 5 {
		return utils.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 5)
	}
	return nil
}

// Must be at least 1 items long
type ExternalContractBalanceType1Code string

func (r ExternalContractBalanceType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalContractBalanceType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentType1Code string

func (r ExternalDocumentType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

func (r ExternalPersonIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPersonIdentification1Code", 1, 4)
	}
	return nil
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

func (r IBAN2007Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("IBAN2007Identifier")
	}
	return nil
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

func (r ISINOct2015Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ISINOct2015Identifier")
	}
	return nil
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

func (r LEIIdentifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z0-9]{18,18}[0-9]{2,2}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("LEIIdentifier")
	}
	return nil
}

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

func (r NamePrefix2Code) Validate() error {
	for _, vv := range []string{
		"DOCT", "MADM", "MISS", "MIST", "MIKS",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("NamePrefix2Code")
}

// May be one of CNTR, ESTM
type PaymentScheduleType1Code string

func (r PaymentScheduleType1Code) Validate() error {
	for _, vv := range []string{
		"CNTR", "ESTM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PaymentScheduleType1Code")
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

func (r PreferredContactMethod1Code) Validate() error {
	for _, vv := range []string{
		"LETT", "MAIL", "PHON", "FAXX", "CELL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PreferredContactMethod1Code")
}

// May be one of HIGH, NORM
type Priority2Code string

func (r Priority2Code) Validate() error {
	for _, vv := range []string{
		"HIGH", "NORM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Priority2Code")
}

// May be one of DAYS, MNTH, WEEK, YEAR
type RateBasis1Code string

func (r RateBasis1Code) Validate() error {
	for _, vv := range []string{
		"DAYS", "MNTH", "WEEK", "YEAR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("RateBasis1Code")
}

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

func (r Max15NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{1,15}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Max15NumericText")
	}
	return nil
}

// May be one of EMAL, FAXI, FILE, ONLI, PHON, POST, PROP, SWMT, SWMX
type CommunicationMethod4Code string

func (r CommunicationMethod4Code) Validate() error {
	for _, vv := range []string{
		"EMAL", "FAXI", "FILE", "ONLI", "PHON", "POST", "PROP", "SWMT", "SWMX",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CommunicationMethod4Code")
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

func (r TaxExemptReason1Code) Validate() error {
	for _, vv := range []string{
		"NONE", "MASA", "MISA", "SISA", "IISA", "CUYP", "PRYP", "ASTR", "EMPY", "EMCY", "EPRY", "ECYE", "NFPI",
		"NFQP", "DECP", "IRAC", "IRAR", "KEOG", "PFSP", "401K", "SIRA", "403B", "457X", "RIRA", "RIAN", "RCRF",
		"RCIP", "EIFP", "EIOP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TaxExemptReason1Code")
}

// Must be at least 1 items long
type ExternalContractClosureReason1Code string

func (r ExternalContractClosureReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalContractClosureReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashAccountType1Code string

func (r ExternalCashAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCashAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

func (r ExternalProxyAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalProxyAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalShipmentCondition1Code string

func (r ExternalShipmentCondition1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalShipmentCondition1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalValidationRuleIdentification1Code string

func (r ExternalValidationRuleIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalValidationRuleIdentification1Code", 1, 4)
	}
	return nil
}

// May be one of ALLL, CHNG, MODF
type QueryType3Code string

func (r QueryType3Code) Validate() error {
	for _, vv := range []string{
		"ALLL", "CHNG", "MODF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("QueryType3Code")
}

// May be one of LFBK, LTBK, SUPP
type SupportDocumentType1Code string

func (r SupportDocumentType1Code) Validate() error {
	for _, vv := range []string{
		"LFBK", "LTBK", "SUPP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SupportDocumentType1Code")
}

// Must be at least 1 items long
type ExternalStatusReason1Code string

// May be one of ACPT, ACTC, PART, PDNG, RCVD, RJCT, RMDR, INCF, CRPT
type StatisticalReportingStatus1Code string

func (r StatisticalReportingStatus1Code) Validate() error {
	for _, vv := range []string{
		"ACPT", "ACTC", "PART", "PDNG", "RCVD", "RJCT", "RMDR", "INCF", "CRPT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("StatisticalReportingStatus1Code")
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}
