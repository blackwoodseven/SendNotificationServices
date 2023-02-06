#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

export CGO_ENABLED=0

TARGETS=(
	./cmd/...
	./pkg/...
)

if [[ ${#@} -ne 0 ]]; then
	TARGETS=("$@")
fi

echo "Running coverage:" "${TARGETS[@]}"

go test -installsuffix "static" -timeout 60s "${TARGETS[@]}" -coverprofile=coverage.out
go tool cover -func=coverage.out
echo "Success!"
