# mywebapp (crypto skeleton)

This is a minimal Go web project skeleton with a `pkg/crypto` facade and subpackages:
- `aesx` (GCM preferred, CBC for legacy)
- `rsax` (OAEP encrypt, PSS sign)
- `hashx`, `randx`, `codex` (small helpers)
- OpenSSL interop demo for AES-CBC.

> Change module path in `go.mod` to your own before `go build`.

## Quick start
```bash
cd golang_template
go test ./...
go run ./cmd/server
```
