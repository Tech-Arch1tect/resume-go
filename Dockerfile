FROM golang:1.20.5 AS build

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o /resume-go

## Deploy
FROM alpine:3.18.2

RUN apk add libc6-compat

WORKDIR /app

COPY --from=build /resume-go /resume-go

ENTRYPOINT ["/resume-go"]