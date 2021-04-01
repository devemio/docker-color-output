#!/usr/bin/env bash

DOCKER_COLOR_OUTPUT_PATH=/absolute/path/to/dco

di() {
  docker images $@ | $DOCKER_COLOR_OUTPUT_PATH
}

dps() {
  docker ps $@ | $DOCKER_COLOR_OUTPUT_PATH
}

dcps() {
  docker-compose ps $@ | $DOCKER_COLOR_OUTPUT_PATH
}