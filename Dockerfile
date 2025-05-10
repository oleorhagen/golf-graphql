# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY graph ./graph

RUN CGO_ENABLED=0 GOOS=linux go build  -o golf-graphql ./server.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/golf-graphql /bin/golf-graphql

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bin/golf-graphql"]
