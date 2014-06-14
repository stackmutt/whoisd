
# Whois Daemon

*A quick and easy way to setup your own WHOIS server with Elasticsearch/MySQL storage*

[![Build Status](https://travis-ci.org/takama/whoisd.png?branch=master)](https://travis-ci.org/takama/whoisd)

**Whois Daemon** is a light server which provide fast way to present whois information. The daemon based on Elasticsearch storage (Mysql storage still in development). 

### Install

This package is "go-gettable", just do:

```sh
go get github.com/takama/whoisd
```

_NOTE_: you need go 1.2+ and need set PATH="$GOPATH/bin". Please check your installation with

```sh
go version
go env
```

### Running

Start of whoisd:

```sh
sudo whoisd
```

This will bring up whoisd listening on port 43 for client communication.

### Usage

```
whoisd - Whois Daemon

Usage:
  whoisd -option
  whoisd -h | --help
  whoisd -v | --version

Options:
  -h --help         Show this screen
  -v --version      Show version
  -config=<path>    Path to configuration file
  -mapper=<path>    Path to mapper file
  -host=<host/IP>   Host name or IP address
  -port=<port>      Port number
  -work=<number>    Number of active workers (default 1000)
  -conn=<number>    Number of active connections (default 1000)
  -storage=<type>   Type of storage (Elasticsearch or Mysql)
  -shost=<host/IP>  Storage host name or IP address
  -sport=<port>     Storage port number
  -base=<name>      Storage index or database name
  -table=<name>     Storage type or table name
```

### Config

Default configuration file placed in /etc/whoisd/whoisd.conf. Of course possible to load a configuration from any other place through parameters on the command line. If the configuration file is absent, then used the default configuration: 

```json
{
  "host": "localhost",
  "port": 43,
  "workers": 1000,
  "connections": 1000,
  "storage": {
    "storageType": "Elasticsearch",
    "host": "localhost",
    "port": 9200,
    "indexBase": "whois",
    "typeTable": "domain"
  }
}
```

### Mapping

All required fields for the whoisd must be defined in the mapping file, the example mapping file is described below:

```json
{
  "Fields" : {
    "01": {
      "key": "Domain Name: ",
      "name": ["name"],
      "related": "name"
    },
    "02": {
      "key": "Registry Domain ID: ",
      "name": ["domainId"],
      "related": "name"
    },
    "03": {
      "key": "Registrar WHOIS Server: ",
      "value": ["whois.yourwhois.eu"]
    }

}
```

- "Fields" - the list of fields from "01" to last number "nn" in ascending order
- "key" - the prompt for the field (the preinstalled configuration file has keys according to ICANN requirements)
- "value" - if the field has prearranged value (not use any field from the database)
- "name" - field name in the database, if the field is not prearranged ("value" is not defined)
- "related" - field name in the database through which the request for 

```json
{

    "06": {
      "key": "Creation Date: ",
      "name": ["creationDate"],
      "format": "{date}",
      "related": "name"
    }

}
```

- "format" - special instructions to indicate how to display field, examples of the use of this will be shown below
- "{date}" - used in the format to indicate that the field is a date and need special formatting of the date RFС3339

```json
{

    "12": {
      "key": "Domain Status: ",
      "name": ["domainStatus"],
      "multiple": true,
      "related": "name"
    },

    "52": {
      "key": "Name Server: ",
      "name": ["name"],
      "multiple": true,
      "related": "nsgroupId",
      "relatedBy": "nsgroupId",
      "relatedTo": "nameserver"
    }

}
```

- "multiple" - if this option is set to 'true', then for each value will be repeated prompt in the whois output like that:
```
Name Server: ns1.example.com
Name Server: ns2.example.com
Name Server: ns3.example.com
```
- "relatedBy" - field name in the database through which the related request for 
- "relatedTo" - table/type name in the database through which made relation

```json
{

    "14": {
      "key": "Registrant Name: ",
      "name": ["name.fullName"],
      "related": "ownerHandle",
      "relatedBy": "handle",
      "relatedTo": "customer"
    },

    "40": {
      "key": "Tech Name: ",
      "name": ["name.firstName", "name.lastName"],
      "related": "techHandle",
      "relatedBy": "handle",
      "relatedTo": "customer"
    }

}
```

- "name": ["name.fullName"] - use dot notation for embedded fields (MySQL storage not allowed)
- "name": ["name.firstName", "name.lastName"] - all these fields will be joined by default

```json
{

    "21": {
      "key": "Registrant Phone: ",
      "name": ["phone.countryCode", "phone.areaCode", "phone.subscriberNumber"],
      "format": "{string}.{string}{string}",
      "related": "ownerHandle",
      "relatedBy": "handle",
      "relatedTo": "customer"
    }

}
```

- "format": "{string}.{string}{string}" - indicate that the fields ["phone.countryCode", "phone.areaCode", "phone.subscriberNumber"] need special formatting by described format (not simple joined by default)
- {string} - represent one string field in the format option

```json
{

    "55": {
      "key": "",
      "value": [""],
      "format": ">>> Last update of WHOIS database: {date} <<<"
    }
}
```

- example of the formating where used {date} and name field has not present, the result is {date} will be replaced by current date in RFC3339 format


Copyright (c) 2014 Igor Dolzhikov

[MIT License](https://github.com/takama/whoisd/blob/master/LICENSE)
