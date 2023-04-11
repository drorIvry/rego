FROM golang:1.20

ARG PORT=4004

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .env ./
COPY *.go ./
COPY config ./config
COPY controllers ./controllers
COPY dao ./dao
COPY initializers ./initializers
COPY k8s ./k8s
COPY models ./models
COPY poller ./poller
COPY tasker ./tasker
COPY LICENSE ./LICENSE
COPY swagger-docs ./swagger-docs

RUN swag init --parseDependency --parseInternal

RUN  CGO_ENABLED=1 go build -o rego

EXPOSE ${PORT}

ENV IN_CLUSTER=true

CMD ["./rego"]
