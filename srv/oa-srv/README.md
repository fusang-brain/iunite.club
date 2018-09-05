# Oa-Srv Service

This is the Oa-Srv service

Generated with

```
micro new github.com/iron-kit/unite-services/oa-srv --namespace=kit.iron.srv --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: kit.iron.srv.srv.oa-srv
- Type: srv
- Alias: oa-srv

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./oa-srv-srv
```

Build a docker image
```
make docker
```