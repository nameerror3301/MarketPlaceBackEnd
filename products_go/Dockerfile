FROM golang:1.19 AS builder

WORKDIR /service

COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /products ./cmd/app/main.go

FROM alpine:latest
COPY --from=builder /products /bin
ENTRYPOINT ["products"]
