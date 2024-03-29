FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o dawn-bot .

ENV TOKEN = ""

CMD ./dawn-bot $TOKEN
