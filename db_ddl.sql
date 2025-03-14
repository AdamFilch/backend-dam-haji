
------- Table Relationship ------

-- User can have many games
-- game can have 2 users
-- game can have many moves



---- USERS TABLE ----
CREATE TABLE user_t (
    user_id_pk SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    total_points INT DEFAULT 0,
    games_won INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---- GAMES TABLE ----
CREATE TABLE games_t (
    game_id_pk UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player1_id INT REFERENCES user_t(user_id_pk) ON DELETE CASCADE,
    player2_id INT REFERENCES user_t(user_id_pk) ON DELETE CASCADE,
    board_state JSONB NOT NULL,  -- Stores the board as JSON
    winner_id INT REFERENCES user_t(user_id_pk) DEFAULT NULL,
    status VARCHAR(20) CHECK (status IN ('ongoing', 'completed', 'abandoned')) DEFAULT 'ongoing',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---- MOVES TABLE ----
CREATE TABLE moves_t (
    move_id_pk SERIAL PRIMARY KEY,
    game_id UUID REFERENCES game_t(game_id_pk) ON DELETE CASCADE,
    user_id INT REFERENCES user_t(user_id_pk) ON DELETE CASCADE,
    move_from VARCHAR(5) NOT NULL,  -- e.g., "E3"
    move_to VARCHAR(5) NOT NULL,    -- e.g., "B6"
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);




