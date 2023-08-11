#!/bin/bash

echo "Creating directories and fixing permissions.."

mkdir -p ./am1/
mkdir -p ./am2/
mkdir -p ./am3/
mkdir -p ./prometheus/

sudo chown -R 65534 ./am*/
sudo chown -R 65534 ./prometheus/

