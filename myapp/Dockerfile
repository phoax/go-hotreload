FROM golang:latest

WORKDIR /go/src/app/

COPY . .

RUN go mod download -x

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
