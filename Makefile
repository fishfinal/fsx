GO = go
PROJECT_DIR = .

test:
	$(GO) test -race $(PROJECT_DIR)/...

test-verbose:
	$(GO) test -race -v $(PROJECT_DIR)/...

.PHONY: test test-verbose
