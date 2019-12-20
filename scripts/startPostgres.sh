#! /bin/bash
PGC=postgresTestContainer

docker run --name $PGC -e POSTGRES_PASSWORD=secret -d postgres:latest
docker container list -a

