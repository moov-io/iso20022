// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

// May be one of LVCO, LVCC, LVRT, EUSU, STSU, LWSU, EUCO, FIRE, STDY, LTNC, CRCO, RECC, LTGC, LTDC, CUSC, IBKC, SYSC, SSSC, REOP, PCOT, NPCT, ESTF
type SystemEventType2Code string

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

// May be one of FTHI, CANC, MODI, DTAU, SAIN, MINE
type CaseForwardingNotification3Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of CLSD, ASGN, INVE, UKNW, ODUE
type CaseStatus2Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// May be one of CARE, UPAR, NSSR, HPAR, THRE
type ReservationType1Code string

// May be one of CARE, UPAR, NSSR, HPAR, THRE, BLKD
type ReservationType2Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalBalanceSubType1Code string

// Must be at least 1 items long
type ExternalBalanceType1Code string

// Must be at least 1 items long
type ExternalEntryStatus1Code string

// May be one of CRED, DEBT, BOTH
type FloorLimitType1Code string

// May be one of ALLL, CHNG, MODF
type QueryType3Code string
