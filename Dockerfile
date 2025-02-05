# Start from the latest golang base image
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN chmod o+rx /root

ENV DICTIONARY_API=https://api.dictionaryapi.dev/api/v2/entries/en/
ENV LINE_CHANNEL_SECRET=c391a341f8ec539a90ac49453e752d17
ENV LINE_CHANNEL_TOKEN=Ix9dFu/BeGP/ViVhhB8NdivmMMRqIMtDFn/kX5X2xxmHH3IeHrK4BOLKDRkoMXJvyaWOzL0JMAVc6e+go8yY4AidEEWM26mYMOBo8fdEYHll0HDvI1oRrblEgVz3AIafEaJQowCDHH8ZsGUpNsdC8gdB04t89/1O/w1cDnyilFU=

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"] 