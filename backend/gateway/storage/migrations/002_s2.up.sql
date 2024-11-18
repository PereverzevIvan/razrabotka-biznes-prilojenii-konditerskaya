drop table if exists `products`;
create table `products`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `sizes` varchar(255) DEFAULT NULL,
    `is_semiproduct` tinyint(1) DEFAULT 0 NOT NULL,

    PRIMARY KEY (`id`)
);



-- Рецепты
drop table if exists `recipe_semiproducts`;
create table `recipe_semiproducts`  (    
    `product_id` int(11) NOT NULL,
    `semiproduct_id` int(11) NOT NULL,

    `count` int(11) NOT NULL,

    PRIMARY KEY (`product_id`, `semiproduct_id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`semiproduct_id`) REFERENCES `products` (`id`)
);


drop table if exists `recipe_components`;
create table `recipe_components`  (    
    `product_id` int(11) NOT NULL,
    `component_id` int(11) NOT NULL,
    
    `count` int(11) NOT NULL,

    PRIMARY KEY (`product_id`, `component_id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`component_id`) REFERENCES `components` (`id`)
);

drop table if exists `recipe_operations`;
create table `recipe_operations`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `product_id` int(11) NOT NULL,
    `tool_type_id` int(11) DEFAULT NULL,


    `description` varchar(255) NOT NULL,
    `duration_minutes` int(11) NOT NULL,
    `order_idx` int(11) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`product_id`, `order_idx`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`tool_type_id`) REFERENCES `tool_types` (`id`)
);