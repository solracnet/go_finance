create table users (
    id serial not null primary key,
    username varchar not null,
    password varchar not null,
    email varchar unique not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now())
);

create table categories (
    id serial not null primary key,
    user_id int not null,
    title varchar not null,
    type varchar not null,
    description varchar not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now())
);

alter table categories add foreign key (user_id) references users (id);

create table accounts (
    id serial not null primary key,
    user_id int not null,
    category_id int not null,
    title varchar not null,
    type varchar not null,
    description varchar not null,
    value integer not null,
    "date" date not null,
    created_at timestamp not null default (now()),
    updated_at timestamp not null default (now())
);

alter table accounts add foreign key (user_id) references users (id);
alter table accounts add foreign key (category_id) references categories (id);