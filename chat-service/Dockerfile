FROM golang:1.22.2-alpine as builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o msazoom

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/msazoom /msazoom
COPY .env .env

CMD ["/msazoom"]