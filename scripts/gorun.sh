#!/bin/bash

echo "Go Run script"

if [[ -f "main.go" ]]
then
  googleCredentialsPath="./internal/adminSdkSecret-Dev.json"

  firebaseReplace="github.com/JohnVicke/pkg/firebase=../../pkg/firebase"
  replaces=($firebaseReplace)

  export GOOGLE_APPLICATION_CREDENTIALS=$googleCredentialsPath
  cd cmd
  for d in */
  do
    echo "Adding replace to $d"
    cd $d
    # for each value in replaces
    for r in "${replaces[@]}"
    do
      echo "Replacing $r"
      go mod edit -replace $r
    done
    cd ..
  done

  cd ..

  echo "Running main.go"
  go run main.go

else
  echo "Make sure you are in the root before running!"
fi 
