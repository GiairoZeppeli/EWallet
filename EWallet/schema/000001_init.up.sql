CREATE TABLE wallets
(
    id serial not null unique,
    address varchar(255) not null,
    amount decimal not null
);

CREATE TABLE transactions
(
    id serial not null unique,
    wallet_from varchar(255) not null,
    wallet_to varchar(255) not null,
    amount decimal not null
);