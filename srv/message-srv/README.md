# Message-Srv Service

This is the Message-Srv service

Generated with

```
micro new github.com/iron-kit/unite-services/message-srv --namespace=kit.iron.srv --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: kit.iron.srv.srv.message-srv
- Type: srv
- Alias: message-srv

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
./message-srv-srv
```

Build a docker image
```
make docker
```