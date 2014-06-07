
client: syntax
	go build ./...
	pushd deis/ && go build && popd

upx:
	upx -q ./deis/deis

format:
	@go fmt ./...

syntax: format
	@go vet ./...
	@golint .

test: syntax
	go test -v ./...
