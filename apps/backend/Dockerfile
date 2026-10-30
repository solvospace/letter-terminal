FROM golang:1.27.3-bookworm AS builder

WORKDIR /app

COPY apps/backend/go.mod apps/backend/go.sum ./

RUN go mod download

COPY apps/backend/ .

RUN go build -o server ./cmd/server

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server ./server

EXPOSE 9090

CMD ["./server"]