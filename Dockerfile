FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /openproject-discord-webhook-proxy/server

EXPOSE 5001

CMD [ "/openproject-discord-webhook-proxy/server" ]
