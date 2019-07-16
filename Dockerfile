FROM golang:1.12-alpine

RUN apk add git
RUN apk --update add ca-certificates

RUN adduser -D app


RUN su - app
WORKDIR /home/app

COPY . .

RUN go build -o twitter-serverless-events .

ENTRYPOINT ["/home/app/twitter-serverless-events"]

