CREATE TABLE wallets
(
    address varchar(255) not null,
    amount decimal not null
);

CREATE TABLE transactions
(
    wallet_from varchar(255) not null,
    wallet_to varchar(255) not null,
    amount decimal not null
);