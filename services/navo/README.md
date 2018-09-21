# Navo Service

This is the Navo service

Generated with

```
micro new iunite.club/services/navo --namespace=iunite.club --alias=navo --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.api.navo
- Type: api
- Alias: navo

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
./navo-api
```

Build a docker image
```
make docker
```