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

CREATE TABLE wallets_transactions
(
    id serial not null unique,
    wallets_id int references wallets(id) on delete cascade not null,
    transactions_id int references transactions(id) on delete cascade not null
);


INSERT INTO wallets (address, amount)
VALUES
(MD5(random()::text), 100);
