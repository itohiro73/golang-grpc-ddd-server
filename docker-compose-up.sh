#!/bin/sh

cp .realize.normal.server.yaml .realize.server.yaml
cp .realize.normal.client.yaml .realize.client.yaml

docker-compose up --build -d
