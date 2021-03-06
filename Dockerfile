# Build Stage
# First pull Golang image
FROM golang:1.18-alpine as development
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install air for development
RUN go install github.com/cosmtrek/air@latest

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN export PATH=$(go env GOPATH)/bin:$PATH

# Expose port
EXPOSE 8080
# Start app
CMD air