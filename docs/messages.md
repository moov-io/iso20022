## ISO 20022 Message Name Format

An ISO 20022 message name consists of a four-character business area code, three-digit message number, three-digit variant number, and two-digit version number in that order. These elements are each separated by a period and left padded with zeros.  For example, in the FeeCollectionInitiation message `cafc.002.001.01`, the business area is Fee Collection, message number is 2, variant is 1, and version is 1.

Variants are usually simplified messages that remove portions of their source global definition that are rarely used. They reduce complexity during implementation and still cover most use cases. Version numbers increase whenever message contents are updated by ISO.


## ISO 20022 Business Areas

| Code | Business Area | Usage |
|-|-|-|
| acmt* | Account Management | Management of account-related activities, such as the opening and maintenance of an account. |
| admi | Administration | Generic messages like system event notifications, generic rejections, etc. |
| auth* | Authorities | The provision of miscellaneous financial information to authorities, such as regulators, police, customs, tax authorities, enforcement authorities, ministries, etc. |
| caaa | Acceptor to Acquirer Card Transactions | Any card payment-related transactions and services between a card acceptor and card transaction acquirer. It includes the authorization, cancellation, and capture of card transactions. |
| caad | Card Administration | Batch management, batch transfers, and reconciliation. |
| caam | ATM Management | Card-related terminal management services between an ATM and Acquirer. |
| cafc | Fee Collection | -- |
| cafm | File Management | -- |
| cafr | Fraud Reporting and Disposition | -- |
| cain | Acquirer to Issuer Card Transactions | Any card payment-related transactions and services between a card transaction acquirer and a card issuer. It includes the authorization, reversal, and financial presentment of card transactions. |
| camt* | Cash Management | The reporting and advising of the cash side of any financial transactions, including cash movements, transactions, and balances, plus any exceptions and investigations related to cash transactions. |
| canm | Network Management | Includes key exchanges. |
| casp | Sale to POI Card Transactions | Any card-related transactions and services between a sale system and Point of Interaction (POI) system. |
| casr | Settlement Reporting | -- |
| catm | Terminal Management | Card-related terminal management services between a Terminal Management System (TMS) and a Point of Interaction (POI). |
| catp | ATM Card Transactions | Any card-related ATM transactions and services between ATM equipment and an ATM acquirer. These services include cash withdrawals, kiosk functions, and card account management transactions. |
| colr | Collateral Management | Includes proposals, disputes, reports, etc. |
| fxtr | Foreign Exchange Traded | Trade and post-trade processes for foreign exchange contracts, including orders to buy or sell, execution, affirmation, confirmation, allocation, and notification. |
| head | Business Application Header | Can be combined with any other ISO20022 message definition to form a business message. |
| pacs* | Payments Clearing and Settlement | The clearing and settlement processes for payment transactions between financial institutions. |
| pain* | Payments Initiation | The initiation of a payment from the ordering customer to a financial institution that services a cash account and reports its status. |
| reda* | Reference Data | The communication of reference data related to financial instruments, parties, accounts, prices, and other business information required to support financial activities. |
| remt* | Payments Remittance Advice | Communication between creditors and debtors regarding remittance details associated with payments. |
| secl | Securities Clearing | The clearing process for securities, including management of post-trading, pre-settlement credit exposure, netting, margining, borrowing, and conformance with market settlement rules. |
| seev | Securities Events | Asset servicing, including proxy voting and corporate actions. |
| semt | Securities Management | Post-settlement processes for securities (including reporting on securities movements, trades, and balances), the processes required to protect beneficial owner's rights throughout settlement, plus any exceptions and investigations related to securities transactions. |
| sese | Securities Settlement | The settlement process for securities and reporting its status and confirmation. |
| setr | Securities Trade | Trade and post-trade processes for securities, including orders to buy or sell, trade execution, affirmation, confirmation, allocation, and notification. |
| tsin | Trade Services Initiation | The request for a trade service, including any related application, instruction, request, acknowledgement, or advice. |
| tsmt | Trade Services Management | Ancillary commercial trade services functions, including checking, matching, and reporting, plus any exceptions and investigations related to trade services transactions. |
| tsrv | Trade Services | The issuance of a trade services instrument including any related reimbursement, acceptance, authorisation, claims, enquiries, invoicing, or financing. |

\* Moov ISO20022 currently supports these business areas.

Sources - iso20022.org, iotafinance.com, swift.com

