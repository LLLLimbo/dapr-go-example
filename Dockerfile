FROM golang:1.16.5-alpine AS build
WORKDIR /app
ENV GOPROXY="https://goproxy.cn,direct"
ENV CGO_ENABLED=0
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /dapr-go-example

FROM scratch
WORKDIR /
COPY --from=build /dapr-go-example /dapr-go-example
EXPOSE 40002
ENTRYPOINT ["/dapr-go-example"]


MAINTAINER limbo