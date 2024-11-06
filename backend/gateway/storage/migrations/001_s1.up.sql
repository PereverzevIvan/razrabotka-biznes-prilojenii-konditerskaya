use `confectionary`;

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);

INSERT INTO `roles` VALUES 
    (1, 'user'),
    (2, 'costumer'),
    (3, 'customer service manager'),
    (4, 'purchasing manager'),
    (5, 'master'),
    (6, 'director');

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `role_id` int(11) NOT NULL,
    `login` varchar(255) NOT NULL UNIQUE,
    `password` varchar(1024) NOT NULL,
    -- `created_at` datetime DEFAULT NULL,
    -- `updated_at` datetime DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    UNIQUE (`login`),
    FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) 
);

-- Поставщики

drop table if exists `suppliers`;
create table `suppliers`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `contact_person` varchar(255) NOT NULL,
    `phone` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
};

-- Инструменты

drop table if exists `tool_types`;
create table `tool_types`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`name`)
};

drop table if exists `tools`;
create table `tools`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `type_id` int(11) NOT NULL,
    `supplier_id` int(11) NOT NULL,

    `name` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,

    `degree_of_wear` int(11) NOT NULL,

    `purchase_date` datetime DEFAULT NULL,

    PRIMARY KEY (`id`),
    CHECK (0 <= `degree_of_wear` AND `degree_of_wear` <= 100)
};


-- Ингредиенты и украшения для тортов

drop table if exists `component_types`;
create table `ingredient_types`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`)
};

INSERT INTO `component_types` VALUES
    (1, "ingredients"),
    (2, "decorations");


drop table if exists `components`;
create table `ingredient_types`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `type_id` int(11) NOT NULL,

    `name` varchar(255) NOT NULL,
    `article` varchar(255) NOT NULL,
    `measure_unit` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`article`),
    FOREIGN KEY (`type_id`) REFERENCES `component_types` (`id`)
};

drop table if exists `purchased_components`;
create table `ingredients`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `component_id` int(11) NOT NULL,
    `supplier_id` int(11) NOT NULL,

    `count` int(11) NOT NULL,

    `purchase_price` decimal(10, 2) NOT NULL,

    `shelf_life` datetime DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`component_id`) REFERENCES `components` (`id`),
    FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
};


-- Доставка

create table `supplier_components` {
    `supplier_id` int(11) NOT NULL,
    `component_id` int(11) NOT NULL,
    `price` decimal(10, 2) NOT NULL

    PRIMARY KEY (`supplier_id`, `component_id`),
    FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`),
    FOREIGN KEY (`component_id`) REFERENCES `components` (`id`)
};