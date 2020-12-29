// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package common

import (
	"bytes"
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
type Max16Text string

func (r Max16Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 16 {
		return utils.NewErrTextLengthInvalid("Max16Text", 1, 16)
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
