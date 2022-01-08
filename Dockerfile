FROM golang:1.17.6-alpine3.15

WORKDIR /go/src/app

# Add gcc for Go unit testing
RUN apk add build-base

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD ["passwordgenerator"]