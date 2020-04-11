FROM golang:1.14.2-alpine

WORKDIR /go/src

ENV GO111MODULE=on
ADD go.mod go.sum main.go ./

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o /go/src/main
CMD ["./main"]