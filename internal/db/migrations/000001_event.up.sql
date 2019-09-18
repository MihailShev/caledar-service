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