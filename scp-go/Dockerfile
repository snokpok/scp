FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go get github.com/gin-gonic/gin/binding@v1.7.7
RUN go mod download

COPY src/ ./src/
COPY main.go .
RUN go build -o scp

EXPOSE 4000
CMD ["./scp"]
