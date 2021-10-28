CREATE TABLE IF NOT EXISTS emploees (
        id              SERIAL ,
        firstname       varchar(50) not null,
        lastname        varchar(50) not null,
        middlename      varchar(50) not null,
        date_of_birth           date not null,
        addres          varchar(2000) not null,
        department      varchar(100) NOT NULL,
        about_me         VARCHAR(3000),
        phone         VARCHAR(20) NOT NULL,
        email           VARCHAR(320) NOT NULL
);
INSERT INTO public.emploees (firstname, lastname, middlename, date_of_birth,
                             addres, department, about_me, phone, email)
    VALUES ('Иванов', 'Иван', 'Иванович', '10.11.1980', 'Москва', 'HR',
            'qwe', '9000000000', 'exaple@ex.ex') RETURNING id;

INSERT INTO public.emploees (firstname, lastname, middlename, date_of_birth,
                             addres, department, about_me, phone, email)
    VALUES ('Иванов', 'Иван', 'Иванович', '10.11.1999', 'Москва', 'разраб',
            'HR', '1234567890', 'exaple@ex.ex') RETURNING id;