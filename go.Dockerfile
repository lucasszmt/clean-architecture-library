FROM golang:buster as builder

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

CMD go build -o ./tmp/cleanarch main.go && ./tmp/cleanarch