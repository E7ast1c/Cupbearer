-- +goose up
-- sql in this section is executed when the migration is applied.
create table if not exists public.birthdays
(
    id            bigserial primary key,
    -- user who create row
    user_id       int                      not null,
    person_name   text unique              not null,
    remind_at     timestamptz              not null,
--     reminder_type timestamp with time zone not null,
    payload       text                     not null,
    birthday_date timestamp with time zone not null
);

-- +goose down
-- sql in this section is executed when the migration is rolled back.
drop table if exists public.birthdays