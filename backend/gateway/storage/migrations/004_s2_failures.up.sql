-- Сбои оборудования

create table `tools_failures`  {
    `id` int(11) NOT NULL AUTO_INCREMENT,
    
    `tool_id` int(11) NOT NULL,
    -- `name` varchar(255) NOT NULL,
    `reason` varchar(255) DEFAULT NULL,
    
    `failure_at` datetime DEFAULT NOW(),
    `fix_at` datetime DEFAULT NULL,
    

    PRIMARY KEY (`id`),
    FOREIGN KEY (`tool_id`) REFERENCES `tools` (`id`)
};
