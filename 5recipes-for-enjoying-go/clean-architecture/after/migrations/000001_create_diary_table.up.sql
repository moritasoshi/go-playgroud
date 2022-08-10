create table if not exists diary (
  id serial primary key,
  title varchar(128) not null,
  description text not null,
  created_at timestamp default current_timestamp
);
