# Use an official Go runtime as a parent image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY ./src /go/src/app

# Build the Go app

RUN go mod tidy
RUN go build -o main .

# Specify the command to run on container startup
CMD ["./main"]
