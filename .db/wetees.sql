CREATE DATABASE `wetees` /*!40100 DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci */;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`),
  UNIQUE KEY `users_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.users (username,email,password,created_at) VALUES
	 ('Admin','roni.connect@gmail.com','$2y$12$y9tWFFMGT95QNAIhoh1CwOiLNL8ZlK6PWZhewqFXfTAkA6Y8Uyzs6','2025-01-03 13:32:50'),
	 ('Pep Guardiola','pep.guardiola@gmail.com','$2y$12$y9tWFFMGT95QNAIhoh1CwOiLNL8ZlK6PWZhewqFXfTAkA6Y8Uyzs6','2025-01-03 13:32:50'),
	 ('Jose Mourinho','jose.mourinho@gmail.com','$2y$12$y9tWFFMGT95QNAIhoh1CwOiLNL8ZlK6PWZhewqFXfTAkA6Y8Uyzs6','2025-01-03 13:32:50'),
	 ('Jurgen Klopp','jurgen.klopp@gmail.com','$2y$12$y9tWFFMGT95QNAIhoh1CwOiLNL8ZlK6PWZhewqFXfTAkA6Y8Uyzs6','2025-01-03 13:32:50'),
	 ('Carlo Anceloti','carlo.anceloti@gmail.com','$2y$12$y9tWFFMGT95QNAIhoh1CwOiLNL8ZlK6PWZhewqFXfTAkA6Y8Uyzs6','2025-01-03 13:32:50'),
	 ('Antonio Conte','antonio.conte@gmail.com','$2a$14$MaMHOhfRrer37DKZ/qC4euTjuMRR/a/VV9d724tlXjQFpTeRThvei','2025-01-03 13:32:50'),
	 ('Alex Ferguson','alex.ferguson@gmail.com','$2a$14$x8wQXyxEqXCIn8Gy3IG.OOsdn3dO8gFzUDMPrhPo9hczKfkTCViKG','2025-01-04 21:03:16');


CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `acc_type` varchar(20) NOT NULL DEFAULT 'saving' COMMENT 'savings, checking, credit',
  `acc_number` char(10) NOT NULL,
  `balance` decimal(10,2) NOT NULL DEFAULT 0.00,
  `dto` date NOT NULL COMMENT 'Date Opened',
  `dtc` date DEFAULT NULL COMMENT 'Date Closed',
  `status` varchar(20) NOT NULL COMMENT 'active, suspended, closed',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account_unique` (`acc_number`),
  KEY `account_users_FK` (`user_id`),
  KEY `account_acc_number_IDX` (`acc_number`) USING BTREE,
  CONSTRAINT `account_users_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.account (user_id,acc_type,acc_number,balance,dto,dtc,status) VALUES
	 (1,'savings','5550000001',250.00,'2025-01-05',NULL,'active'),
	 (2,'savings','5550000002',0.00,'2025-01-05',NULL,'active');


