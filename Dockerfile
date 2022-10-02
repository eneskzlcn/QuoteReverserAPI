FROM   golang:1.18-alpine3.16
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy -go=1.18
RUN go build -o bin/quote-reverser-client .cmd/quote-reverser-client
CMD [ "bin/quote-reverser-client" ]
