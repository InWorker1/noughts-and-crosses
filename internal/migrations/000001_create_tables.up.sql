CREATE TABLE games(
                      id uuid not null unique, -- id игры
                      grid int[][], -- поле игры
                      waiting BOOLEAN not null, -- ожидание игроков
                      move_player uuid DEFAULT NULL, -- кто ходит
                      winner uuid DEFAULT NULL, -- кто выиграл
                      draw BOOLEAN DEFAULT FALSE, -- ничья
                      x_player uuid,
                      o_player uuid,
                      constraint fk_user_id_win foreign key (winner) references users(id),
                      constraint fk_user_id_move foreign key (move_player) references users(id),
                      constraint fk_user_id_x foreign key (x_player) references users(id),
                      constraint fk_user_id_o foreign key (o_player) references users(id)
);

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