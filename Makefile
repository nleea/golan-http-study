build: 
	@go build -o bin/ecom cmd/main.go && make run

test:
	@go test -v ./..

run:
	@./bin/ecom