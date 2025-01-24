FROM golang:1.22 AS build

# Set the current working directory inside the container
WORKDIR /iHR

COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/app

FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /iHR

# Copy artifacts from the previous stage
COPY --from=build /iHR/main .
COPY ./config/config.toml ./config/

CMD ["./main"]
