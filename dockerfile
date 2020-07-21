FROM golang:1.14

ENV APP_USER app
ENV APP_HOME /go/src/marcus

# setting working directory
WORKDIR /go/src/app

COPY / /go/src/app/

# installing dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure 

RUN go build -o marcus

CMD ["./marcus"]