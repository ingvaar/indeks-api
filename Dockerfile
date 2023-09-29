#--- Build stage
FROM golang:1.21-bullseye AS go-builder

WORKDIR /src

COPY . /src/

RUN make build CGO_ENABLED=0

#--- Image stage
FROM alpine:3.18.4

COPY --from=go-builder /src/target/dist/indeks-api /usr/bin/indeks-api

WORKDIR /opt

ENTRYPOINT ["/usr/bin/indeks-api"]
