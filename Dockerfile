FROM golang:alpine AS build-env

RUN apk update && apk add --no-cache gcc musl-dev

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN go build -o app cmd/main.go

FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=build-env /app/app .

EXPOSE 8080

CMD ["app"]
