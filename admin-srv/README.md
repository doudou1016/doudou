# Admin Service

This is the Admin service

Generated with

```
micro new admin-srv --namespace=com.lcb123 --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.lcb123.srv.admin
- Type: srv
- Alias: admin

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./admin-srv
```

Build a docker image
```
make docker
```