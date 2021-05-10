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

// May be one of DUPL, AGNT, CURR, CUST, UPAY, CUTA, TECH, FRAD
type CancellationReason5Code string

func (r CancellationReason5Code) Validate() error {
	for _, vv := range []string{
		"DUPL", " AGNT", " CURR", " CUST", " UPAY", " CUTA", " TECH", " FRAD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CancellationReason5Code")
}

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

func (r DocumentType3Code) Validate() error {
	for _, vv := range []string{
		"RADM", "RPIN", "FXDR", "DISP", "PUOR", "SCOR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("DocumentType3Code")
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

func (r AddressType2Code) Validate() error {
	for _, vv := range []string{
		"ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AddressType2Code")
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

func (r DocumentType6Code) Validate() error {
	for _, vv := range []string{
		"MSIN", "CNFA", "DNFA", "CINV", "CREN", "DEBN", "HIRI", "SBIN", "CMCN", "SOAC", "DISP", "BOLD", "VCHR", "AROI", "TSUT", "PUOR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("DocumentType6Code")
}

// Must be at least 1 items long
type ExternalServiceLevel1Code string

func (r ExternalServiceLevel1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalServiceLevel1Code", 1, 0)
	}
	return nil
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

func (r SettlementMethod1Code) Validate() error {
	for _, vv := range []string{
		"INDA", "INGA", "COVE", "CLRG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SettlementMethod1Code")
}

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 0)
	}
	return nil
}

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

func (r TaxRecordPeriod1Code) Validate() error {
	for _, vv := range []string{
		"MM01", "MM02", "MM03", "MM04", "MM05", "MM06", "MM07", "MM08", "MM09", "MM10", "MM11", "MM12", "QTR1", "QTR2", "QTR3", "QTR4", "HLF1", "HLF2",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TaxRecordPeriod1Code")
}

// May be one of IN01, IN02, IN03, IN04, IN05, IN06, IN07, IN08, IN09, IN10, IN11, IN12, IN13, IN14, IN15, IN16, IN17, IN18, IN19, MM20, MM21, MM22, MM25, MM26, MM27, MM28, MM29, MM30, MM31, MM32, IN33, MM34, MM35, IN36, IN37, IN38, IN39, NARR
type UnableToApplyIncorrectInformation4Code string

func (r UnableToApplyIncorrectInformation4Code) Validate() error {
	for _, vv := range []string{
		"IN01", "IN02", "IN03", "IN04", "IN05", "IN06", "IN07", "IN08", "IN09", "IN10", "IN11", "IN12", "IN13",
		"IN14", "IN15", "IN16", "IN17", "IN18", "IN19", "MM20", "MM21", "MM22", "MM25", "MM26", "MM27", "MM28",
		"MM29", "MM30", "MM31", "MM32", "IN33", "MM34", "MM35", "IN36", "IN37", "IN38", "IN39", "NARR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("UnableToApplyIncorrectInformation4Code")
}

// May be one of MS01, MS02, MS03, MS04, MS05, MS06, MS07, MS08, MS09, MS10, MS11, MS12, MS13, MS14, MS15, MS16, MS17, NARR
type UnableToApplyMissingInformation3Code string

func (r UnableToApplyMissingInformation3Code) Validate() error {
	for _, vv := range []string{
		"MS01", "MS02", "MS03", "MS04", "MS05", "MS06", "MS07", "MS08", "MS09", "MS10", "MS11", "MS12", "MS13", "MS14", "MS15", "MS16", "MS17", "NARR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("UnableToApplyMissingInformation3Code")
}

// Must be at least 1 items long
type ExternalPurpose1Code string

func (r ExternalPurpose1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalPurpose1Code", 1, 0)
	}
	return nil
}

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

func (r ChargeBearerType1Code) Validate() error {
	for _, vv := range []string{
		"DEBT", "CRED", "SHAR", "SLEV",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ChargeBearerType1Code")
}

// May be one of CHQB, HOLD, PHOB, TELB
type Instruction3Code string

func (r Instruction3Code) Validate() error {
	for _, vv := range []string{
		"CHQB", "HOLD", "PHOB", "TELB",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Instruction3Code")
}

// May be one of PHOA, TELA
type Instruction4Code string

func (r Instruction4Code) Validate() error {
	for _, vv := range []string{
		"PHOA", "TELA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Instruction4Code")
}

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

func (r PaymentMethod4Code) Validate() error {
	for _, vv := range []string{
		"CHK", "TRF", "DD", "TRA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PaymentMethod4Code")
}

// May be one of HIGH, NORM
type Priority2Code string

func (r Priority2Code) Validate() error {
	for _, vv := range []string{
		"HIGH", "NORM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Priority2Code")
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

func (r ClearingChannel2Code) Validate() error {
	for _, vv := range []string{
		"RTGS", "RTNS", "MPNS", "BOOK",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ClearingChannel2Code")
}

// May be one of FRST, RCUR, FNAL, OOFF, RPRE
type SequenceType3Code string

func (r SequenceType3Code) Validate() error {
	for _, vv := range []string{
		"FRST", "RCUR", "FNAL", "OOFF", "RPRE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SequenceType3Code")
}

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

func (r ExternalCategoryPurpose1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalCategoryPurpose1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

func (r ExternalCashClearingSystem1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalCashClearingSystem1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

func (r ExternalDiscountAmountType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalDiscountAmountType1Code", 1, 0)
	}
	return nil
}

// May be one of CRDT, DBIT
type CreditDebitCode string

func (r CreditDebitCode) Validate() error {
	for _, vv := range []string{
		"CRDT", "DBIT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CreditDebitCode")
}

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

func (r ExternalDocumentLineType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentLineType1Code", 1, 0)
	}
	return nil
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

func (r Frequency6Code) Validate() error {
	for _, vv := range []string{
		"YEAR", "MNTH", "QURT", "MIAN", "WEEK", "DAIL", "ADHO", "INDA", "FRTN",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Frequency6Code")
}

// Must be at least 1 items long
type ExternalGarnishmentType1Code string

func (r ExternalGarnishmentType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalGarnishmentType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

func (r ExternalLocalInstrument1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalLocalInstrument1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

func (r ExternalMandateSetupReason1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalMandateSetupReason1Code"+
			"", 1, 0)
	}
	return nil
}
