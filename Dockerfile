#********************** FIRST STAGE ***************************
FROM golang:1.14-alpine AS builder

RUN apk add --no-cache ca-certificates git
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/joho/godotenv

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
COPY --from=builder /app/.env .

EXPOSE 80

CMD ["sh", "-c", "source .env && ./app"]
