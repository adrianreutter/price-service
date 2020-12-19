FROM golang:1.14-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /builder

COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod
RUN go mod download

COPY . .
RUN adduser -D -g '' appuser
# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /builder/bin/app ./cmd/run/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /builder/bin/app /priceservice
COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/ld-musl-x86_64.so.1

USER appuser
EXPOSE 8080
CMD ["/priceservice"]
