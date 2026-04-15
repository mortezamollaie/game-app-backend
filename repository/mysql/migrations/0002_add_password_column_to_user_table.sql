-- +migrate Up
ALTER TABLE users ADD column password text not null;

-- +migrate Down
ALTER TABLE users DROP column password;
