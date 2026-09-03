CREATE TABLE games(
                      id uuid not null unique,
                      grid int[][]
--     created_at timestamp with time zone default current_timestamp not null
);

CREATE TABLE online_games(
                             id uuid not null unique, -- id игры
                             grid int[][], -- поле игры
                             waiting BOOLEAN not null, -- ожидание игроков
                             move_player uuid, -- кто ходит
                             win_player uuid, -- кто выиграл
                             draw BOOLEAN -- ничья
); -- таблица для состояний онлайн игры

CREATE TABLE users(
                      id uuid not null unique,
                      login TEXT not null unique,
                      password TEXT not null,
                      role VARCHAR(2),
                      game_now uuid unique,
--     games uuid[],
                      constraint fk_id_game foreign key (game_now) references games(id)
);

-- CREATE TABLE users(
--     id uuid not null unique,
--     login VARCHAR(40) NOT NULL,
--     password VARCHAR NOT NULL
-- );