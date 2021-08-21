CREATE TABLE IF NOT EXISTS emploees (
    firstname       varchar(50) not null,
    middlename      varchar(50) not null,
    lastname        varchar(50) not null,
    bdate           date not null,
    addres          varchar(2000) not null,
    department      varchar(100) NOT NULL,
    aboutMe         VARCHAR(3000),
    tnumber         VARCHAR(20) NOT NULL,
    email           VARCHAR(320) NOT NULL
);