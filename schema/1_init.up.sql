CREATE TABLE IF NOT EXISTS clients (
    id          serial      primary key,
    surname     varchar(20) NOT NULL,
    name        varchar(20) NOT NULL,
    patronymic  varchar(20),
    email       varchar(50) NOT NULL,
    approval    integer     NOT NULL default 0,
    created_at  timestamp   NOT NULL default now(),
    update_at   timestamp   NOT NULL default now()

    CONSTRAINT client_status CHECK (approval in(0,1,-1))
);