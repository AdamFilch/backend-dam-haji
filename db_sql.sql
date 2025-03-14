
------- Table Relationship ------

-- User can have many games
-- game can have 2 users
-- game can have many moves



---- USER TABLE ----
CREATE TABLE USER_T (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    points INT DEFAULT 0,
    games_won INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---- GAME TABLE ----
CREATE TABLE GAME_T (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player1_id INT REFERENCES users(id),
    player2_id INT REFERENCES users(id),
    board_state JSONB NOT NULL,  -- Stores the board as JSON
    winner_id INT REFERENCES users(id) DEFAULT NULL,
    status VARCHAR(20) CHECK (status IN ('ongoing', 'completed', 'abandoned')) DEFAULT 'ongoing',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


--- MOVE TABLE ----
CREATE TABLE moves (
    id SERIAL PRIMARY KEY,
    game_id UUID REFERENCES games(id) ON DELETE CASCADE,
    player_id INT REFERENCES users(id),
    move_from VARCHAR(5) NOT NULL,  -- e.g., "E3"
    move_to VARCHAR(5) NOT NULL,    -- e.g., "B6"
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);




