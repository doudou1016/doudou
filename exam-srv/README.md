# Exam Service

This is the Exam service

Generated with

```
micro new exam-srv --namespace=com.lcb123 --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.lcb123.srv.exam
- Type: srv
- Alias: exam

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
./exam-srv
```

Build a docker image
```
make docker
```