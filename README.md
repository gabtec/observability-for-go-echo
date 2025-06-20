# Go Echo Example App, instrumented

This is an example app, to learn observability.

## Endpoints

| Method | Endpoint    | Description                                                       |
| ------ | ----------- | ----------------------------------------------------------------- |
| GET    | /           | index page - will show a list of available endpoints              |
| GET    | /random     | will generate random logs, either success or error                |
| GET    | /log/{type} | will generate random logs, but we define the type "ok" or "error" |
| GET    | /metrics    | the metrics endpoint, using prometheus (only in v2)               |

## Versions

- v1 - will not be instrumented. Only outputs logs to stdout
- v2 - will be instrumented:
  - <u>metrics</u> exposed in /metrics endpoint
  - <u>logs</u> will be send to stdout, since for now, openTelemetry for Go is not implemented
  - <u>traces</u> send to an endpoint defined using env variable: OTEL_ENDPOINT
    - example "OTEL_ENDPOINT=otel-collector:4816"
    - IMPORTANT: port must be provided (to allow http, grpc or custom collectors)

```ini
# REQUIRED only in v2  (OpenTelemetryCollector Service)
OTEL_ENDPOINT="localhost:4316"
# Optional: will default to "repo name"
OTEL_SERVICE_NAME="my-echo-log-app"
# Optional: will default to ":1323"
SERVER_ADDR=":1323"
```

## How I did it

```sh
go mod init gabtec/go-echo-obs-app
# deps
go get github.com/labstack/echo/v4
go install github.com/air-verse/air@latest

air init
# v1
```

## Usage

```sh
# start local server
go run main.go

# OR start with live reload
air
```

## How to build container image

```sh
# option 1, run:
docker build -t _IMAGE_NAME_:_IMAGE_TAG_

# option 2, use github action included:
# it is "manually" dispatched, and we can point to one og the repo tag
```
