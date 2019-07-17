GRANT ALL PRIVILEGES ON DATABASE baseball TO user1;
-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user1;
-- GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO user1;

-- ALTER DEFAULT PRIVILEGES
-- FOR USER user1
-- IN SCHEMA public
-- GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO user1;

create table game (
game_date date,
number_of_game varchar,
day_of_week varchar,

visiting_team varchar,
visiting_team_league varchar,
visiting_game_number int,

home_team varchar,
home_team_league varchar,
home_team_game_number int,

visiting_team_score int,
home_team_score int,

game_length_in_outs int,
day_night_indicator varchar,

completion_information varchar,
forfeit_information varchar,
protest_information varchar,

park_id varchar,
attendance int,
time_of_game_in_mins int,
visiting_line_score varchar,
home_line_score varchar,

visiting_ab   int,
visiting_h    int,
visiting_2B   int,
visiting_3B   int,
visiting_hr   int,
visiting_rbi  int,
visiting_sh   int,
visiting_sf   int,
visiting_hbp  int,
visiting_bb   int,
visiting_ibb  int,
visiting_k    int,
visiting_sb   int,
visiting_cs   int,
visiting_gidp int,
visiting_ci   int,
visiting_lob  int,

visiting_pitchers_used         int,
visiting_individual_earned_runs int,
visiting_team_earned_runs       int,
visiting_wild_pitches          int,
visiting_balks                 int,

visiting_putouts     int,
visiting_assists     int,
visiting_errors      int,
visiting_passed_balls int,
visiting_double_plays int,
visiting_triple_plays int,

home_ab   int,
home_h    int,
home_2B   int,
home_3B   int,
home_hr   int,
home_rbi  int,
home_sh   int,
home_sf   int,
home_hbp  int,
home_bb   int,
home_ibb  int,
home_k    int,
home_sb   int,
home_cs   int,
home_gidp int,
home_ci   int,
home_lob  int,

home_pitchers_used         int,
home_individual_earned_runs int,
home_team_earned_runs       int,
home_wild_pitches          int,
home_balks          int,

home_putouts     int,
home_assists     int,
home_errors      int,
home_passed_balls int,
home_double_plays int,
home_triple_plays int,

home_plate_umpire_id    varchar,
home_plate_umpire_name  varchar,
first_base_umpire_id    varchar,
first_base_umpire_name  varchar,
second_base_umpire_id   varchar,
second_base_umpire_name varchar,
third_base_umpire_id    varchar,
third_base_umpire_name  varchar,
left_field_umpire_id    varchar,
left_field_umpire_name  varchar,
right_field_umpire_id   varchar,
right_field_umpire_name varchar,

visiting_manager_id   varchar,
visiting_manager_name varchar,
home_manager_id       varchar,
home_manager_name     varchar,

winning_pitcher_id   varchar,
winning_pitcher_name varchar,
losing_pitcher_id    varchar,
losing_pitcher_name  varchar,
saving_pitcher_id    varchar,
saving_pitcher_name  varchar,

game_winning_rbi_batter_id   varchar,
game_winning_rbi_batter_name varchar,

visiting_starting_pitcher_id   varchar,
visiting_starting_pitcher_name varchar,
home_starting_pitcher_id       varchar,
home_starting_pitcher_name     varchar,

visiting_player1_id       varchar,
visiting_player1_name     varchar,
visiting_player1_position int,
visiting_player2_id       varchar,
visiting_player2_name     varchar,
visiting_player2_position int,
visiting_player3_id       varchar,
visiting_player3_name     varchar,
visiting_player3_position int,
visiting_player4_id       varchar,
visiting_player4_name     varchar,
visiting_player4_position int,
visiting_player5_id       varchar,
visiting_player5_name     varchar,
visiting_player5_position int,
visiting_player6_id       varchar,
visiting_player6_name     varchar,
visiting_player6_position int,
visiting_player7_id       varchar,
visiting_player7_name     varchar,
visiting_player7_position int,
visiting_player8_id       varchar,
visiting_player8_name     varchar,
visiting_player8_position int,
visiting_player9_id       varchar,
visiting_player9_name     varchar,
visiting_player9_position int,

home_player1_id       varchar,
home_player1_name     varchar,
home_player1_position int,
home_player2_id       varchar,
home_player2_name     varchar,
home_player2_position int,
home_player3_id       varchar,
home_player3_name     varchar,
home_player3_position int,
home_player4_id       varchar,
home_player4_name     varchar,
home_player4_position int,
home_player5_id       varchar,
home_player5_name     varchar,
home_player5_position int,
home_player6_id       varchar,
home_player6_name     varchar,
home_player6_position int,
home_player7_id       varchar,
home_player7_name     varchar,
home_player7_position int,
home_player8_id       varchar,
home_player8_name     varchar,
home_player8_position int,
home_player9_id       varchar,
home_player9_name     varchar,
home_player9_position int,

additional_information  varchar,
acquisition_information varchar,

primary key(visiting_team, home_team, game_date, number_of_game)
);

create index i_game_date_teams on game(visiting_team, home_team, game_date);

-- Teams

create table team (
    team_symbol varchar,
    founded int,
    league varchar,
    location varchar,
    name varchar,

    primary key(team_symbol)
);

-- Parks

create table park (
    park_id varchar,
    name varchar,
    nickname varchar,
    city varchar,
    state varchar,
    start_date date,
    end_date date,
    league varchar,

    primary key(park_id)
);