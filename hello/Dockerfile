FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o app

FROM alpine
COPY --from=builder /go/src/app /app/app
ENTRYPOINT "/app/app"