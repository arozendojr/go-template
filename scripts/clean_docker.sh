#!/usr/bin/env bash

# Clear and delete container not run
docker system prune --all --force --volumes

# Create kubernets local
kind create cluster
