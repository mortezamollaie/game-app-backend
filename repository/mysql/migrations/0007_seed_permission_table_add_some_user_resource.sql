-- +migrate Up
INSERT INTO `permissions`(`id`, `title`) values (1, 'user-delete');
INSERT INTO `permissions`(`id`, `title`) values (2, 'user-list');

INSERT INTO 'access_controls' (`actor_type`, `actor_id`, `permission_id`) VALUES ('role', '2', '1');
INSERT INTO 'access_controls' (`actor_type`, `actor_id`, `permission_id`) VALUES ('role', '2', '2');

-- +migrate Down
DELETE FROM `permissions` WHERE id in (1,2);