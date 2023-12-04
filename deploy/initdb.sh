#!/bin/bash
psql -d template1 -c 'create extension if not exists "uuid-ossp";'
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d "$POSTGRES_DB"  <<-EOSQL
    CREATE ROLE $APP_DB_USER noinherit login password '$APP_DB_PASSWORD';
    CREATE DATABASE $APP_DB_NAME;
    GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME to "$APP_DB_USER";
EOSQL