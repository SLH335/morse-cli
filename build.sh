#!/usr/bin/env bash

platforms=("linux/amd64" "linux/arm64" "windows/amd64")

build() {
    GOOS=$1
    GOARCH=$2

    # check os and set variables accordingly
    if [ $GOOS = windows ]
    then
        filename="$appname.exe"
        archiver="zip"
        archiveext="zip"
    else
        filename="$appname"
        archiver="tar cfz"
        archiveext="tar.gz"
    fi

    echo "Building $GOOS-$GOARCH binary"

    export GOOS
    export GOARCH
    go build -ldflags "$ldflags" -o "$buildpath/$filename"

    # check if build completed successfully
    if [ $? -ne 0 ]
    then
        echo "Build failed for $GOOS-$GOARCH"
        exit 1
    fi

    (cd $buildpath && $archiver "./${appname}_${RELEASE_VERSION}_$GOOS-$GOARCH.$archiveext" "$filename" > /dev/null && rm $filename)
}

main() {
    # check if release version is specified
    if [ -z "$1" ]
    then
        echo "usage: $0 <version>"
        exit 1
    fi

    RELEASE_VERSION="$1"

    # check if release version follows semantic versioning (e.g. 1.2.3)
    if [[ ! "$1" =~ [0-9]+\.[0-9]+\.[0-9]+ ]]
    then
        echo "error: version number must follow semantic versioning (e.g. 1.2.3)"
        exit 1
    fi


    # check if required programs are installed
    required=("go" "tar" "zip")
    for program in "${required[@]}"
    do
        if ! command -v $program &> /dev/null
        then
            echo "error: $program is not installed"
            exit 1
        fi
    done

    # check if required Go version is installed
    if ! go version | grep -Eq 'go1\.(2[1-9]|[3-9][0-9])'
    then
        echo "error: Go version must be 1.21 or above"
        exit 1
    fi


    appname=morse
    buildpath=./build
    ldflags="-X main.AppVersion=$RELEASE_VERSION"

    # build for specified configurations
    build linux amd64
    build windows amd64
}

main "$@"
