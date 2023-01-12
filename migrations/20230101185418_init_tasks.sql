-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table if not exists public.task
(
    id         bigserial primary key,
    task_type  int8                     not null references public.task_types (id),
    status     int8                     not null,
    created_at timestamp with time zone not null default now()
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table if exists public.task cascade