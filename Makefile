BIN := dist/edifact-ls
GO  := go

.PHONY: build
build:
	$(GO) build -o $(BIN) ./cmd/edifact-ls

.PHONY: test
test:
	$(GO) vet ./...
	$(GO) test ./...

.PHONY: test-e2e
test-e2e: build
	EDIFACT_LS_BIN=$(abspath $(BIN)) ./scripts/e2e.sh

.PHONY: clean
clean:
	rm -f $(BIN)
