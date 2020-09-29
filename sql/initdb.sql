CREATE DATABASE IF NOT EXISTS pismo;
USE pismo;

CREATE TABLE accounts(
    id int auto_increment primary key,
    document_number varchar(11) not null
);

CREATE TABLE transactions(
    id int auto_increment primary key,
    account_id int not null, FOREIGN KEY(account_id) REFERENCES accounts(id) ON DELETE CASCADE,
    operation_type_id int,
    amount float,
    event_date timestamp
);
