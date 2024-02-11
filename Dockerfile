FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY config ./config
COPY controllers ./controllers
COPY dao ./dao
COPY initializers ./initializers
COPY k8s ./k8s
COPY models ./models
COPY poller ./poller
COPY tasker ./tasker

RUN go mod download
RUN CGO_ENABLED=1 go build -o rego


FROM golang:1.22

ARG PORT=4004

WORKDIR /app

COPY --from=builder /app/rego ./rego

RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY swagger-docs ./swagger-docs
RUN swag init --parseDependency --parseInternal

EXPOSE ${PORT}

ENV IN_CLUSTER=true

CMD ["./rego"]
