FROM golang:latest
RUN apt-get update
RUN apt-get install -y default-mysql-client default-libmysqlclient-dev default-mysql-server
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"
RUN mkdir /app 
#ADD . /app/ 
WORKDIR /app 
#RUN go build -o main . 
#CMD ["/app/main"]
CMD tail -f /dev/null

