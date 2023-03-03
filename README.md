# vault-pgpass

[![GoDoc](https://godoc.org/github.com/frederikhs/vault-pgpass?status.svg)](https://godoc.org/github.com/frederikhs/vault-pgpass)
[![Quality](https://goreportcard.com/badge/github.com/frederikhs/vault-pgpass)](https://goreportcard.com/report/github.com/frederikhs/vault-pgpass)
[![Test](https://github.com/frederikhs/vault-pgpass/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/frederikhs/vault-pgpass/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/frederikhs/vault-pgpass.svg)](https://github.com/frederikhs/vault-pgpass/releases/latest)
[![License](https://img.shields.io/github/license/frederikhs/vault-pgpass)](LICENSE)

*generate a [pgpass](https://www.postgresql.org/docs/current/libpq-pgpass.html) file from configuration using secrets stored in vault*

Uses [KVv2](https://developer.hashicorp.com/vault/docs/secrets/kv#kv-version-2)

## Usage
```text
Usage of vault-pgpass:
  -f string
    	specify location for your configuration file
  -g	generate an example configuration file
  -o string
    	output file or omit for stdout
  -t string
    	vault token to use
```

## Example
Generate a new example configuration

```shell
vault-pgpass -g -o example.yml
```

Write a pgpass file to `.pgpass` using `configuration.yml` and a token

```shell
vault-pgpass -f configuration.yml -t <TOKEN> -o .pgpass
```

Omitting `-o` writes to stdout
