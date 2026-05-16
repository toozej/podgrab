# setup project and deps
FROM golang:1.26-trixie AS init

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

WORKDIR /go/podgrab/

COPY go.mod* go.sum* ./
RUN go mod download

COPY . ./

FROM init AS vet
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN go vet ./...

# run tests
FROM init AS test
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN go test -coverprofile c.out -v ./...

# build binary
FROM init AS build
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ARG LDFLAGS

RUN CGO_ENABLED=0 go build -ldflags="${LDFLAGS}"

# Install coreutils for sleep and other utilities utilized in devcontainer
RUN apt-get update && apt-get install --no-install-recommends -y coreutils

# runtime image
FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /go/bin/
# Copy our static executable.
COPY --from=build /go/podgrab/podgrab /go/bin/podgrab
# Expose port for publishing as web service
EXPOSE 8080
# Setup volumes for data
ENV CONFIG=/config
ENV DATA=/assets
ENV UID=998
ENV PID=100
ENV GIN_MODE=release
VOLUME ["/config", "/assets"]

USER nonroot
# Run the binary.
ENTRYPOINT ["/go/bin/podgrab"]
