FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY ..

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amb6 go build

ENTRYPOINT ["/app/main"]
