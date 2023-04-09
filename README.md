# Rego

<img align="right" width="300px" src="./rego.png">


[![Go Report Card](https://goreportcard.com/badge/github.com/drorivry/rego)](https://goreportcard.com/report/github.com/drorivry/rego)
[![GoDoc](https://pkg.go.dev/badge/github.com/drorivry/rego?status.svg)](https://pkg.go.dev/github.com/drorivry/rego?tab=doc)
[![Release](https://img.shields.io/github/release/drorivry/rego.svg?style=flat-square)](https://github.com/drorivry/rego/releases)
[![pages-build-deployment](https://github.com/drorIvry/rego/actions/workflows/pages/pages-build-deployment/badge.svg)](https://github.com/drorIvry/rego/actions/workflows/pages/pages-build-deployment)

[![](https://dcbadge.vercel.app/api/server/J6qKw7Zx)](https://discord.gg/J6qKw7Zx)

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

Clone the repo and run:

```sh
go run main.go

```

### Quick Start

With rego you can use the API to create and run k8s jobs with a managed API.

```sh
curl -X POST -d '{"image": "hello-world"}' localhost:4004/api/v1/task
```

This will start a job on your k8s cluster that'll run the docker image [hello-world](https://hub.docker.com/_/hello-world/)

the response will look something like 

```js
{
  "definition_id": "a36fbd9b-bf8a-4c59-94c1-9938b6707e8f",
  "message": "created"
}
```

we can use the definition ID to see the task's running status

```sh
curl http://localhost:4004/api/v1/task/a36fbd9b-bf8a-4c59-94c1-9938b6707e8f/latest
```

which will respond with

```js
{
  "ID": 0,
  "CreatedAt": "2023-04-02T11:53:08.2008054+03:00",
  "UpdatedAt": "2023-04-02T11:53:16.2032147+03:00",
  "DeletedAt": null,
  "id": "7eb53d97-7380-4e0b-82a6-b38fbf9119d2",
  "task_definition_id": "a36fbd9b-bf8a-4c59-94c1-9938b6707e8f",
  "status_code": 500,
  "status": "SUCCESS",
  "image": "hello-world",
  "name": "test",
  "ttl_seconds_after_finished": 10,
  "namespace": "test",
  "args": "[\"1111\", \"33333\"]",
  "metadata": {
    "ttlSecondsAfterFinished": 1
  }
}
```

which indicates the success.


### API Swagger
browse to `http://localhost:4004/swagger/index.htm`

## Contributing
We welcome contributions from the community! If you'd like to contribute to the project, please follow these guidelines:

- Fork the repository
- Create a new branch: git checkout -b new-feature
- Make your changes and commit them: git commit -m "Add new feature"
- Push your changes to your fork: git push origin new-feature
- Create a pull request to the main repository

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com//go-weather-app/blob/main/LICENSE) file for details.

## TODOs

- implement a CLI 
- docker deployment
- support deployment kickoff
- add workflow options
- support kubernetes metadata
- support actual external DBs