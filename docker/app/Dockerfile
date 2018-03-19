FROM golang:1.9-alpine

WORKDIR /go/src/github.com/marjoram/duke-crd
COPY . .
RUN ./scripts/build.sh

FROM alpine:latest
RUN apk --no-cache add \
  ca-certificates
COPY --from=0 /go/src/github.com/duke-crd/bin/linux/duke-crd /usr/local/bin

ENTRYPOINT ["/usr/local/bin/duke-crd"]
