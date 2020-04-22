#!/bin/bash

echo "Creating directories and fixing permissions.."

mkdir -p ./am1/
mkdir -p ./am2/
mkdir -p ./am3/
mkdir -p ./prometheus/
mkdir -p ./grafana/

sudo chown 65534 ./am1/
sudo chown 65534 ./am2/
sudo chown 65534 ./am3/
sudo chown 65534 ./prometheus/
sudo chown 472:472 ./grafana/

