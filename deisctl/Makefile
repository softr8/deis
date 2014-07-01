COMPONENTS=builder cache controller database logger registry router

test:
	godep go test -v -cover ./...

build:
	cd deisctl && godep go build .
	@echo "binary placed in ./deisctl/deisctl"

install:
	godep go install ./...
	@echo "binaries placed in \$$GOPATH/bin"

package:
	docker build -t deis/deisctl .
	CONTAINER=`docker run -d deis/deisctl` ; \
	docker cp $$CONTAINER:/tmp/deisctl.tar.gz package/ ; \
	docker rm -f $$CONTAINER
	@echo "package placed in ./package/deisctl.tar.gz"
