# Rego

<img align="right" width="300px" src="./rego.png">


[![Go Report Card](https://goreportcard.com/badge/github.com/drorivry/rego)](https://goreportcard.com/report/github.com/drorivry/rego)
[![GoDoc](https://pkg.go.dev/badge/github.com/drorivry/rego?status.svg)](https://pkg.go.dev/github.com/drorivry/rego?tab=doc)
[![Release](https://img.shields.io/github/release/drorivry/rego.svg?style=flat-square)](https://github.com/drorivry/rego/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/drorivry/rego)](https://www.tickgit.com/browse?repo=github.com/drorivry/rego)


> rego /ˈre.ɡoː/
verb
Definitions:
manage, direct
rule, guide

Rego is a 

- 🔥  blazingly fast
- 🥇 API first
- 🌈 lightweight
- 🕜 Task orchestrator

It is designed to allow asynchronous workloads to be deployed over Kubernetes with minimal effort, while also providing a management API that can keep track of progress and run history.

### Use cases

- Run async workloads that need s to be managed (or visible to) a UI
- integrate non production-grade code (data scientist R code for example) within your production environment in a contained way
- use to run stuff periodically with run history

## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release) (we test it with these).

### Getting Rego

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply run

```sh
go run main.go

```

### API Swagger
browse to `http://localhost:4004/swagger/index.htm`

## TODOs

- implement a CLI 
- docker deployment
- support deployment kickoff
- add workflow options
- support kubernetes metadata
- support actual external DBs