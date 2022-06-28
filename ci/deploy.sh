#!/usr/bin/env bash

echo "$@"

args=("$@")

echo Number of arguments: $#
echo 1st argument: ${args[0]}
echo 2nd argument: ${args[1]}