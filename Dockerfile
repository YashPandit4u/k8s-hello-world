FROM golang:tip-trixie as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server main/main.go

FROM alpine:edge
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD [ "./server" ]
