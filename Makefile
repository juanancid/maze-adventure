SOURCES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
TARGET=maze-adventure

.PHONY: all
all: $(TARGET)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race ./...

.PHONY: validate
validate: lint test

.PHONY: run
run: validate $(TARGET)
	./$(TARGET)

$(TARGET): $(SOURCES)
	go build -o $@ ./cmd/main
