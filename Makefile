build:
	@go build .

test:
	@go test -v .

run: build
	@./go-practice