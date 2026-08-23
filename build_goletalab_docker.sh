#!/bin/bash
set -euo pipefail
DOCKER_BUILDKIT=1 docker build --platform linux/amd64 -f goletalab.Dockerfile -t "${1:-mapstructure-task}" .
