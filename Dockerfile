FROM golang:1.23.3-alpine

RUN apk add build-base

COPY go.mod go.sum ./

RUN go mod download

RUN mkdir /app
WORKDIR /app
ADD . /app

EXPOSE 8080     

RUN CGO_ENABLED=1 GOOS=linux go build -o example/hello

CMD [ "example/hello" ] 
