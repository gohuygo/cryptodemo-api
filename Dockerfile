FROM golang:latest AS builder

COPY . $GOPATH/src/github.com/gohuygo/cryptodemo-api
WORKDIR $GOPATH/src/github.com/gohuygo/cryptodemo-api

RUN go get ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM centurylink/ca-certs
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
