-- DDL
DROP DATABASE IF EXISTS `users_bootcamp_db`;

CREATE DATABASE `users_bootcamp_db`;

USE `users_bootcamp_db`;

CREATE TABLE `users` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `age` int(11) NOT NULL,
    `email` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);