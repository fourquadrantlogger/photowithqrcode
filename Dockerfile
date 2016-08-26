# Build
FROM golang:latest

COPY ./ $GOPATH/src/code.aliyun.com/mougew/photowithqrcode

RUN go get && go build

WORKDIR $GOPATH/src/code.aliyun.com/mougew/photowithqrcode

EXPOSE 8005

CMD ["./photowithqrcode"]
