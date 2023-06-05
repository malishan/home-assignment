FROM golang:alpine

LABEL maintainer="MD ALISHAN"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...

RUN go mod tidy

RUN go mod vendor

RUN go install -v ./...

RUN go build -o /build

EXPOSE 8080

ENTRYPOINT [ "/build", "--port", "8080", "--env", "dev"]