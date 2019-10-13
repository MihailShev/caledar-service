FROM golang:1.12-alpine as builder
MAINTAINER mshev
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -mod=vendor -o ./services/notify/notify ./services/notify/

FROM alpine:3.10
COPY --from=builder /build/services/notify/config.yml /go/bin/
COPY --from=builder /build/services/notify/notify /go/bin/notify
WORKDIR /go/bin
RUN chmod +x ./notify
CMD ["./notify"]