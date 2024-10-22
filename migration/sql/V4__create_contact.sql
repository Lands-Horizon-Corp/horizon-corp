CREATE TABLE IF NOT EXISTS `contacts` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `contact_number` VARCHAR(15) NOT NULL,
    description TEXT,


    PRIMARY KEY (`id`),
    `deleted_at` DATETIME(3) DEFAULT NULL,
     INDEX `idx_contact_deleted_at` (`deleted_at`),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `contacts` (`first_name`, `last_name`, `email`, `contact_number`, `description`)
VALUES
('John', 'Doe', 'john.doe@example.com', '1234567890', 'Hello! This is a test message.'),
('Jane', 'Smith', 'jane.smith@example.com', '0987654321', 'This is another test message.');