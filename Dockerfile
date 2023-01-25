FROM golang:alpine AS builder
WORKDIR /bin
ADD go.mod .
COPY . .
RUN go build -o app .
FROM alpine
WORKDIR /bin
COPY --from=builder /bin/app /bin/app
CMD ["./app"]