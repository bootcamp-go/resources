-- DDL
DROP DATABASE IF EXISTS `products_test_db`;

CREATE DATABASE `products_test_db`;

USE `products_test_db`;

CREATE TABLE `warehouses` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE `products` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `type` varchar(255) NOT NULL,
    `count` int NOT NULL,
    `price` decimal(10, 2) NOT NULL,
    `warehouse_id` int(11) NOT NULL,
    PRIMARY KEY (`id`)
    KEY `idx_products_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fk_products_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8;