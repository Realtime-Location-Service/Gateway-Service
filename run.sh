#!/usr/bin/env bash

docker-compose down
docker build -t rls-gateway .
docker-compose rm -f
docker-compose up
