// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v06

import (
	"reflect"

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
type ExternalPaymentControlRequestType1Code string

func (r ExternalPaymentControlRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPaymentControlRequestType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

func (r ExternalSystemErrorHandling1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalSystemErrorHandling1Code", 1, 4)
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

// Must be at least 1 items long
type ExternalSystemEventType1Code string

func (r ExternalSystemEventType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalSystemEventType1Code", 1, 4)
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
type ExternalMarketInfrastructure1Code string

func (r ExternalMarketInfrastructure1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return utils.NewErrTextLengthInvalid("ExternalMarketInfrastructure1Code", 1, 3)
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
type ExternalPurpose1Code string

func (r ExternalPurpose1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPurpose1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 4)
	}
	return nil
}

// May be one of HIGH, NORM, LOWW
type Priority1Code string

func (r Priority1Code) Validate() error {
	for _, vv := range []string{
		"HIGH", "NORM", "LOWW",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Priority1Code")
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

func (r Frequency2Code) Validate() error {
	for _, vv := range []string{
		"YEAR", "MNTH", "QURT", "MIAN", "WEEK", "DAIL", "ADHO", "INDA", "OVNG",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Frequency2Code")
}

// May be one of NFND, NAUT, UKNW, PCOR, WMSG, RNCR, MROI
type InvestigationRejection1Code string

func (r InvestigationRejection1Code) Validate() error {
	for _, vv := range []string{
		"NFND", "NAUT", "UKNW", "PCOR", "WMSG", "RNCR", "MROI",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("InvestigationRejection1Code")
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

// May be one of ENAB, DISA, DELD, REQD, BLKD
type ReservationStatus1Code string

func (r ReservationStatus1Code) Validate() error {
	for _, vv := range []string{
		"ENAB", "DISA", "DELD", "REQD", "BLKD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ReservationStatus1Code")
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

// May be one of RCBD, RCVD, NRCD
type NotificationStatus3Code string

func (r NotificationStatus3Code) Validate() error {
	for _, vv := range []string{
		"RCBD", "RCVD", "NRCD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("NotificationStatus3Code")
}
