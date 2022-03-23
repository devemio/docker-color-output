#!/usr/bin/env bash

di() {
  docker images $@ | docker-color-output
}

dps() {
  docker ps $@ | docker-color-output
}

dcps() {
  docker compose ps $@ | docker-color-output
}
