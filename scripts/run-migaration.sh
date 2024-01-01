#!/usr/bin/env sh
set -e

>&2 echo "Running migration ..."
#migrate -path=./migrations -database=postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:5432/$POSTGRES_DB?sslmode=disable up
  migrate -path=./migrations -database=postgres://user:password@localhost:5432/test-db?sslmode=disable up
# migrate -path=./migrations -database=postgres://abbos:root@localhost:5432/loan_service?sslmode=disable up 1 

#migrate -path=./migrations -database=cassandra://$CASSANDRA_HOST:$CASSANDRA_PORT/$CASSANDRA_DB?sslmode=disable up
tail -f /dev/null