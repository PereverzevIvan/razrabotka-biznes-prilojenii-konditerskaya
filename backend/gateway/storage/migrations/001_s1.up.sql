
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);

INSERT INTO `roles` VALUES 
    (1, 'Заказчик'),
    (2, 'Менеджер по работе с клиентами'),
    (3, 'Монеджер по закупкам'),
    (4, 'Мастер'),
    (5, 'Директор');

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `role_id` int(11) NOT NULL,
    
    `login` varchar(255) NOT NULL UNIQUE,
    `password` varchar(1024) NOT NULL,

    `first_name` varchar(255) NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `patronymic` varchar(255) DEFAULT NULL,
    
    -- `created_at` datetime DEFAULT NULL,
    -- `updated_at` datetime DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    UNIQUE (`login`),
    FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) 
);

-- Поставщики

drop table if exists `suppliers`;
create table `suppliers`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `address` varchar(255) DEFAULT NULL,
    `delivery_time_minutes` int(11) NOT NULL,
    
    PRIMARY KEY (`id`)
);

INSERT INTO `suppliers` 
VALUES 
	(1, "ИП Дмитрий Delivery", "г. Москва ул. Шарикоподшипниковская 42а", 120),
    (2, "Fast Express", "Московская область, г. Красногорск, ул. Маяковская д.3", 300);

-- Инструменты

drop table if exists `tool_types`;
create table `tool_types`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);

insert into `tool_types` (`id`, `name`) values
    (1, "Формы"),
    (2, "Ножи"),
    (3, "Кастрюли")
    ;

create table `degree_of_wears`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);

insert into `degree_of_wears` (`id`, `name`) values
    (1, "Новый"),
    (2, "Сломанный"),
    (3, "Изношенный"),
    (4, "Утилизированный")
    ;

drop table if exists `tools`;
create table `tools`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `type_id` int(11) NOT NULL,
    `supplier_id` int(11) NOT NULL,

    `name` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,

    `degree_of_wear_id` int(11) NOT NULL,

    `purchase_date` datetime DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`type_id`) REFERENCES `tool_types` (`id`),
    FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`),
    FOREIGN KEY (`degree_of_wear_id`) REFERENCES `degree_of_wears` (`id`)
);


-- Ингредиенты и украшения для тортов

drop table if exists `component_categories`;
create table `component_categories`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,

    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`)
);

INSERT INTO `component_categories` VALUES
    (1, "ingredients"),
    (2, "decorations")
    ;

drop table if exists `component_types`;
create table `component_types`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,

    `component_category_id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`component_category_id`) REFERENCES `component_categories` (`id`)
);


drop table if exists `components`;
create table `components`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `component_type_id` int(11) NOT NULL,

    `name` varchar(255) NOT NULL,
    `article` varchar(255) NOT NULL,
    `measure_unit` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`article`),
    FOREIGN KEY (`component_type_id`) REFERENCES `component_types` (`id`)
);

drop table if exists `purchased_components`;
create table `purchased_components`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `component_id` int(11) NOT NULL,
    `supplier_id` int(11) NOT NULL,

    `quantity` int(11) NOT NULL,

    `purchase_price` decimal(10, 2) NOT NULL,

    `shelf_life` datetime DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`component_id`) REFERENCES `components` (`id`),
    FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
);


-- Доставка

create table `supplier_components` (
    `supplier_id` int(11) NOT NULL,
    `component_id` int(11) NOT NULL,
    `price` decimal(10, 2) NOT NULL,

    PRIMARY KEY (`supplier_id`, `component_id`),
    FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`),
    FOREIGN KEY (`component_id`) REFERENCES `components` (`id`)
);
