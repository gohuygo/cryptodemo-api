FROM golang:latest AS builder

COPY . $GOPATH/src/github.com/gohuygo/cryptodemo-api
WORKDIR $GOPATH/src/github.com/gohuygo/cryptodemo-api

RUN go get ./
# RUN go get github.com/coincircle/go-coinmarketcap
# RUN go get github.com/gorilla/mux
# RUN go get github.com/dgrijalva/jwt-go
# RUN go get github.com/gorilla/context


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .
EXPOSE 80
FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
