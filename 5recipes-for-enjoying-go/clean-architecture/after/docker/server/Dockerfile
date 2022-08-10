FROM golang:1.18

WORKDIR /wwgt-diary-server
COPY . .
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD [ "air" ]
