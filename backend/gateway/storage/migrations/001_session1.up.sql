use `confectionary`;

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `roles` VALUES 
    (1, 'user'), 
    (2, 'admin');

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `role_id` int(11) NOT NULL,
    `email` varchar(255) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL,
    -- `created_at` datetime DEFAULT NULL,
    -- `updated_at` datetime DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) 
)