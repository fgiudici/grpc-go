all: vet test testrace

build: deps
	go build github.com/fgiudici/grpc-go/...

clean:
	go clean -i github.com/fgiudici/grpc-go/...

deps:
	go get -d -v github.com/fgiudici/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/fgiudici/grpc-go/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/fgiudici/grpc-go/...

testsubmodule: testdeps
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m github.com/fgiudici/grpc-go/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m github.com/fgiudici/grpc-go/security/authorization/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/fgiudici/grpc-go/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/fgiudici/grpc-go/...

testdeps:
	go get -d -v -t github.com/fgiudici/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/fgiudici/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/fgiudici/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/fgiudici/grpc-go/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
