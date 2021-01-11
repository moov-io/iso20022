// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

// May be one of PDNG, STLD
type BalanceStatus1Code string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalSystemBalanceType1Code string

// Must be at least 1 items long
type ExternalSystemErrorHandling1Code string

// Must be at least 1 items long
type ExternalSystemEventType1Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, OVNG
type Frequency2Code string

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of RJCT, CVHD, RSVT, BLCK, EARM, EFAC, DLVR, COLD, CSDB
type ProcessingType1Code string

// May be one of USTO, PSTO
type StandingOrderType1Code string

// May be one of OPNG, INTM, CLSG, BOOK, CRRT, PDNG, LRLD, AVLB, LTSF, CRDT, EAST, PYMT, BLCK, XPCD, DLOD, XCRD, XDBT, ADJT, PRAV, DBIT, THRE, NOTE, FSET, BLOC, OTHB, CUST, FORC, COLC, FUND, PIPO, XCHG, CCPS, TOHB, COHB, DOHB, TPBL, CPBL, DPBL, FUTB, REJB, FCOL, FCOU, SCOL, SCOU, CUSA, XCHC, XCHN, DSET, LACK, NSET, OTCC, OTCG, OTCN, SAPD, SAPC, REPD, REPC, BSCD, BSCC, SAPP, IRLT, IRDR, DWRD, ADWR, AIDR
type SystemBalanceType2Code string

// May be one of PDNG, FINL
type CashPaymentStatus2Code string

// May be one of BOOK, PDNG, FUTR
type EntryStatus1Code string

// Must match the pattern [BEOVW]{1,1}[0-9]{2,2}|DUM
type EntryTypeIdentifier string

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

// May be one of STLD, RJTD, CAND, FNLD
type FinalStatusCode string

// May be one of PBEN, TTIL, TFRO
type Instruction1Code string

// May be one of BDT, BCT, CDT, CCT, CHK, BKT, DCP, CCP, RTI, CAN
type PaymentInstrument1Code string

// May be one of CBS, BCK, BAL, CLS, CTR, CBH, CBP, DPG, DPN, EXP, TCH, LMT, LIQ, DPP, DPH, DPS, STF, TRP, TCS, LOA, LOR, TCP, OND, MGL
type PaymentType3Code string

// May be one of ACPD, VALD, MATD, AUTD, INVD, UMAC, STLE, STLM, SSPD, PCAN, PSTL, PFST, SMLR, RMLR, SRBL, AVLB, SRML
type PendingStatus4Code string

// May be one of HIGH, LOWW, NORM, URGT
type Priority5Code string

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

// May be one of STND, PRPR
type ReportIndicator1Code string

// May be one of CANI, CANS, CSUB
type CancelledStatusReason1Code string

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

// May be one of STLD, RJTD, CAND, FNLD
type FinalStatus1Code string

// May be one of AWMO, AWSH, LAAW, DOCY, CLAT, CERT, MINO, PHSE, SBLO, DKNY, STCD, BENO, LACK, LATE, CANR, MLAT, OBJT, DOCC, BLOC, CHAS, NEWI, CLAC, PART, CMON, COLL, DEPO, FLIM, NOFX, INCA, LINK, BYIY, CAIS, LALO, MONY, NCON, YCOL, REFS, SDUT, CYCL, BATC, GUAD, PREA, GLOB, CPEC, MUNO
type PendingFailingSettlement1Code string

// May be one of AWMO, CAIS, REFU, AWSH, PHSE, TAMM, DOCY, DOCC, BLOC, CHAS, NEWI, CLAC, MUNO, GLOB, PREA, GUAD, PART, NMAS, CMON, YCOL, COLL, DEPO, FLIM, NOFX, INCA, LINK, FUTU, LACK, LALO, MONY, NCON, REFS, SDUT, BATC, CYCL, SBLO, CPEC, MINO, PCAP
type PendingSettlement2Code string

// May be one of SUBY, SUBS
type SuspendedStatusReason1Code string

// May be one of CMIS, DDAT, DELN, DEPT, DMON, DDEA, DQUA, CADE, SETR, DSEC, VASU, DTRA, RSPR, REPO, CLAT, RERT, REPA, REPP, PHYS, IIND, FRAP, PLCE, PODU, FORF, REGD, RTGS, ICAG, CPCA, CHAR, IEXE, NCRR, NMAS, SAFE, DTRD, LATE, TERM, ICUS
type UnmatchedStatusReason1Code string

