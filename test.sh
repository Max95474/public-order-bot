#!/usr/bin/env bash
# Load environment variables from .env file
export $(cat .env | grep -v ^# | xargs)

go test test/*.go
