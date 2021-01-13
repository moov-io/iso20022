// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

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
type ExternalEnquiryRequestType1Code string

func (r ExternalEnquiryRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalEnquiryRequestType1Code", 1, 4)
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
type ExternalPaymentControlRequestType1Code string

func (r ExternalPaymentControlRequestType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalPaymentControlRequestType1Code", 1, 4)
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

// Must be at least 1 items long
type ExternalSystemBalanceType1Code string

func (r ExternalSystemBalanceType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalSystemBalanceType1Code", 1, 4)
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
type ExternalSystemEventType1Code string

func (r ExternalSystemEventType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalSystemEventType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

func (r ExternalCashClearingSystem1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 3 {
		return utils.NewErrTextLengthInvalid("ExternalCashClearingSystem1Code", 1, 3)
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
type ExternalCancellationReason1Code string

func (r ExternalCancellationReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCancellationReason1Code", 1, 4)
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
type ExternalMandateSetupReason1Code string

func (r ExternalMandateSetupReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalMandateSetupReason1Code", 1, 4)
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

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

func (r ExternalTaxAmountType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalTaxAmountType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalAgentInstruction1Code string

func (r ExternalAgentInstruction1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalAgentInstruction1Code", 1, 4)
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
type ExternalCardTransactionCategory1Code string

func (r ExternalCardTransactionCategory1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCardTransactionCategory1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalChargeType1Code string

func (r ExternalChargeType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalChargeType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalCreditLineType1Code string

func (r ExternalCreditLineType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalCreditLineType1Code", 1, 4)
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

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

func (r ExternalFinancialInstrumentIdentificationType1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalFinancialInstrumentIdentificationType1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalRePresentmentReason1Code string

func (r ExternalRePresentmentReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalRePresentmentReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalReportingSource1Code string

func (r ExternalReportingSource1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalReportingSource1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalReturnReason1Code string

func (r ExternalReturnReason1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalReturnReason1Code", 1, 4)
	}
	return nil
}

// Must be at least 1 items long
type ExternalTechnicalInputChannel1Code string

func (r ExternalTechnicalInputChannel1Code) Validate() error {
	if len(string(r)) < 1 || len(string(r)) > 4 {
		return utils.NewErrTextLengthInvalid("ExternalTechnicalInputChannel1Code", 1, 4)
	}
	return nil
}

// May be one of PDNG, STLD
type BalanceStatus1Code string

func (r BalanceStatus1Code) Validate() error {
	for _, vv := range []string{
		"PDNG", "STLD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("BalanceStatus1Code")
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

// May be one of RJCT, CVHD, RSVT, BLCK, EARM, EFAC, DLVR, COLD, CSDB
type ProcessingType1Code string

func (r ProcessingType1Code) Validate() error {
	for _, vv := range []string{
		"RJCT", "CVHD", "RSVT", "BLCK", "EARM", "EFAC", "DLVR", "COLD", "CSDB",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ProcessingType1Code")
}

// May be one of USTO, PSTO
type StandingOrderType1Code string

func (r StandingOrderType1Code) Validate() error {
	for _, vv := range []string{
		"USTO", "PSTO",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("StandingOrderType1Code")
}

// May be one of OPNG, INTM, CLSG, BOOK, CRRT, PDNG, LRLD, AVLB, LTSF, CRDT, EAST, PYMT, BLCK, XPCD, DLOD, XCRD, XDBT, ADJT, PRAV, DBIT, THRE, NOTE, FSET, BLOC, OTHB, CUST, FORC, COLC, FUND, PIPO, XCHG, CCPS, TOHB, COHB, DOHB, TPBL, CPBL, DPBL, FUTB, REJB, FCOL, FCOU, SCOL, SCOU, CUSA, XCHC, XCHN, DSET, LACK, NSET, OTCC, OTCG, OTCN, SAPD, SAPC, REPD, REPC, BSCD, BSCC, SAPP, IRLT, IRDR, DWRD, ADWR, AIDR
type SystemBalanceType2Code string

func (r SystemBalanceType2Code) Validate() error {
	for _, vv := range []string{
		"OPNG", "INTM", "CLSG", "BOOK", "CRRT", "PDNG", "LRLD", "AVLB", "LTSF", "CRDT", "EAST", "PYMT", "BLCK", "XPCD", "DLOD", "XCRD", "XDBT", "ADJT", "PRAV", "DBIT", "THRE", "NOTE", "FSET", "BLOC", "OTHB", "CUST", "FORC", "COLC", "FUND", "PIPO", "XCHG", "CCPS", "TOHB", "COHB", "DOHB", "TPBL", "CPBL", "DPBL", "FUTB", "REJB", "FCOL", "FCOU", "SCOL", "SCOU", "CUSA", "XCHC", "XCHN", "DSET", "LACK", "NSET", "OTCC", "OTCG", "OTCN", "SAPD", "SAPC", "REPD", "REPC", "BSCD", "BSCC", "SAPP", "IRLT", "IRDR", "DWRD", "ADWR", "AIDR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SystemBalanceType2Code")
}

// May be one of PDNG, FINL
type CashPaymentStatus2Code string

func (r CashPaymentStatus2Code) Validate() error {
	for _, vv := range []string{
		"PDNG", "FINL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CashPaymentStatus2Code")
}

// May be one of BOOK, PDNG, FUTR
type EntryStatus1Code string

func (r EntryStatus1Code) Validate() error {
	for _, vv := range []string{
		"BOOK", "PDNG", "FUTR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("EntryStatus1Code")
}

// May be one of STLD, RJTD, CAND, FNLD
type FinalStatusCode string

func (r FinalStatusCode) Validate() error {
	for _, vv := range []string{
		"STLD", "RJTD", "CAND", "FNLD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("FinalStatusCode")
}

// May be one of PBEN, TTIL, TFRO
type Instruction1Code string

func (r Instruction1Code) Validate() error {
	for _, vv := range []string{
		"PBEN", "TTIL", "TFRO",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Instruction1Code")
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

// May be one of CBS, BCK, BAL, CLS, CTR, CBH, CBP, DPG, DPN, EXP, TCH, LMT, LIQ, DPP, DPH, DPS, STF, TRP, TCS, LOA, LOR, TCP, OND, MGL
type PaymentType3Code string

func (r PaymentType3Code) Validate() error {
	for _, vv := range []string{
		"CBS", "BCK", "BAL", "CLS", "CTR", "CBH", "CBP", "DPG", "DPN", "EXP", "TCH", "LMT", "LIQ", "DPP", "DPH", "DPS", "STF", "TRP", "TCS", "LOA", "LOR", "TCP", "OND", "MGL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PaymentType3Code")
}

// May be one of ACPD, VALD, MATD, AUTD, INVD, UMAC, STLE, STLM, SSPD, PCAN, PSTL, PFST, SMLR, RMLR, SRBL, AVLB, SRML
type PendingStatus4Code string

func (r PendingStatus4Code) Validate() error {
	for _, vv := range []string{
		"ACPD", "VALD", "MATD", "AUTD", "INVD", "UMAC", "STLE", "STLM", "SSPD", "PCAN", "PSTL", "PFST", "SMLR", "RMLR", "SRBL", "AVLB", "SRML",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PendingStatus4Code")
}

// May be one of HIGH, LOWW, NORM, URGT
type Priority5Code string

func (r Priority5Code) Validate() error {
	for _, vv := range []string{
		"HIGH", "LOWW", "NORM", "URGT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("Priority5Code")
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

// May be one of STND, PRPR
type ReportIndicator1Code string

func (r ReportIndicator1Code) Validate() error {
	for _, vv := range []string{
		"STND", "PRPR",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("ReportIndicator1Code")
}

// May be one of CANI, CANS, CSUB
type CancelledStatusReason1Code string

func (r CancelledStatusReason1Code) Validate() error {
	for _, vv := range []string{
		"CANI", "CANS", "CSUB",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CancelledStatusReason1Code")
}

// May be one of STLD, RJTD, CAND, FNLD
type FinalStatus1Code string

func (r FinalStatus1Code) Validate() error {
	for _, vv := range []string{
		"STLD", "RJTD", "CAND", "FNLD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("FinalStatus1Code")
}

// May be one of AWMO, AWSH, LAAW, DOCY, CLAT, CERT, MINO, PHSE, SBLO, DKNY, STCD, BENO, LACK, LATE, CANR, MLAT, OBJT, DOCC, BLOC, CHAS, NEWI, CLAC, PART, CMON, COLL, DEPO, FLIM, NOFX, INCA, LINK, BYIY, CAIS, LALO, MONY, NCON, YCOL, REFS, SDUT, CYCL, BATC, GUAD, PREA, GLOB, CPEC, MUNO
type PendingFailingSettlement1Code string

func (r PendingFailingSettlement1Code) Validate() error {
	for _, vv := range []string{
		"AWMO", "AWSH", "LAAW", "DOCY", "CLAT", "CERT", "MINO", "PHSE", "SBLO", "DKNY", "STCD", "BENO", "LACK", "LATE", "CANR", "MLAT", "OBJT", "DOCC", "BLOC", "CHAS", "NEWI", "CLAC", "PART", "CMON", "COLL", "DEPO", "FLIM", "NOFX", "INCA", "LINK", "BYIY", "CAIS", "LALO", "MONY", "NCON", "YCOL", "REFS", "SDUT", "CYCL", "BATC", "GUAD", "PREA", "GLOB", "CPEC", "MUNO",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PendingFailingSettlement1Code")
}

// May be one of AWMO, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, GUAD, PART, NMAS, CMON, YCOL, COLL, DEPO, FLIM, NOFX, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, PCAP
type PendingSettlement2Code string

func (r PendingSettlement2Code) Validate() error {
	for _, vv := range []string{
		"AWMO", "CAIS", "REFU", "AWSH", "PHSE", "TAMM", "DOCY", "DOCC", "BLOC", "CHAS", "NEWI", "CLAC", "MUNO", "GLOB", "PREA", "GUAD", "PART", "NMAS", "CMON", "YCOL", "COLL", "DEPO", "FLIM", "NOFX", "INCA", "LINK", "FUTU", "LACK", "LALO", "MONY", "NCON", "REFS", "SDUT", "BATC", "CYCL", "SBLO", "CPEC", "MINO", "PCAP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PendingSettlement2Code")
}

// May be one of SUBY, SUBS
type SuspendedStatusReason1Code string

func (r SuspendedStatusReason1Code) Validate() error {
	for _, vv := range []string{
		"SUBY", "SUBS",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("SuspendedStatusReason1Code")
}

// May be one of CMIS, DDAT, DELN, DEPT, DMON, DDEA, DQUA, CADE, SETR, DSEC, VASU, DTRA, RSPR, REPO, CLAT, RERT, REPA, REPP, PHYS, IIND, FRAP, PLCE, PODU, FORF, REGD, RTGS, ICAG, CPCA, CHAR, IEXE, NCRR, NMAS, SAFE, DTRD, LATE, TERM, ICUS
type UnmatchedStatusReason1Code string

func (r UnmatchedStatusReason1Code) Validate() error {
	for _, vv := range []string{
		"CMIS", "DDAT", "DELN", "DEPT", "DMON", "DDEA", "DQUA", "CADE", "SETR", "DSEC", "VASU", "DTRA", "RSPR", "REPO", "CLAT", "RERT", "REPA", "REPP", "PHYS", "IIND", "FRAP", "PLCE", "PODU", "FORF", "REGD", "RTGS", "ICAG", "CPCA", "CHAR", "IEXE", "NCRR", "NMAS", "SAFE", "DTRD", "LATE", "TERM", "ICUS",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("UnmatchedStatusReason1Code")
}

// May be one of ENAB, DISA, DELD, REQD
type LimitStatus1Code string

func (r LimitStatus1Code) Validate() error {
	for _, vv := range []string{
		"ENAB", "DISA", "DELD", "REQD",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("LimitStatus1Code")
}

// May be one of MULT, BILI, MAND, DISC, NELI, INBI, GLBL, DIDB, SPLC, SPLF, TDLC, TDLF, UCDT, ACOL, EXGT
type LimitType3Code string

func (r LimitType3Code) Validate() error {
	for _, vv := range []string{
		"MULT", "BILI", "MAND", "DISC", "NELI", "INBI", "GLBL", "DIDB", "SPLC", "SPLF", "TDLC", "TDLF", "UCDT", "ACOL", "EXGT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("LimitType3Code")
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
		"IN01", "IN02", "IN03", "IN04", "IN05", "IN06", "IN07", "IN08", "IN09", "IN10", "IN11", "IN12", "IN13", "IN14", "IN15", "IN16", "IN17", "IN18", "IN19", "MM20", "MM21", "MM22", "MM25", "MM26", "MM27", "MM28", "MM29", "MM30", "MM31", "MM32", "IN33", "MM34", "MM35", "IN36", "IN37", "IN38", "IN39", "NARR",
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

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

func (r AttendanceContext1Code) Validate() error {
	for _, vv := range []string{
		"ATTD", "SATT", "UATT",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AttendanceContext1Code")
}

// May be one of ICCD, AGNT, MERC
type AuthenticationEntity1Code string

func (r AuthenticationEntity1Code) Validate() error {
	for _, vv := range []string{
		"ICCD", "AGNT", "MERC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AuthenticationEntity1Code")
}

// May be one of UKNW, BYPS, NPIN, FPIN, CPSG, PPSG, MANU, MERC, SCRT, SNCT, SCNL
type AuthenticationMethod1Code string

func (r AuthenticationMethod1Code) Validate() error {
	for _, vv := range []string{
		"UKNW", "BYPS", "NPIN", "FPIN", "CPSG", "PPSG", "MANU", "MERC", "SCRT", "SNCT", "SCNL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("AuthenticationMethod1Code")
}

// May be one of PRST, BYPS, UNRD, NCSC
type CSCManagement1Code string

func (r CSCManagement1Code) Validate() error {
	for _, vv := range []string{
		"PRST", "BYPS", "UNRD", "NCSC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CSCManagement1Code")
}

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

func (r CardDataReading1Code) Validate() error {
	for _, vv := range []string{
		"TAGC", "PHYS", "BRCD", "MGST", "CICC", "DFLE", "CTLS", "ECTL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CardDataReading1Code")
}

// May be one of AGGR, DCCV, GRTT, INSP, LOYT, NRES, PUCO, RECP, SOAF, UNAF, VCAU
type CardPaymentServiceType2Code string

func (r CardPaymentServiceType2Code) Validate() error {
	for _, vv := range []string{
		"AGGR", "DCCV", "GRTT", "INSP", "LOYT", "NRES", "PUCO", "RECP", "SOAF", "UNAF", "VCAU",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CardPaymentServiceType2Code")
}

// May be one of MNSG, NPIN, FCPN, FEPN, FDSG, FBIO, MNVR, FBIG, APKI, PKIS, CHDT, SCEC
type CardholderVerificationCapability1Code string

func (r CardholderVerificationCapability1Code) Validate() error {
	for _, vv := range []string{
		"MNSG", "NPIN", "FCPN", "FEPN", "FDSG", "FBIO", "MNVR", "FBIG", "APKI", "PKIS", "CHDT", "SCEC",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("CardholderVerificationCapability1Code")
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

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

func (r OnLineCapability1Code) Validate() error {
	for _, vv := range []string{
		"OFLN", "ONLN", "SMON",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("OnLineCapability1Code")
}

// May be one of SOFT, EMVK, EMVO, MRIT, CHIT, SECM, PEDV
type POIComponentType1Code string

func (r POIComponentType1Code) Validate() error {
	for _, vv := range []string{
		"SOFT", "EMVK", "EMVO", "MRIT", "CHIT", "SECM", "PEDV",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("POIComponentType1Code")
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

func (r PartyType3Code) Validate() error {
	for _, vv := range []string{
		"OPOI", "MERC", "ACCP", "ITAG", "ACQR", "CISS", "DLIS",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PartyType3Code")
}

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

func (r PartyType4Code) Validate() error {
	for _, vv := range []string{
		"MERC", "ACCP", "ITAG", "ACQR", "CISS", "TAXH",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PartyType4Code")
}

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

func (r PriceValueType1Code) Validate() error {
	for _, vv := range []string{
		"DISC", "PREM", "PARV",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("PriceValueType1Code")
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

// May be one of MAIL, TLPH, ECOM, TVPY
type TransactionChannel1Code string

func (r TransactionChannel1Code) Validate() error {
	for _, vv := range []string{
		"MAIL", "TLPH", "ECOM", "TVPY",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TransactionChannel1Code")
}

// May be one of MERC, PRIV, PUBL
type TransactionEnvironment1Code string

func (r TransactionEnvironment1Code) Validate() error {
	for _, vv := range []string{
		"MERC", "PRIV", "PUBL",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("TransactionEnvironment1Code")
}

// May be one of PIEC, TONS, FOOT, GBGA, USGA, GRAM, INCH, KILO, PUND, METR, CMET, MMET, LITR, CELI, MILI, GBOU, USOU, GBQA, USQA, GBPI, USPI, MILE, KMET, YARD, SQKI, HECT, ARES, SMET, SCMT, SMIL, SQMI, SQYA, SQFO, SQIN, ACRE
type UnitOfMeasure1Code string

func (r UnitOfMeasure1Code) Validate() error {
	for _, vv := range []string{
		"PIEC", "TONS", "FOOT", "GBGA", "USGA", "GRAM", "INCH", "KILO", "PUND", "METR", "CMET", "MMET", "LITR", "CELI", "MILI", "GBOU", "USOU", "GBQA", "USQA", "GBPI", "USPI", "MILE", "KMET", "YARD", "SQKI", "HECT", "ARES", "SMET", "SCMT", "SMIL", "SQMI", "SQYA", "SQFO", "SQIN", "ACRE",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("UnitOfMeasure1Code")
}

// May be one of MDSP, CDSP
type UserInterface2Code string

func (r UserInterface2Code) Validate() error {
	for _, vv := range []string{
		"MDSP", "CDSP",
	} {
		if reflect.DeepEqual(string(r), vv) {
			return nil
		}
	}
	return utils.NewErrValueInvalid("UserInterface2Code")
}

// Must match the pattern [0-9]
type Exact1NumericText string

func (r Exact1NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Exact1NumericText")
	}
	return nil
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

func (r Exact3NumericText) Validate() error {
	reg := regexp.MustCompile(`[0-9]{3}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("Exact3NumericText")
	}
	return nil
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

func (r ISINOct2015Identifier) Validate() error {
	reg := regexp.MustCompile(`[A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ISINOct2015Identifier")
	}
	return nil
}

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

func (r ISO2ALanguageCode) Validate() error {
	reg := regexp.MustCompile(`[a-z]{2,2}`)
	if !reg.MatchString(string(r)) {
		return utils.NewErrValueInvalid("ISO2ALanguageCode")
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
