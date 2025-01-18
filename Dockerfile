FROM alpine:latest

RUN apk update
RUN apk add go

RUN go install github.com/mikerybka/brass/cmd/api@latest

ENTRYPOINT ["/root/go/bin/api"]
