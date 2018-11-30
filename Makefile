export GO111MODULE=on
BINARY_NAME=triangleClassifier

all: clean build

install:
	go install
build:
	go build -v
clean:
	go clean
	go mod tidy
	rm -f $(BINARY_NAME)
test:
	go test -v ./cmd/classifier