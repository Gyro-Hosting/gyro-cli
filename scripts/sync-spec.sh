#!/bin/sh
# Copy the Gyro Public API spec from the API module and regenerate the command table.
#   scripts/sync-spec.sh [path/to/openapi.yaml]
set -eu
cd "$(dirname "$0")/.."
SRC=${1:-../gyro-api/openapi.yaml}
cp -p "$SRC" internal/spec/openapi.yaml
go run ./cmd/gyro-gen -spec internal/spec/openapi.yaml -out internal/spec/zz_generated_ops.go
go build ./... && go test ./internal/spec/ >/dev/null && echo "spec synced: $(grep -m1 '^  version:' internal/spec/openapi.yaml)"
