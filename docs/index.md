# Welcome REGO
<p align="center">
    <img align="center" width="300px" src="./rego.png">
</p>

<p align="center">
    <em>rego /ˈre.ɡoː/ verb. Definitions: manage, direct rule, guide</em>
</p>

<p align="center">
    <em>REGO is a modern, fast (high-performance), framework for running any docker image as a kubernetes job with a fully managed api</em>
</p>


[![Go Report Card](https://goreportcard.com/badge/github.com/drorivry/rego)](https://goreportcard.com/report/github.com/drorivry/rego)
[![GoDoc](https://pkg.go.dev/badge/github.com/drorivry/rego?status.svg)](https://pkg.go.dev/github.com/drorivry/rego?tab=doc)
[![Release](https://img.shields.io/github/release/drorivry/rego.svg?style=flat-square)](https://github.com/drorivry/rego/releases)
[![pages-build-deployment](https://github.com/drorIvry/rego/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/drorIvry/rego/actions/workflows/pages/pages-build-deployment)
[![Build](https://github.com/drorIvry/rego/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/drorIvry/rego/actions/workflows/go.yml)

[![](https://dcbadge.vercel.app/api/server/J6qKw7Zx)](https://discord.gg/J6qKw7Zx)

Rego is a 

- 🔥  blazingly fast
- 🥇 API first
- 🌈 lightweight
- 🕜 Task orchestrator

It is designed to allow asynchronous workloads to be deployed over Kubernetes with minimal effort, while also providing a management API that can keep track of progress and run history.


Star us on [github](https://www.github.com/drorivry/rego).

## Use cases

- Run async workloads that needs to be managed (or visible to) a UI
- Integrate non production-grade code (data scientist's R code for example) within your production environment in a contained way
- Use to run stuff periodically with run history


## Requirements


- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release) (we test it with these).

- **[Kubernetes](https://kubernetes.io/)**: tested on v1.24+

## Installation

### Run Kubernetes

```sh
kubectl apply -f https://raw.githubusercontent.com/drorIvry/rego/main/deploy/deployment.yml
```

### Run on local machine

```sh
curl -L  https://raw.githubusercontent.com/drorIvry/rego/main/local-deploy/rego.sh | sh
```

## Quick Start

## Contributing
We welcome contributions from the community! If you'd like to contribute to the project, please follow these guidelines:

- Fork the repository
- Create a new branch: git checkout -b new-feature
- Make your changes and commit them: git commit -m "Add new feature"
- Push your changes to your fork: git push origin new-feature
- Create a pull request to the main repository

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com//go-weather-app/blob/main/LICENSE) file for details.