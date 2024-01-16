USE `melisprint`;

TRUNCATE TABLE `sellers`;
TRUNCATE TABLE `warehouses`;
TRUNCATE TABLE `sections`;
TRUNCATE TABLE `products`;
TRUNCATE TABLE `employees`;
TRUNCATE TABLE `buyers`;

-- DML
INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`) VALUES
(1, 'Company A', '123 Main St', '123-456-7890'),
(2, 'Company B', '456 Elm St', '123-456-7891'),
(3, 'Company C', '789 Oak St', '123-456-7892'),
(4, 'Company D', '101 Pine St', '123-456-7893'),
(5, 'Company E', '102 Maple St', '123-456-7894'),
(6, 'Company F', '103 Cedar St', '123-456-7895'),
(7, 'Company G', '104 Birch St', '123-456-7896'),
(8, 'Company H', '105 Willow St', '123-456-7897'),
(9, 'Company I', '106 Cherry St', '123-456-7898'),
(10, 'Company J', '107 Walnut St', '123-456-7899');

INSERT INTO `warehouses` (`warehouse_code`, `address`, `telephone`, `minimum_capacity`, `minimum_temperature`) VALUES
('WH01', '200 Warehouse Rd', '234-567-8901', 100, 0),
('WH02', '201 Warehouse Ln', '234-567-8902', 150, -5),
('WH03', '202 Storage Blvd', '234-567-8903', 120, 2),
('WH04', '203 Distribution Ave', '234-567-8904', 200, -2),
('WH05', '204 Inventory St', '234-567-8905', 180, 0),
('WH06', '205 Logistics Way', '234-567-8906', 160, -3),
('WH07', '206 Depot Dr', '234-567-8907', 140, 1),
('WH08', '207 Supply Ct', '234-567-8908', 170, -4),
('WH09', '208 Goods Rd', '234-567-8909', 130, 3),
('WH10', '209 Freight St', '234-567-8910', 190, -1);

INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES
(1, 0, -5, 50, 20, 100, 1, 1),
(2, -2, -6, 60, 30, 110, 2, 2),
(3, 1, -4, 70, 40, 120, 3, 3),
(4, -3, -7, 80, 50, 130, 4, 4),
(5, 2, -5, 90, 60, 140, 5, 5),
(6, -4, -8, 100, 70, 150, 6, 6),
(7, 3, -6, 110, 80, 160, 7, 7),
(8, -5, -9, 120, 90, 170, 8, 8),
(9, 4, -7, 130, 100, 180, 9, 9),
(10, -6, -10, 140, 110, 190, 10, 10);

INSERT INTO `products` (`product_code`, `description`, `height`, `lenght`, `width`, `weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `seller_id`, `product_type_id`) VALUES
('P1001', 'Product 1', 10, 5, 8, 2, 0.1, 0.2, -5, 1, 1),
('P1002', 'Product 2', 12, 6, 9, 2.5, 0.15, 0.25, -6, 2, 2),
('P1003', 'Product 3', 14, 7, 10, 3, 0.2, 0.3, -7, 3, 3),
('P1004', 'Product 4', 16, 8, 11, 3.5, 0.25, 0.35, -8, 4, 4),
('P1005', 'Product 5', 18, 9, 12, 4, 0.3, 0.4, -9, 5, 5),
('P1006', 'Product 6', 20, 10, 13, 4.5, 0.35, 0.45, -10, 6, 6),
('P1007', 'Product 7', 22, 11, 14, 5, 0.4, 0.5, -11, 7, 7),
('P1008', 'Product 8', 24, 12, 15, 5.5, 0.45, 0.55, -12, 8, 8),
('P1009', 'Product 9', 26, 13, 16, 6, 0.5, 0.6, -13, 9, 9),
('P1010', 'Product 10', 28, 14, 17, 6.5, 0.55, 0.65, -14, 10, 10);

INSERT INTO `employees` (`card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES
('E1001', 'John', 'Doe', 1),
('E1002', 'Jane', 'Smith', 2),
('E1003', 'Michael', 'Johnson', 3),
('E1004', 'Emily', 'Davis', 4),
('E1005', 'David', 'Miller', 5),
('E1006', 'Sarah', 'Wilson', 6),
('E1007', 'Robert', 'Moore', 7),
('E1008', 'Jennifer', 'Taylor', 8),
('E1009', 'William', 'Anderson', 9),
('E1010', 'Jessica', 'Thomas', 10);

INSERT INTO `buyers` (`card_number_id`, `first_name`, `last_name`) VALUES
('B1001', 'Alice', 'Brown'),
('B1002', 'Mark', 'Jones'),
('B1003', 'Linda', 'Garcia'),
('B1004', 'Brian', 'Williams'),
('B1005', 'Susan', 'Martinez'),
('B1006', 'Richard', 'Lee'),
('B1007', 'Karen', 'Harris'),
('B1008', 'Steven', 'Clark'),
('B1009', 'Betty', 'Lopez'),
('B1010', 'Edward', 'Gonzalez');