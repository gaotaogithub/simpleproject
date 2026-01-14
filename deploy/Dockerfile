# ---------- build stage ----------
ARG GO_VERSION=1.24.2
FROM golang:${GO_VERSION}-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# ---------- runtime stage ----------
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8280
ENTRYPOINT ["./app"]