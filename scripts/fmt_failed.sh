#!/bin/sh

if ! gofmt -d .; then
    echo "format failed"
    exit 1
fi

if ! go vet ./...; then
    echo "format failed"
    exit 1
fi