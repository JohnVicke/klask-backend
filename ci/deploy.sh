#!/usr/bin/env bash

echo "Using shell: $SHELL"

files=("$@")

VERSION=v1.0.0-beta.1

env_vars_array=(
  "VAR1=${VAR1}" 
  "name=${VERSION}"
)

# TODO: Filter for main .go file

printf -v joined '%s,' "${env_vars_array[@]}"
env_vars=${joined%,}

for file in "${files[@]}"
do
  entry_point="$(grep -oP '(?<=func )\w+' $file)"
  name=${entry_point,}

  echo "Deploying function $entry_point to gcloud..."

  echo "entry_point: $entry_point"
  echo "name: $name"
  echo "env_vars: $env_vars"

  gcloud functions deploy $name \
    --entry-point $entry_point \
    --set-env-vars $env_vars \
    --runtime go113 \
    --allow-unauthenticated \
    --verbosity debug \
    --trigger-http

done
