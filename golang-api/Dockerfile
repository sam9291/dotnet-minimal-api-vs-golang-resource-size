############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/package/app/
COPY . .
# Build the binary.
RUN go build -o /go/bin/golang-api ./main.go
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/golang-api /go/bin/golang-api
# Run the binary.
ENTRYPOINT ["/go/bin/golang-api"]