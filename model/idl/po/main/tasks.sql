create table tasks
(
    id                varchar(50)              not null
        constraint tasks_pk
            primary key,
    xlid              varchar(50) default ''   not null,
    qbid              varchar(50) default ''   not null,
    unavailable_links TEXT        default '[]' not null,
    available_links   text        default '[]' not null
);

create unique index tasks__idx_qbid
    on tasks (qbid);

create unique index tasks_idx_xlid
    on tasks (xlid);

