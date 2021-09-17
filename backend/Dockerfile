FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . . 
 
RUN go build -o /go-server

EXPOSE 8081

CMD ["/go-server"]
#docker run --name y-redis -p 7001:6379 -d redis
