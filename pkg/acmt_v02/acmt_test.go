// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocumentAcmt02200102(t *testing.T) {
	sample := DocumentAcmt02200102{}
	err := sample.Validate()
	assert.Equal(t, nil, err)
}
