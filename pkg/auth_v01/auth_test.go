// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package auth_v01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentAuth001001V01(t *testing.T) {
	sample := DocumentAuth001001V01{}
	err := sample.Validate()
	assert.Equal(t, nil, err)
}
