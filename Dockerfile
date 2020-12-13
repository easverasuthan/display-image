#********************** FIRST STAGE ***************************
FROM golang:1.14-alpine AS builder

#Set the Current Working Directory
WORKDIR /app

# Copy everything from the current directory to the Working Directory
COPY . .

# Start the build and save it as `user`
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/app -v


#********************** SECOND STAGE ***************************

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
#Copy
COPY --from=builder /app/build/app .

EXPOSE 80

CMD ["./app"]
