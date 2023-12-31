#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$DATABASE_NAME" <<-EOSQL
    GRANT USAGE, CREATE ON SCHEMA $SCHEMA_NAME TO $DATABASE_USER;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA $SCHEMA_NAME TO $DATABASE_USER;
EOSQL
