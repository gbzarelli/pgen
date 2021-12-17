FROM golang:1.17
EXPOSE 5000

ENV PROTOCOL_DECIMAL_PLACES_AFTER_DATE=8
ENV REDIS_ADDRESS='redis:6379'

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN go mod tidy
RUN go build -o main .

CMD ["/app/main"]
