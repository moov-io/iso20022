// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentAcmt03600101(t *testing.T) {
	sample := DocumentAcmt03600101{}
	err := sample.Validate()
	assert.Equal(t, nil, err)
}
