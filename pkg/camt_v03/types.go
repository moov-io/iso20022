// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v03

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

// May be one of SLST, SDTL, TAPS, SLSL, SWLS
type StandingOrderQueryType1Code string

// May be one of USTO, PSTO
type StandingOrderType1Code string

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

// Must be at least 1 items long
type ExternalBillingBalanceType1Code string

// Must be at least 1 items long
type ExternalBillingCompensationType1Code string

// Must be at least 1 items long
type ExternalBillingRateIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// May be one of NOCP, DBTD, INVD, DDBT
type CompensationMethod1Code string

// May be one of LBOX, STOR, BILA, SEQN, MACT
type BillingSubServiceQualifier1Code string

// May be one of NTAX, MTDA, MTDB, MTDC, MTDD, UDFD
type BillingTaxCalculationMethod1Code string

// May be one of ORGN, RPLC, TEST
type BillingStatementStatus1Code string

// May be one of ACCT, STLM, PRCG
type BillingCurrencyType1Code string

// May be one of ACCT, STLM, PRCG, HOST
type BillingCurrencyType2Code string

// May be one of UPRC, STAM, BCHG, DPRC, FCHG, LPRC, MCHG, MXRD, TIR1, TIR2, TIR3, TIR4, TIR5, TIR6, TIR7, TIR8, TIR9, TPRC, ZPRC, BBSE
type BillingChargeMethod1Code string

// May be one of LDGR, FLOT, CLLD
type BalanceAdjustmentType1Code string

// May be one of INTM, SMRY
type AccountLevel1Code string

// May be one of INTM, SMRY, DETL
type AccountLevel2Code string

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of COMP, NCMP
type ServiceAdjustmentType1Code string

// May be one of BCMP, FLAT, PVCH, INVS, WVED, FREE
type ServicePaymentMethod1Code string

// May be one of XMPT, ZERO, TAXE
type ServiceTaxDesignation1Code string
