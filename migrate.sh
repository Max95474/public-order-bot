#!/usr/bin/env bash

source .env
go run db/migrations/*.go up
