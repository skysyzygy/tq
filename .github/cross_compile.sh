#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
	
platforms=("windows/amd64" "windows/386" "darwin/arm64" "darwin/amd64" "darwin/386" "linux/amd64" "linux/386") 

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package'-'$GOOS'-'$GOARCH
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	
	echo 'Compiling '$output_name' ...'
	env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$output_name .
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done