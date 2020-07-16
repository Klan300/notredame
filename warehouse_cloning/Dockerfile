FROM golang:1.14-alpine

ENV WAREHOUSE_CLONING_HOME "$GOPATH/src/ava.fund/alpha/Post-Covid/warehouse_cloning"

COPY ./ $WAREHOUSE_CLONING_HOME
WORKDIR $WAREHOUSE_CLONING_HOME/src

RUN go mod download
RUN go build -o ../main ./cmd/main.go 
RUN crontab ../cron

ENTRYPOINT ["crond" , "-f"]