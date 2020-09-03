FROM golang:1.13-alpine

ARG git_tag
ARG git_commit

RUN apk add --no-cache git build-base curl
RUN go get -u github.com/gobuffalo/packr/v2/packr2

ADD . /go/src/github.com/oeoen/push-notifications 
WORKDIR /go/src/github.com/oeoen/push-notifications
RUN packr2
RUN go mod vendor
RUN go build github.com/oeoen/push-notifications
RUN packr2 clean

FROM alpine

COPY --from=0 /go/src/github.com/oeoen/push-notifications/push-notifications /usr/bin/push-notifications

ENTRYPOINT ["push-notifications"]

CMD ["serve"]