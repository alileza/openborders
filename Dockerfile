FROM golang:1.17-alpine AS builder

COPY . /app
WORKDIR /app
RUN go build -o /bin/openborders

FROM alpine:3.15.0

COPY --from=builder /bin/openborders /bin/openborders
COPY ./index.html /index.html

ENTRYPOINT [ "/bin/openborders" ]

