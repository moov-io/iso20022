// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v05

import (
	"reflect"
	"regexp"

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
type ExternalAuthenticationChannel1Code string

func (r ExternalAuthenticationChannel1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalAuthenticationChannel1Code", 1, 4)
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
type ExternalCategoryPurpose1Code string

func (r ExternalCategoryPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCategoryPurpose1Code", 1, 4)
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
type ExternalLocalInstrument1Code string

func (r ExternalLocalInstrument1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 35 {
		return utils.NewErrTextLengthInvalid("ExternalLocalInstrument1Code", 1, 35)
	}
	return nil
}

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

func (r ExternalMandateSetupReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalMandateSetupReason1Code", 1, 4)
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
type ExternalServiceLevel1Code string

func (r ExternalServiceLevel1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalServiceLevel1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalMandateReason1Code string

func (r ExternalMandateReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalMandateReason1Code", 1, 4)
	}
	return nil
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

// May be one of NEVR, YEAR, RATE, MIAN, QURT
type Frequency10Code string

func (r Frequency10Code) Validate() error {
	for _, vv := range []string{
		"NEVR", "YEAR", "RATE", "MIAN", "QURT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Frequency10Code")
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

// May be one of RCUR, OOFF
type SequenceType2Code string

func (r SequenceType2Code) Validate() error {
	for _, vv := range []string{
		"RCUR", "OOFF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SequenceType2Code")
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

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

func (r RemittanceLocationMethod2Code) Validate() error {
	for _, vv := range []string{
		"FAXI", "EDIC", "URID", "EMAL", "POST", "SMSM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("RemittanceLocationMethod2Code")
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

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

func (r AnyBICIdentifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("AnyBICIdentifier")
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

// May be one of MLDB, MLCD, MLFA, CRDB, CRCD, CRFA, PUDB, PUCD, PUFA, RGDB, RGCD, RGFA
type ChequeDelivery1Code string

func (r ChequeDelivery1Code) Validate() error {
	for _, vv := range []string{
		"MLDB", "MLCD", "MLFA", "CRDB", "CRCD", "CRFA", "PUDB", "PUCD", "PUFA", "RGDB", "RGCD", "RGFA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ChequeDelivery1Code")
}

// May be one of CCHQ, CCCH, BCHQ, DRFT, ELDR
type ChequeType2Code string

func (r ChequeType2Code) Validate() error {
	for _, vv := range []string{
		"CCHQ", "CCCH", "BCHQ", "DRFT", "ELDR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ChequeType2Code")
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

// May be one of CHK, TRF
type PaymentMethod7Code string

func (r PaymentMethod7Code) Validate() error {
	for _, vv := range []string{
		"CHK", "TRF",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PaymentMethod7Code")
}

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

func (r RegulatoryReportingType1Code) Validate() error {
	for _, vv := range []string{
		"CRED", "DEBT", "BOTH",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("RegulatoryReportingType1Code")
}

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

func (r ExternalDiscountAmountType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalDiscountAmountType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

func (r ExternalDocumentLineType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentLineType1Code", 1, 0)
	}
	return nil
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
type ExternalPurpose1Code string

func (r ExternalPurpose1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalPurpose1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalStatusReason1Code string

func (r ExternalStatusReason1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalStatusReason1Code", 1, 0)
	}
	return nil
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

// May be one of ACTC, RCVD, PART, RJCT, PDNG, ACCP, ACSP, ACSC, ACWC
type TransactionGroupStatus3Code string

func (r TransactionGroupStatus3Code) Validate() error {
	for _, vv := range []string{
		"ACTC", "RCVD", "PART", "RJCT", "PDNG", "ACCP", "ACSP", "ACSC", "ACWC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TransactionGroupStatus3Code")
}

// May be one of ACTC, RJCT, PDNG, ACCP, ACSP, ACSC, ACWC
type TransactionIndividualStatus3Code string

func (r TransactionIndividualStatus3Code) Validate() error {
	for _, vv := range []string{
		"ACTC", "RJCT", "PDNG", "ACCP", "ACSP", "ACSC", "ACWC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TransactionIndividualStatus3Code")
}
