FROM golang:1.12-alpine as builder
MAINTAINER mshev
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -mod=vendor -o ./services/api/server/server ./services/api/server/

FROM alpine:3.10
COPY --from=builder /build/services/api/server/server /go/bin/server
COPY --from=builder /build/services/api/config.yml /go/bin/server
WORKDIR /go/bin
CMD ["server"]

