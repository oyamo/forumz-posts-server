create table if not exists post_media (
    id UUID not null primary key,
    post_id UUID not null,
    mime_type varchar(64),
    file_name varchar(64),
    datetime_created timestamp not null default current_timestamp,
    last_modified timestamp with time zone  not null default current_timestamp,
    constraint foreign key (post_id) references post(id)
);