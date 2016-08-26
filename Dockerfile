# Build
FROM golang:latest

COPY ./ $GOPATH/src/code.aliyun.com/mougew/photowithqrcode

WORKDIR $GOPATH/src/code.aliyun.com/mougew/photowithqrcode

RUN go get && go build

EXPOSE 8005

CMD ["./photowithqrcode"]
