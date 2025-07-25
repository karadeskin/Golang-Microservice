#official Golang image
FROM golang:1.22-alpine

#working directory
WORKDIR /app

#copy source files
COPY . .

#build the Go app
RUN go build -o server

#run the app
CMD ["./server"]