FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -C ./src/cmd -a -installsuffix cgo -o /app/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY --from=builder /app/.env .

EXPOSE 8080

CMD [ "./app" ]
