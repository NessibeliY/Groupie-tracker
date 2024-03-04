FROM golang:1.20.1-alpine
WORKDIR /groupie-tracker
COPY . .
RUN go build -o groupie-tracker ./cmd/myapp/
CMD ["./groupie-tracker"]