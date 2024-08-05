.PHONY: start
start:
	go run cmd/main.go

.PHONY: example
example:
	go run cmd/main.go https://ron-swanson-quotes.herokuapp.com/v2/quotes -H 'Content-Type: application/json' -v