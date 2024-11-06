-- Контроль качества

drop table `quality_control`;
create table `quality_control`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `order_id` int(11) NOT NULL,
    `master_id` int(11) NOT NULL,

    `is_satisfies` tinyint(1) NOT NULL,
    `comment` varchar(255) DEFAULT NULL,
    
    `parameter` varchar(255) NOT NULL,
    -- Или использовать FK по таблице `parameters`
    -- `parameter_id` int(11) NOT NULL,
    
    PRIMARY KEY (`id`),
    FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
    FOREIGN KEY (`master_id`) REFERENCES `users` (`id`)
};