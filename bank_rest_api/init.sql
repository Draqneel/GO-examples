CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    login TEXT,
    hash_password TEXT,
    full_name TEXT,
    phone_number TEXT,
    email TEXT
);

CREATE TABLE client
(
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    PRIMARY KEY(user_id),
    address TEXT,
    digital_signature TEXT
);

CREATE TABLE branch
(
    branch_id SERIAL PRIMARY KEY ,
    address TEXT
);

CREATE TABLE  employee
(
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    branch_id INTEGER,
    FOREIGN KEY (branch_id) REFERENCES branch(branch_id),
    PRIMARY KEY(user_id),
    role TEXT,
    experience INTEGER
);

CREATE TABLE balance
(
    balance_id SERIAL PRIMARY KEY ,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES client(user_id),
    type TEXT,
    creating_date DATE,
    capitalization INTEGER,
    balance INTEGER
);

CREATE TABLE transactions
(
    transaction_id  SERIAL PRIMARY KEY ,
    from_id INTEGER,
    FOREIGN KEY (from_id) REFERENCES balance(balance_id),
    where_id INTEGER,
    FOREIGN KEY (where_id) REFERENCES balance(balance_id),
    secure_signature TEXT,
    value INTEGER,
    date_time TIMESTAMP
);