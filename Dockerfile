# Step 1: Modules caching
FROM golang:1.21-alpine3.17 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Add custom certificates
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Step 2: Builder
FROM golang:1.21-alpine3.17 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd

# Step 3: Final

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/app /app
CMD ["/app"]