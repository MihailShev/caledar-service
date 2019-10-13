#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER mshev WITH PASSWORD '123';
    ALTER ROLE mshev WITH LOGIN;
    CREATE DATABASE calendar OWNER mshev encoding 'UTF8' lc_collate 'ru_RU.UTF-8' LC_CTYPE 'ru_RU.UTF-8' template template0;
EOSQL

psql --username mshev --dbname calendar <<-EOSQL
  CREATE TABLE IF NOT EXISTS "event"
(
    "uuid"        serial primary key,
    "user_id"     serial,
    "title"       text                     not null,
    "description" text,
    "start"       timestamp with time zone not null,
    "end"         timestamp with time zone not null,
    "notice_time"  timestamp with time zone
);
EOSQL