-- +migrate Up
CREATE TABLE `permissions`(
                      `id` INT PRIMARY KEY AUTO_INCREMENT,
                      `title` VARCHAR(191) NOT NULL UNIQUE,
                      `created_at` datetime DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO permissions(title) values ('user-delete');
INSERT INTO permissions(title) values ('user-list');

-- +migrate Down
DROP TABLE permissions;