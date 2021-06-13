#!/bin/bash
set -e

echo "#################### migrating database"
go run cmd/dbmigrate/main.go
./main
