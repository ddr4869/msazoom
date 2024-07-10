FROM golang:1.22.2-alpine as builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o go_template

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/go_template /go_template
COPY .env .env

CMD ["/go_template"]