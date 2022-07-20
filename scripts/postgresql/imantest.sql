create table if not exists posts
(
    id int not null primary key,
    user_id int default null,
    title varchar default '',
    body varchar default ''
);