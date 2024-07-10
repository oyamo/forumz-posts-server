create table if not exists posts (
     id UUID not null primary key,
     first_name varchar(32) not null,
     person_id UUID not null,
     content varchar(512) not null,
     datetime_created timestamp not null default current_timestamp,
     last_modified timestamp with time zone  not null default current_timestamp,
     constraint foreign key (person_id) references person(id)
);