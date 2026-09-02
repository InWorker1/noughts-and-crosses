CREATE TABLE games(
                      id uuid not null unique,
                      grid int[][]
--     created_at timestamp with time zone default current_timestamp not null
);

CREATE TABLE users(
                      id uuid not null unique,
                      login TEXT not null unique,
                      password TEXT not null,
                      game_now uuid not null unique,
--     games uuid[],
                      constraint fk_id_game foreign key (game_now) references games(id)
);

-- CREATE TABLE users(
--     id uuid not null unique,
--     login VARCHAR(40) NOT NULL,
--     password VARCHAR NOT NULL
-- );