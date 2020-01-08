#!/bin/sh

cp .realize.debug.server.yaml .realize.server.yaml
cp .realize.debug.client.yaml .realize.client.yaml

docker-compose up --build -d
