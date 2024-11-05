FROM golang:1.23.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o . ./...

FROM alpine:latest

EXPOSE 8080

COPY --from=build /app/upfluencett .

ENTRYPOINT ["./upfluencett", "-upfluence-url=https://stream.upfluence.co/stream", "-server-port=8080"]