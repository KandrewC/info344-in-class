#1/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t andrewk7/testserver .
docker push andrewk7/testserver
