FROM golang:latest

RUN mkdir /app

ADD . /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all the dependencies
RUN go build -o main . 

# Run the executable
CMD ["/app/main"]