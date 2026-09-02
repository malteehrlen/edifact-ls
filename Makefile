BIN     := dist/edifact-ls
GO      := go
VERSION ?= dev
LDFLAGS := -X github.com/malteehrlen/edifact-ls/internal/lspserver.Version=$(VERSION)

.PHONY: build
build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BIN) ./cmd/edifact-ls

.PHONY: test
test:
	$(GO) vet ./...
	$(GO) test ./...

.PHONY: test-e2e
test-e2e: build
	EDIFACT_LS_BIN=$(abspath $(BIN)) ./scripts/e2e.sh

.PHONY: install
install:
	$(GO) install -ldflags "$(LDFLAGS)" ./cmd/edifact-ls

.PHONY: docs
docs:
	$(GO) run ./tools/gendocs > docs/SUPPORTED_MESSAGES.md

.PHONY: clean
clean:
	rm -f $(BIN)
