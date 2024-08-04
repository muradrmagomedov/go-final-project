FROM golang:1.22.2

RUN mkdir todo
WORKDIR /todo
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o app ./cmd/.
EXPOSE 8080
CMD ["./app"]

