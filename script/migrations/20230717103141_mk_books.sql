-- +goose Up
-- +goose StatementBegin

create schema if not exists library;
comment on schema library is 'Includes everything about books and their authors';

-- TABLES --

create table if not exists library.books
(
    id           uuid primary key default gen_random_uuid(),
    title        text                              not null,
    release_year int2,
    created_at   timestamptz      default now()    not null,
    created_by   uuid references k_user.users (id) not null,
    modified_at  timestamptz      default now()    not null,
    modified_by  uuid references k_user.users (id) not null
);

comment on table library.books
    is 'Table including info on books';

create table if not exists library.authors
(
    id          uuid primary key default gen_random_uuid(),
    name        text                              not null,
    last_name   text                              not null,
    father_name text,
    birthday    date,
    created_at  timestamptz      default now(),
    created_by  uuid references k_user.users (id) not null,
    modified_at timestamptz      default now(),
    modified_by uuid references k_user.users (id) not null
);

comment on table library.authors
    is 'Table including info on authors';

create table if not exists library.authorship
(
    id         uuid primary key default gen_random_uuid(),
    book_id    uuid references library.books (id)   not null,
    author_id  uuid references library.authors (id) not null,
    created_at timestamptz      default now(),
    created_by uuid references k_user.users (id)    not null
);

comment on table library.authorship
    is 'Table including relations between books and authors';

-- TABLES --

-- API --

create type library.book as
(
    title        text,
    release_year int8
);

create type library.author as
(
    name        text,
    last_name   text,
    father_name text,
    birthday    date
);

create or replace function library.get_book_by_id(_id uuid)
    returns table
            (
                book_id      uuid,
                title        text,
                release_year int2,
                author_id    uuid,
                name         text,
                last_name    text,
                father_name  text,
                created_at   timestamptz,
                created_by   uuid,
                modified_at  timestamptz,
                modified_by  uuid
            )
as
$function$
begin
    return query
        select b.id           as book_id,
               b.title        as title,
               b.release_year as release_year,
               a.id           as author_id,
               a.name         as name,
               a.last_name    as last_name,
               a.father_name  as fathr_name,
               b.created_at   as created_at,
               b.created_by   as created_by,
               b.modified_at  as modified_at,
               b.modified_by  as modified_by
        from library.books b
                 join library.authorship sh
                      on b.id = sh.book_id
                 join library.authors a
                      on a.id = sh.author_id
        where b.id = _id;
end
$function$
    language plpgsql volatile;

create or replace function library.get_author_by_id(_id uuid)
    returns table
            (
                author_id   uuid,
                name        text,
                last_name   text,
                father_name text,
                birthday    date,
                created_at  timestamptz,
                created_by  uuid,
                modified_at timestamptz,
                modified_by uuid
            )
as
$function$
begin
    return query
        select a.id          as author_id,
               a.name        as name,
               a.last_name   as last_name,
               a.father_name as father_name,
               a.birthday    as birthday,
               a.created_at  as created_at,
               a.created_by  as created_by,
               a.modified_at as modified_at,
               a.modified_by as modified_by
        from library.authors a
        where a.id = _id;
end
$function$
    language plpgsql volatile;

create or replace function library.create_book(_user_id uuid, _author_id uuid, _book library.book)
    returns table
            (
                book_id      uuid,
                title        text,
                release_year int2,
                author_id    uuid,
                name         text,
                last_name    text,
                father_name  text,
                created_at   timestamptz,
                created_by   uuid,
                modified_at  timestamptz,
                modified_by  uuid
            )
as
$function$
declare
    id_ uuid;
begin
    perform from library.authors a where a.id = _author_id;
    if not FOUND then raise 'AUTHOR_NOT_FOUND'; end if;

    insert into library.books (title, release_year, created_by, modified_by)
    values (_book.title, _book.release_year, _user_id, _user_id)
    returning id into id_;

    insert into library.authorship (book_id, author_id, created_by)
    values (id_, _author_id, _user_id);

    return query
        select * from library.get_book_by_id(id_);
end
$function$
    language plpgsql volatile;

create or replace function library.create_author(_user_id uuid, _author library.author)
    returns table
            (
                author_id   uuid,
                name        text,
                last_name   text,
                father_name text,
                birthday    date,
                created_at  timestamptz,
                created_by  uuid,
                modified_at timestamptz,
                modified_by uuid
            )
as
$function$
declare
    author_id_ uuid;
begin
    perform
    from library.authors a
    where a.name = _author.name
      and a.last_name = _author.last_name
      and a.father_name = _author.father_name
      and a.birthday is not distinct from _author.birthday;
    if FOUND then raise 'AUTHOR_ALREADY_EXISTS'; end if;

    insert into library.authors
        (name, last_name, father_name, birthday, created_by, modified_by)
    values (_author.name,
            _author.last_name,
            _author.birthday,
            _user_id,
            _user_id)
    returning id into author_id_;

    return query
        select * from library.get_author_by_id(author_id_);
end
$function$
    language plpgsql volatile;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop function if exists library.create_author(_user_id uuid, _author library.author);

drop function if exists library.create_book(uuid, uuid, library.book);

drop function if exists library.get_author_by_id(uuid);

drop function if exists library.get_book_by_id(uuid);

drop type if exists library.author;

drop type if exists library.book;

drop table if exists library.authorship;

drop table if exists library.authors;

drop table if exists library.books;

-- +goose StatementEnd
