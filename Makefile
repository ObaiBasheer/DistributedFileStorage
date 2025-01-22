build: 
	@go build -o bin/dfs -buildvcs=false

run: build
	@./bin/dfs

test:
	@go test -v ./...