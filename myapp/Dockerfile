FROM golang:latest

RUN apt-get update && apt-get install make bash

WORKDIR /go/src/

COPY . .

WORKDIR /go/src/myapp/

RUN go mod download -x

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
