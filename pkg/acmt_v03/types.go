// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package acmt_v03

// May be one of EMAL, FAXI, FILE, ONLI, POST
type CommunicationMethod2Code string

// May be one of EMAL, FAXI, POST, PHON, FILE, ONLI
type CommunicationMethod3Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalCommunicationFormat1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// May be one of YEAR, DAIL, MNTH, QURT, MIAN, TEND, MOVE, WEEK, INDA
type Frequency7Code string

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// Must match the pattern UNLIMITED
type Unlimited9Text string

// May be one of NOCH, MODI, DELE, ADDD
type Modification1Code string

// May be one of DAYH, EARL
type BalanceTransferWindow1Code string

// May be one of FWNG, PREC
type BusinessDayConvention1Code string

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

// May be one of MLDB, MLCD, MLFA, CRDB, CRCD, CRFA, PUDB, PUCD, PUFA, RGDB, RGCD, RGFA
type ChequeDelivery1Code string

// May be one of CCHQ, CCCH, BCHQ, DRFT, ELDR
type ChequeType2Code string

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

// May be one of NEVR, YEAR, RATE, MIAN, QURT
type Frequency10Code string

// May be one of FEMA, MALE
type Gender1Code string

// May be one of CIOC, CHAR, CICC, GENP, IAPS, LLPP, PCLG, LIMP, PCLS, PCLC, SOLE, UNLC, UNLT
type OrganisationLegalStatus1Code string

// May be one of AREG, CPFA, DRLC, EMID, IDCD, NRIN, OTHR, PASS, POCD, SOCS, SRSA, GUNL
type PersonIdentificationType5Code string

// May be one of HIGH, NORM
type Priority2Code string

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

// May be one of RESI, PRES, NRES
type ResidentialStatus1Code string

// May be one of FULL, PART
type SwitchType1Code string

// May be one of ALPR, ALIT, GRSS
type TaxRateMarker1Code string

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// May be one of ACPT, BTRQ, BTRS, COMP, REDT, REDE, REJT, REQU, TMTN
type SwitchStatus1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalCreditorAgentInstruction1Code string

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

// Must be at least 1 items long
type ExternalGarnishmentType1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalPurpose1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

// May be one of CHK, TRF, TRA
type PaymentMethod3Code string

// May be one of OPEN, MNTN, CLSG, VIEW
type UseCases1Code string
