# Build stage
FROM golang:1.22.0 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags '-s -w' \
    -o goapp main.go

# Run stage
FROM alpine
COPY --from=build /app/goapp .
ENTRYPOINT ["./goapp"]
