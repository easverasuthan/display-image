## Downloading Latest go image
FROM golang:latest

## Creae a new directory within our image
RUN mkdir /app

## Copy the code from source to the container
COPY . /app

## Set the work directory as app
WORKDIR /app

## Build the go application
RUN go build -o main .

## Expose a port to connect with our application
EXPOSE 80

## Command to run when starting the container
CMD ["/app/main"]
