#!/usr/bin/env bash
# Always run in project root
# ./scripts/tests_all.sh
go test -coverprofile=coverage.out    ./... &&
go tool cover -html=coverage.out

