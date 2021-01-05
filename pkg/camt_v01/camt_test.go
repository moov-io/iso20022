// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentAuth018001V02(t *testing.T) {
	sample := DocumentCamt10100101{}
	err := sample.Validate()
	assert.Equal(t, nil, err)
}
