FROM golang:1.14-alpine

ENV WAREHOUSE_API_HOME "$GOPATH/src/ava.fund/alpha/notredame/warehouse_api"

COPY ./ $WAREHOUSE_API_HOME
WORKDIR $WAREHOUSE_API_HOME/src

RUN go mod download
RUN go build -o ../server ./cmd/server.go 

EXPOSE 1323

ENTRYPOINT ["../server", "../config.yaml"]