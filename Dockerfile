# syntax=docker/dockerfile:1

# Build stage — pick an image tag that matches your go.mod (1.25.x)
FROM golang:1.25-bookworm AS build

WORKDIR /src

# Dependencies first (you have no go.sum yet; only go.mod is enough for cache)
COPY go.mod ./
RUN go mod download

# App source
COPY . .

# Static binary: no libc/CGO issues in minimal runtime image
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/server .

# Runtime — small, non-root
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Binary + config (paths are relative to WORKDIR, same as local "repo root")
COPY --from=build --chown=nonroot:nonroot /out/server /app/server
COPY --chown=nonroot:nonroot config/switch_file.json /app/config/switch_file.json

# Your code uses PORT or JSON default
ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]