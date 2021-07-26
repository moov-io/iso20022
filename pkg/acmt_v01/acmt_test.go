// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	var type1 BalanceTransferWindow1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.NotNil(t, type1.Validate())
	type1 = "DAYH"
	assert.Nil(t, type1.Validate())

	var type2 SwitchStatus1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.NotNil(t, type2.Validate())
	type2 = "TMTN"
	assert.Nil(t, type2.Validate())

	var type3 SwitchType1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.NotNil(t, type3.Validate())
	type3 = "PART"
	assert.Nil(t, type3.Validate())
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountSwitchDetails1{}.Validate())
	assert.NotNil(t, AccountSwitchTerminationSwitchV01{}.Validate())
	assert.NotNil(t, MessageIdentification1{}.Validate())
	assert.NotNil(t, ResponseDetails1{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
}
