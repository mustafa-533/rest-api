FROM golang:1.19-alpine3.16 as builder

RUN apk add --no-cache git bash sed build-base

RUN mkdir -p /go/src/github.com/mustafa-533/rest-api

WORKDIR /go/src/github.com/mustafa-533/rest-api

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPRIVATE=*.benzinga.io
ENV GOFLAGS=-mod=vendor

RUN export GIT_HEAD=$(git rev-parse --verify HEAD) && \
    echo "Applying build tag ${GIT_HEAD:0:8}" && \
    go build -v -a -installsuffix cgo -o  rest-api\
    -ldflags "-X main.build=${GIT_HEAD:0:8} -w -extldflags '-static'" -a -tags netgo \
    /go/src/github.com/mustafa-533/rest-api/main.go

# actual container
FROM golang:1.19-alpine3.16

WORKDIR /app

COPY --from=builder /go/src/github.com/mustafa-533/rest-api .

CMD ["./rest-api"]

# docker build -t registry.gitlab.benzinga.io/benzinga/feed-engine/feed-engine . -f Dockerfile