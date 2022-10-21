FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY src/ ./src/
COPY main.go .
RUN go build -o scp

EXPOSE 4000
CMD ["./scp"]
