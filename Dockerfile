ARG GO_VERSION=1.22.1
ARG PORT=8080

FROM golang:${GO_VERSION}-alpine AS Builder

ENV PORT=$PORT

WORKDIR /app
COPY . .

RUN go mod download

RUN go build -o /app/main

FROM alpine:latest 

ENV PORT=$PORT

WORKDIR /app

COPY --from=Builder /app/main .

EXPOSE $PORT

CMD ["./main"]