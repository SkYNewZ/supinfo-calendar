FROM golang:1.10-alpine3.8 as BUILDER
WORKDIR /go/src/github.com/skynewz/supinfo-calendar

RUN apk add --update --no-cache \
        git \
        curl \
        ca-certificates && \
        go get -u github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

ENV OUTPUT_PATH /out

WORKDIR /
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /go/src/github.com/skynewz/supinfo-calendar/main .

VOLUME /out

ENTRYPOINT ["/main"]