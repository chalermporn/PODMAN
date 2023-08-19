FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .

RUN go get \
&& go build -o /go/bin/app


FROM alpine

RUN apk update \
&& apk add ca-certificates \
&& rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]