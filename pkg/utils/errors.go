// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
)

// NewErrTextLength returns a error that the length of value is invalid
func NewErrTextLengthInvalid(typeStr string, min, max int) error {
	errStr := fmt.Sprintf("The value of %s has invalid length (minLength:%d, maxLength:%d)",
		typeStr, min, max)
	if max == 0 {
		errStr = fmt.Sprintf("The value of %s has invalid length (minLength:%d)", typeStr, min)
	}
	return fmt.Errorf(errStr)
}

// NewErrTextLength returns a error that the length of value is invalid
func NewErrValueInvalid(typeStr string) error {
	errStr := fmt.Sprintf("The value of %s is invalid", typeStr)
	return fmt.Errorf(errStr)
}

// NewErrInvalidNameSpace returns a error that namespace is invalid
func NewErrInvalidNameSpace() error {
	errStr := fmt.Sprintf("The namespace of %s is invalid", "document")
	return fmt.Errorf(errStr)
}

// NewErrUnsupportedNameSpace returns a error that namespace is unsupported
func NewErrUnsupportedNameSpace() error {
	errStr := fmt.Sprintf("The namespace of %s is unsupported", "document")
	return fmt.Errorf(errStr)
}

// NewErrOmittedNameSpace returns a error that namespace is omitted
func NewErrOmittedNameSpace() error {
	errStr := fmt.Sprintf("The namespace of %s is omitted", "document")
	return fmt.Errorf(errStr)
}

// NewErrInvalidFileType returns a error that type is invalid
func NewErrInvalidFileType() error {
	errStr := fmt.Sprintf("The type of %s is invalid", "file")
	return fmt.Errorf(errStr)
}
