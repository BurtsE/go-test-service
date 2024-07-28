FROM golang:1.21-alpine AS builder

WORKDIR /build


ADD go.mod .
ADD go.sum .
RUN go mod download



COPY cmd cmd
COPY internal internal

RUN GOOS=linux go build -o app ./cmd

FROM golang:1.21-alpine
WORKDIR /root/

COPY configs configs

ENV PG_USER=${PG_USER}
ENV PG_PASSWORD=${PG_PASSWORD}
COPY --from=builder /build/app .


CMD ["./app"]