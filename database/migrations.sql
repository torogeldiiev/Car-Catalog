CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50)
    );

CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(20) NOT NULL,
    mark VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INTEGER NOT NULL,
    owner_id INTEGER,
    FOREIGN KEY (owner_id) REFERENCES people(id)
    );

