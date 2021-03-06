#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/backend/backend ./cmd/backend
#npm run --prefix ./ui build
#docker build -f Dockerfile -t hardstylez72/bzdacs .