CREATE TABLE `transaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `acc_number` char(10) NOT NULL,
  `trans_type` varchar(20) NOT NULL COMMENT 'deposit, withdrawal',
  `amount` decimal(10,2) NOT NULL DEFAULT 0.00,
  `trans_date` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `transaction_account_FK` (`acc_number`),
  CONSTRAINT `transaction_account_FK` FOREIGN KEY (`acc_number`) REFERENCES `account` (`acc_number`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.`transaction` (acc_number,trans_type,amount,trans_date) VALUES
	 ('5550000001','deposit',100.00,'2025-01-04 00:00:00'),
	 ('5550000001','deposit',100.00,'2025-01-04 00:00:00'),
	 ('5550000001','deposit',100.00,'2025-01-05 08:56:29'),
	 ('5550000001','deposit',100.00,'2025-01-05 08:56:30'),
	 ('5550000001','deposit',100.00,'2025-01-05 08:56:31'),
	 ('5550000001','deposit',100.00,'2025-01-05 08:56:32'),
	 ('5550000001','deposit',200.00,'2025-01-05 08:59:44'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:05:10'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:08:23'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:03');
INSERT INTO wetees.`transaction` (acc_number,trans_type,amount,trans_date) VALUES
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:06'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:06'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:07'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:08'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:08'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:09'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:10:10'),
	 ('5550000001','withdrawal',50.00,'2025-01-05 09:29:25'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:43:59'),
	 ('5550000001','deposit',200.00,'2025-01-05 09:44:01');
INSERT INTO wetees.`transaction` (acc_number,trans_type,amount,trans_date) VALUES
	 ('5550000001','deposit',200.00,'2025-01-05 09:44:02'),
	 ('5550000001','withdrawal',100.00,'2025-01-05 10:11:02'),
	 ('5550000001','withdrawal',100.00,'2025-01-05 10:11:04'),
	 ('5550000001','withdrawal',100.00,'2025-01-05 10:11:04'),
	 ('5550000001','withdrawal',1000.00,'2025-01-05 10:11:22'),
	 ('5550000001','withdrawal',1000.00,'2025-01-05 10:11:27'),
	 ('5550000001','withdrawal',1000.00,'2025-01-05 10:11:31'),
	 ('5550000001','deposit',200.00,'2025-01-05 10:13:43'),
	 ('5550000001','deposit',200.00,'2025-01-05 10:13:44'),
	 ('5550000001','deposit',200.00,'2025-01-05 10:13:44');
INSERT INTO wetees.`transaction` (acc_number,trans_type,amount,trans_date) VALUES
	 ('5550000001','deposit',200.00,'2025-01-05 10:13:45'),
	 ('5550000001','deposit',200.00,'2025-01-05 10:13:46'),
	 ('5550000001','deposit',200.00,'2025-01-05 10:13:46'),
	 ('5550000001','withdrawal',1000.00,'2025-01-05 10:13:51');


CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cat_name` varchar(50) NOT NULL,
  `parent_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.categories (cat_name,parent_id) VALUES
	 ('Pakaian',0),
	 ('Elektronik',0),
	 ('Peralatan rumah tangga',0),
	 ('Makanan',0);

CREATE TABLE `merchants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_id` int(11) NOT NULL,
  `merchant_name` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `merchants_users_FK` (`admin_id`),
  CONSTRAINT `merchants_users_FK` FOREIGN KEY (`admin_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.merchants (admin_id,merchant_name,created_at) VALUES
	 (1,'Blueberry Indonesia','2025-01-03 13:32:50'),
	 (2,'Karya Kampoeng','2025-01-03 13:32:50');

CREATE TABLE `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(50) NOT NULL,
  `category_id` int(11) NOT NULL,
  `merchant_id` int(11) NOT NULL,
  `price` decimal(10,2) NOT NULL DEFAULT 0.00,
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '1 is available, 0 is out of stock',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `products_merchants_FK` (`merchant_id`),
  KEY `products_categories_FK` (`category_id`),
  CONSTRAINT `products_categories_FK` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `products_merchants_FK` FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.products (name,category_id,merchant_id,price,status,created_at) VALUES
	 ('Kemeja Tactical Lengan Panjang',1,1,90000.00,1,'2025-01-03 15:44:42'),
	 ('Jaket WP',1,1,145000.00,1,'2025-01-03 15:44:42'),
	 ('Rompi Sholat Hodoo',1,1,50000.00,1,'2025-01-03 15:48:56'),
	 ('Kaos Kerah Pria Knit Shirt Lengan Pendek',1,1,1200000.00,1,'2025-01-03 15:48:56'),
	 ('Celana Olahraga Knit Shirt',1,1,790000.00,1,'2025-01-03 15:48:56'),
	 ('Oxone Giant Oven',2,2,2250000.00,1,'2025-01-03 15:54:38'),
	 ('Gabor Rice Cooker',2,2,245000.00,1,'2025-01-03 15:54:38'),
	 ('Smart Chef Presto',2,2,1200000.00,1,'2025-01-03 15:54:38'),
	 ('Blender Cosmos Jumbo',2,2,790000.00,1,'2025-01-03 15:54:38'),
	 ('Kipas Angin Maspion',2,2,144000.00,1,'2025-01-03 15:54:53');


CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `status` char(5) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.orders (user_id,status,created_at) VALUES
	 (16,'AWP','2025-01-04 08:30:18'),
	 (1,'AWP','2025-01-04 08:31:14'),
	 (16,'CNC','2025-01-04 08:37:25');


CREATE TABLE `order_items` (
  `order_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` smallint(6) NOT NULL DEFAULT 1,
  PRIMARY KEY (`order_id`,`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

INSERT INTO wetees.order_items (order_id,product_id,quantity) VALUES
	 (1,1,2),
	 (1,2,2),
	 (1,4,2),
	 (2,5,10);


