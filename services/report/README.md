# Report Service

This is the Report service

Generated with

```
micro new iunite.club/services/report --namespace=iunite.club --alias=report --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: iunite.club.srv.report
- Type: srv
- Alias: report

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
./report-srv
```

Build a docker image
```
make docker
```