CREATE TABLE IF NOT EXIST users (
    nik VARCHAR(16) PRIMARY KEY,
    name VARCHAR,
    phone_num VARCHAR(12)
);

CREATE TABLE IF NOT EXIST accounts (
    account_num VARCHAR(10) PRIMARY KEY,
    nik VARCHAR(16) REFERENCES users (nik),
    balance DECIMAL(12, 2)
);