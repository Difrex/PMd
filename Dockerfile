FROM golang

RUN go get -t -v github.com/Difrex/PMd

ENTRYPOINT cd /go/src/github.com/Difrex/PMd && go build && cp PMd /out
