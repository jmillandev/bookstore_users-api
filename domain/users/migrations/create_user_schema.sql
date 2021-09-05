CREATE TABLE users(
    id INT UNSIGNED AUTO_INCREMENT,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    status VARCHAR(10) NOT NULL,
    password VARCHAR(250) NOT NULL,
    date_created VARCHAR(45) NOT NULL,
    PRIMARY KEY(id)
);