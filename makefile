clean:
	@rm -rf bin

build:
	@go build -o bin/kal

run: build
	@./bin/kal
