# Build
RUN go get && go build

EXPOSE 80

CMD ["./photowithqrcode"]
