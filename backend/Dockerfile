FROM golang:alpine

# Set necessary environmet variables needed for our image
# PRODUCTION = true|false
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
 
# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY . .
RUN go mod download

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# REMUEVE TODO EL CODIGO NO NECESARIO
WORKDIR /
RUN rm -rf /build

CMD [ "/dist/main" ]