# ------------------------------------------------------------------------------
# Build flare
# ------------------------------------------------------------------------------
FROM golang:1.17 AS flare

ARG FLARE_VERSION

RUN git clone https://github.com/flare-foundation/flare.git \
  /go/src/github.com/flare-foundation/flare

WORKDIR /go/src/github.com/flare-foundation/flare

RUN git checkout $FLARE_VERSION && \
    ./scripts/build.sh

# ------------------------------------------------------------------------------
# Build flare rosetta
# ------------------------------------------------------------------------------
FROM golang:1.17 AS rosetta

ARG ROSETTA_VERSION

RUN git clone https://github.com/flare-foundation/flare-rosetta.git \
  /go/src/github.com/flare-foundation/flare-rosetta

WORKDIR /go/src/github.com/flare-foundation/flare-rosetta

RUN git checkout $ROSETTA_VERSION && \
    go mod download

RUN \
  GO_VERSION=$(go version | awk {'print $3'}) \
  GIT_COMMIT=$(git rev-parse HEAD) \
  make build

# ------------------------------------------------------------------------------
# Target container for running the node and rosetta server
# ------------------------------------------------------------------------------
FROM ubuntu:20.04

# Install dependencies
RUN apt-get update -y && \
    apt-get install -y wget

WORKDIR /app

# Install flare daemon
COPY --from=flare \
  /go/src/github.com/flare-foundation/flare/build/flare \
  /app/flare

# Install evm plugin
COPY --from=flare \
  /go/src/github.com/flare-foundation/flare/build/plugins/evm \
  /app/plugins/evm

# Install list of FBA validators
COPY --from=flare \
  /go/src/github.com/flare-foundation/flare/scripts/configs/songbird/fba_validators.json \
  /app/songbird_validators.json

# Install rosetta server
COPY --from=rosetta \
  /go/src/github.com/flare-foundation/flare-rosetta/rosetta-server \
  /app/rosetta-server

# Install rosetta runner
COPY --from=rosetta \
  /go/src/github.com/flare-foundation/flare-rosetta/rosetta-runner \
  /app/rosetta-runner

# Install service start script
COPY --from=rosetta \
  /go/src/github.com/flare-foundation/flare-rosetta/docker/entrypoint.sh \
  /app/entrypoint.sh

EXPOSE 9650
EXPOSE 9651
EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]
