# Variables

GO_CMD=go
API_PATH=cmd/api/main.go
WORKER_PATH=cmd/worker/main.go
DEBUG_PATH=cmd/debug/debug.go

# Phony targets

.PHONY: run-api run-worker

# Run targets

run-api:
	$(GO_CMD) run $(API_PATH)

run-worker:
	$(GO_CMD) run $(WORKER_PATH)

run-debug:
	$(GO_CMD) run $(DEBUG_PATH)