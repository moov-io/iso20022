// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v01

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

// May be one of MULT, BILI, MAND, DISC, NELI, INBI, GLBL, DIDB, SPLC, SPLF, TDLC, TDLF, UCDT, ACOL, EXGT
type LimitType3Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemEventType1Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

// May be one of CARE, UPAR, NSSR, HPAR, THRE, BLKD
type ReservationType2Code string

// May be one of LQMG, LMMG, PYMG, REDR, BKMG, STMG
type PaymentRole1Code string
