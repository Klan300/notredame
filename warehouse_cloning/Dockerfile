FROM golang:1.14-alpine

WORKDIR /go/src/ava.fund/alpha/Post-Covid/warehouse_cloning

COPY ./ ./

WORKDIR /go/src/ava.fund/alpha/Post-Covid/warehouse_cloning/src

RUN go mod download

RUN go build -o ../main ./cmd/main.go 

ENTRYPOINT ["../main" , "../config.yaml"]