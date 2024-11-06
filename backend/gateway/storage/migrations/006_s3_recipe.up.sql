-- Рецепты

drop table if exists `recipe_operations`;
create table `recipe_operations`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `product_id` int(11) NOT NULL,
    `tool_id` int(11) DEFAULT NULL,

    `description` varchar(255) NOT NULL,
    `duration_minutes` int(11) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`tool_id`) REFERENCES `tools` (`id`)
};

drop table if exists `recipe_semiproducts`;
create table `recipe_semiproducts`  {    
    `product_id` int(11) NOT NULL,
    `semiproduct_id` int(11) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`semiproduct_id`) REFERENCES `products` (`id`)
};


drop table if exists `recipe_components`;
create table `recipe_ingredients`  {    
    `product_id` int(11) NOT NULL,
    `component_id` int(11) NOT NULL,
    
    `count` int(11) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
    FOREIGN KEY (`component_id`) REFERENCES `components` (`id`)
};
