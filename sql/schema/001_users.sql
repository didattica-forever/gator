
-- The -- +Goose Up and -- +Goose Down comments are case sensitive and required. 
-- They tell Goose how to run the migration in each direction.
-- 
-- Connection string
-- postgres://postgres:manager@localhost:5432/gator


-- +goose Up
CREATE TABLE users (
    
    id uuid primary key, -- a UUID that will serve as the primary key
    created_at timestamp not null, -- a TIMESTAMP that can not be null
    updated_at timestamp not null, -- a TIMESTAMP that can not be null
    name varchar(255) not null unique -- a unique string that can not be null

);

-- +goose Down
DROP TABLE users;