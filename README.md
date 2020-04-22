# Prometheus Proof of Concept + more!

This Dockerized Prometheus proof-of-concept project spins up:
* A Prometheus server
* A clustered Alertmanager (three containers) 
* A helper service (to receive alerts and echoing them to `stdout`)
* An Nginx server, serving a default page and the `stub_status` page
* An nginx-exporter, exports nginx metrics by parsing the `stub_status`
* A blackbox-exporter that probes URLs and exports metrics
* Grafana server

The Alertmanager cluster is fully meshed, every Alertmanager are configured with itself and all of the other Alertmanagers as cluster peers. To check the cluster status browse to: http://localhost:9001 (any active Alertmanger will be fine)

Currently there's one alarm configured that fires as long as the active connections to the Nginx container is above 2, it's easy enough to trigger (basically just to send a couple of requests to it).

The script `send_alert.sh` sends a test alert to each one of the alertmanagers, you should see `stdout` messages from the `echo-service` container if tailing the logs.

Check out the `config/` directory for all configuration files used in this proof-of-concept. 

## How-to:

1. Run the `init.sh` bash script to fix permissions on directories used by the different components of this PoC.
2. Run `docker-compose` to start the environment: `docker-compose up -d`

## Exposed services

* Prometheus: http://localhost:9090
* Alertmanager 1: http://localhost:9001
* Alertmanager 2: http://localhost:9002
* Alertmanager 3: http://localhost:9003
* Nginx: http://localhost:81
* Grafana: http://localhost:3000

## Todo

* Un-uglify the random permission changing script..
* Restructure this repo a bit (like the data directories)