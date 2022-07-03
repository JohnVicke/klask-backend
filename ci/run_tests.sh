#!/usr/bin/env bash

run_tests_in_subdirs () {
  cd $1 
  dirs=$(find . -maxdepth 1 -type d -not -name ".*" -printf '%f\n')
  for dir in $dirs; do
    if [ -d $dir ]; then
      cd $dir
      go test ./... -v -cover
      return_code=$?
      if [ $return_code -ne 0 ]; then
        echo "Tests in $dir failed"
        exit $return_code
      fi
      cd ..
      printf "\n"
    fi
  done
  cd $2
}

base_dir=$(pwd)
run_tests_in_subdirs "./cmd" $base_dir
