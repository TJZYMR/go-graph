FROM golang:alpine
COPY . /GO-GRAPH
WORKDIR /GO-GRAPH
CMD go run main.go 
