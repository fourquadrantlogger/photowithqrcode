# Build
RUN go get && go build

EXPOSE 8005

CMD ["./photowithqrcode"]
