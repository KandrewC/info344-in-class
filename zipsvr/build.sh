#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t andrewk7/zipsvr .
docker push andrewk7/zipsvr
go clean
