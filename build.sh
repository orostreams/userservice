#!/bin/sh

MODE=$1

if [ "${MODE}" == "dev" ]
then
    go run ./cmd/user-service.go
elif [ "${MODE}" == "prod" ]
then
    docker build -t orostreams/userservice:latest .
else
    echo Please Provide a build mode dev or prod
fi
