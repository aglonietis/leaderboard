FROM golang:1.18.3

RUN mkdir /app
WORKDIR /app
COPY . .
RUN rm main || true
RUN env GOOS=linux GOARCH=amd64 go build -o main cmd/api/main.go
RUN cp .env.example .env

CMD ["/app/main"]
EXPOSE 8080