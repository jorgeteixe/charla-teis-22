CREATE DATABASE IF NOT EXISTS `example_app`;

USE `example_app`;

CREATE TABLE IF NOT EXISTS `user` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255),
    `login` varchar(255),
    `dni` varchar(9)
) ENGINE=InnoDB;
