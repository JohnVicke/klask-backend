#!/usr/bin/env bash

# verify input arguments
current_dir=$(pwd)
command_name=$1

if [ $# -lt 1 ]; then
  echo "Usage: $0 <commandname>"
  exit 1
fi

if [[ $command_name == *['!'@#\$%^\&*()_+]* ]]
then
  echo "$command_name is not a valid command name"
  exit 1
fi

if [[ $command_name =~ ^[a-z] ]]
then
  echo "$command_name is not a valid command name, must start with a uppercase letter"
  exit 
fi


dir_name=${command_name,,}

if [ -d "./cmd/$dir_name" ]; then
  echo "Command $command_name already exists"
  exit 1
fi


echo "Boostrapping command $command_name..."
cd ./cmd && mkdir $dir_name && cd $dir_name

echo "- v1.0.0-beta.1: Initializing $command_name function" > .version

go mod init github.com/JohnVicke/klask-backend/cmd/$dir_name

echo "package $dir_name" >> $dir_name.go

cp $dir_name.go "$dir_name"_test.go

echo "

import (
	\"log\"
	\"net/http\"
)

func init() {
	log.Println(\"Healthcheck initialized\")
}

// $command_name is a function that returns a string.
func $command_name(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(\"Hello from $command_name\"))
	w.WriteHeader(http.StatusOK)
}" >> $dir_name.go


echo "
import (
	\"net/http\"
	\"net/http/httptest\"
	\"testing\"

	\"github.com/stretchr/testify/assert\"
)

func Test$command_name(t *testing.T) {
	req, err:= http.NewRequest(\"GET\", \"/$dir_name\", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc($command_name)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
" >> "$dir_name"_test.go

go mod tidy

echo "Finished bootstrapping command $command_name, cd into ./cmd/$dir_name to start coding"