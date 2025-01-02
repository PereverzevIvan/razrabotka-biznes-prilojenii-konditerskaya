-- Заказы

drop table if exists `order_statuses`;
create table `order_statuses`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    
    PRIMARY KEY (`id`)
);

insert into `order_statuses` (`id`, `name`) values
    (1, "Новый"),
    (2, "Отменен"),
    (3, "Составление спецификации"),
    (4, "Подтверждение"),
    (5, "Закупка"),
    (6, "Производство"),
    (7, "Контроль"),
    (8, "Готов"),
    (9, "Выполнен");


drop table if exists `orders`;
create table `orders`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `product_id` int(11) NOT NULL,
    
    `customer_id` int(11) NOT NULL,
    `manager_id` int(11) DEFAULT NULL,

    `status_id` int(11) NOT NULL,

    `number` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `description` varchar(255) DEFAULT NULL,
    `cost` decimal(10, 2) DEFAULT NULL,

    `decline_reason` varchar(255) DEFAULT NULL,
    
    `created_at` datetime DEFAULT NULL,
    `planned_completion_at` datetime DEFAULT NULL,
    -- `actual_completion_date` datetime DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`customer_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`manager_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`status_id`) REFERENCES `order_statuses` (`id`)
);
