#!/usr/bin/env bash

files=("$@")

# loop over files
for file in "${files[@]}"
do
  echo "Verifying version $file"
  # deploy file
  # ...
done
