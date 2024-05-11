FROM golang:1.16-alpine

WORKDIR /app

# COPY go.mod, go.sum and download the dependencies
COPY src/go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY src .
RUN go build -o /app/build/myapp .

EXPOSE 8080
ENTRYPOINT [ "/app/build/myapp" ]