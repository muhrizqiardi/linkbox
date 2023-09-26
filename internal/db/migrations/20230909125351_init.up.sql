create table if not exists users (
    id serial primary key,
    username text unique not null,
    password text not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table if not exists folders (
    id serial primary key,
    unique_name text not null,
    user_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    foreign key (user_id)
        references users(id),

    unique (user_id, unique_name)
);

create table if not exists tags (
    id serial primary key,
    unique_name text unique not null,
    user_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    foreign key (user_id)
        references users(id),

    unique (user_id, unique_name)
);

create table if not exists links (
    id serial primary key,
    url text unique not null,
    title text not null,
    description text not null,
    user_id int not null,
    folder_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    foreign key (user_id)
        references users(id),
    foreign key (folder_id)
        references folders(id)
);

create table if not exists link_tags (
    id serial primary key,
    link_id int not null,
    tag_id int not null,
    user_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,

    foreign key (user_id)
        references users(id),

    unique (link_id, tag_id, user_id)
);

