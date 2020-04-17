FROM golang
RUN apt-get clean && apt-get update
WORKDIR /go/src/github.com/zaffka/newwords
ADD . .
RUN apt-get install -qy nano
RUN go get -d -v
RUN go install -v
ENTRYPOINT /go/bin/backend_task_advertising_site
EXPOSE 8080