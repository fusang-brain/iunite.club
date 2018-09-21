# Organization Service

This is the Organization service

Generated with

```
micro new iunite.club/services/organization --namespace=iunite.club --alias=organization --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.srv.organization
- Type: srv
- Alias: organization

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
./organization-srv
```

Build a docker image
```
make docker
```