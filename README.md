[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://github.com/moov-io/iso20022/tree/master/docs">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/iso20022/api/">API Endpoints</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/iso20022?status.svg)](https://godoc.org/github.com/moov-io/iso20022)
[![Build Status](https://github.com/moov-io/iso20022/workflows/Go/badge.svg)](https://github.com/moov-io/iso20022/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/iso20022/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/iso20022)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/iso20022)](https://goreportcard.com/report/github.com/moov-io/iso20022)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/iso20022?label=project%20size)](https://github.com/moov-io/iso20022)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/iso20022/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/iso20022)](https://hub.docker.com/r/moov/iso20022)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/iso20022)](https://github.com/moov-io/iso20022)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/iso20022

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

ISO 20022 is a standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The standard covers financial information transferred between financial institutions that includes payment transactions, securities trading and settlement information, credit and debit card transactions, and other financial information.

ISO20022 implements a message reader and writer written in Go, decorated with an HTTP API for creating, parsing, and validating ISO 20022 messages used in financial networks. ISO20022 supports XML and JSON format for messages. The types of messages supported in this package are for business domain catalogue payments.

## Project Status

ISO 20022 is a large and evolving set of specifications. We are deprecating this repository in favor of targeted repositories such as [rtp20022](https://github.com/moov-io/rtp20022).

Please star the project if you are interested in its progress. Feedback on this early version of ISO 20022 is appreciated and vital to its success. Please let us know if you encounter any bugs/unclear documentation or have feature suggestions by opening up an issue. Thanks!

## Usage

### Go Library

This project offers a Go library which can read and write ISO 20022 messages. We write unit tests and fuzz the code to help ensure it is production ready for everyone. ISO20022 uses [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies and suggests Go 1.14 or greater.

To clone our code and verify our tests on your system, run:

```
$ git clone git@github.com:moov-io/iso20022.git
$ cd iso20022

$ go test ./...
ok      github.com/moov-io/iso20022      0.015s
ok      github.com/moov-io/iso20022/cmd/iso20022  21.908s
...
```

### Formats and Configuration

ISO20022 supports two message types: JSON and XML. The general ISO 20022 specification defines a message structure, but doesn't define JSON and XML format. Our ISO20022 package also includes a specification file (configuration file) that is used to define message structure.

JSON format example:
```
{
	"Xmlns": "urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03",
    "AcctOpngReq": {}, // real document body
}
```

XML format example:
```
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:acmt.007.001.03" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<AcctOpngReq></AcctOpngReq> // real document body
</Document>
```

### Docker (under construction)

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

### Command Line (under construction)

ISO20022 has a command line interface to manage messages and launch a web service.

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
`convert` | The convert command allows users to convert between message formats. The output will create a new message.
`print` | The print command allows users to print a message in a specified file format (JSON, XML).
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

- The `output` parameter represents the full path name for the new iso20022 file.
- The `format` parameter determines the output file format and supports “json”, “xml”, and "iso20022".
- The `input` parameter is the source iso20022 file to be converted, and can be “json”, “xml”, or "iso20022".
- The `spec` parameter is the specification file.

Example:
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

Example:
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

Example:
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

Example:
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

## Getting Help

 channel | info
 ------- | -------
[Project Documentation](https://github.com/moov-io/iso20022/tree/master/docs) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov) | You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/iso20022/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#iso20022`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! [Check out our issues](https://github.com/moov-io/iso20022/issues) for something to help out with.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and uses Go 1.14 or higher. See [Golang's install instructions](https://golang.org/doc/install) for help setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/iso20022/releases/latest) as well. We highly recommend you use a tagged release for production.

## Related Projects
As part of Moov's initiative to offer open source fintech infrastructure, we have a large collection of active projects you may find useful:

- [Moov ACH](https://github.com/moov-io/ach) provides ACH file generation and parsing, supporting all Standard Entry Codes for the primary method of money movement throughout the United States.

- [Moov Watchman](https://github.com/moov-io/watchman) offers search functions over numerous trade sanction lists from the United States and European Union.

- [Moov Fed](https://github.com/moov-io/fed) implements utility services for searching the United States Federal Reserve System such as ABA routing numbers, financial institution name lookup, and FedACH and Fedwire routing information.

- [Moov Wire](https://github.com/moov-io/wire) implements an interface to write files for the Fedwire Funds Service, a real-time gross settlement funds transfer system operated by the United States Federal Reserve Banks.

- [Moov Image Cash Letter](https://github.com/moov-io/imagecashletter) implements Image Cash Letter (ICL) files used for Check21, X.9 or check truncation files for exchange and remote deposit in the U.S.

- [Moov Metro 2](https://github.com/moov-io/metro2) provides a way to easily read, create, and validate Metro 2 format, which is used for consumer credit history reporting by the United States credit bureaus.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
