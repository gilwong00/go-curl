.PHONY: start
start:
	go run cmd/main.go https://example.com/get

.PHONY: start-header
start-header:
	go run cmd/main.go https://example.com/get -H 'Content-Type: application/json'