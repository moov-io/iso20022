// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"fmt"
)

// NewErrTextLength returns a error that the length of value is invalid
func NewErrTextLengthInvalid(typeStr string, min, max int) error {
	errStr := fmt.Sprintf("The value of %s has invalid length (minLength:%d, maxLength:%d)",
		typeStr, min, max)
	return fmt.Errorf(errStr)
}

// NewErrTextLength returns a error that the length of value is invalid
func NewErrValueInvalid(typeStr string) error {
	errStr := fmt.Sprintf("The value of %s is invalid", typeStr)
	return fmt.Errorf(errStr)
}

// NewErrInvalidNameSpace returns a error that namespace is invalid
func NewErrInvalidNameSpace() error {
	return errors.New("The namespace of document is invalid")
}
