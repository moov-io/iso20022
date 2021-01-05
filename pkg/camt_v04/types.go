// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v04

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// Must be at least 1 items long
type ExternalSystemMemberType1Code string

// May be one of ENBL, DSBL, DLTD, JOIN
type MemberStatus1Code string

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

// May be one of X020, X030, X050
type ErrorHandling1Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalPaymentRole1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// May be one of LQMG, LMMG, PYMG, REDR, BKMG, STMG
type PaymentRole1Code string

// May be one of HIGH, NORM, LOWW
type Priority1Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

// Must be at least 1 items long
type ExternalSystemEventType1Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

// May be one of SLST, SDTL, TAPS, SLSL, SWLS
type StandingOrderQueryType1Code string

// May be one of USTO, PSTO
type StandingOrderType1Code string
