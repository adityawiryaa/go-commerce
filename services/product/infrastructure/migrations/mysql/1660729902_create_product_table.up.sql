CREATE TABLE IF NOT EXISTS `products` (
  `id` char(36) PRIMARY KEY NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` int DEFAULT NULL,
  `quantity` decimal DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `deleted_at` timestamp DEFAULT NULL
);