FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app/subsea

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/subsea .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/subsea"]