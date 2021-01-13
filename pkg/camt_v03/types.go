// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

import (
	"reflect"

	"github.com/moov-io/iso20022/pkg/utils"
)

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

func (r ExternalAccountIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalAccountIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashAccountType1Code string

func (r ExternalCashAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCashAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 5 {
		return utils.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 5)
	}
	return nil
}

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

func (r ExternalProxyAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalProxyAccountType1Code", 1, 4)
	}
	return nil
}

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

func (r QueryType2Code) Validate() error {
	for _, vv := range []string{
		"ALLL", "CHNG", "MODF", "DELD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("QueryType2Code")
}

// May be one of SLST, SDTL, TAPS, SLSL, SWLS
type StandingOrderQueryType1Code string

func (r StandingOrderQueryType1Code) Validate() error {
	for _, vv := range []string{
		"SLST", "SDTL", "TAPS", "SLSL", "SWLS",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("StandingOrderQueryType1Code")
}

// May be one of USTO, PSTO
type StandingOrderType1Code string

func (r StandingOrderType1Code) Validate() error {
	for _, vv := range []string{
		"USTO", "PSTO",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("StandingOrderType1Code")
}

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

func (r ExternalBankTransactionDomain1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBankTransactionDomain1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

func (r ExternalBankTransactionFamily1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBankTransactionFamily1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

func (r ExternalBankTransactionSubFamily1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBankTransactionSubFamily1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBillingBalanceType1Code string

func (r ExternalBillingBalanceType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBillingBalanceType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBillingCompensationType1Code string

func (r ExternalBillingCompensationType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBillingCompensationType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBillingRateIdentification1Code string

func (r ExternalBillingRateIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBillingRateIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 4)
	}
	return nil
}

// May be one of NOCP, DBTD, INVD, DDBT
type CompensationMethod1Code string

func (r CompensationMethod1Code) Validate() error {
	for _, vv := range []string{
		"NOCP", "DBTD", "INVD", "DDBT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CompensationMethod1Code")
}

// May be one of LBOX, STOR, BILA, SEQN, MACT
type BillingSubServiceQualifier1Code string

func (r BillingSubServiceQualifier1Code) Validate() error {
	for _, vv := range []string{
		"LBOX", "STOR", "BILA", "SEQN", "MACT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BillingSubServiceQualifier1Code")
}

// May be one of NTAX, MTDA, MTDB, MTDC, MTDD, UDFD
type BillingTaxCalculationMethod1Code string

func (r BillingTaxCalculationMethod1Code) Validate() error {
	for _, vv := range []string{
		"NTAX", "MTDA", "MTDB", "MTDC", "MTDD", "UDFD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BillingTaxCalculationMethod1Code")
}

// May be one of ORGN, RPLC, TEST
type BillingStatementStatus1Code string

func (r BillingStatementStatus1Code) Validate() error {
	for _, vv := range []string{
		"ORGN", "RPLC", "TEST",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BillingStatementStatus1Code")
}

// May be one of ACCT, STLM, PRCG
type BillingCurrencyType1Code string

func (r BillingCurrencyType1Code) Validate() error {
	for _, vv := range []string{
		"ACCT", "STLM", "PRCG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BillingCurrencyType1Code")
}

// May be one of ACCT, STLM, PRCG, HOST
type BillingCurrencyType2Code string

func (r BillingCurrencyType2Code) Validate() error {
	for _, vv := range []string{
		"ACCT", "STLM", "PRCG", "HOST",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BillingCurrencyType2Code")
}

// May be one of UPRC, STAM, BCHG, DPRC, FCHG, LPRC, MCHG, MXRD, TIR1, TIR2, TIR3, TIR4, TIR5, TIR6, TIR7, TIR8, TIR9, TPRC, ZPRC, BBSE
type BillingChargeMethod1Code string

func (r BillingChargeMethod1Code) Validate() error {
	for _, vv := range []string{
		"UPRC", "STAM", "BCHG", "DPRC", "FCHG", "LPRC", "MCHG", "MXRD", "TIR1", "TIR2", "TIR3", "TIR4", "TIR5", "TIR6", "TIR7", "TIR8", "TIR9", "TPRC", "ZPRC", "BBSE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BillingChargeMethod1Code")
}

// May be one of LDGR, FLOT, CLLD
type BalanceAdjustmentType1Code string

func (r BalanceAdjustmentType1Code) Validate() error {
	for _, vv := range []string{
		"LDGR", "FLOT", "CLLD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BalanceAdjustmentType1Code")
}

// May be one of INTM, SMRY
type AccountLevel1Code string

func (r AccountLevel1Code) Validate() error {
	for _, vv := range []string{
		"INTM", "SMRY",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AccountLevel1Code")
}

// May be one of INTM, SMRY, DETL
type AccountLevel2Code string

func (r AccountLevel2Code) Validate() error {
	for _, vv := range []string{
		"INTM", "SMRY", "DETL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AccountLevel2Code")
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

func (r PreferredContactMethod1Code) Validate() error {
	for _, vv := range []string{
		"LETT", "MAIL", "PHON", "FAXX", "CELL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PreferredContactMethod1Code")
}

// May be one of COMP, NCMP
type ServiceAdjustmentType1Code string

func (r ServiceAdjustmentType1Code) Validate() error {
	for _, vv := range []string{
		"COMP", "NCMP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ServiceAdjustmentType1Code")
}

// May be one of BCMP, FLAT, PVCH, INVS, WVED, FREE
type ServicePaymentMethod1Code string

func (r ServicePaymentMethod1Code) Validate() error {
	for _, vv := range []string{
		"BCMP", "FLAT", "PVCH", "INVS", "WVED", "FREE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ServicePaymentMethod1Code")
}

// May be one of XMPT, ZERO, TAXE
type ServiceTaxDesignation1Code string

func (r ServiceTaxDesignation1Code) Validate() error {
	for _, vv := range []string{
		"XMPT", "ZERO", "TAXE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ServiceTaxDesignation1Code")
}
