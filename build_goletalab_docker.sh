#!/usr/bin/env bash
set -euo pipefail
docker build -f goletalab.Dockerfile -t mapstructure-task .
docker run --rm mapstructure-task "$@"
