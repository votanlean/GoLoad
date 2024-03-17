CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT UNSIGNED PRIMARY KEY,
    `username` VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS `user_passwords` (
    `hash` VARCHAR(128) NOT NULL,
    `user_id` BIGINT UNSIGNED PRIMARY KEY,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `download_tasks` (
    `id` BIGINT UNSIGNED PRIMARY KEY,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `type` SMALLINT NOT NULL,
    `url` TEXT NOT NULL,
    `status` SMALLINT NOT NULL,
    `metadata` TEXT NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
