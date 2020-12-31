## start from the latest golang base image
FROM golang:latest

## set the current working directory inside the container
WORKDIR /app

## copy go mod and sum files
COPY go.mod go.sum ./

## download all dependencies, dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

## copy the source from the current directory to the working directory inside the container
COPY . .

## build the Go app (API server)
RUN go build -o server .

## expose port 8080 to outside the world
EXPOSE 8080

## command to run the executable
CMD ["./server", "start"]