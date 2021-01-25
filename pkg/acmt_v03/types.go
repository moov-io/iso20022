// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

import (
	"reflect"

	"github.com/moov-io/iso20022/pkg/utils"
)

// May be one of EMAL, FAXI, FILE, ONLI, POST
type CommunicationMethod2Code string

func (r CommunicationMethod2Code) Validate() error {
	for _, vv := range []string{
		"EMAL", "FAXI", "FILE", "ONLI", "POST",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CommunicationMethod2Code")
}

// May be one of EMAL, FAXI, POST, PHON, FILE, ONLI
type CommunicationMethod3Code string

func (r CommunicationMethod3Code) Validate() error {
	for _, vv := range []string{
		"EMAL", "FAXI", "POST", "PHON", "FILE", "ONLI",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CommunicationMethod3Code")
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
type ExternalCommunicationFormat1Code string

func (r ExternalCommunicationFormat1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCommunicationFormat1Code", 1, 4)
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
type ExternalProxyAccountType1Code string

func (r ExternalProxyAccountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalProxyAccountType1Code", 1, 4)
	}
	return nil
}

// May be one of YEAR, DAIL, MNTH, QURT, MIAN, TEND, MOVE, WEEK, INDA
type Frequency7Code string

func (r Frequency7Code) Validate() error {
	for _, vv := range []string{
		"YEAR", "DAIL", "MNTH", "QURT", "MIAN", "TEND", "MOVE", "WEEK", "INDA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Frequency7Code")
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

// Must match the pattern UNLIMITED
type Unlimited9Text string

func (r Unlimited9Text) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 9 {
		return utils.NewErrTextLengthInvalid("Unlimited9Text", 1, 9)
	}
	return nil
}

// May be one of NOCH, MODI, DELE, ADDD
type Modification1Code string

func (r Modification1Code) Validate() error {
	for _, vv := range []string{
		"NOCH", "MODI", "DELE", "ADDD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Modification1Code")
}

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

// May be one of FWNG, PREC
type BusinessDayConvention1Code string

func (r BusinessDayConvention1Code) Validate() error {
	for _, vv := range []string{
		"FWNG", "PREC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BusinessDayConvention1Code")
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
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 4)
	}
	return nil
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

// May be one of FEMA, MALE
type Gender1Code string

func (r Gender1Code) Validate() error {
	for _, vv := range []string{
		"FEMA", "MALE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Gender1Code")
}

// May be one of CIOC, CHAR, CICC, GENP, IAPS, LLPP, PCLG, LIMP, PCLS, PCLC, SOLE, UNLC, UNLT
type OrganisationLegalStatus1Code string

func (r OrganisationLegalStatus1Code) Validate() error {
	for _, vv := range []string{
		"CIOC", "CHAR", "CICC", "GENP", "IAPS", "LLPP", "PCLG", "LIMP", "PCLS", "PCLC", "SOLE", "UNLC", "UNLT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("OrganisationLegalStatus1Code")
}

// May be one of AREG, CPFA, DRLC, EMID, IDCD, NRIN, OTHR, PASS, POCD, SOCS, SRSA, GUNL
type PersonIdentificationType5Code string

func (r PersonIdentificationType5Code) Validate() error {
	for _, vv := range []string{
		"AREG", "CPFA", "DRLC", "EMID", "IDCD", "NRIN", "OTHR", "PASS", "POCD", "SOCS", "SRSA", "GUNL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PersonIdentificationType5Code")
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

// May be one of RESI, PRES, NRES
type ResidentialStatus1Code string

func (r ResidentialStatus1Code) Validate() error {
	for _, vv := range []string{
		"RESI", "PRES", "NRES",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ResidentialStatus1Code")
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

// May be one of ALPR, ALIT, GRSS
type TaxRateMarker1Code string

func (r TaxRateMarker1Code) Validate() error {
	for _, vv := range []string{
		"ALPR", "ALIT", "GRSS",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TaxRateMarker1Code")
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

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

func (r ExternalCategoryPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 5 {
		return utils.NewErrTextLengthInvalid("ExternalCategoryPurpose1Code", 1, 5)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCreditorAgentInstruction1Code string

func (r ExternalCreditorAgentInstruction1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCreditorAgentInstruction1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

func (r ExternalDiscountAmountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDiscountAmountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

func (r ExternalDocumentLineType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentLineType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalGarnishmentType1Code string

func (r ExternalGarnishmentType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalGarnishmentType1Code", 1, 4)
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
type ExternalPurpose1Code string

func (r ExternalPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPurpose1Code", 1, 4)
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

// May be one of CHK, TRF, TRA
type PaymentMethod3Code string

func (r PaymentMethod3Code) Validate() error {
	for _, vv := range []string{
		"CHK", "TRF", "TRA",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PaymentMethod3Code")
}

// May be one of OPEN, MNTN, CLSG, VIEW
type UseCases1Code string

func (r UseCases1Code) Validate() error {
	for _, vv := range []string{
		"OPEN", "MNTN", "CLSG", "VIEW",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("UseCases1Code")
}