// Must be at least 1 items long
type ExternalCancellationReason1Code string

// May be one of ENAB, DISA, DELD, REQD
type LimitStatus1Code string

// May be one of MULT, BILI, MAND, DISC, NELI, INBI, GLBL, DIDB, SPLC, SPLF, TDLC, TDLF, UCDT, ACOL, EXGT
type LimitType3Code string

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must match the pattern [0-9]{2}
type Exact2NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

// Must be at least 1 items long
type ExternalGarnishmentType1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

// Must be at least 1 items long
type ExternalPurpose1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

// May be one of HIGH, NORM
type Priority2Code string

// May be one of FRST, RCUR, FNAL, OOFF, RPRE
type SequenceType3Code string

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// May be one of IN01, IN02, IN03, IN04, IN05, IN06, IN07, IN08, IN09, IN10, IN11, IN12, IN13, IN14, IN15, IN16, IN17, IN18, IN19, MM20, MM21, MM22, MM25, MM26, MM27, MM28, MM29, MM30, MM31, MM32, IN33, MM34, MM35, IN36, IN37, IN38, IN39, NARR
type UnableToApplyIncorrectInformation4Code string

// May be one of MS01, MS02, MS03, MS04, MS05, MS06, MS07, MS08, MS09, MS10, MS11, MS12, MS13, MS14, MS15, MS16, MS17, NARR
type UnableToApplyMissingInformation3Code string

// Must be at least 1 items long
type ExternalAgentInstruction1Code string

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of ICCD, AGNT, MERC
type AuthenticationEntity1Code string

// May be one of UKNW, BYPS, NPIN, FPIN, CPSG, PPSG, MANU, MERC, SCRT, SNCT, SCNL
type AuthenticationMethod1Code string

// May be one of PRST, BYPS, UNRD, NCSC
type CSCManagement1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL
type CardDataReading1Code string

// May be one of AGGR, DCCV, GRTT, INSP, LOYT, NRES, PUCO, RECP, SOAF, UNAF, VCAU
type CardPaymentServiceType2Code string

// May be one of MNSG, NPIN, FCPN, FEPN, FDSG, FBIO, MNVR, FBIG, APKI, PKIS, CHDT, SCEC
type CardholderVerificationCapability1Code string

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

// Must match the pattern [0-9]
type Exact1NumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must be at least 1 items long
type ExternalBalanceSubType1Code string

// Must be at least 1 items long
type ExternalBalanceType1Code string

// Must be at least 1 items long
type ExternalBankTransactionDomain1Code string

// Must be at least 1 items long
type ExternalBankTransactionFamily1Code string

// Must be at least 1 items long
type ExternalBankTransactionSubFamily1Code string

// Must be at least 1 items long
type ExternalCardTransactionCategory1Code string

// Must be at least 1 items long
type ExternalChargeType1Code string

// Must be at least 1 items long
type ExternalCreditLineType1Code string

// Must be at least 1 items long
type ExternalEntryStatus1Code string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

// Must be at least 1 items long
type ExternalRePresentmentReason1Code string

// Must be at least 1 items long
type ExternalReportingSource1Code string

// Must be at least 1 items long
type ExternalReturnReason1Code string

// Must be at least 1 items long
type ExternalTechnicalInputChannel1Code string

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

// Must match the pattern [a-z]{2,2}
type ISO2ALanguageCode string

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

// May be one of SOFT, EMVK, EMVO, MRIT, CHIT, SECM, PEDV
type POIComponentType1Code string

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of DISC, PREM, PARV
type PriceValueType1Code string

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

// May be one of MAIL, TLPH, ECOM, TVPY
type TransactionChannel1Code string

// May be one of MERC, PRIV, PUBL
type TransactionEnvironment1Code string

// May be one of PIEC, TONS, FOOT, GBGA, USGA, GRAM, INCH, KILO, PUND, METR, CMET, MMET, LITR, CELI, MILI, GBOU, USOU, GBQA, USQA, GBPI, USPI, MILE, KMET, YARD, SQKI, HECT, ARES, SMET, SCMT, SMIL, SQMI, SQYA, SQFO, SQIN, ACRE
type UnitOfMeasure1Code string

// May be one of MDSP, CDSP
type UserInterface2Code string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string
