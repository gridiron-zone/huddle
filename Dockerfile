# To build the Huddle image, just run:
# > docker build -t huddle .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.huddle:/root/.huddle huddle huddle init
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.huddle:/root/.huddle huddle huddle start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 26657:26657 -p 26656:26656 -v ~/.huddle:/root/.huddle --name huddle huddle
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it huddle bash
#
# To exit the bash, just execute
# > exit
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-env
COPY --from=desmoslabs/desmos-build:latest /core/build/huddle /usr/bin/huddle

EXPOSE 26656 26657 1317 9090

# Run huddle by default, omit entrypoint to ease using container with huddle
CMD ["huddle"]
