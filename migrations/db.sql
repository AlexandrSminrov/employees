CREATE TABLE IF NOT EXISTS emploees (
        id              SERIAL ,
        firstname       varchar(50) not null,
        lastname        varchar(50) not null,
        middlename      varchar(50) not null,
        bdate           date not null,
        addres          varchar(2000) not null,
        department      varchar(100) NOT NULL,
        aboutMe         VARCHAR(3000),
        tnumber         VARCHAR(20) NOT NULL,
        email           VARCHAR(320) NOT NULL
);
INSERT INTO public.emploees (firstname, lastname, middlename, bdate,
                             addres, department, aboutMe, tnumber, email)
    VALUES ('Иванов', 'Иван', 'Иванович', '10.11.1980', 'Москва', 'HR',
            'qwe', '9000000000', 'exaple@ex.ex') RETURNING id;

INSERT INTO public.emploees (firstname, lastname, middlename, bdate,
                             addres, department, aboutMe, tnumber, email)
    VALUES ('Иванов', 'Иван', 'Иванович', '10.11.1999', 'Москва', 'разраб',
            'HR', '1234567890', 'exaple@ex.ex') RETURNING id;