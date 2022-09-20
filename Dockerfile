# ------------------------------------------------------------------------------
# Build flare-rosetta
# ------------------------------------------------------------------------------
FROM golang:1.18.1 AS rosetta

ARG ROSETTA_VERSION=main

COPY . /go/repo

WORKDIR /go/repo

ENV CGO_ENABLED=1
ENV GOARCH=amd64
ENV GOOS=linux

RUN git checkout $ROSETTA_VERSION

WORKDIR /go/repo/rosetta

RUN go mod download

RUN \
  GO_VERSION=$(go version | awk {'print $3'}) \
  GIT_COMMIT=main \
  make setup && \
  make build

# ------------------------------------------------------------------------------
# Target container for running the rosetta server
# ------------------------------------------------------------------------------
FROM ubuntu:20.04

# Install dependencies
RUN apt-get update -y && \
    apt-get install -y wget

WORKDIR /app

# Install rosetta server
COPY --from=rosetta \
  /go/repo/rosetta/rosetta-server \
  /app/rosetta-server

# Install rosetta runner
COPY --from=rosetta \
  /go/repo/rosetta/rosetta-runner \
  /app/rosetta-runner

# Install service start script
COPY --from=rosetta \
  /go/repo/rosetta/docker/entrypoint_flare.sh \
  /app/entrypoint_flare.sh

# Install configs
COPY --from=rosetta \
  /go/repo/config \
  /app/config

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint_flare.sh"]
CMD ["flare"]
