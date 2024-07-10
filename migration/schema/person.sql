CREATE TABLE IF NOT EXISTS person (
  id UUID not null primary key,
  first_name varchar(32) not null,
  status varchar(12) not null default 'Active',
  datetime_created timestamp not null default current_timestamp,
  last_modified timestamp with time zone  not null default current_timestamp
);
