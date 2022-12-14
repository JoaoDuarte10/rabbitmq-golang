FROM golang:1.19
WORKDIR /app/rabbitmq-golang
COPY . /app/rabbitmq-golang
RUN go get ./... && touch order.db
EXPOSE 3000
CMD [ "go", "run", "src/main.go" ]