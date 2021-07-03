FROM golang:latest
WORKDIR /go/src/app
COPY . .
CMD go run GoRH.go