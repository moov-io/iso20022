package pacs_v06

import (
	"reflect"
	"regexp"

	"github.com/moov-io/iso20022/pkg/utils"
)

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

func (r ActiveOrHistoricCurrencyCode) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{3,3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ActiveOrHistoricCurrencyCode")
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

// May be one of URGT, HIGH, NORM
type Priority3Code string

func (r Priority3Code) Validate() error {
	for _, vv := range []string{
		"URGT", "HIGH", "NORM",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Priority3Code")
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
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalServiceLevel1Code string

func (r ExternalServiceLevel1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalServiceLevel1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

func (r ExternalAccountIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalAccountIdentification1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashAccountType1Code string

func (r ExternalCashAccountType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalCashAccountType1Code", 1, 0)
	}
	return nil
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
type ExternalClearingSystemIdentification1Code string

func (r ExternalClearingSystemIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalClearingSystemIdentification1Code", 1, 0)
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

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

func (r ExternalDocumentLineType1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentLineType1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

func (r ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalFinancialInstitutionIdentification1Code", 1, 0)
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
type ExternalLocalInstrument1Code string

func (r ExternalLocalInstrument1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalLocalInstrument1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

func (r ExternalOrganisationIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalOrganisationIdentification1Code", 1, 0)
	}
	return nil
}

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

func (r ExternalPersonIdentification1Code) Validate() error {
	if len(string(r)) < 1 {
		return utils.NewErrTextLengthInvalid("ExternalPersonIdentification1Code", 1, 0)
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
