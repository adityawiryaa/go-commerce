#!/bin/sh
export $(cat ./environment/$1/.env | xargs)

go run services/migrations/mysql/main.go migrations:up
#go run services/migrations/main.go migrations:status
#go run services/migrations/mysql/main.go migrations:create create_product_table --table=products