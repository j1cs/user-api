FROM golang:bookworm AS builder
WORKDIR /home/app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o command ./cmd/api/main.go

FROM golang:bookworm
RUN apt-get update -q && apt-get install -yq bash curl
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /home/app
COPY --from=builder /home/app/command /home/app/command
COPY ./build/cmd/start.sh /home/app/start.sh
RUN chmod +x /home/app/start.sh
ENTRYPOINT ["/home/app/start.sh"]
