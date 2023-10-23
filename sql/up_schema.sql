CREATE TABLE `roles` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(64) NOT NULL
);

CREATE TABLE `users` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(64) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    `firstname` VARCHAR(64) NOT NULL,
    `lastname` VARCHAR(64) NOT NULL,
    `roles_id` INT NOT NULL DEFAULT 2,
    FOREIGN KEY (`roles_id`) REFERENCES `roles`(`id`)
);

CREATE TABLE `restaurants` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(128) NOT NULL,
    `address` VARCHAR(128) NOT NULL,
    `glink` VARCHAR(128) NOT NULL,
    `lat` DECIMAL(11,7) NOT NULL,
    `lon` DECIMAL(11,7) NOT NULL,
    `users_id` INT,
    FOREIGN KEY (`users_id`) REFERENCES users(`id`)
);

CREATE TABLE `daily_offers` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `link` VARCHAR(128) NOT NULL,
    `date` DATE NOT NULL,
    `restaurants_id` INT NOT NULL,
    FOREIGN KEY (`restaurants_id`) REFERENCES restaurants(`id`)
);

INSERT INTO `roles` (`id`, `title`)
VALUES 
    (1, 'admin'),
    (2, 'moderator');