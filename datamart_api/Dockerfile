FROM golang:1.14-alpine

ENV DATAMART_API_HOME "$GOPATH/src/ava.fund/alpha/Post-Covid/datamart_api"

COPY ./ DATAMART_API_HOME
WORKDIR $DATAMART_API_HOME/src

RUN go mod download
RUN go build -o ../server ./cmd/server.go 

EXPOSE 1324

ENTRYPOINT ["../server", "../config.yaml"]