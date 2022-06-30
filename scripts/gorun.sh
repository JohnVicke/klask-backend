#!/bin/bash

echo "Go Run script"

if [[ -f "main.go" ]]
then
  firebaseReplace="github.com/JohnVicke/pkg/firebase=../../pkg/firebase"
  replaces=($firebaseReplace)

  cd cmd

  for d in */
  do
    cd $d
    # for each value in replaces
    for r in "${replaces[@]}"
    do
      echo "Replacing $r in $d"
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
