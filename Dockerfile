FROM golang:1.25-alpine as builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o gonverter-service ./app/**/*.go
RUN chmod  +x ./gonverter-service

FROM alpine:latest
RUN apk --update --no-cache add curl

RUN mkdir /app
COPY --from=builder /app/gonverter-service /app

EXPOSE 8080

CMD [ "/app/gonverter-service" ]
