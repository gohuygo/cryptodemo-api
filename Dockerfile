FROM golang:latest

RUN go get github.com/coincircle/go-coinmarketcap
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gorilla/context

WORKDIR /go/src/app

ADD vendor vendor
ADD . src

CMD ["go", "run", "src/main.go"]
