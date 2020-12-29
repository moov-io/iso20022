// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v01

import (
	"reflect"

	"github.com/moov-io/iso20022/pkg/utils"
)

// May be one of DAYH, EARL
type BalanceTransferWindow1Code string

func (r BalanceTransferWindow1Code) Validate() error {
	for _, vv := range []string{
		"DAYH", "EARL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BalanceTransferWindow1Code")
}

// May be one of ACPT, BTRQ, BTRS, COMP, REDT, REDE, REJT, REQU, TMTN
type SwitchStatus1Code string

func (r SwitchStatus1Code) Validate() error {
	for _, vv := range []string{
		"ACPT", "BTRQ", "BTRS", "COMP", "REDT", "REDE", "REJT", "REQU", "TMTN",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SwitchStatus1Code")
}

// May be one of FULL, PART
type SwitchType1Code string

func (r SwitchType1Code) Validate() error {
	for _, vv := range []string{
		"FULL", "PART",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SwitchType1Code")
}
