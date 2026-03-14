create table tasks
(
    id                varchar(50) not null
        constraint tasks_pk
            primary key,
    xlid              varchar(50),
    qbid              varchar(50),
    unavailable_links TEXT default '[]',
    available_links   TEXT default '[]'
);

