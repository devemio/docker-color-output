#!/usr/bin/env bash

# Usage:
#  alias docker='bash decorator.sh docker'
#  alias docker-compose='bash decorator.sh docker-compose'
if [[ "$1" == "docker" ]]; then
  if [[ "$2" == "ps" || "$2" == "images" ]]; then
    "$@" | docker-color-output
  elif [[ "$2" == "compose" && "$3" == "ps" ]]; then
    "$@" | docker-color-output
  else
    "$@"
  fi
elif [[ "$1" == "docker-compose" && "$2" == "ps" ]]; then
  "$@" | docker-color-output
else
  "$@"
fi
