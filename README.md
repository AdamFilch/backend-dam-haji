# backend-checkers



Sample return from api from gameplay move
```
{
    "Adam": {
        "points": 3,
        "letter": "O",
        "winner": true 
    },
    "Nazmi": {
        "points": 10,
        "letter": "X"
    },
    "board": {
        "0": ["1","2","3","4","5","6","7","8","9","10"],
        "A": [" ","X"," ","X"," ","X"," ","X"," ","X"],
        "B": ["X"," ","X"," ","X"," ","X"," ","X"," "],
        "C": [" "," "," "," "," "," "," "," "," "," "],
        "D": [" "," "," "," "," "," "," "," "," "," "],
        "E": [" "," "," "," "," "," "," "," "," "," "],
        "F": [" "," "," "," "," "," "," "," "," "," "],
        "G": [" "," "," "," "," "," "," "," "," "," "],
        "H": [" "," "," "," "," "," "," "," "," "," "],
        "I": [" ","0"," ","0"," ","0"," ","0"," ","0"],
        "J": ["0"," ","0"," ","0"," ","0"," ","0"," "],
        "Z": ["1","2","3","4","5","6","7","8","9","10"]
    },
    "data": {
        "time": "30:00",
        "error": "You are not allowed to move in that direction",
        "end": "Congratulations, you have won the game",
        "playback": "Adam has just moved E3 to B6"
    } 
}
```


Database Schema
---

- USER_T
    - user_id_pk
    - user_username
    - user_games_won


- GAME_T
    - game_id_pk
    - game_player_one_id_fk
    - game_player_two_id_fk
    - game_board_state
    - game_status
    - game_winner
    - game_started
    - game_last_updated

- MOVES_T
    - move_id_pk
    - move_game_id_fk
    - move_from
    - move_to
    - move_created

