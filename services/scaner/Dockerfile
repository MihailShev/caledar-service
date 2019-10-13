FROM golang:1.12-alpine as builder
MAINTAINER mshev
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags="-w -s" -mod=vendor -o ./services/scaner/cmd/scaner ./services/scaner/cmd/

FROM alpine:3.10
COPY --from=builder /build/services/scaner/config.yml /go/bin/
COPY --from=builder /build/services/scaner/cmd/scaner /go/bin/scaner
WORKDIR /go/bin
RUN chmod +x ./scaner
CMD ["./scaner"]