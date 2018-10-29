# Approved Service

This is the Approved service

Generated with

```
micro new iunite.club/services/approved --namespace=iunite.club --alias=approved --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.srv.approved
- Type: srv
- Alias: approved

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
./approved-srv
```

Build a docker image
```
make docker
```