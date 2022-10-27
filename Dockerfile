FROM golang:1.19 AS builder

WORKDIR /backend

COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /service-backend ./cmd/app/main.go

FROM alpine:latest

COPY --from=builder /service-backend /bin

ENTRYPOINT ["service-backend"]