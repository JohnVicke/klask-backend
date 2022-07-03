#!/usr/bin/env bash

run_tests_in_subdirs () {
  cd $1 
  dirs=$(find . -maxdepth 1 -type d -not -name ".*" -printf '%f\n')
  echo $dirs
  for dir in $dirs; do
    if [ -d $dir ]; then
      cd $dir
      echo "Running tests in $dir"
      go test ./... -v -cover
      cd ..
    fi
  cd $2
  done
}

base_dir=$(pwd)
run_tests_in_subdirs "./cmd" $base_dir
