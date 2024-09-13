# make sure targets do not conflict with file and folder names
.PHONY: build clean test

# build the project
build: build-prerelease

build-prerelease:
	@nu main.nu build

build-all:
	@nu main.nu build --all --version "$$(sbot predict version --debug)"

# run quality assessment checks
check:
	@echo "Running go version check ..."
	@go version | grep -q 'go1.21' || (echo "Go version 1.21 required. Your version is '$$(go version)'" && exit 1)
	@echo "Ok!"

	@echo "Running gofmt ..."
	@! gofmt -s -d -l . 2>&1 | grep -vE '^\.git/'
	@echo "Ok!"

	@echo "Running go vet ..."
	@go vet ./...
	@echo "Ok!"

	@echo "Running goimports ..."
	@! goimports -l . | grep -vF 'No Exceptions'
	@echo "Ok!"

# clean
clean:
	rm -rf bin out

# format
format:
	go fmt ./...
	goimports -w .

# get all dependencies
provision:
	@echo "Getting dependencies ..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/gregoryv/uncover/cmd/uncover@latest
	@go mod download
	@echo "Done!"

# run the binary
run:
	./bin/sbot $(arg)

# run tests
test:
	mkdir -p ./out
	go test ./... -cover -v -coverprofile ./out/coverage.txt
	uncover ./out/coverage.txt
