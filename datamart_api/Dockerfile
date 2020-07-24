FROM golang:1.14-alpine

ENV DATAMART_API_HOME "$GOPATH/src/ava.fund/alpha/Notre-Dame/datamart_api"

COPY ./ DATAMART_API_HOME
WORKDIR $DATAMART_API_HOME/src

RUN go mod download
RUN go build -o ../server ./cmd/server.go 

EXPOSE 1323

ENTRYPOINT ["../server", "../config.yaml"]