create table if not exists post (
     id UUID not null primary key,
     person_id UUID not null,
     content varchar(512) not null,
     datetime_created timestamp not null default current_timestamp,
     last_modified timestamp with time zone  not null default current_timestamp,
     constraint fk_person_posts foreign key (person_id) references person(id)
);