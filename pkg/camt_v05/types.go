// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"reflect"
	"regexp"

	"github.com/moov-io/iso20022/pkg/utils"
)

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

func (r ExternalEnquiryRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalEnquiryRequestType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

func (r ExternalMarketInfrastructure1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return utils.NewErrTextLengthInvalid("ExternalMarketInfrastructure1Code", 1, 3)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

func (r ExternalPaymentControlRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPaymentControlRequestType1Code", 1, 4)
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
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

func (r ExternalPersonIdentification1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPersonIdentification1Code", 1, 4)
	}
	return nil
}

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
type ExternalProxyAccountType1Code string

func (r ExternalProxyAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalProxyAccountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBalanceSubType1Code string

func (r ExternalBalanceSubType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBalanceSubType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalBalanceType1Code string

func (r ExternalBalanceType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalBalanceType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalEntryStatus1Code string

func (r ExternalEntryStatus1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalEntryStatus1Code", 1, 4)
	}
	return nil
}

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

func (r EntryTypeIdentifier) Validate() error {
	reg := regexp.MustCompile(`[BEOVW]{1,1}[0-9]{2,2}|DUM`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("EntryTypeIdentifier")
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

// May be one of LVCO, LVCC, LVRT, EUSU, STSU, LWSU, EUCO, FIRE, STDY, LTNC, CRCO, RECC, LTGC, LTDC, CUSC, IBKC, SYSC, SSSC, REOP, PCOT, NPCT, ESTF
type SystemEventType2Code string

func (r SystemEventType2Code) Validate() error {
	for _, vv := range []string{
		"LVCO", "LVCC", "LVRT", "EUSU", "STSU", "LWSU", "EUCO", "FIRE", "STDY", "LTNC", "CRCO", "RECC", "LTGC", "LTDC", "CUSC", "IBKC", "SYSC", "SSSC", "REOP", "PCOT", "NPCT", "ESTF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SystemEventType2Code")
}

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

func (r PaymentInstrument1Code) Validate() error {
	for _, vv := range []string{
		"BDT", "BCT", "CDT", "CCT", "CHK", "BKT", "DCP", "CCP", "RTI", "CAN",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PaymentInstrument1Code")
}

// May be one of FTHI, CANC, MODI, DTAU, SAIN, MINE
type CaseForwardingNotification3Code string

func (r CaseForwardingNotification3Code) Validate() error {
	for _, vv := range []string{
		"FTHI", "CANC", "MODI", "DTAU", "SAIN", "MINE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CaseForwardingNotification3Code")
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

// May be one of CLSD, ASGN, INVE, UKNW, ODUE
type CaseStatus2Code string

func (r CaseStatus2Code) Validate() error {
	for _, vv := range []string{
		"CLSD", "ASGN", "INVE", "UKNW", "ODUE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CaseStatus2Code")
}

// May be one of CARE, UPAR, NSSR, HPAR, THRE
type ReservationType1Code string

func (r ReservationType1Code) Validate() error {
	for _, vv := range []string{
		"CARE", "UPAR", "NSSR", "HPAR", "THRE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ReservationType1Code")
}

// May be one of CARE, UPAR, NSSR, HPAR, THRE, BLKD
type ReservationType2Code string

func (r ReservationType2Code) Validate() error {
	for _, vv := range []string{
		"CARE", "UPAR", "NSSR", "HPAR", "THRE", "BLKD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ReservationType2Code")
}

// May be one of CRED, DEBT, BOTH
type FloorLimitType1Code string

func (r FloorLimitType1Code) Validate() error {
	for _, vv := range []string{
		"CRED", "DEBT", "BOTH",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("FloorLimitType1Code")
}

// May be one of ALLL, CHNG, MODF
type QueryType3Code string

func (r QueryType3Code) Validate() error {
	for _, vv := range []string{
		"ALLL", "CHNG", "MODF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("QueryType3Code")
}
