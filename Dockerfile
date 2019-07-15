FROM golang:1.12-alpine

RUN apk add git
RUN adduser -D app

RUN apk --update add ca-certificates

USER app
WORKDIR /home/app

COPY . .

RUN go build -o GO-Serverless .

CMD /home/app/GO-Serverless
