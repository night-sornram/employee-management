# Start from the latest golang base image
#FROM golang:latest AS builder-common-utils
#
#WORKDIR /app/common-utils
#
#COPY ../common-utils/go.mod ./
#
#RUN go mod download
#
#COPY ../common-utils/ ./
#
#RUN go build -o common-utils .

FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

#COPY --from=builder-common-utils /app/common-utils /app/common-utils

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]