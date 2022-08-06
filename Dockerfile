FROM golang:1.18 as builder

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.* ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -v -o srvtwo main.go

FROM debian:buster-slim

COPY --from=builder /app/srvtwo /app/srvtwo

CMD ["app/srvtwo"]