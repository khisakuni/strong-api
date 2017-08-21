
-- +migrate Up
CREATE TABLE workouts(id BIGSERIAL, name VARCHAR, description text);

-- +migrate Down
DROP TABLE workouts;
