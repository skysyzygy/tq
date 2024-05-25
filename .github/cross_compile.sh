#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
	
platforms=("windows/amd64" "windows/386" "darwin/arm64" "darwin/amd64" "linux/amd64" "linux/386") 

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	case $GOARCH in 
		386) arch="x86";;
		amd64) arch="x86_64";;
		arm64) arch="M1";;
	esac
	output_name=$package'-'$GOOS'-'$arch
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
