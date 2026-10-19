# go-27 entry point in cmd/server
Single main at cmd/server/main.go (NOT at the repo root), with pkg/cart. No third-party deps.
Expected: Octiq discovers cmd/server as the entry point and builds with go build ./cmd/server.
