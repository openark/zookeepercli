FROM gliderlabs/alpine:3.2

MAINTAINER Ryan Eschinger <ryanesc@gmail.com>

COPY . /go/

RUN apk add --update go git \
  && cd /go/src/github.com/outbrain/zookeepercli/ \
  && export GOPATH=/go \
  && go get \
  && go build -o /bin/zookeepercli \
  && rm -rf /go \
  && apk del --purge go git

ENTRYPOINT ["/bin/zookeepercli"]
