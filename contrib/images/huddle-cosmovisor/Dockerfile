ARG HUDDLE_VERSION="v0.17.2"
FROM golang:1.18-alpine AS build-env

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

ENV GOBIN /go/bin

# Set working directory for the build
WORKDIR /go/src/github.com/cosmos/cosmos-sdk
RUN git clone https://github.com/cosmos/cosmos-sdk.git /go/src/github.com/cosmos/cosmos-sdk
RUN git checkout tags/cosmovisor/v1.0.0
RUN cd /go/src/github.com/cosmos/cosmos-sdk/cosmovisor
RUN make cosmovisor

# Final image
FROM huddlelabs/huddle:$HUDDLE_VERSION

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/cosmos/cosmos-sdk/cosmovisor/cosmovisor /usr/bin/cosmovisor

ARG UID=1000
ARG GID=1000
USER ${UID}:${GID}

COPY wrapper.sh /usr/bin/wrapper.sh
ENTRYPOINT ["bash", "/usr/bin/wrapper.sh"]

# Run cosmovisor by default, omit entrypoint to ease using container with Huddle
CMD ["cosmovisor"]
STOPSIGNAL SIGTERM