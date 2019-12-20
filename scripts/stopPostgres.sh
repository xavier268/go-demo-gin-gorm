#! /bin/bash
PGC=postgresTestContainer

docker container kill $PGC
docker container rm $PGC
docker container list  -a