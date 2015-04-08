FROM golang:1.4.1

ADD . /go/src/github.com/mbrevoort/go-canary

RUN go install github.com/mbrevoort/go-canary

ENTRYPOINT /go/bin/go-canary

ENV PORT 3000

EXPOSE 3000