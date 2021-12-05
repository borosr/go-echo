FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/borosr/go-echo/

COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN go get -d -v && go mod vendor

ARG TARGETOS
ARG TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o /go/bin/go-echo

FROM scratch

COPY --from=builder /go/bin/go-echo /go/bin/go-echo
EXPOSE 8080
#ENTRYPOINT ["/go/bin/go-echo"]
CMD ["/go/bin/go-echo"]
