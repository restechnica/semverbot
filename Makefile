# make sure targets do not conflict with file and folder names
.PHONY: build clean test

BIN = bin
GOBUILD = go build

# build the project
build:
	$(GOBUILD) -o $(BIN)/sbot
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BIN)/sbot-windows-amd64.exe
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN)/sbot-linux-amd64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BIN)/sbot-darwin-amd64

# run quality assessment checks
check:
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
	./bin/sbot

# run tests
test:
	mkdir -p ./out
	go test ./... -cover -v -coverprofile ./out/coverage.txt
	uncover ./out/coverage.txt
