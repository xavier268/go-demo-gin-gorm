#! /bin/bash
PGC=postgresTestContainer

docker run --rm -d --name $PGC -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres  -p 5432:5432 postgres:latest
docker container list -a

