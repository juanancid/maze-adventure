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

.PHONY: fresh-run
fresh-run: clean $(TARGET)
	./$(TARGET)

.PHONY: clean
clean:
	rm -f $(TARGET)

$(TARGET): $(SOURCES)
	go build -o $@ ./cmd/main
