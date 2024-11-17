-- Цеха
drop table if exists `workshops`;

create table `workshops` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `scheme_image` varchar(255) NOT NULL,
  `rotated` tinyint (1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`name`)
);

drop table if exists `workshop_object_types`;

create table `workshop_object_types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`name`)
);

insert into
  `workshop_object_types` (`id`, `name`, `image`)
values
  (1, "Equipment", "equipment.png"),
  (2, "Exit", "exit.jpg"),
  (3, "Fire extinguisher", "FireExtinguisher.png"),
  (4, "First aid", "FirstAid.png");

drop table if exists `workshop_objects`;

create table `workshop_objects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `workshop_id` int(11) NOT NULL,
  `object_type_id` int(11) NOT NULL,
  `x` DOUBLE PRECISION NOT NULL,
  `y` DOUBLE PRECISION NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`workshop_id`) REFERENCES `workshops` (`id`),
  FOREIGN KEY (`object_type_id`) REFERENCES `workshop_object_types` (`id`)
);
