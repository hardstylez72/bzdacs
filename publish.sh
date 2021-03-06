#!/bin/bash

docker build -f Dockerfile -t hardstylez72/bzdacs .
docker login
docker push hardstylez72/bzdacs:latest
