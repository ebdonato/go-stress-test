# Build stage
FROM golang:1.21.6 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags '-s -w' \
    -o goapp main.go

# Run stage
FROM scratch
ARG url
ARG requests
ARG concurrency
COPY --from=build /app/goapp .
CMD ["./goapp"]
