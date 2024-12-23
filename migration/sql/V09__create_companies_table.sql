CREATE TABLE IF NOT EXISTS `companies` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `address` VARCHAR(500),
    `longitude` DECIMAL(10,7),
    `latitude` DECIMAL(10,7),
    `contact_number` VARCHAR(255) NOT NULL UNIQUE,
    `owner_id` BIGINT UNSIGNED NULL,
    `media_id` BIGINT UNSIGNED NULL,
    `is_admin_verified` BOOLEAN DEFAULT FALSE,
    `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`owner_id`) REFERENCES `owners`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;