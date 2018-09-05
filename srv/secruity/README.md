# Secruity Service

This is the Secruity service

Generated with

```
micro new github.com/iron-kit/unite-services/secruity --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.secruity
- Type: srv
- Alias: secruity

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
./secruity-srv
```

Build a docker image
```
make docker
```