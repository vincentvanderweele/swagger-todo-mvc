FROM golang:1

RUN go get -v github.com/pilu/fresh

RUN mkdir -p /go/src/github.com/vincentvanderweele/swagger-todo-mvc
WORKDIR /go/src/github.com/vincentvanderweele/swagger-todo-mvc

COPY . /go/src/github.com/vincentvanderweele/swagger-todo-mvc
RUN go-wrapper install

CMD go-wrapper run
