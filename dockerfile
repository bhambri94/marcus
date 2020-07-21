FROM golang:1.14

ENV APP_USER app
ENV APP_HOME /go/src/marcus

RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD . /app

# setting working directory
WORKDIR /go/src/app

# installing dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure 

COPY / /go/src/app/
RUN go build -o main .

CMD ["/app/main"]