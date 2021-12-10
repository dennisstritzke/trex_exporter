###########
# BUILDER #
###########

FROM golang:1.15 as builder

# set work directory
WORKDIR /usr/src/app

COPY . .

# set build environments and build the app
ARG GOOS=linux
ARG GOARCH=$TARGETARCH
ARG CGO_ENABLED=0
RUN go build -a -installsuffix cgo -o bin/trex
RUN chmod +x bin/trex

#########
# FINAL #
#########

FROM alpine:3.15

# create directory for the app and set it as a work dircetory
RUN mkdir -p /trex_exporter
WORKDIR /trex_exporter

# copy built files
COPY --from=builder /usr/src/app/bin .

# create non-privelleged user for running the app
RUN addgroup -S app && adduser -S app -G app
RUN chown -R app:app /trex_exporter
USER app

# define environments
ENV TREX_EXPORTER_PORT=9788
ENV TREX_EXPORTER_BIND_ADDRESS=0.0.0.0
ENV TREX_MINER_URL=http://localhost:4057
ENV TREX_WORKER_NAME=trex

# run the app
CMD ./trex --api-address=$TREX_MINER_URL --web.listen-address=$TREX_EXPORTER_BIND_ADDRESS:$TREX_EXPORTER_PORT --worker=$TREX_WORKER_NAME
