#!/usr/bin/env bash

files=("$@")

env_vars_array=(
  "VAR1=${VAR1}" 
  "VAR2=${VAR1}"
)

printf -v joined '%s,' "${env_vars_array[@]}"
env_vars=${joined%,}

for file in "${files[@]}"
do
  echo "Verifying version $file"

  entry_point="$(grep -oP '(?<=func )\w+' $file)"
  name=${func_name,}

#  gcloud functions deploy $name \
#    --entry-point $entry_point \
#    --set-env-vars $env_vars \
#    --runtime go113 \
#    --allow-unauthenticated \
#    --trigger-http

done
