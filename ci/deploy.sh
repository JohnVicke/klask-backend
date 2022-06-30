#!/usr/bin/env bash

echo "Using shell: $SHELL"

files=("$@")

VERSION=v1.0.0-beta.1

env_vars_array=(
  "version=${VERSION}"
)


workdir=${PWD}
printf -v joined '%s,' "${env_vars_array[@]}"
env_vars=${joined%,}

for file in "${files[@]}"
do
  dir_name=$(basename "$file" .go)
  entry_point="$(grep -oP '(?<=func )\w+' $file | grep -i $dir_name)"
  name=${entry_point,}

  echo "Deploying function $entry_point to gcloud..."
  cd $(dirname $file)

  gcloud functions deploy $name \
    --entry-point $entry_point \
    --set-env-vars $env_vars \
    --runtime go113 \
    --allow-unauthenticated \
    --verbosity debug \
    --trigger-http

  cd $workdir 

done
