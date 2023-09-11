# Official Golang runtime
FROM golang:alpine as builder

# Git Installation
RUN apk update && apk add --no-cache git

# Set WORKDIR
WORKDIR /app

# Copy go.mod & go.sum
COPY go.mod go.sum ./

# Downloading all dependencies
RUN go mod download

# Copying project to container
COPY . .

# Building the project
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Starting new stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

# Expose 8080
EXPOSE 8080

# Run
CMD ["./main"]