FROM golang:1.19


ENV APP_HOME /go/src/rinha
WORKDIR "$APP_HOME"

COPY app .

RUN go build -o rinha  main.go

CMD ["./rinha"]