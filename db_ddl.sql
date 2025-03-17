
------- Table Relationship ------

-- User can have many games
-- game can have 2 users
-- game can have many moves



---- USERS TABLE ----
CREATE TABLE users (
    username VARCHAR(50) PRIMARY KEY,  -- Now username is the unique ID
    total_points INT DEFAULT 0,
    games_won INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---- GAMES TABLE ----
CREATE TABLE games (
    game_id_pk UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player1_username VARCHAR(50) REFERENCES users(username) ON DELETE CASCADE NOT NULL,
    player2_username VARCHAR(50) REFERENCES users(username) ON DELETE CASCADE NULL,  -- Allow NULL initially
    board_state JSONB NOT NULL,  -- Stores the board as JSON
    winner_username VARCHAR(50) REFERENCES users(username) DEFAULT NULL, -- Changed from winner_id
    status VARCHAR(20) CHECK (status IN ('ongoing', 'completed', 'abandoned')) DEFAULT 'ongoing',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
---- MOVES TABLE ----
CREATE TABLE moves (
    move_id_pk SERIAL PRIMARY KEY,
    game_id UUID REFERENCES games(game_id_pk) ON DELETE CASCADE,
    username VARCHAR(50) REFERENCES users(username) ON DELETE CASCADE, -- Changed from user_id
    move_from VARCHAR(5) NOT NULL,  -- e.g., "E3"
    move_to VARCHAR(5) NOT NULL,    -- e.g., "B6"
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

