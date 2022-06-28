#!/usr/bin/env bash

files=("$@")

# loop over files
for file in "${files[@]}"
do
  echo "Verifying version $file"
  grep -oP '(?<=func)\w+' $file
  # deploy file
  # ...
done
