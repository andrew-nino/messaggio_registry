################################################################################
# Registry
################################################################################
FROM golang:alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/registry

FROM scratch AS registry
COPY --from=builder /app/config /config
COPY --from=builder /app/schema /schema
COPY --from=builder /bin/app /app
ENV  CONFIG_PATH=./config/config.yaml
CMD ["/app"]
