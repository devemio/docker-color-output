#!/usr/bin/env bash

# Usage:
#   alias docker='bash path/to/docker.sh docker'
if [[ "$1" == "docker" ]]; then
  if [[ "$2" == "ps" || "$2" == "images" ]]; then
    "$@" | docker-color-output
  elif [[ "$2" == "compose" && "$3" == "ps" ]]; then
    "$@" | docker-color-output
  else
    "$@"
  fi
else
  "$@"
fi
