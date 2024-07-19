CREATE TABLE IF NOT EXISTS clients (
    id          serial primary key,
    surname     varchar(20) NOT NULL,
    name        varchar(20) NOT NULL,
    patronymic  varchar(20),
    birth_date  date        NOT NULL,
    status      varchar(10) NOT NULL default 'processed'

    CONSTRAINT client_status CHECK (status in('processed','accepted','rejected'))
);