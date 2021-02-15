# moov-io/iso20022

[![GoDoc](https://godoc.org/github.com/moov-io/iso20022?status.svg)](https://godoc.org/github.com/moov-io/iso20022)
[![Build Status](https://github.com/moov-io/iso20022/workflows/Go/badge.svg)](https://github.com/moov-io/iso20022/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/iso20022/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/iso20022)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/iso20022)](https://goreportcard.com/report/github.com/moov-io/iso20022)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/iso20022/master/LICENSE)

ISO 20022 is an ISO standard for electronic data interchange between financial institutions. 
It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The standard covers financial information transferred between financial institutions that includes payment transactions, securities trading and settlement information, credit and debit card transactions and other financial information.

The repository contains a huge amount of financial services metadata that has been shared and standardized across the industry.
The metadata is stored in UML models with a special ISO 20022 UML Profile. 
Underlying all of this is the ISO 20022 metamodel - a model of the models. 
The UML profile is the metamodel transformed into UML. 
The metadata is transformed into the syntax of messages used in financial networks.
 
The first syntax supported for messages was XML Schema. 

Package `github.com/moov-io/iso20022` implements a message reader and writer written in Go decorated with a HTTP API for creating, parsing, and validating messages used in financial networks.
Package ISO20022 supported xml and json format for message
The meessages supported in this package are messages of payments in business domain catalogue.

Docs: [API Endpoints](https://moov-io.github.io/iso20022/api/)


## Getting Started

### Docker

We publish a [public Docker image `moov/iso20022`](https://hub.docker.com/r/moov/iso20022/tags) on Docker Hub with tagged release of the package. No configuration is required to serve on `:8080`.


Start the Docker image:
```
docker run -p 8080:8080 moov/iso20022:latest
```

Upload a file and validate it
```
curl -XPOST --form "input=@./test/testdata/valid_acmt_v03.xml" http://localhost:8080/validator
```
```
{"status":"valid file"}
```

Convert a message between formats
```
curl -XPOST --form "file=@./test/testdata/valid_acmt_v03.xml" --form "format=json" http://localhost:8080/convert
```
```
{
	"Xmlns": "urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03",
	"AcctOpngReq": {
		"Refs": {
			"MsgId": {
				"Id": "Id",
				"CreDtTm": "2014-11-12T11:45:26.371"
			},
			"PrcId": {
				"Id": "Id",
				"CreDtTm": "2014-11-12T11:45:26.371"
			}
		},
		"Acct": {
			"Ccy": "ABC"
		},
		"AcctSvcrId": {
			"FinInstnId": {}
		},
		"Org": {
			"FullLglNm": "FullLglNm",
			"CtryOfOpr": "AA",
			"LglAdr": {},
			"OrgId": {}
		}
	}
}
```
```
curl -XPOST --form "file=@./test/testdata/valid_acmt_v03.json" http://localhost:8080/convert
```
```
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<AcctOpngReq>
		<Refs>
			<MsgId>
				<Id>Id</Id>
				<CreDtTm>2014-11-12T11:45:26.371</CreDtTm>
			</MsgId>
			<PrcId>
				<Id>Id</Id>
				<CreDtTm>2014-11-12T11:45:26.371</CreDtTm>
			</PrcId>
		</Refs>
		<Acct>
			<Ccy>ABC</Ccy>
		</Acct>
		<AcctSvcrId>
			<FinInstnId></FinInstnId>
		</AcctSvcrId>
		<Org>
			<FullLglNm>FullLglNm</FullLglNm>
			<CtryOfOpr>AA</CtryOfOpr>
			<LglAdr></LglAdr>
			<OrgId></OrgId>
		</Org>
	</AcctOpngReq>
</Document>
```

### Go Library

There is a Go library which can read and write iso20022 message. We write unit tests and fuzz the code to help ensure our code is production ready for everyone. Iso20022 uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies and suggests Go 1.14 or greater.

To clone our code and verify our tests on your system run:

```
$ git clone git@github.com:moov-io/iso20022.git
$ cd iso20022

$ go test ./...
ok      github.com/moov-io/iso20022      0.015s
ok      github.com/moov-io/iso20022/cmd/iso20022  21.908s
...
```

## Formats and configuration file
### message formats
Iso20022 have supported 2 message types: json, xml.
Iso20022 specification defines a message format, but don't define json and xml format.
Iso20022 package have defined json and xml format of message, specification file (configuration file) that use to define message structure.

Json format:
```
{
	"Xmlns": "urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03",
    "AcctOpngReq": {}, // real document body
}
```

XML format:
```
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<AcctOpngReq></AcctOpngReq> // real document body
</Document>
```

## Commands

iso20022 has command line interface to manage iso20022 messages and to lunch web service.

```
iso20022 --help

Usage:
   [command]

Available Commands:
  convert     Convert iso20022 document file format
  help        Help about any command
  print       Print iso20022 message
  validator   Validate iso20022 message
  web         Launches web server

Flags:
  -h, --help           help for this command
      --input string   iso20022 document (valid types are xml, json. default is $PWD/iso20022_document.xml)

Use " [command] --help" for more information about a command.
```

Each interaction that the library supports is exposed in a command-line option:

 Command | Info
 ------- | -------
`convert` | The convert command allows users to convert from a message to another message format. Result will create a message.
`print` | The print command allows users to print a message with special file format (json, xml).
`validator` | The validator command allows users to validate a message.
`web` | The web command will launch a web server with endpoints to manage messages.

### message convert

```
Usage:
   convert [output] [flags]

Flags:
      --format string   format of document file (default "xml")
  -h, --help            help for convert

Global Flags:
      --input string   iso20022 document (valid types are xml, json. default is $PWD/iso20022_document.xml)
```

The output parameter is the full path name to convert new iso20022 message.
The format parameter is supported 3 types that are "json", "xml" and  "iso20022".
The input parameter is source iso20022 message, supported "json", "xml" and  "iso20022".
The spec parameter is specification file.

example:
```
iso20022 converted --input test/testdata/valid_acmt_v03.json
```

### message print

```
iso20022 print --help

Usage:
   print [flags]

Flags:
      --format string   print format (default "xml")
  -h, --help            help for print

Global Flags:
      --input string   iso20022 document (valid types are xml, json. default is $PWD/iso20022_document.xml)
```

example:
```
iso20022 print --input test/testdata/valid_acmt_v03.json
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<AcctOpngReq>
		<Refs>
			<MsgId>
				<Id>Id</Id>
				<CreDtTm>2014-11-12T11:45:26.371</CreDtTm>
			</MsgId>
			<PrcId>
				<Id>Id</Id>
				<CreDtTm>2014-11-12T11:45:26.371</CreDtTm>
			</PrcId>
		</Refs>
		<Acct>
			<Ccy>ABC</Ccy>
		</Acct>
		<AcctSvcrId>
			<FinInstnId></FinInstnId>
		</AcctSvcrId>
		<Org>
			<FullLglNm>FullLglNm</FullLglNm>
			<CtryOfOpr>AA</CtryOfOpr>
			<LglAdr></LglAdr>
			<OrgId></OrgId>
		</Org>
	</AcctOpngReq>
</Document>

```

### message validate

```
iso20022 validator --help

Usage:
   validator [flags]

Flags:
  -h, --help   help for validator

Global Flags:
      --input string   iso20022 document (valid types are xml, json. default is $PWD/iso20022_document.xml)
```

The input parameter is source iso20022 message, supported "json", "xml" and  "iso20022".

example:
```
iso20022 validator --input testdata/valid_acmt_v03.json
```

### web server

```
iso20022 web --help

Usage:
   web [flags]

Flags:
  -h, --help   help for web
  -t, --test   test server

Global Flags:
      --input string   iso20022 document (valid types are xml, json. default is $PWD/iso20022_document.xml)
```

The port parameter is port number of web service.

example:
```
iso20022 web
```

Web server have some endpoints to manage iso20022 messages

Method | Endpoint | Content-Type | Info
 ------- | ------- | ------- | -------
 `POST` | `/convert` | multipart/form-data | convert iso20022 messages. will download new file.
 `GET` | `/health` | text/plain | check web server.
 `POST` | `/print` | multipart/form-data | print iso20022 messages.
 `POST` | `/validator` | multipart/form-data | validate iso20022 messages.

web page example to use iso20022 web server:

```
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Test ISO20022 APIs</title>
</head>
<body>
<h1>Upload single file with fields</h1>

<form action="http://localhost:8080/convert" method="post" enctype="multipart/form-data">
    Format: <input type="text" name="format"><br>
    Input File: <input type="file" name="input"><br><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>
```

## Docker

You can run the [moov/iso20022 Docker image](https://hub.docker.com/r/moov/iso20022) which defaults to starting the HTTP server.

```
docker run -p 8080:8080 moov/iso20022:latest
```

## Getting Help

 channel | info
 ------- | -------
  Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/iso20022/issues) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#iso20022`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! [Checkout our issues](https://github.com/moov-io/iso20022/issues) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/iso20022/releases/latest) as well. We highly recommend you use a tagged release for production.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.

