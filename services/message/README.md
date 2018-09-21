# Message Service

This is the Message service

Generated with

```
micro new iunite.club/services/message --namespace=iunite.club --alias=message --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.srv.message
- Type: srv
- Alias: message

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
./message-srv
```

Build a docker image
```
make docker
```