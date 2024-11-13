-- Сбои оборудования

create table `tool_failure_reasons`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`name`)
);

create table `tool_failures`  (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `tool_id` int(11) NOT NULL,
    
    `master_id` int(11) NOT NULL,

    `failure_reason_id` int(11) NOT NULL,
    
    `failure_at` datetime DEFAULT NOW(),
    `fixed_at` datetime DEFAULT NULL,
    

    PRIMARY KEY (`id`),
    FOREIGN KEY (`tool_id`) REFERENCES `tools` (`id`),
    FOREIGN KEY (`master_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`failure_reason_id`) REFERENCES `tool_failure_reasons` (`id`)
);
