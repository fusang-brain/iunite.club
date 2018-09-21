# Core Service

This is the Core service

Generated with

```
micro new iunite.club/services/core --namespace=iunite.club --alias=core --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.srv.core
- Type: srv
- Alias: core

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
./core-srv
```

Build a docker image
```
make docker
```