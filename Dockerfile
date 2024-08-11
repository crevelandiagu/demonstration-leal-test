FROM golang:1.18-alpine

WORKDIR /app

COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#EXPOSE the port
EXPOSE 5000

# Run the executable
CMD ["./api"]