version_file=$1
version=$(echo "$(head -n 1 $version_file)" | awk -F':' '{print $1}' | egrep -o "v.*")
echo $version
