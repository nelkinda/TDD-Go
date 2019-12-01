FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
#CMD ["/bin/bash", "-l", "-c", "go test"]
#CMD CGO_ENABLED=0 go test
CMD  go test
#CMD ls
