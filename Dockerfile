FROM golang:1.9.7 AS builder

RUN mkdir -p /go/src/github.com/mzbac/logService

WORKDIR /go/src/github.com/mzbac/logService

COPY . .

ENV CGO_ENABLED=0

ENV GOOS=linux

RUN go build -a -installsuffix cgo -o logService

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /go/src/github.com/mzbac/logService .

EXPOSE 8080

CMD ["./logService"]