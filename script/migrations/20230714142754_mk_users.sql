-- +goose Up
-- +goose StatementBegin

create schema if not exists k_user;
comment on schema k_user
    is 'Schema including everything related to users and auth';

create table if not exists k_user.users
(
    id          uuid primary key default gen_random_uuid(),
    nickname    text,
    first_name  text        not null,
    last_name   text,
    email       text unique not null check (email like '%@%.%'),
    created_at  timestamptz      default now(),
    modified_at timestamptz      default now()
);

comment on table k_user.users
    is 'Table including info on users';

create type k_user._user as
(
    nickname   text,
    first_name text,
    last_name  text,
    email      text
);

create or replace function k_user.get_user_by_id(_id uuid)
    returns table
            (
                id           uuid,
                nickname     text,
                first_schema text,
                last_name    text,
                email        text,
                created_at   timestamptz,
                modified_at  timestamptz
            )
as
$function$
begin
    return query
        select u.id,
               u.nickname,
               u.first_name,
               u.last_name,
               u.email,
               u.created_at,
               u.modified_at
        from k_user.users as u
        where u.id = _id;
end
$function$
    language plpgsql volatile;

create or replace function k_user.create_user(_args k_user._user)
    returns table
            (
                id           uuid,
                nickname     text,
                first_schema text,
                last_name    text,
                email        text,
                created_at   timestamptz,
                modified_at  timestamptz
            )
as
$function$
declare
    _id uuid;
begin
    perform from k_user.users as u where u.email = _args.email;
    if FOUND then raise exception 'USER_ALREADY_EXISTS'; end if;

    insert into k_user.users (nickname, first_name, last_name, email)
    values (_args.nickname, _args.first_name, _args.last_name, _args.email)
    returning id into _id;

    return query select * from k_user.get_user_by_id(_id);
end
$function$
    language plpgsql volatile;

-- +goose StatementEnd




-- +goose Down
-- +goose StatementBegin

drop function if exists k_user.create_user(k_user._user);

drop function if exists k_user.get_user_by_id(_id uuid);

drop type if exists  k_user._user;

drop table k_user.users;

drop schema if exists k_user;

-- +goose StatementEnd
