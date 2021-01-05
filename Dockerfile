FROM golang:1.15.6 as builder

# Copy the code from the host and compile it
WORKDIR /go/src
COPY . .
RUN ls
RUN go mod download
RUN go build -o ./components-mock ./example/main.go

CMD ["./components-mock"]
