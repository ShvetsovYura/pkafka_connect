FROM golang:1.22 As builder
ENV GOOS linux
WORKDIR /app
COPY . .
RUN go build -o metrics_api ./cmd

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/metrics_api .

EXPOSE 8001
CMD ["./metrics_api"]