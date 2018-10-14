# Restful Service

This is the Restful service

Generated with

```
micro new iunite.club/services/restful --namespace=iunite.club --alias=restful --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.web.restful
- Type: web
- Alias: restful

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
./restful-web
```

Build a docker image
```
make docker
```