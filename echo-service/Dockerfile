# Build
FROM golang:alpine AS builder
RUN mkdir /src

COPY go.mod /src
COPY echo.go /src
RUN cd /src; go build -o echo

# Final
FROM alpine
WORKDIR /app
COPY --from=builder /src/echo /app/
ENTRYPOINT ["./echo"]