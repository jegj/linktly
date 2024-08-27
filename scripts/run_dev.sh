#!/bin/bash
ENV_FILE="configs/.env"

set -o allexport
source $ENV_FILE
set +o allexport

go run cmd/main.go
