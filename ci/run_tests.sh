#!/usr/bin/env bash

run_tests_in_subdirs () {
  cd $1 
  dirs=$(find . -maxdepth 1 -type d -not -name ".*" -printf '%f\n')
  for dir in $dirs; do
    if [ -d $dir ]; then
      cd $dir
      go test ./... -v -cover
      if [ $? -ne 0 ]; then
        return 1
      fi
      cd ..
      printf "\n"
    fi
  done
  cd $2
  return 0
}

base_dir=$(pwd)
run_tests_in_subdirs "./cmd" $base_dir || exit 1
