// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package reda_v01

import (
	"reflect"

	"github.com/moov-io/iso20022/pkg/utils"
)

// Must be at least 1 items long
type ExternalDocumentFormat1Code string

func (r ExternalDocumentFormat1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentFormat1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDocumentType1Code string

func (r ExternalDocumentType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDocumentType1Code", 1, 4)
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
type ExternalCreditorEnrolmentAmendmentReason1Code string

func (r ExternalCreditorEnrolmentAmendmentReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCreditorEnrolmentAmendmentReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCreditorEnrolmentCancellationReason1Code string

func (r ExternalCreditorEnrolmentCancellationReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCreditorEnrolmentCancellationReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCreditorEnrolmentStatusReason1Code string

func (r ExternalCreditorEnrolmentStatusReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCreditorEnrolmentCancellationReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDebtorActivationAmendmentReason1Code string

func (r ExternalDebtorActivationAmendmentReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDebtorActivationAmendmentReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDebtorActivationCancellationReason1Code string

func (r ExternalDebtorActivationCancellationReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDebtorActivationCancellationReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalDebtorActivationStatusReason1Code string

func (r ExternalDebtorActivationStatusReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalDebtorActivationStatusReason1Code", 1, 4)
	}
	return nil
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

// May be one of ACPT, RJCT
type ServiceRequestStatus1Code string

func (r ServiceRequestStatus1Code) Validate() error {
	for _, vv := range []string{
		"ACPT", "RJCT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ServiceRequestStatus1Code")
}

// May be one of FULL, PAYD
type PresentmentType1Code string

func (r PresentmentType1Code) Validate() error {
	for _, vv := range []string{
		"FULL", "PAYD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PresentmentType1Code")
}
