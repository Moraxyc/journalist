ARG ARCH=
FROM ${ARCH}golang:alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go build

FROM ${ARCH}alpine:latest AS container

COPY --from=builder /go/src/app/journalist /usr/bin/journalist

CMD ["journalist"]
