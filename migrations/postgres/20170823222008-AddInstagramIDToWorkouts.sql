
-- +migrate Up
ALTER TABLE workouts
ADD instagram_id VARCHAR;

-- +migrate Down
ALTER TABLE workouts
DROP COLUMN instagram_id;
