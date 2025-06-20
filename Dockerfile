# syntax=docker/dockerfile:1.7   # <-- enables $TARGETARCH, $BUILDPLATFORM, etc.

###############################
# 1. Build stage (per‑arch)
###############################
FROM --platform=$BUILDPLATFORM golang:1.24 AS build

# Build‑time architecture vars supplied by BuildKit
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT   # e.g. "v7" for ARMv7

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Compile for the *target* platform, not the build platform
RUN CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
    GOARM=${TARGETVARIANT#v} \
    go build -o /out/server ./main.go

###############################
# 2. Minimal runtime stage
###############################
FROM --platform=$TARGETPLATFORM gcr.io/distroless/static:nonroot

COPY --from=build /out/server /server
USER nonroot:nonroot
ENTRYPOINT ["/server"]
