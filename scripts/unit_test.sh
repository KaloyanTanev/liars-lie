#!/usr/bin/env bash

PROJECT_HOME="$(cd "$(dirname "${0}")/.." && pwd)"
cd "$PROJECT_HOME"

go test ./...
