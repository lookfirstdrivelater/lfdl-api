#!/usr/bin/env bash

cd ..

rm -rf tempapi

mkdir tempapi
mkdir -p tempapi/opt/services/lookfirstdrivelater/
mkdir -p tempapi/etc/systemd/system


cd cmd/lookfirstdrivelater/

if go build -o lookfirstdrivelater
then
    echo "build ok"
else
    exit 1
fi

cd ..
cd ..

cp ./init/lfdlapi.service tempapi/etc/systemd/system

cp  cmd/lookfirstdrivelater/lookfirstdrivelater tempapi/opt/services/lookfirstdrivelater/
cp  cmd/lookfirstdrivelater/.env tempapi/opt/services/lookfirstdrivelater/

cd cmd/lookfirstdrivelater

VERSION=$(./lookfirstdrivelater -version)
NAME=lfdlapi

cd ..
cd ..


fpm -s dir -t deb -v $VERSION -n $NAME --after-install scripts/sysd-post-exec.sh -C tempapi .

rm -rf tempapi