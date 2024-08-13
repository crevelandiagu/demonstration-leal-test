FROM golang:1.22-alpine

WORKDIR /app

COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o app cmd/main.go

#EXPOSE the port
EXPOSE 5000

# Run the executable
CMD ["./app"]