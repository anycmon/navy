FROM golang:1.20 AS builder
WORKDIR /build
COPY go.mod ./
RUN go mod download
ADD . /build/
RUN go build -o navy-ports cmd/main.go

FROM golang:1.20
COPY --from=builder /build/navy-ports /app/
WORKDIR /app

EXPOSE ${PORT}
CMD ["./navy-ports"]