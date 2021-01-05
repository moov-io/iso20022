// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package common

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"github.com/moov-io/iso20022/pkg/utils"
	"reflect"
	"regexp"
	"time"
)

// Must be at least 1 items long
type Max140Text string

func (r Max140Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 140 {
		return utils.NewErrTextLengthInvalid("Max140Text", 1, 140)
	}
	return nil
}

// Must be at least 1 items long
type Max10Text string

func (r Max10Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 10 {
		return utils.NewErrTextLengthInvalid("Max10Text", 1, 10)
	}
	return nil
}

// Must be at least 1 items long
type Max6Text string

func (r Max6Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 6 {
		return utils.NewErrTextLengthInvalid("Max6Text", 1, 6)
	}
	return nil
}

// Must be at least 1 items long
type Max16Text string

func (r Max16Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 16 {
		return utils.NewErrTextLengthInvalid("Max16Text", 1, 16)
	}
	return nil
}

// Must be at least 1 items long
type Max40Text string

func (r Max40Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 40 {
		return utils.NewErrTextLengthInvalid("Max40Text", 1, 40)
	}
	return nil
}

// Must be at least 1 items long
type Max12Text string

func (r Max12Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 12 {
		return utils.NewErrTextLengthInvalid("Max12Text", 1, 12)
	}
	return nil
}

// Must be at least 1 items long
type Max8Text string

func (r Max8Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 8 {
		return utils.NewErrTextLengthInvalid("Max8Text", 1, 8)
	}
	return nil
}

// Must be at least 1 items long
type Max25Text string

func (r Max25Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 25 {
		return utils.NewErrTextLengthInvalid("Max25Text", 1, 25)
	}
	return nil
}

// Must be at least 1 items long
type Max2048Text string

func (r Max2048Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 2048 {
		return utils.NewErrTextLengthInvalid("Max2048Text", 1, 2048)
	}
	return nil
}

// Must be at least 1 items long
type Max34Text string

func (r Max34Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 34 {
		return utils.NewErrTextLengthInvalid("Max34Text", 1, 2048)
	}
	return nil
}

// Must be at least 1 items long
type Max350Text string

func (r Max350Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 350 {
		return utils.NewErrTextLengthInvalid("Max350Text", 1, 350)
	}
	return nil
}

// Must be at least 1 items long
type Max35Text string

func (r Max35Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 35 {
		return utils.NewErrTextLengthInvalid("Max35Text", 1, 35)
	}
	return nil
}

// Must be at least 1 items long
type Max500Text string

func (r Max500Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 500 {
		return utils.NewErrTextLengthInvalid("Max500Text", 1, 500)
	}
	return nil
}

// Must be at least 1 items long
type Max70Text string

func (r Max70Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 70 {
		return utils.NewErrTextLengthInvalid("Max70Text", 1, 70)
	}
	return nil
}

// Must be at least 1 items long
type Max1025Text string

func (r Max1025Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 1025 {
		return utils.NewErrTextLengthInvalid("Max1025Text", 1, 1025)
	}
	return nil
}

// Must be at least 1 items long
type Max256Text string

func (r Max256Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 256 {
		return utils.NewErrTextLengthInvalid("Max256Text", 1, 256)
	}
	return nil
}

// Must be at least 1 items long
type Max4Text string

func (r Max4Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("Max4Text", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type Max128Text string

func (r Max128Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 128 {
		return utils.NewErrTextLengthInvalid("Max128Text", 1, 128)
	}
	return nil
}

// Must be at least 1 items long
type Max105Text string

func (r Max105Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 105 {
		return utils.NewErrTextLengthInvalid("Max105Text", 1, 105)
	}
	return nil
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

func (r PhoneNumber) Validate() error {
	reg := regexp.MustCompile(`\+[0-9]{1,3}-[0-9()+\-]{1,30}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("PhoneNumber")
	}
	return nil
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

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
type CreditDebitCode string

func (r CreditDebitCode) Validate() error {
	for _, vv := range []string{
		"CRDT", "DBIT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CreditDebitCode")
}

// May be one of CODU, COPY, DUPL
type CopyDuplicate1Code string

func (r CopyDuplicate1Code) Validate() error {
	for _, vv := range []string{
		"CODU", "COPY", "DUPL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CopyDuplicate1Code")
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

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

func (r Max15NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{1,15}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Max15NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

func (r Max3NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{1,3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Max3NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{4,4}
type MerchantCategoryCodeIdentifier string

func (r MerchantCategoryCodeIdentifier) Validate() error {
	reg := regexp.MustCompile(`[0-9]{4,4}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("MerchantCategoryCodeIdentifier")
	}
	return nil
}

// May be one of FIXE, USGB, VARI
type MandateClassification1Code string

func (r MandateClassification1Code) Validate() error {
	for _, vv := range []string{
		"FIXE", "USGB", "VARI",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("MandateClassification1Code")
}

// May be one of INDY, OVRN
type InterestType1Code string

func (r InterestType1Code) Validate() error {
	for _, vv := range []string{
		"INDY", "OVRN",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("InterestType1Code")
}

// Must match the pattern [\+]{0,1}[0-9]{1,15}
type Max15PlusSignedNumericText string

func (r Max15PlusSignedNumericText) Validate() error {
	reg := regexp.MustCompile(`[\+]{0,1}[0-9]{1,15}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Max15PlusSignedNumericText")
	}
	return nil
}

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

func (r Max4AlphaNumericText) Validate() error {
	reg := regexp.MustCompile(`[a-zA-Z0-9]{1,4}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Max4AlphaNumericText")
	}
	return nil
}

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

func (r Max5NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{1,5}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Max5NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

func (r Min2Max3NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{2,3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Min2Max3NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{3,4}
type Min3Max4NumericText string

func (r Min3Max4NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{3,4}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Min3Max4NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

func (r Min8Max28NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{8,28}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Min8Max28NumericText")
	}
	return nil
}

// May be one of AUTH, FDET, FSUM, ILEV
type Authorisation1Code string

func (r Authorisation1Code) Validate() error {
	for _, vv := range []string{
		"AUTH", "FDET", "FSUM", "ILEV",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Authorisation1Code")
}

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

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

func (r AnyBICIdentifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("AnyBICIdentifier")
	}
	return nil
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

func (r BICFIIdentifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("BICFIIdentifier")
	}
	return nil
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

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

func (r NamePrefix1Code) Validate() error {
	for _, vv := range []string{
		"DOCT", "MIST", "MISS", "MADM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("NamePrefix1Code")
}

// May be one of ENAB, DISA, DELE, FORM
type AccountStatus3Code string

func (r AccountStatus3Code) Validate() error {
	for _, vv := range []string{
		"ENAB", "DISA", "DELE", "FORM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AccountStatus3Code")
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

func (r UUIDv4Identifier) Validate() error {
	reg := regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("UUIDv4Identifier")
	}
	return nil
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("15:04:05.999999999")), nil
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
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

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01")), nil
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type Max10MbBinary []byte

func (t *Max10MbBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10MbBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}
