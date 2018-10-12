FROM golang:latest

RUN mkdir -p "/usr/local/go/src/github.com/setgetdb/setgetdb"
ADD . /usr/local/go/src/github.com/setgetdb/setgetdb/
WORKDIR /usr/local/go/src/github.com/setgetdb/setgetdb

RUN go build -o server .
EXPOSE 10101/tcp
CMD ["./server"]
