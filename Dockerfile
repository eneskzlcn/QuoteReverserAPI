FROM    golang:1.17.0-alpine3.14
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy -go=1.16 && go mod tidy -go=1.17
RUN make build
CMD [ "bin/main" ]
