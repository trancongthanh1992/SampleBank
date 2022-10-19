#!/bin/sh
# Alpine just has shell script.
# start.sh support for run migration before run Dockerfile command `CMD [ "/app/main" ]`
# Check exists
set -e

echo "run db migration"

# /app/migration/ is directory of container from Dockerfile.
# Line 11 is run migrate command.
/app/migrate -path /app/migration/ -database "$DB_SOURCE" -verbose up

echo "start the app"
# Take all the parameters passed to the script and run
exec "$@"