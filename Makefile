build:
	go build -o bin/quote-reverser-client .cmd/quote-reverser-client
clean:
	rm -rf bin
run:
	./bin/quote-reverser-client