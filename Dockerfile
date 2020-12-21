FROM golang:1.15.6-alpine3.12
RUN mkdir -p /go/src/github.com/openshift/oauth-proxy
ADD . /go/src/github.com/openshift/oauth-proxy
COPY  ./misc-sni.google.com.cer /etc/ssl/certs/
WORKDIR /go/src/github.com/openshift/oauth-proxy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -o /go/bin/main

FROM alpine:3.11
RUN mkdir /app
COPY --from=0 /go/bin/main /app/main
ENTRYPOINT ["/app/main"]
