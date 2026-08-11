CREATE TABLE games(
                      id uuid not null unique,
                      grid int[][]
--     created_at timestamp with time zone default current_timestamp not null
);

-- CREATE TABLE users(
--     id uuid not null unique,
--     login VARCHAR(40) NOT NULL,
--     password VARCHAR NOT NULL
-- );