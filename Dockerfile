FROM golang:1.19.4-alpine AS builder

WORKDIR /go/src/app
COPY ./ .
RUN go build -o main main.go 

FROM alpine:latest
COPY --from=builder /go/src/app/main /usr/local/bin/
CMD ["main"]