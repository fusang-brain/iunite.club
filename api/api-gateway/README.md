# Api-Gateway Service

This is the Api-Gateway service

Generated with

```
micro new github.com/iron-kit/unite-services/api-gateway --namespace=kit.iron --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: kit.iron.api.api-gateway
- Type: api
- Alias: api-gateway

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
./api-gateway-api
```

Build a docker image
```
make docker
```