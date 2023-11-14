create table users
(
    id           uuid default gen_random_uuid() primary key not null,
    username     text                                       not null,
    password     text default ''                            not null
);