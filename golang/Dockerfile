FROM golang:1.19

WORKDIR /app

# Download Go modules
COPY golang/go.mod .
COPY golang/go.sum .
RUN go mod download

# Build
COPY *.go .
RUN go build -o hello-world

EXPOSE 8000

# Run
CMD [ "./hello-world" ]
