#!/usr/bin/env bash

DOCKER_COLOR_OUTPUT_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" &> /dev/null && pwd)/bin/dco-darwin-amd64"

di() {
  docker images $@ | $DOCKER_COLOR_OUTPUT_PATH
}

dps() {
  docker ps $@ | $DOCKER_COLOR_OUTPUT_PATH
}

dcps() {
  docker-compose ps $@ | $DOCKER_COLOR_OUTPUT_PATH
}